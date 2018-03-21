package hundred

import (
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

const (
	England  dip.Nation = "England"
	Burgundy dip.Nation = "Burgundy"
	France   dip.Nation = "France"
)

var Nations = []dip.Nation{England, Burgundy, France}

var HundredVariant = common.Variant{
	Name:        "Hundred",
	Graph:       func() dip.Graph { return HundredGraph() },
	Start:       HundredStart,
	Blank:       HundredBlank,
	Phase:       classical.Phase,
	ParseOrders: orders.ParseAll,
	ParseOrder:  orders.Parse,
	OrderTypes:  orders.OrderTypes(),
	Nations:     Nations,
	PhaseTypes:  cla.PhaseTypes,
	Seasons:     cla.Seasons,
	UnitTypes:   cla.UnitTypes,
	SoloWinner:  common.SCCountWinner(9),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/hundredmap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules:       "",
}

func HundredBlank(phase dip.Phase) *state.State {
	return state.New(HundredGraph(), phase, classical.BackupRule)
}

func HundredStart() (result *state.State, err error) {
	startPhase := classical.Phase(1425, cla.Spring, cla.Movement)
	result = state.New(HundredGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[dip.Province]dip.Unit{
		"lon": dip.Unit{cla.Fleet, England},
		"dev": dip.Unit{cla.Fleet, England},
		"cal": dip.Unit{cla.Army, England},
		"guy": dip.Unit{cla.Army, England},
		"nom": dip.Unit{cla.Army, England},
		"hol": dip.Unit{cla.Fleet, Burgundy},
		"dij": dip.Unit{cla.Army, Burgundy},
		"lux": dip.Unit{cla.Army, Burgundy},
		"fla": dip.Unit{cla.Army, Burgundy},
		"dau": dip.Unit{cla.Army, France},
		"orl": dip.Unit{cla.Army, France},
		"par": dip.Unit{cla.Army, France},
		"tou": dip.Unit{cla.Army, France},
		"pro": dip.Unit{cla.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
		"lon": England,
		"dev": England,
		"cal": England,
		"guy": England,
		"nom": England,
		"hol": Burgundy,
		"dij": Burgundy,
		"lux": Burgundy,
		"fla": Burgundy,
		"dau": France,
		"orl": France,
		"par": France,
		"tou": France,
	})
	return
}

func HundredGraph() *graph.Graph {
	return graph.New().
		// Atlantic Sea
		Prov("atl").Conn("med", cla.Sea).Conn("cas", cla.Sea).Conn("bis", cla.Sea).Conn("brs", cla.Sea).Conn("iri", cla.Sea).Conn("thp", cla.Sea).Flag(cla.Sea).
		// Normandy
		Prov("nom").Conn("str", cla.Sea).Conn("eng", cla.Sea).Conn("brt", cla.Coast...).Conn("anj", cla.Land).Conn("orl", cla.Land).Conn("par", cla.Land).Conn("cal", cla.Coast...).Flag(cla.Coast...).SC(England).
		// Dauphine
		Prov("dau").Conn("orl", cla.Land).Conn("lim", cla.Land).Conn("pro", cla.Land).Conn("sav", cla.Land).Conn("can", cla.Land).Conn("dij", cla.Land).Conn("cha", cla.Land).Conn("par", cla.Land).Flag(cla.Land).SC(France).
		// Anjou
		Prov("anj").Conn("orl", cla.Land).Conn("nom", cla.Land).Conn("brt", cla.Land).Flag(cla.Land).
		// Guyenne
		Prov("guy").Conn("ara", cla.Coast...).Conn("tou", cla.Land).Conn("poi", cla.Land).Conn("brt", cla.Coast...).Conn("bis", cla.Sea).Flag(cla.Coast...).SC(England).
		// Friesland
		Prov("fri").Conn("lux", cla.Land).Conn("thw", cla.Sea).Conn("hol", cla.Coast...).Flag(cla.Coast...).
		// North Sea
		Prov("nos").Conn("thp", cla.Sea).Conn("iri", cla.Sea).Conn("sco", cla.Sea).Conn("not", cla.Sea).Conn("ang", cla.Sea).Conn("thw", cla.Sea).Flag(cla.Sea).
		// The Wash
		Prov("thw").Conn("nos", cla.Sea).Conn("ang", cla.Sea).Conn("str", cla.Sea).Conn("hol", cla.Sea).Conn("fri", cla.Sea).Flag(cla.Sea).
		// Devon
		Prov("dev").Conn("wal", cla.Coast...).Conn("brs", cla.Sea).Conn("eng", cla.Sea).Conn("lon", cla.Coast...).Conn("ang", cla.Land).Conn("not", cla.Land).Flag(cla.Coast...).SC(England).
		// London
		Prov("lon").Conn("cal", cla.Coast...).Conn("dev", cla.Coast...).Conn("eng", cla.Sea).Conn("str", cla.Sea).Conn("ang", cla.Coast...).Flag(cla.Coast...).SC(England).
		// Calais
		Prov("cal").Conn("lon", cla.Coast...).Conn("str", cla.Sea).Conn("nom", cla.Coast...).Conn("par", cla.Land).Conn("dij", cla.Land).Conn("fla", cla.Coast...).Flag(cla.Coast...).SC(England).
		// Alsace
		Prov("als").Conn("lor", cla.Land).Conn("can", cla.Land).Flag(cla.Land).
		// Poitou
		Prov("poi").Conn("lim", cla.Land).Conn("orl", cla.Land).Conn("brt", cla.Land).Conn("guy", cla.Land).Conn("tou", cla.Land).Flag(cla.Land).
		// Biscay
		Prov("bis").Conn("cas", cla.Sea).Conn("ara", cla.Sea).Conn("guy", cla.Sea).Conn("brt", cla.Sea).Conn("brs", cla.Sea).Conn("atl", cla.Sea).Flag(cla.Sea).
		// Savoy
		Prov("sav").Conn("can", cla.Land).Conn("dau", cla.Land).Conn("pro", cla.Coast...).Conn("med", cla.Sea).Flag(cla.Coast...).
		// Orleanais
		Prov("orl").Conn("dau", cla.Land).Conn("par", cla.Land).Conn("nom", cla.Land).Conn("anj", cla.Land).Conn("brt", cla.Land).Conn("poi", cla.Land).Conn("lim", cla.Land).Flag(cla.Land).SC(France).
		// Strait of Dover
		Prov("str").Conn("cal", cla.Sea).Conn("fla", cla.Sea).Conn("hol", cla.Sea).Conn("thw", cla.Sea).Conn("ang", cla.Sea).Conn("lon", cla.Sea).Conn("eng", cla.Sea).Conn("nom", cla.Sea).Flag(cla.Sea).
		// Mediterranean
		Prov("med").Conn("sav", cla.Sea).Conn("pro", cla.Sea).Conn("tou", cla.Sea).Conn("ara", cla.Sea).Conn("cas", cla.Sea).Conn("atl", cla.Sea).Flag(cla.Sea).
		// Lorraine
		Prov("lor").Conn("lux", cla.Land).Conn("dij", cla.Land).Conn("can", cla.Land).Conn("als", cla.Land).Flag(cla.Land).
		// Flanders
		Prov("fla").Conn("str", cla.Sea).Conn("cal", cla.Coast...).Conn("dij", cla.Land).Conn("lux", cla.Land).Conn("hol", cla.Coast...).Flag(cla.Coast...).SC(Burgundy).
		// Bristol Channel
		Prov("brs").Conn("wal", cla.Sea).Conn("iri", cla.Sea).Conn("atl", cla.Sea).Conn("bis", cla.Sea).Conn("brt", cla.Sea).Conn("eng", cla.Sea).Conn("dev", cla.Sea).Flag(cla.Sea).
		// Cantons
		Prov("can").Conn("als", cla.Land).Conn("lor", cla.Land).Conn("dij", cla.Land).Conn("dau", cla.Land).Conn("sav", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Northumbria
		Prov("not").Conn("wal", cla.Coast...).Conn("dev", cla.Land).Conn("ang", cla.Coast...).Conn("nos", cla.Sea).Conn("sco", cla.Coast...).Conn("iri", cla.Sea).Flag(cla.Coast...).
		// Provence
		Prov("pro").Conn("lim", cla.Land).Conn("tou", cla.Coast...).Conn("med", cla.Sea).Conn("sav", cla.Coast...).Conn("dau", cla.Land).Flag(cla.Coast...).
		// Paris
		Prov("par").Conn("orl", cla.Land).Conn("dau", cla.Land).Conn("cha", cla.Land).Conn("dij", cla.Land).Conn("cal", cla.Land).Conn("nom", cla.Land).Flag(cla.Land).SC(France).
		// Toulouse
		Prov("tou").Conn("med", cla.Sea).Conn("pro", cla.Coast...).Conn("lim", cla.Land).Conn("poi", cla.Land).Conn("guy", cla.Land).Conn("ara", cla.Coast...).Flag(cla.Coast...).SC(France).
		// Irish Sea
		Prov("iri").Conn("nos", cla.Sea).Conn("thp", cla.Sea).Conn("atl", cla.Sea).Conn("brs", cla.Sea).Conn("wal", cla.Sea).Conn("not", cla.Sea).Conn("sco", cla.Sea).Flag(cla.Sea).
		// Dijon
		Prov("dij").Conn("lux", cla.Land).Conn("fla", cla.Land).Conn("cal", cla.Land).Conn("par", cla.Land).Conn("cha", cla.Land).Conn("dau", cla.Land).Conn("can", cla.Land).Conn("lor", cla.Land).Flag(cla.Land).SC(Burgundy).
		// Scotland
		Prov("sco").Conn("nos", cla.Sea).Conn("iri", cla.Sea).Conn("not", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Brittany
		Prov("brt").Conn("bis", cla.Sea).Conn("guy", cla.Coast...).Conn("poi", cla.Land).Conn("orl", cla.Land).Conn("anj", cla.Land).Conn("nom", cla.Coast...).Conn("eng", cla.Sea).Conn("brs", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Limousin
		Prov("lim").Conn("poi", cla.Land).Conn("tou", cla.Land).Conn("pro", cla.Land).Conn("dau", cla.Land).Conn("orl", cla.Land).Flag(cla.Land).
		// Luxembourg
		Prov("lux").Conn("fri", cla.Land).Conn("hol", cla.Land).Conn("fla", cla.Land).Conn("dij", cla.Land).Conn("lor", cla.Land).Flag(cla.Land).SC(Burgundy).
		// Wales
		Prov("wal").Conn("brs", cla.Sea).Conn("dev", cla.Coast...).Conn("not", cla.Coast...).Conn("iri", cla.Sea).Flag(cla.Coast...).
		// English Channel
		Prov("eng").Conn("str", cla.Sea).Conn("lon", cla.Sea).Conn("dev", cla.Sea).Conn("brs", cla.Sea).Conn("brt", cla.Sea).Conn("nom", cla.Sea).Flag(cla.Sea).
		// Anglia
		Prov("ang").Conn("thw", cla.Sea).Conn("nos", cla.Sea).Conn("not", cla.Coast...).Conn("dev", cla.Land).Conn("lon", cla.Coast...).Conn("str", cla.Sea).Flag(cla.Coast...).
		// Aragon
		Prov("ara").Conn("med", cla.Sea).Conn("tou", cla.Coast...).Conn("guy", cla.Coast...).Conn("bis", cla.Sea).Conn("cas", cla.Coast...).Flag(cla.Coast...).
		// Castile
		Prov("cas").Conn("bis", cla.Sea).Conn("atl", cla.Sea).Conn("med", cla.Sea).Conn("ara", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Charolais
		Prov("cha").Conn("dau", cla.Land).Conn("dij", cla.Land).Conn("par", cla.Land).Flag(cla.Land).
		// The Pale
		Prov("thp").Conn("atl", cla.Sea).Conn("iri", cla.Sea).Conn("nos", cla.Sea).Flag(cla.Coast...).
		// Holland
		Prov("hol").Conn("str", cla.Sea).Conn("fla", cla.Coast...).Conn("lux", cla.Land).Conn("fri", cla.Coast...).Conn("thw", cla.Sea).Flag(cla.Coast...).SC(Burgundy).
		Done()
}
