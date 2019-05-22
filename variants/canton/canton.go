package canton

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Turkey  godip.Nation = "Turkey"
	Britain godip.Nation = "Britain"
	China   godip.Nation = "China"
	Holland godip.Nation = "Holland"
	Japan   godip.Nation = "Japan"
	Russia  godip.Nation = "Russia"
	France  godip.Nation = "France"
)

var Nations = []godip.Nation{Turkey, Britain, China, Holland, Japan, Russia, France}

var CantonVariant = common.Variant{
	Name:       "Canton",
	Graph:      func() godip.Graph { return CantonGraph() },
	Start:      CantonStart,
	Blank:      CantonBlank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/cantonmap.svg")
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
	CreatedBy: "Paul Webb",
	Version: "3",
	Description: "Canton is a seven-player Diplomacy variant set in Asia at the beginning of the 20th century.  The \"Great Powers\" are Britain, China, France, Holland, Japan, Russia, and Turkey.  Return to this era and determine the fate of Asia.",
	Rules: "Canton follows the same general principles of Classic Diplomacy. Units can only be built on home (starting) home supply centers. Constantinople (Con) and Egypt (Egy) work as canals (same as Kiel in Classic). Four provinces have dual coasts: Damascus (South/West), Bulgaria (South/East), Siam (West/East) and Canton (South/East). Provinces without names are impassable.",
}

func CantonBlank(phase godip.Phase) *state.State {
	return state.New(CantonGraph(), phase, classical.BackupRule, nil)
}

func CantonStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1901, godip.Spring, godip.Movement)
	result = CantonBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"con": godip.Unit{godip.Fleet, Turkey},
		"dam": godip.Unit{godip.Army, Turkey},
		"bag": godip.Unit{godip.Army, Turkey},
		"mad": godip.Unit{godip.Fleet, Britain},
		"cal": godip.Unit{godip.Army, Britain},
		"del": godip.Unit{godip.Army, Britain},
		"pek": godip.Unit{godip.Fleet, China},
		"sha": godip.Unit{godip.Army, China},
		"chu": godip.Unit{godip.Army, China},
		"tib": godip.Unit{godip.Army, China},
		"sum": godip.Unit{godip.Fleet, Holland},
		"jav": godip.Unit{godip.Fleet, Holland},
		"bor": godip.Unit{godip.Fleet, Holland},
		"tok": godip.Unit{godip.Fleet, Japan},
		"sas": godip.Unit{godip.Fleet, Japan},
		"kyo": godip.Unit{godip.Army, Japan},
		"sev": godip.Unit{godip.Fleet, Russia},
		"kha": godip.Unit{godip.Fleet, Russia},
		"mos": godip.Unit{godip.Army, Russia},
		"irk": godip.Unit{godip.Army, Russia},
		"sai": godip.Unit{godip.Fleet, France},
		"han": godip.Unit{godip.Army, France},
		"hue": godip.Unit{godip.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"con": Turkey,
		"dam": Turkey,
		"bag": Turkey,
		"mad": Britain,
		"cal": Britain,
		"del": Britain,
		"pek": China,
		"sha": China,
		"chu": China,
		"tib": China,
		"sum": Holland,
		"jav": Holland,
		"bor": Holland,
		"tok": Japan,
		"sas": Japan,
		"kyo": Japan,
		"sev": Russia,
		"kha": Russia,
		"mos": Russia,
		"irk": Russia,
		"sai": France,
		"han": France,
		"hue": France,
	})
	return
}

