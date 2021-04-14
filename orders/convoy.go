package orders

import (
	"fmt"
	"log"
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

func (self *convoy) Corroborate(v godip.Validator) []error {
	unit, _, _ := v.Unit(self.targets[0])
	me := unit.Nation
	targetUnit, _, _ := v.Unit(self.targets[1])
	if targetUnit.Nation == me {
		potentialInconsistencies := []error{godip.InconsistencyMismatchedConvoyer{
			Convoyee: self.targets[1].Super(),
		}}
		ord, _, found := v.Order(self.targets[1].Super())
		if !found {
			return potentialInconsistencies
		}
		if ord.Type() != godip.Move {
			return potentialInconsistencies
		}
		if len(ord.Targets()) != 2 || ord.Targets()[0].Super() != self.targets[1].Super() || ord.Targets()[1].Super() != self.targets[2].Super() {
			return potentialInconsistencies
		}
	}
	return nil
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
	if v.Graph().Flags(actualSrc.Super())[godip.Land] && !v.Graph().Flags(actualSrc.Super())[godip.Convoyable] {
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
			if path := (ConvoyPathFinder{
				ConvoyPathFilter: ConvoyPathFilter{
					Validator:   v,
					Source:      actualSrc,
					Destination: endpoint,
				},
			}).Path(); path != nil {
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
	if len((ConvoyPathFinder{
		ConvoyPathFilter: ConvoyPathFilter{
			Validator:              v,
			Source:                 self.targets[1],
			Destination:            self.targets[2],
			MinLengthAtDestination: 1,
		}}).Any()) < 2 {
		return "", godip.ErrIllegalConvoyMove
	}
	return convoyer.Nation, nil
}

func (self *convoy) Execute(state godip.State) {
}

// ConvoyEndPoints returns all possible end points for a convoy route starting at
// startPoint. If reverse is false then the route will be for a unit moving from
// startPoint; if it's true then the route will be for a unit moving to startPoint
// instead. If noConvoy is set then the route cannot pass through it.
func ConvoyEndPoints(v godip.Validator, startPoint godip.Province, reverse bool, noConvoy *godip.Province) (result []godip.Province) {
	potentialConvoyCoasts := map[godip.Province]bool{}
	v.Graph().Path(startPoint, "-", reverse, func(prov godip.Province, edgeFlags, provFlags map[godip.Flag]bool, sc *godip.Nation, trace []godip.Province) (okStep bool) {
		if !edgeFlags[godip.Sea] {
			return false
		}
		if v.Graph().Flags(prov.Super())[godip.Land] {
			if len(trace) > 0 && prov.Super() != startPoint.Super() {
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

type ConvoyPathFilter struct {
	Validator   godip.Validator
	Source      godip.Province
	Destination godip.Province
	// If true, convoys will be resolved using Validator (cast to a Resolver)
	// to be considered OK.
	// Used during adjudication.
	ResolveConvoys bool
	// If true, destination will be checked the same way as every other step.
	// Used when validating convoy path _via_ provinces.
	DestinationMustConvoy bool
	// If not DestinationMustConvoy, i.e. destination is the landing point for the
	// convoyed unit, then the path up to the destination has to be at least
	// MinLengthAtDestination long for the step to destination to be OK.
	// Used to avoid finding convoy paths without any convoying fleets.
	MinLengthAtDestination int
	// If not nil, all fleets along the path have to be of this nation.
	// Used when warning about mismatched orders.
	OnlyNation *godip.Nation
	// If !ResolveConvoys, but VerifyConvoyOrderes, the participating fleets
	// need to at least give the convoy order to be considered OK.
	VerifyConvoyOrders bool
	// If set, this province will not be considered.
	// Used to find convoy paths not using a fleet attacked by someone supported
	// by the convoy target province, which is used to check if a support
	// is cut by the convoyed unit or not.
	AvoidProvince *godip.Province
}

func (p ConvoyPathFilter) PathFilter(
	name godip.Province,
	edgeFlags,
	nodeFlags map[godip.Flag]bool,
	sc *godip.Nation,
	trace []godip.Province,
) bool {
	superFlags := p.Validator.Graph().Flags(name.Super())
	if !p.DestinationMustConvoy && len(trace) >= p.MinLengthAtDestination && name.Contains(p.Destination) && superFlags[godip.Land] {
		return true
	}
	if (superFlags[godip.Land] || !superFlags[godip.Sea]) && !superFlags[godip.Convoyable] {
		return false
	}
	if p.AvoidProvince != nil && name.Super() == p.AvoidProvince.Super() {
		return false
	}
	if u, _, ok := p.Validator.Unit(name); ok && u.Type == godip.Fleet && (p.OnlyNation == nil || u.Nation == *p.OnlyNation) {
		if !p.ResolveConvoys && !p.VerifyConvoyOrders {
			return true
		}
		if order, prov, ok := p.Validator.Order(name); ok &&
			order.Type() == godip.Convoy &&
			order.Targets()[1].Contains(p.Source) &&
			order.Targets()[2].Contains(p.Destination) {
			if !p.ResolveConvoys {
				return true
			}
			if err := p.Validator.(godip.Resolver).Resolve(prov); err != nil {
				return false
			}
			return true
		}
	}
	return false
}

type ConvoyPathFinder struct {
	ConvoyPathFilter
	ViaNation   *godip.Nation
	ViaProvince *godip.Province
}

func (c ConvoyPathFinder) Path() []godip.Province {
	if c.ViaNation != nil && c.ViaProvince != nil {
		log.Panicf("Should never call Path with both ViaNation and ViaProvince: %+v", c)
	}
	if c.Source == c.Destination {
		return nil
	}
	if c.ViaNation != nil {
		// Find all fleets that could or will convoy.
		waypoints, _, _ := c.Validator.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
			// (not on land or is convoyable)
			if (!c.Validator.Graph().Flags(p)[godip.Land] || c.Validator.Graph().Flags(p)[godip.Convoyable]) &&
				// and exists
				u != nil &&
				// and is the viaNation
				u.Nation == *c.ViaNation &&
				// and is a fleet
				u.Type == godip.Fleet &&
				// and is not _at_ src or dst
				p.Super() != c.Source.Super() &&
				p.Super() != c.Destination.Super() {
				if !c.ResolveConvoys {
					if o != nil && o.Type() == godip.Convoy && o.Targets()[1].Contains(c.Source) && o.Targets()[2].Contains(c.Destination) {
						return true
					}
					return false
				}
				if o != nil && o.Type() == godip.Convoy && o.Targets()[1].Contains(c.Source) && o.Targets()[2].Contains(c.Destination) {
					if err := c.Validator.(godip.Resolver).Resolve(p); err != nil {
						return false
					}
					return true
				}
			}
			return false
		})
		// Run ViaProvince for each found fleet.
		c.ViaNation = nil
		for _, waypoint := range waypoints {
			c.ViaProvince = &waypoint
			if path := c.Path(); path != nil {
				return path
			}
		}
		return nil
	} else if c.ViaProvince != nil {
		part1Filter := c.ConvoyPathFilter
		part1Filter.DestinationMustConvoy = true
		// Find any path from src to via.
		if part1 := c.Validator.Graph().Path(c.Source, *c.ViaProvince, false, part1Filter.PathFilter); len(part1) > 0 {
			// Find any path from via to dst.
			if part2 := c.Validator.Graph().Path(*c.ViaProvince, c.Destination, false, c.ConvoyPathFilter.PathFilter); len(part2) > 0 {
				return append(part1, part2...)
			}
		}
		return nil
	}
	rval := c.Validator.Graph().Path(c.Source, c.Destination, false, c.ConvoyPathFilter.PathFilter)
	return rval
}

func (c ConvoyPathFinder) Any() []godip.Province {
	if !c.Validator.Graph().AllFlags(c.Source)[godip.Sea] || !c.Validator.Graph().AllFlags(c.Destination)[godip.Sea] {
		return nil
	}
	for _, srcCoast := range c.Validator.Graph().Coasts(c.Source) {
		for _, dstCoast := range c.Validator.Graph().Coasts(c.Destination) {
			coastToCoast := c
			coastToCoast.Source = srcCoast
			coastToCoast.Destination = dstCoast
			if result := coastToCoast.Path(); len(result) > 1 {
				return result
			}
		}
	}
	return nil
}

// MustConvoy returns whether the unit at src must convoy.
// Used during adjudication to find mandatory convoy path, i.e. if there is no other option,
// or there is a convoy option and the move is via convoy, or the owner has told at least one fleet
// to convoy the unit that way.
func MustConvoy(r godip.Resolver, src godip.Province) bool {
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
	rval := (!HasEdge(r, unit.Type, order.Targets()[0], order.Targets()[1]) ||
		(order.Flags()[godip.ViaConvoy] && len((ConvoyPathFinder{
			ConvoyPathFilter: ConvoyPathFilter{
				Validator:              r,
				Source:                 order.Targets()[0],
				Destination:            order.Targets()[1],
				VerifyConvoyOrders:     true,
				MinLengthAtDestination: 1,
			}}).Any()) > 1) ||
		len((ConvoyPathFinder{
			ConvoyPathFilter: ConvoyPathFilter{
				Validator:   r,
				Source:      order.Targets()[0],
				Destination: order.Targets()[1],
			},
			ViaNation: &unit.Nation,
		}).Any()) > 1)
	return rval
}
