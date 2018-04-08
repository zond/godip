package pure

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"

	ord "github.com/zond/godip/orders"
)

var pureParser = ord.NewParser([]godip.Order{
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
	Phase:      classical.PhaseGenerator(pureParser),
	Parser:     pureParser,
	Nations:    classical.Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  []godip.UnitType{godip.Army},
	SoloWinner: common.SCCountWinner(4),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/puremap.svg")
	},
	SVGVersion: "2",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
	},
	CreatedBy:   "Danny Loeb",
	Version:     "vb10",
	Description: "A very minimal version of classical Diplomacy where each country is a single province.",
	Rules:       "Each of the seven nations has a single supply center, and each is adjacent to all of the others. The first player to own four of these centers is the winner.",
}

func PureBlank(phase godip.Phase) *state.State {
	return state.New(PureGraph(), phase, classical.BackupRule, nil)
}

func PureStart() (result *state.State, err error) {
	if result, err = classical.Start(); err != nil {
		return
	}
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
