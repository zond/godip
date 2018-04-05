package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"

	cla "github.com/zond/godip/variants/classical/common"
)

var MoveOrder = &move{}

var MoveViaConvoyOrder = &move{
	flags: map[godip.Flag]bool{
		cla.ViaConvoy: true,
	},
}

func Move(source, dest godip.Province) *move {
	return &move{
		targets: []godip.Province{source, dest},
		flags:   make(map[godip.Flag]bool),
	}
}

type move struct {
	targets []godip.Province
	flags   map[godip.Flag]bool
}

func (self *move) String() string {
	via := ""
	if self.flags[cla.ViaConvoy] {
		via = " via convoy"
	}
	return fmt.Sprintf("%v %v %v%v", self.targets[0], cla.Move, self.targets[1], via)
}

func (self *move) ViaConvoy() *move {
	self.flags[cla.ViaConvoy] = true
	return self
}

func (self *move) Type() godip.OrderType {
	return cla.Move
}

func (self *move) DisplayType() godip.OrderType {
	if self.flags[cla.ViaConvoy] {
		return cla.MoveViaConvoy
	}
	return cla.Move
}

func (self *move) Targets() []godip.Province {
	return self.targets
}

func (self *move) At() time.Time {
	return time.Now()
}

func (self *move) Adjudicate(r godip.Resolver) error {
	if r.Phase().Type() == cla.Movement {
		return self.adjudicateMovementPhase(r)
	}
	return self.adjudicateRetreatPhase(r)
}

func (self *move) adjudicateRetreatPhase(r godip.Resolver) error {
	for prov, order := range r.Orders() {
		if prov.Super() != self.targets[0].Super() && order.Type() == cla.Move && order.Targets()[1].Super() == self.targets[1].Super() {
			return godip.ErrBounce{order.Targets()[0]}
		}
	}
	return nil
}

func (self *move) Flags() map[godip.Flag]bool {
	return self.flags
}

func (self *move) adjudicateAgainstCompetition(r godip.Resolver, forbiddenSupporter *godip.Nation) error {
	_, competingOrders, competingUnits := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		return o != nil && u != nil && o.Type() == cla.Move && o.Targets()[0] != self.targets[0] && self.targets[1].Super() == o.Targets()[1].Super()
	})
	for index, competingOrder := range competingOrders {
		var forbiddenSupporters []godip.Nation
		if forbiddenSupporter != nil {
			forbiddenSupporters = append(forbiddenSupporters, *forbiddenSupporter)
		}
		attackStrength := cla.MoveSupport(r, self.targets[0], self.targets[1], forbiddenSupporters) + 1
		godip.Logf("'%v' vs '%v': %v", self, competingOrder, attackStrength)
		if as := cla.MoveSupport(r, competingOrder.Targets()[0], competingOrder.Targets()[1], nil) + 1; as >= attackStrength {
			if cla.MustConvoy(r, competingOrder.Targets()[0]) {
				if cla.AnyConvoyPath(r, competingOrder.Targets()[0], competingOrder.Targets()[1], true, nil) != nil {
					godip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
					r.AddBounce(self.targets[0], self.targets[1])
					return godip.ErrBounce{competingOrder.Targets()[0]}
				}
			} else {
				godip.Logf("H2HDisl(%v)", self.targets[1])
				godip.Indent("  ")
				if dislodgers, _, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
					res := o != nil && // is an order
						u != nil && // is a unit
						o.Type() == cla.Move && // move
						o.Targets()[1].Super() == competingOrder.Targets()[0].Super() && // against the competition
						o.Targets()[0].Super() == competingOrder.Targets()[1].Super() && // from their destination
						u.Nation != competingUnits[index].Nation // not from themselves
					if res {
						if !cla.MustConvoy(r, o.Targets()[0]) && r.Resolve(p) == nil {
							return true
						}
					}
					return false
				}); len(dislodgers) == 0 {
					godip.DeIndent()
					godip.Logf("Not dislodged")
					godip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
					r.AddBounce(self.targets[0], self.targets[1])
					return godip.ErrBounce{competingOrder.Targets()[0]}
				} else {
					godip.DeIndent()
					godip.Logf("Dislodged by %v", dislodgers)
				}
			}
		} else {
			godip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
		}
	}
	return nil
}

