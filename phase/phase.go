package phase

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
)

func Generator(parser orders.Parser, adjustSCs func(*Phase) bool) func(int, godip.Season, godip.PhaseType) godip.Phase {
	return func(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
		return &Phase{
			Yr:        year,
			Se:        season,
			Ty:        typ,
			Parser:    parser,
			AdjustSCs: adjustSCs,
		}
	}
}

type Phase struct {
	Yr        int
	Se        godip.Season
	Ty        godip.PhaseType
	Parser    orders.Parser
	AdjustSCs func(*Phase) bool
}

func (self *Phase) Corroborate(v godip.Validator, nat godip.Nation) []godip.Inconsistency {
	rval := []godip.Inconsistency{}
	switch self.Ty {
	case godip.Retreat:
		for prov, unit := range v.Dislodgeds() {
			if unit.Nation == nat {
				if _, _, found := v.Order(prov); !found {
					rval = append(rval, godip.Inconsistency{
						Province: prov.Super(),
						Errors:   []error{godip.InconsistencyMissingOrder},
					})
				}
			}
		}
	case godip.Adjustment:
		foundBuilds := 0
		foundDisbands := 0
		for _, ord := range v.Orders() {
			owner, err := ord.Validate(v)
			if err == nil && owner == nat {
				if ord.Type() == godip.Build {
					foundBuilds += 1
				} else if ord.Type() == godip.Disband {
					foundDisbands += 1
				}
			}
		}
		if balance := self.allowedBuildBalance(v)[nat]; balance >= 0 {
			if foundBuilds != balance {
				rval = append(rval, godip.Inconsistency{
					Errors: []error{godip.InconsistencyOrderTypeCount{
						OrderType: godip.Build,
						Found:     foundBuilds,
						Want:      balance,
					}},
				})
			}
			if foundDisbands != 0 {
				rval = append(rval, godip.Inconsistency{
					Errors: []error{godip.InconsistencyOrderTypeCount{
						OrderType: godip.Disband,
						Found:     foundDisbands,
						Want:      0,
					}},
				})
			}
		} else {
			if foundDisbands != -balance {
				rval = append(rval, godip.Inconsistency{
					Errors: []error{godip.InconsistencyOrderTypeCount{
						OrderType: godip.Disband,
						Found:     foundDisbands,
						Want:      -balance,
					}},
				})
			}
			if foundBuilds != 0 {
				rval = append(rval, godip.Inconsistency{
					Errors: []error{godip.InconsistencyOrderTypeCount{
						OrderType: godip.Build,
						Found:     foundBuilds,
						Want:      0,
					}},
				})
			}
		}
	case godip.Movement:
		for prov, unit := range v.Units() {
			if unit.Nation == nat {
				if _, _, found := v.Order(prov); !found {
					rval = append(rval, godip.Inconsistency{
						Province: prov.Super(),
						Errors:   []error{godip.InconsistencyMissingOrder},
					})
				}
			}
		}
	}
	for prov, ord := range v.Orders() {
		owner, err := ord.Validate(v)
		if nat == owner {
			if err != nil {
				rval = append(rval, godip.Inconsistency{
					Province: prov.Super(),
					Errors:   []error{err},
				})
			} else {
				if errs := ord.Corroborate(v); len(errs) > 0 {
					rval = append(rval, godip.Inconsistency{
						Province: prov.Super(),
						Errors:   errs,
					})
				}
			}
		}
	}
	return rval
}

func (self *Phase) String() string {
	return fmt.Sprintf("%s %d, %s", self.Se, self.Yr, self.Ty)
}

func (self *Phase) Options(s godip.Validator, nation godip.Nation) godip.Options {
	return s.Options(self.Parser.Orders(), nation)
}

// Determine if the given nation is potentially allowed to build in a particular province.
func isEligiableForBuild(s godip.Validator, sc godip.Province, nat godip.Nation) bool {
	if s.Flags()[godip.Anywhere] {
		// Build anywhere rules.
		return true
	}
	originalOwner := s.Graph().SC(sc)
	if originalOwner == nil {
		// Not a home center.
		return false
	}
	if s.Flags()[godip.AnyHomeCenter] {
		// A home center in a build any-home center game.
		return true
	}
	// Standard rules.
	return originalOwner != nil && *originalOwner == nat
}

