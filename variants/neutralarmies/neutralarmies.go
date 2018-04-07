package neutralarmies

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var NeutralArmiesVariant = common.Variant{
	Name:  "Neutral Armies",
	Graph: func() godip.Graph { return start.Graph() },
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		for _, sc := range start.Graph().SCs(godip.Neutral) {
			if err = result.SetUnit(godip.Province(sc), godip.Unit{
				Type:   godip.Army,
				Nation: godip.Neutral,
			}); err != nil {
				return
			}
		}
		return
	},
	Blank:      classical.Blank,
	Phase:      classical.Phase,
	Parser:     orders.ClassicalParser,
	Nations:    classical.Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
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
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules:       "",
}
