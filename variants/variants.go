package variants

import (
	"github.com/zond/godip/variants/ancientmediterranean"
	"github.com/zond/godip/variants/chaos"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/coldwar"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/fleetrome"
	"github.com/zond/godip/variants/franceaustria"
	"github.com/zond/godip/variants/hundred"
	"github.com/zond/godip/variants/pure"
	"github.com/zond/godip/variants/twentytwenty"
	"github.com/zond/godip/variants/westernworld901"
	"github.com/zond/godip/variants/youngstownredux"
)

func init() {
	for _, variant := range OrderedVariants {
		Variants[variant.Name] = variant
	}
}

var Variants = map[string]common.Variant{}

var OrderedVariants = []common.Variant{
	classical.ClassicalVariant,
	coldwar.ColdWarVariant,
	chaos.ChaosVariant,
	fleetrome.FleetRomeVariant,
	franceaustria.FranceAustriaVariant,
	hundred.HundredVariant,
	pure.PureVariant,
	ancientmediterranean.AncientMediterraneanVariant,
	twentytwenty.TwentyTwentyVariant,
	westernworld901.WesternWorld901Variant,
	youngstownredux.YoungstownReduxVariant,
}
