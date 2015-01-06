package common

import (
	"fmt"
	. "github.com/zond/godip/common"
	"sort"
	"time"
)

const (
	Sea  Flag = "Sea"
	Land Flag = "Land"

	Army  UnitType = "Army"
	Fleet UnitType = "Fleet"

	England Nation = "England"
	France  Nation = "France"
	Germany Nation = "Germany"
	Russia  Nation = "Russia"
	Austria Nation = "Austria"
	Italy   Nation = "Italy"
	Turkey  Nation = "Turkey"
	Neutral Nation = "Neutral"

	Spring Season = "Spring"
	Fall   Season = "Fall"

	Movement   PhaseType = "Movement"
	Retreat    PhaseType = "Retreat"
	Adjustment PhaseType = "Adjustment"

	Build         OrderType = "Build"
	Move          OrderType = "Move"
	MoveViaConvoy OrderType = "MoveViaConvoy"
	Hold          OrderType = "Hold"
	Convoy        OrderType = "Convoy"
	Support       OrderType = "Support"
	Disband       OrderType = "Disband"

	ViaConvoy Flag = "C"
)

var Coast = []Flag{Sea, Land}

var Nations = []Nation{Austria, England, France, Germany, Italy, Turkey, Russia}
var PhaseTypes = []PhaseType{Movement, Retreat, Adjustment}
var Seasons = []Season{Spring, Fall}
var UnitTypes = []UnitType{Army, Fleet}

// Invalid is not understood
// Illegal is understood but not allowed
var ErrInvalidSource = fmt.Errorf("ErrInvalidSource")
var ErrInvalidDestination = fmt.Errorf("ErrInvalidDestination")
var ErrInvalidTarget = fmt.Errorf("ErrInvalidTarget")
var ErrInvalidPhase = fmt.Errorf("ErrInvalidPhase")
var ErrMissingUnit = fmt.Errorf("ErrMissingUnit")
var ErrIllegalDestination = fmt.Errorf("ErrIllegalDestination")
var ErrMissingConvoyPath = fmt.Errorf("ErrMissignConvoyPath")
var ErrIllegalMove = fmt.Errorf("ErrIllegalMove")
var ErrConvoyParadox = fmt.Errorf("ErrConvoyParadox")
var ErrIllegalSupportPosition = fmt.Errorf("ErrIllegalSupportPosition")
var ErrIllegalSupportDestination = fmt.Errorf("ErrIllegalSupportDestination")
var ErrIllegalSupportDestinationNation = fmt.Errorf("ErrIllegalSupportDestinationNation")
var ErrMissingSupportUnit = fmt.Errorf("ErrMissingSupportUnit")
var ErrIllegalSupportMove = fmt.Errorf("ErrIllegalSupportMove")
var ErrIllegalConvoyUnit = fmt.Errorf("ErrIllegalConvoyUnit")
var ErrIllegalConvoyPath = fmt.Errorf("ErrIllegalConvoyPath")
var ErrIllegalConvoyMove = fmt.Errorf("ErrIllegalConvoyMove")
var ErrMissingConvoyee = fmt.Errorf("ErrMissingConvoyee")
var ErrIllegalConvoyer = fmt.Errorf("ErrIllegalConvoyer")
var ErrIllegalConvoyee = fmt.Errorf("ErrIllegalConvoyee")
var ErrIllegalBuild = fmt.Errorf("ErrIllegalBuild")
var ErrIllegalDisband = fmt.Errorf("ErrIllegalDisband")
var ErrOccupiedSupplyCenter = fmt.Errorf("ErrOccupiedSupplyCenter")
var ErrMissingSupplyCenter = fmt.Errorf("ErrMissingSupplyCenter")
var ErrMissingSurplus = fmt.Errorf("ErrMissingSurplus")
var ErrIllegalUnitType = fmt.Errorf("ErrIllegalUnitType")
var ErrMissingDeficit = fmt.Errorf("ErrMissingDeficit")
var ErrOccupiedDestination = fmt.Errorf("ErrOccupiedDestination")
var ErrIllegalRetreat = fmt.Errorf("ErrIllegalRetreat")
var ErrForcedDisband = fmt.Errorf("ErrForcedDisband")
var ErrHostileSupplyCenter = fmt.Errorf("ErrHostileSupplyCenter")

type ErrDoubleBuild struct {
	Provinces []Province
}

func (self ErrDoubleBuild) Error() string {
	return fmt.Sprintf("ErrDoubleBuild:%v", self.Provinces)
}

type ErrConvoyDislodged struct {
	Province Province
}

func (self ErrConvoyDislodged) Error() string {
	return fmt.Sprintf("ErrConvoyDislodged:%v", self.Province)
}

