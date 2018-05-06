package hundred

import (
	"fmt"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/variants/classical"
)

const (
	YearSeason godip.Season = "Year"
)

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	if season != YearSeason {
		fmt.Errorf("Warning - Hundred only supports YearSeason, but got {}", season)
	}
	return &phase{year, typ}
}

type phase struct {
	year int
	typ  godip.PhaseType
}

func (self *phase) String() string {
	return fmt.Sprintf("%s %d, %s", YearSeason, self.year, self.typ)
}

func (self *phase) Options(s godip.Validator, nation godip.Nation) (result godip.Options) {
	return s.Options(BuildAnywhereParser.Orders(), nation)
}

func (self *phase) DefaultOrder(p godip.Province) godip.Adjudicator {
	if self.typ == godip.Movement {
		return orders.Hold(p)
	}
	return nil
}

func (self *phase) PreProcess(s godip.State) (err error) {
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
		if self.year%10 == 0 {
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
			_, _, balance := orders.AdjustmentStatus(s, nationality)
			if balance < 0 {
				var su []godip.Province
				if su, err = classical.SortedUnits(s, nationality); err != nil {
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
	return
}

func (self *phase) Year() int {
	return self.year
}

func (self *phase) Season() godip.Season {
	return YearSeason
}

func (self *phase) Type() godip.PhaseType {
	return self.typ
}

func (self *phase) Next() godip.Phase {
	if self.typ == godip.Movement {
		return &phase{
			year: self.year,
			typ:  godip.Retreat,
		}
	} else if self.typ == godip.Retreat {
		if self.year%10 == 5 {
			return &phase{
				year: self.year + 5,
				typ:  godip.Movement,
			}
		} else {
			return &phase{
				year: self.year,
				typ:  godip.Adjustment,
			}
		}
	} else {
		return &phase{
			year: self.year + 5,
			typ:  godip.Movement,
		}
	}
	return nil
}
