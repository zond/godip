package fleetrome

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/variants/classical/common"
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
	Phase:      classical.Phase,
	Parser:     orders.ClassicalParser,
	Nations:    cla.Nations,
	PhaseTypes: cla.PhaseTypes,
	Seasons:    cla.Seasons,
	UnitTypes:  cla.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion: "1482957154",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Richard Sharp",
	Version:     "",
	Description: "Classical Diplomacy, but Italy starts with a fleet in Rome.",
	Rules:       "The first to 18 supply centers is the winner.  Rules are as per classical Diplomacy, but Italy starts with a fleet in Rome rather than an army.",
}