type ErrSupportBroken struct {
	Province Province
}

func (self ErrSupportBroken) Error() string {
	return fmt.Sprintf("ErrSupportBroken:%v", self.Province)
}

type ErrBounce struct {
	Province Province
}

func (self ErrBounce) Error() string {
	return fmt.Sprintf("ErrBounce:%v", self.Province)
}

func PossibleConvoyPathFilter(v Validator, src, dst Province, resolveConvoys, dstOk bool) PathFilter {
	return func(name Province, edgeFlags, nodeFlags map[Flag]bool, sc *Nation) bool {
		if dstOk && name.Contains(dst) {
			return true
		}
		if nodeFlags[Land] {
			return false
		}
		if u, _, ok := v.Unit(name); ok && u.Type == Fleet {
			if !resolveConvoys {
				return true
			}
			if order, prov, ok := v.Order(name); ok && order.Type() == Convoy && order.Targets()[1].Contains(src) && order.Targets()[2].Contains(dst) {
				if err := v.(Resolver).Resolve(prov); err != nil {
					return false
				}
				return true
			}
		}
		return false
	}
}

func convoyPath(v Validator, src, dst Province, resolveConvoys bool, viaNation *Nation) []Province {
	if src == dst {
		return nil
	}
	waypoints, _, _ := v.Find(func(p Province, o Order, u *Unit) bool {
		if !v.Graph().Flags(p)[Land] && u != nil && (viaNation == nil || u.Nation == *viaNation) && u.Type == Fleet {
			if !resolveConvoys {
				if viaNation == nil || (o != nil && o.Type() == Convoy && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst)) {
					return true
				}
				return false
			}
			if o != nil && o.Type() == Convoy && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst) {
				if err := v.(Resolver).Resolve(p); err != nil {
					return false
				}
				return true
			}
		}
		return false
	})
	for _, waypoint := range waypoints {
		if part1 := v.Graph().Path(src, waypoint, PossibleConvoyPathFilter(v, src, dst, resolveConvoys, false)); part1 != nil {
			if part2 := v.Graph().Path(waypoint, dst, PossibleConvoyPathFilter(v, src, dst, resolveConvoys, true)); part2 != nil {
				return append(part1, part2...)
			}
		}
	}
	return nil
}

func HasEdge(v Validator, typ UnitType, src, dst Province) bool {
	if typ == Army {
		return v.Graph().Flags(dst)[Land] && v.Graph().Edges(src)[dst][Land]
	} else {
		return v.Graph().Flags(dst)[Sea] && v.Graph().Edges(src)[dst][Sea]
	}
}

func MustConvoy(r Resolver, src Province) bool {
	unit, _, ok := r.Unit(src)
	if !ok {
		return false
	}
	if unit.Type != Army {
		return false
	}
	order, _, ok := r.Order(src)
	if !ok {
		return false
	}
	if order.Type() != Move {
		return false
	}
	return (!HasEdge(r, unit.Type, order.Targets()[0], order.Targets()[1]) ||
		(order.Flags()[ViaConvoy] && AnyConvoyPath(r, order.Targets()[0], order.Targets()[1], true, nil) != nil) ||
		AnyConvoyPath(r, order.Targets()[0], order.Targets()[1], false, &unit.Nation) != nil)
}

func AnyConvoyPath(v Validator, src, dst Province, resolveConvoys bool, viaNation *Nation) (result []Province) {
	if !v.Graph().AllFlags(src)[Sea] || !v.Graph().AllFlags(dst)[Sea] {
		return
	}
	if result = convoyPath(v, src, dst, resolveConvoys, viaNation); result != nil {
		return
	}
	for _, srcCoast := range v.Graph().Coasts(src) {
		for _, dstCoast := range v.Graph().Coasts(dst) {
			if result = convoyPath(v, srcCoast, dstCoast, resolveConvoys, viaNation); result != nil {
				return
			}
		}
	}
	return
}

