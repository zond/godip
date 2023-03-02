package threekingdoms

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Wu  godip.Nation = "Wu"
	Shu godip.Nation = "Shu"
	Wei godip.Nation = "Wei"
)

var Nations = []godip.Nation{Wu, Shu, Wei}
var SVGFlags = map[godip.Nation]func() ([]byte, error){
	Wu: func() ([]byte, error) {
		return Asset("svg/wu.svg")
	},
	Shu: func() ([]byte, error) {
		return Asset("svg/shu.svg")
	},
	Wei: func() ([]byte, error) {
		return Asset("svg/wei.svg")
	},
}

var ThreeKingdomsVariant = common.Variant{
	Name:              "Three Kingdoms",
	NationColors: map[godip.Nation]string{
		Wu:       "#7F5E13",
		Shu:      "#126176",
		Wei:      "#7D1C12",
	},
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
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	},
	SVGFlags:    SVGFlags,
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
		"nah": godip.Unit{godip.Fleet, Wu},
		"lin": godip.Unit{godip.Army, Wu},
		"yuz": godip.Unit{godip.Army, Wu},
		"jiy": godip.Unit{godip.Army, Wu},
		"qin": godip.Unit{godip.Fleet, Shu},
		"che": godip.Unit{godip.Army, Shu},
		"han": godip.Unit{godip.Army, Shu},
		"zan": godip.Unit{godip.Army, Shu},
		"don": godip.Unit{godip.Fleet, Wei},
		"luo": godip.Unit{godip.Army, Wei},
		"qio": godip.Unit{godip.Army, Wei},
		"tay": godip.Unit{godip.Army, Wei},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"nah": Wu,
		"lin": Wu,
		"yuz": Wu,
		"jiy": Wu,
		"qin": Shu,
		"che": Shu,
		"han": Shu,
		"zan": Shu,
		"don": Wei,
		"luo": Wei,
		"qio": Wei,
		"tay": Wei,
	})
	return
}

