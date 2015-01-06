package classical

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	cla "github.com/zond/godip/classical/common"
	"github.com/zond/godip/classical/orders"
	dip "github.com/zond/godip/common"
)

func Phase(year int, season dip.Season, typ dip.PhaseType) dip.Phase {
	return &phase{year, season, typ}
}

type phase struct {
	year   int
	season dip.Season
	typ    dip.PhaseType
}

func (self *phase) Options(s dip.Validator, nation dip.Nation) (result dip.Options) {
	return s.Options(orders.Orders(), nation)
}

func (self *phase) Winner(s dip.Validator) *dip.Nation {
	sizes := map[dip.Nation]int{}
	for _, nat := range s.SupplyCenters() {
		sizes[nat]++
	}
	var biggest *dip.Nation
	var bigSize int
	for nat, size := range sizes {
		if biggest == nil || size > bigSize {
			natCpy := nat
			biggest = &natCpy
			bigSize = size
		}
	}
	if bigSize > 17 {
		return biggest
	}
	return nil
}

func (self *phase) shortestDistance(s dip.State, src dip.Province, dst []dip.Province) (result int, err error) {
	var unit dip.Unit
	var ok bool
	unit, src, ok = s.Unit(src)
	if !ok {
		err = fmt.Errorf("No unit at %v", src)
		return
	}
	var filter dip.PathFilter
	found := false
	for _, destination := range dst {
		if unit.Type == cla.Fleet {
			filter = func(p dip.Province, edgeFlags, nodeFlags map[dip.Flag]bool, sc *dip.Nation) bool {
				return edgeFlags[cla.Sea] && nodeFlags[cla.Sea]
			}
		} else {
			filter = func(p dip.Province, edgeFlags, nodeFlags map[dip.Flag]bool, sc *dip.Nation) bool {
				if p.Super() == destination.Super() {
					return true
				}
				u, _, ok := s.Unit(p)
				return (edgeFlags[cla.Land] && nodeFlags[cla.Land]) || (ok && !nodeFlags[cla.Land] && u.Nation == unit.Nation && u.Type == cla.Fleet)
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
	provinces []dip.Province
	distances map[dip.Province]int
	units     map[dip.Province]dip.Unit
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
		if u1.Type == cla.Fleet && u2.Type == cla.Army {
			return true
		}
		if u2.Type == cla.Fleet && u1.Type == cla.Army {
			return false
		}
		return bytes.Compare([]byte(self.provinces[i]), []byte(self.provinces[j])) < 0
	}
	return self.distances[self.provinces[i]] > self.distances[self.provinces[j]]
}

func (self *phase) sortedUnits(s dip.State, n dip.Nation) (result []dip.Province, err error) {
	provs := remoteUnitSlice{
		distances: make(map[dip.Province]int),
		units:     make(map[dip.Province]dip.Unit),
	}
	provs.provinces, _, _ = s.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
		if u != nil && u.Nation == n {
			if provs.distances[p], err = self.shortestDistance(s, p, s.Graph().SCs(n)); err != nil {
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
	dip.Logf("Sorted units for %v is %v", n, provs)
	result = provs.provinces
	return
}

func (self *phase) DefaultOrder(p dip.Province) dip.Adjudicator {
	if self.typ == cla.Movement {
		return orders.Hold(p)
	}
	return nil
}

func (self *phase) PostProcess(s dip.State) (err error) {
	if self.typ == cla.Retreat {
		for prov, _ := range s.Dislodgeds() {
			s.RemoveDislodged(prov)
			s.SetResolution(prov, cla.ErrForcedDisband)
		}
		s.ClearDislodgers()
		s.ClearBounces()
		if self.season == cla.Fall {
			s.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
				if u != nil {
					if s.Graph().SC(p) != nil {
						s.SetSC(p.Super(), u.Nation)
					}
				}
				return false
			})
		}
	} else if self.typ == cla.Adjustment {
		for _, nationality := range cla.Nations {
			_, _, balance := cla.AdjustmentStatus(s, nationality)
			if balance < 0 {
				var su []dip.Province
				if su, err = self.sortedUnits(s, nationality); err != nil {
					return
				}
				su = su[:-balance]
				for _, prov := range su {
					dip.Logf("Removing %v due to forced disband", prov)
					s.RemoveUnit(prov)
					s.SetResolution(prov, cla.ErrForcedDisband)
				}
			}
		}
	} else if self.typ == cla.Movement {
		for prov, unit := range s.Dislodgeds() {
			hasRetreat := false
			for edge, _ := range s.Graph().Edges(prov) {
				if _, _, ok := s.Unit(edge); !ok && !s.Bounce(prov, edge) {
					if cla.HasEdge(s, unit.Type, prov, edge) {
						dip.Logf("%v can retreat to %v", prov, edge)
						hasRetreat = true
						break
					}
				}
			}
			if !hasRetreat {
				s.RemoveDislodged(prov)
				dip.Logf("Removing %v since it has no retreat", prov)
			}
		}
	}
	return
}

func (self *phase) Year() int {
	return self.year
}

func (self *phase) Season() dip.Season {
	return self.season
}

func (self *phase) Type() dip.PhaseType {
	return self.typ
}

func (self *phase) Prev() dip.Phase {
	if self.typ == cla.Retreat {
		return &phase{
			year:   self.year,
			season: self.season,
			typ:    cla.Movement,
		}
	} else if self.typ == cla.Movement {
		if self.season == cla.Spring {
			if self.year == 1901 {
				return nil
			}
			return &phase{
				year:   self.year - 1,
				season: cla.Fall,
				typ:    cla.Adjustment,
			}
		} else {
			return &phase{
				year:   self.year,
				season: cla.Spring,
				typ:    cla.Retreat,
			}
		}
	} else {
		return &phase{
			year:   self.year,
			season: cla.Fall,
			typ:    cla.Retreat,
		}
	}
	return nil
}

func (self *phase) Next() dip.Phase {
	if self.typ == cla.Movement {
		return &phase{
			year:   self.year,
			season: self.season,
			typ:    cla.Retreat,
		}
	} else if self.typ == cla.Retreat {
		if self.season == cla.Spring {
			return &phase{
				year:   self.year,
				season: cla.Fall,
				typ:    cla.Movement,
			}
		} else {
			return &phase{
				year:   self.year,
				season: cla.Fall,
				typ:    cla.Adjustment,
			}
		}
	} else {
		return &phase{
			year:   self.year + 1,
			season: cla.Spring,
			typ:    cla.Movement,
		}
	}
	return nil
}