func AnySupportPossible(v Validator, typ UnitType, src, dst Province) (err error) {
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

func PossibleMoves(v Validator, src Province, allowConvoy, dislodged bool) (result []Province) {
	dsts := map[Province]bool{}
	var unit Unit
	var realSrc Province
	var found bool
	if dislodged {
		unit, realSrc, found = v.Dislodged(src)
	} else {
		unit, realSrc, found = v.Unit(src)
	}
	if found {
		if unit.Type == Army && !allowConvoy {
			for dst, flags := range v.Graph().Edges(realSrc) {
				if flags[Land] && v.Graph().Flags(dst)[Land] {
					dsts[dst] = true
				}
			}
		} else if unit.Type == Fleet {
			for dst, flags := range v.Graph().Edges(realSrc) {
				if flags[Sea] && v.Graph().Flags(dst)[Sea] {
					dsts[dst] = true
				}
			}
		} else {
			for _, prov := range v.Graph().Provinces() {
				if err := movePossible(v, unit.Type, realSrc, prov, allowConvoy, false); err == nil {
					dsts[prov] = true
				}
			}
		}
	}
	for dst, _ := range dsts {
		if dst.Super() == dst {
			result = append(result, dst)
		} else if !dsts[dst.Super()] {
			foundCoasts := 0
			for _, coast := range v.Graph().Coasts(dst) {
				if dsts[coast] {
					foundCoasts += 1
				}
			}
			if foundCoasts == 1 {
				result = append(result, dst.Super())
			} else {
				result = append(result, dst)
			}
		}
	}
	return
}

func AnyMovePossible(v Validator, typ UnitType, src, dst Province, lax, allowConvoy, resolveConvoys bool) (dstCoast Province, err error) {
	dstCoast = dst
	if err = movePossible(v, typ, src, dst, allowConvoy, resolveConvoys); err == nil {
		return
	}
	if lax || dst.Super() == dst {
		var options []Province
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

func movePossible(v Validator, typ UnitType, src, dst Province, allowConvoy, resolveConvoys bool) error {
	defer v.Profile("movePossible", time.Now())
	if !v.Graph().Has(src) {
		return ErrInvalidSource
	}
	if !v.Graph().Has(dst) {
		return ErrInvalidDestination
	}
	if typ == Army {
		defer v.Profile("movePossible (army)", time.Now())
		if !v.Graph().Flags(dst)[Land] {
			return ErrIllegalDestination
		}
		if !allowConvoy {
			flags, found := v.Graph().Edges(src)[dst]
			if !found {
				return ErrIllegalMove
			}
			if !flags[Land] {
				return ErrIllegalDestination
			}
			return nil
		}
		if resolveConvoys {
			if MustConvoy(v.(Resolver), src) {
				if AnyConvoyPath(v, src, dst, true, nil) == nil {
					return ErrMissingConvoyPath
				}
				return nil
			}
		}
		if !HasEdge(v, typ, src, dst) {
			if cp := AnyConvoyPath(v, src, dst, false, nil); cp == nil {
				return ErrMissingConvoyPath
			}
			return nil
		}
		return nil
	} else if typ == Fleet {
		defer v.Profile("movePossible (fleet)", time.Now())
		if !v.Graph().Flags(dst)[Sea] {
			return ErrIllegalDestination
		}
		if !HasEdge(v, typ, src, dst) {
			return ErrIllegalMove
		}
	}
	return nil
}

func AdjustmentStatus(v Validator, me Nation) (builds Orders, disbands Orders, balance int) {
	scs := 0
	for prov, nat := range v.SupplyCenters() {
		if nat == me {
			scs += 1
			if order, _, ok := v.Order(prov); ok && order.Type() == Build {
				builds = append(builds, order)
			}
		}
	}

	units := 0
	for prov, unit := range v.Units() {
		if unit.Nation == me {
			units += 1
			if order, _, ok := v.Order(prov); ok && order.Type() == Disband {
				disbands = append(disbands, order)
			}
		}
	}

	sort.Sort(builds)
	sort.Sort(disbands)

	balance = scs - units
	if balance > 0 {
		disbands = nil
		builds = builds[:Max(0, Min(len(builds), balance))]
	} else if balance < 0 {
		builds = nil
		disbands = disbands[:Max(0, Min(len(disbands), -balance))]
	} else {
		builds = nil
		disbands = nil
	}

	return
}

/*
HoldSupport returns successful supports of a hold in prov.
*/
func HoldSupport(r Resolver, prov Province) int {
	_, supports, _ := r.Find(func(p Province, o Order, u *Unit) bool {
		if o != nil && u != nil && o.Type() == Support && p.Super() != prov.Super() && len(o.Targets()) == 2 && o.Targets()[1].Super() == prov.Super() {
			if err := r.Resolve(p); err == nil {
				return true
			}
		}
		return false
	})
	return len(supports)
}

/*
MoveSupport returns the successful supports of movement from src to dst, discounting the nations in forbiddenSupports.
*/
func MoveSupport(r Resolver, src, dst Province, forbiddenSupports []Nation) int {
	_, supports, _ := r.Find(func(p Province, o Order, u *Unit) bool {
		if o != nil && u != nil {
			if o.Type() == Support && len(o.Targets()) == 3 && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst) {
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
