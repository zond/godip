package hundred

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

const (
	YearSeason dip.Season = "Year"
)

func Phase(year int, typ dip.PhaseType) dip.Phase {
	return &phase{year, typ}
}

type phase struct {
	year int
	typ  dip.PhaseType
}

func (self *phase) String() string {
	return fmt.Sprintf("%s %d, %s", YearSeason, self.year, self.typ)
}

func (self *phase) Options(s dip.Validator, nation dip.Nation) (result dip.Options) {
	return s.Options(orders.ClassicalParser.Orders(), nation)
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
		if self.year%10 == 0 {
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
		for _, nationality := range s.Graph().Nations() {
			_, _, balance := cla.AdjustmentStatus(s, nationality)
			if balance < 0 {
				var su []dip.Province
				if su, err = classical.SortedUnits(s, nationality); err != nil {
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
	return YearSeason
}

func (self *phase) Type() dip.PhaseType {
	return self.typ
}

func (self *phase) Next() dip.Phase {
	if self.typ == cla.Movement {
		return &phase{
			year: self.year,
			typ:  cla.Retreat,
		}
	} else if self.typ == cla.Retreat {
		if self.year%10 == 5 {
			return &phase{
				year: self.year + 5,
				typ:  cla.Movement,
			}
		} else {
			return &phase{
				year: self.year,
				typ:  cla.Adjustment,
			}
		}
	} else {
		return &phase{
			year: self.year + 5,
			typ:  cla.Movement,
		}
	}
	return nil
}
