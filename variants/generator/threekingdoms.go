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
	SoloWinner:        common.SCCountWinner(9),
	SoloSCCount:       func(*state.State) int { return 9 },
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
		"qin": godip.Unit{godip.Army, Wei},
		"zan": godip.Unit{godip.Army, Shu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"qin": Wei,
		"zan": Shu,
	})
	return
}

func ThreeKingdomsGraph() *graph.Graph {
	return graph.New().
		// Zhongshan
		Prov("zho").Conn("yer", godip.Sea).Conn("boh", godip.Coast...).Conn("tay", godip.Land).Conn("hed", godip.Coast...).Flag(godip.Coast...).
		// Cangwu
		Prov("can").Conn("nah", godip.Sea).Conn("hep", godip.Sea).Conn("zan", godip.Sea).Conn("pea", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Taiyuan
		Prov("tay").Conn("and", godip.Land).Conn("hed", godip.Land).Conn("zho", godip.Land).Conn("boh", godip.Land).Conn("cen", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Runan
		Prov("run").Conn("luo", godip.Land).Conn("luj", godip.Land).Conn("gua", godip.Land).Conn("qio", godip.Land).Flag(godip.Land).
		// Guangling
		Prov("gua").Conn("don", godip.Coast...).Conn("qio", godip.Land).Conn("run", godip.Land).Conn("luj", godip.Coast...).Conn("eay", godip.Sea).Conn("ecs", godip.Sea).Conn("yes", godip.Sea).Conn("lan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Hepu
		Prov("hep").Conn("lin", godip.Land).Conn("wul", godip.Land).Conn("zan", godip.Coast...).Conn("can", godip.Sea).Conn("nah", godip.Coast...).Flag(godip.Coast...).
		// Hanzhong
		Prov("han").Conn("sha", godip.Land).Conn("jig", godip.Land).Conn("jic", godip.Land).Conn("che", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wuling
		Prov("wul").Conn("bad", godip.Land).Conn("ful", godip.Land).Conn("zan", godip.Land).Conn("hep", godip.Land).Conn("lin", godip.Land).Conn("cha", godip.Land).Flag(godip.Land).
		// Hedong
		Prov("hed").Conn("tay", godip.Land).Conn("and", godip.Land).Conn("jig", godip.Coast...).Conn("yer", godip.Sea).Conn("zho", godip.Coast...).Flag(godip.Coast...).
		// Pearl River
		Prov("pea").Conn("jio", godip.Coast...).Conn("sou", godip.Sea).Conn("can", godip.Sea).Conn("zan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Yuzhang
		Prov("yuz").Conn("lin", godip.Land).Conn("nah", godip.Land).Conn("kua", godip.Land).Conn("jiy", godip.Coast...).Conn("eay", godip.Sea).Conn("luj", godip.Coast...).Conn("nan", godip.Sea).Conn("cha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Xiangyang
		Prov("xia").Conn("luo", godip.Land).Conn("sha", godip.Coast...).Conn("wes", godip.Sea).Conn("bad", godip.Coast...).Conn("nan", godip.Sea).Flag(godip.Coast...).
		// Kuaiji
		Prov("kua").Conn("jiy", godip.Coast...).Conn("yuz", godip.Land).Conn("nah", godip.Coast...).Conn("ecs", godip.Sea).Flag(godip.Coast...).
		// Anding
		Prov("and").Conn("tay", godip.Land).Conn("cen", godip.Land).Conn("jic", godip.Land).Conn("jig", godip.Land).Conn("hed", godip.Land).Flag(godip.Land).
		// Jincheng
		Prov("jic").Conn("and", godip.Land).Conn("cen", godip.Land).Conn("che", godip.Land).Conn("han", godip.Land).Conn("jig", godip.Land).Flag(godip.Land).
		// Lujiang
		Prov("luj").Conn("nan", godip.Sea).Conn("yuz", godip.Coast...).Conn("eay", godip.Sea).Conn("gua", godip.Coast...).Conn("run", godip.Land).Flag(godip.Coast...).
		// Langya
		Prov("lan").Conn("don", godip.Coast...).Conn("gua", godip.Coast...).Conn("yes", godip.Sea).Flag(godip.Coast...).
		// Fuling
		Prov("ful").Conn("wul", godip.Land).Conn("bad", godip.Coast...).Conn("wes", godip.Sea).Conn("qin", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// Jingzhao
		Prov("jig").Conn("and", godip.Land).Conn("jic", godip.Land).Conn("han", godip.Land).Conn("sha", godip.Land).Conn("luo", godip.Coast...).Conn("yer", godip.Sea).Conn("hed", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Badong
		Prov("bad").Conn("wes", godip.Sea).Conn("ful", godip.Coast...).Conn("wul", godip.Land).Conn("cha", godip.Coast...).Conn("nan", godip.Sea).Conn("xia", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Zangke
		Prov("zan").Conn("ful", godip.Land).Conn("qin", godip.Land).Conn("yon", godip.Land).Conn("jio", godip.Land).Conn("pea", godip.Coast...).Conn("can", godip.Sea).Conn("hep", godip.Coast...).Conn("wul", godip.Land).Flag(godip.Coast...).SC(Shu).
		// Lingling
		Prov("lin").Conn("hep", godip.Land).Conn("nah", godip.Land).Conn("yuz", godip.Land).Conn("cha", godip.Land).Conn("wul", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Jianye
		Prov("jiy").Conn("kua", godip.Coast...).Conn("ecs", godip.Sea).Conn("eay", godip.Sea).Conn("yuz", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Nanhai
		Prov("nah").Conn("can", godip.Sea).Conn("sou", godip.Sea).Conn("ecs", godip.Sea).Conn("kua", godip.Coast...).Conn("yuz", godip.Land).Conn("lin", godip.Land).Conn("hep", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Bohai
		Prov("boh").Conn("yer", godip.Sea).Conn("yes", godip.Sea).Conn("cen", godip.Coast...).Conn("tay", godip.Land).Conn("zho", godip.Coast...).Flag(godip.Coast...).
		// Yellow River
		Prov("yer").Conn("yes", godip.Sea).Conn("boh", godip.Sea).Conn("zho", godip.Sea).Conn("hed", godip.Sea).Conn("jig", godip.Sea).Conn("luo", godip.Sea).Conn("qio", godip.Sea).Conn("don", godip.Sea).Flag(godip.Sea).
		// Shangyong
		Prov("sha").Conn("han", godip.Land).Conn("che", godip.Land).Conn("qin", godip.Coast...).Conn("wes", godip.Sea).Conn("xia", godip.Coast...).Conn("luo", godip.Land).Conn("jig", godip.Land).Flag(godip.Coast...).
		// Qiao
		Prov("qio").Conn("luo", godip.Coast...).Conn("run", godip.Land).Conn("gua", godip.Land).Conn("don", godip.Coast...).Conn("yer", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Nan
		Prov("nan").Conn("luj", godip.Sea).Conn("xia", godip.Sea).Conn("bad", godip.Sea).Conn("cha", godip.Sea).Conn("yuz", godip.Sea).Flag(godip.Sea).
		// West Yangtze
		Prov("wes").Conn("bad", godip.Sea).Conn("xia", godip.Sea).Conn("sha", godip.Sea).Conn("qin", godip.Sea).Conn("ful", godip.Sea).Flag(godip.Sea).
		// Luoyang
		Prov("luo").Conn("xia", godip.Land).Conn("run", godip.Land).Conn("qio", godip.Coast...).Conn("yer", godip.Sea).Conn("jig", godip.Coast...).Conn("sha", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Changsha
		Prov("cha").Conn("lin", godip.Land).Conn("yuz", godip.Coast...).Conn("nan", godip.Sea).Conn("bad", godip.Coast...).Conn("wul", godip.Land).Flag(godip.Coast...).
		// Yongchang
		Prov("yon").Conn("jio", godip.Land).Conn("zan", godip.Land).Conn("qin", godip.Land).Conn("cen", godip.Land).Flag(godip.Land).
		// SOUTH CHINA SEA
		Prov("sou").Conn("ecs", godip.Sea).Conn("ecs", godip.Sea).Conn("nah", godip.Sea).Conn("can", godip.Sea).Conn("pea", godip.Sea).Conn("jio", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Sea).
		// EAST CHINA SEA
		Prov("ecs").Conn("cen", godip.Sea).Conn("yes", godip.Sea).Conn("gua", godip.Sea).Conn("eay", godip.Sea).Conn("jiy", godip.Sea).Conn("kua", godip.Sea).Conn("nah", godip.Sea).Conn("sou", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Jiaozhi
		Prov("jio").Conn("yon", godip.Land).Conn("cen", godip.Coast...).Conn("sou", godip.Sea).Conn("pea", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// East Yangtze
		Prov("eay").Conn("jiy", godip.Sea).Conn("ecs", godip.Sea).Conn("gua", godip.Sea).Conn("luj", godip.Sea).Conn("yuz", godip.Sea).Flag(godip.Sea).
		// Donglai
		Prov("don").Conn("gua", godip.Coast...).Conn("lan", godip.Coast...).Conn("yes", godip.Sea).Conn("yer", godip.Sea).Conn("qio", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Chengdu
		Prov("che").Conn("han", godip.Land).Conn("jic", godip.Land).Conn("cen", godip.Land).Conn("qin", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Central Yangtze
		Prov("cen").Conn("sou", godip.Sea).Conn("jio", godip.Coast...).Conn("yon", godip.Land).Conn("qin", godip.Land).Conn("che", godip.Land).Conn("jic", godip.Land).Conn("and", godip.Land).Conn("tay", godip.Land).Conn("boh", godip.Coast...).Conn("yes", godip.Sea).Conn("ecs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Qianwei
		Prov("qin").Conn("yon", godip.Land).Conn("zan", godip.Land).Conn("ful", godip.Coast...).Conn("wes", godip.Sea).Conn("sha", godip.Coast...).Conn("che", godip.Land).Conn("cen", godip.Land).Flag(godip.Coast...).SC(Wei).
		// YELLOW SEA
		Prov("yes").Conn("yer", godip.Sea).Conn("don", godip.Sea).Conn("lan", godip.Sea).Conn("gua", godip.Sea).Conn("ecs", godip.Sea).Conn("cen", godip.Sea).Conn("boh", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"zho": "Zhongshan",
	"can": "Cangwu",
	"tay": "Taiyuan",
	"run": "Runan",
	"gua": "Guangling",
	"hep": "Hepu",
	"han": "Hanzhong",
	"wul": "Wuling",
	"hed": "Hedong",
	"pea": "Pearl River",
	"yuz": "Yuzhang",
	"xia": "Xiangyang",
	"kua": "Kuaiji",
	"and": "Anding",
	"jic": "Jincheng",
	"luj": "Lujiang",
	"lan": "Langya",
	"ful": "Fuling",
	"jig": "Jingzhao",
	"bad": "Badong",
	"zan": "Zangke",
	"lin": "Lingling",
	"jiy": "Jianye",
	"nah": "Nanhai",
	"boh": "Bohai",
	"yer": "Yellow River",
	"sha": "Shangyong",
	"qio": "Qiao",
	"nan": "Nan",
	"wes": "West Yangtze",
	"luo": "Luoyang",
	"cha": "Changsha",
	"yon": "Yongchang",
	"sou": "SOUTH CHINA SEA",
	"ecs": "EAST CHINA SEA",
	"jio": "Jiaozhi",
	"eay": "East Yangtze",
	"don": "Donglai",
	"che": "Chengdu",
	"cen": "Central Yangtze",
	"qin": "Qianwei",
	"yes": "YELLOW SEA",
}
