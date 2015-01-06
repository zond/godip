package orders

import (
	"fmt"
	"time"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

func init() {
	generators = append(generators, func() dip.Order { return &support{} })
}

func SupportHold(prov, target dip.Province) *support {
	return &support{
		targets: []dip.Province{prov, target},
	}
}

func SupportMove(prov, from, to dip.Province) *support {
	return &support{
		targets: []dip.Province{prov, from, to},
	}
}

type support struct {
	targets []dip.Province
}

func (self *support) Flags() map[dip.Flag]bool {
	return nil
}

func (self *support) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], cla.Support, self.targets[1:])
}

func (self *support) At() time.Time {
	return time.Now()
}

func (self *support) Type() dip.OrderType {
	return cla.Support
}

func (self *support) DisplayType() dip.OrderType {
	return cla.Support
}

func (self *support) Targets() []dip.Province {
	return self.targets
}

func (self *support) Adjudicate(r dip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	if breaks, _, _ := r.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
		if o != nil && // is an order
			u != nil && // is a unit
			o.Type() == cla.Move && // move
			o.Targets()[1].Super() == self.targets[0].Super() && // against us
			(len(self.targets) == 2 || o.Targets()[0].Super() != self.targets[2].Super()) && // not from something we support attacking
			u.Nation != unit.Nation { // not from ourselves

			_, err := cla.AnyMovePossible(r, u.Type, o.Targets()[0], o.Targets()[1], u.Type == cla.Army, true, true) // and legal move counting convoy success
			return err == nil
		}
		return false
	}); len(breaks) > 0 {
		dip.Logf("%v: broken by: %v", self, breaks)
		return cla.ErrSupportBroken{breaks[0]}
	}

	if dislodgers, _, _ := r.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
		return o != nil && // is an order
			u != nil && // is a unit
			o.Type() == cla.Move && // move
			o.Targets()[1].Super() == self.targets[0].Super() && // against us
			u.Nation != unit.Nation && // not from ourselves
			r.Resolve(p) == nil // and it succeeded
	}); len(dislodgers) > 0 {
		dip.Logf("%v: dislodged by: %v", self, dislodgers)
		return cla.ErrSupportBroken{dislodgers[0]}
	}
	return nil
}

func (self *support) Options(v dip.Validator, nation dip.Nation, src dip.Province) (result dip.Options) {
	if v.Phase().Type() == cla.Movement {
		if v.Graph().Has(src) {
			if supporter, actualSrc, ok := v.Unit(src); ok {
				if supporter.Nation == nation {
					for _, supportable := range cla.PossibleMoves(v, src, false, false) {
						if _, supporteeSrc, ok := v.Unit(supportable); ok {
							if result == nil {
								result = dip.Options{}
							}
							if result[dip.SrcProvince(actualSrc)] == nil {
								result[dip.SrcProvince(actualSrc)] = dip.Options{}
							}
							opt, f := result[dip.SrcProvince(actualSrc)][supporteeSrc.Super()]
							if !f {
								opt = dip.Options{}
								result[dip.SrcProvince(actualSrc)][supporteeSrc.Super()] = opt
							}
							opt[supporteeSrc.Super()] = nil
						}
						for mvSrc, unit := range v.Units() {
							if mvSrc != actualSrc {
								if mvDst, err := cla.AnyMovePossible(v, unit.Type, mvSrc, supportable, true, true, false); err == nil {
									if result == nil {
										result = dip.Options{}
									}
									if result[dip.SrcProvince(actualSrc)] == nil {
										result[dip.SrcProvince(actualSrc)] = dip.Options{}
									}
									opt, f := result[dip.SrcProvince(actualSrc)][mvSrc.Super()]
									if !f {
										opt = dip.Options{}
										result[dip.SrcProvince(actualSrc)][mvSrc.Super()] = opt
									}
									opt[mvDst.Super()] = nil
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

func (self *support) Validate(v dip.Validator) error {
	if v.Phase().Type() != cla.Movement {
		return cla.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return cla.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return cla.ErrInvalidTarget
	}
	var ok bool
	var unit, supported dip.Unit
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return cla.ErrMissingUnit
	}
	if supported, self.targets[1], ok = v.Unit(self.targets[1]); !ok {
		return cla.ErrMissingSupportUnit
	}
	if len(self.targets) == 2 {
		if err := cla.AnySupportPossible(v, unit.Type, self.targets[0], self.targets[1]); err != nil {
			return cla.ErrIllegalSupportPosition
		}
	} else {
		if !v.Graph().Has(self.targets[2]) {
			return cla.ErrInvalidTarget
		}
		if err := cla.AnySupportPossible(v, unit.Type, self.targets[0], self.targets[2]); err != nil {
			return cla.ErrIllegalSupportDestination
		}

		if _, err := cla.AnyMovePossible(v, supported.Type, self.targets[1], self.targets[2], true, true, false); err != nil {
			return cla.ErrIllegalSupportMove
		}
	}
	return nil
}

func (self *support) Execute(state dip.State) {
}
