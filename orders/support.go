package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"
)

var SupportOrder = &support{}

func SupportHold(prov, target godip.Province) *support {
	return &support{
		targets: []godip.Province{prov, target},
	}
}

func SupportMove(prov, from, to godip.Province) *support {
	return &support{
		targets: []godip.Province{prov, from, to},
	}
}

type support struct {
	targets []godip.Province
}

func (self *support) Corroborate(v godip.Validator) []error {
	unit, _, _ := v.Unit(self.targets[0])
	me := unit.Nation
	supportee, _, _ := v.Unit(self.targets[1])
	if supportee.Nation == me {
		potentialInconsistencies := []error{godip.InconsistencyMismatchedSupporter{
			Supportee: self.targets[1].Super(),
		}}
		supporteeOrd, _, found := v.Order(self.targets[1])
		if found && supporteeOrd.Type() == godip.Move {
			if len(self.targets) != 3 || supporteeOrd.Targets()[1] != self.targets[2] {
				return potentialInconsistencies
			}
		} else if len(self.targets) > 2 {
			return potentialInconsistencies
		}
	}
	return nil
}

func (self *support) Flags() map[godip.Flag]bool {
	return nil
}

func (self *support) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], godip.Support, self.targets[1:])
}

func (self *support) At() time.Time {
	return time.Now()
}

func (self *support) Type() godip.OrderType {
	return godip.Support
}

func (self *support) DisplayType() godip.OrderType {
	return godip.Support
}

func (self *support) Targets() []godip.Province {
	return self.targets
}

func (self *support) Adjudicate(r godip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	if breaks, _, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		if o != nil && // is an order
			u != nil && // is a unit
			o.Type() == godip.Move && // move
			o.Targets()[1].Super() == self.targets[0].Super() && // against us
			(len(self.targets) == 2 || o.Targets()[0].Super() != self.targets[2].Super()) && // not from something we support attacking
			u.Nation != unit.Nation { // not from ourselves

			_, err := AnyMovePossible(r, u.Type, o.Targets()[0], o.Targets()[1], u.Type == godip.Army, true, true) // and legal move counting convoy success
			return err == nil
		}
		return false
	}); len(breaks) > 0 {
		godip.Logf("%v: broken by: %v", self, breaks)
		return godip.ErrSupportBroken{breaks[0]}
	}

	if dislodgers, _, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		return o != nil && // is an order
			u != nil && // is a unit
			o.Type() == godip.Move && // move
			o.Targets()[1].Super() == self.targets[0].Super() && // against us
			u.Nation != unit.Nation && // not from ourselves
			r.Resolve(p) == nil // and it succeeded
	}); len(dislodgers) > 0 {
		godip.Logf("%v: dislodged by: %v", self, dislodgers)
		return godip.ErrSupportBroken{dislodgers[0]}
	}
	return nil
}

func (self *support) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 4 {
			if bits[2] == bits[3] {
				result = SupportHold(godip.Province(bits[0]), godip.Province(bits[2]))
			} else {
				result = SupportMove(godip.Province(bits[0]), godip.Province(bits[2]), godip.Province(bits[3]))
			}
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *support) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() != godip.Movement {
		return
	}
	if !v.Graph().Has(src) {
		return
	}
	supporter, actualSrc, ok := v.Unit(src)
	if !ok {
		return
	}
	if supporter.Nation != nation {
		return
	}
	for _, supportable := range PossibleMoves(v, src, false, false) {
		// Support HOLD.
		if _, supporteeSrc, ok := v.Unit(supportable); ok {
			if result == nil {
				result = godip.Options{}
			}
			if result[godip.SrcProvince(actualSrc)] == nil {
				result[godip.SrcProvince(actualSrc)] = godip.Options{}
			}
			opt, f := result[godip.SrcProvince(actualSrc)][supporteeSrc.Super()]
			if !f {
				opt = godip.Options{}
				result[godip.SrcProvince(actualSrc)][supporteeSrc.Super()] = opt
			}
			opt[supporteeSrc.Super()] = nil
		}
		// Support MOVE.
		for _, mvDst := range v.Graph().Coasts(supportable) {
			if mvDst.Super() == actualSrc.Super() {
				continue
			}
			// For everyone able to move to the possible destination.
			for _, moveSupportable := range PossibleMovesUnit(v, godip.Fleet, mvDst, true, false, nil) {
				if moveSupportable.Super() == actualSrc.Super() {
					continue
				}
				supportee, mvSrc, ok := v.Unit(moveSupportable.Super())
				if !ok || supportee.Type != godip.Fleet || mvSrc != moveSupportable {
					continue
				}
				if result == nil {
					result = godip.Options{}
				}
				if result[godip.SrcProvince(actualSrc)] == nil {
					result[godip.SrcProvince(actualSrc)] = godip.Options{}
				}
				opt, f := result[godip.SrcProvince(actualSrc)][mvSrc.Super()]
				if !f {
					opt = godip.Options{}
					result[godip.SrcProvince(actualSrc)][mvSrc.Super()] = opt
				}
				opt[mvDst.Super()] = nil
			}
			// For everyone able to convoy to the possible destination, avoiding convoy by the supporter.
			for _, moveSupportable := range PossibleMovesUnit(v, godip.Army, mvDst, true, true, &actualSrc) {
				if moveSupportable.Super() == actualSrc.Super() {
					continue
				}
				supportee, mvSrc, ok := v.Unit(moveSupportable)
				if !ok || supportee.Type != godip.Army {
					continue
				}
				if result == nil {
					result = godip.Options{}
				}
				if result[godip.SrcProvince(actualSrc)] == nil {
					result[godip.SrcProvince(actualSrc)] = godip.Options{}
				}
				opt, f := result[godip.SrcProvince(actualSrc)][mvSrc.Super()]
				if !f {
					opt = godip.Options{}
					result[godip.SrcProvince(actualSrc)][mvSrc.Super()] = opt
				}
				opt[mvDst.Super()] = nil
			}
		}
	}
	return
}

func (self *support) Validate(v godip.Validator) (godip.Nation, error) {
	if v.Phase().Type() != godip.Movement {
		return "", godip.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return "", godip.ErrInvalidTarget
	}
	var ok bool
	var unit, supported godip.Unit
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return "", godip.ErrMissingUnit
	}
	if supported, self.targets[1], ok = v.Unit(self.targets[1]); !ok {
		return "", godip.ErrMissingSupportUnit
	}
	if len(self.targets) == 2 {
		if err := AnySupportPossible(v, unit.Type, self.targets[0], self.targets[1]); err != nil {
			return "", godip.ErrIllegalSupportPosition
		}
	} else {
		if !v.Graph().Has(self.targets[2]) {
			return "", godip.ErrInvalidTarget
		}
		if err := AnySupportPossible(v, unit.Type, self.targets[0], self.targets[2]); err != nil {
			return "", godip.ErrIllegalSupportDestination
		}

		if _, err := AnyMovePossible(v, supported.Type, self.targets[1], self.targets[2], true, true, false); err != nil {
			return "", godip.ErrIllegalSupportMove
		}
	}
	return unit.Nation, nil
}

func (self *support) Execute(state godip.State) {
}

func AnySupportPossible(v godip.Validator, typ godip.UnitType, src, dst godip.Province) (err error) {
	if err = movePossible(v, typ, src, dst, false, false); err == nil {
		return
	}
	for _, coast := range v.Graph().Coasts(dst) {
		if err = movePossible(v, typ, src, coast, false, false); err == nil {
			return
		}
	}
	return
}
