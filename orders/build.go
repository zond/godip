package orders

import (
	"fmt"
	"time"

	"github.com/zond/godip"
	cla "github.com/zond/godip/variants/classical/common"
)

var BuildOrder = &build{}

var BuildAnywhereOrder = &build{
	flags: map[godip.Flag]bool{
		cla.Anywhere: true,
	},
}

func Build(source godip.Province, typ godip.UnitType, at time.Time) *build {
	return &build{
		targets: []godip.Province{source},
		typ:     typ,
		at:      at,
	}
}

func BuildAnywhere(source godip.Province, typ godip.UnitType, at time.Time) *build {
	return &build{
		targets: []godip.Province{source},
		typ:     typ,
		at:      at,
		flags: map[godip.Flag]bool{
			cla.Anywhere: true,
		},
	}
}

type build struct {
	targets []godip.Province
	typ     godip.UnitType
	at      time.Time
	flags   map[godip.Flag]bool
}

func (self *build) Type() godip.OrderType {
	return cla.Build
}

func (self *build) DisplayType() godip.OrderType {
	return cla.Build
}

func (self *build) Flags() map[godip.Flag]bool {
	return nil
}

func (self *build) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], cla.Build, self.typ)
}

func (self *build) Targets() []godip.Province {
	return self.targets
}

func (self *build) At() time.Time {
	return self.at
}

func (self *build) Adjudicate(r godip.Resolver) error {
	me := r.SupplyCenters()[self.targets[0].Super()]
	builds, _, _ := cla.AdjustmentStatus(r, me)
	if len(builds) == 0 || self.at.After(builds[len(builds)-1].At()) {
		return godip.ErrIllegalBuild
	}
	return nil
}

func (self *build) Parse(bits []string) (godip.Adjudicator, error) {
	var result godip.Adjudicator
	var err error
	if len(bits) > 1 && godip.OrderType(bits[1]) == self.DisplayType() {
		if len(bits) == 3 {
			result = Build(godip.Province(bits[0]), godip.UnitType(bits[2]), time.Now())
		}
		if result == nil {
			err = fmt.Errorf("Can't parse as %+v", bits)
		}
	}
	return result, err
}

func (self *build) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	if len(v.Graph().Coasts(src)) > 1 && src == src.Super() {
		return
	}
	if v.Phase().Type() != cla.Adjustment {
		return
	}
	// To avoid having build order for a coast and the main province at the same time...
	otherOrders := 0
	for _, prov := range v.Graph().Coasts(src) {
		if _, foundProv, ok := v.Order(prov); ok && foundProv == src {
			otherOrders += 1
		}
	}
	if otherOrders > 0 {
		return
	}
	me, _, ok := v.SupplyCenter(src)
	if !ok {
		return
	}
	if nation != me {
		return
	}
	if !self.flags[cla.Anywhere] {
		owner := v.Graph().SC(src.Super())
		if owner == nil || *owner != me {
			return
		}
	}
	if _, _, ok = v.Unit(src); ok {
		return
	}
	if _, _, balance := cla.AdjustmentStatus(v, me); balance < 1 {
		return
	}
	if v.Graph().Flags(src)[cla.Land] || v.Graph().Flags(src.Super())[cla.Land] {
		if result == nil {
			result = godip.Options{}
		}
		if result[cla.Army] == nil {
			result[cla.Army] = godip.Options{}
		}
		result[cla.Army][godip.SrcProvince(src.Super())] = nil
	}
	if v.Graph().Flags(src)[cla.Sea] || v.Graph().Flags(src.Super())[cla.Sea] {
		if result == nil {
			result = godip.Options{}
		}
		if result[cla.Fleet] == nil {
			result[cla.Fleet] = godip.Options{}
		}
		result[cla.Fleet][godip.SrcProvince(src)] = nil
	}
	return
}

func (self *build) Validate(v godip.Validator) (godip.Nation, error) {
	// right phase type
	if v.Phase().Type() != cla.Adjustment {
		return "", godip.ErrInvalidPhase
	}
	// does someone own this
	var me godip.Nation
	var ok bool
	if me, _, ok = v.SupplyCenter(self.targets[0]); !ok {
		return "", godip.ErrMissingSupplyCenter
	}
	if !self.flags[cla.Anywhere] {
		// is there a home sc here
		owner := v.Graph().SC(self.targets[0].Super())
		if owner == nil {
			return "", fmt.Errorf("Should be SOME owner of %v", self.targets[0])
		} else if *owner != me {
			return "", godip.ErrHostileSupplyCenter
		}
	}
	// is there a unit here
	if _, _, ok := v.Unit(self.targets[0]); ok {
		return "", godip.ErrOccupiedSupplyCenter
	}
	// is there another build order here
	for _, prov := range v.Graph().Coasts(self.targets[0]) {
		if other, foundProv, ok := v.Order(prov); ok && foundProv == prov && other != self {
			return "", godip.ErrDoubleBuild{
				Provinces: []godip.Province{prov, foundProv},
			}
		}
	}
	// can i build
	if _, _, balance := cla.AdjustmentStatus(v, me); balance < 1 {
		return "", godip.ErrMissingSurplus
	}
	// can i build THIS here
	if self.typ == cla.Army && !v.Graph().Flags(self.targets[0])[cla.Land] {
		return "", godip.ErrIllegalUnitType
	}
	if self.typ == cla.Fleet && !v.Graph().Flags(self.targets[0])[cla.Sea] {
		return "", godip.ErrIllegalUnitType
	}
	return me, nil
}

func (self *build) Execute(state godip.State) {
	me := state.SupplyCenters()[self.targets[0].Super()]
	state.SetUnit(self.targets[0], godip.Unit{self.typ, me})
}