func ThreeKingdomsGraph() *graph.Graph {
	return graph.New().
		// Zhongshan
		Prov("zho").Conn("tay", godip.Land).Conn("hed", godip.Coast...).Conn("yer", godip.Sea).Conn("boh", godip.Coast...).Flag(godip.Coast...).
		// Cangwu
		Prov("can").Conn("hep", godip.Sea).Conn("zan", godip.Sea).Conn("pea", godip.Sea).Conn("sou", godip.Sea).Conn("nah", godip.Sea).Flag(godip.Sea).
		// Taiyuan
		Prov("tay").Conn("zho", godip.Land).Conn("boh", godip.Land).Conn("and", godip.Land).Conn("hed", godip.Land).Flag(godip.Land).SC(Wei).
		// Runan
		Prov("run").Conn("luj", godip.Land).Conn("gua", godip.Land).Conn("qio", godip.Land).Conn("luo", godip.Land).Conn("nan", godip.Land).Flag(godip.Land).
		// Guangling
		Prov("gua").Conn("don", godip.Land).Conn("qio", godip.Land).Conn("run", godip.Land).Conn("luj", godip.Coast...).Conn("eay", godip.Sea).Conn("ecs", godip.Sea).Conn("yes", godip.Sea).Conn("lan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Hepu
		Prov("hep").Conn("can", godip.Sea).Conn("nah", godip.Coast...).Conn("lin", godip.Land).Conn("wul", godip.Land).Conn("zan", godip.Coast...).Flag(godip.Coast...).
		// Hanzhong
		Prov("han").Conn("jic", godip.Land).Conn("che", godip.Land).Conn("sha", godip.Land).Conn("jig", godip.Land).Flag(godip.Land).SC(Shu).
		// Wuling
		Prov("wul").Conn("hep", godip.Land).Conn("lin", godip.Land).Conn("cha", godip.Land).Conn("bad", godip.Land).Conn("ful", godip.Land).Conn("zan", godip.Land).Flag(godip.Land).
		// Hedong
		Prov("hed").Conn("jig", godip.Coast...).Conn("yer", godip.Sea).Conn("zho", godip.Coast...).Conn("tay", godip.Land).Conn("and", godip.Land).Flag(godip.Coast...).
		// Pearl River
		Prov("pea").Conn("can", godip.Sea).Conn("zan", godip.Coast...).Conn("jio", godip.Coast...).Conn("sou", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Yuzhang
		Prov("yuz").Conn("kua", godip.Land).Conn("jiy", godip.Coast...).Conn("eay", godip.Sea).Conn("luj", godip.Coast...).Conn("cen", godip.Sea).Conn("cha", godip.Coast...).Conn("lin", godip.Land).Conn("nah", godip.Land).Flag(godip.Coast...).SC(Wu).
		// Xiangyang
		Prov("xia").Conn("luo", godip.Land).Conn("sha", godip.Coast...).Conn("wes", godip.Sea).Conn("bad", godip.Coast...).Conn("cen", godip.Sea).Conn("nan", godip.Coast...).Flag(godip.Coast...).
		// Kuaiji
		Prov("kua").Conn("ecs", godip.Sea).Conn("jiy", godip.Coast...).Conn("yuz", godip.Land).Conn("nah", godip.Coast...).Flag(godip.Coast...).
		// Anding
		Prov("and").Conn("jic", godip.Land).Conn("jig", godip.Land).Conn("hed", godip.Land).Conn("tay", godip.Land).Flag(godip.Land).
		// Jincheng
		Prov("jic").Conn("han", godip.Land).Conn("jig", godip.Land).Conn("and", godip.Land).Conn("che", godip.Land).Flag(godip.Land).
		// Lujiang
		Prov("luj").Conn("run", godip.Land).Conn("nan", godip.Coast...).Conn("cen", godip.Sea).Conn("yuz", godip.Coast...).Conn("eay", godip.Sea).Conn("gua", godip.Coast...).Flag(godip.Coast...).
		// Langya
		Prov("lan").Conn("yes", godip.Sea).Conn("don", godip.Coast...).Conn("gua", godip.Coast...).Flag(godip.Coast...).
		// Fuling
		Prov("ful").Conn("bad", godip.Coast...).Conn("wes", godip.Sea).Conn("qin", godip.Coast...).Conn("zan", godip.Land).Conn("wul", godip.Land).Flag(godip.Coast...).
		// Jingzhao
		Prov("jig").Conn("jic", godip.Land).Conn("han", godip.Land).Conn("sha", godip.Land).Conn("luo", godip.Coast...).Conn("yer", godip.Sea).Conn("hed", godip.Coast...).Conn("and", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Badong
		Prov("bad").Conn("ful", godip.Coast...).Conn("wul", godip.Land).Conn("cha", godip.Coast...).Conn("cen", godip.Sea).Conn("xia", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Zangke
		Prov("zan").Conn("qin", godip.Land).Conn("yon", godip.Land).Conn("jio", godip.Land).Conn("pea", godip.Coast...).Conn("can", godip.Sea).Conn("hep", godip.Coast...).Conn("wul", godip.Land).Conn("ful", godip.Land).Flag(godip.Coast...).SC(Shu).
		// Lingling
		Prov("lin").Conn("nah", godip.Land).Conn("yuz", godip.Land).Conn("cha", godip.Land).Conn("wul", godip.Land).Conn("hep", godip.Land).Flag(godip.Land).SC(Wu).
		// Jianye
		Prov("jiy").Conn("yuz", godip.Coast...).Conn("kua", godip.Coast...).Conn("ecs", godip.Sea).Conn("eay", godip.Sea).Flag(godip.Coast...).SC(Wu).
		// Nanhai
		Prov("nah").Conn("lin", godip.Land).Conn("hep", godip.Coast...).Conn("can", godip.Sea).Conn("sou", godip.Sea).Conn("ecs", godip.Sea).Conn("kua", godip.Coast...).Conn("yuz", godip.Land).Flag(godip.Coast...).SC(Wu).
		// Bohai
		Prov("boh").Conn("tay", godip.Land).Conn("zho", godip.Coast...).Conn("yer", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Coast...).
		// Yellow River
		Prov("yer").Conn("yes", godip.Sea).Conn("boh", godip.Sea).Conn("zho", godip.Sea).Conn("hed", godip.Sea).Conn("jig", godip.Sea).Conn("luo", godip.Sea).Conn("qio", godip.Sea).Conn("don", godip.Sea).Flag(godip.Sea).
		// Shangyong
		Prov("sha").Conn("han", godip.Land).Conn("che", godip.Land).Conn("qin", godip.Coast...).Conn("wes", godip.Sea).Conn("xia", godip.Coast...).Conn("luo", godip.Land).Conn("jig", godip.Land).Flag(godip.Coast...).
		// Qiao
		Prov("qio").Conn("don", godip.Coast...).Conn("yer", godip.Sea).Conn("luo", godip.Coast...).Conn("run", godip.Land).Conn("gua", godip.Land).Flag(godip.Coast...).SC(Wei).
		// Nan
		Prov("nan").Conn("cen", godip.Sea).Conn("luj", godip.Coast...).Conn("run", godip.Land).Conn("luo", godip.Land).Conn("xia", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// West Yangtze
		Prov("wes").Conn("cen", godip.Sea).Conn("qin", godip.Sea).Conn("ful", godip.Sea).Conn("bad", godip.Sea).Conn("xia", godip.Sea).Conn("sha", godip.Sea).Flag(godip.Sea).
		// Luoyang
		Prov("luo").Conn("xia", godip.Land).Conn("nan", godip.Land).Conn("run", godip.Land).Conn("qio", godip.Coast...).Conn("yer", godip.Sea).Conn("jig", godip.Coast...).Conn("sha", godip.Land).Flag(godip.Coast...).SC(Wei).
		// Changsha
		Prov("cha").Conn("bad", godip.Coast...).Conn("wul", godip.Land).Conn("lin", godip.Land).Conn("yuz", godip.Coast...).Conn("cen", godip.Sea).Flag(godip.Coast...).
		// Yongchang
		Prov("yon").Conn("qin", godip.Land).Conn("jio", godip.Land).Conn("zan", godip.Land).Flag(godip.Land).
		// SOUTH CHINA SEA
		Prov("sou").Conn("ecs", godip.Sea).Conn("ecs", godip.Sea).Conn("nah", godip.Sea).Conn("can", godip.Sea).Conn("pea", godip.Sea).Conn("jio", godip.Sea).Flag(godip.Sea).
		// EAST CHINA SEA
		Prov("ecs").Conn("yes", godip.Sea).Conn("gua", godip.Sea).Conn("eay", godip.Sea).Conn("jiy", godip.Sea).Conn("kua", godip.Sea).Conn("nah", godip.Sea).Conn("sou", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Jiaozhi
		Prov("jio").Conn("zan", godip.Land).Conn("yon", godip.Land).Conn("sou", godip.Sea).Conn("pea", godip.Coast...).Flag(godip.Coast...).
		// East Yangtze
		Prov("eay").Conn("cen", godip.Sea).Conn("gua", godip.Sea).Conn("luj", godip.Sea).Conn("yuz", godip.Sea).Conn("jiy", godip.Sea).Conn("ecs", godip.Sea).Flag(godip.Sea).
		// Donglai
		Prov("don").Conn("gua", godip.Land).Conn("lan", godip.Coast...).Conn("yes", godip.Sea).Conn("yer", godip.Sea).Conn("qio", godip.Coast...).Flag(godip.Coast...).SC(Wei).
		// Chengdu
		Prov("che").Conn("han", godip.Land).Conn("jic", godip.Land).Conn("qin", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).SC(Shu).
		// Central Yangtze
		Prov("cen").Conn("eay", godip.Sea).Conn("wes", godip.Sea).Conn("nan", godip.Sea).Conn("xia", godip.Sea).Conn("bad", godip.Sea).Conn("cha", godip.Sea).Conn("yuz", godip.Sea).Conn("luj", godip.Sea).Flag(godip.Sea).
		// Qianwei
		Prov("qin").Conn("zan", godip.Land).Conn("ful", godip.Coast...).Conn("wes", godip.Sea).Conn("sha", godip.Coast...).Conn("che", godip.Land).Conn("yon", godip.Land).Flag(godip.Coast...).SC(Shu).
		// YELLOW SEA
		Prov("yes").Conn("lan", godip.Sea).Conn("gua", godip.Sea).Conn("ecs", godip.Sea).Conn("boh", godip.Sea).Conn("yer", godip.Sea).Conn("don", godip.Sea).Flag(godip.Sea).
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
	"sou": "South China Sea",
	"ecs": "East China Sea",
	"jio": "Jiaozhi",
	"eay": "East Yangtze",
	"don": "Donglai",
	"che": "Chengdu",
	"cen": "Central Yangtze",
	"qin": "Qianwei",
	"yes": "Yellow Sea",
}
