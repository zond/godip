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
	SoloWinner:        common.SCCountWinner(2),
	SoloSCCount:       func(*state.State) int { return 2 },
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
		// Guyana
		Prov("guy").Conn("bra", godip.Sea).Conn("sou", godip.Sea).Conn("sou", godip.Sea).Conn("mid", godip.Coast...).Flag(godip.Coast...).SC(Columbia).
		// Southwest Atlantic
		Prov("sou").Conn("bra", godip.Sea).Conn("bel", godip.Sea).Conn("mid", godip.Sea).Conn("mid", godip.Sea).Conn("guy", godip.Sea).Conn("guy", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// Brazilian Sea
		Prov("bra").Conn("sou", godip.Sea).Conn("guy", godip.Sea).Conn("mid", godip.Sea).Conn("bel", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Belem
		Prov("bel").Conn("bra", godip.Sea).Conn("mid", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Mid Atlantic Ocean
		Prov("mid").Conn("bra", godip.Sea).Conn("guy", godip.Coast...).Conn("sou", godip.Sea).Conn("sou", godip.Sea).Conn("bel", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"guy": "Guyana",
	"sou": "Southwest Atlantic",
	"bra": "Brazilian Sea",
	"bel": "Belem",
	"mid": "Mid Atlantic Ocean",
}
