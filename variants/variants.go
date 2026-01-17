package variants

import (
	"github.com/zond/godip/variants/ancientmediterranean"
	"github.com/zond/godip/variants/beta/atlanticcolonies"
	"github.com/zond/godip/variants/beta/gatewaywest"
	"github.com/zond/godip/variants/beta/threekingdoms"
	"github.com/zond/godip/variants/canton"
	"github.com/zond/godip/variants/chaos"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classicalcrowded"
	"github.com/zond/godip/variants/coldwar"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/empiresandcoalitions"
	"github.com/zond/godip/variants/europe1939"
	"github.com/zond/godip/variants/fleetrome"
	"github.com/zond/godip/variants/franceaustria"
	"github.com/zond/godip/variants/hundred"
	"github.com/zond/godip/variants/italygermany"
	"github.com/zond/godip/variants/northseawars"
	"github.com/zond/godip/variants/pure"
	"github.com/zond/godip/variants/sengoku"
	"github.com/zond/godip/variants/twentytwenty"
	"github.com/zond/godip/variants/unconstitutional"
	"github.com/zond/godip/variants/vietnamwar"
	"github.com/zond/godip/variants/westernworld901"
	"github.com/zond/godip/variants/year1908"
	"github.com/zond/godip/variants/youngstownredux"
)

func init() {
	for _, variant := range OrderedVariants {
		Variants[variant.Name] = variant
	}
}

var Variants = map[string]common.Variant{}

var OrderedVariants = []common.Variant{
	atlanticcolonies.AtlanticColoniesVariant,
	gatewaywest.GatewayWestVariant,
	classicalcrowded.ClassicalCrowdedVariant,
	threekingdoms.ThreeKingdomsVariant,
	ancientmediterranean.AncientMediterraneanVariant,
	canton.CantonVariant,
	chaos.ChaosVariant,
	classical.ClassicalVariant,
	coldwar.ColdWarVariant,
	empiresandcoalitions.EmpiresAndCoalitionsVariant,
	europe1939.Europe1939Variant,
	fleetrome.FleetRomeVariant,
	franceaustria.FranceAustriaVariant,
	hundred.HundredVariant,
	italygermany.ItalyGermanyVariant,
	northseawars.NorthSeaWarsVariant,
	pure.PureVariant,
	sengoku.SengokuVariant,
	twentytwenty.TwentyTwentyVariant,
	unconstitutional.UnconstitutionalVariant,
	vietnamwar.VietnamWarVariant,
	westernworld901.WesternWorld901Variant,
	year1908.Year1908Variant,
	youngstownredux.YoungstownReduxVariant,
}
