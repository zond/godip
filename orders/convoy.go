package orders

import (
	"fmt"
	"time"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

var ConvoyOrder = &convoy{}

func Convoy(source, from, to dip.Province) *convoy {
	return &convoy{
		targets: []dip.Province{source, from, to},
	}
}

type convoy struct {
	targets []dip.Province
}

func (self *convoy) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], cla.Convoy, self.targets[1:])
}

func (self *convoy) Flags() map[dip.Flag]bool {
	return nil
}

func (self *convoy) At() time.Time {
	return time.Now()
}

func (self *convoy) Type() dip.OrderType {
	return cla.Convoy
}

func (self *convoy) DisplayType() dip.OrderType {
	return cla.Convoy
}

func (self *convoy) Targets() []dip.Province {
	return self.targets
}

func (self *convoy) Adjudicate(r dip.Resolver) error {
	unit, _, _ := r.Unit(self.targets[0])
	if breaks, _, _ := r.Find(func(p dip.Province, o dip.Order, u *dip.Unit) bool {
		return (o.Type() == cla.Move && // move
			o.Targets()[1] == self.targets[0] && // against us
			u.Nation != unit.Nation && // not friendly
			r.Resolve(p) == nil)
	}); len(breaks) > 0 {
		return cla.ErrConvoyDislodged{breaks[0]}
	}
	return nil
}

func (self *convoy) Parse(bits []string) (dip.Adjudicator, error) {
	var result dip.Adjudicator
	var err error
	if len(bits) > 1 && dip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 4 {
			result = Convoy(dip.Province(bits[0]), dip.Province(bits[2]), dip.Province(bits[3]))
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *convoy) Options(v dip.Validator, nation dip.Nation, src dip.Province) (result dip.Options) {
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
	possibleSources := []dip.Province{}
	possibleDestinations := []dip.Province{}
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
				result = dip.Options{}
			}
			if result[dip.SrcProvince(actualSrc)] == nil {
				result[dip.SrcProvince(actualSrc)] = dip.Options{}
			}
			opt, f := result[dip.SrcProvince(actualSrc)][src]
			if !f {
				opt = dip.Options{}
				result[dip.SrcProvince(actualSrc)][src] = opt
			}
			opt[dst] = nil
		}
	}
	return
}

func (self *convoy) Validate(v dip.Validator) (dip.Nation, error) {
	if v.Phase().Type() != cla.Movement {
		return "", cla.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return "", cla.ErrInvalidSource
	}
	if !v.Graph().Has(self.targets[1]) {
		return "", cla.ErrInvalidTarget
	}
	if !v.Graph().Has(self.targets[2]) {
		return "", cla.ErrInvalidTarget
	}
	for _, src := range v.Graph().Coasts(self.targets[0]) {
		if v.Graph().Flags(src)[cla.Land] && !v.Graph().Flags(src)[cla.Convoyable] {
			return "", cla.ErrIllegalConvoyPath
		}
	}
	var convoyer dip.Unit
	var ok bool
	convoyer, self.targets[0], ok = v.Unit(self.targets[0])
	if !ok {
		return "", cla.ErrMissingUnit
	} else if convoyer.Type != cla.Fleet {
		return "", cla.ErrIllegalConvoyer
	}
	var convoyee dip.Unit
	if convoyee, self.targets[1], ok = v.Unit(self.targets[1]); !ok {
		return "", cla.ErrMissingConvoyee
	} else if convoyee.Type != cla.Army {
		return "", cla.ErrIllegalConvoyee
	}
	if cla.AnyConvoyPath(v, self.targets[1], self.targets[2], false, nil) == nil {
		return "", cla.ErrIllegalConvoyMove
	}
	return convoyer.Nation, nil
}

func (self *convoy) Execute(state dip.State) {
}