func CantonGraph() *graph.Graph {
	return graph.New().
		// Bombay
		Prov("bom").Conn("wes", godip.Sea).Conn("mad", godip.Coast...).Conn("del", godip.Land).Conn("bal", godip.Coast...).Conn("ars", godip.Sea).Flag(godip.Coast...).
		// Sevastopol
		Prov("sev").Conn("rum", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Conn("kaz", godip.Land).Conn("mos", godip.Land).Flag(godip.Coast...).SC(Russia).
		// Amur
		Prov("amu").Conn("irk", godip.Land).Conn("mon", godip.Land).Conn("man", godip.Land).Conn("mar", godip.Land).Conn("kha", godip.Land).Flag(godip.Land).
		// Kashmir
		Prov("kas").Conn("bal", godip.Land).Conn("del", godip.Land).Conn("tib", godip.Land).Conn("sin", godip.Land).Conn("afg", godip.Land).Flag(godip.Land).
		// Hong Kong
		Prov("hon").Conn("scs", godip.Sea).Conn("can", godip.Coast...).Conn("can", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Siam
		Prov("sia").Conn("mal", godip.Coast...).Conn("gos", godip.Sea).Conn("cam", godip.Coast...).Conn("lao", godip.Land).Conn("bur", godip.Coast...).Conn("gom", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Afghanistan
		Prov("afg").Conn("tur", godip.Land).Conn("pes", godip.Land).Conn("bal", godip.Land).Conn("kas", godip.Land).Conn("sin", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South China Sea
		Prov("scs").Conn("sar", godip.Sea).Conn("cel", godip.Sea).Conn("phl", godip.Sea).Conn("phs", godip.Sea).Conn("for", godip.Sea).Conn("ecs", godip.Sea).Conn("can", godip.Sea).Conn("hon", godip.Sea).Conn("can", godip.Sea).Conn("han", godip.Sea).Conn("hue", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Yunnan
		Prov("yun").Conn("han", godip.Land).Conn("can", godip.Land).Conn("chu", godip.Land).Conn("asm", godip.Land).Conn("bur", godip.Land).Conn("lao", godip.Land).Flag(godip.Land).
		// Korea
		Prov("kor").Conn("ecs", godip.Sea).Conn("soj", godip.Sea).Conn("mar", godip.Coast...).Conn("man", godip.Coast...).Conn("yel", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Borneo
		Prov("bor").Conn("cel", godip.Sea).Conn("sar", godip.Coast...).Conn("kar", godip.Sea).Conn("jse", godip.Sea).Flag(godip.Coast...).SC(Holland).
		// Akita
		Prov("aki").Conn("shi", godip.Land).Conn("tok", godip.Land).Conn("yes", godip.Coast...).Conn("soj", godip.Sea).Conn("kyo", godip.Coast...).Flag(godip.Coast...).
		// Hanoi
		Prov("han").Conn("yun", godip.Land).Conn("lao", godip.Land).Conn("hue", godip.Coast...).Conn("scs", godip.Sea).Conn("can", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Tibet
		Prov("tib").Conn("asm", godip.Land).Conn("chu", godip.Land).Conn("kan", godip.Land).Conn("sin", godip.Land).Conn("kas", godip.Land).Flag(godip.Land).SC(China).
		// Baghdad
		Prov("bag").Conn("arb", godip.Coast...).Conn("peg", godip.Sea).Conn("pes", godip.Coast...).Conn("con", godip.Land).Conn("dam", godip.Land).Flag(godip.Coast...).SC(Turkey).
		// Mediterranean Sea
		Prov("med").Conn("egy", godip.Sea).Conn("dam", godip.Sea).Conn("con", godip.Sea).Conn("bul", godip.Sea).Flag(godip.Sea).
		// Karimata Strait
		Prov("kar").Conn("scs", godip.Sea).Conn("hue", godip.Sea).Conn("sai", godip.Sea).Conn("gos", godip.Sea).Conn("jse", godip.Sea).Conn("bor", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// South Indian Ocean
		Prov("sio").Conn("ban", godip.Sea).Conn("jav", godip.Sea).Conn("jse", godip.Sea).Conn("sum", godip.Sea).Conn("eio", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// Armenia
		Prov("arm").Conn("kaz", godip.Land).Conn("sev", godip.Coast...).Conn("bla", godip.Sea).Conn("con", godip.Coast...).Conn("pes", godip.Land).Flag(godip.Coast...).
		// Tokyo
		Prov("tok").Conn("cen", godip.Sea).Conn("yes", godip.Coast...).Conn("aki", godip.Land).Conn("shi", godip.Coast...).Conn("phs", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Peking
		Prov("pek").Conn("sha", godip.Coast...).Conn("yel", godip.Sea).Conn("man", godip.Coast...).Conn("mon", godip.Land).Conn("kan", godip.Land).Conn("chu", godip.Land).Flag(godip.Coast...).SC(China).
		// Arabia
		Prov("arb").Conn("bag", godip.Coast...).Conn("dam", godip.Coast...).Conn("red", godip.Sea).Conn("ars", godip.Sea).Conn("peg", godip.Sea).Flag(godip.Coast...).
		// Gulf of Martaban
		Prov("gom").Conn("sum", godip.Sea).Conn("jse", godip.Sea).Conn("mal", godip.Sea).Conn("sia", godip.Sea).Conn("bur", godip.Sea).Conn("bay", godip.Sea).Conn("eio", godip.Sea).Flag(godip.Sea).
		// Baluchistan
		Prov("bal").Conn("ars", godip.Sea).Conn("bom", godip.Coast...).Conn("del", godip.Land).Conn("kas", godip.Land).Conn("afg", godip.Land).Conn("pes", godip.Coast...).Conn("peg", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sasebo
		Prov("sas").Conn("kyo", godip.Coast...).Conn("soj", godip.Sea).Conn("ecs", godip.Sea).Conn("phs", godip.Sea).Conn("shi", godip.Coast...).Flag(godip.Coast...).SC(Japan).
		// Calcutta
		Prov("cal").Conn("del", godip.Land).Conn("mad", godip.Coast...).Conn("bay", godip.Sea).Conn("bur", godip.Coast...).Conn("asm", godip.Land).Flag(godip.Coast...).SC(Britain).
		// Kyoto
		Prov("kyo").Conn("sas", godip.Coast...).Conn("shi", godip.Land).Conn("aki", godip.Coast...).Conn("soj", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Bulgaria
		Prov("bul").Conn("bla", godip.Sea).Conn("rum", godip.Coast...).Conn("med", godip.Sea).Conn("con", godip.Coast...).Flag(godip.Coast...).
		// Delhi
		Prov("del").Conn("bom", godip.Land).Conn("mad", godip.Land).Conn("cal", godip.Land).Conn("kas", godip.Land).Conn("bal", godip.Land).Flag(godip.Land).SC(Britain).
		// Java Sea
		Prov("jse").Conn("gos", godip.Sea).Conn("mal", godip.Sea).Conn("gom", godip.Sea).Conn("sum", godip.Sea).Conn("sio", godip.Sea).Conn("jav", godip.Sea).Conn("ban", godip.Sea).Conn("cel", godip.Sea).Conn("bor", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Chungking
		Prov("chu").Conn("pek", godip.Land).Conn("kan", godip.Land).Conn("tib", godip.Land).Conn("asm", godip.Land).Conn("yun", godip.Land).Conn("can", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).SC(China).
		// Sumatra
		Prov("sum").Conn("gom", godip.Sea).Conn("eio", godip.Sea).Conn("sio", godip.Sea).Conn("jse", godip.Sea).Flag(godip.Coast...).SC(Holland).
		// Sarawak
		Prov("sar").Conn("scs", godip.Sea).Conn("kar", godip.Sea).Conn("bor", godip.Coast...).Conn("cel", godip.Sea).Flag(godip.Coast...).
		// Assam
		Prov("asm").Conn("tib", godip.Land).Conn("cal", godip.Land).Conn("bur", godip.Land).Conn("yun", godip.Land).Conn("chu", godip.Land).Flag(godip.Land).
		// Maritime Province
		Prov("mar").Conn("amu", godip.Land).Conn("man", godip.Land).Conn("kor", godip.Coast...).Conn("soj", godip.Sea).Conn("kha", godip.Coast...).Flag(godip.Coast...).
		// Sea of Okhotsk
		Prov("soo").Conn("kha", godip.Sea).Conn("soj", godip.Sea).Conn("soj", godip.Sea).Conn("yes", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Sea).
		// Central Pacific Ocean
		Prov("cen").Conn("soo", godip.Sea).Conn("yes", godip.Sea).Conn("tok", godip.Sea).Conn("phs", godip.Sea).Conn("phl", godip.Sea).Conn("cel", godip.Sea).Flag(godip.Sea).
		// Philippine Sea
		Prov("phs").Conn("shi", godip.Sea).Conn("sas", godip.Sea).Conn("ecs", godip.Sea).Conn("for", godip.Sea).Conn("scs", godip.Sea).Conn("phl", godip.Sea).Conn("cen", godip.Sea).Conn("tok", godip.Sea).Flag(godip.Sea).
		// Saigon
		Prov("sai").Conn("kar", godip.Sea).Conn("hue", godip.Coast...).Conn("cam", godip.Coast...).Conn("gos", godip.Sea).Flag(godip.Coast...).SC(France).
		// Java
		Prov("jav").Conn("jse", godip.Sea).Conn("sio", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Coast...).SC(Holland).
		// Malaya
		Prov("mal").Conn("sia", godip.Coast...).Conn("gom", godip.Sea).Conn("jse", godip.Sea).Conn("gos", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Arabian Sea
		Prov("ars").Conn("bal", godip.Sea).Conn("peg", godip.Sea).Conn("arb", godip.Sea).Conn("red", godip.Sea).Conn("wes", godip.Sea).Conn("bom", godip.Sea).Flag(godip.Sea).
		// Celebes Sea
		Prov("cel").Conn("cen", godip.Sea).Conn("phl", godip.Sea).Conn("scs", godip.Sea).Conn("sar", godip.Sea).Conn("bor", godip.Sea).Conn("jse", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Sea).
		// Philippines
		Prov("phl").Conn("phs", godip.Sea).Conn("scs", godip.Sea).Conn("cel", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Manchuria
		Prov("man").Conn("amu", godip.Land).Conn("mon", godip.Land).Conn("pek", godip.Coast...).Conn("yel", godip.Sea).Conn("kor", godip.Coast...).Conn("mar", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Siam
		Prov("gos").Conn("jse", godip.Sea).Conn("kar", godip.Sea).Conn("sai", godip.Sea).Conn("cam", godip.Sea).Conn("sia", godip.Sea).Conn("mal", godip.Sea).Flag(godip.Sea).
		// Constantinople
		Prov("con").Conn("bla", godip.Sea).Conn("bul", godip.Coast...).Conn("med", godip.Sea).Conn("dam", godip.Coast...).Conn("bag", godip.Land).Conn("pes", godip.Land).Conn("arm", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Bay of Bengal
		Prov("bay").Conn("eio", godip.Sea).Conn("gom", godip.Sea).Conn("bur", godip.Sea).Conn("cal", godip.Sea).Conn("mad", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// Canton
		Prov("can").Conn("yun", godip.Land).Conn("han", godip.Coast...).Conn("scs", godip.Sea).Conn("hon", godip.Coast...).Conn("hon", godip.Coast...).Conn("scs", godip.Sea).Conn("ecs", godip.Sea).Conn("sha", godip.Coast...).Conn("chu", godip.Land).Flag(godip.Coast...).
		// Kansu
		Prov("kan").Conn("sin", godip.Land).Conn("tib", godip.Land).Conn("chu", godip.Land).Conn("pek", godip.Land).Conn("mon", godip.Land).Flag(godip.Land).
		// Persian Gulf
		Prov("peg").Conn("arb", godip.Sea).Conn("ars", godip.Sea).Conn("bal", godip.Sea).Conn("pes", godip.Sea).Conn("bag", godip.Sea).Flag(godip.Sea).
		// Laos
		Prov("lao").Conn("bur", godip.Land).Conn("sia", godip.Land).Conn("cam", godip.Land).Conn("hue", godip.Land).Conn("han", godip.Land).Conn("yun", godip.Land).Flag(godip.Land).
		// Sea of Japan
		Prov("soj").Conn("soo", godip.Sea).Conn("soo", godip.Sea).Conn("kha", godip.Sea).Conn("mar", godip.Sea).Conn("kor", godip.Sea).Conn("ecs", godip.Sea).Conn("sas", godip.Sea).Conn("kyo", godip.Sea).Conn("aki", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Sea).
		// Hue
		Prov("hue").Conn("han", godip.Coast...).Conn("lao", godip.Land).Conn("cam", godip.Land).Conn("sai", godip.Coast...).Conn("kar", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Coast...).SC(France).
		// Kazakhstan
		Prov("kaz").Conn("tur", godip.Land).Conn("tur", godip.Land).Conn("tom", godip.Land).Conn("sib", godip.Land).Conn("mos", godip.Land).Conn("sev", godip.Land).Conn("arm", godip.Land).Flag(godip.Land).
		// West Indian Ocean
		Prov("wes").Conn("sio", godip.Sea).Conn("eio", godip.Sea).Conn("bay", godip.Sea).Conn("mad", godip.Sea).Conn("bom", godip.Sea).Conn("ars", godip.Sea).Conn("red", godip.Sea).Flag(godip.Sea).
		// Sinkiang
		Prov("sin").Conn("tom", godip.Land).Conn("tur", godip.Land).Conn("afg", godip.Land).Conn("kas", godip.Land).Conn("tib", godip.Land).Conn("kan", godip.Land).Conn("mon", godip.Land).Flag(godip.Land).
		// Shikoku
		Prov("shi").Conn("aki", godip.Land).Conn("kyo", godip.Land).Conn("sas", godip.Coast...).Conn("phs", godip.Sea).Conn("tok", godip.Coast...).Flag(godip.Coast...).
		// Mongolia
		Prov("mon").Conn("irk", godip.Land).Conn("tom", godip.Land).Conn("sin", godip.Land).Conn("kan", godip.Land).Conn("pek", godip.Land).Conn("man", godip.Land).Conn("amu", godip.Land).Flag(godip.Land).
		// Banda Sea
		Prov("ban").Conn("cel", godip.Sea).Conn("jse", godip.Sea).Conn("jav", godip.Sea).Conn("sio", godip.Sea).Flag(godip.Sea).
		// Yellow Sea
		Prov("yel").Conn("ecs", godip.Sea).Conn("kor", godip.Sea).Conn("man", godip.Sea).Conn("pek", godip.Sea).Conn("sha", godip.Sea).Flag(godip.Sea).
		// Turkestan
		Prov("tur").Conn("tom", godip.Land).Conn("tom", godip.Land).Conn("kaz", godip.Land).Conn("kaz", godip.Land).Conn("pes", godip.Land).Conn("afg", godip.Land).Conn("sin", godip.Land).Flag(godip.Land).
		// East China Sea
		Prov("ecs").Conn("yel", godip.Sea).Conn("sha", godip.Sea).Conn("can", godip.Sea).Conn("scs", godip.Sea).Conn("for", godip.Sea).Conn("phs", godip.Sea).Conn("sas", godip.Sea).Conn("soj", godip.Sea).Conn("kor", godip.Sea).Flag(godip.Sea).
		// Rumania
		Prov("rum").Conn("bul", godip.Coast...).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shanghai
		Prov("sha").Conn("pek", godip.Coast...).Conn("chu", godip.Land).Conn("can", godip.Coast...).Conn("ecs", godip.Sea).Conn("yel", godip.Sea).Flag(godip.Coast...).SC(China).
		// Black Sea
		Prov("bla").Conn("bul", godip.Sea).Conn("con", godip.Sea).Conn("arm", godip.Sea).Conn("sev", godip.Sea).Conn("rum", godip.Sea).Flag(godip.Sea).
		// Irkutsk
		Prov("irk").Conn("sib", godip.Land).Conn("tom", godip.Land).Conn("mon", godip.Land).Conn("amu", godip.Land).Conn("kha", godip.Land).Flag(godip.Land).SC(Russia).
		// Egypt
		Prov("egy").Conn("red", godip.Sea).Conn("dam", godip.Coast...).Conn("med", godip.Sea).Flag(godip.Coast...).
		// Formosa
		Prov("for").Conn("ecs", godip.Sea).Conn("scs", godip.Sea).Conn("phs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Red Sea
		Prov("red").Conn("dam", godip.Sea).Conn("egy", godip.Sea).Conn("wes", godip.Sea).Conn("ars", godip.Sea).Conn("arb", godip.Sea).Flag(godip.Sea).
		// East Indian Ocean
		Prov("eio").Conn("sio", godip.Sea).Conn("sum", godip.Sea).Conn("gom", godip.Sea).Conn("bay", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// Khabarovsk
		Prov("kha").Conn("irk", godip.Land).Conn("amu", godip.Land).Conn("mar", godip.Coast...).Conn("soj", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Coast...).SC(Russia).
		// Burma
		Prov("bur").Conn("lao", godip.Land).Conn("yun", godip.Land).Conn("asm", godip.Land).Conn("cal", godip.Coast...).Conn("bay", godip.Sea).Conn("gom", godip.Sea).Conn("sia", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Yesso
		Prov("yes").Conn("cen", godip.Sea).Conn("soo", godip.Sea).Conn("soj", godip.Sea).Conn("aki", godip.Coast...).Conn("tok", godip.Coast...).Flag(godip.Coast...).
		// Siberia
		Prov("sib").Conn("mos", godip.Land).Conn("kaz", godip.Land).Conn("tom", godip.Land).Conn("irk", godip.Land).Flag(godip.Land).
		// Persia
		Prov("pes").Conn("arm", godip.Land).Conn("con", godip.Land).Conn("bag", godip.Coast...).Conn("peg", godip.Sea).Conn("bal", godip.Coast...).Conn("afg", godip.Land).Conn("tur", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Cambodia
		Prov("cam").Conn("lao", godip.Land).Conn("sia", godip.Coast...).Conn("gos", godip.Sea).Conn("sai", godip.Coast...).Conn("hue", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Tomsk
		Prov("tom").Conn("sin", godip.Land).Conn("mon", godip.Land).Conn("irk", godip.Land).Conn("sib", godip.Land).Conn("kaz", godip.Land).Conn("tur", godip.Land).Conn("tur", godip.Land).Flag(godip.Land).
		// Moscow
		Prov("mos").Conn("sev", godip.Land).Conn("kaz", godip.Land).Conn("sib", godip.Land).Flag(godip.Land).SC(Russia).
		// Madras
		Prov("mad").Conn("bay", godip.Sea).Conn("cal", godip.Coast...).Conn("del", godip.Land).Conn("bom", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Damascus
		Prov("dam").Conn("red", godip.Sea).Conn("arb", godip.Coast...).Conn("bag", godip.Land).Conn("con", godip.Coast...).Conn("med", godip.Sea).Conn("egy", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		Done()
}
