package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"
)

var ConvoyOrder = &convoy{}

func Convoy(source, from, to godip.Province) *convoy {
	return &convoy{
		targets: []godip.Province{source, from, to},
	}
}

type convoy struct {
	targets []godip.Province
}

func (self *convoy) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], godip.Convoy, self.targets[1:])
}

func (self *convoy) Flags() map[godip.Flag]bool {
	return nil
}

func (self *convoy) At() time.Time {
	return time.Now()
}

func (self *convoy) Type() godip.OrderType {
	return godip.Convoy
}

func (self *convoy) DisplayType() godip.OrderType {
	return godip.Convoy
}

func (self *convoy) Targets() []godip.Province {
	return self.targets
}

func (self *convoy) Adjudicate(r godip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	if breaks, _, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		return (o.Type() == godip.Move && // move
			o.Targets()[1] == self.targets[0] && // against us
			u.Nation != unit.Nation && // not friendly
			r.Resolve(p) == nil)
	}); len(breaks) > 0 {
		return godip.ErrConvoyDislodged{breaks[0]}
	}
	return nil
}

func (self *convoy) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 4 {
			result = Convoy(godip.Province(bits[0]), godip.Province(bits[2]), godip.Province(bits[3]))
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *convoy) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() != godip.Movement {
		return
	}
	if !v.Graph().Has(src) {
		return
	}
	convoyer, actualSrc, ok := v.Unit(src)
	if !ok || convoyer.Type != godip.Fleet {
		return
	}
	if v.Graph().Flags(actualSrc)[godip.Land] && !v.Graph().Flags(actualSrc)[godip.Convoyable] {
		return
	}
	if convoyer.Nation != nation {
		return
	}
	possibleSources := []godip.Province{}
	possibleDestinations := []godip.Province{}
	for _, endpointProv := range v.Graph().Provinces() {
		for _, endpoint := range v.Graph().Coasts(endpointProv) {
			if !v.Graph().Flags(endpoint)[godip.Land] {
				continue
			}
			if path := ConvoyParticipationPossible(v, actualSrc, endpoint); path != nil {
				possibleDestinations = append(possibleDestinations, endpoint)
				if endpointUnit, _, ok := v.Unit(endpoint); ok && endpointUnit.Type == godip.Army {
					possibleSources = append(possibleSources, endpoint)
				}
			}
		}
	}
	for _, src := range possibleSources {
		for _, dst := range possibleDestinations {
			if src.Super() == dst.Super() {
				continue
			}
			if result == nil {
				result = godip.Options{}
			}
			if result[godip.SrcProvince(actualSrc)] == nil {
				result[godip.SrcProvince(actualSrc)] = godip.Options{}
			}
			opt, f := result[godip.SrcProvince(actualSrc)][src]
			if !f {
				opt = godip.Options{}
				result[godip.SrcProvince(actualSrc)][src] = opt
			}
			opt[dst] = nil
		}
	}
	return
}

func (self *convoy) Validate(v godip.Validator) (godip.Nation, error) {
	if v.Phase().Type() != godip.Movement {
		return "", godip.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return "", godip.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return "", godip.ErrInvalidTarget
	}
	if !v.Graph().Has(self.targets[2]) {
		return "", godip.ErrInvalidTarget
	}
	for _, src := range v.Graph().Coasts(self.targets[0]) {
		if v.Graph().Flags(src)[godip.Land] && !v.Graph().Flags(src)[godip.Convoyable] {
			return "", godip.ErrIllegalConvoyPath
		}
	}
	var convoyer godip.Unit
	var ok bool
	convoyer, self.targets[0], ok = v.Unit(self.targets[0])
	if !ok {
		return "", godip.ErrMissingUnit
	} else if convoyer.Type != godip.Fleet {
		return "", godip.ErrIllegalConvoyer
	}
	var convoyee godip.Unit
	if convoyee, self.targets[1], ok = v.Unit(self.targets[1]); !ok {
		return "", godip.ErrMissingConvoyee
	} else if convoyee.Type != godip.Army {
		return "", godip.ErrIllegalConvoyee
	}
	if AnyConvoyPath(v, self.targets[1], self.targets[2], false, nil) == nil {
		return "", godip.ErrIllegalConvoyMove
	}
	return convoyer.Nation, nil
}

func (self *convoy) Execute(state godip.State) {
}

func ConvoyDestinations(v godip.Validator, src godip.Province, noConvoy *godip.Province) (result []godip.Province) {
	defer v.Profile("ConvoyDestinations", time.Now())
	potentialConvoyCoasts := map[godip.Province]bool{}
	v.Graph().Path(src, "-", func(prov godip.Province, edgeFlags, provFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) (okStep bool) {
		if !edgeFlags[godip.Sea] {
			return false
		}
		if v.Graph().Flags(prov.Super())[godip.Land] {
			if len(trace) > 0 {
				potentialConvoyCoasts[prov] = true
			}
			if !provFlags[godip.Convoyable] {
				return false
			}
		}
		if noConvoy != nil && *noConvoy == prov {
			return false
		}
		unit, _, found := v.Unit(prov)
		if !found {
			return false
		}
		if unit.Type != godip.Fleet {
			return false
		}
		return true
	})
	result = make([]godip.Province, 0, len(potentialConvoyCoasts))
	for prov := range potentialConvoyCoasts {
		result = append(result, prov)
	}
	return result
}

