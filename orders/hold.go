package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"

	cla "github.com/zond/godip/variants/classical/common"
)

var HoldOrder = &hold{}

func Hold(source godip.Province) *hold {
	return &hold{
		targets: []godip.Province{source},
	}
}

type hold struct {
	targets []godip.Province
}

func (self *hold) String() string {
	return fmt.Sprintf("%v %v", self.targets[0], cla.Hold)
}

func (self *hold) Flags() map[godip.Flag]bool {
	return nil
}

func (self *hold) Type() godip.OrderType {
	return cla.Hold
}

func (self *hold) DisplayType() godip.OrderType {
	return cla.Hold
}

func (self *hold) Targets() []godip.Province {
	return self.targets
}

func (self *hold) At() time.Time {
	return time.Now()
}

func (self *hold) Adjudicate(r godip.Resolver) error {
	return nil
}

func (self *hold) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 2 {
			result = Hold(godip.Province(bits[0]))
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *hold) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() == cla.Movement {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
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

func (self *hold) Validate(v godip.Validator) (godip.Nation, error) {
	if v.Phase().Type() != cla.Movement {
		return "", godip.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidTarget
	}
	var ok bool
	var unit godip.Unit
	unit, self.targets[0], ok = v.Unit(self.targets[0])
	if !ok {
		return "", godip.ErrMissingUnit
	}
	return unit.Nation, nil
}

func (self *hold) Execute(state godip.State) {
}
