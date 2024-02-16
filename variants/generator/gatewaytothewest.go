package gatewaytothewest

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Missouria godip.Nation = "Missouria"
	Oto       godip.Nation = "Oto"
)

var Nations = []godip.Nation{Missouria, Oto}

var GatewayToTheWestVariant = common.Variant{
	Name:              "GatewayToTheWest",
	Graph:             func() godip.Graph { return GatewayToTheWestGraph() },
	Start:             GatewayToTheWestStart,
	Blank:             GatewayToTheWestBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(1),
	SoloSCCount:       func(*state.State) int { return 1 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/gatewaytothewestmap.svg")
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

func GatewayToTheWestBlank(phase godip.Phase) *state.State {
	return state.New(GatewayToTheWestGraph(), phase, classical.BackupRule, nil, nil)
}

func GatewayToTheWestStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(2024, godip.Spring, godip.Movement)
	result = GatewayToTheWestBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"neo": godip.Unit{godip.Army, Missouria},
		"pla": godip.Unit{godip.Army, Oto},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"neo": Missouria,
	})
	return
}

func GatewayToTheWestGraph() *graph.Graph {
	return graph.New().
		// North Mississippi River
		Prov("nor").Conn("sou", godip.Sea).Conn("pla", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// Platte
		Prov("pla").Conn("nwm", godip.Sea).Conn("nwm", godip.Sea).Conn("wes", godip.Sea).Conn("nor", godip.Sea).Flag(godip.Coast...).
		// West Missouri River
		Prov("wes").Conn("neo", godip.Sea).Conn("sou", godip.Sea).Conn("nor", godip.Sea).Conn("pla", godip.Sea).Conn("nwm", godip.Sea).Flag(godip.Sea).
		// South Mississippi River
		Prov("sou").Conn("nor", godip.Sea).Conn("wes", godip.Sea).Conn("neo", godip.Sea).Conn("neo", godip.Sea).Flag(godip.Sea).
		// Neosho
		Prov("neo").Conn("sou", godip.Sea).Conn("sou", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(Missouria).
		// North West Missouri River
		Prov("nwm").Conn("wes", godip.Sea).Conn("pla", godip.Sea).Conn("pla", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"nor": "North Mississippi River",
	"pla": "Platte",
	"wes": "West Missouri River",
	"sou": "South Mississippi River",
	"neo": "Neosho",
	"nwm": "North West Missouri River",
}
