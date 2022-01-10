package fleetrome

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var FleetRomeVariant = common.Variant{
	Name:  "Fleet Rome",
	Graph: func() godip.Graph { return start.Graph() },
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		result.RemoveUnit(godip.Province("rom"))
		if err = result.SetUnit(godip.Province("rom"), godip.Unit{
			Type:   godip.Fleet,
			Nation: godip.Italy,
		}); err != nil {
			return
		}
		return
	},
	Blank:      classical.Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    classical.Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	ProvinceLongNames: classical.ClassicalVariant.ProvinceLongNames,
	SVGVersion:        "1",
	SVGUnits:          classical.SVGUnits,
	SVGFlags:          classical.SVGFlags,
	CreatedBy:         "Richard Sharp",
	Version:           "",
	Description:       "Classical Diplomacy, but Italy starts with a fleet in Rome.",
	SoloSCCount:       func(*state.State) int { return 18 },
	Rules: `The first to 18 supply centers is the winner.  
Italy starts with a fleet in Rome rather than an army.`,
}
