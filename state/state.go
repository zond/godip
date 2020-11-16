package state

import (
	"fmt"
	"time"

	"github.com/zond/godip"
)

func New(graph godip.Graph, phase godip.Phase, backupRule godip.BackupRule, flags map[godip.Flag]bool, neutralOrders func(State) map[godip.Province]godip.Adjudicator) *State {
	return &State{
		graph:              graph,
		phase:              phase,
		backupRule:         backupRule,
		neutralOrders:      neutralOrders,
		orders:             make(map[godip.Province]godip.Adjudicator),
		units:              make(map[godip.Province]godip.Unit),
		dislodgeds:         make(map[godip.Province]godip.Unit),
		forceDisbands:      make(map[godip.Province]bool),
		supplyCenters:      make(map[godip.Province]godip.Nation),
		dislodgers:         make(map[godip.Province]godip.Province),
		bounces:            make(map[godip.Province]map[godip.Province]bool),
		profile:            make(map[string]time.Duration),
		profileCounts:      make(map[string]int),
		memoizedProvSlices: make(map[string][]godip.Province),
		flags:              flags,
	}
}

type movement struct {
	src            godip.Province
	dst            godip.Province
	unit           godip.Unit
	preventRetreat bool
}

func (self *movement) prepare(s *State) (err error) {
	var ok bool
	if self.unit, self.src, ok = s.Unit(self.src); !ok {
		err = fmt.Errorf("No unit at %v?", self.src)
		return
	} else {
		s.RemoveUnit(self.src)
	}
	godip.Logf("Lifted %v from %v", self.unit, self.src)
	return
}

func (self *movement) execute(s *State) (err error) {
	if dislodged, prov, ok := s.Unit(self.dst); ok {
		s.RemoveUnit(prov)
		if err = s.SetDislodged(prov, dislodged); err != nil {
			return
		}
		if self.preventRetreat {
			s.SetDislodger(self.src, prov)
		}
		godip.Logf("Dislodged %v from %v", dislodged, self.dst)
	}
	if err = s.SetUnit(self.dst, self.unit); err != nil {
		return
	}
	godip.Logf("Dropped %v in %v", self.unit, self.dst)
	return
}

type State struct {
	orders             map[godip.Province]godip.Adjudicator
	units              map[godip.Province]godip.Unit
	dislodgeds         map[godip.Province]godip.Unit
	supplyCenters      map[godip.Province]godip.Nation
	graph              godip.Graph
	phase              godip.Phase
	backupRule         godip.BackupRule
	neutralOrders      func(State) map[godip.Province]godip.Adjudicator
	resolutions        map[godip.Province]error
	dislodgers         map[godip.Province]godip.Province
	forceDisbands      map[godip.Province]bool
	movements          []*movement
	bounces            map[godip.Province]map[godip.Province]bool
	profile            map[string]time.Duration
	profileCounts      map[string]int
	memoizedProvSlices map[string][]godip.Province
	flags              map[godip.Flag]bool
}

func (self *State) Profile(a string, t time.Time) {
	self.profile[a] += time.Now().Sub(t)
	self.profileCounts[a] += 1
}

func (self *State) MemoizeProvSlice(key string, f func() []godip.Province) []godip.Province {
	old, found := self.memoizedProvSlices[key]
	if found {
		return old
	}
	neu := f()
	self.memoizedProvSlices[key] = neu
	return neu
}

func (self *State) Flags() map[godip.Flag]bool {
	return self.flags
}

func (self *State) GetProfile() (map[string]time.Duration, map[string]int) {
	return self.profile, self.profileCounts
}

func (self *State) resolver() *resolver {
	return &resolver{
		State:     self,
		guesses:   make(map[godip.Province]error),
		resolving: make(map[godip.Province]bool),
	}
}

func (self *State) Graph() godip.Graph {
	return self.graph
}

func (self *State) Options(orders []godip.Order, nation godip.Nation) (result godip.Options) {
	defer self.Profile("Options", time.Now())
	result = godip.Options{}
	for _, prov := range self.graph.Provinces() {
		for _, order := range orders {
			before := time.Now()
			opts := order.Options(self, nation, prov)
			self.Profile(string(order.DisplayType())+".Options", before)
			if len(opts) > 0 {
				provOpts, found := result[prov]
				if !found {
					provOpts = godip.Options{}
					result[prov] = provOpts
				}
				provOpts[order.DisplayType()] = opts
			}
		}
	}
	return
}

func (self *State) Find(filter godip.StateFilter) (provinces []godip.Province, orders []godip.Order, units []*godip.Unit) {
	visitedProvinces := make(map[godip.Province]bool)
	for prov, unit := range self.units {
		visitedProvinces[prov] = true
		var order godip.Order
		var ok bool
		if order, _, ok = self.Order(prov); !ok {
			order = nil
		}
		unitCopy := unit
		if filter(prov, order, &unit) {
			provinces = append(provinces, prov)
			orders = append(orders, order)
			units = append(units, &unitCopy)
		}
	}
	for prov, order := range self.orders {
		if !visitedProvinces[prov] {
			if filter(prov, order, nil) {
				provinces = append(provinces, prov)
				orders = append(orders, order)
				units = append(units, nil)
			}
		}
	}
	return
}

