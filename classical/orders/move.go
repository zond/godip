package orders

import (
	"fmt"
	"time"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

func init() {
	generators = append(generators, func() dip.Order { return &move{} })
	generators = append(generators, func() dip.Order {
		return &move{
			flags: map[dip.Flag]bool{
				cla.ViaConvoy: true,
			},
		}
	})
}

func Move(source, dest dip.Province) *move {
	return &move{
		targets: []dip.Province{source, dest},
		flags:   make(map[dip.Flag]bool),
	}
}

type move struct {
	targets []dip.Province
	flags   map[dip.Flag]bool
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

func (self *move) Type() dip.OrderType {
	return cla.Move
}

func (self *move) DisplayType() dip.OrderType {
	if self.flags[cla.ViaConvoy] {
		return cla.MoveViaConvoy
	}
	return cla.Move
}

func (self *move) Targets() []dip.Province {
	return self.targets
}

func (self *move) At() time.Time {
	return time.Now()
}

func (self *move) Adjudicate(r dip.Resolver) error {
	if r.Phase().Type() == cla.Movement {
		return self.adjudicateMovementPhase(r)
	}
	return self.adjudicateRetreatPhase(r)
}

func (self *move) adjudicateRetreatPhase(r dip.Resolver) error {
	for prov, order := range r.Orders() {
		if prov.Super() != self.targets[0].Super() && order.Type() == cla.Move && order.Targets()[1].Super() == self.targets[1].Super() {
			return cla.ErrBounce{order.Targets()[0]}
		}
	}
	return nil
}

func (self *move) Flags() map[dip.Flag]bool {
	return self.flags
}

func (self *move) adjudicateAgainstCompetition(r dip.Resolver, forbiddenSupporter *dip.Nation) error {
	_, competingOrders, competingUnits := r.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
		return o != nil && u != nil && o.Type() == cla.Move && o.Targets()[0] != self.targets[0] && self.targets[1].Super() == o.Targets()[1].Super()
	})
	for index, competingOrder := range competingOrders {
		var forbiddenSupporters []dip.Nation
		if forbiddenSupporter != nil {
			forbiddenSupporters = append(forbiddenSupporters, *forbiddenSupporter)
		}
		attackStrength := cla.MoveSupport(r, self.targets[0], self.targets[1], forbiddenSupporters) + 1
		dip.Logf("'%v' vs '%v': %v", self, competingOrder, attackStrength)
		if as := cla.MoveSupport(r, competingOrder.Targets()[0], competingOrder.Targets()[1], nil) + 1; as >= attackStrength {
			if cla.MustConvoy(r, competingOrder.Targets()[0]) {
				if cla.AnyConvoyPath(r, competingOrder.Targets()[0], competingOrder.Targets()[1], true, nil) != nil {
					dip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
					r.AddBounce(self.targets[0], self.targets[1])
					return cla.ErrBounce{competingOrder.Targets()[0]}
				}
			} else {
				dip.Logf("H2HDisl(%v)", self.targets[1])
				dip.Indent("  ")
				if dislodgers, _, _ := r.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
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
					dip.DeIndent()
					dip.Logf("Not dislodged")
					dip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
					r.AddBounce(self.targets[0], self.targets[1])
					return cla.ErrBounce{competingOrder.Targets()[0]}
				} else {
					dip.DeIndent()
					dip.Logf("Dislodged by %v", dislodgers)
				}
			}
		} else {
			dip.Logf("'%v' vs '%v': %v", competingOrder, self, as)
		}
	}
	return nil
}

