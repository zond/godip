package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"
)

var MoveOrder = &move{}

var MoveViaConvoyOrder = &move{
	flags: map[godip.Flag]bool{
		godip.ViaConvoy: true,
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
	if self.flags[godip.ViaConvoy] {
		via = " via convoy"
	}
	return fmt.Sprintf("%v %v %v%v", self.targets[0], godip.Move, self.targets[1], via)
}

func (self *move) ViaConvoy() *move {
	self.flags[godip.ViaConvoy] = true
	return self
}

func (self *move) Type() godip.OrderType {
	return godip.Move
}

func (self *move) DisplayType() godip.OrderType {
	if self.flags[godip.ViaConvoy] {
		return godip.MoveViaConvoy
	}
	return godip.Move
}

func (self *move) Targets() []godip.Province {
	return self.targets
}

func (self *move) At() time.Time {
	return time.Now()
}

func (self *move) Adjudicate(r godip.Resolver) error {
	if r.Phase().Type() == godip.Movement {
		return self.adjudicateMovementPhase(r)
	}
	return self.adjudicateRetreatPhase(r)
}

func (self *move) adjudicateRetreatPhase(r godip.Resolver) error {
	for prov, order := range r.Orders() {
		if prov.Super() != self.targets[0].Super() && order.Type() == godip.Move && order.Targets()[1].Super() == self.targets[1].Super() {
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
		return o != nil && u != nil && o.Type() == godip.Move && o.Targets()[0] != self.targets[0] && self.targets[1].Super() == o.Targets()[1].Super()
	})
	for index, competingOrder := range competingOrders {
		var forbiddenSupporters []godip.Nation
		if forbiddenSupporter != nil {
			forbiddenSupporters = append(forbiddenSupporters, *forbiddenSupporter)
		}
		attackStrength := MoveSupport(r, self.targets[0], self.targets[1], forbiddenSupporters) + 1
		godip.Logf("'%v' vs '%v': %v", self, competingOrder, attackStrength)
		if as := MoveSupport(r, competingOrder.Targets()[0], competingOrder.Targets()[1], nil) + 1; as >= attackStrength {
			if MustConvoy(r, competingOrder.Targets()[0]) {
				if AnyConvoyPath(r, competingOrder.Targets()[0], competingOrder.Targets()[1], true, nil) != nil {
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
						o.Type() == godip.Move && // move
						o.Targets()[1].Super() == competingOrder.Targets()[0].Super() && // against the competition
						o.Targets()[0].Super() == competingOrder.Targets()[1].Super() && // from their destination
						u.Nation != competingUnits[index].Nation // not from themselves
					if res {
						if !MustConvoy(r, o.Targets()[0]) && r.Resolve(p) == nil {
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

	convoyed := MustConvoy(r, self.targets[0])
	if convoyed {
		if AnyConvoyPath(r, self.targets[0], self.targets[1], true, nil) == nil {
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
		attackStrength := MoveSupport(r, self.targets[0], self.targets[1], []godip.Nation{victim.Nation}) + 1
		order, prov, _ := r.Order(self.targets[1])
		godip.Logf("'%v' vs '%v': %v", self, order, attackStrength)
		if order.Type() == godip.Move {
			victimConvoyed := MustConvoy(r, order.Targets()[0])
			if !convoyed && !victimConvoyed && order.Targets()[1].Super() == self.targets[0].Super() {
				as := MoveSupport(r, order.Targets()[0], order.Targets()[1], []godip.Nation{unit.Nation}) + 1
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
			hs := HoldSupport(r, self.targets[1]) + 1
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
	if v.Phase().Type() == godip.Movement {
		return self.validateMovementPhase(v)
	} else if v.Phase().Type() == godip.Retreat {
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
	if self.targets[1], err = AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == godip.Army, false, false); err != nil {
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
	if self.targets[1], err = AnyMovePossible(v, unit.Type, self.targets[0], self.targets[1], unit.Type == godip.Army, true, false); err != nil {
		return "", err
	}
	return unit.Nation, nil
}

func (self *move) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if !self.flags[godip.ViaConvoy] {
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
	if v.Phase().Type() == godip.Retreat {
		if !self.flags[godip.ViaConvoy] {
			if v.Graph().Has(src) {
				if unit, actualSrc, ok := v.Dislodged(src); ok {
					if unit.Nation == nation {
						for _, dst := range PossibleMoves(v, src, false, true) {
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
	} else if v.Phase().Type() == godip.Movement {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
				if unit.Nation == nation {
					if !self.flags[godip.ViaConvoy] || unit.Type == godip.Army {
						for _, dst := range PossibleMoves(v, src, true, false) {
							if !self.flags[godip.ViaConvoy] {
								if result == nil {
									result = godip.Options{}
								}
								if result[godip.SrcProvince(actualSrc)] == nil {
									result[godip.SrcProvince(actualSrc)] = godip.Options{}
								}
								result[godip.SrcProvince(actualSrc)][dst] = nil
							} else {
								if cp := AnyConvoyPath(v, src, dst, false, nil); len(cp) > 1 {
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
	if state.Phase().Type() == godip.Retreat {
		state.Retreat(self.targets[0], self.targets[1])
	} else {
		state.Move(self.targets[0], self.targets[1], !MustConvoy(state, self.targets[0]))
	}
}

/*
MoveSupport returns the successful supports of movement from src to dst, discounting the nations in forbiddenSupports.
*/
func MoveSupport(r godip.Resolver, src, dst godip.Province, forbiddenSupports []godip.Nation) int {
	_, supports, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		if o != nil && u != nil {
			if o.Type() == godip.Support && len(o.Targets()) == 3 && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst) {
				for _, ban := range forbiddenSupports {
					if ban == u.Nation {
						return false
					}
				}
				if err := r.Resolve(p); err == nil {
					return true
				}
			}
		}
		return false
	})
	return len(supports)
}

func HasEdge(v godip.Validator, typ godip.UnitType, src, dst godip.Province) bool {
	if typ == godip.Army {
		return v.Graph().Flags(dst)[godip.Land] && v.Graph().Edges(src, false)[dst][godip.Land]
	} else {
		return v.Graph().Flags(dst)[godip.Sea] && v.Graph().Edges(src, false)[dst][godip.Sea]
	}
}

// PossibleMovesUnit returns the possible provinces that a unit can move to or
// from. If reverse is false then the unit must be in start, and if true then the
// potential units must be able to move to start. Setting allowConvoy allows convoy
// routes to be considered, and these must avoid the province noConvoy (if given).
func PossibleMovesUnit(v godip.Validator, unitType godip.UnitType, start godip.Province, reverse bool, allowConvoy bool, noConvoy *godip.Province) (result []godip.Province) {
	defer v.Profile("PossibleMovesUnit", time.Now())
	noConvoyStr := ""
	if noConvoy != nil {
		noConvoyStr = string(*noConvoy)
	}
	return v.MemoizeProvSlice(fmt.Sprintf("PossibleMovesUnit(%v,%v,%v,%v,%v)", unitType, start, reverse, allowConvoy, noConvoyStr), func() []godip.Province {
		neighbours := v.Graph().Edges(start, reverse)
		ends := map[godip.Province]bool{}
		if unitType == godip.Army {
			if v.Graph().Flags(start)[godip.Land] {
				for end, flags := range neighbours {
					if flags[godip.Land] && v.Graph().Flags(end)[godip.Land] {
						ends[end] = true
					}
				}
				if allowConvoy {
					for _, coast := range v.Graph().Coasts(start) {
						for _, end := range ConvoyEndPoints(v, coast, reverse, noConvoy) {
							ends[end] = true
						}
					}
				}
			}
		} else if unitType == godip.Fleet {
			for end, flags := range neighbours {
				if flags[godip.Sea] && v.Graph().Flags(end)[godip.Sea] {
					ends[end] = true
				}
			}
		} else {
			panic(fmt.Errorf("unknown unit type %q", unitType))
		}
		for end, _ := range ends {
			if end.Super() == end || !ends[end.Super()] {
				result = append(result, end)
			}
		}
		return result
	})
}

func PossibleMoves(v godip.Validator, src godip.Province, allowConvoy, dislodged bool) (result []godip.Province) {
	defer v.Profile("PossibleMoves", time.Now())
	var unit godip.Unit
	var realSrc godip.Province
	var found bool
	if dislodged {
		unit, realSrc, found = v.Dislodged(src)
	} else {
		unit, realSrc, found = v.Unit(src)
	}
	if found {
		return PossibleMovesUnit(v, unit.Type, realSrc, false, allowConvoy, nil)
	}
	return nil
}

func AnyMovePossible(v godip.Validator, typ godip.UnitType, src, dst godip.Province, lax, allowConvoy, resolveConvoys bool) (dstCoast godip.Province, err error) {
	defer v.Profile("AnyMovePossible", time.Now())
	dstCoast = dst
	if err = movePossible(v, typ, src, dst, allowConvoy, resolveConvoys); err == nil {
		return
	}
	if lax || dst.Super() == dst {
		var options []godip.Province
		for _, coast := range v.Graph().Coasts(dst) {
			if err2 := movePossible(v, typ, src, coast, allowConvoy, resolveConvoys); err2 == nil {
				options = append(options, coast)
			}
		}
		if len(options) > 0 {
			if lax || len(options) == 1 {
				dstCoast, err = options[0], nil
			}
		}
	}
	return
}

func movePossible(v godip.Validator, typ godip.UnitType, src, dst godip.Province, allowConvoy, resolveConvoys bool) error {
	defer v.Profile("movePossible", time.Now())
	if !v.Graph().Has(src) {
		return godip.ErrInvalidSource
	}
	if !v.Graph().Has(dst) {
		return godip.ErrInvalidDestination
	}
	if typ == godip.Army {
		defer v.Profile("movePossible (army)", time.Now())
		if !v.Graph().Flags(dst)[godip.Land] {
			return godip.ErrIllegalDestination
		}
		if !allowConvoy {
			flags, found := v.Graph().Edges(src, false)[dst]
			if !found {
				return godip.ErrIllegalMove
			}
			if !flags[godip.Land] {
				return godip.ErrIllegalDestination
			}
			return nil
		}
		if resolveConvoys {
			if MustConvoy(v.(godip.Resolver), src) {
				if AnyConvoyPath(v, src, dst, true, nil) == nil {
					return godip.ErrMissingConvoyPath
				}
				return nil
			}
		}
		if !HasEdge(v, typ, src, dst) {
			if cp := AnyConvoyPath(v, src, dst, false, nil); cp == nil {
				return godip.ErrMissingConvoyPath
			}
			return nil
		}
		return nil
	} else if typ == godip.Fleet {
		defer v.Profile("movePossible (fleet)", time.Now())
		if !v.Graph().Flags(dst)[godip.Sea] {
			return godip.ErrIllegalDestination
		}
		if !HasEdge(v, typ, src, dst) {
			return godip.ErrIllegalMove
		}
	}
	return nil
}
