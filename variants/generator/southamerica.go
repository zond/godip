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
	Chile    godip.Nation = "Chile"
	Peru     godip.Nation = "Peru"
)

var Nations = []godip.Nation{Brazil, Columbia, Chile, Peru}

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
	SoloWinner:        common.SCCountWinner(13),
	SoloSCCount:       func(*state.State) int { return 13 },
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
		"con": godip.Unit{godip.Fleet, Chile},
		"ant": godip.Unit{godip.Army, Chile},
		"sat": godip.Unit{godip.Army, Chile},
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
		"ant": Chile,
		"sat": Chile,
		"lim": Peru,
		"chi": Peru,
		"are": Peru,
	})
	return
}

func SouthAmericaGraph() *graph.Graph {
	return graph.New().
		// Pampas
		Prov("pam").Conn("mar", godip.Land).Conn("bue", godip.Land).Conn("gra", godip.Land).Conn("sat", godip.Land).Conn("con", godip.Land).Flag(godip.Land).
		// Belem
		Prov("bel").Conn("bra", godip.Sea).Conn("guy", godip.Coast...).Conn("ror", godip.Land).Conn("acr", godip.Land).Conn("goi", godip.Land).Conn("rdj", godip.Coast...).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(Brazil).
		// Cordillera Oriental
		Prov("cor").Conn("mon", godip.Land).Conn("ama", godip.Land).Conn("bog", godip.Land).Conn("val", godip.Land).Flag(godip.Land).
		// Santiago
		Prov("sat").Conn("bdc", godip.Sea).Conn("pat", godip.Coast...).Conn("con", godip.Land).Conn("pam", godip.Land).Conn("gra", godip.Land).Conn("ant", godip.Coast...).Flag(godip.Coast...).SC(Chile).
		// Roraima
		Prov("ror").Conn("caa", godip.Land).Conn("ori", godip.Land).Conn("man", godip.Land).Conn("acr", godip.Land).Conn("bel", godip.Land).Conn("guy", godip.Land).Flag(godip.Land).
		// Southwest Atlantic
		Prov("soa").Conn("bra", godip.Sea).Conn("mid", godip.Sea).Conn("coa", godip.Sea).Conn("tie", godip.Sea).Conn("sco", godip.Sea).Flag(godip.Sea).
		// Uruguay
		Prov("uru").Conn("bue", godip.Coast...).Conn("coa", godip.Sea).Conn("mid", godip.Sea).Conn("rgd", godip.Coast...).Conn("saf", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Trujillo
		Prov("tru").Conn("ori", godip.Land).Conn("caa", godip.Coast...).Conn("cas", godip.Sea).Conn("bar", godip.Coast...).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Rio de Janeiro
		Prov("rdj").Conn("mid", godip.Sea).Conn("bel", godip.Coast...).Conn("goi", godip.Land).Conn("rgd", godip.Coast...).Flag(godip.Coast...).SC(Brazil).
		// Bahia da Coquimbo
		Prov("bdc").Conn("sat", godip.Sea).Conn("ant", godip.Sea).Conn("bda", godip.Sea).Conn("isl", godip.Sea).Conn("sco", godip.Sea).Conn("pat", godip.Sea).Flag(godip.Sea).
		// Montana
		Prov("mon").Conn("chi", godip.Land).Conn("are", godip.Land).Conn("ama", godip.Land).Conn("cor", godip.Land).Conn("val", godip.Land).Conn("ecu", godip.Land).Flag(godip.Land).
		// Costa Rica
		Prov("cos").Conn("sop", godip.Sea).Conn("gal", godip.Sea).Conn("pan", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Mid Atlantic Ocean
		Prov("mid").Conn("uru", godip.Sea).Conn("coa", godip.Sea).Conn("soa", godip.Sea).Conn("bra", godip.Sea).Conn("bel", godip.Sea).Conn("rdj", godip.Sea).Conn("rgd", godip.Sea).Flag(godip.Sea).
		// Chiclayo
		Prov("chi").Conn("gal", godip.Sea).Conn("lim", godip.Coast...).Conn("are", godip.Land).Conn("mon", godip.Land).Conn("ecu", godip.Coast...).Flag(godip.Coast...).SC(Peru).
		// Patagonia
		Prov("pat").Conn("sat", godip.Coast...).Conn("bdc", godip.Sea).Conn("sco", godip.Sea).Conn("tie", godip.Coast...).Conn("con", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Potosi
		Prov("pot").Conn("lap", godip.Land).Conn("des", godip.Land).Conn("ant", godip.Land).Conn("gra", godip.Land).Conn("mes", godip.Land).Conn("par", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bogota
		Prov("bog").Conn("bar", godip.Land).Conn("med", godip.Land).Conn("val", godip.Land).Conn("cor", godip.Land).Conn("ama", godip.Land).Conn("ori", godip.Land).Conn("tru", godip.Land).Flag(godip.Land).SC(Columbia).
		// Ecuador
		Prov("ecu").Conn("chi", godip.Coast...).Conn("mon", godip.Land).Conn("val", godip.Coast...).Conn("gol", godip.Sea).Conn("gal", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Gran Chaco
		Prov("gra").Conn("mes", godip.Land).Conn("pot", godip.Land).Conn("ant", godip.Land).Conn("sat", godip.Land).Conn("pam", godip.Land).Conn("bue", godip.Land).Conn("saf", godip.Land).Flag(godip.Land).
		// Amazon Basin
		Prov("ama").Conn("cor", godip.Land).Conn("mon", godip.Land).Conn("are", godip.Land).Conn("lap", godip.Land).Conn("man", godip.Land).Conn("ori", godip.Land).Conn("bog", godip.Land).Flag(godip.Land).
		// Goias
		Prov("goi").Conn("yun", godip.Land).Conn("par", godip.Land).Conn("rgd", godip.Land).Conn("rdj", godip.Land).Conn("bel", godip.Land).Conn("acr", godip.Land).Conn("mat", godip.Land).Flag(godip.Land).
		// Caracas
		Prov("caa").Conn("ror", godip.Land).Conn("guy", godip.Coast...).Conn("bra", godip.Sea).Conn("cas", godip.Sea).Conn("tru", godip.Coast...).Conn("ori", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Islas Juan Fernandez
		Prov("isl").Conn("sco", godip.Sea).Conn("bdc", godip.Sea).Conn("bda", godip.Sea).Conn("sop", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Medellin
		Prov("med").Conn("pan", godip.Coast...).Conn("gol", godip.Sea).Conn("val", godip.Coast...).Conn("bog", godip.Land).Conn("bar", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// La Paz
		Prov("lap").Conn("pot", godip.Land).Conn("par", godip.Land).Conn("yun", godip.Land).Conn("mat", godip.Land).Conn("man", godip.Land).Conn("ama", godip.Land).Conn("are", godip.Land).Conn("des", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Manaus
		Prov("man").Conn("mat", godip.Land).Conn("acr", godip.Land).Conn("ror", godip.Land).Conn("ori", godip.Land).Conn("ama", godip.Land).Conn("lap", godip.Land).Flag(godip.Land).
		// Southeast Pacific
		Prov("sop").Conn("sco", godip.Sea).Conn("isl", godip.Sea).Conn("bda", godip.Sea).Conn("gal", godip.Sea).Conn("cos", godip.Sea).Flag(godip.Sea).
		// Galapagos Sea
		Prov("gal").Conn("chi", godip.Sea).Conn("ecu", godip.Sea).Conn("gol", godip.Sea).Conn("pan", godip.Sea).Conn("cos", godip.Sea).Conn("sop", godip.Sea).Conn("bda", godip.Sea).Conn("lim", godip.Sea).Flag(godip.Sea).
		// Santa Fe
		Prov("saf").Conn("bue", godip.Land).Conn("uru", godip.Land).Conn("rgd", godip.Land).Conn("mes", godip.Land).Conn("gra", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Caribbean Sea
		Prov("cas").Conn("cos", godip.Sea).Conn("pan", godip.Sea).Conn("bar", godip.Sea).Conn("tru", godip.Sea).Conn("caa", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Lima
		Prov("lim").Conn("gal", godip.Sea).Conn("bda", godip.Sea).Conn("are", godip.Coast...).Conn("chi", godip.Coast...).Flag(godip.Coast...).SC(Peru).
		// Yungaz
		Prov("yun").Conn("goi", godip.Land).Conn("mat", godip.Land).Conn("lap", godip.Land).Conn("par", godip.Land).Flag(godip.Land).
		// Guyana
		Prov("guy").Conn("bel", godip.Coast...).Conn("bra", godip.Sea).Conn("caa", godip.Coast...).Conn("ror", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Arequipa
		Prov("are").Conn("chi", godip.Land).Conn("lim", godip.Coast...).Conn("bda", godip.Sea).Conn("des", godip.Coast...).Conn("lap", godip.Land).Conn("ama", godip.Land).Conn("mon", godip.Land).Flag(godip.Coast...).SC(Peru).
		// Tierra del Fuego
		Prov("tie").Conn("coa", godip.Sea).Conn("con", godip.Coast...).Conn("pat", godip.Coast...).Conn("sco", godip.Sea).Conn("soa", godip.Sea).Flag(godip.Coast...).
		// Desierto Atacama
		Prov("des").Conn("pot", godip.Land).Conn("lap", godip.Land).Conn("are", godip.Coast...).Conn("bda", godip.Sea).Conn("ant", godip.Coast...).Flag(godip.Coast...).
		// Panama
		Prov("pan").Conn("med", godip.Coast...).Conn("bar", godip.Coast...).Conn("cas", godip.Sea).Conn("cos", godip.Coast...).Conn("gal", godip.Sea).Conn("gol", godip.Sea).Flag(godip.Coast...).
		// Valle Magdalena
		Prov("val").Conn("med", godip.Coast...).Conn("gol", godip.Sea).Conn("ecu", godip.Coast...).Conn("mon", godip.Land).Conn("cor", godip.Land).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Concepcion
		Prov("con").Conn("pat", godip.Land).Conn("tie", godip.Coast...).Conn("coa", godip.Sea).Conn("mar", godip.Coast...).Conn("pam", godip.Land).Conn("sat", godip.Land).Flag(godip.Coast...).
		// Coast of Argentina
		Prov("coa").Conn("mar", godip.Sea).Conn("con", godip.Sea).Conn("tie", godip.Sea).Conn("soa", godip.Sea).Conn("mid", godip.Sea).Conn("uru", godip.Sea).Conn("bue", godip.Sea).Flag(godip.Sea).
		// Acre
		Prov("acr").Conn("ror", godip.Land).Conn("man", godip.Land).Conn("mat", godip.Land).Conn("goi", godip.Land).Conn("bel", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bahia da Arica
		Prov("bda").Conn("ant", godip.Sea).Conn("des", godip.Sea).Conn("are", godip.Sea).Conn("lim", godip.Sea).Conn("gal", godip.Sea).Conn("sop", godip.Sea).Conn("isl", godip.Sea).Conn("bdc", godip.Sea).Flag(godip.Sea).
		// Buenos Aires
		Prov("bue").Conn("saf", godip.Land).Conn("gra", godip.Land).Conn("pam", godip.Land).Conn("mar", godip.Coast...).Conn("coa", godip.Sea).Conn("uru", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Golfo da Panama
		Prov("gol").Conn("med", godip.Sea).Conn("pan", godip.Sea).Conn("gal", godip.Sea).Conn("ecu", godip.Sea).Conn("val", godip.Sea).Flag(godip.Sea).
		// Mesopotamia
		Prov("mes").Conn("gra", godip.Land).Conn("saf", godip.Land).Conn("rgd", godip.Land).Conn("par", godip.Land).Conn("pot", godip.Land).Flag(godip.Land).
		// Rio Grande do Sul
		Prov("rgd").Conn("mid", godip.Sea).Conn("rdj", godip.Coast...).Conn("goi", godip.Land).Conn("par", godip.Land).Conn("mes", godip.Land).Conn("saf", godip.Land).Conn("uru", godip.Coast...).Flag(godip.Coast...).
		// Barranquilla
		Prov("bar").Conn("bog", godip.Land).Conn("tru", godip.Coast...).Conn("cas", godip.Sea).Conn("pan", godip.Coast...).Conn("med", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// Scotia Sea
		Prov("sco").Conn("soa", godip.Sea).Conn("tie", godip.Sea).Conn("pat", godip.Sea).Conn("bdc", godip.Sea).Conn("isl", godip.Sea).Conn("sop", godip.Sea).Flag(godip.Sea).
		// Mar del Plata
		Prov("mar").Conn("coa", godip.Sea).Conn("bue", godip.Coast...).Conn("pam", godip.Land).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Brazilian Sea
		Prov("bra").Conn("cas", godip.Sea).Conn("caa", godip.Sea).Conn("guy", godip.Sea).Conn("bel", godip.Sea).Conn("mid", godip.Sea).Conn("soa", godip.Sea).Flag(godip.Sea).
		// Antofagasta
		Prov("ant").Conn("bda", godip.Sea).Conn("bdc", godip.Sea).Conn("sat", godip.Coast...).Conn("gra", godip.Land).Conn("pot", godip.Land).Conn("des", godip.Coast...).Flag(godip.Coast...).SC(Chile).
		// Mato Grosso
		Prov("mat").Conn("man", godip.Land).Conn("lap", godip.Land).Conn("yun", godip.Land).Conn("goi", godip.Land).Conn("acr", godip.Land).Flag(godip.Land).
		// Orinoco Springs
		Prov("ori").Conn("tru", godip.Land).Conn("bog", godip.Land).Conn("ama", godip.Land).Conn("man", godip.Land).Conn("ror", godip.Land).Conn("caa", godip.Land).Flag(godip.Land).
		// Paraguay
		Prov("par").Conn("lap", godip.Land).Conn("pot", godip.Land).Conn("mes", godip.Land).Conn("rgd", godip.Land).Conn("goi", godip.Land).Conn("yun", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"pam": "Pampas",
	"bel": "Belem",
	"cor": "Cordillera Oriental",
	"sat": "Santiago",
	"ror": "Roraima",
	"soa": "Southwest Atlantic",
	"uru": "Uruguay",
	"tru": "Trujillo",
	"rdj": "Rio de Janeiro",
	"bdc": "Bahia da Coquimbo",
	"mon": "Montana",
	"cos": "Costa Rica",
	"mid": "Mid Atlantic Ocean",
	"chi": "Chiclayo",
	"pat": "Patagonia",
	"pot": "Potosi",
	"bog": "Bogota",
	"ecu": "Ecuador",
	"gra": "Gran Chaco",
	"ama": "Amazon Basin",
	"goi": "Goias",
	"caa": "Caracas",
	"isl": "Islas Juan Fernandez",
	"med": "Medellin",
	"lap": "La Paz",
	"man": "Manaus",
	"sop": "Southeast Pacific",
	"gal": "Galapagos Sea",
	"saf": "Santa Fe",
	"cas": "Caribbean Sea",
	"lim": "Lima",
	"yun": "Yungaz",
	"guy": "Guyana",
	"are": "Arequipa",
	"tie": "Tierra del Fuego",
	"des": "Desierto Atacama",
	"pan": "Panama",
	"val": "Valle Magdalena",
	"con": "Concepcion",
	"coa": "Coast of Argentina",
	"acr": "Acre",
	"bda": "Bahia da Arica",
	"bue": "Buenos Aires",
	"gol": "Golfo da Panama",
	"mes": "Mesopotamia",
	"rgd": "Rio Grande do Sul",
	"bar": "Barranquilla",
	"sco": "Scotia Sea",
	"mar": "Mar del Plata",
	"bra": "Brazilian Sea",
	"ant": "Antofagasta",
	"mat": "Mato Grosso",
	"ori": "Orinoco Springs",
	"par": "Paraguay",
}