func (self *move) adjudicateMovementPhase(r dip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])

	convoyed := cla.MustConvoy(r, self.targets[0])
	if convoyed {
		if cla.AnyConvoyPath(r, self.targets[0], self.targets[1], true, nil) == nil {
			return cla.ErrMissingConvoyPath
		}
	}

	if err := self.adjudicateAgainstCompetition(r, nil); err != nil {
		return err
	}

	var forbiddenSupporter *dip.Nation
	// at destination
	if victim, _, hasVictim := r.Unit(self.targets[1]); hasVictim {
		forbiddenSupporter = &victim.Nation
		attackStrength := cla.MoveSupport(r, self.targets[0], self.targets[1], []dip.Nation{victim.Nation}) + 1
		order, prov, _ := r.Order(self.targets[1])
		dip.Logf("'%v' vs '%v': %v", self, order, attackStrength)
		if order.Type() == cla.Move {
			victimConvoyed := cla.MustConvoy(r, order.Targets()[0])
			if !convoyed && !victimConvoyed && order.Targets()[1].Super() == self.targets[0].Super() {
				as := cla.MoveSupport(r, order.Targets()[0], order.Targets()[1], []dip.Nation{unit.Nation}) + 1
				dip.Logf("'%v' vs '%v': %v", order, self, as)
				if victim.Nation == unit.Nation || as >= attackStrength {
					return cla.ErrBounce{self.targets[1]}
				}
			} else {
				dip.Logf("Esc(%v)", order.Targets()[0])
				dip.Indent("  ")
				if err := r.Resolve(prov); err == nil {
					dip.DeIndent()
					dip.Logf("Success")
					forbiddenSupporter = nil
				} else {
					dip.DeIndent()
					dip.Logf("Failure: %v", err)
					if victim.Nation == unit.Nation || 1 >= attackStrength {
						return cla.ErrBounce{self.targets[1]}
					}
				}
			}
		} else {
			hs := cla.HoldSupport(r, self.targets[1]) + 1
			dip.Logf("'%v': %v", order, hs)
			if victim.Nation == unit.Nation || hs >= attackStrength {
				return cla.ErrBounce{self.targets[1]}
			}
		}
	}

	if err := self.adjudicateAgainstCompetition(r, forbiddenSupporter); err != nil {
		return err
	}

	return nil
}

func (self *move) Validate(v dip.Validator) error {
	if v.Phase().Type() == cla.Movement {
		return self.validateMovementPhase(v)
	} else if v.Phase().Type() == cla.Retreat {
		return self.validateRetreatPhase(v)
	}
	return cla.ErrInvalidPhase
}

func (self *move) validateRetreatPhase(v dip.Validator) error {
	if !v.Graph().Has(self.targets[0]) {
		return cla.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return cla.ErrInvalidDestination
	}
	if self.targets[0] == self.targets[1] {
		return cla.ErrIllegalMove
	}
	var unit dip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Dislodged(self.targets[0]); !ok {
		return cla.ErrMissingUnit
	}
	var err error
	if self.targets[1], err = cla.AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == cla.Army, false, false); err != nil {
		return cla.ErrIllegalMove
	}
	if _, _, ok := v.Unit(self.targets[1]); ok {
		return cla.ErrIllegalRetreat
	}
	if v.Bounce(self.targets[0], self.targets[1]) {
		return cla.ErrIllegalRetreat
	}
	return nil
}

func (self *move) validateMovementPhase(v dip.Validator) error {
	if !v.Graph().Has(self.targets[0]) {
		return cla.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return cla.ErrInvalidDestination
	}
	if self.targets[0] == self.targets[1] {
		return cla.ErrIllegalMove
	}
	var unit dip.Unit
	var ok bool
	if unit, self.targets[0], ok = v.Unit(self.targets[0]); !ok {
		return cla.ErrMissingUnit
	}
	var err error
	if self.targets[1], err = cla.AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == cla.Army, true, false); err != nil {
		return err
	}
	return nil
}

func (self *move) Options(v dip.Validator, nation dip.Nation, src dip.Province) (result dip.Options) {
	if v.Phase().Type() == cla.Retreat {
		if !self.flags[cla.ViaConvoy] {
			if v.Graph().Has(src) {
				if unit, actualSrc, ok := v.Dislodged(src); ok {
					if unit.Nation == nation {
						for _, dst := range cla.PossibleMoves(v, src, false, true) {
							if _, _, foundUnit := v.Unit(dst); !foundUnit {
								if !v.Bounce(src, dst) {
									if result == nil {
										result = dip.Options{}
									}
									if result[dip.SrcProvince(actualSrc)] == nil {
										result[dip.SrcProvince(actualSrc)] = dip.Options{}
									}
									result[dip.SrcProvince(actualSrc)][dst] = nil
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
									result = dip.Options{}
								}
								if result[dip.SrcProvince(actualSrc)] == nil {
									result[dip.SrcProvince(actualSrc)] = dip.Options{}
								}
								result[dip.SrcProvince(actualSrc)][dst] = nil
							} else {
								if cp := cla.AnyConvoyPath(v, src, dst, false, nil); len(cp) > 1 {
									if result == nil {
										result = dip.Options{}
									}
									if result[dip.SrcProvince(actualSrc)] == nil {
										result[dip.SrcProvince(actualSrc)] = dip.Options{}
									}
									result[dip.SrcProvince(actualSrc)][dst] = nil
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

func (self *move) Execute(state dip.State) {
	if state.Phase().Type() == cla.Retreat {
		state.Retreat(self.targets[0], self.targets[1])
	} else {
		state.Move(self.targets[0], self.targets[1], !cla.MustConvoy(state, self.targets[0]))
	}
}
