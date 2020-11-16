package pure

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

var pureParser = orders.NewParser([]godip.Order{
	orders.BuildOrder,
	orders.DisbandOrder,
	orders.HoldOrder,
	orders.MoveOrder,
	orders.SupportOrder,
})

var PureVariant = common.Variant{
	Name:       "Pure",
	Graph:      func() godip.Graph { return PureGraph() },
	Start:      PureStart,
	Blank:      PureBlank,
	Phase:      phase.Generator(pureParser, classical.AdjustSCs),
	Parser:     pureParser,
	Nations:    classical.Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  []godip.UnitType{godip.Army},
	SoloWinner: common.SCCountWinner(4),
	ProvinceLongNames: map[godip.Province]string{
		"lon": "London",
		"ber": "Berlin",
		"par": "Paris",
		"rom": "Rome",
		"con": "Constantinople",
		"vie": "Vienna",
		"mos": "Moscow",
	},
	SVGMap: func() ([]byte, error) {
		return Asset("svg/puremap.svg")
	},
	SVGVersion: "5",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
	},
	SVGFlags:    classical.SVGFlags,
	CreatedBy:   "Danny Loeb",
	Version:     "vb10",
	Description: "A minimal version of Diplomacy where each country is a single province.",
	SoloSCCount: func(*state.State) int { return 4 },
	Rules: `First to 4 Supply Centers (SC) is the winner.
	Each nation has only one SC, and each is adjacent to all others.`,
}

func PureBlank(phase godip.Phase) *state.State {
	return state.New(PureGraph(), phase, classical.BackupRule, nil, nil)
}

func PureStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1901, godip.Spring, godip.Movement)
	result = PureBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"ber": godip.Unit{godip.Army, godip.Germany},
		"lon": godip.Unit{godip.Army, godip.England},
		"par": godip.Unit{godip.Army, godip.France},
		"rom": godip.Unit{godip.Army, godip.Italy},
		"con": godip.Unit{godip.Army, godip.Turkey},
		"vie": godip.Unit{godip.Army, godip.Austria},
		"mos": godip.Unit{godip.Army, godip.Russia},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"ber": godip.Germany,
		"lon": godip.England,
		"par": godip.France,
		"rom": godip.Italy,
		"con": godip.Turkey,
		"vie": godip.Austria,
		"mos": godip.Russia,
	})
	return
}

func PureGraph() *graph.Graph {
	return graph.New().
		// ber
		Prov("ber").Conn("lon", godip.Land).Conn("par", godip.Land).Conn("rom", godip.Land).Conn("con", godip.Land).Conn("vie", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.Germany).
		// lon
		Prov("lon").Conn("ber", godip.Land).Conn("par", godip.Land).Conn("rom", godip.Land).Conn("con", godip.Land).Conn("vie", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.England).
		// par
		Prov("par").Conn("ber", godip.Land).Conn("lon", godip.Land).Conn("rom", godip.Land).Conn("con", godip.Land).Conn("vie", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.France).
		// rom
		Prov("rom").Conn("ber", godip.Land).Conn("lon", godip.Land).Conn("par", godip.Land).Conn("con", godip.Land).Conn("vie", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.Italy).
		// con
		Prov("con").Conn("ber", godip.Land).Conn("lon", godip.Land).Conn("par", godip.Land).Conn("rom", godip.Land).Conn("vie", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.Turkey).
		// vie
		Prov("vie").Conn("ber", godip.Land).Conn("lon", godip.Land).Conn("par", godip.Land).Conn("rom", godip.Land).Conn("con", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(godip.Austria).
		// mos
		Prov("mos").Conn("ber", godip.Land).Conn("lon", godip.Land).Conn("par", godip.Land).Conn("rom", godip.Land).Conn("con", godip.Land).Conn("vie", godip.Land).Flag(godip.Land).SC(godip.Russia).
		Done()
}
