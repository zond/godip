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
	Wu  godip.Nation = "Wu"
	Shu godip.Nation = "Shu"
)

var Nations = []godip.Nation{Wei, Wu, Shu}

var ThreeKingdomsVariant = common.Variant{
	Name:              "Three Kingdoms",
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
	SVGVersion:  "1",
	SVGUnits: 	 sengoku.SengokuVariant.SVGUnits,
	CreatedBy:   "Kuthador",
	Version:     "0.9",
	Description: "THIS IS A BETA MAP. IT MIGHT BE UPDATED AND CHANGED DURING YOUR GAME, WITHOUT WARNING. IT IS ONLY ACCESSIBLE OR VISIBLE FROM THE DIPLICITY BETA VERSION. The Three Kingdoms from 220 to 280 AD was the tripartite division of China among the dynastic states of Cao Wei, Shu Han, and Eastern Wu. The period is one of the bloodiest in Chinese history. The term 'Three Kingdoms' is something of a misnomer, since each state was eventually headed not by a king, but by an emperor who claimed suzerainty over all China. Fight for suzerainty with your two adversaries!",
	Rules:       "The upper, middle and lower Yangtze count as adjacent; fleets can move and convoy between them directly",
}

func ThreeKingdomsBlank(phase godip.Phase) *state.State {
	return state.New(ThreeKingdomsGraph(), phase, classical.BackupRule, nil, nil)
}

func ThreeKingdomsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(220, godip.Spring, godip.Movement)
	result = ThreeKingdomsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"don": godip.Unit{godip.Fleet, Wei},
		"qio": godip.Unit{godip.Army, Wei},
		"luo": godip.Unit{godip.Army, Wei},
		"tai": godip.Unit{godip.Army, Wei},
		"nah": godip.Unit{godip.Fleet, Wu},
		"lin": godip.Unit{godip.Army, Wu},
		"jiy": godip.Unit{godip.Army, Wu},
		"yuz": godip.Unit{godip.Army, Wu},
		"qin": godip.Unit{godip.Fleet, Shu},
		"che": godip.Unit{godip.Army, Shu},
		"han": godip.Unit{godip.Army, Shu},
		"zan": godip.Unit{godip.Army, Shu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"qio": Wei,
		"luo": Wei,
		"tai": Wei,
		"nah": Wu,
		"lin": Wu,
		"jiy": Wu,
		"yuz": Wu,
		"qin": Shu,
		"che": Shu,
		"han": Shu,
		"zan": Shu,
	})
	return
}