// Returns number of allowed (after considering free owned SCs where builds are allowed considering the validator
// flags) builds/needed disbands per nation still in the game.
func (self *Phase) allowedBuildBalance(s godip.Validator) map[godip.Nation]int {
	unitsPerNat := map[godip.Nation]int{}
	scsPerNat := map[godip.Nation]int{}
	nats := map[godip.Nation]bool{}
	freePerNat := map[godip.Nation]int{}

	for _, unit := range s.Units() {
		unitsPerNat[unit.Nation] += 1
		nats[unit.Nation] = true
	}
	for sc, nat := range s.SupplyCenters() {
		scsPerNat[nat] += 1
		nats[nat] = true
		if _, _, found := s.Unit(sc); !found {
			if isEligiableForBuild(s, sc, nat) {
				freePerNat[nat] += 1
			}
		}
	}

	result := map[godip.Nation]int{}
	for _, nat := range s.Graph().Nations() {
		delta := scsPerNat[nat] - unitsPerNat[nat]
		if delta > freePerNat[nat] {
			delta = freePerNat[nat]
		}
		result[nat] = delta
	}
	return result
}

func (self *Phase) Messages(s godip.Validator, nation godip.Nation) []string {
	messages := []string{}
	if self.Ty == godip.Adjustment {
		for nat, delta := range self.allowedBuildBalance(s) {
			if nat == nation {
				if delta < 0 {
					messages = append(messages, fmt.Sprintf("MustDisband:%v", -delta))
				} else {
					messages = append(messages, fmt.Sprintf("MayBuild:%v", delta))
				}
			} else {
				if delta < 0 {
					messages = append(messages, fmt.Sprintf("OtherMustDisband:%v:%v", nat, -delta))
				} else {
					messages = append(messages, fmt.Sprintf("OtherMayBuild:%v:%v", nat, delta))
				}
			}
		}
	}
	return messages
}

func shortestDistance(s godip.State, src godip.Province, dst []godip.Province) (result int, err error) {
	var unit godip.Unit
	var ok bool
	unit, src, ok = s.Unit(src)
	if !ok {
		err = fmt.Errorf("No unit at %v", src)
		return
	}
	var filter godip.PathFilter
	found := false
	for _, destination := range dst {
		if unit.Type == godip.Fleet {
			filter = func(p godip.Province, edgeFlags, nodeFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) bool {
				return edgeFlags[godip.Sea] && nodeFlags[godip.Sea]
			}
		} else {
			filter = func(p godip.Province, edgeFlags, nodeFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) bool {
				if p.Super() == destination.Super() {
					return true
				}
				u, _, ok := s.Unit(p)
				return (edgeFlags[godip.Land] && nodeFlags[godip.Land]) || (ok && !nodeFlags[godip.Land] && u.Nation == unit.Nation && u.Type == godip.Fleet)
			}
		}
		for _, coast := range s.Graph().Coasts(destination) {
			for _, srcCoast := range s.Graph().Coasts(src) {
				if srcCoast == destination {
					result = 0
					found = true
				} else {
					if path := s.Graph().Path(srcCoast, coast, false, filter); path != nil {
						if !found || len(path) < result {
							result = len(path)
							found = true
						}
					}
					if path := s.Graph().Path(srcCoast, coast, false, nil); path != nil {
						if !found || len(path) < result {
							result = len(path)
							found = true
						}
					}
				}
			}
		}
	}
	return
}

type remoteUnitSlice struct {
	provinces []godip.Province
	distances map[godip.Province]int
	units     map[godip.Province]godip.Unit
}

func (self remoteUnitSlice) String() string {
	var l []string
	for _, prov := range self.provinces {
		l = append(l, fmt.Sprintf("%v:%v", prov, self.distances[prov]))
	}
	return strings.Join(l, ", ")
}

func (self remoteUnitSlice) Len() int {
	return len(self.provinces)
}

func (self remoteUnitSlice) Swap(i, j int) {
	self.provinces[i], self.provinces[j] = self.provinces[j], self.provinces[i]
}

func (self remoteUnitSlice) Less(i, j int) bool {
	if self.distances[self.provinces[i]] == self.distances[self.provinces[j]] {
		u1 := self.units[self.provinces[i]]
		u2 := self.units[self.provinces[j]]
		if u1.Type == godip.Fleet && u2.Type == godip.Army {
			return true
		}
		if u2.Type == godip.Fleet && u1.Type == godip.Army {
			return false
		}
		return bytes.Compare([]byte(self.provinces[i]), []byte(self.provinces[j])) < 0
	}
	return self.distances[self.provinces[i]] > self.distances[self.provinces[j]]
}

