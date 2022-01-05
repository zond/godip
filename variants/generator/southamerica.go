package southamerica

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Brazil   godip.Nation = "Brazil"
	Columbia godip.Nation = "Columbia"
	Peru     godip.Nation = "Peru"
)

var Nations = []godip.Nation{Brazil, Columbia, Peru}

var SouthAmericaVariant = common.Variant{
	Name:              "South America",
	Graph:             func() godip.Graph { return SouthAmericaGraph() },
	Start:             SouthAmericaStart,
	Blank:             SouthAmericaBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(9),
	SoloSCCount:       func(*state.State) int { return 9 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/southamericamap.svg")
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
	Description: "",
	Rules:       "",
}

func SouthAmericaBlank(phase godip.Phase) *state.State {
	return state.New(SouthAmericaGraph(), phase, classical.BackupRule, nil, nil)
}

func SouthAmericaStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1901, godip.Spring, godip.Movement)
	result = SouthAmericaBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"bel": godip.Unit{godip.Fleet, Brazil},
		"man": godip.Unit{godip.Army, Brazil},
		"rdj": godip.Unit{godip.Army, Brazil},
		"bar": godip.Unit{godip.Fleet, Columbia},
		"med": godip.Unit{godip.Army, Columbia},
		"bog": godip.Unit{godip.Army, Columbia},
		"lim": godip.Unit{godip.Fleet, Peru},
		"chi": godip.Unit{godip.Army, Peru},
		"are": godip.Unit{godip.Army, Peru},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"bel": Brazil,
		"rdj": Brazil,
		"bar": Columbia,
		"med": Columbia,
		"bog": Columbia,
		"lim": Peru,
		"chi": Peru,
		"are": Peru,
	})
	return
}

