package franceaustria

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var FranceAustriaVariant = common.Variant{
	Name: "France vs Austria",
	Graph: func() godip.Graph {
		okNations := map[godip.Nation]bool{
			godip.France:  true,
			godip.Austria: true,
			godip.Neutral: true,
		}
		neutral := godip.Neutral
		result := start.Graph()
		for _, node := range result.Nodes {
			if node.SC != nil && !okNations[*node.SC] {
				node.SC = &neutral
			}
		}
		return result
	},
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		if err = result.SetUnits(map[godip.Province]godip.Unit{
			"bre": godip.Unit{godip.Fleet, godip.France},
			"par": godip.Unit{godip.Army, godip.France},
			"mar": godip.Unit{godip.Army, godip.France},
			"tri": godip.Unit{godip.Fleet, godip.Austria},
			"vie": godip.Unit{godip.Army, godip.Austria},
			"bud": godip.Unit{godip.Army, godip.Austria},
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[godip.Province]godip.Nation{
			"bre": godip.France,
			"par": godip.France,
			"mar": godip.France,
			"tri": godip.Austria,
			"vie": godip.Austria,
			"bud": godip.Austria,
		})
		return
	},
	Blank:      classical.Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    []godip.Nation{godip.Austria, godip.France},
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion: "1",
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
	Description: "A two player variant on the classical map.",
	Rules:       "The first to 18 supply centers is the winner. The rules are as per classical Diplomacy, but with only France and Austria.",
}