func SortedUnits(s godip.State, n godip.Nation) (result []godip.Province, err error) {
	provs := remoteUnitSlice{
		distances: make(map[godip.Province]int),
		units:     make(map[godip.Province]godip.Unit),
	}
	provs.provinces, _, _ = s.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		if u != nil && u.Nation == n {
			if provs.distances[p], err = shortestDistance(s, p, s.Graph().SCs(n)); err != nil {
				return false
			}
			provs.units[p] = *u
			return true
		}
		return false
	})
	if err != nil {
		return
	}
	sort.Sort(provs)
	godip.Logf("Sorted units for %v is %v", n, provs)
	result = provs.provinces
	return
}

func (self *Phase) DefaultOrder(p godip.Province) godip.Adjudicator {
	if self.Ty == godip.Movement {
		return orders.Hold(p)
	}
	return nil
}

func (self *Phase) PreProcess(s godip.State) (err error) {
	return nil
}

func (self *Phase) PostProcess(s godip.State) (err error) {
	if self.Ty == godip.Retreat {
		for prov, _ := range s.Dislodgeds() {
			s.RemoveDislodged(prov)
			s.ForceDisband(prov)
			godip.Logf("Removing %v since it didn't retreat", prov)
		}
		s.ClearDislodgers()
		s.ClearBounces()
	} else if self.Ty == godip.Adjustment {
		for _, nationality := range s.Graph().Nations() {
			_, _, balance := orders.AdjustmentStatus(s, nationality)
			if balance < 0 {
				var su []godip.Province
				if su, err = SortedUnits(s, nationality); err != nil {
					return
				}
				su = su[:-balance]
				for _, prov := range su {
					s.RemoveUnit(prov)
					s.ForceDisband(prov)
					godip.Logf("Removing %v since it wasn't disbanded by order", prov)
				}
			}
		}
	} else if self.Ty == godip.Movement {
		for prov, unit := range s.Dislodgeds() {
			hasRetreat := false
			for edge, _ := range s.Graph().Edges(prov, false) {
				if _, _, ok := s.Unit(edge); !ok && !s.Bounce(prov, edge) {
					if orders.HasEdge(s, unit.Type, prov, edge) {
						godip.Logf("%v can retreat to %v", prov, edge)
						hasRetreat = true
						break
					}
				}
			}
			if !hasRetreat {
				s.RemoveDislodged(prov)
				s.ForceDisband(prov)
				godip.Logf("Removing %v since it has no retreat", prov)
			}
		}
	}
	if self.AdjustSCs(self) {
		s.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
			if u != nil {
				if s.Graph().SC(p) != nil {
					godip.Logf("%v now belongs to %v", p.Super(), u.Nation)
					s.SetSC(p.Super(), u.Nation)
				}
			}
			return false
		})
	}
	return
}

func (self *Phase) Year() int {
	return self.Yr
}

func (self *Phase) Season() godip.Season {
	return self.Se
}

func (self *Phase) Type() godip.PhaseType {
	return self.Ty
}

func (self *Phase) Next() godip.Phase {
	if self.Ty == godip.Movement {
		return &Phase{
			Yr:        self.Yr,
			Se:        self.Se,
			Ty:        godip.Retreat,
			Parser:    self.Parser,
			AdjustSCs: self.AdjustSCs,
		}
	} else if self.Ty == godip.Retreat {
		if self.Se == godip.Spring {
			return &Phase{
				Yr:        self.Yr,
				Se:        godip.Fall,
				Ty:        godip.Movement,
				Parser:    self.Parser,
				AdjustSCs: self.AdjustSCs,
			}
		} else {
			return &Phase{
				Yr:        self.Yr,
				Se:        godip.Fall,
				Ty:        godip.Adjustment,
				Parser:    self.Parser,
				AdjustSCs: self.AdjustSCs,
			}
		}
	} else {
		return &Phase{
			Yr:        self.Yr + 1,
			Se:        godip.Spring,
			Ty:        godip.Movement,
			Parser:    self.Parser,
			AdjustSCs: self.AdjustSCs,
		}
	}
	return nil
}