func ThreeKingdomsGraph() *graph.Graph {
	return graph.New().
		// Zhongshan
		Prov("zho").Conn("boh", godip.Coast...).Conn("tai", godip.Land).Conn("jic", godip.Coast...).Conn("yer", godip.Sea).Flag(godip.Coast...).
		// Cangwu
		Prov("can").Conn("nah", godip.Coast...).Conn("lin", godip.Land).Conn("wul", godip.Land).Conn("zan", godip.Coast...).Conn("hep", godip.Sea).Flag(godip.Coast...).
		// Taiyuan
		Prov("tai").Conn("boh", godip.Land).Conn("jic", godip.Land).Conn("zho", godip.Land).Flag(godip.Land).SC(Wei).
		// Runan
		Prov("run").Conn("nan", godip.Land).Conn("luj", godip.Land).Conn("gua", godip.Land).Conn("qio", godip.Land).Conn("luo", godip.Land).Flag(godip.Land).
		// Hepu
		Prov("hep").Conn("nah", godip.Sea).Conn("can", godip.Sea).Conn("zan", godip.Sea).Conn("pea", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Hanzhong
		Prov("han").Conn("jig", godip.Land).Conn("and", godip.Land).Conn("che", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).SC(Shu).
		// Wuling
		Prov("wul").Conn("ful", godip.Land).Conn("zan", godip.Land).Conn("can", godip.Land).Conn("lin", godip.Land).Conn("cha", godip.Land).Conn("bad", godip.Land).Flag(godip.Land).
		// Chengdu
		Prov("che").Conn("qin", godip.Land).Conn("sha", godip.Land).Conn("han", godip.Land).Conn("and", godip.Land).Flag(godip.Land).SC(Shu).
		// Guangling
		Prov("gua").Conn("don", godip.Coast...).Conn("lan", godip.Coast...).Conn("qio", godip.Land).Conn("run", godip.Land).Conn("luj", godip.Coast...).Conn("low", godip.Sea).Conn("eas", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Badong
		Prov("bad").Conn("mid", godip.Sea).Conn("xia", godip.Coast...).Conn("upp", godip.Sea).Conn("ful", godip.Coast...).Conn("wul", godip.Land).Conn("cha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Xiangyang
		Prov("xia").Conn("mid", godip.Sea).Conn("nan", godip.Coast...).Conn("luo", godip.Land).Conn("sha", godip.Coast...).Conn("upp", godip.Sea).Conn("bad", godip.Coast...).Flag(godip.Coast...).
		// South China Sea
		Prov("sou").Conn("eas", godip.Sea).Conn("nah", godip.Sea).Conn("hep", godip.Sea).Conn("pea", godip.Sea).Conn("jio", godip.Sea).Flag(godip.Sea).
		// Kuaiji
		Prov("kua").Conn("jiy", godip.Coast...).Conn("yuz", godip.Land).Conn("nah", godip.Coast...).Conn("eas", godip.Sea).Flag(godip.Coast...).
		// East China Sea
		Prov("eas").Conn("yes", godip.Sea).Conn("gua", godip.Sea).Conn("low", godip.Sea).Conn("jiy", godip.Sea).Conn("kua", godip.Sea).Conn("nah", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Anding
		Prov("and").Conn("che", godip.Land).Conn("han", godip.Land).Conn("jig", godip.Land).Conn("jic", godip.Land).Flag(godip.Land).
		// Middle Yangtze
		Prov("mid").Conn("low", godip.Sea).Conn("upp", godip.Sea).Conn("bad", godip.Sea).Conn("cha", godip.Sea).Conn("yuz", godip.Sea).Conn("luj", godip.Sea).Conn("nan", godip.Sea).Conn("xia", godip.Sea).Flag(godip.Sea).
		// Lujiang
		Prov("luj").Conn("nan", godip.Coast...).Conn("mid", godip.Sea).Conn("yuz", godip.Coast...).Conn("low", godip.Sea).Conn("gua", godip.Coast...).Conn("run", godip.Land).Flag(godip.Coast...).
		// Langya
		Prov("lan").Conn("yer", godip.Sea).Conn("qio", godip.Coast...).Conn("gua", godip.Coast...).Conn("don", godip.Coast...).Conn("yes", godip.Sea).Flag(godip.Coast...).
		// Fuling
		Prov("ful").Conn("wul", godip.Land).Conn("bad", godip.Coast...).Conn("upp", godip.Sea).Conn("qin", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// Jingzhao
		Prov("jig").Conn("sha", godip.Land).Conn("luo", godip.Coast...).Conn("yer", godip.Sea).Conn("jic", godip.Coast...).Conn("and", godip.Land).Conn("han", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Yuzhang
		Prov("yuz").Conn("nah", godip.Land).Conn("kua", godip.Land).Conn("jiy", godip.Coast...).Conn("low", godip.Sea).Conn("luj", godip.Coast...).Conn("mid", godip.Sea).Conn("cha", godip.Coast...).Conn("lin", godip.Land).Flag(godip.Coast...).SC(Wu).
		// Zangke
		Prov("zan").Conn("qin", godip.Land).Conn("yon", godip.Land).Conn("jio", godip.Land).Conn("pea", godip.Coast...).Conn("hep", godip.Sea).Conn("can", godip.Coast...).Conn("wul", godip.Land).Conn("ful", godip.Land).Flag(godip.Coast...).SC(Shu).
		// Upper Yangtze
		Prov("upp").Conn("mid", godip.Sea).Conn("sha", godip.Sea).Conn("qin", godip.Sea).Conn("ful", godip.Sea).Conn("bad", godip.Sea).Conn("xia", godip.Sea).Flag(godip.Sea).
		// Jincheng
		Prov("jic").Conn("yer", godip.Sea).Conn("zho", godip.Coast...).Conn("tai", godip.Land).Conn("and", godip.Land).Conn("jig", godip.Coast...).Flag(godip.Coast...).
		// Lingling
		Prov("lin").Conn("nah", godip.Land).Conn("yuz", godip.Land).Conn("cha", godip.Land).Conn("wul", godip.Land).Conn("can", godip.Land).Flag(godip.Land).SC(Wu).
		// Jianye
		Prov("jiy").Conn("kua", godip.Coast...).Conn("eas", godip.Sea).Conn("low", godip.Sea).Conn("yuz", godip.Coast...).Flag(godip.Coast...).SC(Wu).
		// Nanhai
		Prov("nah").Conn("can", godip.Coast...).Conn("hep", godip.Sea).Conn("sou", godip.Sea).Conn("eas", godip.Sea).Conn("kua", godip.Coast...).Conn("yuz", godip.Land).Conn("lin", godip.Land).Flag(godip.Coast...).SC(Wu).
		// Bohai
		Prov("boh").Conn("zho", godip.Coast...).Conn("yer", godip.Sea).Conn("yes", godip.Sea).Conn("tai", godip.Land).Flag(godip.Coast...).
		// Shangyong
		Prov("sha").Conn("jig", godip.Land).Conn("han", godip.Land).Conn("che", godip.Land).Conn("qin", godip.Coast...).Conn("upp", godip.Sea).Conn("xia", godip.Coast...).Conn("luo", godip.Land).Flag(godip.Coast...).
		// Qiao
		Prov("qio").Conn("lan", godip.Coast...).Conn("yer", godip.Sea).Conn("luo", godip.Coast...).Conn("run", godip.Land).Conn("gua", godip.Land).Flag(godip.Coast...).SC(Wei).
		// Nan
		Prov("nan").Conn("run", godip.Land).Conn("luo", godip.Land).Conn("xia", godip.Coast...).Conn("mid", godip.Sea).Conn("luj", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Luoyang
		Prov("luo").Conn("nan", godip.Land).Conn("run", godip.Land).Conn("qio", godip.Coast...).Conn("yer", godip.Sea).Conn("jig", godip.Coast...).Conn("sha", godip.Land).Conn("xia", godip.Land).Flag(godip.Coast...).SC(Wei).
		// Changsha
		Prov("cha").Conn("bad", godip.Coast...).Conn("wul", godip.Land).Conn("lin", godip.Land).Conn("yuz", godip.Coast...).Conn("mid", godip.Sea).Flag(godip.Coast...).
		// Yongchang
		Prov("yon").Conn("jio", godip.Land).Conn("zan", godip.Land).Conn("qin", godip.Land).Flag(godip.Land).
		// Jiaozhi
		Prov("jio").Conn("yon", godip.Land).Conn("sou", godip.Sea).Conn("pea", godip.Coast...).Conn("zan", godip.Land).Flag(godip.Coast...).
		// Donglai
		Prov("don").Conn("gua", godip.Coast...).Conn("yes", godip.Sea).Conn("lan", godip.Coast...).Flag(godip.Coast...).
		// Hedong
		Prov("hed").Flag(godip.Land).
		// Yellow river
		Prov("yer").Conn("jic", godip.Sea).Conn("jig", godip.Sea).Conn("luo", godip.Sea).Conn("qio", godip.Sea).Conn("lan", godip.Sea).Conn("yes", godip.Sea).Conn("boh", godip.Sea).Conn("zho", godip.Sea).Flag(godip.Sea).
		// Yellow Sea
		Prov("yes").Conn("boh", godip.Sea).Conn("yer", godip.Sea).Conn("lan", godip.Sea).Conn("don", godip.Sea).Conn("gua", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Lower Yangtze
		Prov("low").Conn("mid", godip.Sea).Conn("gua", godip.Sea).Conn("luj", godip.Sea).Conn("yuz", godip.Sea).Conn("jiy", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Pearl river
		Prov("pea").Conn("sou", godip.Sea).Conn("hep", godip.Sea).Conn("zan", godip.Coast...).Conn("jio", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Qianwei
		Prov("qin").Conn("zan", godip.Land).Conn("ful", godip.Coast...).Conn("upp", godip.Sea).Conn("sha", godip.Coast...).Conn("che", godip.Land).Conn("yon", godip.Land).Flag(godip.Coast...).SC(Shu).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"zho": "Zhongshan",
	"can": "Cangwu",
	"tai": "Taiyuan",
	"run": "Runan",
	"hep": "Hepu",
	"han": "Hanzhong",
	"wul": "Wuling",
	"che": "Chengdu",
	"gua": "Guangling",
	"bad": "Badong",
	"xia": "Xiangyang",
	"sou": "South China Sea",
	"kua": "Kuaiji",
	"eas": "East China Sea",
	"and": "Anding",
	"mid": "Middle Yangtze",
	"luj": "Lujiang",
	"lan": "Langya",
	"ful": "Fuling",
	"jig": "Jingzhao",
	"yuz": "Yuzhang",
	"zan": "Zangke",
	"upp": "Upper Yangtze",
	"jic": "Jincheng",
	"lin": "Lingling",
	"jiy": "Jianye",
	"nah": "Nanhai",
	"boh": "Bohai",
	"sha": "Shangyong",
	"qio": "Qiao",
	"nan": "Nan",
	"luo": "Luoyang",
	"cha": "Changsha",
	"yon": "Yongchang",
	"jio": "Jiaozhi",
	"don": "Donglai",
	"hed": "Hedong",
	"yer": "Yellow river",
	"yes": "Yellow Sea",
	"low": "Lower Yangtze",
	"pea": "Pearl river",
	"qin": "Qianwei",
}
