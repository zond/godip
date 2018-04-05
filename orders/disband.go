package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"
	cla "github.com/zond/godip/variants/classical/common"
)

var DisbandOrder = &disband{}

func Disband(source godip.Province, at time.Time) *disband {
	return &disband{
		targets: []godip.Province{source},
		at:      at,
	}
}

type disband struct {
	targets []godip.Province
	at      time.Time
}

func (self *disband) String() string {
	return fmt.Sprintf("%v %v", self.targets[0], cla.Disband)
}

func (self *disband) Type() godip.OrderType {
	return cla.Disband
}

func (self *disband) DisplayType() godip.OrderType {
	return cla.Disband
}

func (self *disband) Flags() map[godip.Flag]bool {
	return nil
}

func (self *disband) Targets() []godip.Province {
	return self.targets
}

func (self *disband) At() time.Time {
	return self.at
}

func (self *disband) adjudicateBuildPhase(r godip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	_, disbands, _ := cla.AdjustmentStatus(r, unit.Nation)
	if len(disbands) == 0 || self.at.After(disbands[len(disbands)-1].At()) {
		return godip.ErrIllegalDisband
	}
	return nil
}

func (self *disband) adjudicateRetreatPhase(r godip.Resolver) error {
	return nil
}

func (self *disband) Adjudicate(r godip.Resolver) error {
	if r.Phase().Type() == cla.Adjustment {
		return self.adjudicateBuildPhase(r)
	}
	return self.adjudicateRetreatPhase(r)
}

func (self *disband) validateRetreatPhase(v godip.Validator) (godip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidTarget
	}
	var ok bool
	var dislodged godip.Unit
	dislodged, self.targets[0], ok = v.Dislodged(self.targets[0])
	if !ok {
		return "", godip.ErrMissingUnit
	}
	return dislodged.Nation, nil
}

func (self *disband) validateBuildPhase(v godip.Validator) (godip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidTarget
	}
	var unit godip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return "", godip.ErrMissingUnit
	}
	if _, _, balance := cla.AdjustmentStatus(v, unit.Nation); balance > -1 {
		return "", godip.ErrMissingDeficit
	}
	return unit.Nation, nil
}

func (self *disband) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 2 {
			result = Disband(godip.Province(bits[0]), time.Now())
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *disband) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() == cla.Adjustment {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
				if unit.Nation == nation {
					if _, _, balance := cla.AdjustmentStatus(v, unit.Nation); balance < 0 {
						result = godip.Options{
							godip.SrcProvince(actualSrc): nil,
						}
					}
				}
			}
		}
	} else if v.Phase().Type() == cla.Retreat {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Dislodged(src); ok {
				if unit.Nation == nation {
					result = godip.Options{
						godip.SrcProvince(actualSrc): nil,
					}
				}
			}
		}
	}
	return
}

func (self *disband) Validate(v godip.Validator) (godip.Nation, error) {
	if v.Phase().Type() == cla.Adjustment {
		return self.validateBuildPhase(v)
	} else if v.Phase().Type() == cla.Retreat {
		return self.validateRetreatPhase(v)
	}
	return "", godip.ErrInvalidPhase
}

func (self *disband) Execute(state godip.State) {
	if state.Phase().Type() == cla.Adjustment {
		state.RemoveUnit(self.targets[0])
	} else {
		state.RemoveDislodged(self.targets[0])
	}
}
