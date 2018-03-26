package pure

import (
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	dip "github.com/zond/godip/common"
	ord "github.com/zond/godip/orders"
	cla "github.com/zond/godip/variants/classical/common"
)

var pureParser = ord.NewParser([]dip.Order{
	orders.BuildOrder,
	orders.DisbandOrder,
	orders.HoldOrder,
	orders.MoveOrder,
	orders.SupportOrder,
})

var PureVariant = common.Variant{
	Name:       "Pure",
	Graph:      func() dip.Graph { return PureGraph() },
	Start:      PureStart,
	Blank:      PureBlank,
	Phase:      classical.Phase,
	Parser:     pureParser,
	Nations:    cla.Nations,
	PhaseTypes: cla.PhaseTypes,
	Seasons:    cla.Seasons,
	UnitTypes:  []dip.UnitType{cla.Army},
	SoloWinner: common.SCCountWinner(4),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/puremap.svg")
	},
	SVGVersion: "2",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
	},
	CreatedBy:   "Danny Loeb",
	Version:     "vb10",
	Description: "A very minimal version of classical Diplomacy where each country is a single province.",
	Rules:       "Each of the seven nations has a single supply center, and each is adjacent to all of the others. The first player to own four of these centers is the winner.",
}

func PureBlank(phase dip.Phase) *state.State {
	return state.New(PureGraph(), phase, classical.BackupRule)
}

func PureStart() (result *state.State, err error) {
	if result, err = classical.Start(); err != nil {
		return
	}
	if err = result.SetUnits(map[dip.Province]dip.Unit{
		"ber": dip.Unit{cla.Army, cla.Germany},
		"lon": dip.Unit{cla.Army, cla.England},
		"par": dip.Unit{cla.Army, cla.France},
		"rom": dip.Unit{cla.Army, cla.Italy},
		"con": dip.Unit{cla.Army, cla.Turkey},
		"vie": dip.Unit{cla.Army, cla.Austria},
		"mos": dip.Unit{cla.Army, cla.Russia},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
		"ber": cla.Germany,
		"lon": cla.England,
		"par": cla.France,
		"rom": cla.Italy,
		"con": cla.Turkey,
		"vie": cla.Austria,
		"mos": cla.Russia,
	})
	return
}

func PureGraph() *graph.Graph {
	return graph.New().
		// ber
		Prov("ber").Conn("lon", cla.Land).Conn("par", cla.Land).Conn("rom", cla.Land).Conn("con", cla.Land).Conn("vie", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.Germany).
		// lon
		Prov("lon").Conn("ber", cla.Land).Conn("par", cla.Land).Conn("rom", cla.Land).Conn("con", cla.Land).Conn("vie", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.England).
		// par
		Prov("par").Conn("ber", cla.Land).Conn("lon", cla.Land).Conn("rom", cla.Land).Conn("con", cla.Land).Conn("vie", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.France).
		// rom
		Prov("rom").Conn("ber", cla.Land).Conn("lon", cla.Land).Conn("par", cla.Land).Conn("con", cla.Land).Conn("vie", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.Italy).
		// con
		Prov("con").Conn("ber", cla.Land).Conn("lon", cla.Land).Conn("par", cla.Land).Conn("rom", cla.Land).Conn("vie", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.Turkey).
		// vie
		Prov("vie").Conn("ber", cla.Land).Conn("lon", cla.Land).Conn("par", cla.Land).Conn("rom", cla.Land).Conn("con", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).SC(cla.Austria).
		// mos
		Prov("mos").Conn("ber", cla.Land).Conn("lon", cla.Land).Conn("par", cla.Land).Conn("rom", cla.Land).Conn("con", cla.Land).Conn("vie", cla.Land).Flag(cla.Land).SC(cla.Russia).
		Done()
}
