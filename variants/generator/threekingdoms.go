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
	SoloWinner:        common.SCCountWinner(6),
	SoloSCCount:       func(*state.State) int { return 6 },
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
		"qia": godip.Unit{godip.Army, Wei},
		"zan": godip.Unit{godip.Army, Shu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"qia": Wei,
		"zan": Shu,
	})
	return
}

func ThreeKingdomsGraph() *graph.Graph {
	return graph.New().
		// Cangwu
		Prov("can").Conn("nah", godip.Sea).Conn("hep", godip.Sea).Conn("zan", godip.Sea).Conn("pea", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Runan
		Prov("run").Conn("cen", godip.Sea).Conn("gre", godip.Coast...).Conn("luj", godip.Land).Conn("gua", godip.Coast...).Conn("cen", godip.Sea).Flag(godip.Coast...).
		// Hepu
		Prov("hep").Conn("lin", godip.Land).Conn("wul", godip.Land).Conn("zan", godip.Coast...).Conn("can", godip.Sea).Conn("nah", godip.Coast...).Flag(godip.Coast...).
		// Wuling
		Prov("wul").Conn("bad", godip.Land).Conn("ful", godip.Land).Conn("zan", godip.Land).Conn("hep", godip.Land).Conn("lin", godip.Land).Conn("cha", godip.Land).Flag(godip.Land).
		// Chengdu
		Prov("che").Conn("cen", godip.Sea).Conn("cen", godip.Sea).Conn("qia", godip.Land).Conn("sha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Great Space
		Prov("gre").Conn("cen", godip.Sea).Conn("xia", godip.Coast...).Conn("nan", godip.Sea).Conn("luj", godip.Coast...).Conn("run", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Pearl River
		Prov("pea").Conn("jio", godip.Coast...).Conn("sou", godip.Sea).Conn("can", godip.Sea).Conn("zan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Yuzhang
		Prov("yuz").Conn("lin", godip.Land).Conn("nah", godip.Land).Conn("kua", godip.Land).Conn("jin", godip.Coast...).Conn("eas", godip.Sea).Conn("luj", godip.Coast...).Conn("nan", godip.Sea).Conn("cha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Xiangyang
		Prov("xia").Conn("cen", godip.Sea).Conn("sha", godip.Coast...).Conn("wes", godip.Sea).Conn("bad", godip.Coast...).Conn("nan", godip.Sea).Conn("gre", godip.Coast...).Flag(godip.Coast...).
		// South China Sea
		Prov("sou").Conn("cen", godip.Sea).Conn("cen", godip.Sea).Conn("nah", godip.Sea).Conn("can", godip.Sea).Conn("pea", godip.Sea).Conn("jio", godip.Sea).Flag(godip.Sea).
		// Kuaiji
		Prov("kua").Conn("jin", godip.Coast...).Conn("yuz", godip.Land).Conn("nah", godip.Coast...).Conn("cen", godip.Sea).Flag(godip.Coast...).
		// Lujiang
		Prov("luj").Conn("nan", godip.Sea).Conn("yuz", godip.Coast...).Conn("eas", godip.Sea).Conn("gua", godip.Coast...).Conn("run", godip.Land).Conn("gre", godip.Coast...).Flag(godip.Coast...).
		// Fuling
		Prov("ful").Conn("wul", godip.Land).Conn("bad", godip.Coast...).Conn("wes", godip.Sea).Conn("qia", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// Badong
		Prov("bad").Conn("wes", godip.Sea).Conn("ful", godip.Coast...).Conn("wul", godip.Land).Conn("cha", godip.Coast...).Conn("nan", godip.Sea).Conn("xia", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Zangke
		Prov("zan").Conn("ful", godip.Land).Conn("qia", godip.Land).Conn("yon", godip.Land).Conn("jio", godip.Land).Conn("pea", godip.Coast...).Conn("can", godip.Sea).Conn("hep", godip.Coast...).Conn("wul", godip.Land).Flag(godip.Coast...).SC(Shu).
		// Lingling
		Prov("lin").Conn("hep", godip.Land).Conn("nah", godip.Land).Conn("yuz", godip.Land).Conn("cha", godip.Land).Conn("wul", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Guangling
		Prov("gua").Conn("cen", godip.Sea).Conn("cen", godip.Sea).Conn("run", godip.Coast...).Conn("luj", godip.Coast...).Conn("eas", godip.Sea).Conn("cen", godip.Sea).Conn("cen", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Jianye
		Prov("jin").Conn("kua", godip.Coast...).Conn("cen", godip.Sea).Conn("eas", godip.Sea).Conn("yuz", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shangyong
		Prov("sha").Conn("cen", godip.Sea).Conn("che", godip.Coast...).Conn("qia", godip.Coast...).Conn("wes", godip.Sea).Conn("xia", godip.Coast...).Conn("cen", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Coast...).
		// Nan
		Prov("nan").Conn("luj", godip.Sea).Conn("gre", godip.Sea).Conn("xia", godip.Sea).Conn("bad", godip.Sea).Conn("cha", godip.Sea).Conn("yuz", godip.Sea).Flag(godip.Sea).
		// West Yangtze
		Prov("wes").Conn("bad", godip.Sea).Conn("xia", godip.Sea).Conn("sha", godip.Sea).Conn("qia", godip.Sea).Conn("ful", godip.Sea).Flag(godip.Sea).
		// Changsha
		Prov("cha").Conn("lin", godip.Land).Conn("yuz", godip.Coast...).Conn("nan", godip.Sea).Conn("bad", godip.Coast...).Conn("wul", godip.Land).Flag(godip.Coast...).
		// Yongchang
		Prov("yon").Conn("jio", godip.Land).Conn("zan", godip.Land).Conn("qia", godip.Land).Flag(godip.Land).
		// Jiaozhi
		Prov("jio").Conn("yon", godip.Land).Conn("sou", godip.Sea).Conn("pea", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// East Yangtze
		Prov("eas").Conn("jin", godip.Sea).Conn("cen", godip.Sea).Conn("gua", godip.Sea).Conn("luj", godip.Sea).Conn("yuz", godip.Sea).Flag(godip.Sea).
		// Nanhai
		Prov("nah").Conn("can", godip.Sea).Conn("sou", godip.Sea).Conn("cen", godip.Sea).Conn("kua", godip.Coast...).Conn("yuz", godip.Land).Conn("lin", godip.Land).Conn("hep", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Central Yangtze
		Prov("cen").Conn("che", godip.Sea).Conn("che", godip.Sea).Conn("sha", godip.Sea).Conn("sha", godip.Sea).Conn("sha", godip.Sea).Conn("xia", godip.Sea).Conn("gre", godip.Sea).Conn("run", godip.Sea).Conn("run", godip.Sea).Conn("gua", godip.Sea).Conn("gua", godip.Sea).Conn("gua", godip.Sea).Conn("gua", godip.Sea).Conn("gua", godip.Sea).Conn("eas", godip.Sea).Conn("jin", godip.Sea).Conn("kua", godip.Sea).Conn("nah", godip.Sea).Conn("sou", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Qianwei
		Prov("qia").Conn("yon", godip.Land).Conn("zan", godip.Land).Conn("ful", godip.Coast...).Conn("wes", godip.Sea).Conn("sha", godip.Coast...).Conn("che", godip.Land).Flag(godip.Coast...).SC(Wei).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"can": "Cangwu",
	"run": "Runan",
	"hep": "Hepu",
	"wul": "Wuling",
	"che": "Chengdu",
	"gre": "Great Space",
	"pea": "Pearl River",
	"yuz": "Yuzhang",
	"xia": "Xiangyang",
	"sou": "South China Sea",
	"kua": "Kuaiji",
	"luj": "Lujiang",
	"ful": "Fuling",
	"bad": "Badong",
	"zan": "Zangke",
	"lin": "Lingling",
	"gua": "Guangling",
	"jin": "Jianye",
	"sha": "Shangyong",
	"nan": "Nan",
	"wes": "West Yangtze",
	"cha": "Changsha",
	"yon": "Yongchang",
	"jio": "Jiaozhi",
	"eas": "East Yangtze",
	"nah": "Nanhai",
	"cen": "Central Yangtze",
	"qia": "Qianwei",
}