func (self *State) Corroborate(nat godip.Nation) []godip.Inconsistency {
	return self.Phase().Corroborate(self, nat)
}

func (self *State) Next() (err error) {
	/*
	   Sanitize orders.
	*/
	self.resolutions = make(map[godip.Province]error)
	for prov, order := range self.orders {
		if _, err := order.Validate(self); err != nil {
			self.resolutions[prov] = err
			delete(self.orders, prov)
			godip.Logf("Deleted %v due to %v", prov, err)
		}
	}

	/*
		Create orders for neutral units.
	*/
	if self.neutralOrders != nil {
		for prov, order := range self.neutralOrders(*self) {
			self.orders[prov] = order
		}
	}

	/*
	   Preprocess the phase.
	*/
	if err = self.phase.PreProcess(self.resolver()); err != nil {
		return
	}

	/*
		Add hold to units missing orders.
	*/
	for prov, _ := range self.units {
		if _, ok := self.orders[prov]; !ok {
			if _, ok := self.orders[prov.Super()]; !ok {
				if def := self.phase.DefaultOrder(prov); def != nil {
					self.orders[prov] = def
				}
			}
		}
	}

	/*
	   Adjudicate orders.
	*/
	for prov, _ := range self.orders {
		err := self.resolver().Resolve(prov)
		self.resolutions[prov] = err
	}

	/*
	   Execute orders.
	*/
	self.movements = nil
	for prov, order := range self.orders {
		if err, ok := self.resolutions[prov]; ok && err == nil {
			order.Execute(self.resolver())
		}
	}
	self.orders = make(map[godip.Province]godip.Adjudicator)

	/*
	   Execute movements.
	*/
	for _, movement := range self.movements {
		if err = movement.prepare(self); err != nil {
			return
		}
	}
	for _, movement := range self.movements {
		movement.execute(self)
	}

	/*
	   Change phase.
	*/
	if err = self.phase.PostProcess(self.resolver()); err != nil {
		return
	}
	self.phase = self.phase.Next()

	self.memoizedProvSlices = map[string][]godip.Province{}
	return
}

func (self *State) Phase() godip.Phase {
	return self.phase
}

// Bulk setters

func (self *State) SetOrders(orders map[godip.Province]godip.Adjudicator) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.orders = make(map[godip.Province]godip.Adjudicator)
	for prov, order := range orders {
		self.SetOrder(prov, order)
	}
}

func (self *State) SetUnits(units map[godip.Province]godip.Unit) (err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.units = make(map[godip.Province]godip.Unit)
	for prov, unit := range units {
		if err = self.SetUnit(prov, unit); err != nil {
			return
		}
	}
	return
}

func (self *State) SetDislodgeds(dislodgeds map[godip.Province]godip.Unit) (err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.dislodgeds = make(map[godip.Province]godip.Unit)
	for prov, unit := range dislodgeds {
		if err = self.SetDislodged(prov, unit); err != nil {
			return
		}
	}
	return
}

func (self *State) SetSupplyCenters(supplyCenters map[godip.Province]godip.Nation) *State {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.supplyCenters = supplyCenters
	return self
}

func (self *State) ClearBounces() {
	self.bounces = make(map[godip.Province]map[godip.Province]bool)
}

func (self *State) ClearDislodgers() {
	self.dislodgers = make(map[godip.Province]godip.Province)
}

func (self *State) Load(
	units map[godip.Province]godip.Unit,
	supplyCenters map[godip.Province]godip.Nation,
	dislodgeds map[godip.Province]godip.Unit,
	dislodgers map[godip.Province]godip.Province,
	bounces map[godip.Province]map[godip.Province]bool,
	orders map[godip.Province]godip.Adjudicator) *State {

	self.units, self.supplyCenters, self.dislodgeds, self.dislodgers, self.bounces, self.orders =
		units, supplyCenters, dislodgeds, dislodgers, bounces, orders

	return self
}

// Singular setters

func (self *State) SetDislodger(attacker, victim godip.Province) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.dislodgers[attacker.Super()] = victim.Super()
}

func (self *State) AddBounce(src, dst godip.Province) {
	if existing, ok := self.bounces[dst.Super()]; ok {
		existing[src.Super()] = true
	} else {
		self.bounces[dst.Super()] = map[godip.Province]bool{
			src.Super(): true,
		}
	}
}

func (self *State) SetResolution(p godip.Province, err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.resolutions[p] = err
}

func (self *State) SetSC(p godip.Province, n godip.Nation) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	self.supplyCenters[p] = n
}

func (self *State) SetDislodged(prov godip.Province, unit godip.Unit) (err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	if found, _, ok := self.Dislodged(prov); ok {
		err = fmt.Errorf("%v is already at %v", found, prov)
		return
	}
	self.dislodgeds[prov] = unit
	return
}

