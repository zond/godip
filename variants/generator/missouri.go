package missouri

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

var MissouriVariant = common.Variant{
	Name:              "Missouri",
	Graph:             func() godip.Graph { return MissouriGraph() },
	Start:             MissouriStart,
	Blank:             MissouriBlank,
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
		return Asset("svg/missourimap.svg")
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

func MissouriBlank(phase godip.Phase) *state.State {
	return state.New(MissouriGraph(), phase, classical.BackupRule, nil, nil)
}

func MissouriStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(2024, godip.Spring, godip.Movement)
	result = MissouriBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"and": godip.Unit{godip.Army, Missouria},
		"atc": godip.Unit{godip.Army, Oto},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"and": Missouria,
		"atc": Oto,
	})
	return
}

func MissouriGraph() *graph.Graph {
	return graph.New().
		// Mercer
		Prov("mer").Conn("gru", godip.Land).Conn("MIO", godip.Land).Conn("MIO", godip.Land).Conn("har", godip.Land).Flag(godip.Land).
		// Atchison
		Prov("atc").Conn("hol", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).SC(Oto).
		// Harrison
		Prov("har").Conn("wor", godip.Land).Conn("gen", godip.Land).Conn("dav", godip.Land).Conn("gru", godip.Land).Conn("mer", godip.Land).Flag(godip.Land).
		// Gentry
		Prov("gen").Conn("dav", godip.Land).Conn("har", godip.Land).Conn("wor", godip.Land).Conn("nod", godip.Land).Conn("and", godip.Land).Conn("dek", godip.Land).Flag(godip.Land).
		// Missouri
		Prov("MIO").Conn("ray", godip.Land).Conn("ray", godip.Land).Conn("cla", godip.Land).Conn("mer", godip.Land).Conn("mer", godip.Land).Conn("gru", godip.Land).Conn("gru", godip.Land).Conn("liv", godip.Land).Conn("liv", godip.Land).Conn("car", godip.Land).Conn("car", godip.Land).Conn("car", godip.Land).Flag(godip.Land).
		// Buchanan
		Prov("buc").Conn("pla", godip.Land).Conn("cli", godip.Land).Conn("dek", godip.Land).Conn("and", godip.Land).Flag(godip.Land).
		// Holt
		Prov("hol").Conn("atc", godip.Land).Conn("and", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).
		// Carroll
		Prov("car").Conn("MIO", godip.Land).Conn("MIO", godip.Land).Conn("liv", godip.Land).Conn("cal", godip.Land).Conn("ray", godip.Land).Conn("MIO", godip.Land).Flag(godip.Land).
		// Andrew
		Prov("and").Conn("buc", godip.Land).Conn("dek", godip.Land).Conn("gen", godip.Land).Conn("nod", godip.Land).Conn("hol", godip.Land).Flag(godip.Land).SC(Missouria).
		// Livingston
		Prov("liv").Conn("MIO", godip.Land).Conn("gru", godip.Land).Conn("dav", godip.Land).Conn("cal", godip.Land).Conn("car", godip.Land).Conn("MIO", godip.Land).Flag(godip.Land).
		// Platte
		Prov("pla").Conn("cla", godip.Land).Conn("cli", godip.Land).Conn("buc", godip.Land).Flag(godip.Land).
		// Clay
		Prov("cla").Conn("MIO", godip.Land).Conn("ray", godip.Land).Conn("cli", godip.Land).Conn("pla", godip.Land).Flag(godip.Land).
		// De Kalb
		Prov("dek").Conn("cal", godip.Land).Conn("dav", godip.Land).Conn("gen", godip.Land).Conn("and", godip.Land).Conn("buc", godip.Land).Conn("cli", godip.Land).Flag(godip.Land).
		// Caldwell
		Prov("cal").Conn("dek", godip.Land).Conn("cli", godip.Land).Conn("ray", godip.Land).Conn("car", godip.Land).Conn("liv", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).
		// Clinton
		Prov("cli").Conn("buc", godip.Land).Conn("pla", godip.Land).Conn("cla", godip.Land).Conn("ray", godip.Land).Conn("cal", godip.Land).Conn("dek", godip.Land).Flag(godip.Land).
		// Ray
		Prov("ray").Conn("MIO", godip.Land).Conn("car", godip.Land).Conn("cal", godip.Land).Conn("cli", godip.Land).Conn("cla", godip.Land).Conn("MIO", godip.Land).Flag(godip.Land).
		// Grundy
		Prov("gru").Conn("liv", godip.Land).Conn("MIO", godip.Land).Conn("MIO", godip.Land).Conn("mer", godip.Land).Conn("har", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).
		// Daviess
		Prov("dav").Conn("cal", godip.Land).Conn("liv", godip.Land).Conn("gru", godip.Land).Conn("har", godip.Land).Conn("gen", godip.Land).Conn("dek", godip.Land).Flag(godip.Land).
		// Worth
		Prov("wor").Conn("har", godip.Land).Conn("nod", godip.Land).Conn("gen", godip.Land).Flag(godip.Land).
		// Nodaway
		Prov("nod").Conn("atc", godip.Land).Conn("hol", godip.Land).Conn("and", godip.Land).Conn("gen", godip.Land).Conn("wor", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"mer": "Mercer",
	"atc": "Atchison",
	"har": "Harrison",
	"gen": "Gentry",
	"MIO": "Missouri",
	"buc": "Buchanan",
	"hol": "Holt",
	"car": "Carroll",
	"and": "Andrew",
	"liv": "Livingston",
	"pla": "Platte",
	"cla": "Clay",
	"dek": "De Kalb",
	"cal": "Caldwell",
	"cli": "Clinton",
	"ray": "Ray",
	"gru": "Grundy",
	"dav": "Daviess",
	"wor": "Worth",
	"nod": "Nodaway",
}
