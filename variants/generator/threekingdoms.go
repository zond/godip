package threekingdoms

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Wei godip.Nation = "Wei"
	Shu godip.Nation = "Shu"
)

var Nations = []godip.Nation{Wei, Shu}

var ThreeKingdomsVariant = common.Variant{
	Name:              "ThreeKingdoms",
	Graph:             func() godip.Graph { return ThreeKingdomsGraph() },
	Start:             ThreeKingdomsStart,
	Blank:             ThreeKingdomsBlank,
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
		return Asset("svg/threekingdomsmap.svg")
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

func ThreeKingdomsBlank(phase godip.Phase) *state.State {
	return state.New(ThreeKingdomsGraph(), phase, classical.BackupRule, nil, nil)
}

func ThreeKingdomsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(220, godip.Spring, godip.Movement)
	result = ThreeKingdomsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"jzh": godip.Unit{godip.Army, Wei},
		"yon": godip.Unit{godip.Army, Shu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"jzh": Wei,
		"yon": Shu,
	})
	return
}

func ThreeKingdomsGraph() *graph.Graph {
	return graph.New().
		// Jiaozhi
		Prov("jzh").Conn("yon", godip.Land).Flag(godip.Land).SC(Wei).
		// Yongchang
		Prov("yon").Conn("jzh", godip.Land).Flag(godip.Land).SC(Shu).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"jzh": "Jiaozhi",
	"yon": "Yongchang",
}
