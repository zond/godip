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
	SoloSupplyCenters: 9,
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
	Rules: "",
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
		"nmd": dip.Unit{cla.Army, England},
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
		"nmd": England,
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
		// sco
		Prov("sco").Conn("num", cla.Land).Conn("nth", cla.Coast...).Conn("iri", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// ang
		Prov("ang").Conn("dov", cla.Coast...).Conn("was", cla.Coast...).Conn("nth", cla.Coast...).Conn("num", cla.Land).Conn("dev", cla.Land).Conn("lon", cla.Land).Flag(cla.Coast...).
		// dxx
		Prov("dxx").Conn("nmd", cla.Coast...).Conn("cal", cla.Coast...).Conn("lon", cla.Coast...).Conn("ech", cla.Sea).Flag(cla.Sea).
		// pro
		Prov("pro").Conn("dau", cla.Land).Conn("lim", cla.Land).Conn("tou", cla.Land).Conn("med", cla.Coast...).Conn("sav", cla.Land).Flag(cla.Coast...).
		// brt
		Prov("brt").Conn("ech", cla.Coast...).Conn("bch", cla.Coast...).Conn("bis", cla.Coast...).Conn("guy", cla.Land).Conn("poi", cla.Land).Conn("orl", cla.Land).Conn("anj", cla.Land).Conn("nmd", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// anj
		Prov("anj").Conn("nmd", cla.Land).Conn("brt", cla.Land).Conn("orl", cla.Land).Flag(cla.Land).
		// ara
		Prov("ara").Conn("med", cla.Coast...).Conn("tou", cla.Land).Conn("guy", cla.Land).Conn("bis", cla.Coast...).Conn("cas", cla.Land).Flag(cla.Coast...).
		// num
		Prov("num").Conn("nth", cla.Coast...).Conn("sco", cla.Land).Conn("iri", cla.Coast...).Conn("wal", cla.Land).Conn("dev", cla.Land).Conn("ang", cla.Land).Flag(cla.Coast...).
		// poi
		Prov("poi").Conn("tou", cla.Land).Conn("lim", cla.Land).Conn("orl", cla.Land).Conn("brt", cla.Land).Conn("guy", cla.Land).Flag(cla.Land).
		// tou
		Prov("tou").Conn("poi", cla.Land).Conn("guy", cla.Land).Conn("ara", cla.Land).Conn("med", cla.Coast...).Conn("pro", cla.Land).Conn("lim", cla.Land).Flag(cla.Coast...).SC(France).
		// sav
		Prov("sav").Conn("lim", cla.Land).Conn("tou", cla.Land).Conn("can", cla.Land).Conn("dau", cla.Land).Conn("pro", cla.Land).Conn("med", cla.Coast...).Flag(cla.Coast...).
		// cha
		Prov("cha").Conn("dij", cla.Land).Conn("par", cla.Land).Conn("dau", cla.Land).Flag(cla.Land).
		// par
		Prov("par").Conn("dij", cla.Land).Conn("cal", cla.Land).Conn("nmd", cla.Land).Conn("orl", cla.Land).Conn("dau", cla.Land).Conn("cha", cla.Land).Flag(cla.Land).SC(France).
		// wal
		Prov("wal").Conn("num", cla.Land).Conn("iri", cla.Coast...).Conn("bch", cla.Coast...).Conn("dev", cla.Land).Flag(cla.Coast...).
		// lux
		Prov("lux").Conn("fla", cla.Land).Conn("dij", cla.Land).Conn("lor", cla.Land).Conn("ger", cla.Land).Conn("fri", cla.Land).Conn("hol", cla.Land).Flag(cla.Land).SC(Burgundy).
		// lim
		Prov("lim").Conn("orl", cla.Land).Conn("poi", cla.Land).Conn("tou", cla.Land).Conn("pro", cla.Land).Conn("dau", cla.Land).Flag(cla.Land).
		// hol
		Prov("hol").Conn("dov", cla.Coast...).Conn("fla", cla.Land).Conn("lux", cla.Land).Conn("fri", cla.Land).Conn("was", cla.Coast...).Flag(cla.Coast...).SC(Burgundy).
		// lon
		Prov("lon").Conn("ech", cla.Coast...).Conn("dxx", cla.Coast...).Conn("dov", cla.Coast...).Conn("ang", cla.Land).Conn("dev", cla.Land).Flag(cla.Coast...).SC(England).
		// als
		Prov("als").Conn("lon", cla.Land).Conn("dev", cla.Land).Conn("ger", cla.Land).Conn("lor", cla.Land).Conn("can", cla.Land).Flag(cla.Land).
		// lor
		Prov("lor").Conn("dij", cla.Land).Conn("can", cla.Land).Conn("als", cla.Land).Conn("ger", cla.Land).Conn("lux", cla.Land).Flag(cla.Land).
		// iri
		Prov("iri").Conn("nth", cla.Sea).Conn("pal", cla.Coast...).Conn("atl", cla.Sea).Conn("bch", cla.Sea).Conn("wal", cla.Coast...).Conn("num", cla.Coast...).Conn("sco", cla.Coast...).Flag(cla.Sea).
		// pal
		Prov("pal").Conn("atl", cla.Coast...).Conn("iri", cla.Coast...).Conn("nth", cla.Coast...).Conn("ire", cla.Land).Flag(cla.Coast...).
		// med
		Prov("med").Conn("pal", cla.Coast...).Conn("ire", cla.Coast...).Conn("sav", cla.Coast...).Conn("pro", cla.Coast...).Conn("tou", cla.Coast...).Conn("ara", cla.Coast...).Conn("cas", cla.Coast...).Conn("atl", cla.Sea).Conn("atl", cla.Sea).Flag(cla.Sea).
		// ech
		Prov("ech").Conn("lon", cla.Coast...).Conn("dev", cla.Coast...).Conn("bch", cla.Sea).Conn("brt", cla.Coast...).Conn("nmd", cla.Coast...).Conn("dxx", cla.Sea).Flag(cla.Sea).
		// dau
		Prov("dau").Conn("par", cla.Land).Conn("orl", cla.Land).Conn("lim", cla.Land).Conn("pro", cla.Land).Conn("sav", cla.Land).Conn("can", cla.Land).Conn("dij", cla.Land).Conn("cha", cla.Land).Flag(cla.Land).SC(France).
		// nmd
		Prov("nmd").Conn("dxx", cla.Coast...).Conn("ech", cla.Coast...).Conn("brt", cla.Land).Conn("anj", cla.Land).Conn("orl", cla.Land).Conn("par", cla.Land).Conn("cal", cla.Land).Flag(cla.Coast...).SC(England).
		// bis
		Prov("bis").Conn("atl", cla.Sea).Conn("cas", cla.Coast...).Conn("ara", cla.Coast...).Conn("guy", cla.Coast...).Conn("brt", cla.Coast...).Conn("bch", cla.Sea).Flag(cla.Sea).
		// bch
		Prov("bch").Conn("iri", cla.Sea).Conn("atl", cla.Sea).Conn("bis", cla.Sea).Conn("brt", cla.Coast...).Conn("ech", cla.Sea).Conn("dev", cla.Coast...).Conn("wal", cla.Coast...).Flag(cla.Sea).
		// dij
		Prov("dij").Conn("par", cla.Land).Conn("cha", cla.Land).Conn("dau", cla.Land).Conn("can", cla.Land).Conn("lor", cla.Land).Conn("lux", cla.Land).Conn("fla", cla.Land).Conn("cal", cla.Land).Flag(cla.Land).SC(Burgundy).
		// nth
		Prov("nth").Conn("dij", cla.Coast...).Conn("cal", cla.Coast...).Conn("dij", cla.Coast...).Conn("cal", cla.Coast...).Conn("ire", cla.Coast...).Conn("pal", cla.Coast...).Conn("iri", cla.Sea).Conn("sco", cla.Coast...).Conn("num", cla.Coast...).Conn("ang", cla.Coast...).Conn("was", cla.Sea).Flag(cla.Sea).
		// cas
		Prov("cas").Conn("med", cla.Coast...).Conn("ara", cla.Land).Conn("bis", cla.Coast...).Conn("atl", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// was
		Prov("was").Conn("cas", cla.Coast...).Conn("atl", cla.Sea).Conn("nth", cla.Sea).Conn("ang", cla.Coast...).Conn("dov", cla.Sea).Conn("hol", cla.Coast...).Conn("fri", cla.Coast...).Conn("fri", cla.Coast...).Flag(cla.Sea).
		// atl
		Prov("atl").Conn("was", cla.Sea).Conn("fri", cla.Coast...).Conn("med", cla.Sea).Conn("cas", cla.Coast...).Conn("bis", cla.Sea).Conn("bch", cla.Sea).Conn("iri", cla.Sea).Conn("pal", cla.Coast...).Conn("ire", cla.Coast...).Conn("ire", cla.Coast...).Flag(cla.Sea).
		// orl
		Prov("orl").Conn("nmd", cla.Land).Conn("anj", cla.Land).Conn("brt", cla.Land).Conn("poi", cla.Land).Conn("lim", cla.Land).Conn("dau", cla.Land).Conn("par", cla.Land).Flag(cla.Land).SC(France).
		// dev
		Prov("dev").Conn("ech", cla.Coast...).Conn("lon", cla.Land).Conn("ang", cla.Land).Conn("num", cla.Land).Conn("wal", cla.Land).Conn("bch", cla.Coast...).Flag(cla.Coast...).SC(England).
		// can
		Prov("can").Conn("bch", cla.Coast...).Conn("dev", cla.Land).Conn("als", cla.Land).Conn("lor", cla.Land).Conn("dij", cla.Land).Conn("dau", cla.Land).Conn("sav", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// cal
		Prov("cal").Conn("dxx", cla.Coast...).Conn("nmd", cla.Land).Conn("par", cla.Land).Conn("dij", cla.Land).Conn("fla", cla.Land).Conn("dov", cla.Coast...).Flag(cla.Coast...).SC(England).
		// fla
		Prov("fla").Conn("lux", cla.Land).Conn("hol", cla.Land).Conn("dov", cla.Coast...).Conn("cal", cla.Land).Conn("dij", cla.Land).Flag(cla.Coast...).SC(Burgundy).
		// guy
		Prov("guy").Conn("brt", cla.Land).Conn("bis", cla.Coast...).Conn("ara", cla.Land).Conn("tou", cla.Land).Conn("poi", cla.Land).Flag(cla.Coast...).SC(England).
		// dov
		Prov("dov").Conn("ang", cla.Coast...).Conn("lon", cla.Coast...).Conn("cal", cla.Coast...).Conn("fla", cla.Coast...).Conn("hol", cla.Coast...).Conn("was", cla.Sea).Flag(cla.Sea).
		// fri
		Prov("fri").Conn("was", cla.Coast...).Conn("hol", cla.Land).Conn("lux", cla.Land).Conn("ger", cla.Land).Flag(cla.Coast...).
		Done()
}
