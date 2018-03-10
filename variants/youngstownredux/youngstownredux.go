package youngstownredux

import (
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

const (
	Turkey  dip.Nation = "Turkey"
	Austria dip.Nation = "Austria"
	Britain dip.Nation = "Britain"
	China   dip.Nation = "China"
	Japan   dip.Nation = "Japan"
	Italy   dip.Nation = "Italy"
	Germany dip.Nation = "Germany"
	India   dip.Nation = "India"
	Russia  dip.Nation = "Russia"
	France  dip.Nation = "France"
)

var Nations = []dip.Nation{Turkey, Austria, Britain, China, Japan, Italy, Germany, India, Russia, France}

var YoungstownReduxVariant = common.Variant{
	Name:        "Youngstown Redux",
	Graph:       func() dip.Graph { return YoungstownReduxGraph() },
	Start:       YoungstownReduxStart,
	Blank:       YoungstownReduxBlank,
	Phase:       classical.Phase,
	ParseOrders: orders.ParseAll,
	ParseOrder:  orders.Parse,
	OrderTypes:  orders.OrderTypes(),
	Nations:     Nations,
	PhaseTypes:  cla.PhaseTypes,
	Seasons:     cla.Seasons,
	UnitTypes:   cla.UnitTypes,
	SoloWinner:  common.SCCountWinner(28),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/youngstownreduxmap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "airborne",
	Version:     "I",
	Description: "A ten player variant that adds China, India and Japan to the standard seven nations.",
	Rules:       "Rules are as per classical Diplomacy. There are eight box sea regions which are each connected " +
		"to the other boxes in the same row and column, and allow fleets to travel 'around the world'. Six provinces " +
		"have two coasts (Spain, St. Petersburg, Levant, Arabia, Hebei and Thailand), and all other coastal regions have a " +
		"single coast. The winner is the first nation to 28 supply centers, or the player with the most in the case of " +
		"multiple nations reaching 28 in the same turn. If the leading two nations both have the same number of centers " +
		"then the game will continue for another year.  This variant is based on the Youngstown variant by Rod Walker, " + 
		"A. Phillips, Ken Lowe and Jon Monsarret.",
}

func YoungstownReduxBlank(phase dip.Phase) *state.State {
	return state.New(YoungstownReduxGraph(), phase, classical.BackupRule)
}

func YoungstownReduxStart() (result *state.State, err error) {
	startPhase := classical.Phase(1901, cla.Spring, cla.Movement)
	result = state.New(YoungstownReduxGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[dip.Province]dip.Unit{
		"ank":    dip.Unit{cla.Fleet, Turkey},
		"con":    dip.Unit{cla.Army, Turkey},
		"bag":    dip.Unit{cla.Army, Turkey},
		"mec":    dip.Unit{cla.Army, Turkey},
		"sar":    dip.Unit{cla.Fleet, Austria},
		"vnn":    dip.Unit{cla.Army, Austria},
		"bud":    dip.Unit{cla.Army, Austria},
		"tes":    dip.Unit{cla.Army, Austria},
		"lon":    dip.Unit{cla.Fleet, Britain},
		"lie":    dip.Unit{cla.Fleet, Britain},
		"edi":    dip.Unit{cla.Fleet, Britain},
		"ade":    dip.Unit{cla.Fleet, Britain},
		"sig":    dip.Unit{cla.Fleet, Britain},
		"sha":    dip.Unit{cla.Fleet, China},
		"pek":    dip.Unit{cla.Army, China},
		"gua":    dip.Unit{cla.Army, China},
		"wuh":    dip.Unit{cla.Army, China},
		"tok":    dip.Unit{cla.Fleet, Japan},
		"osa":    dip.Unit{cla.Fleet, Japan},
		"sap":    dip.Unit{cla.Fleet, Japan},
		"kyo":    dip.Unit{cla.Army, Japan},
		"nap":    dip.Unit{cla.Fleet, Italy},
		"mog":    dip.Unit{cla.Fleet, Italy},
		"rom":    dip.Unit{cla.Army, Italy},
		"mil":    dip.Unit{cla.Army, Italy},
		"tsi":    dip.Unit{cla.Fleet, Germany},
		"kie":    dip.Unit{cla.Fleet, Germany},
		"ber":    dip.Unit{cla.Army, Germany},
		"mun":    dip.Unit{cla.Army, Germany},
		"col":    dip.Unit{cla.Army, Germany},
		"bom":    dip.Unit{cla.Fleet, India},
		"mad":    dip.Unit{cla.Fleet, India},
		"del":    dip.Unit{cla.Army, India},
		"cal":    dip.Unit{cla.Army, India},
		"sev":    dip.Unit{cla.Fleet, Russia},
		"stp/sc": dip.Unit{cla.Fleet, Russia},
		"vla":    dip.Unit{cla.Fleet, Russia},
		"mos":    dip.Unit{cla.Army, Russia},
		"oms":    dip.Unit{cla.Army, Russia},
		"war":    dip.Unit{cla.Army, Russia},
		"bre":    dip.Unit{cla.Fleet, France},
		"sai":    dip.Unit{cla.Fleet, France},
		"par":    dip.Unit{cla.Army, France},
		"mar":    dip.Unit{cla.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
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
		Prov("bxh").Conn("eio", cla.Sea).Conn("eio", cla.Sea).Conn("eio", cla.Sea).Conn("bxc", cla.Sea).Conn("bxf", cla.Sea).Conn("bxg", cla.Sea).Flag(cla.Sea).
		// Tunisia
		Prov("tun").Conn("alg", cla.Coast...).Conn("sah", cla.Land).Conn("trp", cla.Coast...).Conn("ion", cla.Sea).Conn("tyh", cla.Sea).Flag(cla.Coast...).
		// Bombay
		Prov("bom").Conn("mad", cla.Coast...).Conn("dec", cla.Land).Conn("del", cla.Land).Conn("sid", cla.Coast...).Conn("ars", cla.Sea).Conn("wio", cla.Sea).Flag(cla.Coast...).SC(India).
		// Hebei
		Prov("heb").Conn("pek", cla.Land).Conn("qin", cla.Land).Conn("wuh", cla.Land).Conn("sha", cla.Land).Conn("tsi", cla.Land).Flag(cla.Land).
		// Hebei (North Coast)
		Prov("heb/nc").Conn("yel", cla.Sea).Conn("pek", cla.Sea).Conn("tsi", cla.Sea).Flag(cla.Sea).
		// Hebei (South Coast)
		Prov("heb/sc").Conn("sha", cla.Sea).Conn("ecs", cla.Sea).Conn("tsi", cla.Sea).Flag(cla.Sea).
		// Silesia
		Prov("sil").Conn("pru", cla.Land).Conn("ber", cla.Land).Conn("sax", cla.Land).Conn("boh", cla.Land).Conn("gal", cla.Land).Conn("war", cla.Land).Flag(cla.Land).
		// Sevastopol
		Prov("sev").Conn("cau", cla.Land).Conn("mos", cla.Land).Conn("ukr", cla.Land).Conn("rum", cla.Coast...).Conn("bla", cla.Sea).Conn("arm", cla.Coast...).Flag(cla.Coast...).SC(Russia).
		// Albania
		Prov("alb").Conn("ion", cla.Sea).Conn("gre", cla.Coast...).Conn("mac", cla.Land).Conn("ser", cla.Land).Conn("sar", cla.Coast...).Conn("adr", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// St. Petersburg
		Prov("stp").Conn("fin", cla.Land).Conn("lia", cla.Land).Conn("mos", cla.Land).Conn("oms", cla.Land).Conn("nay", cla.Land).Flag(cla.Land).SC(Russia).
		// St. Petersburg (North Coast)
		Prov("stp/nc").Conn("bar", cla.Sea).Conn("nay", cla.Sea).Flag(cla.Sea).
		// St. Petersburg (South Coast)
		Prov("stp/sc").Conn("fin", cla.Sea).Conn("gob", cla.Sea).Conn("lia", cla.Sea).Flag(cla.Sea).
		// Kashmir
		Prov("kas").Conn("del", cla.Land).Conn("tib", cla.Land).Conn("afg", cla.Land).Conn("sid", cla.Land).Flag(cla.Land).
		// Red Sea
		Prov("red").Conn("lev", cla.Sea).Conn("lev/sc", cla.Sea).Conn("egy", cla.Sea).Conn("sud", cla.Sea).Conn("eth", cla.Sea).Conn("goa", cla.Sea).Conn("ade", cla.Sea).Conn("ara", cla.Sea).Conn("ara/sc", cla.Sea).Conn("mec", cla.Sea).Flag(cla.Sea).
		// London
		Prov("lon").Conn("not", cla.Sea).Conn("yor", cla.Coast...).Conn("wal", cla.Coast...).Conn("eng", cla.Sea).Flag(cla.Coast...).SC(Britain).
		// Galicia
		Prov("gal").Conn("ukr", cla.Land).Conn("war", cla.Land).Conn("sil", cla.Land).Conn("boh", cla.Land).Conn("vnn", cla.Land).Conn("bud", cla.Land).Conn("rum", cla.Land).Flag(cla.Land).
		// Yemen
		Prov("yem").Conn("goa", cla.Sea).Conn("ars", cla.Sea).Conn("oma", cla.Coast...).Conn("ara", cla.Land).Conn("ade", cla.Coast...).Flag(cla.Coast...).
		// Afghanistan
		Prov("afg").Conn("tib", cla.Land).Conn("tur", cla.Land).Conn("per", cla.Land).Conn("sid", cla.Land).Conn("kas", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// South China Sea
		Prov("scs").Conn("guh", cla.Sea).Conn("bor", cla.Sea).Conn("cel", cla.Sea).Conn("phi", cla.Sea).Conn("ecs", cla.Sea).Conn("for", cla.Sea).Conn("ecs", cla.Sea).Conn("sha", cla.Sea).Conn("gua", cla.Sea).Conn("goo", cla.Sea).Conn("ann", cla.Sea).Conn("sai", cla.Sea).Flag(cla.Sea).
		// Levant
		Prov("lev").Conn("kon", cla.Land).Conn("egy", cla.Land).Conn("mec", cla.Land).Conn("ara", cla.Land).Conn("bag", cla.Land).Conn("arm", cla.Land).Flag(cla.Land).
		// Levant (North Coast)
		Prov("lev/nc").Conn("kon", cla.Sea).Conn("ems", cla.Sea).Conn("egy", cla.Sea).Flag(cla.Sea).
		// Levant (South Coast)
		Prov("lev/sc").Conn("egy", cla.Sea).Conn("red", cla.Sea).Conn("mec", cla.Sea).Flag(cla.Sea).
		// Guangzhou
		Prov("gua").Conn("sha", cla.Coast...).Conn("wuh", cla.Land).Conn("yun", cla.Land).Conn("vit", cla.Coast...).Conn("goo", cla.Sea).Conn("scs", cla.Sea).Flag(cla.Coast...).SC(China).
		// Yunnan
		Prov("yun").Conn("bum", cla.Land).Conn("lao", cla.Land).Conn("vit", cla.Land).Conn("gua", cla.Land).Conn("wuh", cla.Land).Conn("qin", cla.Land).Conn("tib", cla.Land).Flag(cla.Land).
		// Rome
		Prov("rom").Conn("ven", cla.Land).Conn("mil", cla.Land).Conn("pie", cla.Coast...).Conn("gol", cla.Sea).Conn("tyh", cla.Sea).Conn("nap", cla.Coast...).Conn("apu", cla.Land).Flag(cla.Coast...).SC(Italy).
		// Brest
		Prov("bre").Conn("par", cla.Land).Conn("pic", cla.Coast...).Conn("eng", cla.Sea).Conn("mid", cla.Sea).Conn("gas", cla.Coast...).Flag(cla.Coast...).SC(France).
		// Tyrol
		Prov("tyo").Conn("boh", cla.Land).Conn("mun", cla.Land).Conn("swi", cla.Land).Conn("mil", cla.Land).Conn("ven", cla.Land).Conn("tes", cla.Land).Conn("vnn", cla.Land).Flag(cla.Land).
		// Paris
		Prov("par").Conn("bre", cla.Land).Conn("gas", cla.Land).Conn("bug", cla.Land).Conn("pic", cla.Land).Flag(cla.Land).SC(France).
		// Korea
		Prov("kor").Conn("yel", cla.Sea).Conn("ecs", cla.Sea).Conn("soj", cla.Sea).Conn("vla", cla.Coast...).Conn("man", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Ionian Sea
		Prov("ion").Conn("trp", cla.Sea).Conn("cyr", cla.Sea).Conn("ems", cla.Sea).Conn("aeg", cla.Sea).Conn("gre", cla.Sea).Conn("alb", cla.Sea).Conn("adr", cla.Sea).Conn("apu", cla.Sea).Conn("nap", cla.Sea).Conn("tyh", cla.Sea).Conn("tun", cla.Sea).Flag(cla.Sea).
		// South Pacific Ocean
		Prov("spo").Conn("npo", cla.Sea).Conn("tok", cla.Sea).Conn("shi", cla.Sea).Conn("osa", cla.Sea).Conn("ecs", cla.Sea).Conn("phi", cla.Sea).Conn("cel", cla.Sea).Conn("tim", cla.Sea).Conn("bxe", cla.Sea).Conn("bxe", cla.Sea).Conn("bxe", cla.Sea).Flag(cla.Sea).
		// Eastern Indian Ocean
		Prov("eio").Conn("tim", cla.Sea).Conn("jav", cla.Sea).Conn("jvs", cla.Sea).Conn("sum", cla.Sea).Conn("and", cla.Sea).Conn("bay", cla.Sea).Conn("mad", cla.Sea).Conn("wio", cla.Sea).Conn("cey", cla.Sea).Conn("cey", cla.Sea).Conn("wio", cla.Sea).Conn("bxh", cla.Sea).Conn("bxh", cla.Sea).Conn("bxh", cla.Sea).Flag(cla.Sea).
		// Wuhan
		Prov("wuh").Conn("sha", cla.Land).Conn("heb", cla.Land).Conn("qin", cla.Land).Conn("yun", cla.Land).Conn("gua", cla.Land).Flag(cla.Land).SC(China).
		// Portugal
		Prov("por").Conn("spa", cla.Land).Conn("spa/nc", cla.Sea).Conn("spa/sc", cla.Sea).Conn("mid", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Akita
		Prov("aki").Conn("tok", cla.Coast...).Conn("npo", cla.Sea).Conn("soj", cla.Sea).Conn("kyo", cla.Coast...).Flag(cla.Coast...).
		// Sumatra
		Prov("sum").Conn("eio", cla.Sea).Conn("jvs", cla.Sea).Conn("and", cla.Sea).Flag(cla.Coast...).
		// Tibet
		Prov("tib").Conn("afg", cla.Land).Conn("kas", cla.Land).Conn("del", cla.Land).Conn("nep", cla.Land).Conn("cal", cla.Land).Conn("bum", cla.Land).Conn("yun", cla.Land).Conn("qin", cla.Land).Conn("xin", cla.Land).Conn("tur", cla.Land).Flag(cla.Land).
		// Baghdad
		Prov("bag").Conn("ara", cla.Land).Conn("ara/nc", cla.Sea).Conn("psg", cla.Sea).Conn("per", cla.Coast...).Conn("arm", cla.Land).Conn("lev", cla.Land).Flag(cla.Coast...).SC(Turkey).
		// Switzerland
		Prov("swi").Conn("swa", cla.Land).Conn("bug", cla.Land).Conn("mar", cla.Land).Conn("pie", cla.Land).Conn("mil", cla.Land).Conn("tyo", cla.Land).Conn("mun", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Gulf of Lyons
		Prov("gol").Conn("mar", cla.Sea).Conn("spa/sc", cla.Sea).Conn("spa", cla.Sea).Conn("wms", cla.Sea).Conn("tyh", cla.Sea).Conn("rom", cla.Sea).Conn("pie", cla.Sea).Flag(cla.Sea).
		// Skagerrak
		Prov("ska").Conn("swe", cla.Sea).Conn("nay", cla.Sea).Conn("not", cla.Sea).Conn("den", cla.Sea).Flag(cla.Sea).
		// Western Indian Ocean
		Prov("wio").Conn("eio", cla.Sea).Conn("cey", cla.Sea).Conn("eio", cla.Sea).Conn("mad", cla.Sea).Conn("bom", cla.Sea).Conn("ars", cla.Sea).Conn("goa", cla.Sea).Conn("hor", cla.Sea).Conn("bxg", cla.Sea).Conn("bxg", cla.Sea).Conn("bxg", cla.Sea).Flag(cla.Sea).
		// Mogadishu
		Prov("mog").Conn("ken", cla.Coast...).Conn("hor", cla.Sea).Conn("goa", cla.Sea).Conn("awd", cla.Coast...).Conn("eth", cla.Coast...).Flag(cla.Coast...).SC(Italy).
		// Box A bcd
		Prov("bxa").Conn("nao", cla.Sea).Conn("nao", cla.Sea).Conn("bxb", cla.Sea).Conn("bxc", cla.Sea).Conn("bxd", cla.Sea).Flag(cla.Sea).
		// Saxony
		Prov("sax").Conn("mun", cla.Land).Conn("boh", cla.Land).Conn("sil", cla.Land).Conn("ber", cla.Land).Conn("kie", cla.Land).Conn("col", cla.Land).Flag(cla.Land).
		// Ethiopia
		Prov("eth").Conn("goa", cla.Sea).Conn("red", cla.Sea).Conn("sud", cla.Coast...).Conn("ken", cla.Land).Conn("mog", cla.Coast...).Conn("awd", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Tokyo
		Prov("tok").Conn("spo", cla.Sea).Conn("npo", cla.Sea).Conn("aki", cla.Coast...).Conn("kyo", cla.Land).Conn("shi", cla.Coast...).Flag(cla.Coast...).SC(Japan).
		// Peking
		Prov("pek").Conn("yel", cla.Sea).Conn("man", cla.Coast...).Conn("inn", cla.Land).Conn("qin", cla.Land).Conn("heb", cla.Land).Conn("heb/nc", cla.Sea).Flag(cla.Coast...).SC(China).
		// Arabia
		Prov("ara").Conn("bag", cla.Land).Conn("lev", cla.Land).Conn("mec", cla.Land).Conn("ade", cla.Land).Conn("yem", cla.Land).Conn("oma", cla.Land).Flag(cla.Land).
		// Arabia (North Coast)
		Prov("ara/nc").Conn("bag", cla.Sea).Conn("oma", cla.Sea).Conn("psg", cla.Sea).Flag(cla.Sea).
		// Arabia (South Coast)
		Prov("ara/sc").Conn("mec", cla.Sea).Conn("red", cla.Sea).Conn("ade", cla.Sea).Flag(cla.Sea).
		// Barents Sea
		Prov("bar").Conn("noi", cla.Sea).Conn("nay", cla.Sea).Conn("stp", cla.Sea).Conn("stp/nc", cla.Sea).Flag(cla.Sea).
		// North Sea
		Prov("not").Conn("lon", cla.Sea).Conn("eng", cla.Sea).Conn("bel", cla.Sea).Conn("hol", cla.Sea).Conn("hel", cla.Sea).Conn("den", cla.Sea).Conn("ska", cla.Sea).Conn("nay", cla.Sea).Conn("noi", cla.Sea).Conn("edi", cla.Sea).Conn("yor", cla.Sea).Flag(cla.Sea).
		// Inner Mongolia
		Prov("inn").Conn("man", cla.Land).Conn("mon", cla.Land).Conn("xin", cla.Land).Conn("qin", cla.Land).Conn("pek", cla.Land).Flag(cla.Land).
		// Kamchatka
		Prov("kam").Conn("npo", cla.Sea).Conn("sib", cla.Land).Conn("vla", cla.Coast...).Conn("soo", cla.Sea).Flag(cla.Coast...).
		// South Atlantic Ocean
		Prov("sao").Conn("mal", cla.Sea).Conn("mor", cla.Sea).Conn("mid", cla.Sea).Conn("bxc", cla.Sea).Conn("bxc", cla.Sea).Flag(cla.Sea).
		// Deccan
		Prov("dec").Conn("mad", cla.Land).Conn("cal", cla.Land).Conn("del", cla.Land).Conn("bom", cla.Land).Flag(cla.Land).
		// Karafuto
		Prov("kar").Conn("sak", cla.Coast...).Conn("soj", cla.Sea).Conn("npo", cla.Sea).Conn("soo", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Algeria
		Prov("alg").Conn("wms", cla.Sea).Conn("mor", cla.Coast...).Conn("sah", cla.Land).Conn("tun", cla.Coast...).Conn("tyh", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Awdal
		Prov("awd").Conn("eth", cla.Coast...).Conn("mog", cla.Coast...).Conn("goa", cla.Sea).Flag(cla.Coast...).
		// Baltic Sea
		Prov("bal").Conn("pru", cla.Sea).Conn("lia", cla.Sea).Conn("gob", cla.Sea).Conn("swe", cla.Sea).Conn("den", cla.Sea).Conn("kie", cla.Sea).Conn("ber", cla.Sea).Flag(cla.Sea).
		// Calcutta
		Prov("cal").Conn("bay", cla.Sea).Conn("bum", cla.Coast...).Conn("tib", cla.Land).Conn("nep", cla.Land).Conn("del", cla.Land).Conn("dec", cla.Land).Conn("mad", cla.Coast...).Flag(cla.Coast...).SC(India).
		// Box D aef
		Prov("bxd").Conn("npo", cla.Sea).Conn("npo", cla.Sea).Conn("bxa", cla.Sea).Conn("bxe", cla.Sea).Conn("bxf", cla.Sea).Flag(cla.Sea).
		// Edinburgh
		Prov("edi").Conn("cly", cla.Coast...).Conn("lie", cla.Land).Conn("yor", cla.Coast...).Conn("not", cla.Sea).Conn("noi", cla.Sea).Flag(cla.Coast...).SC(Britain).
		// Piedmont
		Prov("pie").Conn("swi", cla.Land).Conn("mar", cla.Coast...).Conn("gol", cla.Sea).Conn("rom", cla.Coast...).Conn("mil", cla.Land).Flag(cla.Coast...).
		// Budapest
		Prov("bud").Conn("gal", cla.Land).Conn("vnn", cla.Land).Conn("tes", cla.Land).Conn("sar", cla.Land).Conn("ser", cla.Land).Conn("rum", cla.Land).Flag(cla.Land).SC(Austria).
		// Vladivostok
		Prov("vla").Conn("soj", cla.Sea).Conn("soo", cla.Sea).Conn("kam", cla.Coast...).Conn("sib", cla.Land).Conn("man", cla.Land).Conn("kor", cla.Coast...).Flag(cla.Coast...).SC(Russia).
		// Kyoto
		Prov("kyo").Conn("shi", cla.Land).Conn("tok", cla.Land).Conn("aki", cla.Coast...).Conn("soj", cla.Sea).Conn("kag", cla.Coast...).Conn("osa", cla.Land).Flag(cla.Coast...).SC(Japan).
		// Bulgaria
		Prov("bul").Conn("mac", cla.Land).Conn("con", cla.Coast...).Conn("bla", cla.Sea).Conn("rum", cla.Coast...).Conn("ser", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Horn of Africa
		Prov("hor").Conn("wio", cla.Sea).Conn("goa", cla.Sea).Conn("mog", cla.Sea).Conn("ken", cla.Sea).Flag(cla.Sea).
		// Delhi
		Prov("del").Conn("sid", cla.Land).Conn("bom", cla.Land).Conn("dec", cla.Land).Conn("cal", cla.Land).Conn("nep", cla.Land).Conn("tib", cla.Land).Conn("kas", cla.Land).Flag(cla.Land).SC(India).
		// Java Sea
		Prov("jvs").Conn("cel", cla.Sea).Conn("bor", cla.Sea).Conn("guh", cla.Sea).Conn("sig", cla.Sea).Conn("and", cla.Sea).Conn("sum", cla.Sea).Conn("eio", cla.Sea).Conn("jav", cla.Sea).Flag(cla.Sea).
		// Morocco
		Prov("mor").Conn("sao", cla.Sea).Conn("mal", cla.Coast...).Conn("sah", cla.Land).Conn("alg", cla.Coast...).Conn("wms", cla.Sea).Conn("mid", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Serbia
		Prov("ser").Conn("sar", cla.Land).Conn("alb", cla.Land).Conn("mac", cla.Land).Conn("bul", cla.Land).Conn("rum", cla.Land).Conn("bud", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Borneo
		Prov("bor").Conn("jvs", cla.Sea).Conn("cel", cla.Sea).Conn("scs", cla.Sea).Conn("guh", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Ceylon
		Prov("cey").Conn("eio", cla.Sea).Conn("eio", cla.Sea).Conn("wio", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Belgium
		Prov("bel").Conn("pic", cla.Coast...).Conn("bug", cla.Land).Conn("col", cla.Land).Conn("hol", cla.Coast...).Conn("not", cla.Sea).Conn("eng", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Kagoshima
		Prov("kag").Conn("osa", cla.Coast...).Conn("kyo", cla.Coast...).Conn("soj", cla.Sea).Conn("ecs", cla.Sea).Flag(cla.Coast...).
		// Ireland
		Prov("ire").Conn("nao", cla.Sea).Conn("mid", cla.Sea).Conn("iri", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Western Mediterranean Sea
		Prov("wms").Conn("alg", cla.Sea).Conn("tyh", cla.Sea).Conn("gol", cla.Sea).Conn("spa/sc", cla.Sea).Conn("spa", cla.Sea).Conn("mid", cla.Sea).Conn("mor", cla.Sea).Flag(cla.Sea).
		// Mali
		Prov("mal").Conn("sah", cla.Land).Conn("mor", cla.Coast...).Conn("sao", cla.Sea).Flag(cla.Coast...).
		// Liverpool
		Prov("lie").Conn("iri", cla.Sea).Conn("wal", cla.Coast...).Conn("yor", cla.Land).Conn("edi", cla.Land).Conn("cly", cla.Coast...).Conn("nao", cla.Sea).Flag(cla.Coast...).SC(Britain).
		// Sapporo
		Prov("sap").Conn("npo", cla.Sea).Conn("soj", cla.Sea).Conn("npo", cla.Sea).Flag(cla.Coast...).SC(Japan).
		// Sudan
		Prov("sud").Conn("ken", cla.Land).Conn("eth", cla.Coast...).Conn("red", cla.Sea).Conn("egy", cla.Coast...).Conn("fez", cla.Land).Conn("cen", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Eastern Mediterranean Sea
		Prov("ems").Conn("kon", cla.Sea).Conn("aeg", cla.Sea).Conn("ion", cla.Sea).Conn("cyr", cla.Sea).Conn("egy", cla.Sea).Conn("lev", cla.Sea).Conn("lev/nc", cla.Sea).Flag(cla.Sea).
		// Denmark
		Prov("den").Conn("not", cla.Sea).Conn("hel", cla.Sea).Conn("kie", cla.Coast...).Conn("bal", cla.Sea).Conn("swe", cla.Coast...).Conn("ska", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Aegean Sea
		Prov("aeg").Conn("gre", cla.Sea).Conn("ion", cla.Sea).Conn("ems", cla.Sea).Conn("kon", cla.Sea).Conn("con", cla.Sea).Conn("mac", cla.Sea).Flag(cla.Sea).
		// Adriatic Sea
		Prov("adr").Conn("tes", cla.Sea).Conn("ven", cla.Sea).Conn("apu", cla.Sea).Conn("ion", cla.Sea).Conn("alb", cla.Sea).Conn("sar", cla.Sea).Flag(cla.Sea).
		// Sweden
		Prov("swe").Conn("bal", cla.Sea).Conn("gob", cla.Sea).Conn("fin", cla.Coast...).Conn("nay", cla.Coast...).Conn("ska", cla.Sea).Conn("den", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Heligoland Bight
		Prov("hel").Conn("den", cla.Sea).Conn("not", cla.Sea).Conn("hol", cla.Sea).Conn("kie", cla.Sea).Flag(cla.Sea).
		// English Channel
		Prov("eng").Conn("wal", cla.Sea).Conn("iri", cla.Sea).Conn("mid", cla.Sea).Conn("bre", cla.Sea).Conn("pic", cla.Sea).Conn("bel", cla.Sea).Conn("not", cla.Sea).Conn("lon", cla.Sea).Flag(cla.Sea).
		// Caucasus
		Prov("cau").Conn("sev", cla.Land).Conn("arm", cla.Land).Conn("per", cla.Land).Conn("oms", cla.Land).Conn("mos", cla.Land).Flag(cla.Land).
		// Armenia
		Prov("arm").Conn("ank", cla.Coast...).Conn("kon", cla.Land).Conn("lev", cla.Land).Conn("bag", cla.Land).Conn("per", cla.Land).Conn("cau", cla.Land).Conn("sev", cla.Coast...).Conn("bla", cla.Sea).Flag(cla.Coast...).
		// Fez
		Prov("fez").Conn("egy", cla.Land).Conn("cyr", cla.Land).Conn("trp", cla.Land).Conn("sah", cla.Land).Conn("cen", cla.Land).Conn("sud", cla.Land).Flag(cla.Land).
		// Norway
		Prov("nay").Conn("swe", cla.Coast...).Conn("fin", cla.Land).Conn("stp", cla.Land).Conn("stp/nc", cla.Sea).Conn("bar", cla.Sea).Conn("noi", cla.Sea).Conn("not", cla.Sea).Conn("ska", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Java
		Prov("jav").Conn("jvs", cla.Sea).Conn("eio", cla.Sea).Conn("tim", cla.Sea).Conn("cel", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Holland
		Prov("hol").Conn("hel", cla.Sea).Conn("not", cla.Sea).Conn("bel", cla.Coast...).Conn("col", cla.Land).Conn("kie", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Central Africa
		Prov("cen").Conn("sud", cla.Land).Conn("fez", cla.Land).Conn("sah", cla.Land).Flag(cla.Land).
		// Arabian Sea
		Prov("ars").Conn("yem", cla.Sea).Conn("goa", cla.Sea).Conn("wio", cla.Sea).Conn("bom", cla.Sea).Conn("sid", cla.Sea).Conn("per", cla.Sea).Conn("psg", cla.Sea).Conn("oma", cla.Sea).Flag(cla.Sea).
		// Tyrrhenian Sea
		Prov("tyh").Conn("wms", cla.Sea).Conn("alg", cla.Sea).Conn("tun", cla.Sea).Conn("ion", cla.Sea).Conn("nap", cla.Sea).Conn("rom", cla.Sea).Conn("gol", cla.Sea).Flag(cla.Sea).
		// Thailand
		Prov("tha").Conn("cam", cla.Land).Conn("lao", cla.Land).Conn("bum", cla.Land).Conn("sig", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Thailand (East Coast)
		Prov("tha/ec").Conn("cam", cla.Sea).Conn("sig", cla.Sea).Conn("guh", cla.Sea).Flag(cla.Sea).
		// Thailand (West Coast)
		Prov("tha/wc").Conn("bum", cla.Sea).Conn("and", cla.Sea).Conn("sig", cla.Sea).Flag(cla.Sea).
		// Cyrene
		Prov("cyr").Conn("fez", cla.Land).Conn("egy", cla.Coast...).Conn("ems", cla.Sea).Conn("ion", cla.Sea).Conn("trp", cla.Coast...).Flag(cla.Coast...).
		// Gulf of Bothnia
		Prov("gob").Conn("bal", cla.Sea).Conn("lia", cla.Sea).Conn("stp", cla.Sea).Conn("stp/sc", cla.Sea).Conn("fin", cla.Sea).Conn("swe", cla.Sea).Flag(cla.Sea).
		// Gascony
		Prov("gas").Conn("mar", cla.Land).Conn("par", cla.Land).Conn("bre", cla.Coast...).Conn("mid", cla.Sea).Conn("spa", cla.Land).Conn("spa/nc", cla.Sea).Conn("bug", cla.Land).Flag(cla.Coast...).
		// Celebes Sea
		Prov("cel").Conn("jvs", cla.Sea).Conn("jav", cla.Sea).Conn("tim", cla.Sea).Conn("spo", cla.Sea).Conn("phi", cla.Sea).Conn("scs", cla.Sea).Conn("bor", cla.Sea).Flag(cla.Sea).
		// Timor Sea
		Prov("tim").Conn("spo", cla.Sea).Conn("cel", cla.Sea).Conn("jav", cla.Sea).Conn("eio", cla.Sea).Conn("bxf", cla.Sea).Conn("bxf", cla.Sea).Flag(cla.Sea).
		// Sea of Okhotsk
		Prov("soo").Conn("sak", cla.Sea).Conn("kar", cla.Sea).Conn("npo", cla.Sea).Conn("kam", cla.Sea).Conn("vla", cla.Sea).Conn("soj", cla.Sea).Flag(cla.Sea).
		// Philippines
		Prov("phi").Conn("ecs", cla.Sea).Conn("scs", cla.Sea).Conn("cel", cla.Sea).Conn("spo", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Cologne
		Prov("col").Conn("bel", cla.Land).Conn("bug", cla.Land).Conn("swa", cla.Land).Conn("mun", cla.Land).Conn("sax", cla.Land).Conn("kie", cla.Land).Conn("hol", cla.Land).Flag(cla.Land).SC(Germany).
		// Annam
		Prov("ann").Conn("cam", cla.Land).Conn("sai", cla.Coast...).Conn("scs", cla.Sea).Conn("goo", cla.Sea).Conn("vit", cla.Coast...).Conn("lao", cla.Land).Flag(cla.Coast...).
		// Manchuria
		Prov("man").Conn("yel", cla.Sea).Conn("kor", cla.Coast...).Conn("vla", cla.Land).Conn("sib", cla.Land).Conn("mon", cla.Land).Conn("inn", cla.Land).Conn("pek", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Constantinople
		Prov("con").Conn("mac", cla.Coast...).Conn("aeg", cla.Sea).Conn("kon", cla.Coast...).Conn("ank", cla.Coast...).Conn("bla", cla.Sea).Conn("bul", cla.Coast...).Flag(cla.Coast...).SC(Turkey).
		// Bay of Bengal
		Prov("bay").Conn("mad", cla.Sea).Conn("eio", cla.Sea).Conn("and", cla.Sea).Conn("bum", cla.Sea).Conn("cal", cla.Sea).Flag(cla.Sea).
		// Marseilles
		Prov("mar").Conn("gol", cla.Sea).Conn("pie", cla.Coast...).Conn("swi", cla.Land).Conn("bug", cla.Land).Conn("gas", cla.Land).Conn("spa", cla.Land).Conn("spa/sc", cla.Sea).Flag(cla.Coast...).SC(France).
		// York
		Prov("yor").Conn("not", cla.Sea).Conn("edi", cla.Coast...).Conn("lie", cla.Land).Conn("wal", cla.Land).Conn("lon", cla.Coast...).Flag(cla.Coast...).
		// Ukraine
		Prov("ukr").Conn("rum", cla.Land).Conn("sev", cla.Land).Conn("mos", cla.Land).Conn("war", cla.Land).Conn("gal", cla.Land).Flag(cla.Land).
		// Mid Atlantic Ocean
		Prov("mid").Conn("sao", cla.Sea).Conn("mor", cla.Sea).Conn("wms", cla.Sea).Conn("spa/sc", cla.Sea).Conn("spa/nc", cla.Sea).Conn("spa", cla.Sea).Conn("por", cla.Sea).Conn("gas", cla.Sea).Conn("bre", cla.Sea).Conn("eng", cla.Sea).Conn("iri", cla.Sea).Conn("ire", cla.Sea).Conn("nao", cla.Sea).Conn("bxb", cla.Sea).Conn("bxb", cla.Sea).Conn("bxb", cla.Sea).Flag(cla.Sea).
		// Saigon
		Prov("sai").Conn("scs", cla.Sea).Conn("ann", cla.Coast...).Conn("cam", cla.Coast...).Conn("guh", cla.Sea).Flag(cla.Coast...).SC(France).
		// Gulf of Tonkin
		Prov("goo").Conn("ann", cla.Sea).Conn("scs", cla.Sea).Conn("gua", cla.Sea).Conn("vit", cla.Sea).Flag(cla.Sea).
		// Qinghai
		Prov("qin").Conn("wuh", cla.Land).Conn("heb", cla.Land).Conn("pek", cla.Land).Conn("inn", cla.Land).Conn("xin", cla.Land).Conn("tib", cla.Land).Conn("yun", cla.Land).Flag(cla.Land).
		// Cambodia
		Prov("cam").Conn("ann", cla.Land).Conn("lao", cla.Land).Conn("tha", cla.Land).Conn("tha/ec", cla.Sea).Conn("guh", cla.Sea).Conn("sai", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// East China Sea
		Prov("ecs").Conn("phi", cla.Sea).Conn("spo", cla.Sea).Conn("osa", cla.Sea).Conn("kag", cla.Sea).Conn("soj", cla.Sea).Conn("kor", cla.Sea).Conn("yel", cla.Sea).Conn("tsi", cla.Sea).Conn("heb", cla.Sea).Conn("heb/sc", cla.Sea).Conn("sha", cla.Sea).Conn("scs", cla.Sea).Conn("for", cla.Sea).Conn("for", cla.Sea).Conn("scs", cla.Sea).Flag(cla.Sea).
		// North Atlantic Ocean
		Prov("nao").Conn("bxa", cla.Sea).Conn("bxa", cla.Sea).Conn("mid", cla.Sea).Conn("ire", cla.Sea).Conn("iri", cla.Sea).Conn("lie", cla.Sea).Conn("cly", cla.Sea).Conn("noi", cla.Sea).Flag(cla.Sea).
		// Swabia
		Prov("swa").Conn("swi", cla.Land).Conn("mun", cla.Land).Conn("col", cla.Land).Conn("bug", cla.Land).Flag(cla.Land).
		// Kenya
		Prov("ken").Conn("hor", cla.Sea).Conn("mog", cla.Coast...).Conn("eth", cla.Land).Conn("sud", cla.Land).Flag(cla.Coast...).
		// Sea of Japan
		Prov("soj").Conn("kor", cla.Sea).Conn("ecs", cla.Sea).Conn("kag", cla.Sea).Conn("kyo", cla.Sea).Conn("aki", cla.Sea).Conn("npo", cla.Sea).Conn("sap", cla.Sea).Conn("npo", cla.Sea).Conn("kar", cla.Sea).Conn("sak", cla.Sea).Conn("soo", cla.Sea).Conn("vla", cla.Sea).Flag(cla.Sea).
		// Sarajevo
		Prov("sar").Conn("alb", cla.Coast...).Conn("ser", cla.Land).Conn("bud", cla.Land).Conn("tes", cla.Coast...).Conn("adr", cla.Sea).Flag(cla.Coast...).SC(Austria).
		// Konya
		Prov("kon").Conn("ank", cla.Land).Conn("con", cla.Coast...).Conn("aeg", cla.Sea).Conn("ems", cla.Sea).Conn("lev", cla.Land).Conn("lev/nc", cla.Sea).Conn("arm", cla.Land).Flag(cla.Coast...).
		// Aden
		Prov("ade").Conn("goa", cla.Sea).Conn("yem", cla.Coast...).Conn("ara", cla.Land).Conn("ara/sc", cla.Sea).Conn("red", cla.Sea).Flag(cla.Coast...).SC(Britain).
		// Sindh
		Prov("sid").Conn("del", cla.Land).Conn("kas", cla.Land).Conn("afg", cla.Land).Conn("per", cla.Coast...).Conn("ars", cla.Sea).Conn("bom", cla.Coast...).Flag(cla.Coast...).
		// Spain
		Prov("spa").Conn("mar", cla.Land).Conn("gas", cla.Land).Conn("por", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Spain (North Coast)
		Prov("spa/nc").Conn("gas", cla.Sea).Conn("mid", cla.Sea).Conn("por", cla.Sea).Flag(cla.Sea).
		// Spain (South Coast)
		Prov("spa/sc").Conn("wms", cla.Sea).Conn("mar", cla.Sea).Conn("mid", cla.Sea).Conn("gol", cla.Sea).Conn("por", cla.Sea).Flag(cla.Sea).
		// Warsaw
		Prov("war").Conn("mos", cla.Land).Conn("lia", cla.Land).Conn("pru", cla.Land).Conn("sil", cla.Land).Conn("gal", cla.Land).Conn("ukr", cla.Land).Flag(cla.Land).SC(Russia).
		// Norwegian Sea
		Prov("noi").Conn("nao", cla.Sea).Conn("cly", cla.Sea).Conn("edi", cla.Sea).Conn("not", cla.Sea).Conn("nay", cla.Sea).Conn("bar", cla.Sea).Flag(cla.Sea).
		// Singapore
		Prov("sig").Conn("jvs", cla.Sea).Conn("guh", cla.Sea).Conn("tha", cla.Land).Conn("tha/ec", cla.Sea).Conn("tha/wc", cla.Sea).Conn("and", cla.Sea).Flag(cla.Coast...).SC(Britain).
		// Gulf of Aden
		Prov("goa").Conn("eth", cla.Sea).Conn("awd", cla.Sea).Conn("mog", cla.Sea).Conn("hor", cla.Sea).Conn("wio", cla.Sea).Conn("ars", cla.Sea).Conn("yem", cla.Sea).Conn("ade", cla.Sea).Conn("red", cla.Sea).Flag(cla.Sea).
		// Mongolia
		Prov("mon").Conn("sib", cla.Land).Conn("xin", cla.Land).Conn("inn", cla.Land).Conn("man", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Wales
		Prov("wal").Conn("eng", cla.Sea).Conn("lon", cla.Coast...).Conn("yor", cla.Land).Conn("lie", cla.Coast...).Conn("iri", cla.Sea).Flag(cla.Coast...).
		// Yellow Sea
		Prov("yel").Conn("kor", cla.Sea).Conn("man", cla.Sea).Conn("pek", cla.Sea).Conn("heb", cla.Sea).Conn("heb/nc", cla.Sea).Conn("tsi", cla.Sea).Conn("ecs", cla.Sea).Flag(cla.Sea).
		// Greece
		Prov("gre").Conn("aeg", cla.Sea).Conn("mac", cla.Coast...).Conn("alb", cla.Coast...).Conn("ion", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Venice
		Prov("ven").Conn("rom", cla.Land).Conn("apu", cla.Coast...).Conn("adr", cla.Sea).Conn("tes", cla.Coast...).Conn("tyo", cla.Land).Conn("mil", cla.Land).Flag(cla.Coast...).
		// Vienna
		Prov("vnn").Conn("boh", cla.Land).Conn("tyo", cla.Land).Conn("tes", cla.Land).Conn("bud", cla.Land).Conn("gal", cla.Land).Flag(cla.Land).SC(Austria).
		// Andaman Sea
		Prov("and").Conn("sum", cla.Sea).Conn("jvs", cla.Sea).Conn("sig", cla.Sea).Conn("tha", cla.Sea).Conn("tha/wc", cla.Sea).Conn("bum", cla.Sea).Conn("bay", cla.Sea).Conn("eio", cla.Sea).Flag(cla.Sea).
		// Nepal
		Prov("nep").Conn("tib", cla.Land).Conn("del", cla.Land).Conn("cal", cla.Land).Flag(cla.Land).
		// Box E bdf
		Prov("bxe").Conn("spo", cla.Sea).Conn("spo", cla.Sea).Conn("spo", cla.Sea).Conn("bxb", cla.Sea).Conn("bxd", cla.Sea).Conn("bxf", cla.Sea).Flag(cla.Sea).
		// Box F cdegh
		Prov("bxf").Conn("tim", cla.Sea).Conn("tim", cla.Sea).Conn("bxc", cla.Sea).Conn("bxd", cla.Sea).Conn("bxe", cla.Sea).Conn("bxg", cla.Sea).Conn("bxh", cla.Sea).Flag(cla.Sea).
		// Bohemia
		Prov("boh").Conn("vnn", cla.Land).Conn("gal", cla.Land).Conn("sil", cla.Land).Conn("sax", cla.Land).Conn("mun", cla.Land).Conn("tyo", cla.Land).Flag(cla.Land).
		// Laos
		Prov("lao").Conn("bum", cla.Land).Conn("tha", cla.Land).Conn("cam", cla.Land).Conn("ann", cla.Land).Conn("vit", cla.Land).Conn("yun", cla.Land).Flag(cla.Land).
		// Rumania
		Prov("rum").Conn("bul", cla.Coast...).Conn("bla", cla.Sea).Conn("sev", cla.Coast...).Conn("ukr", cla.Land).Conn("gal", cla.Land).Conn("bud", cla.Land).Conn("ser", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Shanghai
		Prov("sha").Conn("wuh", cla.Land).Conn("gua", cla.Coast...).Conn("scs", cla.Sea).Conn("ecs", cla.Sea).Conn("heb", cla.Land).Conn("heb/sc", cla.Sea).Flag(cla.Coast...).SC(China).
		// Milan
		Prov("mil").Conn("ven", cla.Land).Conn("tyo", cla.Land).Conn("swi", cla.Land).Conn("pie", cla.Land).Conn("rom", cla.Land).Flag(cla.Land).SC(Italy).
		// Macedonia
		Prov("mac").Conn("bul", cla.Land).Conn("ser", cla.Land).Conn("alb", cla.Land).Conn("gre", cla.Coast...).Conn("aeg", cla.Sea).Conn("con", cla.Coast...).Flag(cla.Coast...).
		// Black Sea
		Prov("bla").Conn("bul", cla.Sea).Conn("con", cla.Sea).Conn("ank", cla.Sea).Conn("arm", cla.Sea).Conn("sev", cla.Sea).Conn("rum", cla.Sea).Flag(cla.Sea).
		// Omsk
		Prov("oms").Conn("tur", cla.Land).Conn("sib", cla.Land).Conn("stp", cla.Land).Conn("mos", cla.Land).Conn("cau", cla.Land).Flag(cla.Land).SC(Russia).
		// Xinjiang
		Prov("xin").Conn("qin", cla.Land).Conn("inn", cla.Land).Conn("mon", cla.Land).Conn("sib", cla.Land).Conn("tur", cla.Land).Conn("tib", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// Egypt
		Prov("egy").Conn("lev", cla.Land).Conn("lev/sc", cla.Sea).Conn("lev/nc", cla.Sea).Conn("ems", cla.Sea).Conn("cyr", cla.Coast...).Conn("fez", cla.Land).Conn("sud", cla.Coast...).Conn("red", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Formosa
		Prov("for").Conn("ecs", cla.Sea).Conn("ecs", cla.Sea).Conn("scs", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Box G cfh
		Prov("bxg").Conn("wio", cla.Sea).Conn("wio", cla.Sea).Conn("wio", cla.Sea).Conn("bxc", cla.Sea).Conn("bxf", cla.Sea).Conn("bxh", cla.Sea).Flag(cla.Sea).
		// Osaka
		Prov("osa").Conn("kag", cla.Coast...).Conn("ecs", cla.Sea).Conn("spo", cla.Sea).Conn("shi", cla.Coast...).Conn("kyo", cla.Land).Flag(cla.Coast...).SC(Japan).
		// Turkmenistan
		Prov("tur").Conn("sib", cla.Land).Conn("oms", cla.Land).Conn("per", cla.Land).Conn("afg", cla.Land).Conn("tib", cla.Land).Conn("xin", cla.Land).Flag(cla.Land).
		// Tsingtao
		Prov("tsi").Conn("ecs", cla.Sea).Conn("yel", cla.Sea).Conn("heb", cla.Land).Conn("heb/nc", cla.Sea).Conn("heb/sc", cla.Sea).Flag(cla.Coast...).SC(Germany).
		// Shizuoka
		Prov("shi").Conn("kyo", cla.Land).Conn("osa", cla.Coast...).Conn("spo", cla.Sea).Conn("tok", cla.Coast...).Flag(cla.Coast...).
		// Clyde
		Prov("cly").Conn("edi", cla.Coast...).Conn("noi", cla.Sea).Conn("nao", cla.Sea).Conn("lie", cla.Coast...).Flag(cla.Coast...).
		// Sahara
		Prov("sah").Conn("cen", cla.Land).Conn("fez", cla.Land).Conn("trp", cla.Land).Conn("tun", cla.Land).Conn("alg", cla.Land).Conn("mor", cla.Land).Conn("mal", cla.Land).Flag(cla.Land).
		// North Pacific Ocean
		Prov("npo").Conn("kam", cla.Sea).Conn("soo", cla.Sea).Conn("kar", cla.Sea).Conn("soj", cla.Sea).Conn("sap", cla.Sea).Conn("sap", cla.Sea).Conn("soj", cla.Sea).Conn("aki", cla.Sea).Conn("tok", cla.Sea).Conn("spo", cla.Sea).Conn("bxd", cla.Sea).Conn("bxd", cla.Sea).Flag(cla.Sea).
		// Apulia
		Prov("apu").Conn("nap", cla.Coast...).Conn("ion", cla.Sea).Conn("adr", cla.Sea).Conn("ven", cla.Coast...).Conn("rom", cla.Land).Flag(cla.Coast...).
		// Burma
		Prov("bum").Conn("yun", cla.Land).Conn("tib", cla.Land).Conn("cal", cla.Coast...).Conn("bay", cla.Sea).Conn("and", cla.Sea).Conn("tha", cla.Land).Conn("tha/wc", cla.Sea).Conn("lao", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Box B ace
		Prov("bxb").Conn("mid", cla.Sea).Conn("mid", cla.Sea).Conn("mid", cla.Sea).Conn("bxa", cla.Sea).Conn("bxc", cla.Sea).Conn("bxe", cla.Sea).Flag(cla.Sea).
		// Box C abfgh
		Prov("bxc").Conn("sao", cla.Sea).Conn("sao", cla.Sea).Conn("bxa", cla.Sea).Conn("bxb", cla.Sea).Conn("bxf", cla.Sea).Conn("bxg", cla.Sea).Conn("bxh", cla.Sea).Flag(cla.Sea).
		// Siberia
		Prov("sib").Conn("tur", cla.Land).Conn("xin", cla.Land).Conn("mon", cla.Land).Conn("man", cla.Land).Conn("vla", cla.Land).Conn("kam", cla.Land).Conn("oms", cla.Land).Flag(cla.Land).
		// Oman
		Prov("oma").Conn("ars", cla.Sea).Conn("psg", cla.Sea).Conn("ara", cla.Land).Conn("ara/nc", cla.Sea).Conn("yem", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Gulf of Thailand
		Prov("guh").Conn("scs", cla.Sea).Conn("sai", cla.Sea).Conn("cam", cla.Sea).Conn("tha", cla.Sea).Conn("tha/ec", cla.Sea).Conn("sig", cla.Sea).Conn("jvs", cla.Sea).Conn("bor", cla.Sea).Flag(cla.Sea).
		// Irish Sea
		Prov("iri").Conn("nao", cla.Sea).Conn("ire", cla.Sea).Conn("mid", cla.Sea).Conn("eng", cla.Sea).Conn("wal", cla.Sea).Conn("lie", cla.Sea).Flag(cla.Sea).
		// Finland
		Prov("fin").Conn("stp", cla.Land).Conn("stp/sc", cla.Sea).Conn("nay", cla.Land).Conn("swe", cla.Coast...).Conn("gob", cla.Sea).Flag(cla.Coast...).
		// Prussia
		Prov("pru").Conn("sil", cla.Land).Conn("war", cla.Land).Conn("lia", cla.Coast...).Conn("bal", cla.Sea).Conn("ber", cla.Coast...).Flag(cla.Coast...).
		// Berlin
		Prov("ber").Conn("kie", cla.Coast...).Conn("sax", cla.Land).Conn("sil", cla.Land).Conn("pru", cla.Coast...).Conn("bal", cla.Sea).Flag(cla.Coast...).SC(Germany).
		// Persia
		Prov("per").Conn("tur", cla.Land).Conn("cau", cla.Land).Conn("arm", cla.Land).Conn("bag", cla.Coast...).Conn("psg", cla.Sea).Conn("ars", cla.Sea).Conn("sid", cla.Coast...).Conn("afg", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Livonia
		Prov("lia").Conn("bal", cla.Sea).Conn("pru", cla.Coast...).Conn("war", cla.Land).Conn("mos", cla.Land).Conn("stp", cla.Land).Conn("stp/sc", cla.Sea).Conn("gob", cla.Sea).Flag(cla.Coast...).
		// Burgundy
		Prov("bug").Conn("swi", cla.Land).Conn("swa", cla.Land).Conn("col", cla.Land).Conn("bel", cla.Land).Conn("pic", cla.Land).Conn("par", cla.Land).Conn("mar", cla.Land).Conn("gas", cla.Land).Flag(cla.Land).
		// Mecca
		Prov("mec").Conn("lev", cla.Land).Conn("lev/sc", cla.Sea).Conn("red", cla.Sea).Conn("ara", cla.Land).Conn("ara/sc", cla.Sea).Flag(cla.Coast...).SC(Turkey).
		// Persian Gulf
		Prov("psg").Conn("per", cla.Sea).Conn("bag", cla.Sea).Conn("ara", cla.Sea).Conn("ara/nc", cla.Sea).Conn("oma", cla.Sea).Conn("ars", cla.Sea).Flag(cla.Sea).
		// Naples
		Prov("nap").Conn("apu", cla.Coast...).Conn("rom", cla.Coast...).Conn("tyh", cla.Sea).Conn("ion", cla.Sea).Flag(cla.Coast...).SC(Italy).
		// Tripolitania
		Prov("trp").Conn("ion", cla.Sea).Conn("tun", cla.Coast...).Conn("sah", cla.Land).Conn("fez", cla.Land).Conn("cyr", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Sakhalin
		Prov("sak").Conn("soo", cla.Sea).Conn("soj", cla.Sea).Conn("kar", cla.Coast...).Flag(cla.Coast...).
		// Kiel
		Prov("kie").Conn("ber", cla.Coast...).Conn("bal", cla.Sea).Conn("den", cla.Coast...).Conn("hel", cla.Sea).Conn("hol", cla.Coast...).Conn("col", cla.Land).Conn("sax", cla.Land).Flag(cla.Coast...).SC(Germany).
		// Moscow
		Prov("mos").Conn("war", cla.Land).Conn("ukr", cla.Land).Conn("sev", cla.Land).Conn("cau", cla.Land).Conn("oms", cla.Land).Conn("stp", cla.Land).Conn("lia", cla.Land).Flag(cla.Land).SC(Russia).
		// Ankara
		Prov("ank").Conn("kon", cla.Land).Conn("arm", cla.Coast...).Conn("bla", cla.Sea).Conn("con", cla.Coast...).Flag(cla.Coast...).SC(Turkey).
		// Madras
		Prov("mad").Conn("bom", cla.Coast...).Conn("wio", cla.Sea).Conn("eio", cla.Sea).Conn("bay", cla.Sea).Conn("cal", cla.Coast...).Conn("dec", cla.Land).Flag(cla.Coast...).SC(India).
		// Trieste
		Prov("tes").Conn("adr", cla.Sea).Conn("sar", cla.Coast...).Conn("bud", cla.Land).Conn("vnn", cla.Land).Conn("tyo", cla.Land).Conn("ven", cla.Coast...).Flag(cla.Coast...).SC(Austria).
		// Vietnaam
		Prov("vit").Conn("lao", cla.Land).Conn("ann", cla.Coast...).Conn("goo", cla.Sea).Conn("gua", cla.Coast...).Conn("yun", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Picardy
		Prov("pic").Conn("bel", cla.Coast...).Conn("eng", cla.Sea).Conn("bre", cla.Coast...).Conn("par", cla.Land).Conn("bug", cla.Land).Flag(cla.Coast...).
		// Munich
		Prov("mun").Conn("sax", cla.Land).Conn("col", cla.Land).Conn("swa", cla.Land).Conn("swi", cla.Land).Conn("tyo", cla.Land).Conn("boh", cla.Land).Flag(cla.Land).SC(Germany).
		Done()
}
