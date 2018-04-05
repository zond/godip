package classical

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"

	cla "github.com/zond/godip/variants/classical/common"
	ord "github.com/zond/godip/variants/classical/orders"
)

func PhaseGenerator(parser orders.Parser) func(int, godip.Season, godip.PhaseType) godip.Phase {
	return func(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
		return &phase{year, season, typ, parser}
	}
}

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return PhaseGenerator(ord.ClassicalParser)(year, season, typ)
}

type phase struct {
	year   int
	season godip.Season
	typ    godip.PhaseType
	parser orders.Parser
}

func (self *phase) String() string {
	return fmt.Sprintf("%s %d, %s", self.season, self.year, self.typ)
}

func (self *phase) Options(s godip.Validator, nation godip.Nation) (result godip.Options) {
	return s.Options(ord.ClassicalParser.Orders(), nation)
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
					if path := s.Graph().Path(srcCoast, coast, filter); path != nil {
						if !found || len(path) < result {
							result = len(path)
							found = true
						}
					}
					if path := s.Graph().Path(srcCoast, coast, nil); path != nil {
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

func (self *phase) DefaultOrder(p godip.Province) godip.Adjudicator {
	if self.typ == godip.Movement {
		return orders.Hold(p)
	}
	return nil
}

func (self *phase) PostProcess(s godip.State) (err error) {
	if self.typ == godip.Retreat {
		for prov, _ := range s.Dislodgeds() {
			s.RemoveDislodged(prov)
			s.SetResolution(prov, godip.ErrForcedDisband)
		}
		s.ClearDislodgers()
		s.ClearBounces()
		if self.season == godip.Fall {
			s.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
				if u != nil {
					if s.Graph().SC(p) != nil {
						s.SetSC(p.Super(), u.Nation)
					}
				}
				return false
			})
		}
	} else if self.typ == godip.Adjustment {
		for _, nationality := range s.Graph().Nations() {
			_, _, balance := cla.AdjustmentStatus(s, nationality)
			if balance < 0 {
				var su []godip.Province
				if su, err = SortedUnits(s, nationality); err != nil {
					return
				}
				su = su[:-balance]
				for _, prov := range su {
					godip.Logf("Removing %v due to forced disband", prov)
					s.RemoveUnit(prov)
					s.SetResolution(prov, godip.ErrForcedDisband)
				}
			}
		}
	} else if self.typ == godip.Movement {
		for prov, unit := range s.Dislodgeds() {
			hasRetreat := false
			for edge, _ := range s.Graph().Edges(prov) {
				if _, _, ok := s.Unit(edge); !ok && !s.Bounce(prov, edge) {
					if cla.HasEdge(s, unit.Type, prov, edge) {
						godip.Logf("%v can retreat to %v", prov, edge)
						hasRetreat = true
						break
					}
				}
			}
			if !hasRetreat {
				s.RemoveDislodged(prov)
				godip.Logf("Removing %v since it has no retreat", prov)
			}
		}
	}
	return
}

func (self *phase) Year() int {
	return self.year
}

func (self *phase) Season() godip.Season {
	return self.season
}

func (self *phase) Type() godip.PhaseType {
	return self.typ
}

func (self *phase) Next() godip.Phase {
	if self.typ == godip.Movement {
		return &phase{
			year:   self.year,
			season: self.season,
			typ:    godip.Retreat,
		}
	} else if self.typ == godip.Retreat {
		if self.season == godip.Spring {
			return &phase{
				year:   self.year,
				season: godip.Fall,
				typ:    godip.Movement,
			}
		} else {
			return &phase{
				year:   self.year,
				season: godip.Fall,
				typ:    godip.Adjustment,
			}
		}
	} else {
		return &phase{
			year:   self.year + 1,
			season: godip.Spring,
			typ:    godip.Movement,
		}
	}
	return nil
}
