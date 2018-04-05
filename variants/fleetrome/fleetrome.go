package fleetrome

import (
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"

	dip "github.com/zond/godip"
	cla "github.com/zond/godip/variants/classical/common"
)

var FleetRomeVariant = common.Variant{
	Name:  "Fleet Rome",
	Graph: func() dip.Graph { return start.Graph() },
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		result.RemoveUnit(dip.Province("rom"))
		if err = result.SetUnit(dip.Province("rom"), dip.Unit{
			Type:   cla.Fleet,
			Nation: cla.Italy,
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
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Richard Sharp",
	Version:     "",
	Description: "Classical Diplomacy, but Italy starts with a fleet in Rome.",
	Rules:       "The first to 18 supply centers is the winner.  Rules are as per classical Diplomacy, but Italy starts with a fleet in Rome rather than an army.",
}
