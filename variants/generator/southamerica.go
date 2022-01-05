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
)

var Nations = []godip.Nation{Brazil, Columbia}

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
	SoloWinner:        common.SCCountWinner(5),
	SoloSCCount:       func(*state.State) int { return 5 },
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
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"bel": Brazil,
		"rdj": Brazil,
		"bar": Columbia,
		"med": Columbia,
		"bog": Columbia,
	})
	return
}

func SouthAmericaGraph() *graph.Graph {
	return graph.New().
		// Belem
		Prov("bel").Conn("ror", godip.Land).Conn("acr", godip.Land).Conn("goi", godip.Land).Conn("rdj", godip.Coast...).Conn("sou", godip.Sea).Conn("bra", godip.Sea).Conn("guy", godip.Coast...).Flag(godip.Coast...).SC(Brazil).
		// Cordillera Oriental
		Prov("cor").Conn("val", godip.Coast...).Conn("mid", godip.Sea).Conn("ama", godip.Coast...).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Roraima
		Prov("ror").Conn("bel", godip.Land).Conn("guy", godip.Land).Conn("caa", godip.Land).Conn("ori", godip.Land).Conn("man", godip.Land).Conn("acr", godip.Land).Flag(godip.Land).
		// Rio de Janeiro
		Prov("rdj").Conn("bel", godip.Coast...).Conn("goi", godip.Land).Conn("rgd", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).SC(Brazil).
		// Costa Rica
		Prov("cos").Conn("mid", godip.Sea).Conn("pan", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Mid Atlantic Ocean
		Prov("mid").Conn("bra", godip.Sea).Conn("sou", godip.Sea).Conn("rgd", godip.Sea).Conn("goi", godip.Sea).Conn("mat", godip.Sea).Conn("man", godip.Sea).Conn("ama", godip.Sea).Conn("cor", godip.Sea).Conn("val", godip.Sea).Conn("med", godip.Sea).Conn("pan", godip.Sea).Conn("cos", godip.Sea).Flag(godip.Sea).
		// Mato Grosso
		Prov("mat").Conn("man", godip.Coast...).Conn("mid", godip.Sea).Conn("goi", godip.Coast...).Conn("acr", godip.Land).Flag(godip.Coast...).
		// Bogota
		Prov("bog").Conn("ori", godip.Land).Conn("tru", godip.Land).Conn("bar", godip.Land).Conn("med", godip.Land).Conn("val", godip.Land).Conn("cor", godip.Land).Conn("ama", godip.Land).Flag(godip.Land).SC(Columbia).
		// Amazon Basin
		Prov("ama").Conn("cor", godip.Coast...).Conn("mid", godip.Sea).Conn("man", godip.Coast...).Conn("ori", godip.Land).Conn("bog", godip.Land).Flag(godip.Coast...).
		// Goias
		Prov("goi").Conn("acr", godip.Land).Conn("mat", godip.Coast...).Conn("mid", godip.Sea).Conn("rgd", godip.Coast...).Conn("rdj", godip.Land).Conn("bel", godip.Land).Flag(godip.Coast...).
		// Caracas
		Prov("caa").Conn("ror", godip.Land).Conn("guy", godip.Coast...).Conn("bra", godip.Sea).Conn("cas", godip.Sea).Conn("tru", godip.Coast...).Conn("ori", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Manaus
		Prov("man").Conn("mat", godip.Coast...).Conn("acr", godip.Land).Conn("ror", godip.Land).Conn("ori", godip.Land).Conn("ama", godip.Coast...).Conn("mid", godip.Sea).Flag(godip.Coast...).
		// Caribbean Sea
		Prov("cas").Conn("cos", godip.Sea).Conn("pan", godip.Sea).Conn("bar", godip.Sea).Conn("tru", godip.Sea).Conn("caa", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Medellin
		Prov("med").Conn("pan", godip.Coast...).Conn("mid", godip.Sea).Conn("val", godip.Coast...).Conn("bog", godip.Land).Conn("bar", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// Trujillo
		Prov("tru").Conn("bar", godip.Coast...).Conn("bog", godip.Land).Conn("ori", godip.Land).Conn("caa", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).
		// Guyana
		Prov("guy").Conn("caa", godip.Coast...).Conn("ror", godip.Land).Conn("bel", godip.Coast...).Conn("bra", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Panama
		Prov("pan").Conn("med", godip.Coast...).Conn("bar", godip.Coast...).Conn("cas", godip.Sea).Conn("cos", godip.Coast...).Conn("mid", godip.Sea).Flag(godip.Coast...).
		// Valle Magdalena
		Prov("val").Conn("mid", godip.Sea).Conn("cor", godip.Coast...).Conn("bog", godip.Land).Conn("med", godip.Coast...).Flag(godip.Coast...).
		// Southwest Atlantic
		Prov("sou").Conn("mid", godip.Sea).Conn("bra", godip.Sea).Conn("bel", godip.Sea).Conn("rdj", godip.Sea).Conn("rgd", godip.Sea).Flag(godip.Sea).
		// Acre
		Prov("acr").Conn("goi", godip.Land).Conn("bel", godip.Land).Conn("ror", godip.Land).Conn("man", godip.Land).Conn("mat", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Rio Grande do Sul
		Prov("rgd").Conn("mid", godip.Sea).Conn("sou", godip.Sea).Conn("rdj", godip.Coast...).Conn("goi", godip.Coast...).Flag(godip.Coast...).
		// Barranquilla
		Prov("bar").Conn("tru", godip.Coast...).Conn("cas", godip.Sea).Conn("pan", godip.Coast...).Conn("med", godip.Land).Conn("bog", godip.Land).Flag(godip.Coast...).SC(Columbia).
		// Brazilian Sea
		Prov("bra").Conn("cas", godip.Sea).Conn("caa", godip.Sea).Conn("guy", godip.Sea).Conn("bel", godip.Sea).Conn("sou", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Sea).
		// Orinoco Springs
		Prov("ori").Conn("bog", godip.Land).Conn("ama", godip.Land).Conn("man", godip.Land).Conn("ror", godip.Land).Conn("caa", godip.Land).Conn("tru", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"bel": "Belem",
	"cor": "Cordillera Oriental",
	"ror": "Roraima",
	"rdj": "Rio de Janeiro",
	"cos": "Costa Rica",
	"mid": "Mid Atlantic Ocean",
	"mat": "Mato Grosso",
	"bog": "Bogota",
	"ama": "Amazon Basin",
	"goi": "Goias",
	"caa": "Caracas",
	"man": "Manaus",
	"cas": "Caribbean Sea",
	"med": "Medellin",
	"tru": "Trujillo",
	"guy": "Guyana",
	"pan": "Panama",
	"val": "Valle Magdalena",
	"sou": "Southwest Atlantic",
	"acr": "Acre",
	"rgd": "Rio Grande do Sul",
	"bar": "Barranquilla",
	"bra": "Brazilian Sea",
	"ori": "Orinoco Springs",
}
