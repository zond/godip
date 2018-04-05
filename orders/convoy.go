package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"

	cla "github.com/zond/godip/variants/classical/common"
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
	return fmt.Sprintf("%v %v %v", self.targets[0], cla.Convoy, self.targets[1:])
}

func (self *convoy) Flags() map[godip.Flag]bool {
	return nil
}

func (self *convoy) At() time.Time {
	return time.Now()
}

func (self *convoy) Type() godip.OrderType {
	return cla.Convoy
}

func (self *convoy) DisplayType() godip.OrderType {
	return cla.Convoy
}

func (self *convoy) Targets() []godip.Province {
	return self.targets
}

func (self *convoy) Adjudicate(r godip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	if breaks, _, _ := r.Find(func(p godip.Province, o godip.Order, u *godip.Unit) bool {
		return (o.Type() == cla.Move && // move
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
	if v.Phase().Type() != cla.Movement {
		return
	}
	if !v.Graph().Has(src) {
		return
	}
	convoyer, actualSrc, ok := v.Unit(src)
	if !ok || convoyer.Type != cla.Fleet {
		return
	}
	if v.Graph().Flags(actualSrc)[cla.Land] && !v.Graph().Flags(actualSrc)[cla.Convoyable] {
		return
	}
	if convoyer.Nation != nation {
		return
	}
	possibleSources := []godip.Province{}
	possibleDestinations := []godip.Province{}
	for _, endpointProv := range v.Graph().Provinces() {
		for _, endpoint := range v.Graph().Coasts(endpointProv) {
			if !v.Graph().Flags(endpoint)[cla.Land] {
				continue
			}
			if path := cla.ConvoyParticipationPossible(v, actualSrc, endpoint); path != nil {
				possibleDestinations = append(possibleDestinations, endpoint)
				if endpointUnit, _, ok := v.Unit(endpoint); ok && endpointUnit.Type == cla.Army {
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
	if v.Phase().Type() != cla.Movement {
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
		if v.Graph().Flags(src)[cla.Land] && !v.Graph().Flags(src)[cla.Convoyable] {
			return "", godip.ErrIllegalConvoyPath
		}
	}
	var convoyer godip.Unit
	var ok bool
	convoyer, self.targets[0], ok = v.Unit(self.targets[0])
	if !ok {
		return "", godip.ErrMissingUnit
	} else if convoyer.Type != cla.Fleet {
		return "", godip.ErrIllegalConvoyer
	}
	var convoyee godip.Unit
	if convoyee, self.targets[1], ok = v.Unit(self.targets[1]); !ok {
		return "", godip.ErrMissingConvoyee
	} else if convoyee.Type != cla.Army {
		return "", godip.ErrIllegalConvoyee
	}
	if cla.AnyConvoyPath(v, self.targets[1], self.targets[2], false, nil) == nil {
		return "", godip.ErrIllegalConvoyMove
	}
	return convoyer.Nation, nil
}

func (self *convoy) Execute(state godip.State) {
}
