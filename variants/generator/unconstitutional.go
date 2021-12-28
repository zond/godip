package unconstitutional

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Pennsylvania  godip.Nation = "Pennsylvania"
	SouthCarolina godip.Nation = "SouthCarolina"
)

var Nations = []godip.Nation{Pennsylvania, SouthCarolina}

var UnconstitutionalVariant = common.Variant{
	Name:              "Unconstitutional",
	Graph:             func() godip.Graph { return UnconstitutionalGraph() },
	Start:             UnconstitutionalStart,
	Blank:             UnconstitutionalBlank,
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
		return Asset("svg/unconstitutionalmap.svg")
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

func UnconstitutionalBlank(phase godip.Phase) *state.State {
	return state.New(UnconstitutionalGraph(), phase, classical.BackupRule, nil, nil)
}

func UnconstitutionalStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1805, godip.Spring, godip.Movement)
	result = UnconstitutionalBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"cho": godip.Unit{godip.Army, Pennsylvania},
		"sal": godip.Unit{godip.Army, SouthCarolina},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"cho": Pennsylvania,
		"sal": SouthCarolina,
	})
	return
}

func UnconstitutionalGraph() *graph.Graph {
	return graph.New().
		// Artibonite
		Prov("art").Conn("azu", godip.Land).Conn("cib", godip.Coast...).Conn("old", godip.Sea).Conn("win", godip.Sea).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// New Orleans
		Prov("new").Conn("gue", godip.Sea).Conn("wes", godip.Coast...).Conn("qua", godip.Land).Flag(godip.Coast...).
		// Quapow
		Prov("qua").Conn("new", godip.Land).Conn("wes", godip.Land).Conn("cho", godip.Land).Conn("sal", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Windward Passage
		Prov("win").Conn("por", godip.Sea).Conn("art", godip.Sea).Conn("old", godip.Sea).Conn("gof", godip.Sea).Flag(godip.Sea).
		// Choctaw
		Prov("cho").Conn("cad", godip.Coast...).Conn("gua", godip.Sea).Conn("sal", godip.Coast...).Conn("qua", godip.Land).Conn("wes", godip.Land).Flag(godip.Coast...).SC(Pennsylvania).
		// Illiniwek
		Prov("ill").Conn("gua", godip.Sea).Conn("mis", godip.Coast...).Conn("sal", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Maine
		Prov("gua").Conn("mis", godip.Sea).Conn("ill", godip.Sea).Conn("sal", godip.Sea).Conn("cho", godip.Sea).Conn("cad", godip.Sea).Conn("flo", godip.Sea).Conn("bah", godip.Sea).Conn("sar", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Cibao
		Prov("cib").Conn("sar", godip.Sea).Conn("old", godip.Sea).Conn("art", godip.Coast...).Conn("azu", godip.Land).Conn("sad", godip.Coast...).Flag(godip.Coast...).
		// Port au Prince
		Prov("por").Conn("gof", godip.Sea).Conn("azu", godip.Coast...).Conn("art", godip.Coast...).Conn("win", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Saint Domingue
		Prov("sad").Conn("gof", godip.Sea).Conn("atl", godip.Sea).Conn("sar", godip.Sea).Conn("cib", godip.Coast...).Conn("azu", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Turks and Caicos
		Prov("tur").Conn("old", godip.Sea).Conn("sar", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// West Florida
		Prov("wes").Conn("new", godip.Coast...).Conn("gue", godip.Sea).Conn("flo", godip.Sea).Conn("cad", godip.Coast...).Conn("cho", godip.Land).Conn("qua", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Caddo
		Prov("cad").Conn("cho", godip.Coast...).Conn("wes", godip.Coast...).Conn("flo", godip.Sea).Conn("gua", godip.Sea).Flag(godip.Coast...).
		// Gulf of Gonave
		Prov("gof").Conn("atl", godip.Sea).Conn("sad", godip.Sea).Conn("azu", godip.Sea).Conn("por", godip.Sea).Conn("win", godip.Sea).Conn("old", godip.Sea).Conn("gue", godip.Sea).Flag(godip.Sea).
		// Missouri
		Prov("mis").Conn("osa", godip.Land).Conn("sal", godip.Coast...).Conn("ill", godip.Coast...).Conn("gua", godip.Sea).Flag(godip.Coast...).
		// Osage
		Prov("osa").Conn("qua", godip.Land).Conn("sal", godip.Land).Conn("mis", godip.Land).Flag(godip.Land).
		// Florida Bight
		Prov("flo").Conn("wes", godip.Sea).Conn("gue", godip.Sea).Conn("old", godip.Sea).Conn("bah", godip.Sea).Conn("gua", godip.Sea).Conn("cad", godip.Sea).Flag(godip.Sea).
		// Azua
		Prov("azu").Conn("gof", godip.Sea).Conn("sad", godip.Coast...).Conn("cib", godip.Land).Conn("art", godip.Land).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// Bahama Banks
		Prov("bah").Conn("tur", godip.Sea).Conn("sar", godip.Sea).Conn("gua", godip.Sea).Conn("flo", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Saint Louis
		Prov("sal").Conn("qua", godip.Land).Conn("cho", godip.Coast...).Conn("gua", godip.Sea).Conn("ill", godip.Coast...).Conn("mis", godip.Coast...).Conn("osa", godip.Land).Flag(godip.Coast...).SC(SouthCarolina).
		// Gulf of Mexico
		Prov("gue").Conn("gof", godip.Sea).Conn("old", godip.Sea).Conn("flo", godip.Sea).Conn("wes", godip.Sea).Conn("new", godip.Sea).Flag(godip.Sea).
		// Old Bahama Channel
		Prov("old").Conn("tur", godip.Sea).Conn("bah", godip.Sea).Conn("flo", godip.Sea).Conn("gue", godip.Sea).Conn("gof", godip.Sea).Conn("win", godip.Sea).Conn("art", godip.Sea).Conn("cib", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// Atlantic Ocean
		Prov("atl").Conn("gua", godip.Sea).Conn("sar", godip.Sea).Conn("sad", godip.Sea).Conn("gof", godip.Sea).Flag(godip.Sea).
		// Sargasso Sea
		Prov("sar").Conn("tur", godip.Sea).Conn("old", godip.Sea).Conn("cib", godip.Sea).Conn("sad", godip.Sea).Conn("atl", godip.Sea).Conn("gua", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"art": "Artibonite",
	"new": "New Orleans",
	"qua": "Quapow",
	"win": "Windward Passage",
	"cho": "Choctaw",
	"ill": "Illiniwek",
	"gua": "Gulf of Maine",
	"cib": "Cibao",
	"por": "Port au Prince",
	"sad": "Saint Domingue",
	"tur": "Turks and Caicos",
	"wes": "West Florida",
	"cad": "Caddo",
	"gof": "Gulf of Gonave",
	"mis": "Missouri",
	"osa": "Osage",
	"flo": "Florida Bight",
	"azu": "Azua",
	"bah": "Bahama Banks",
	"sal": "Saint Louis",
	"gue": "Gulf of Mexico",
	"old": "Old Bahama Channel",
	"atl": "Atlantic Ocean",
	"sar": "Sargasso Sea",
}
