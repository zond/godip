package youngstownredux

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/variants/classical/common"
)

const (
	Turkey  godip.Nation = "Turkey"
	Austria godip.Nation = "Austria"
	Britain godip.Nation = "Britain"
	China   godip.Nation = "China"
	Japan   godip.Nation = "Japan"
	Italy   godip.Nation = "Italy"
	Germany godip.Nation = "Germany"
	India   godip.Nation = "India"
	Russia  godip.Nation = "Russia"
	France  godip.Nation = "France"
)

var Nations = []godip.Nation{Turkey, Austria, Britain, China, Japan, Italy, Germany, India, Russia, France}

var YoungstownReduxVariant = common.Variant{
	Name:       "Youngstown Redux",
	Graph:      func() godip.Graph { return YoungstownReduxGraph() },
	Start:      YoungstownReduxStart,
	Blank:      YoungstownReduxBlank,
	Phase:      classical.Phase,
	Parser:     orders.ClassicalParser,
	Nations:    Nations,
	PhaseTypes: cla.PhaseTypes,
	Seasons:    cla.Seasons,
	UnitTypes:  cla.UnitTypes,
	SoloWinner: common.SCCountWinner(28),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/youngstownreduxmap.svg")
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
	CreatedBy:   "airborne",
	Version:     "I",
	Description: "A ten player variant that adds China, India and Japan to the standard seven nations.",
	Rules: "Rules are as per classical Diplomacy. There are eight box sea regions which are each connected " +
		"to the other boxes in the same row and column, and allow fleets to travel 'around the world'. Six provinces " +
		"have two coasts (Spain, St. Petersburg, Levant, Arabia, Hebei and Thailand), and all other coastal regions have a " +
		"single coast. The winner is the first nation to 28 supply centers, or the player with the most in the case of " +
		"multiple nations reaching 28 in the same turn. If the leading two nations both have the same number of centers " +
		"then the game will continue for another year.  This variant is based on the Youngstown variant by Rod Walker, " +
		"A. Phillips, Ken Lowe and Jon Monsarret.",
}

func YoungstownReduxBlank(phase godip.Phase) *state.State {
	return state.New(YoungstownReduxGraph(), phase, classical.BackupRule)
}

func YoungstownReduxStart() (result *state.State, err error) {
	startPhase := classical.Phase(1901, godip.Spring, godip.Movement)
	result = state.New(YoungstownReduxGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"ank":    godip.Unit{godip.Fleet, Turkey},
		"con":    godip.Unit{godip.Army, Turkey},
		"bag":    godip.Unit{godip.Army, Turkey},
		"mec":    godip.Unit{godip.Army, Turkey},
		"sar":    godip.Unit{godip.Fleet, Austria},
		"vnn":    godip.Unit{godip.Army, Austria},
		"bud":    godip.Unit{godip.Army, Austria},
		"tes":    godip.Unit{godip.Army, Austria},
		"lon":    godip.Unit{godip.Fleet, Britain},
		"lie":    godip.Unit{godip.Fleet, Britain},
		"edi":    godip.Unit{godip.Fleet, Britain},
		"ade":    godip.Unit{godip.Fleet, Britain},
		"sig":    godip.Unit{godip.Fleet, Britain},
		"sha":    godip.Unit{godip.Fleet, China},
		"pek":    godip.Unit{godip.Army, China},
		"gua":    godip.Unit{godip.Army, China},
		"wuh":    godip.Unit{godip.Army, China},
		"tok":    godip.Unit{godip.Fleet, Japan},
		"osa":    godip.Unit{godip.Fleet, Japan},
		"sap":    godip.Unit{godip.Fleet, Japan},
		"kyo":    godip.Unit{godip.Army, Japan},
		"nap":    godip.Unit{godip.Fleet, Italy},
		"mog":    godip.Unit{godip.Fleet, Italy},
		"rom":    godip.Unit{godip.Army, Italy},
		"mil":    godip.Unit{godip.Army, Italy},
		"tsi":    godip.Unit{godip.Fleet, Germany},
		"kie":    godip.Unit{godip.Fleet, Germany},
		"ber":    godip.Unit{godip.Army, Germany},
		"mun":    godip.Unit{godip.Army, Germany},
		"col":    godip.Unit{godip.Army, Germany},
		"bom":    godip.Unit{godip.Fleet, India},
		"mad":    godip.Unit{godip.Fleet, India},
		"del":    godip.Unit{godip.Army, India},
		"cal":    godip.Unit{godip.Army, India},
		"sev":    godip.Unit{godip.Fleet, Russia},
		"stp/sc": godip.Unit{godip.Fleet, Russia},
		"vla":    godip.Unit{godip.Fleet, Russia},
		"mos":    godip.Unit{godip.Army, Russia},
		"oms":    godip.Unit{godip.Army, Russia},
		"war":    godip.Unit{godip.Army, Russia},
		"bre":    godip.Unit{godip.Fleet, France},
		"sai":    godip.Unit{godip.Fleet, France},
		"par":    godip.Unit{godip.Army, France},
		"mar":    godip.Unit{godip.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"ank": Turkey,
		"con": Turkey,
		"bag": Turkey,
		"mec": Turkey,
		"sar": Austria,
		"vnn": Austria,
		"bud": Austria,
		"tes": Austria,
		"lon": Britain,
		"lie": Britain,
		"edi": Britain,
		"ade": Britain,
		"sig": Britain,
		"sha": China,
		"pek": China,
		"gua": China,
		"wuh": China,
		"tok": Japan,
		"osa": Japan,
		"sap": Japan,
		"kyo": Japan,
		"nap": Italy,
		"mog": Italy,
		"rom": Italy,
		"mil": Italy,
		"tsi": Germany,
		"kie": Germany,
		"ber": Germany,
		"mun": Germany,
		"col": Germany,
		"bom": India,
		"mad": India,
		"del": India,
		"cal": India,
		"sev": Russia,
		"stp": Russia,
		"vla": Russia,
		"mos": Russia,
		"oms": Russia,
		"war": Russia,
		"bre": France,
		"sai": France,
		"par": France,
		"mar": France,
	})
	return
}