func (self *move) adjudicateMovementPhase(r godip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])

	convoyed := cla.MustConvoy(r, self.targets[0])
	if convoyed {
		if cla.AnyConvoyPath(r, self.targets[0], self.targets[1], true, nil) == nil {
			return godip.ErrMissingConvoyPath
		}
	}

	if err := self.adjudicateAgainstCompetition(r, nil); err != nil {
		return err
	}

	var forbiddenSupporter *godip.Nation
	// at destination
	if victim, _, hasVictim := r.Unit(self.targets[1]); hasVictim {
		forbiddenSupporter = &victim.Nation
		attackStrength := cla.MoveSupport(r, self.targets[0], self.targets[1], []godip.Nation{victim.Nation}) + 1
		order, prov, _ := r.Order(self.targets[1])
		godip.Logf("'%v' vs '%v': %v", self, order, attackStrength)
		if order.Type() == cla.Move {
			victimConvoyed := cla.MustConvoy(r, order.Targets()[0])
			if !convoyed && !victimConvoyed && order.Targets()[1].Super() == self.targets[0].Super() {
				as := cla.MoveSupport(r, order.Targets()[0], order.Targets()[1], []godip.Nation{unit.Nation}) + 1
				godip.Logf("'%v' vs '%v': %v", order, self, as)
				if victim.Nation == unit.Nation || as >= attackStrength {
					return godip.ErrBounce{self.targets[1]}
				}
			} else {
				godip.Logf("Esc(%v)", order.Targets()[0])
				godip.Indent("  ")
				if err := r.Resolve(prov); err == nil {
					godip.DeIndent()
					godip.Logf("Success")
					forbiddenSupporter = nil
				} else {
					godip.DeIndent()
					godip.Logf("Failure: %v", err)
					if victim.Nation == unit.Nation || 1 >= attackStrength {
						return godip.ErrBounce{self.targets[1]}
					}
				}
			}
		} else {
			hs := cla.HoldSupport(r, self.targets[1]) + 1
			godip.Logf("'%v': %v", order, hs)
			if victim.Nation == unit.Nation || hs >= attackStrength {
				return godip.ErrBounce{self.targets[1]}
			}
		}
	}

	if err := self.adjudicateAgainstCompetition(r, forbiddenSupporter); err != nil {
		return err
	}

	return nil
}

func (self *move) Validate(v godip.Validator) (godip.Nation, error) {
	if v.Phase().Type() == cla.Movement {
		return self.validateMovementPhase(v)
	} else if v.Phase().Type() == cla.Retreat {
		return self.validateRetreatPhase(v)
	}
	return "", godip.ErrInvalidPhase
}

func (self *move) validateRetreatPhase(v godip.Validator) (godip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return "", godip.ErrInvalidDestination
	}
	if self.targets[0] == self.targets[1] {
		return "", godip.ErrIllegalMove
	}
	var unit godip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Dislodged(self.targets[0]); !ok {
		return "", godip.ErrMissingUnit
	}
	var err error
	if self.targets[1], err = cla.AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == cla.Army, false, false); err != nil {
		return "", godip.ErrIllegalMove
	}
	if _, _, ok := v.Unit(self.targets[1]); ok {
		return "", godip.ErrIllegalRetreat
	}
	if v.Bounce(self.targets[0], self.targets[1]) {
		return "", godip.ErrIllegalRetreat
	}
	return unit.Nation, nil
}

func (self *move) validateMovementPhase(v godip.Validator) (godip.Nation, error) {
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return "", godip.ErrInvalidDestination
	}
	if self.targets[0] == self.targets[1] {
		return "", godip.ErrIllegalMove
	}
	var unit godip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return "", godip.ErrMissingUnit
	}
	var err error
	if self.targets[1], err = cla.AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == cla.Army, true, false); err != nil {
		return "", err
	}
	return unit.Nation, nil
}

func (self *move) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if !self.flags[cla.ViaConvoy] {
		if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
			if len(bits) == 3 {
				result = Move(godip.Province(bits[0]), godip.Province(bits[2]))
			}
			if result == nil {
				err = fmt.Errorf("Can't parse as %+v", bits)
			}
		}
	} else {
		if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
			if len(bits) == 3 {
				result = Move(godip.Province(bits[0]), godip.Province(bits[2])).ViaConvoy()
			}
			if result == nil {
				err = fmt.Errorf("Can't parse as %+v", bits)
			}
		}
	}
	return result, err
}

func (self *move) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() == cla.Retreat {
		if !self.flags[cla.ViaConvoy] {
			if v.Graph().Has(src) {
				if unit, actualSrc, ok := v.Dislodged(src); ok {
					if unit.Nation == nation {
						for _, dst := range cla.PossibleMoves(v, src, false, true) {
							if _, _, foundUnit := v.Unit(dst); !foundUnit {
								if !v.Bounce(src, dst) {
									if result == nil {
										result = godip.Options{}
									}
									if result[godip.SrcProvince(actualSrc)] == nil {
										result[godip.SrcProvince(actualSrc)] = godip.Options{}
									}
									result[godip.SrcProvince(actualSrc)][dst] = nil
								}
							}
						}
					}
				}
			}
		}
	} else if v.Phase().Type() == cla.Movement {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
				if unit.Nation == nation {
					if !self.flags[cla.ViaConvoy] || unit.Type == cla.Army {
						for _, dst := range cla.PossibleMoves(v, src, true, false) {
							if !self.flags[cla.ViaConvoy] {
								if result == nil {
									result = godip.Options{}
								}
								if result[godip.SrcProvince(actualSrc)] == nil {
									result[godip.SrcProvince(actualSrc)] = godip.Options{}
								}
								result[godip.SrcProvince(actualSrc)][dst] = nil
							} else {
								if cp := cla.AnyConvoyPath(v, src, dst, false, nil); len(cp) > 1 {
									if result == nil {
										result = godip.Options{}
									}
									if result[godip.SrcProvince(actualSrc)] == nil {
										result[godip.SrcProvince(actualSrc)] = godip.Options{}
									}
									result[godip.SrcProvince(actualSrc)][dst] = nil
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

func (self *move) Execute(state godip.State) {
	if state.Phase().Type() == cla.Retreat {
		state.Retreat(self.targets[0], self.targets[1])
	} else {
		state.Move(self.targets[0], self.targets[1], !cla.MustConvoy(state, self.targets[0]))
	}
}