func ConvoySources(v godip.Validator, dst godip.Province, noConvoy *godip.Province) (result []godip.Province) {
	defer v.Profile("ConvoyDestinations", time.Now())
	potentialConvoyCoasts := map[godip.Province]bool{}
	v.Graph().ReversePath("-", dst, func(prov godip.Province, edgeFlags, provFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) (okStep bool) {
		if !edgeFlags[godip.Sea] {
			return false
		}
		if v.Graph().Flags(prov.Super())[godip.Land] {
			if len(trace) > 0 {
				potentialConvoyCoasts[prov] = true
			}
			if !provFlags[godip.Convoyable] {
				return false
			}
		}
		if noConvoy != nil && *noConvoy == prov {
			return false
		}
		unit, _, found := v.Unit(prov)
		if !found {
			return false
		}
		if unit.Type != godip.Fleet {
			return false
		}
		return true
	})
	result = make([]godip.Province, 0, len(potentialConvoyCoasts))
	for prov := range potentialConvoyCoasts {
		result = append(result, prov)
	}
	return result
}

// PossibleConvoyPathFilter returns a path filter for Graph that only accepts nodes that can partake in a convoy from
// src to dst. If resolveConvoys, then the convoys have to be successful. If dstOk then the dst is acceptable as convoying
// node.
func PossibleConvoyPathFilter(v godip.Validator, src, dst godip.Province, resolveConvoys, dstOk bool) godip.PathFilter {
	return func(name godip.Province, edgeFlags, nodeFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) bool {
		if dstOk && name.Contains(dst) && nodeFlags[godip.Land] {
			return true
		}
		if (nodeFlags[godip.Land] || !nodeFlags[godip.Sea]) && !nodeFlags[godip.Convoyable] {
			return false
		}
		if u, _, ok := v.Unit(name); ok && u.Type == godip.Fleet {
			if !resolveConvoys {
				return true
			}
			if order, prov, ok := v.Order(name); ok && order.Type() == godip.Convoy && order.Targets()[1].Contains(src) && order.Targets()[2].Contains(dst) {
				if err := v.(godip.Resolver).Resolve(prov); err != nil {
					return false
				}
				return true
			}
		}
		return false
	}
}

// ConvoyParticipantionPossible returns a path that participant (assumed to be a fleet at a convoyable position)
// could send an army to endpoint.
func ConvoyParticipationPossible(v godip.Validator, participant, endpoint godip.Province) []godip.Province {
	defer v.Profile("ConvoyParticipationPossible", time.Now())
	return v.Graph().Path(participant, endpoint, PossibleConvoyPathFilter(v, participant, endpoint, false, true))
}

func ConvoyPathPossibleVia(v godip.Validator, via, src, dst godip.Province, resolveConvoys bool) []godip.Province {
	defer v.Profile("ConvoyPathPossibleVia", time.Now())
	if part1 := v.Graph().Path(src, via, PossibleConvoyPathFilter(v, src, dst, resolveConvoys, false)); part1 != nil {
		t2 := time.Now()
		if part2 := v.Graph().Path(via, dst, PossibleConvoyPathFilter(v, src, dst, resolveConvoys, true)); part2 != nil {
			return append(part1, part2...)
		}
		v.Profile("ConvoyPathPossbleVia { [ check second half ] }", t2)
	}
	return nil
}

func convoyPath(v godip.Validator, src, dst godip.Province, resolveConvoys bool, viaNation *godip.Nation) []godip.Province {
	defer v.Profile("convoyPath", time.Now())
	if src == dst {
		return nil
	}
	// Find all fleets that could or will convoy.
	t := time.Now()
	waypoints, _, _ := v.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		//  (not on land               or is convoyable)                 and exists  and is the viaNation, if provided               and is a fleet     and is not _at_ src or dst.
		if (!v.Graph().Flags(p)[godip.Land] || v.Graph().Flags(p)[godip.Convoyable]) && u != nil && (viaNation == nil || u.Nation == *viaNation) && u.Type == godip.Fleet && p.Super() != src.Super() && p.Super() != dst.Super() {
			if !resolveConvoys {
				if viaNation == nil || (o != nil && o.Type() == godip.Convoy && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst)) {
					return true
				}
				return false
			}
			if o != nil && o.Type() == godip.Convoy && o.Targets()[1].Contains(src) && o.Targets()[2].Contains(dst) {
				if err := v.(godip.Resolver).Resolve(p); err != nil {
					return false
				}
				return true
			}
		}
		return false
	})
	v.Profile("convoyPath { v.Find([matching fleets]) }", t)
	for _, waypoint := range waypoints {
		if path := ConvoyPathPossibleVia(v, waypoint, src, dst, resolveConvoys); path != nil {
			return path
		}
	}
	return nil
}

func MustConvoy(r godip.Resolver, src godip.Province) bool {
	defer r.Profile("MustConvoy", time.Now())
	unit, _, ok := r.Unit(src)
	if !ok {
		return false
	}
	if unit.Type != godip.Army {
		return false
	}
	order, _, ok := r.Order(src)
	if !ok {
		return false
	}
	if order.Type() != godip.Move {
		return false
	}
	return (!HasEdge(r, unit.Type, order.Targets()[0], order.Targets()[1]) ||
		(order.Flags()[godip.ViaConvoy] && AnyConvoyPath(r, order.Targets()[0], order.Targets()[1], true, nil) != nil) ||
		AnyConvoyPath(r, order.Targets()[0], order.Targets()[1], false, &unit.Nation) != nil)
}

func AnyConvoyPath(v godip.Validator, src, dst godip.Province, resolveConvoys bool, viaNation *godip.Nation) (result []godip.Province) {
	defer v.Profile("AnyConvoyPath", time.Now())
	if !v.Graph().AllFlags(src)[godip.Sea] || !v.Graph().AllFlags(dst)[godip.Sea] {
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
