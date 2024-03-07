package classicalneutralitaly

import (
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var ClassicalNeutralItalyVariant = common.Variant{
	Name: "Classical - Neutral Italy",
	Graph: func() godip.Graph {
		okNations := map[godip.Nation]bool{
			godip.Austria: true,
			godip.England:  true,
			godip.France: true,
			godip.Germany: true,
			godip.Turkey: true,
			godip.Russia: true,
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
			"edi":    godip.Unit{godip.Fleet, godip.England},
			"lvp":    godip.Unit{godip.Army, godip.England},
			"lon":    godip.Unit{godip.Fleet, godip.England},
			"bre":    godip.Unit{godip.Fleet, godip.France},
			"par":    godip.Unit{godip.Army, godip.France},
			"mar":    godip.Unit{godip.Army, godip.France},
			"kie":    godip.Unit{godip.Fleet, godip.Germany},
			"ber":    godip.Unit{godip.Army, godip.Germany},
			"mun":    godip.Unit{godip.Army, godip.Germany},
			"ven":    godip.Unit{godip.Army, godip.Neutral},
			"rom":    godip.Unit{godip.Army, godip.Neutral},
			"nap":    godip.Unit{godip.Fleet, godip.Neutral},
			"tri":    godip.Unit{godip.Fleet, godip.Austria},
			"vie":    godip.Unit{godip.Army, godip.Austria},
			"bud":    godip.Unit{godip.Army, godip.Austria},
			"stp/sc": godip.Unit{godip.Fleet, godip.Russia},
			"mos":    godip.Unit{godip.Army, godip.Russia},
			"war":    godip.Unit{godip.Army, godip.Russia},
			"sev":    godip.Unit{godip.Fleet, godip.Russia},
			"con":    godip.Unit{godip.Army, godip.Turkey},
			"smy":    godip.Unit{godip.Army, godip.Turkey},
			"ank":    godip.Unit{godip.Fleet, godip.Turkey},
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[godip.Province]godip.Nation{
			"edi": godip.England,
			"lvp": godip.England,
			"lon": godip.England,
			"bre": godip.France,
			"par": godip.France,
			"mar": godip.France,
			"kie": godip.Germany,
			"ber": godip.Germany,
			"mun": godip.Germany,
			"ven": godip.Neutral,
			"rom": godip.Neutral,
			"nap": godip.Neutral,
			"tri": godip.Austria,
			"vie": godip.Austria,
			"bud": godip.Austria,
			"con": godip.Turkey,
			"ank": godip.Turkey,
			"smy": godip.Turkey,
			"sev": godip.Russia,
			"mos": godip.Russia,
			"stp": godip.Russia,
			"war": godip.Russia,
		})
		return
	},
	Blank:             classical.Blank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           []godip.Nation{godip.Austria, godip.England, godip.France, godip.Germany, godip.Turkey, godip.Russia},
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
	Rules: `The first to 18 Supply Centers (SC) is the winner. 
	Kiel and Constantinople have a canal, so fleets can exit on either side. 
	Armies can move from Denmark to Kiel. Italy is a Neutral power. Italy's 
	SCs get an army which always holds and disbands when dislodged.`,
}

func NeutralOrders(state state.State) (ret map[godip.Province]godip.Adjudicator) {
	ret = map[godip.Province]godip.Adjudicator{}
	switch state.Phase().Type() {
	case godip.Movement:
		// Strictly this is unnecessary - because hold is the default order.
		for prov, unit := range state.Units() {
			if unit.Nation == godip.Neutral {
				ret[prov] = orders.Hold(prov)
			}
		}
	case godip.Adjustment:
		// Rebuild any missing units.
		for _, prov := range state.Graph().AllSCs() {
			if n, _, ok := state.SupplyCenter(prov); ok && n == godip.Neutral {
				if _, _, ok := state.Unit(prov); !ok {
					ret[prov] = orders.Build(prov, godip.Army, time.Now())
				}
			}
		}
	}
	return
}