func (self *State) SetUnit(prov godip.Province, unit godip.Unit) (err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	if found, _, ok := self.Unit(prov); ok {
		err = fmt.Errorf("%v is already at %v", found, prov)
		return
	}
	self.units[prov] = unit
	return
}

func (self *State) SetOrder(prov godip.Province, order godip.Adjudicator) (err error) {
	self.memoizedProvSlices = map[string][]godip.Province{}

	if found, _, ok := self.Order(prov); ok {
		err = fmt.Errorf("%v is already at %v", found, prov)
		return
	}
	self.orders[prov] = order
	return
}

func (self *State) RemoveUnit(prov godip.Province) {
	if _, p, ok := self.Unit(prov); ok {
		delete(self.units, p)
	}
}

func (self *State) RemoveDislodged(prov godip.Province) {
	if _, p, ok := self.Dislodged(prov); ok {
		delete(self.dislodgeds, p)
	}
}

// Bulk getters

func (self *State) ForceDisbands() map[godip.Province]bool {
	return self.forceDisbands
}

func (self *State) Resolutions() map[godip.Province]error {
	return self.resolutions
}

func (self *State) SupplyCenters() map[godip.Province]godip.Nation {
	return self.supplyCenters
}

func (self *State) Units() map[godip.Province]godip.Unit {
	return self.units
}

func (self *State) Dislodgeds() map[godip.Province]godip.Unit {
	return self.dislodgeds
}

func (self *State) Orders() map[godip.Province]godip.Adjudicator {
	return self.orders
}

func (self *State) Dump() (
	units map[godip.Province]godip.Unit,
	supplyCenters map[godip.Province]godip.Nation,
	dislodgeds map[godip.Province]godip.Unit,
	dislodgers map[godip.Province]godip.Province,
	bounces map[godip.Province]map[godip.Province]bool,
	resolutions map[godip.Province]error) {

	return self.units,
		self.supplyCenters,
		self.dislodgeds,
		self.dislodgers,
		self.bounces,
		self.resolutions
}

// Singular getters, will search all coasts of a province

func (self *State) Bounce(src, dst godip.Province) bool {
	if sources, ok := self.bounces[dst.Super()]; ok {
		if dislodger, ok := self.dislodgers[dst.Super()]; ok {
			if len(sources) == 1 && sources[dislodger.Super()] {
				return false
			}
		}
		return true
	}
	if self.dislodgers[dst.Super()] == src.Super() {
		return true
	}
	return false
}

func (self *State) Dislodged(prov godip.Province) (u godip.Unit, p godip.Province, ok bool) {
	if u, ok = self.dislodgeds[prov]; ok {
		p = prov
		return
	}
	sup, _ := prov.Split()
	if u, ok = self.dislodgeds[sup]; ok {
		p = sup
		return
	}
	for _, name := range self.graph.Coasts(prov) {
		if u, ok = self.dislodgeds[name]; ok {
			p = name
			return
		}
	}
	return
}

func (self *State) Unit(prov godip.Province) (u godip.Unit, p godip.Province, ok bool) {
	if u, ok = self.units[prov]; ok {
		p = prov
		return
	}
	sup, _ := prov.Split()
	if u, ok = self.units[sup]; ok {
		p = sup
		return
	}
	for _, name := range self.graph.Coasts(prov) {
		if u, ok = self.units[name]; ok {
			p = name
			return
		}
	}
	return
}

func (self *State) SupplyCenter(prov godip.Province) (n godip.Nation, p godip.Province, ok bool) {
	if n, ok = self.supplyCenters[prov]; ok {
		p = prov
		return
	}
	sup, _ := prov.Split()
	if n, ok = self.supplyCenters[sup]; ok {
		p = sup
		return
	}
	for _, name := range self.graph.Coasts(prov) {
		if n, ok = self.supplyCenters[name]; ok {
			p = name
			return
		}
	}
	return
}

func (self *State) Order(prov godip.Province) (o godip.Order, p godip.Province, ok bool) {
	if o, ok = self.orders[prov]; ok {
		p = prov
		return
	}
	sup, _ := prov.Split()
	if o, ok = self.orders[sup]; ok {
		p = sup
		return
	}
	for _, name := range self.graph.Coasts(prov) {
		if o, ok = self.orders[name]; ok {
			p = name
			return
		}
	}
	return
}

// Mutators

func (self *State) Move(src, dst godip.Province, preventRetreat bool) {
	self.movements = append(self.movements, &movement{
		src:            src,
		dst:            dst,
		preventRetreat: preventRetreat,
	})
}

func (self *State) ForceDisband(prov godip.Province) {
	self.forceDisbands[prov] = true
}

func (self *State) Retreat(src, dst godip.Province) (err error) {
	if unit, prov, ok := self.Dislodged(src); !ok {
		err = fmt.Errorf("No dislodged at %v?", src)
		return
	} else {
		self.RemoveDislodged(prov)
		if err = self.SetUnit(dst, unit); err != nil {
			return
		}
		godip.Logf("Moving dislodged %v from %v to %v", unit, src, dst)
	}
	return
}
