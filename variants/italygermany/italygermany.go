package italygermany

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var ItalyGermanyVariant = common.Variant{
	Name: "Italy vs Germany",
	Graph: func() godip.Graph {
		okNations := map[godip.Nation]bool{
			godip.Italy:   true,
			godip.Germany: true,
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
			"nap": godip.Unit{godip.Fleet, godip.Italy},
			"rom": godip.Unit{godip.Army, godip.Italy},
			"ven": godip.Unit{godip.Army, godip.Italy},
			"kie": godip.Unit{godip.Fleet, godip.Germany},
			"ber": godip.Unit{godip.Army, godip.Germany},
			"mun": godip.Unit{godip.Army, godip.Germany},
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[godip.Province]godip.Nation{
			"nap": godip.Italy,
			"rom": godip.Italy,
			"ven": godip.Italy,
			"kie": godip.Germany,
			"ber": godip.Germany,
			"mun": godip.Germany,
		})
		return
	},
	Blank:   classical.Blank,
	Phase:   classical.NewPhase,
	Parser:  classical.Parser,
	Nations: []godip.Nation{godip.Germany, godip.Italy},
	NationColors: map[godip.Nation]string{
		godip.Italy:   "#4CAF50",
		godip.Germany: "#212121",
	},
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(18),
	ProvinceLongNames: classical.ClassicalVariant.ProvinceLongNames,
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion:  "1",
	SVGUnits:    classical.SVGUnits,
	CreatedBy:   "",
	Version:     "",
	Description: "A two player variant on the classical map.",
	SoloSCCount: func(*state.State) int { return 18 },
	Rules: `The first to 18 supply centers is the winner. 
The game only has two nations: Italy and Germany.`,
}
