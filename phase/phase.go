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

func (self *Phase) String() string {
	return fmt.Sprintf("%s %d, %s", self.Se, self.Yr, self.Ty)
}

func (self *Phase) Options(s godip.Validator, nation godip.Nation) (result godip.Options) {
	return s.Options(self.Parser.Orders(), nation)
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
			s.SetResolution(prov, godip.ErrForcedDisband)
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
					godip.Logf("Removing %v due to forced disband", prov)
					s.RemoveUnit(prov)
					s.SetResolution(prov, godip.ErrForcedDisband)
				}
			}
		}
	} else if self.Ty == godip.Movement {
		for prov, unit := range s.Dislodgeds() {
			hasRetreat := false
			for edge, _ := range s.Graph().Edges(prov) {
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
