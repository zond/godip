package orders

import (
	"fmt"
	"time"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

var DisbandGenerator func() dip.Order = func() dip.Order { return &disband{} }

func Disband(source dip.Province, at time.Time) *disband {
	return &disband{
		targets: []dip.Province{source},
		at:      at,
	}
}

type disband struct {
	targets []dip.Province
	at      time.Time
}

func (self *disband) String() string {
	return fmt.Sprintf("%v %v", self.targets[0], cla.Disband)
}

func (self *disband) Type() dip.OrderType {
	return cla.Disband
}

func (self *disband) DisplayType() dip.OrderType {
	return cla.Disband
}

func (self *disband) Flags() map[dip.Flag]bool {
	return nil
}

func (self *disband) Targets() []dip.Province {
	return self.targets
}

func (self *disband) At() time.Time {
	return self.at
}

func (self *disband) adjudicateBuildPhase(r dip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	_, disbands, _ := cla.AdjustmentStatus(r, unit.Nation)
	if len(disbands) == 0 || self.at.After(disbands[len(disbands)-1].At()) {
		return cla.ErrIllegalDisband
	}
	return nil
}

func (self *disband) adjudicateRetreatPhase(r dip.Resolver) error {
	return nil
}

func (self *disband) Adjudicate(r dip.Resolver) error {
	if r.Phase().Type() == cla.Adjustment {
		return self.adjudicateBuildPhase(r)
	}
	return self.adjudicateRetreatPhase(r)
}

func (self *disband) validateRetreatPhase(v dip.Validator) (dip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", cla.ErrInvalidTarget
	}
	var ok bool
	var dislodged dip.Unit
	dislodged, self.targets[0], ok = v.Dislodged(self.targets[0])
	if !ok {
		return "", cla.ErrMissingUnit
	}
	return dislodged.Nation, nil
}

func (self *disband) validateBuildPhase(v dip.Validator) (dip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", cla.ErrInvalidTarget
	}
	var unit dip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return "", cla.ErrMissingUnit
	}
	if _, _, balance := cla.AdjustmentStatus(v, unit.Nation); balance > -1 {
		return "", cla.ErrMissingDeficit
	}
	return unit.Nation, nil
}

func (self *disband) Parse(bits []string) (dip.Adjudicator, error) {
	var result dip.Adjudicator
	var err error
	if len(bits) > 1 && dip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 2 {
			result = Disband(dip.Province(bits[0]), time.Now())
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *disband) Options(v dip.Validator, nation dip.Nation, src dip.Province) (result dip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() == cla.Adjustment {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
				if unit.Nation == nation {
					if _, _, balance := cla.AdjustmentStatus(v, unit.Nation); balance < 0 {
						result = dip.Options{
							dip.SrcProvince(actualSrc): nil,
						}
					}
				}
			}
		}
	} else if v.Phase().Type() == cla.Retreat {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Dislodged(src); ok {
				if unit.Nation == nation {
					result = dip.Options{
						dip.SrcProvince(actualSrc): nil,
					}
				}
			}
		}
	}
	return
}

func (self *disband) Validate(v dip.Validator) (dip.Nation, error) {
	if v.Phase().Type() == cla.Adjustment {
		return self.validateBuildPhase(v)
	} else if v.Phase().Type() == cla.Retreat {
		return self.validateRetreatPhase(v)
	}
	return "", cla.ErrInvalidPhase
}

func (self *disband) Execute(state dip.State) {
	if state.Phase().Type() == cla.Adjustment {
		state.RemoveUnit(self.targets[0])
	} else {
		state.RemoveDislodged(self.targets[0])
	}
}
