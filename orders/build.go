package orders

import (
	"fmt"
	"sort"
	"time"

	"github.com/zond/godip"
)

var BuildOrder = &build{}

var BuildAnywhereOrder = &build{
	flags: map[godip.Flag]bool{
		godip.Anywhere: true,
	},
}

var BuildAnyHomeCenterOrder = &build{
	flags: map[godip.Flag]bool{
		godip.AnyHomeCenter: true,
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
			godip.Anywhere: true,
		},
	}
}

func BuildAnyHomeCenter(source godip.Province, typ godip.UnitType, at time.Time) *build {
	return &build{
		targets: []godip.Province{source},
		typ:     typ,
		at:      at,
		flags: map[godip.Flag]bool{
			godip.AnyHomeCenter: true,
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
	return godip.Build
}

func (self *build) DisplayType() godip.OrderType {
	return godip.Build
}

func (self *build) Flags() map[godip.Flag]bool {
	return self.flags
}

func (self *build) String() string {
	return fmt.Sprintf("%v %v %v", self.targets[0], godip.Build, self.typ)
}

func (self *build) Targets() []godip.Province {
	return self.targets
}

func (self *build) At() time.Time {
	return self.at
}

func (self *build) Adjudicate(r godip.Resolver) error {
	me, _, ok := r.SupplyCenter(self.targets[0].Super())
	if !ok {
		me = godip.Neutral
	}
	builds, _, _ := AdjustmentStatus(r, me)
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
			if self.flags[godip.Anywhere] {
				result = BuildAnywhere(godip.Province(bits[0]), godip.UnitType(bits[2]), time.Now())
			} else if self.flags[godip.AnyHomeCenter] {
				result = BuildAnyHomeCenter(godip.Province(bits[0]), godip.UnitType(bits[2]), time.Now())
			} else {
				result = Build(godip.Province(bits[0]), godip.UnitType(bits[2]), time.Now())
			}
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
	if v.Phase().Type() != godip.Adjustment {
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
	if !self.flags[godip.Anywhere] {
		owner := v.Graph().SC(src.Super())
		if owner == nil || (!self.flags[godip.AnyHomeCenter] && *owner != me) || *owner == godip.Neutral {
			return
		}
	}
	if _, _, ok = v.Unit(src); ok {
		return
	}
	var wrapperFunc func(godip.OptionValue) godip.FilteredOptionValue
	if _, _, balance := AdjustmentStatus(v, me); balance < 1 {
		return
	} else {
		wrapperFunc = func(val godip.OptionValue) godip.FilteredOptionValue {
			return godip.FilteredOptionValue{
				Filter: fmt.Sprintf("MAX:%v:%v", godip.Build, balance),
				Value:  val,
			}
		}
	}
	if v.Graph().Flags(src)[godip.Land] || v.Graph().Flags(src.Super())[godip.Land] {
		if result == nil {
			result = godip.Options{}
		}
		if result[wrapperFunc(godip.Army)] == nil {
			result[wrapperFunc(godip.Army)] = godip.Options{}
		}
		result[wrapperFunc(godip.Army)][godip.SrcProvince(src.Super())] = nil
	}
	if v.Graph().Flags(src)[godip.Sea] || v.Graph().Flags(src.Super())[godip.Sea] {
		if result == nil {
			result = godip.Options{}
		}
		if result[wrapperFunc(godip.Fleet)] == nil {
			result[wrapperFunc(godip.Fleet)] = godip.Options{}
		}
		result[wrapperFunc(godip.Fleet)][godip.SrcProvince(src)] = nil
	}
	return
}

func (self *build) Validate(v godip.Validator) (godip.Nation, error) {
	// right phase type
	if v.Phase().Type() != godip.Adjustment {
		return "", godip.ErrInvalidPhase
	}
	// does someone own this
	var me godip.Nation
	var ok bool
	if me, _, ok = v.SupplyCenter(self.targets[0]); !ok {
		return "", godip.ErrMissingSupplyCenter
	}
	if !self.flags[godip.Anywhere] {
		// is there a home sc here
		owner := v.Graph().SC(self.targets[0].Super())
		if owner == nil {
			return "", fmt.Errorf("Should be SOME owner of %v", self.targets[0])
		} else if (!self.flags[godip.AnyHomeCenter] && *owner != me) || *owner == godip.Neutral {
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
	if _, _, balance := AdjustmentStatus(v, me); balance < 1 {
		return "", godip.ErrMissingSurplus
	}
	// can i build THIS here
	if self.typ == godip.Army && !v.Graph().Flags(self.targets[0])[godip.Land] {
		return "", godip.ErrIllegalUnitType
	}
	if self.typ == godip.Fleet && !v.Graph().Flags(self.targets[0])[godip.Sea] {
		return "", godip.ErrIllegalUnitType
	}
	return me, nil
}

func (self *build) Execute(state godip.State) {
	me, ok := state.SupplyCenters()[self.targets[0].Super()]
	if !ok {
		me = godip.Neutral
	}
	state.SetUnit(self.targets[0], godip.Unit{self.typ, me})
}

func AdjustmentStatus(v godip.Validator, me godip.Nation) (builds godip.Orders, disbands godip.Orders, balance int) {
	scs := 0
	for _, prov := range v.Graph().AllSCs() {
		nat, _, ok := v.SupplyCenter(prov)
		if !ok {
			nat = godip.Neutral
		}
		if nat == me {
			scs += 1
			if order, _, ok := v.Order(prov); ok && order.Type() == godip.Build {
				builds = append(builds, order)
			}
		}
	}

	units := 0
	for prov, unit := range v.Units() {
		if unit.Nation == me {
			units += 1
			if order, _, ok := v.Order(prov); ok && order.Type() == godip.Disband {
				disbands = append(disbands, order)
			}
		}
	}

	sort.Sort(builds)
	sort.Sort(disbands)

	balance = scs - units
	if balance > 0 {
		disbands = nil
		builds = builds[:godip.Max(0, godip.Min(len(builds), balance))]
	} else if balance < 0 {
		builds = nil
		disbands = disbands[:godip.Max(0, godip.Min(len(disbands), -balance))]
	} else {
		builds = nil
		disbands = nil
	}

	return
}
