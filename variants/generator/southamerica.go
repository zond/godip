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
	SoloWinner:        common.SCCountWinner(4),
	SoloSCCount:       func(*state.State) int { return 4 },
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
		"bel": godip.Unit{godip.Army, Brazil},
		"guy": godip.Unit{godip.Fleet, Columbia},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"guy": Columbia,
	})
	return
}

func SouthAmericaGraph() *graph.Graph {
	return graph.New().
		// Belem
		Prov("bel").Conn("mid", godip.Sea).Conn("sou", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Caribbean Sea
		Prov("cas").Conn("cos", godip.Sea).Conn("pan", godip.Sea).Conn("bar", godip.Sea).Conn("tru", godip.Sea).Conn("caa", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Mid Atlantic Ocean
		Prov("mid").Conn("sou", godip.Sea).Conn("bel", godip.Sea).Conn("bra", godip.Sea).Conn("guy", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Trujillo
		Prov("tru").Conn("bar", godip.Coast...).Conn("bog", godip.Land).Conn("ori", godip.Land).Conn("caa", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).
		// Barranquilla
		Prov("bar").Conn("tru", godip.Coast...).Conn("cas", godip.Sea).Conn("pan", godip.Coast...).Conn("med", godip.Land).Conn("bog", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Cordillera Oriental
		Prov("cor").Conn("bog", godip.Coast...).Conn("val", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).
		// Brazilian Sea
		Prov("bra").Conn("cas", godip.Sea).Conn("caa", godip.Sea).Conn("guy", godip.Sea).Conn("mid", godip.Sea).Conn("bel", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Bogota
		Prov("bog").Conn("ori", godip.Coast...).Conn("tru", godip.Land).Conn("bar", godip.Land).Conn("med", godip.Coast...).Conn("val", godip.Coast...).Conn("cor", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Southwest Atlantic
		Prov("sou").Conn("bra", godip.Sea).Conn("bel", godip.Sea).Conn("mid", godip.Sea).Conn("mid", godip.Sea).Conn("guy", godip.Sea).Conn("caa", godip.Sea).Conn("ori", godip.Sea).Conn("ori", godip.Sea).Conn("bog", godip.Sea).Conn("cor", godip.Sea).Conn("val", godip.Sea).Conn("med", godip.Sea).Conn("pan", godip.Sea).Conn("cos", godip.Sea).Flag(godip.Sea).
		// Panama
		Prov("pan").Conn("med", godip.Coast...).Conn("bar", godip.Coast...).Conn("cas", godip.Sea).Conn("cos", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).
		// Caracas
		Prov("caa").Conn("sou", godip.Sea).Conn("guy", godip.Coast...).Conn("bra", godip.Sea).Conn("cas", godip.Sea).Conn("tru", godip.Coast...).Conn("ori", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Guyana
		Prov("guy").Conn("sou", godip.Sea).Conn("mid", godip.Coast...).Conn("bra", godip.Sea).Conn("caa", godip.Coast...).Flag(godip.Coast...).SC(Columbia).
		// Medellin
		Prov("med").Conn("pan", godip.Coast...).Conn("sou", godip.Sea).Conn("val", godip.Coast...).Conn("bog", godip.Coast...).Conn("bar", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Valle Magdalena
		Prov("val").Conn("sou", godip.Sea).Conn("cor", godip.Coast...).Conn("bog", godip.Coast...).Conn("med", godip.Coast...).Flag(godip.Coast...).
		// Orinoco Springs
		Prov("ori").Conn("sou", godip.Sea).Conn("caa", godip.Coast...).Conn("tru", godip.Land).Conn("bog", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).
		// Costa Rica
		Prov("cos").Conn("sou", godip.Sea).Conn("pan", godip.Coast...).Conn("cas", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"bel": "Belem",
	"cas": "Caribbean Sea",
	"mid": "Mid Atlantic Ocean",
	"tru": "Trujillo",
	"bar": "Barranquilla",
	"cor": "Cordillera Oriental",
	"bra": "Brazilian Sea",
	"bog": "Bogota",
	"sou": "Southwest Atlantic",
	"pan": "Panama",
	"caa": "Caracas",
	"guy": "Guyana",
	"med": "Medellin",
	"val": "Valle Magdalena",
	"ori": "Orinoco Springs",
	"cos": "Costa Rica",
}
