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
	SoloWinner:        common.SCCountWinner(2),
	SoloSCCount:       func(*state.State) int { return 2 },
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
		"sai": godip.Unit{godip.Army, SouthCarolina},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"cho": Pennsylvania,
		"sai": SouthCarolina,
	})
	return
}

func UnconstitutionalGraph() *graph.Graph {
	return graph.New().
		// Windward Passage
		Prov("win").Conn("atl", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// Choctaw
		Prov("cho").Conn("qua", godip.Land).Conn("wes", godip.Land).Conn("cad", godip.Land).Conn("atl", godip.Land).Conn("sai", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Caddo
		Prov("cad").Conn("atl", godip.Coast...).Conn("cho", godip.Land).Conn("wes", godip.Coast...).Conn("flo", godip.Sea).Flag(godip.Coast...).
		// Missouri
		Prov("mis").Conn("osa", godip.Land).Conn("sai", godip.Land).Conn("ill", godip.Land).Conn("atl", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("qua", godip.Land).Conn("sai", godip.Land).Conn("mis", godip.Land).Flag(godip.Land).
		// Illiniwek
		Prov("ill").Conn("atl", godip.Land).Conn("mis", godip.Land).Conn("sai", godip.Land).Flag(godip.Land).
		// Saint Louis
		Prov("sai").Conn("qua", godip.Land).Conn("cho", godip.Land).Conn("atl", godip.Land).Conn("ill", godip.Land).Conn("mis", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).SC(SouthCarolina).
		// New Orleans
		Prov("new").Conn("gul", godip.Sea).Conn("wes", godip.Coast...).Conn("qua", godip.Land).Flag(godip.Coast...).
		// Florida Bight
		Prov("flo").Conn("wes", godip.Sea).Conn("gul", godip.Sea).Conn("atl", godip.Sea).Conn("cad", godip.Sea).Flag(godip.Sea).
		// Quapow
		Prov("qua").Conn("new", godip.Land).Conn("wes", godip.Land).Conn("cho", godip.Land).Conn("sai", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Atlantic Ocean
		Prov("atl").Conn("mis", godip.Land).Conn("ill", godip.Land).Conn("sai", godip.Land).Conn("cho", godip.Land).Conn("cad", godip.Coast...).Conn("flo", godip.Sea).Conn("gul", godip.Sea).Conn("win", godip.Sea).Flag(godip.Coast...).
		// West Florida
		Prov("wes").Conn("flo", godip.Sea).Conn("cad", godip.Coast...).Conn("cho", godip.Land).Conn("qua", godip.Land).Conn("new", godip.Coast...).Conn("gul", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Mexico
		Prov("gul").Conn("win", godip.Sea).Conn("atl", godip.Sea).Conn("flo", godip.Sea).Conn("wes", godip.Sea).Conn("new", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"win": "Windward Passage",
	"cho": "Choctaw",
	"cad": "Caddo",
	"mis": "Missouri",
	"osa": "Osage",
	"ill": "Illiniwek",
	"sai": "Saint Louis",
	"new": "New Orleans",
	"flo": "Florida Bight",
	"qua": "Quapow",
	"atl": "Atlantic Ocean",
	"wes": "West Florida",
	"gul": "Gulf of Mexico",
}
