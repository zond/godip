package variants

import (
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/fleetrome"
	"github.com/zond/godip/variants/franceaustria"
	"github.com/zond/godip/variants/pure"
)

func init() {
	for _, variant := range OrderedVariants {
		Variants[variant.Name] = variant
	}
}

var Variants = map[string]common.Variant{}

var OrderedVariants = []common.Variant{
	classical.ClassicalVariant,
	fleetrome.FleetRomeVariant,
	franceaustria.FranceAustriaVariant,
	pure.PureVariant,
}
