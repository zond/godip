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
	Name:              "Canton",
	Graph:             func() godip.Graph { return CantonGraph() },
	Start:             CantonStart,
	Blank:             CantonBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(19),
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/cantonmap.svg")
	},
	SVGVersion: "10",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Paul Webb",
	Version:     "3",
	SoloSCCount: func(*state.State) int { return 19 },
	Description: "Asia at the beginning of the 20th century. It's fate is determined by Britain, China, France, Holland, Japan, Russia, and Turkey.",
	Rules: `First to 19 Supply Centers (SC) is the winner. 
	Constantinople and Egypt have a canal (similar to Kiel in Classic). 
	Four provinces have dual coasts: Damascus, Bulgaria, Siam and Canton.`,
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
		Prov("bom").Conn("wio", godip.Sea).Conn("mad", godip.Coast...).Conn("del", godip.Land).Conn("bal", godip.Coast...).Conn("ase", godip.Sea).Flag(godip.Coast...).
		// Sevastopol
		Prov("sev").Conn("rum", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Conn("kaz", godip.Land).Conn("mos", godip.Land).Flag(godip.Coast...).SC(Russia).
		// Amur
		Prov("amu").Conn("irk", godip.Land).Conn("mon", godip.Land).Conn("man", godip.Land).Conn("mar", godip.Land).Conn("kha", godip.Land).Flag(godip.Land).
		// Kashmir
		Prov("kas").Conn("bal", godip.Land).Conn("del", godip.Land).Conn("tib", godip.Land).Conn("sin", godip.Land).Conn("afg", godip.Land).Flag(godip.Land).
		// Hong Kong
		Prov("hko").Conn("scs", godip.Sea).Conn("can", godip.Land).Conn("can/sc", godip.Sea).Conn("can/ec", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Siam
		Prov("sia").Conn("mal", godip.Land).Conn("cam", godip.Land).Conn("lao", godip.Land).Conn("bur", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Siam (West Coast)
		Prov("sia/wc").Conn("mal", godip.Sea).Conn("bur", godip.Sea).Conn("gom", godip.Sea).Flag(godip.Sea).
		// Siam (East Coast)
		Prov("sia/ec").Conn("mal", godip.Sea).Conn("gos", godip.Sea).Conn("cam", godip.Sea).Flag(godip.Sea).
		// Afghanistan
		Prov("afg").Conn("tur", godip.Land).Conn("per", godip.Land).Conn("bal", godip.Land).Conn("kas", godip.Land).Conn("sin", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South China Sea
		Prov("scs").Conn("sar", godip.Sea).Conn("cel", godip.Sea).Conn("phi", godip.Sea).Conn("pse", godip.Sea).Conn("for", godip.Sea).Conn("ecs", godip.Sea).Conn("can", godip.Sea).Conn("can/sc", godip.Sea).Conn("can/ec", godip.Sea).Conn("hko", godip.Sea).Conn("han", godip.Sea).Conn("hue", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Yunnan
		Prov("yun").Conn("han", godip.Land).Conn("can", godip.Land).Conn("chu", godip.Land).Conn("asm", godip.Land).Conn("bur", godip.Land).Conn("lao", godip.Land).Flag(godip.Land).
		// Korea
		Prov("kor").Conn("ecs", godip.Sea).Conn("soj", godip.Sea).Conn("mar", godip.Coast...).Conn("man", godip.Coast...).Conn("yel", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Borneo
		Prov("bor").Conn("cel", godip.Sea).Conn("sar", godip.Coast...).Conn("kar", godip.Sea).Conn("jse", godip.Sea).Flag(godip.Coast...).SC(Holland).
		// Akita
		Prov("aki").Conn("shi", godip.Land).Conn("tok", godip.Land).Conn("yes", godip.Coast...).Conn("soj", godip.Sea).Conn("kyo", godip.Coast...).Flag(godip.Coast...).
		// Hanoi
		Prov("han").Conn("yun", godip.Land).Conn("lao", godip.Land).Conn("hue", godip.Coast...).Conn("scs", godip.Sea).Conn("can", godip.Land).Conn("can/sc", godip.Sea).Flag(godip.Coast...).SC(France).
		// Tibet
		Prov("tib").Conn("asm", godip.Land).Conn("chu", godip.Land).Conn("kan", godip.Land).Conn("sin", godip.Land).Conn("kas", godip.Land).Flag(godip.Land).SC(China).
		// Baghdad
		Prov("bag").Conn("ara", godip.Coast...).Conn("pgu", godip.Sea).Conn("per", godip.Coast...).Conn("con", godip.Land).Conn("dam", godip.Land).Flag(godip.Coast...).SC(Turkey).
		// Mediterranean Sea
		Prov("med").Conn("egy", godip.Sea).Conn("dam", godip.Sea).Conn("dam/wc", godip.Sea).Conn("con", godip.Sea).Conn("bul", godip.Sea).Conn("bul/sc", godip.Sea).Flag(godip.Sea).
		// Karimata Strait
		Prov("kar").Conn("scs", godip.Sea).Conn("hue", godip.Sea).Conn("sai", godip.Sea).Conn("gos", godip.Sea).Conn("jse", godip.Sea).Conn("bor", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// South Indian Ocean
		Prov("sio").Conn("ban", godip.Sea).Conn("jav", godip.Sea).Conn("jse", godip.Sea).Conn("sum", godip.Sea).Conn("eio", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Sea).
		// Armenia
		Prov("arm").Conn("kaz", godip.Land).Conn("sev", godip.Coast...).Conn("bla", godip.Sea).Conn("con", godip.Coast...).Conn("per", godip.Land).Flag(godip.Coast...).
		// Tokyo
		Prov("tok").Conn("cpo", godip.Sea).Conn("yes", godip.Coast...).Conn("aki", godip.Land).Conn("shi", godip.Coast...).Conn("pse", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Peking
		Prov("pek").Conn("sha", godip.Coast...).Conn("yel", godip.Sea).Conn("man", godip.Coast...).Conn("mon", godip.Land).Conn("kan", godip.Land).Conn("chu", godip.Land).Flag(godip.Coast...).SC(China).
		// Arabia
		Prov("ara").Conn("bag", godip.Coast...).Conn("dam", godip.Land).Conn("dam/sc", godip.Sea).Conn("red", godip.Sea).Conn("ase", godip.Sea).Conn("pgu", godip.Sea).Flag(godip.Coast...).
		// Gulf of Martaban
		Prov("gom").Conn("sum", godip.Sea).Conn("jse", godip.Sea).Conn("mal", godip.Sea).Conn("sia", godip.Sea).Conn("sia/wc", godip.Sea).Conn("bur", godip.Sea).Conn("bob", godip.Sea).Conn("eio", godip.Sea).Flag(godip.Sea).
		// Baluchistan
		Prov("bal").Conn("ase", godip.Sea).Conn("bom", godip.Coast...).Conn("del", godip.Land).Conn("kas", godip.Land).Conn("afg", godip.Land).Conn("per", godip.Coast...).Conn("pgu", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sasebo
		Prov("sas").Conn("kyo", godip.Coast...).Conn("soj", godip.Sea).Conn("ecs", godip.Sea).Conn("pse", godip.Sea).Conn("shi", godip.Coast...).Flag(godip.Coast...).SC(Japan).
		// Calcutta
		Prov("cal").Conn("del", godip.Land).Conn("mad", godip.Coast...).Conn("bob", godip.Sea).Conn("bur", godip.Coast...).Conn("asm", godip.Land).Flag(godip.Coast...).SC(Britain).
		// Kyoto
		Prov("kyo").Conn("sas", godip.Coast...).Conn("shi", godip.Land).Conn("aki", godip.Coast...).Conn("soj", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Bulgaria
		Prov("bul").Conn("rum", godip.Land).Conn("con", godip.Land).Flag(godip.Land).
		// Bulgaria (East Coast)
		Prov("bul/ec").Conn("bla", godip.Sea).Conn("rum", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
		// Bulgaria (South Coast)
		Prov("bul/sc").Conn("med", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
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
		Prov("soo").Conn("kha", godip.Sea).Conn("soj", godip.Sea).Conn("soj", godip.Sea).Conn("yes", godip.Sea).Conn("cpo", godip.Sea).Flag(godip.Sea).
		// Central Pacific Ocean
		Prov("cpo").Conn("soo", godip.Sea).Conn("yes", godip.Sea).Conn("tok", godip.Sea).Conn("pse", godip.Sea).Conn("phi", godip.Sea).Conn("cel", godip.Sea).Flag(godip.Sea).
		// Philippine Sea
		Prov("pse").Conn("shi", godip.Sea).Conn("sas", godip.Sea).Conn("ecs", godip.Sea).Conn("for", godip.Sea).Conn("scs", godip.Sea).Conn("phi", godip.Sea).Conn("cpo", godip.Sea).Conn("tok", godip.Sea).Flag(godip.Sea).
		// Saigon
		Prov("sai").Conn("kar", godip.Sea).Conn("hue", godip.Coast...).Conn("cam", godip.Coast...).Conn("gos", godip.Sea).Flag(godip.Coast...).SC(France).
		// Java
		Prov("jav").Conn("jse", godip.Sea).Conn("sio", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Coast...).SC(Holland).
		// Malaya
		Prov("mal").Conn("sia", godip.Land).Conn("sia/wc", godip.Sea).Conn("sia/ec", godip.Sea).Conn("gom", godip.Sea).Conn("jse", godip.Sea).Conn("gos", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Arabian Sea
		Prov("ase").Conn("bal", godip.Sea).Conn("pgu", godip.Sea).Conn("ara", godip.Sea).Conn("red", godip.Sea).Conn("wio", godip.Sea).Conn("bom", godip.Sea).Flag(godip.Sea).
		// Celebes Sea
		Prov("cel").Conn("cpo", godip.Sea).Conn("phi", godip.Sea).Conn("scs", godip.Sea).Conn("sar", godip.Sea).Conn("bor", godip.Sea).Conn("jse", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Sea).
		// Philippines
		Prov("phi").Conn("pse", godip.Sea).Conn("scs", godip.Sea).Conn("cel", godip.Sea).Conn("cpo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Manchuria
		Prov("man").Conn("amu", godip.Land).Conn("mon", godip.Land).Conn("pek", godip.Coast...).Conn("yel", godip.Sea).Conn("kor", godip.Coast...).Conn("mar", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Siam
		Prov("gos").Conn("jse", godip.Sea).Conn("kar", godip.Sea).Conn("sai", godip.Sea).Conn("cam", godip.Sea).Conn("sia", godip.Sea).Conn("sia/ec", godip.Sea).Conn("mal", godip.Sea).Flag(godip.Sea).
		// Constantinople
		Prov("con").Conn("bla", godip.Sea).Conn("bul", godip.Land).Conn("bul/ec", godip.Sea).Conn("bul/sc", godip.Sea).Conn("med", godip.Sea).Conn("dam", godip.Land).Conn("dam/wc", godip.Sea).Conn("bag", godip.Land).Conn("per", godip.Land).Conn("arm", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Bay of Bengal
		Prov("bob").Conn("eio", godip.Sea).Conn("gom", godip.Sea).Conn("bur", godip.Sea).Conn("cal", godip.Sea).Conn("mad", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Sea).
		// Canton
		Prov("can").Conn("yun", godip.Land).Conn("han", godip.Land).Conn("hko", godip.Land).Conn("sha", godip.Land).Conn("chu", godip.Land).Flag(godip.Land).
		// Canton (South Coast)
		Prov("can/sc").Conn("han", godip.Sea).Conn("scs", godip.Sea).Conn("hko", godip.Sea).Flag(godip.Sea).
		// Canton (East Coast)
		Prov("can/ec").Conn("hko", godip.Sea).Conn("scs", godip.Sea).Conn("ecs", godip.Sea).Conn("sha", godip.Sea).Flag(godip.Sea).
		// Kansu
		Prov("kan").Conn("sin", godip.Land).Conn("tib", godip.Land).Conn("chu", godip.Land).Conn("pek", godip.Land).Conn("mon", godip.Land).Flag(godip.Land).
		// Persian Gulf
		Prov("pgu").Conn("ara", godip.Sea).Conn("ase", godip.Sea).Conn("bal", godip.Sea).Conn("per", godip.Sea).Conn("bag", godip.Sea).Flag(godip.Sea).
		// Laos
		Prov("lao").Conn("bur", godip.Land).Conn("sia", godip.Land).Conn("cam", godip.Land).Conn("hue", godip.Land).Conn("han", godip.Land).Conn("yun", godip.Land).Flag(godip.Land).
		// Sea of Japan
		Prov("soj").Conn("soo", godip.Sea).Conn("soo", godip.Sea).Conn("kha", godip.Sea).Conn("mar", godip.Sea).Conn("kor", godip.Sea).Conn("ecs", godip.Sea).Conn("sas", godip.Sea).Conn("kyo", godip.Sea).Conn("aki", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Sea).
		// Hue
		Prov("hue").Conn("han", godip.Coast...).Conn("lao", godip.Land).Conn("cam", godip.Land).Conn("sai", godip.Coast...).Conn("kar", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Coast...).SC(France).
		// Kazakhstan
		Prov("kaz").Conn("tur", godip.Land).Conn("tur", godip.Land).Conn("tom", godip.Land).Conn("sib", godip.Land).Conn("mos", godip.Land).Conn("sev", godip.Land).Conn("arm", godip.Land).Flag(godip.Land).
		// West Indian Ocean
		Prov("wio").Conn("sio", godip.Sea).Conn("eio", godip.Sea).Conn("bob", godip.Sea).Conn("mad", godip.Sea).Conn("bom", godip.Sea).Conn("ase", godip.Sea).Conn("red", godip.Sea).Flag(godip.Sea).
		// Sinkiang
		Prov("sin").Conn("tom", godip.Land).Conn("tur", godip.Land).Conn("afg", godip.Land).Conn("kas", godip.Land).Conn("tib", godip.Land).Conn("kan", godip.Land).Conn("mon", godip.Land).Flag(godip.Land).
		// Shikoku
		Prov("shi").Conn("aki", godip.Land).Conn("kyo", godip.Land).Conn("sas", godip.Coast...).Conn("pse", godip.Sea).Conn("tok", godip.Coast...).Flag(godip.Coast...).
		// Mongolia
		Prov("mon").Conn("irk", godip.Land).Conn("tom", godip.Land).Conn("sin", godip.Land).Conn("kan", godip.Land).Conn("pek", godip.Land).Conn("man", godip.Land).Conn("amu", godip.Land).Flag(godip.Land).
		// Banda Sea
		Prov("ban").Conn("cel", godip.Sea).Conn("jse", godip.Sea).Conn("jav", godip.Sea).Conn("sio", godip.Sea).Flag(godip.Sea).
		// Yellow Sea
		Prov("yel").Conn("ecs", godip.Sea).Conn("kor", godip.Sea).Conn("man", godip.Sea).Conn("pek", godip.Sea).Conn("sha", godip.Sea).Flag(godip.Sea).
		// Turkestan
		Prov("tur").Conn("tom", godip.Land).Conn("tom", godip.Land).Conn("kaz", godip.Land).Conn("kaz", godip.Land).Conn("per", godip.Land).Conn("afg", godip.Land).Conn("sin", godip.Land).Flag(godip.Land).
		// East China Sea
		Prov("ecs").Conn("yel", godip.Sea).Conn("sha", godip.Sea).Conn("can", godip.Sea).Conn("can/ec", godip.Sea).Conn("scs", godip.Sea).Conn("for", godip.Sea).Conn("pse", godip.Sea).Conn("sas", godip.Sea).Conn("soj", godip.Sea).Conn("kor", godip.Sea).Flag(godip.Sea).
		// Rumania
		Prov("rum").Conn("bul", godip.Land).Conn("bul/ec", godip.Sea).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shanghai
		Prov("sha").Conn("pek", godip.Coast...).Conn("chu", godip.Land).Conn("can", godip.Land).Conn("can/ec", godip.Sea).Conn("ecs", godip.Sea).Conn("yel", godip.Sea).Flag(godip.Coast...).SC(China).
		// Black Sea
		Prov("bla").Conn("bul", godip.Sea).Conn("bul/ec", godip.Sea).Conn("con", godip.Sea).Conn("arm", godip.Sea).Conn("sev", godip.Sea).Conn("rum", godip.Sea).Flag(godip.Sea).
		// Irkutsk
		Prov("irk").Conn("sib", godip.Land).Conn("tom", godip.Land).Conn("mon", godip.Land).Conn("amu", godip.Land).Conn("kha", godip.Land).Flag(godip.Land).SC(Russia).
		// Egypt
		Prov("egy").Conn("red", godip.Sea).Conn("dam", godip.Land).Conn("dam/wc", godip.Sea).Conn("dam/sc", godip.Sea).Conn("med", godip.Sea).Flag(godip.Coast...).
		// Formosa
		Prov("for").Conn("ecs", godip.Sea).Conn("scs", godip.Sea).Conn("pse", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Red Sea
		Prov("red").Conn("dam", godip.Sea).Conn("dam/sc", godip.Sea).Conn("egy", godip.Sea).Conn("wio", godip.Sea).Conn("ase", godip.Sea).Conn("ara", godip.Sea).Flag(godip.Sea).
		// East Indian Ocean
		Prov("eio").Conn("sio", godip.Sea).Conn("sum", godip.Sea).Conn("gom", godip.Sea).Conn("bob", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Sea).
		// Khabarovsk
		Prov("kha").Conn("irk", godip.Land).Conn("amu", godip.Land).Conn("mar", godip.Coast...).Conn("soj", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Coast...).SC(Russia).
		// Burma
		Prov("bur").Conn("lao", godip.Land).Conn("yun", godip.Land).Conn("asm", godip.Land).Conn("cal", godip.Coast...).Conn("bob", godip.Sea).Conn("gom", godip.Sea).Conn("sia", godip.Land).Conn("sia/wc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Yesso
		Prov("yes").Conn("cpo", godip.Sea).Conn("soo", godip.Sea).Conn("soj", godip.Sea).Conn("aki", godip.Coast...).Conn("tok", godip.Coast...).Flag(godip.Coast...).
		// Siberia
		Prov("sib").Conn("mos", godip.Land).Conn("kaz", godip.Land).Conn("tom", godip.Land).Conn("irk", godip.Land).Flag(godip.Land).
		// Persia
		Prov("per").Conn("arm", godip.Land).Conn("con", godip.Land).Conn("bag", godip.Coast...).Conn("pgu", godip.Sea).Conn("bal", godip.Coast...).Conn("afg", godip.Land).Conn("tur", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Cambodia
		Prov("cam").Conn("lao", godip.Land).Conn("sia", godip.Land).Conn("sia/ec", godip.Sea).Conn("gos", godip.Sea).Conn("sai", godip.Coast...).Conn("hue", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Tomsk
		Prov("tom").Conn("sin", godip.Land).Conn("mon", godip.Land).Conn("irk", godip.Land).Conn("sib", godip.Land).Conn("kaz", godip.Land).Conn("tur", godip.Land).Conn("tur", godip.Land).Flag(godip.Land).
		// Moscow
		Prov("mos").Conn("sev", godip.Land).Conn("kaz", godip.Land).Conn("sib", godip.Land).Flag(godip.Land).SC(Russia).
		// Madras
		Prov("mad").Conn("bob", godip.Sea).Conn("cal", godip.Coast...).Conn("del", godip.Land).Conn("bom", godip.Coast...).Conn("wio", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Damascus
		Prov("dam").Conn("ara", godip.Land).Conn("bag", godip.Land).Conn("con", godip.Land).Conn("egy", godip.Land).Flag(godip.Land).SC(Turkey).
		// Damascus (South Coast)
		Prov("dam/sc").Conn("red", godip.Sea).Conn("ara", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Sea).
		// Damascus (West Coast)
		Prov("dam/wc").Conn("con", godip.Sea).Conn("med", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"kyo":    "Kyoto",
	"shi":    "Shikoku",
	"sas":    "Sasebo",
	"con":    "Constantinople",
	"mar":    "Primorsky Krai",
	"kha":    "Khabarovsk",
	"irk":    "Irkutsk",
	"amu":    "Amur",
	"wio":    "West Indian Ocean",
	"ase":    "Arabian Sea",
	"pgu":    "Persian Gulf",
	"ara":    "Arabia",
	"sio":    "South Indian Ocean",
	"eio":    "East Indian Ocean",
	"red":    "Red Sea",
	"egy":    "Egypt",
	"bob":    "Bay of Bengal",
	"gom":    "Gulf of Martaban",
	"sum":    "Sumatra",
	"jav":    "Java",
	"mal":    "Malaya",
	"ban":    "Banda Sea",
	"gos":    "Gulf of Siam",
	"jse":    "Java Sea",
	"kar":    "Karimata Strait",
	"bor":    "Borneo",
	"sar":    "Sarawak",
	"mad":    "Madras",
	"bom":    "Bombay",
	"sai":    "Saigon",
	"hue":    "Hue",
	"lao":    "Laos",
	"cam":    "Cambodia",
	"bur":    "Burma",
	"cal":    "Calcutta",
	"del":    "Delhi",
	"bal":    "Baluchistan",
	"per":    "Persia",
	"bag":    "Baghdad",
	"arm":    "Armenia",
	"rum":    "Rumenia",
	"sev":    "Sevastopol",
	"mos":    "Moscow",
	"kaz":    "Kazakhstan",
	"sib":    "Siberia",
	"tur":    "Turkestan",
	"afg":    "Afghanistan",
	"kas":    "Kashmir",
	"tom":    "Tomsk",
	"sin":    "Sinkiang",
	"tib":    "Tibet",
	"mon":    "Mongolia",
	"man":    "Manchuria",
	"pek":    "Peking",
	"kan":    "Kansu",
	"asm":    "Assam",
	"chu":    "Chunking",
	"yun":    "Yunnan",
	"sha":    "Shanghai",
	"hko":    "Hong Kong",
	"han":    "Hanoi",
	"kor":    "Korea",
	"yes":    "Yeso",
	"aki":    "Akita",
	"tok":    "Tokyo",
	"for":    "Formosa",
	"med":    "Mediterranean Sea",
	"bla":    "Black Sea",
	"phi":    "Philippines",
	"cel":    "Celebes Sea",
	"cpo":    "Central Pacific Ocean",
	"soo":    "Sea of Okhotsk",
	"yel":    "Yellow Sea",
	"soj":    "Sea of Japan",
	"pse":    "Philippine Sea",
	"ecs":    "East China Sea",
	"scs":    "South China Sea",
	"bul/sc": "Bulgaria (SC)",
	"bul/ec": "Bulgaria (EC)",
	"sia/ec": "Siam (EC)",
	"sia/wc": "Siam (WC)",
	"dam/sc": "Damascus (SC)",
	"dam/wc": "Damascus (WC)",
	"can/ec": "Canton (EC)",
	"can/sc": "Canton (SC)",
	"can":    "Canton",
	"bul":    "Bulgaria",
	"sia":    "Siam",
	"dam":    "Damascus",
}