func YoungstownReduxGraph() *graph.Graph {
	return graph.New().
		// Box H cfg
		Prov("bxh").Conn("eio", godip.Sea).Conn("eio", godip.Sea).Conn("eio", godip.Sea).Conn("bxc", godip.Sea).Conn("bxf", godip.Sea).Conn("bxg", godip.Sea).Flag(godip.Sea).
		// Tunisia
		Prov("tun").Conn("alg", godip.Coast...).Conn("sah", godip.Land).Conn("trp", godip.Coast...).Conn("ion", godip.Sea).Conn("tyh", godip.Sea).Flag(godip.Coast...).
		// Bombay
		Prov("bom").Conn("mad", godip.Coast...).Conn("dec", godip.Land).Conn("del", godip.Land).Conn("sid", godip.Coast...).Conn("ars", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Coast...).SC(India).
		// Hebei
		Prov("heb").Conn("pek", godip.Land).Conn("qin", godip.Land).Conn("wuh", godip.Land).Conn("sha", godip.Land).Conn("tsi", godip.Land).Flag(godip.Land).
		// Hebei (North Coast)
		Prov("heb/nc").Conn("yel", godip.Sea).Conn("pek", godip.Sea).Conn("tsi", godip.Sea).Flag(godip.Sea).
		// Hebei (South Coast)
		Prov("heb/sc").Conn("sha", godip.Sea).Conn("ecs", godip.Sea).Conn("tsi", godip.Sea).Flag(godip.Sea).
		// Silesia
		Prov("sil").Conn("pru", godip.Land).Conn("ber", godip.Land).Conn("sax", godip.Land).Conn("boh", godip.Land).Conn("gal", godip.Land).Conn("war", godip.Land).Flag(godip.Land).
		// Sevastopol
		Prov("sev").Conn("cau", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("rum", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Flag(godip.Coast...).SC(Russia).
		// Albania
		Prov("alb").Conn("ion", godip.Sea).Conn("gre", godip.Coast...).Conn("mac", godip.Land).Conn("ser", godip.Land).Conn("sar", godip.Coast...).Conn("adr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// St. Petersburg
		Prov("stp").Conn("fin", godip.Land).Conn("lia", godip.Land).Conn("mos", godip.Land).Conn("oms", godip.Land).Conn("nay", godip.Land).Flag(godip.Land).SC(Russia).
		// St. Petersburg (North Coast)
		Prov("stp/nc").Conn("bar", godip.Sea).Conn("nay", godip.Sea).Flag(godip.Sea).
		// St. Petersburg (South Coast)
		Prov("stp/sc").Conn("fin", godip.Sea).Conn("gob", godip.Sea).Conn("lia", godip.Sea).Flag(godip.Sea).
		// Kashmir
		Prov("kas").Conn("del", godip.Land).Conn("tib", godip.Land).Conn("afg", godip.Land).Conn("sid", godip.Land).Flag(godip.Land).
		// Red Sea
		Prov("red").Conn("lev", godip.Sea).Conn("lev/sc", godip.Sea).Conn("egy", godip.Sea).Conn("sud", godip.Sea).Conn("eth", godip.Sea).Conn("goa", godip.Sea).Conn("ade", godip.Sea).Conn("ara", godip.Sea).Conn("ara/sc", godip.Sea).Conn("mec", godip.Sea).Flag(godip.Sea).
		// London
		Prov("lon").Conn("not", godip.Sea).Conn("yor", godip.Coast...).Conn("wal", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Galicia
		Prov("gal").Conn("ukr", godip.Land).Conn("war", godip.Land).Conn("sil", godip.Land).Conn("boh", godip.Land).Conn("vnn", godip.Land).Conn("bud", godip.Land).Conn("rum", godip.Land).Flag(godip.Land).
		// Yemen
		Prov("yem").Conn("goa", godip.Sea).Conn("ars", godip.Sea).Conn("oma", godip.Coast...).Conn("ara", godip.Land).Conn("ade", godip.Coast...).Flag(godip.Coast...).
		// Afghanistan
		Prov("afg").Conn("tib", godip.Land).Conn("tur", godip.Land).Conn("per", godip.Land).Conn("sid", godip.Land).Conn("kas", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South China Sea
		Prov("scs").Conn("guh", godip.Sea).Conn("bor", godip.Sea).Conn("cel", godip.Sea).Conn("phi", godip.Sea).Conn("ecs", godip.Sea).Conn("for", godip.Sea).Conn("ecs", godip.Sea).Conn("sha", godip.Sea).Conn("gua", godip.Sea).Conn("goo", godip.Sea).Conn("ann", godip.Sea).Conn("sai", godip.Sea).Flag(godip.Sea).
		// Levant
		Prov("lev").Conn("kon", godip.Land).Conn("egy", godip.Land).Conn("mec", godip.Land).Conn("ara", godip.Land).Conn("bag", godip.Land).Conn("arm", godip.Land).Flag(godip.Land).
		// Levant (North Coast)
		Prov("lev/nc").Conn("kon", godip.Sea).Conn("ems", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Sea).
		// Levant (South Coast)
		Prov("lev/sc").Conn("egy", godip.Sea).Conn("red", godip.Sea).Conn("mec", godip.Sea).Flag(godip.Sea).
		// Guangzhou
		Prov("gua").Conn("sha", godip.Coast...).Conn("wuh", godip.Land).Conn("yun", godip.Land).Conn("vit", godip.Coast...).Conn("goo", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Coast...).SC(China).
		// Yunnan
		Prov("yun").Conn("bum", godip.Land).Conn("lao", godip.Land).Conn("vit", godip.Land).Conn("gua", godip.Land).Conn("wuh", godip.Land).Conn("qin", godip.Land).Conn("tib", godip.Land).Flag(godip.Land).
		// Rome
		Prov("rom").Conn("ven", godip.Land).Conn("mil", godip.Land).Conn("pie", godip.Coast...).Conn("gol", godip.Sea).Conn("tyh", godip.Sea).Conn("nap", godip.Coast...).Conn("apu", godip.Land).Flag(godip.Coast...).SC(Italy).
		// Brest
		Prov("bre").Conn("par", godip.Land).Conn("pic", godip.Coast...).Conn("eng", godip.Sea).Conn("mid", godip.Sea).Conn("gas", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Tyrol
		Prov("tyo").Conn("boh", godip.Land).Conn("mun", godip.Land).Conn("swi", godip.Land).Conn("mil", godip.Land).Conn("ven", godip.Land).Conn("tes", godip.Land).Conn("vnn", godip.Land).Flag(godip.Land).
		// Paris
		Prov("par").Conn("bre", godip.Land).Conn("gas", godip.Land).Conn("bug", godip.Land).Conn("pic", godip.Land).Flag(godip.Land).SC(France).
		// Korea
		Prov("kor").Conn("yel", godip.Sea).Conn("ecs", godip.Sea).Conn("soj", godip.Sea).Conn("vla", godip.Coast...).Conn("man", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Ionian Sea
		Prov("ion").Conn("trp", godip.Sea).Conn("cyr", godip.Sea).Conn("ems", godip.Sea).Conn("aeg", godip.Sea).Conn("gre", godip.Sea).Conn("alb", godip.Sea).Conn("adr", godip.Sea).Conn("apu", godip.Sea).Conn("nap", godip.Sea).Conn("tyh", godip.Sea).Conn("tun", godip.Sea).Flag(godip.Sea).
		// South Pacific Ocean
		Prov("spo").Conn("npo", godip.Sea).Conn("tok", godip.Sea).Conn("shi", godip.Sea).Conn("osa", godip.Sea).Conn("ecs", godip.Sea).Conn("phi", godip.Sea).Conn("cel", godip.Sea).Conn("tim", godip.Sea).Conn("bxe", godip.Sea).Conn("bxe", godip.Sea).Conn("bxe", godip.Sea).Flag(godip.Sea).
		// Eastern Indian Ocean
		Prov("eio").Conn("tim", godip.Sea).Conn("jav", godip.Sea).Conn("jvs", godip.Sea).Conn("sum", godip.Sea).Conn("and", godip.Sea).Conn("bay", godip.Sea).Conn("mad", godip.Sea).Conn("wio", godip.Sea).Conn("cey", godip.Sea).Conn("cey", godip.Sea).Conn("wio", godip.Sea).Conn("bxh", godip.Sea).Conn("bxh", godip.Sea).Conn("bxh", godip.Sea).Flag(godip.Sea).
		// Wuhan
		Prov("wuh").Conn("sha", godip.Land).Conn("heb", godip.Land).Conn("qin", godip.Land).Conn("yun", godip.Land).Conn("gua", godip.Land).Flag(godip.Land).SC(China).
		// Portugal
		Prov("por").Conn("spa", godip.Land).Conn("spa/nc", godip.Sea).Conn("spa/sc", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Akita
		Prov("aki").Conn("tok", godip.Coast...).Conn("npo", godip.Sea).Conn("soj", godip.Sea).Conn("kyo", godip.Coast...).Flag(godip.Coast...).
		// Sumatra
		Prov("sum").Conn("eio", godip.Sea).Conn("jvs", godip.Sea).Conn("and", godip.Sea).Flag(godip.Coast...).
		// Tibet
		Prov("tib").Conn("afg", godip.Land).Conn("kas", godip.Land).Conn("del", godip.Land).Conn("nep", godip.Land).Conn("cal", godip.Land).Conn("bum", godip.Land).Conn("yun", godip.Land).Conn("qin", godip.Land).Conn("xin", godip.Land).Conn("tur", godip.Land).Flag(godip.Land).
		// Baghdad
		Prov("bag").Conn("ara", godip.Land).Conn("ara/nc", godip.Sea).Conn("psg", godip.Sea).Conn("per", godip.Coast...).Conn("arm", godip.Land).Conn("lev", godip.Land).Flag(godip.Coast...).SC(Turkey).
		// Switzerland
		Prov("swi").Conn("swa", godip.Land).Conn("bug", godip.Land).Conn("mar", godip.Land).Conn("pie", godip.Land).Conn("mil", godip.Land).Conn("tyo", godip.Land).Conn("mun", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Gulf of Lyons
		Prov("gol").Conn("mar", godip.Sea).Conn("spa/sc", godip.Sea).Conn("spa", godip.Sea).Conn("wms", godip.Sea).Conn("tyh", godip.Sea).Conn("rom", godip.Sea).Conn("pie", godip.Sea).Flag(godip.Sea).
		// Skagerrak
		Prov("ska").Conn("swe", godip.Sea).Conn("nay", godip.Sea).Conn("not", godip.Sea).Conn("den", godip.Sea).Flag(godip.Sea).
		// Western Indian Ocean
		Prov("wio").Conn("eio", godip.Sea).Conn("cey", godip.Sea).Conn("eio", godip.Sea).Conn("mad", godip.Sea).Conn("bom", godip.Sea).Conn("ars", godip.Sea).Conn("goa", godip.Sea).Conn("hor", godip.Sea).Conn("bxg", godip.Sea).Conn("bxg", godip.Sea).Conn("bxg", godip.Sea).Flag(godip.Sea).
		// Mogadishu
		Prov("mog").Conn("ken", godip.Coast...).Conn("hor", godip.Sea).Conn("goa", godip.Sea).Conn("awd", godip.Coast...).Conn("eth", godip.Land).Flag(godip.Coast...).SC(Italy).
		// Box A bcd
		Prov("bxa").Conn("nao", godip.Sea).Conn("nao", godip.Sea).Conn("bxb", godip.Sea).Conn("bxc", godip.Sea).Conn("bxd", godip.Sea).Flag(godip.Sea).
		// Saxony
		Prov("sax").Conn("mun", godip.Land).Conn("boh", godip.Land).Conn("sil", godip.Land).Conn("ber", godip.Land).Conn("kie", godip.Land).Conn("col", godip.Land).Flag(godip.Land).
		// Ethiopia
		Prov("eth").Conn("goa", godip.Sea).Conn("red", godip.Sea).Conn("sud", godip.Coast...).Conn("ken", godip.Land).Conn("mog", godip.Land).Conn("awd", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Tokyo
		Prov("tok").Conn("spo", godip.Sea).Conn("npo", godip.Sea).Conn("aki", godip.Coast...).Conn("kyo", godip.Land).Conn("shi", godip.Coast...).Flag(godip.Coast...).SC(Japan).
		// Peking
		Prov("pek").Conn("yel", godip.Sea).Conn("man", godip.Coast...).Conn("inn", godip.Land).Conn("qin", godip.Land).Conn("heb", godip.Land).Conn("heb/nc", godip.Sea).Flag(godip.Coast...).SC(China).
		// Arabia
		Prov("ara").Conn("bag", godip.Land).Conn("lev", godip.Land).Conn("mec", godip.Land).Conn("ade", godip.Land).Conn("yem", godip.Land).Conn("oma", godip.Land).Flag(godip.Land).
		// Arabia (North Coast)
		Prov("ara/nc").Conn("bag", godip.Sea).Conn("oma", godip.Sea).Conn("psg", godip.Sea).Flag(godip.Sea).
		// Arabia (South Coast)
		Prov("ara/sc").Conn("mec", godip.Sea).Conn("red", godip.Sea).Conn("ade", godip.Sea).Flag(godip.Sea).
		// Barents Sea
		Prov("bar").Conn("noi", godip.Sea).Conn("nay", godip.Sea).Conn("stp", godip.Sea).Conn("stp/nc", godip.Sea).Flag(godip.Sea).
		// North Sea
		Prov("not").Conn("lon", godip.Sea).Conn("eng", godip.Sea).Conn("bel", godip.Sea).Conn("hol", godip.Sea).Conn("hel", godip.Sea).Conn("den", godip.Sea).Conn("ska", godip.Sea).Conn("nay", godip.Sea).Conn("noi", godip.Sea).Conn("edi", godip.Sea).Conn("yor", godip.Sea).Flag(godip.Sea).
		// Inner Mongolia
		Prov("inn").Conn("man", godip.Land).Conn("mon", godip.Land).Conn("xin", godip.Land).Conn("qin", godip.Land).Conn("pek", godip.Land).Flag(godip.Land).
		// Kamchatka
		Prov("kam").Conn("npo", godip.Sea).Conn("sib", godip.Land).Conn("vla", godip.Coast...).Conn("soo", godip.Sea).Flag(godip.Coast...).
		// South Atlantic Ocean
		Prov("sao").Conn("mal", godip.Sea).Conn("mor", godip.Sea).Conn("mid", godip.Sea).Conn("bxc", godip.Sea).Conn("bxc", godip.Sea).Flag(godip.Sea).
		// Deccan
		Prov("dec").Conn("mad", godip.Land).Conn("cal", godip.Land).Conn("del", godip.Land).Conn("bom", godip.Land).Flag(godip.Land).
		// Karafuto
		Prov("kar").Conn("sak", godip.Coast...).Conn("soj", godip.Sea).Conn("npo", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Algeria
		Prov("alg").Conn("wms", godip.Sea).Conn("mor", godip.Coast...).Conn("sah", godip.Land).Conn("tun", godip.Coast...).Conn("tyh", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Awdal
		Prov("awd").Conn("eth", godip.Coast...).Conn("mog", godip.Coast...).Conn("goa", godip.Sea).Flag(godip.Coast...).
		// Baltic Sea
		Prov("bal").Conn("pru", godip.Sea).Conn("lia", godip.Sea).Conn("gob", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Conn("kie", godip.Sea).Conn("ber", godip.Sea).Flag(godip.Sea).
		// Calcutta
		Prov("cal").Conn("bay", godip.Sea).Conn("bum", godip.Coast...).Conn("tib", godip.Land).Conn("nep", godip.Land).Conn("del", godip.Land).Conn("dec", godip.Land).Conn("mad", godip.Coast...).Flag(godip.Coast...).SC(India).
		// Box D aef
		Prov("bxd").Conn("npo", godip.Sea).Conn("npo", godip.Sea).Conn("bxa", godip.Sea).Conn("bxe", godip.Sea).Conn("bxf", godip.Sea).Flag(godip.Sea).
		// Edinburgh
		Prov("edi").Conn("cly", godip.Coast...).Conn("lie", godip.Land).Conn("yor", godip.Coast...).Conn("not", godip.Sea).Conn("noi", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Piedmont
		Prov("pie").Conn("swi", godip.Land).Conn("mar", godip.Coast...).Conn("gol", godip.Sea).Conn("rom", godip.Coast...).Conn("mil", godip.Land).Flag(godip.Coast...).
		// Budapest
		Prov("bud").Conn("gal", godip.Land).Conn("vnn", godip.Land).Conn("tes", godip.Land).Conn("sar", godip.Land).Conn("ser", godip.Land).Conn("rum", godip.Land).Flag(godip.Land).SC(Austria).
		// Vladivostok
		Prov("vla").Conn("soj", godip.Sea).Conn("soo", godip.Sea).Conn("kam", godip.Coast...).Conn("sib", godip.Land).Conn("man", godip.Land).Conn("kor", godip.Coast...).Flag(godip.Coast...).SC(Russia).
		// Kyoto
		Prov("kyo").Conn("shi", godip.Land).Conn("tok", godip.Land).Conn("aki", godip.Coast...).Conn("soj", godip.Sea).Conn("kag", godip.Coast...).Conn("osa", godip.Land).Flag(godip.Coast...).SC(Japan).
		// Bulgaria
		Prov("bul").Conn("mac", godip.Land).Conn("con", godip.Coast...).Conn("bla", godip.Sea).Conn("rum", godip.Coast...).Conn("ser", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Horn of Africa
		Prov("hor").Conn("wio", godip.Sea).Conn("goa", godip.Sea).Conn("mog", godip.Sea).Conn("ken", godip.Sea).Flag(godip.Sea).
		// Delhi
		Prov("del").Conn("sid", godip.Land).Conn("bom", godip.Land).Conn("dec", godip.Land).Conn("cal", godip.Land).Conn("nep", godip.Land).Conn("tib", godip.Land).Conn("kas", godip.Land).Flag(godip.Land).SC(India).
		// Java Sea
		Prov("jvs").Conn("cel", godip.Sea).Conn("bor", godip.Sea).Conn("guh", godip.Sea).Conn("sig", godip.Sea).Conn("and", godip.Sea).Conn("sum", godip.Sea).Conn("eio", godip.Sea).Conn("jav", godip.Sea).Flag(godip.Sea).
		// Morocco
		Prov("mor").Conn("sao", godip.Sea).Conn("mal", godip.Coast...).Conn("sah", godip.Land).Conn("alg", godip.Coast...).Conn("wms", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Serbia
		Prov("ser").Conn("sar", godip.Land).Conn("alb", godip.Land).Conn("mac", godip.Land).Conn("bul", godip.Land).Conn("rum", godip.Land).Conn("bud", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Borneo
		Prov("bor").Conn("jvs", godip.Sea).Conn("cel", godip.Sea).Conn("scs", godip.Sea).Conn("guh", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Ceylon
		Prov("cey").Conn("eio", godip.Sea).Conn("eio", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Belgium
		Prov("bel").Conn("pic", godip.Coast...).Conn("bug", godip.Land).Conn("col", godip.Land).Conn("hol", godip.Coast...).Conn("not", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Kagoshima
		Prov("kag").Conn("osa", godip.Coast...).Conn("kyo", godip.Coast...).Conn("soj", godip.Sea).Conn("ecs", godip.Sea).Flag(godip.Coast...).
		// Ireland
		Prov("ire").Conn("nao", godip.Sea).Conn("mid", godip.Sea).Conn("iri", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Western Mediterranean Sea
		Prov("wms").Conn("alg", godip.Sea).Conn("tyh", godip.Sea).Conn("gol", godip.Sea).Conn("spa/sc", godip.Sea).Conn("spa", godip.Sea).Conn("mid", godip.Sea).Conn("mor", godip.Sea).Flag(godip.Sea).
		// Mali
		Prov("mal").Conn("sah", godip.Land).Conn("mor", godip.Coast...).Conn("sao", godip.Sea).Flag(godip.Coast...).
		// Liverpool
		Prov("lie").Conn("iri", godip.Sea).Conn("wal", godip.Coast...).Conn("yor", godip.Land).Conn("edi", godip.Land).Conn("cly", godip.Coast...).Conn("nao", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Sapporo
		Prov("sap").Conn("npo", godip.Sea).Conn("soj", godip.Sea).Conn("npo", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Sudan
		Prov("sud").Conn("ken", godip.Land).Conn("eth", godip.Coast...).Conn("red", godip.Sea).Conn("egy", godip.Coast...).Conn("fez", godip.Land).Conn("cen", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Eastern Mediterranean Sea
		Prov("ems").Conn("kon", godip.Sea).Conn("aeg", godip.Sea).Conn("ion", godip.Sea).Conn("cyr", godip.Sea).Conn("egy", godip.Sea).Conn("lev", godip.Sea).Conn("lev/nc", godip.Sea).Flag(godip.Sea).
		// Denmark
		Prov("den").Conn("not", godip.Sea).Conn("hel", godip.Sea).Conn("kie", godip.Coast...).Conn("bal", godip.Sea).Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Aegean Sea
		Prov("aeg").Conn("gre", godip.Sea).Conn("ion", godip.Sea).Conn("ems", godip.Sea).Conn("kon", godip.Sea).Conn("con", godip.Sea).Conn("mac", godip.Sea).Flag(godip.Sea).
		// Adriatic Sea
		Prov("adr").Conn("tes", godip.Sea).Conn("ven", godip.Sea).Conn("apu", godip.Sea).Conn("ion", godip.Sea).Conn("alb", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// Sweden
		Prov("swe").Conn("bal", godip.Sea).Conn("gob", godip.Sea).Conn("fin", godip.Coast...).Conn("nay", godip.Coast...).Conn("ska", godip.Sea).Conn("den", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Heligoland Bight
		Prov("hel").Conn("den", godip.Sea).Conn("not", godip.Sea).Conn("hol", godip.Sea).Conn("kie", godip.Sea).Flag(godip.Sea).
		// English Channel
		Prov("eng").Conn("wal", godip.Sea).Conn("iri", godip.Sea).Conn("mid", godip.Sea).Conn("bre", godip.Sea).Conn("pic", godip.Sea).Conn("bel", godip.Sea).Conn("not", godip.Sea).Conn("lon", godip.Sea).Flag(godip.Sea).
		// Caucasus
		Prov("cau").Conn("sev", godip.Land).Conn("arm", godip.Land).Conn("per", godip.Land).Conn("oms", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).
		// Armenia
		Prov("arm").Conn("ank", godip.Coast...).Conn("kon", godip.Land).Conn("lev", godip.Land).Conn("bag", godip.Land).Conn("per", godip.Land).Conn("cau", godip.Land).Conn("sev", godip.Coast...).Conn("bla", godip.Sea).Flag(godip.Coast...).
		// Fez
		Prov("fez").Conn("egy", godip.Land).Conn("cyr", godip.Land).Conn("trp", godip.Land).Conn("sah", godip.Land).Conn("cen", godip.Land).Conn("sud", godip.Land).Flag(godip.Land).
		// Norway
		Prov("nay").Conn("swe", godip.Coast...).Conn("fin", godip.Land).Conn("stp", godip.Land).Conn("stp/nc", godip.Sea).Conn("bar", godip.Sea).Conn("noi", godip.Sea).Conn("not", godip.Sea).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Java
		Prov("jav").Conn("jvs", godip.Sea).Conn("eio", godip.Sea).Conn("tim", godip.Sea).Conn("cel", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Holland
		Prov("hol").Conn("hel", godip.Sea).Conn("not", godip.Sea).Conn("bel", godip.Coast...).Conn("col", godip.Land).Conn("kie", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Central Africa
		Prov("cen").Conn("sud", godip.Land).Conn("fez", godip.Land).Conn("sah", godip.Land).Flag(godip.Land).
		// Arabian Sea
		Prov("ars").Conn("yem", godip.Sea).Conn("goa", godip.Sea).Conn("wio", godip.Sea).Conn("bom", godip.Sea).Conn("sid", godip.Sea).Conn("per", godip.Sea).Conn("psg", godip.Sea).Conn("oma", godip.Sea).Flag(godip.Sea).
		// Tyrrhenian Sea
		Prov("tyh").Conn("wms", godip.Sea).Conn("alg", godip.Sea).Conn("tun", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Sea).Conn("rom", godip.Sea).Conn("gol", godip.Sea).Flag(godip.Sea).
		// Thailand
		Prov("tha").Conn("cam", godip.Land).Conn("lao", godip.Land).Conn("bum", godip.Land).Conn("sig", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Thailand (East Coast)
		Prov("tha/ec").Conn("cam", godip.Sea).Conn("sig", godip.Sea).Conn("guh", godip.Sea).Flag(godip.Sea).
		// Thailand (West Coast)
		Prov("tha/wc").Conn("bum", godip.Sea).Conn("and", godip.Sea).Conn("sig", godip.Sea).Flag(godip.Sea).
		// Cyrene
		Prov("cyr").Conn("fez", godip.Land).Conn("egy", godip.Coast...).Conn("ems", godip.Sea).Conn("ion", godip.Sea).Conn("trp", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Bothnia
		Prov("gob").Conn("bal", godip.Sea).Conn("lia", godip.Sea).Conn("stp", godip.Sea).Conn("stp/sc", godip.Sea).Conn("fin", godip.Sea).Conn("swe", godip.Sea).Flag(godip.Sea).
		// Gascony
		Prov("gas").Conn("mar", godip.Land).Conn("par", godip.Land).Conn("bre", godip.Coast...).Conn("mid", godip.Sea).Conn("spa", godip.Land).Conn("spa/nc", godip.Sea).Conn("bug", godip.Land).Flag(godip.Coast...).
		// Celebes Sea
		Prov("cel").Conn("jvs", godip.Sea).Conn("jav", godip.Sea).Conn("tim", godip.Sea).Conn("spo", godip.Sea).Conn("phi", godip.Sea).Conn("scs", godip.Sea).Conn("bor", godip.Sea).Flag(godip.Sea).
		// Timor Sea
		Prov("tim").Conn("spo", godip.Sea).Conn("cel", godip.Sea).Conn("jav", godip.Sea).Conn("eio", godip.Sea).Conn("bxf", godip.Sea).Conn("bxf", godip.Sea).Flag(godip.Sea).
		// Sea of Okhotsk
		Prov("soo").Conn("sak", godip.Sea).Conn("kar", godip.Sea).Conn("npo", godip.Sea).Conn("kam", godip.Sea).Conn("vla", godip.Sea).Conn("soj", godip.Sea).Flag(godip.Sea).
		// Philippines
		Prov("phi").Conn("ecs", godip.Sea).Conn("scs", godip.Sea).Conn("cel", godip.Sea).Conn("spo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cologne
		Prov("col").Conn("bel", godip.Land).Conn("bug", godip.Land).Conn("swa", godip.Land).Conn("mun", godip.Land).Conn("sax", godip.Land).Conn("kie", godip.Land).Conn("hol", godip.Land).Flag(godip.Land).SC(Germany).
		// Annam
		Prov("ann").Conn("cam", godip.Land).Conn("sai", godip.Coast...).Conn("scs", godip.Sea).Conn("goo", godip.Sea).Conn("vit", godip.Coast...).Conn("lao", godip.Land).Flag(godip.Coast...).
		// Manchuria
		Prov("man").Conn("yel", godip.Sea).Conn("kor", godip.Coast...).Conn("vla", godip.Land).Conn("sib", godip.Land).Conn("mon", godip.Land).Conn("inn", godip.Land).Conn("pek", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Constantinople
		Prov("con").Conn("mac", godip.Coast...).Conn("aeg", godip.Sea).Conn("kon", godip.Coast...).Conn("ank", godip.Coast...).Conn("bla", godip.Sea).Conn("bul", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Bay of Bengal
		Prov("bay").Conn("mad", godip.Sea).Conn("eio", godip.Sea).Conn("and", godip.Sea).Conn("bum", godip.Sea).Conn("cal", godip.Sea).Flag(godip.Sea).
		// Marseilles
		Prov("mar").Conn("gol", godip.Sea).Conn("pie", godip.Coast...).Conn("swi", godip.Land).Conn("bug", godip.Land).Conn("gas", godip.Land).Conn("spa", godip.Land).Conn("spa/sc", godip.Sea).Flag(godip.Coast...).SC(France).
		// York
		Prov("yor").Conn("not", godip.Sea).Conn("edi", godip.Coast...).Conn("lie", godip.Land).Conn("wal", godip.Land).Conn("lon", godip.Coast...).Flag(godip.Coast...).
		// Ukraine
		Prov("ukr").Conn("rum", godip.Land).Conn("sev", godip.Land).Conn("mos", godip.Land).Conn("war", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).
		// Mid Atlantic Ocean
		Prov("mid").Conn("sao", godip.Sea).Conn("mor", godip.Sea).Conn("wms", godip.Sea).Conn("spa/sc", godip.Sea).Conn("spa/nc", godip.Sea).Conn("spa", godip.Sea).Conn("por", godip.Sea).Conn("gas", godip.Sea).Conn("bre", godip.Sea).Conn("eng", godip.Sea).Conn("iri", godip.Sea).Conn("ire", godip.Sea).Conn("nao", godip.Sea).Conn("bxb", godip.Sea).Conn("bxb", godip.Sea).Conn("bxb", godip.Sea).Flag(godip.Sea).
		// Saigon
		Prov("sai").Conn("scs", godip.Sea).Conn("ann", godip.Coast...).Conn("cam", godip.Coast...).Conn("guh", godip.Sea).Flag(godip.Coast...).SC(France).
		// Gulf of Tonkin
		Prov("goo").Conn("ann", godip.Sea).Conn("scs", godip.Sea).Conn("gua", godip.Sea).Conn("vit", godip.Sea).Flag(godip.Sea).
		// Qinghai
		Prov("qin").Conn("wuh", godip.Land).Conn("heb", godip.Land).Conn("pek", godip.Land).Conn("inn", godip.Land).Conn("xin", godip.Land).Conn("tib", godip.Land).Conn("yun", godip.Land).Flag(godip.Land).
		// Cambodia
		Prov("cam").Conn("ann", godip.Land).Conn("lao", godip.Land).Conn("tha", godip.Land).Conn("tha/ec", godip.Sea).Conn("guh", godip.Sea).Conn("sai", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// East China Sea
		Prov("ecs").Conn("phi", godip.Sea).Conn("spo", godip.Sea).Conn("osa", godip.Sea).Conn("kag", godip.Sea).Conn("soj", godip.Sea).Conn("kor", godip.Sea).Conn("yel", godip.Sea).Conn("tsi", godip.Sea).Conn("heb", godip.Sea).Conn("heb/sc", godip.Sea).Conn("sha", godip.Sea).Conn("scs", godip.Sea).Conn("for", godip.Sea).Conn("for", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Sea).
		// North Atlantic Ocean
		Prov("nao").Conn("bxa", godip.Sea).Conn("bxa", godip.Sea).Conn("mid", godip.Sea).Conn("ire", godip.Sea).Conn("iri", godip.Sea).Conn("lie", godip.Sea).Conn("cly", godip.Sea).Conn("noi", godip.Sea).Flag(godip.Sea).
		// Swabia
		Prov("swa").Conn("swi", godip.Land).Conn("mun", godip.Land).Conn("col", godip.Land).Conn("bug", godip.Land).Flag(godip.Land).
		// Kenya
		Prov("ken").Conn("hor", godip.Sea).Conn("mog", godip.Coast...).Conn("eth", godip.Land).Conn("sud", godip.Land).Flag(godip.Coast...).
		// Sea of Japan
		Prov("soj").Conn("kor", godip.Sea).Conn("ecs", godip.Sea).Conn("kag", godip.Sea).Conn("kyo", godip.Sea).Conn("aki", godip.Sea).Conn("npo", godip.Sea).Conn("sap", godip.Sea).Conn("npo", godip.Sea).Conn("kar", godip.Sea).Conn("sak", godip.Sea).Conn("soo", godip.Sea).Conn("vla", godip.Sea).Flag(godip.Sea).
		// Sarajevo
		Prov("sar").Conn("alb", godip.Coast...).Conn("ser", godip.Land).Conn("bud", godip.Land).Conn("tes", godip.Coast...).Conn("adr", godip.Sea).Flag(godip.Coast...).SC(Austria).
		// Konya
		Prov("kon").Conn("ank", godip.Land).Conn("con", godip.Coast...).Conn("aeg", godip.Sea).Conn("ems", godip.Sea).Conn("lev", godip.Land).Conn("lev/nc", godip.Sea).Conn("arm", godip.Land).Flag(godip.Coast...).
		// Aden
		Prov("ade").Conn("goa", godip.Sea).Conn("yem", godip.Coast...).Conn("ara", godip.Land).Conn("ara/sc", godip.Sea).Conn("red", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Sindh
		Prov("sid").Conn("del", godip.Land).Conn("kas", godip.Land).Conn("afg", godip.Land).Conn("per", godip.Coast...).Conn("ars", godip.Sea).Conn("bom", godip.Coast...).Flag(godip.Coast...).
		// Spain
		Prov("spa").Conn("mar", godip.Land).Conn("gas", godip.Land).Conn("por", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Spain (North Coast)
		Prov("spa/nc").Conn("gas", godip.Sea).Conn("mid", godip.Sea).Conn("por", godip.Sea).Flag(godip.Sea).
		// Spain (South Coast)
		Prov("spa/sc").Conn("wms", godip.Sea).Conn("mar", godip.Sea).Conn("mid", godip.Sea).Conn("gol", godip.Sea).Conn("por", godip.Sea).Flag(godip.Sea).
		// Warsaw
		Prov("war").Conn("mos", godip.Land).Conn("lia", godip.Land).Conn("pru", godip.Land).Conn("sil", godip.Land).Conn("gal", godip.Land).Conn("ukr", godip.Land).Flag(godip.Land).SC(Russia).
		// Norwegian Sea
		Prov("noi").Conn("nao", godip.Sea).Conn("cly", godip.Sea).Conn("edi", godip.Sea).Conn("not", godip.Sea).Conn("nay", godip.Sea).Conn("bar", godip.Sea).Flag(godip.Sea).
		// Singapore
		Prov("sig").Conn("jvs", godip.Sea).Conn("guh", godip.Sea).Conn("tha", godip.Land).Conn("tha/ec", godip.Sea).Conn("tha/wc", godip.Sea).Conn("and", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Gulf of Aden
		Prov("goa").Conn("eth", godip.Sea).Conn("awd", godip.Sea).Conn("mog", godip.Sea).Conn("hor", godip.Sea).Conn("wio", godip.Sea).Conn("ars", godip.Sea).Conn("yem", godip.Sea).Conn("ade", godip.Sea).Conn("red", godip.Sea).Flag(godip.Sea).
		// Mongolia
		Prov("mon").Conn("sib", godip.Land).Conn("xin", godip.Land).Conn("inn", godip.Land).Conn("man", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wales
		Prov("wal").Conn("eng", godip.Sea).Conn("lon", godip.Coast...).Conn("yor", godip.Land).Conn("lie", godip.Coast...).Conn("iri", godip.Sea).Flag(godip.Coast...).
		// Yellow Sea
		Prov("yel").Conn("kor", godip.Sea).Conn("man", godip.Sea).Conn("pek", godip.Sea).Conn("heb", godip.Sea).Conn("heb/nc", godip.Sea).Conn("tsi", godip.Sea).Conn("ecs", godip.Sea).Flag(godip.Sea).
		// Greece
		Prov("gre").Conn("aeg", godip.Sea).Conn("mac", godip.Coast...).Conn("alb", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Venice
		Prov("ven").Conn("rom", godip.Land).Conn("apu", godip.Coast...).Conn("adr", godip.Sea).Conn("tes", godip.Coast...).Conn("tyo", godip.Land).Conn("mil", godip.Land).Flag(godip.Coast...).
		// Vienna
		Prov("vnn").Conn("boh", godip.Land).Conn("tyo", godip.Land).Conn("tes", godip.Land).Conn("bud", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).SC(Austria).
		// Andaman Sea
		Prov("and").Conn("sum", godip.Sea).Conn("jvs", godip.Sea).Conn("sig", godip.Sea).Conn("tha", godip.Sea).Conn("tha/wc", godip.Sea).Conn("bum", godip.Sea).Conn("bay", godip.Sea).Conn("eio", godip.Sea).Flag(godip.Sea).
		// Nepal
		Prov("nep").Conn("tib", godip.Land).Conn("del", godip.Land).Conn("cal", godip.Land).Flag(godip.Land).
		// Box E bdf
		Prov("bxe").Conn("spo", godip.Sea).Conn("spo", godip.Sea).Conn("spo", godip.Sea).Conn("bxb", godip.Sea).Conn("bxd", godip.Sea).Conn("bxf", godip.Sea).Flag(godip.Sea).
		// Box F cdegh
		Prov("bxf").Conn("tim", godip.Sea).Conn("tim", godip.Sea).Conn("bxc", godip.Sea).Conn("bxd", godip.Sea).Conn("bxe", godip.Sea).Conn("bxg", godip.Sea).Conn("bxh", godip.Sea).Flag(godip.Sea).
		// Bohemia
		Prov("boh").Conn("vnn", godip.Land).Conn("gal", godip.Land).Conn("sil", godip.Land).Conn("sax", godip.Land).Conn("mun", godip.Land).Conn("tyo", godip.Land).Flag(godip.Land).
		// Laos
		Prov("lao").Conn("bum", godip.Land).Conn("tha", godip.Land).Conn("cam", godip.Land).Conn("ann", godip.Land).Conn("vit", godip.Land).Conn("yun", godip.Land).Flag(godip.Land).
		// Rumania
		Prov("rum").Conn("bul", godip.Coast...).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Conn("ukr", godip.Land).Conn("gal", godip.Land).Conn("bud", godip.Land).Conn("ser", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Shanghai
		Prov("sha").Conn("wuh", godip.Land).Conn("gua", godip.Coast...).Conn("scs", godip.Sea).Conn("ecs", godip.Sea).Conn("heb", godip.Land).Conn("heb/sc", godip.Sea).Flag(godip.Coast...).SC(China).
		// Milan
		Prov("mil").Conn("ven", godip.Land).Conn("tyo", godip.Land).Conn("swi", godip.Land).Conn("pie", godip.Land).Conn("rom", godip.Land).Flag(godip.Land).SC(Italy).
		// Macedonia
		Prov("mac").Conn("bul", godip.Land).Conn("ser", godip.Land).Conn("alb", godip.Land).Conn("gre", godip.Coast...).Conn("aeg", godip.Sea).Conn("con", godip.Coast...).Flag(godip.Coast...).
		// Black Sea
		Prov("bla").Conn("bul", godip.Sea).Conn("con", godip.Sea).Conn("ank", godip.Sea).Conn("arm", godip.Sea).Conn("sev", godip.Sea).Conn("rum", godip.Sea).Flag(godip.Sea).
		// Omsk
		Prov("oms").Conn("tur", godip.Land).Conn("sib", godip.Land).Conn("stp", godip.Land).Conn("mos", godip.Land).Conn("cau", godip.Land).Flag(godip.Land).SC(Russia).
		// Xinjiang
		Prov("xin").Conn("qin", godip.Land).Conn("inn", godip.Land).Conn("mon", godip.Land).Conn("sib", godip.Land).Conn("tur", godip.Land).Conn("tib", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Egypt
		Prov("egy").Conn("lev", godip.Land).Conn("lev/sc", godip.Sea).Conn("lev/nc", godip.Sea).Conn("ems", godip.Sea).Conn("cyr", godip.Coast...).Conn("fez", godip.Land).Conn("sud", godip.Coast...).Conn("red", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Formosa
		Prov("for").Conn("ecs", godip.Sea).Conn("ecs", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Box G cfh
		Prov("bxg").Conn("wio", godip.Sea).Conn("wio", godip.Sea).Conn("wio", godip.Sea).Conn("bxc", godip.Sea).Conn("bxf", godip.Sea).Conn("bxh", godip.Sea).Flag(godip.Sea).
		// Osaka
		Prov("osa").Conn("kag", godip.Coast...).Conn("ecs", godip.Sea).Conn("spo", godip.Sea).Conn("shi", godip.Coast...).Conn("kyo", godip.Land).Flag(godip.Coast...).SC(Japan).
		// Turkmenistan
		Prov("tur").Conn("sib", godip.Land).Conn("oms", godip.Land).Conn("per", godip.Land).Conn("afg", godip.Land).Conn("tib", godip.Land).Conn("xin", godip.Land).Flag(godip.Land).
		// Tsingtao
		Prov("tsi").Conn("ecs", godip.Sea).Conn("yel", godip.Sea).Conn("heb", godip.Land).Conn("heb/nc", godip.Sea).Conn("heb/sc", godip.Sea).Flag(godip.Coast...).SC(Germany).
		// Shizuoka
		Prov("shi").Conn("kyo", godip.Land).Conn("osa", godip.Coast...).Conn("spo", godip.Sea).Conn("tok", godip.Coast...).Flag(godip.Coast...).
		// Clyde
		Prov("cly").Conn("edi", godip.Coast...).Conn("noi", godip.Sea).Conn("nao", godip.Sea).Conn("lie", godip.Coast...).Flag(godip.Coast...).
		// Sahara
		Prov("sah").Conn("cen", godip.Land).Conn("fez", godip.Land).Conn("trp", godip.Land).Conn("tun", godip.Land).Conn("alg", godip.Land).Conn("mor", godip.Land).Conn("mal", godip.Land).Flag(godip.Land).
		// North Pacific Ocean
		Prov("npo").Conn("kam", godip.Sea).Conn("soo", godip.Sea).Conn("kar", godip.Sea).Conn("soj", godip.Sea).Conn("sap", godip.Sea).Conn("sap", godip.Sea).Conn("soj", godip.Sea).Conn("aki", godip.Sea).Conn("tok", godip.Sea).Conn("spo", godip.Sea).Conn("bxd", godip.Sea).Conn("bxd", godip.Sea).Flag(godip.Sea).
		// Apulia
		Prov("apu").Conn("nap", godip.Coast...).Conn("ion", godip.Sea).Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("rom", godip.Land).Flag(godip.Coast...).
		// Burma
		Prov("bum").Conn("yun", godip.Land).Conn("tib", godip.Land).Conn("cal", godip.Coast...).Conn("bay", godip.Sea).Conn("and", godip.Sea).Conn("tha", godip.Land).Conn("tha/wc", godip.Sea).Conn("lao", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Box B ace
		Prov("bxb").Conn("mid", godip.Sea).Conn("mid", godip.Sea).Conn("mid", godip.Sea).Conn("bxa", godip.Sea).Conn("bxc", godip.Sea).Conn("bxe", godip.Sea).Flag(godip.Sea).
		// Box C abfgh
		Prov("bxc").Conn("sao", godip.Sea).Conn("sao", godip.Sea).Conn("bxa", godip.Sea).Conn("bxb", godip.Sea).Conn("bxf", godip.Sea).Conn("bxg", godip.Sea).Conn("bxh", godip.Sea).Flag(godip.Sea).
		// Siberia
		Prov("sib").Conn("tur", godip.Land).Conn("xin", godip.Land).Conn("mon", godip.Land).Conn("man", godip.Land).Conn("vla", godip.Land).Conn("kam", godip.Land).Conn("oms", godip.Land).Flag(godip.Land).
		// Oman
		Prov("oma").Conn("ars", godip.Sea).Conn("psg", godip.Sea).Conn("ara", godip.Land).Conn("ara/nc", godip.Sea).Conn("yem", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Thailand
		Prov("guh").Conn("scs", godip.Sea).Conn("sai", godip.Sea).Conn("cam", godip.Sea).Conn("tha", godip.Sea).Conn("tha/ec", godip.Sea).Conn("sig", godip.Sea).Conn("jvs", godip.Sea).Conn("bor", godip.Sea).Flag(godip.Sea).
		// Irish Sea
		Prov("iri").Conn("nao", godip.Sea).Conn("ire", godip.Sea).Conn("mid", godip.Sea).Conn("eng", godip.Sea).Conn("wal", godip.Sea).Conn("lie", godip.Sea).Flag(godip.Sea).
		// Finland
		Prov("fin").Conn("stp", godip.Land).Conn("stp/sc", godip.Sea).Conn("nay", godip.Land).Conn("swe", godip.Coast...).Conn("gob", godip.Sea).Flag(godip.Coast...).
		// Prussia
		Prov("pru").Conn("sil", godip.Land).Conn("war", godip.Land).Conn("lia", godip.Coast...).Conn("bal", godip.Sea).Conn("ber", godip.Coast...).Flag(godip.Coast...).
		// Berlin
		Prov("ber").Conn("kie", godip.Coast...).Conn("sax", godip.Land).Conn("sil", godip.Land).Conn("pru", godip.Coast...).Conn("bal", godip.Sea).Flag(godip.Coast...).SC(Germany).
		// Persia
		Prov("per").Conn("tur", godip.Land).Conn("cau", godip.Land).Conn("arm", godip.Land).Conn("bag", godip.Coast...).Conn("psg", godip.Sea).Conn("ars", godip.Sea).Conn("sid", godip.Coast...).Conn("afg", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Livonia
		Prov("lia").Conn("bal", godip.Sea).Conn("pru", godip.Coast...).Conn("war", godip.Land).Conn("mos", godip.Land).Conn("stp", godip.Land).Conn("stp/sc", godip.Sea).Conn("gob", godip.Sea).Flag(godip.Coast...).
		// Burgundy
		Prov("bug").Conn("swi", godip.Land).Conn("swa", godip.Land).Conn("col", godip.Land).Conn("bel", godip.Land).Conn("pic", godip.Land).Conn("par", godip.Land).Conn("mar", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).
		// Mecca
		Prov("mec").Conn("lev", godip.Land).Conn("lev/sc", godip.Sea).Conn("red", godip.Sea).Conn("ara", godip.Land).Conn("ara/sc", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Persian Gulf
		Prov("psg").Conn("per", godip.Sea).Conn("bag", godip.Sea).Conn("ara", godip.Sea).Conn("ara/nc", godip.Sea).Conn("oma", godip.Sea).Conn("ars", godip.Sea).Flag(godip.Sea).
		// Naples
		Prov("nap").Conn("apu", godip.Coast...).Conn("rom", godip.Coast...).Conn("tyh", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Italy).
		// Tripolitania
		Prov("trp").Conn("ion", godip.Sea).Conn("tun", godip.Coast...).Conn("sah", godip.Land).Conn("fez", godip.Land).Conn("cyr", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Sakhalin
		Prov("sak").Conn("soo", godip.Sea).Conn("soj", godip.Sea).Conn("kar", godip.Coast...).Flag(godip.Coast...).
		// Kiel
		Prov("kie").Conn("ber", godip.Coast...).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Conn("hel", godip.Sea).Conn("hol", godip.Coast...).Conn("col", godip.Land).Conn("sax", godip.Land).Flag(godip.Coast...).SC(Germany).
		// Moscow
		Prov("mos").Conn("war", godip.Land).Conn("ukr", godip.Land).Conn("sev", godip.Land).Conn("cau", godip.Land).Conn("oms", godip.Land).Conn("stp", godip.Land).Conn("lia", godip.Land).Flag(godip.Land).SC(Russia).
		// Ankara
		Prov("ank").Conn("kon", godip.Land).Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Madras
		Prov("mad").Conn("bom", godip.Coast...).Conn("wio", godip.Sea).Conn("eio", godip.Sea).Conn("bay", godip.Sea).Conn("cal", godip.Coast...).Conn("dec", godip.Land).Flag(godip.Coast...).SC(India).
		// Trieste
		Prov("tes").Conn("adr", godip.Sea).Conn("sar", godip.Coast...).Conn("bud", godip.Land).Conn("vnn", godip.Land).Conn("tyo", godip.Land).Conn("ven", godip.Coast...).Flag(godip.Coast...).SC(Austria).
		// Vietnaam
		Prov("vit").Conn("lao", godip.Land).Conn("ann", godip.Coast...).Conn("goo", godip.Sea).Conn("gua", godip.Coast...).Conn("yun", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Picardy
		Prov("pic").Conn("bel", godip.Coast...).Conn("eng", godip.Sea).Conn("bre", godip.Coast...).Conn("par", godip.Land).Conn("bug", godip.Land).Flag(godip.Coast...).
		// Munich
		Prov("mun").Conn("sax", godip.Land).Conn("col", godip.Land).Conn("swa", godip.Land).Conn("swi", godip.Land).Conn("tyo", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).SC(Germany).
		Done()
}