func SouthAmericaGraph() *graph.Graph {
	return graph.New().
		// Belem
		Prov("bel").Conn("rdj", godip.Coast...).Conn("mid", godip.Sea).Conn("bra", godip.Sea).Conn("guy", godip.Coast...).Conn("ror", godip.Land).Conn("acr", godip.Land).Conn("goi", godip.Land).Flag(godip.Coast...).SC(Brazil).
		// Cordillera Oriental
		Prov("cor").Conn("mon", godip.Land).Conn("ama", godip.Land).Conn("bog", godip.Land).Conn("val", godip.Land).Flag(godip.Land).
		// Mato Grosso
		Prov("mat").Conn("man", godip.Land).Conn("lap", godip.Land).Conn("yun", godip.Land).Conn("goi", godip.Land).Conn("acr", godip.Land).Flag(godip.Land).
		// Trujillo
		Prov("tru").Conn("ori", godip.Land).Conn("caa", godip.Coast...).Conn("cas", godip.Sea).Conn("bar", godip.Coast...).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Rio de Janeiro
		Prov("rdj").Conn("bel", godip.Coast...).Conn("goi", godip.Land).Conn("rgd", godip.Coast...).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(Brazil).
		// Montana
		Prov("mon").Conn("chi", godip.Land).Conn("are", godip.Land).Conn("ama", godip.Land).Conn("cor", godip.Land).Conn("val", godip.Land).Conn("ecu", godip.Land).Flag(godip.Land).
		// Costa Rica
		Prov("cos").Conn("sop", godip.Sea).Conn("gal", godip.Sea).Conn("pan", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Mid Atlantic Ocean
		Prov("mid").Conn("rdj", godip.Sea).Conn("rgd", godip.Sea).Conn("soa", godip.Sea).Conn("bra", godip.Sea).Conn("bel", godip.Sea).Flag(godip.Sea).
		// Chiclayo
		Prov("chi").Conn("are", godip.Land).Conn("mon", godip.Land).Conn("ecu", godip.Coast...).Conn("gal", godip.Sea).Conn("lim", godip.Coast...).Flag(godip.Coast...).SC(Peru).
		// Potosi
		Prov("pot").Conn("lap", godip.Coast...).Conn("soa", godip.Sea).Conn("soa", godip.Sea).Conn("soa", godip.Sea).Conn("soa", godip.Sea).Conn("par", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Bogota
		Prov("bog").Conn("bar", godip.Land).Conn("med", godip.Land).Conn("val", godip.Land).Conn("cor", godip.Land).Conn("ama", godip.Land).Conn("ori", godip.Land).Conn("tru", godip.Land).Flag(godip.Land).SC(Columbia).
		// Ecuador
		Prov("ecu").Conn("chi", godip.Coast...).Conn("mon", godip.Land).Conn("val", godip.Coast...).Conn("gol", godip.Sea).Conn("gal", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Amazon Basin
		Prov("ama").Conn("cor", godip.Land).Conn("mon", godip.Land).Conn("are", godip.Land).Conn("lap", godip.Land).Conn("man", godip.Land).Conn("ori", godip.Land).Conn("bog", godip.Land).Flag(godip.Land).
		// Goias
		Prov("goi").Conn("yun", godip.Land).Conn("par", godip.Land).Conn("rgd", godip.Land).Conn("rdj", godip.Land).Conn("bel", godip.Land).Conn("acr", godip.Land).Conn("mat", godip.Land).Flag(godip.Land).
		// Caracas
		Prov("caa").Conn("ror", godip.Land).Conn("guy", godip.Coast...).Conn("bra", godip.Sea).Conn("cas", godip.Sea).Conn("tru", godip.Coast...).Conn("ori", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Manaus
		Prov("man").Conn("ror", godip.Land).Conn("ori", godip.Land).Conn("ama", godip.Land).Conn("lap", godip.Land).Conn("mat", godip.Land).Conn("acr", godip.Land).Flag(godip.Land).
		// Medellin
		Prov("med").Conn("pan", godip.Coast...).Conn("gol", godip.Sea).Conn("val", godip.Coast...).Conn("bog", godip.Land).Conn("bar", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// La Paz
		Prov("lap").Conn("pot", godip.Coast...).Conn("par", godip.Coast...).Conn("yun", godip.Land).Conn("mat", godip.Land).Conn("man", godip.Land).Conn("ama", godip.Land).Conn("are", godip.Coast...).Conn("soa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Southeast Pacific
		Prov("sop").Conn("soa", godip.Sea).Conn("bah", godip.Sea).Conn("gal", godip.Sea).Conn("cos", godip.Sea).Flag(godip.Sea).
		// Galapagos Sea
		Prov("gal").Conn("sop", godip.Sea).Conn("bah", godip.Sea).Conn("lim", godip.Sea).Conn("chi", godip.Sea).Conn("ecu", godip.Sea).Conn("gol", godip.Sea).Conn("pan", godip.Sea).Conn("cos", godip.Sea).Flag(godip.Sea).
		// Caribbean Sea
		Prov("cas").Conn("cos", godip.Sea).Conn("pan", godip.Sea).Conn("bar", godip.Sea).Conn("tru", godip.Sea).Conn("caa", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Lima
		Prov("lim").Conn("gal", godip.Sea).Conn("bah", godip.Sea).Conn("are", godip.Coast...).Conn("chi", godip.Coast...).Flag(godip.Coast...).SC(Peru).
		// Guyana
		Prov("guy").Conn("bel", godip.Coast...).Conn("bra", godip.Sea).Conn("caa", godip.Coast...).Conn("ror", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Yungaz
		Prov("yun").Conn("goi", godip.Land).Conn("mat", godip.Land).Conn("lap", godip.Land).Conn("par", godip.Land).Flag(godip.Land).
		// Roraima
		Prov("ror").Conn("caa", godip.Land).Conn("ori", godip.Land).Conn("man", godip.Land).Conn("acr", godip.Land).Conn("bel", godip.Land).Conn("guy", godip.Land).Flag(godip.Land).
		// Panama
		Prov("pan").Conn("med", godip.Coast...).Conn("bar", godip.Coast...).Conn("cas", godip.Sea).Conn("cos", godip.Coast...).Conn("gal", godip.Sea).Conn("gol", godip.Sea).Flag(godip.Coast...).
		// Valle Magdalena
		Prov("val").Conn("med", godip.Coast...).Conn("gol", godip.Sea).Conn("ecu", godip.Coast...).Conn("mon", godip.Land).Conn("cor", godip.Land).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Southwest Atlantic
		Prov("soa").Conn("bra", godip.Sea).Conn("mid", godip.Sea).Conn("rgd", godip.Sea).Conn("rgd", godip.Sea).Conn("par", godip.Sea).Conn("pot", godip.Sea).Conn("pot", godip.Sea).Conn("pot", godip.Sea).Conn("pot", godip.Sea).Conn("lap", godip.Sea).Conn("are", godip.Sea).Conn("bah", godip.Sea).Conn("sop", godip.Sea).Flag(godip.Sea).
		// Acre
		Prov("acr").Conn("ror", godip.Land).Conn("man", godip.Land).Conn("mat", godip.Land).Conn("goi", godip.Land).Conn("bel", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bahia da Arica
		Prov("bah").Conn("sop", godip.Sea).Conn("soa", godip.Sea).Conn("are", godip.Sea).Conn("lim", godip.Sea).Conn("gal", godip.Sea).Flag(godip.Sea).
		// Golfo da Panama
		Prov("gol").Conn("med", godip.Sea).Conn("pan", godip.Sea).Conn("gal", godip.Sea).Conn("ecu", godip.Sea).Conn("val", godip.Sea).Flag(godip.Sea).
		// Rio Grande do Sul
		Prov("rgd").Conn("soa", godip.Sea).Conn("soa", godip.Sea).Conn("mid", godip.Sea).Conn("rdj", godip.Coast...).Conn("goi", godip.Land).Conn("par", godip.Coast...).Flag(godip.Coast...).
		// Barranquilla
		Prov("bar").Conn("bog", godip.Land).Conn("tru", godip.Coast...).Conn("cas", godip.Sea).Conn("pan", godip.Coast...).Conn("med", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// Brazilian Sea
		Prov("bra").Conn("cas", godip.Sea).Conn("caa", godip.Sea).Conn("guy", godip.Sea).Conn("bel", godip.Sea).Conn("mid", godip.Sea).Conn("soa", godip.Sea).Flag(godip.Sea).
		// Arequipa
		Prov("are").Conn("chi", godip.Land).Conn("lim", godip.Coast...).Conn("bah", godip.Sea).Conn("soa", godip.Sea).Conn("lap", godip.Coast...).Conn("ama", godip.Land).Conn("mon", godip.Land).Flag(godip.Coast...).SC(Peru).
		// Orinoco Springs
		Prov("ori").Conn("tru", godip.Land).Conn("bog", godip.Land).Conn("ama", godip.Land).Conn("man", godip.Land).Conn("ror", godip.Land).Conn("caa", godip.Land).Flag(godip.Land).
		// Paraguay
		Prov("par").Conn("lap", godip.Coast...).Conn("pot", godip.Coast...).Conn("soa", godip.Sea).Conn("rgd", godip.Coast...).Conn("goi", godip.Land).Conn("yun", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"bel": "Belem",
	"cor": "Cordillera Oriental",
	"mat": "Mato Grosso",
	"tru": "Trujillo",
	"rdj": "Rio de Janeiro",
	"mon": "Montana",
	"cos": "Costa Rica",
	"mid": "Mid Atlantic Ocean",
	"chi": "Chiclayo",
	"pot": "Potosi",
	"bog": "Bogota",
	"ecu": "Ecuador",
	"ama": "Amazon Basin",
	"goi": "Goias",
	"caa": "Caracas",
	"man": "Manaus",
	"med": "Medellin",
	"lap": "La Paz",
	"sop": "Southeast Pacific",
	"gal": "Galapagos Sea",
	"cas": "Caribbean Sea",
	"lim": "Lima",
	"guy": "Guyana",
	"yun": "Yungaz",
	"ror": "Roraima",
	"pan": "Panama",
	"val": "Valle Magdalena",
	"soa": "Southwest Atlantic",
	"acr": "Acre",
	"bah": "Bahia da Arica",
	"gol": "Golfo da Panama",
	"rgd": "Rio Grande do Sul",
	"bar": "Barranquilla",
	"bra": "Brazilian Sea",
	"are": "Arequipa",
	"ori": "Orinoco Springs",
	"par": "Paraguay",
}
