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
		// Old Bahama Channel
		Prov("old").Conn("tur", godip.Sea).Conn("bah", godip.Sea).Conn("flo", godip.Sea).Conn("gue", godip.Sea).Conn("gul", godip.Sea).Conn("win", godip.Sea).Conn("art", godip.Sea).Conn("cib", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// New York Bight
		Prov("nyb").Conn("che", godip.Sea).Conn("sar", godip.Sea).Conn("gua", godip.Sea).Conn("lis", godip.Sea).Conn("loi", godip.Sea).Conn("del", godip.Sea).Flag(godip.Sea).
		// Artibonite
		Prov("art").Conn("azu", godip.Land).Conn("cib", godip.Coast...).Conn("old", godip.Sea).Conn("win", godip.Sea).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// Chesapeake Bay
		Prov("che").Conn("out", godip.Sea).Conn("sar", godip.Sea).Conn("nyb", godip.Sea).Conn("del", godip.Sea).Flag(godip.Sea).
		// New Orleans
		Prov("neo").Conn("gue", godip.Sea).Conn("wes", godip.Coast...).Conn("qua", godip.Land).Flag(godip.Coast...).
		// Quapow
		Prov("qua").Conn("neo", godip.Land).Conn("wes", godip.Land).Conn("cho", godip.Land).Conn("sal", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Atlantic Ocean
		Prov("atl").Conn("gua", godip.Sea).Conn("sar", godip.Sea).Conn("sad", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// Windward Passage
		Prov("win").Conn("por", godip.Sea).Conn("art", godip.Sea).Conn("old", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// Choctaw
		Prov("cho").Conn("cad", godip.Land).Conn("sal", godip.Land).Conn("qua", godip.Land).Conn("wes", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Gulf of Maine
		Prov("gua").Conn("lis", godip.Sea).Conn("nyb", godip.Sea).Conn("sar", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Cibao
		Prov("cib").Conn("sar", godip.Sea).Conn("old", godip.Sea).Conn("art", godip.Coast...).Conn("azu", godip.Land).Conn("sad", godip.Coast...).Flag(godip.Coast...).
		// Missouri
		Prov("mis").Conn("osa", godip.Land).Conn("sal", godip.Land).Conn("ill", godip.Land).Flag(godip.Land).
		// Port au Prince
		Prov("por").Conn("gul", godip.Sea).Conn("azu", godip.Coast...).Conn("art", godip.Coast...).Conn("win", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Saint Domingue
		Prov("sad").Conn("gul", godip.Sea).Conn("atl", godip.Sea).Conn("sar", godip.Sea).Conn("cib", godip.Coast...).Conn("azu", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Delaware Bay
		Prov("del").Conn("che", godip.Sea).Conn("nyb", godip.Sea).Flag(godip.Sea).
		// West Florida
		Prov("wes").Conn("neo", godip.Coast...).Conn("gue", godip.Sea).Conn("flo", godip.Sea).Conn("cad", godip.Coast...).Conn("cho", godip.Land).Conn("qua", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Caddo
		Prov("cad").Conn("cho", godip.Land).Conn("wes", godip.Coast...).Conn("flo", godip.Sea).Flag(godip.Coast...).
		// Gulf of Gonave
		Prov("gul").Conn("atl", godip.Sea).Conn("sad", godip.Sea).Conn("azu", godip.Sea).Conn("por", godip.Sea).Conn("win", godip.Sea).Conn("old", godip.Sea).Conn("gue", godip.Sea).Flag(godip.Sea).
		// Long Island Sound
		Prov("lis").Conn("loi", godip.Sea).Conn("nyb", godip.Sea).Conn("gua", godip.Sea).Flag(godip.Sea).
		// Osage
		Prov("osa").Conn("qua", godip.Land).Conn("sal", godip.Land).Conn("mis", godip.Land).Flag(godip.Land).
		// Florida Bight
		Prov("flo").Conn("wes", godip.Sea).Conn("gue", godip.Sea).Conn("old", godip.Sea).Conn("bah", godip.Sea).Conn("cad", godip.Sea).Flag(godip.Sea).
		// Azua
		Prov("azu").Conn("art", godip.Land).Conn("por", godip.Coast...).Conn("gul", godip.Sea).Conn("sad", godip.Coast...).Conn("cib", godip.Land).Flag(godip.Coast...).
		// Bahama Banks
		Prov("bah").Conn("tur", godip.Sea).Conn("sar", godip.Sea).Conn("geo", godip.Sea).Conn("flo", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Saint Louis
		Prov("sal").Conn("qua", godip.Land).Conn("cho", godip.Land).Conn("ill", godip.Land).Conn("mis", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).SC(SouthCarolina).
		// Illiniwek
		Prov("ill").Conn("mis", godip.Land).Conn("sal", godip.Land).Flag(godip.Land).
		// Turks and Caicos
		Prov("tur").Conn("old", godip.Sea).Conn("sar", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Long Island
		Prov("loi").Conn("lis", godip.Sea).Conn("nyb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Georgia Bight
		Prov("geo").Conn("out", godip.Sea).Conn("bah", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// Outer Banks
		Prov("out").Conn("geo", godip.Sea).Conn("sar", godip.Sea).Conn("che", godip.Sea).Flag(godip.Sea).
		// Sargasso Sea
		Prov("sar").Conn("cib", godip.Sea).Conn("sad", godip.Sea).Conn("atl", godip.Sea).Conn("gua", godip.Sea).Conn("nyb", godip.Sea).Conn("che", godip.Sea).Conn("out", godip.Sea).Conn("geo", godip.Sea).Conn("bah", godip.Sea).Conn("tur", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Gulf of Mexico
		Prov("gue").Conn("gul", godip.Sea).Conn("old", godip.Sea).Conn("flo", godip.Sea).Conn("wes", godip.Sea).Conn("neo", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"old": "Old Bahama Channel",
	"nyb": "New York Bight",
	"art": "Artibonite",
	"che": "Chesapeake Bay",
	"neo": "New Orleans",
	"qua": "Quapow",
	"atl": "Atlantic Ocean",
	"win": "Windward Passage",
	"cho": "Choctaw",
	"gua": "Gulf of Maine",
	"cib": "Cibao",
	"mis": "Missouri",
	"por": "Port au Prince",
	"sad": "Saint Domingue",
	"del": "Delaware Bay",
	"wes": "West Florida",
	"cad": "Caddo",
	"gul": "Gulf of Gonave",
	"lis": "Long Island Sound",
	"osa": "Osage",
	"flo": "Florida Bight",
	"azu": "Azua",
	"bah": "Bahama Banks",
	"sal": "Saint Louis",
	"ill": "Illiniwek",
	"tur": "Turks and Caicos",
	"loi": "Long Island",
	"geo": "Georgia Bight",
	"out": "Outer Banks",
	"sar": "Sargasso Sea",
	"gue": "Gulf of Mexico",
}
