package coldwar

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
	USSR dip.Nation = "USSR"
	NATO dip.Nation = "NATO"
)

var Nations = []dip.Nation{USSR, NATO}

var ColdWarVariant = common.Variant{
	Name:        "Cold War",
	Graph:       func() dip.Graph { return ColdWarGraph() },
	Start:       ColdWarStart,
	Blank:       ColdWarBlank,
	Phase:       classical.Phase,
	ParseOrders: orders.ParseAll,
	ParseOrder:  orders.Parse,
	OrderTypes:  orders.OrderTypes(),
	Nations:     Nations,
	PhaseTypes:  cla.PhaseTypes,
	Seasons:     cla.Seasons,
	UnitTypes:   cla.UnitTypes,
	SoloSupplyCenters: 14,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/coldwarmap.svg")
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
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules: "",
}

func ColdWarBlank(phase dip.Phase) *state.State {
	return state.New(ColdWarGraph(), phase, classical.BackupRule)
}

func ColdWarStart() (result *state.State, err error) {
	startPhase := classical.Phase(1960, cla.Spring, cla.Movement)
	result = state.New(ColdWarGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[dip.Province]dip.Unit{
		"len": dip.Unit{cla.Fleet, USSR},
		"alb": dip.Unit{cla.Fleet, USSR},
		"hav": dip.Unit{cla.Fleet, USSR},
		"mos": dip.Unit{cla.Army, USSR},
		"sha": dip.Unit{cla.Army, USSR},
		"vla": dip.Unit{cla.Army, USSR},
		"lon": dip.Unit{cla.Fleet, NATO},
		"ist": dip.Unit{cla.Fleet, NATO},
		"aus": dip.Unit{cla.Fleet, NATO},
		"new": dip.Unit{cla.Army, NATO},
		"los": dip.Unit{cla.Army, NATO},
		"par": dip.Unit{cla.Army, NATO},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
		"len": USSR,
		"alb": USSR,
		"hav": USSR,
		"mos": USSR,
		"sha": USSR,
		"vla": USSR,
		"lon": NATO,
		"ist": NATO,
		"aus": NATO,
		"new": NATO,
		"los": NATO,
		"par": NATO,
	})
	return
}

func ColdWarGraph() *graph.Graph {
	return graph.New().
		// Tunisia
		Prov("tun").Conn("nof", cla.Coast...).Conn("lib", cla.Coast...).Conn("ion", cla.Sea).Conn("wtm", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// North Vietnam
		Prov("nov").Conn("sai", cla.Coast...).Conn("soc", cla.Sea).Conn("sha", cla.Coast...).Conn("ban", cla.Land).Conn("sta", cla.Coast...).Flag(cla.Coast...).
		// Albania
		Prov("alb").Conn("ion", cla.Sea).Conn("grc", cla.Coast...).Conn("yug", cla.Coast...).Flag(cla.Coast...).SC(USSR).
		// Iran
		Prov("irn").Conn("arm", cla.Land).Conn("irq", cla.Coast...).Conn("arb", cla.Sea).Conn("pak", cla.Coast...).Conn("afg", cla.Land).Conn("ura", cla.Land).Conn("cau", cla.Land).Conn("cau", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Florida
		Prov("flo").Conn("wel", cla.Sea).Conn("new", cla.Coast...).Conn("mid", cla.Land).Conn("sow", cla.Coast...).Conn("gum", cla.Sea).Conn("car", cla.Sea).Flag(cla.Coast...).
		// London
		Prov("lon").Conn("nts", cla.Sea).Conn("nwe", cla.Sea).Conn("nts", cla.Sea).Flag(cla.Coast...).SC(NATO).
		// Afghanistan
		Prov("afg").Conn("pak", cla.Land).Conn("sib", cla.Land).Conn("ura", cla.Land).Conn("irn", cla.Land).Flag(cla.Land).
		// Midwest
		Prov("mid").Conn("new", cla.Land).Conn("tor", cla.Land).Conn("wtn", cla.Land).Conn("los", cla.Land).Conn("sow", cla.Land).Conn("flo", cla.Land).Flag(cla.Land).
		// Levant
		Prov("lev").Conn("etm", cla.Sea).Conn("egy", cla.Coast...).Conn("ara", cla.Land).Conn("irq", cla.Land).Conn("arm", cla.Land).Conn("ist", cla.Coast...).Flag(cla.Coast...).
		// North Korea
		Prov("nok").Conn("yel", cla.Sea).Conn("seo", cla.Coast...).Conn("soj", cla.Sea).Conn("vla", cla.Coast...).Conn("man", cla.Coast...).Flag(cla.Coast...).
		// India
		Prov("ind").Conn("ban", cla.Coast...).Conn("pak", cla.Coast...).Conn("arb", cla.Sea).Conn("inc", cla.Sea).Conn("bay", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// New York
		Prov("new").Conn("que", cla.Coast...).Conn("tor", cla.Land).Conn("mid", cla.Land).Conn("flo", cla.Coast...).Conn("wel", cla.Sea).Flag(cla.Coast...).SC(NATO).
		// Venezuala
		Prov("ven").Conn("col", cla.Coast...).Conn("bra", cla.Coast...).Conn("wel", cla.Sea).Conn("car", cla.Sea).Flag(cla.Coast...).
		// Caribbean Sea
		Prov("car").Conn("hav", cla.Sea).Conn("gum", cla.Sea).Conn("mex", cla.Sea).Conn("cen", cla.Sea).Conn("pan", cla.Sea).Conn("col", cla.Sea).Conn("ven", cla.Sea).Conn("wel", cla.Sea).Conn("flo", cla.Sea).Conn("gum", cla.Sea).Conn("hav", cla.Sea).Flag(cla.Sea).
		// Greenland
		Prov("grd").Conn("arc", cla.Sea).Conn("wel", cla.Sea).Conn("nwe", cla.Sea).Flag(cla.Coast...).
		// Paris
		Prov("par").Conn("wtm", cla.Sea).Conn("ita", cla.Coast...).Conn("weg", cla.Coast...).Conn("nts", cla.Sea).Conn("eal", cla.Sea).Conn("spa", cla.Coast...).Flag(cla.Coast...).SC(NATO).
		// Ionian Sea
		Prov("ion").Conn("grc", cla.Sea).Conn("alb", cla.Sea).Conn("yug", cla.Sea).Conn("ita", cla.Sea).Conn("wtm", cla.Sea).Conn("tun", cla.Sea).Conn("lib", cla.Sea).Conn("etm", cla.Sea).Flag(cla.Sea).
		// Brazil
		Prov("bra").Conn("wel", cla.Sea).Conn("ven", cla.Coast...).Conn("col", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Gulf of Mexico
		Prov("gum").Conn("mex", cla.Sea).Conn("car", cla.Sea).Conn("hav", cla.Sea).Conn("car", cla.Sea).Conn("flo", cla.Sea).Conn("sow", cla.Sea).Flag(cla.Sea).
		// West Atlantic
		Prov("wel").Conn("eal", cla.Sea).Conn("nwe", cla.Sea).Conn("grd", cla.Sea).Conn("arc", cla.Sea).Conn("hud", cla.Sea).Conn("que", cla.Sea).Conn("new", cla.Sea).Conn("flo", cla.Sea).Conn("car", cla.Sea).Conn("ven", cla.Sea).Conn("bra", cla.Sea).Flag(cla.Sea).
		// West China
		Prov("weh").Conn("mon", cla.Land).Conn("sib", cla.Land).Conn("pak", cla.Land).Conn("ban", cla.Land).Conn("sha", cla.Land).Flag(cla.Land).
		// Havana
		Prov("hav").Conn("car", cla.Sea).Conn("car", cla.Sea).Conn("gum", cla.Sea).Flag(cla.Coast...).SC(USSR).
		// Arabia
		Prov("ara").Conn("egy", cla.Coast...).Conn("red", cla.Sea).Conn("arb", cla.Sea).Conn("irq", cla.Coast...).Conn("lev", cla.Land).Flag(cla.Coast...).
		// East Germany
		Prov("eag").Conn("weg", cla.Land).Conn("yug", cla.Land).Conn("ukr", cla.Land).Conn("mos", cla.Land).Conn("fin", cla.Coast...).Conn("bal", cla.Sea).Conn("den", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Leningrad
		Prov("len").Conn("bal", cla.Sea).Conn("fin", cla.Coast...).Conn("nwe", cla.Sea).Conn("noy", cla.Coast...).Conn("swe", cla.Coast...).Flag(cla.Coast...).
		// North Africa
		Prov("nof").Conn("lib", cla.Land).Conn("tun", cla.Coast...).Conn("wtm", cla.Sea).Conn("eal", cla.Sea).Flag(cla.Coast...).
		// Baltic Sea
		Prov("bal").Conn("len", cla.Sea).Conn("swe", cla.Sea).Conn("den", cla.Sea).Conn("eag", cla.Sea).Conn("fin", cla.Sea).Flag(cla.Sea).
		// Yugoslavia
		Prov("yug").Conn("weg", cla.Land).Conn("ita", cla.Coast...).Conn("ion", cla.Sea).Conn("alb", cla.Coast...).Conn("grc", cla.Coast...).Conn("ukr", cla.Land).Conn("eag", cla.Land).Flag(cla.Coast...).
		// Toronto
		Prov("tor").Conn("new", cla.Land).Conn("que", cla.Coast...).Conn("hud", cla.Sea).Conn("wtn", cla.Coast...).Conn("mid", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Norway
		Prov("noy").Conn("nwe", cla.Sea).Conn("nts", cla.Sea).Conn("swe", cla.Coast...).Conn("len", cla.Coast...).Flag(cla.Coast...).
		// Vladivostok
		Prov("vla").Conn("man", cla.Land).Conn("nok", cla.Coast...).Conn("soj", cla.Sea).Conn("ber", cla.Sea).Conn("kam", cla.Coast...).Conn("sib", cla.Land).Flag(cla.Coast...).SC(USSR).
		// East Africa
		Prov("etf").Conn("inc", cla.Sea).Conn("red", cla.Sea).Conn("egy", cla.Coast...).Conn("lib", cla.Land).Flag(cla.Coast...).
		// Libya
		Prov("lib").Conn("etf", cla.Land).Conn("egy", cla.Coast...).Conn("etm", cla.Sea).Conn("ion", cla.Sea).Conn("tun", cla.Coast...).Conn("nof", cla.Land).Flag(cla.Coast...).
		// Japan
		Prov("jap").Conn("soj", cla.Sea).Conn("yel", cla.Sea).Conn("wep", cla.Sea).Conn("ber", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Denmark
		Prov("den").Conn("bal", cla.Sea).Conn("nts", cla.Sea).Conn("weg", cla.Coast...).Conn("eag", cla.Coast...).Flag(cla.Coast...).
		// Seoul
		Prov("seo").Conn("nok", cla.Coast...).Conn("yel", cla.Sea).Conn("soj", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// Bering Sea
		Prov("ber").Conn("eap", cla.Sea).Conn("gus", cla.Sea).Conn("ala", cla.Sea).Conn("arc", cla.Sea).Conn("kam", cla.Sea).Conn("vla", cla.Sea).Conn("soj", cla.Sea).Conn("jap", cla.Sea).Conn("wep", cla.Sea).Flag(cla.Sea).
		// Los Angeles
		Prov("los").Conn("wtn", cla.Coast...).Conn("gus", cla.Sea).Conn("eap", cla.Sea).Conn("mex", cla.Coast...).Conn("sow", cla.Land).Conn("mid", cla.Land).Flag(cla.Coast...).SC(NATO).
		// Caucasus
		Prov("cau").Conn("bla", cla.Sea).Conn("arm", cla.Coast...).Conn("irn", cla.Land).Conn("irn", cla.Land).Conn("ura", cla.Land).Conn("mos", cla.Land).Conn("ukr", cla.Coast...).Flag(cla.Coast...).
		// Armenia
		Prov("arm").Conn("irq", cla.Land).Conn("irn", cla.Land).Conn("cau", cla.Coast...).Conn("bla", cla.Sea).Conn("ist", cla.Coast...).Conn("lev", cla.Land).Flag(cla.Coast...).
		// Panama
		Prov("pan").Conn("col", cla.Coast...).Conn("car", cla.Sea).Conn("cen", cla.Coast...).Conn("eap", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// South West
		Prov("sow").Conn("mid", cla.Land).Conn("los", cla.Land).Conn("mex", cla.Coast...).Conn("gum", cla.Sea).Conn("flo", cla.Coast...).Flag(cla.Coast...).
		// South China Sea
		Prov("soc").Conn("sai", cla.Sea).Conn("sta", cla.Sea).Conn("bay", cla.Sea).Conn("ins", cla.Sea).Conn("phi", cla.Sea).Conn("yel", cla.Sea).Conn("sha", cla.Sea).Conn("nov", cla.Sea).Flag(cla.Sea).
		// Istanbul
		Prov("ist").Conn("grc", cla.Coast...).Conn("etm", cla.Sea).Conn("lev", cla.Coast...).Conn("arm", cla.Coast...).Conn("bla", cla.Sea).Conn("ukr", cla.Coast...).Flag(cla.Coast...).SC(NATO).
		// Arabian Sea
		Prov("arb").Conn("irq", cla.Sea).Conn("ara", cla.Sea).Conn("red", cla.Sea).Conn("inc", cla.Sea).Conn("ind", cla.Sea).Conn("pak", cla.Sea).Conn("irn", cla.Sea).Flag(cla.Sea).
		// Finland
		Prov("fin").Conn("len", cla.Coast...).Conn("bal", cla.Sea).Conn("eag", cla.Coast...).Conn("mos", cla.Land).Conn("ura", cla.Coast...).Conn("nwe", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// East Mediterranean
		Prov("etm").Conn("lev", cla.Sea).Conn("ist", cla.Sea).Conn("grc", cla.Sea).Conn("ion", cla.Sea).Conn("lib", cla.Sea).Conn("egy", cla.Sea).Flag(cla.Sea).
		// North Sea
		Prov("nts").Conn("swe", cla.Sea).Conn("noy", cla.Sea).Conn("nwe", cla.Sea).Conn("lon", cla.Sea).Conn("lon", cla.Sea).Conn("nwe", cla.Sea).Conn("eal", cla.Sea).Conn("par", cla.Sea).Conn("weg", cla.Sea).Conn("den", cla.Sea).Flag(cla.Sea).
		// Urals
		Prov("ura").Conn("nwe", cla.Sea).Conn("fin", cla.Coast...).Conn("mos", cla.Land).Conn("cau", cla.Land).Conn("irn", cla.Land).Conn("afg", cla.Land).Conn("sib", cla.Coast...).Conn("arc", cla.Sea).Flag(cla.Coast...).
		// Manchuria
		Prov("man").Conn("vla", cla.Land).Conn("sib", cla.Land).Conn("mon", cla.Land).Conn("sha", cla.Coast...).Conn("yel", cla.Sea).Conn("nok", cla.Coast...).Flag(cla.Coast...).
		// East Atlantic
		Prov("eal").Conn("nof", cla.Sea).Conn("wtm", cla.Sea).Conn("spa", cla.Sea).Conn("par", cla.Sea).Conn("nts", cla.Sea).Conn("nwe", cla.Sea).Conn("wel", cla.Sea).Flag(cla.Sea).
		// Alaska
		Prov("ala").Conn("arc", cla.Sea).Conn("ber", cla.Sea).Conn("gus", cla.Sea).Conn("wtn", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Bay of Bengal
		Prov("bay").Conn("ins", cla.Sea).Conn("soc", cla.Sea).Conn("sta", cla.Sea).Conn("ban", cla.Sea).Conn("ind", cla.Sea).Conn("inc", cla.Sea).Conn("inc", cla.Sea).Flag(cla.Sea).
		// Ukraine
		Prov("ukr").Conn("cau", cla.Coast...).Conn("mos", cla.Land).Conn("eag", cla.Land).Conn("yug", cla.Land).Conn("grc", cla.Land).Conn("ist", cla.Coast...).Conn("bla", cla.Sea).Flag(cla.Coast...).
		// Saigon
		Prov("sai").Conn("soc", cla.Sea).Conn("nov", cla.Coast...).Conn("sta", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Bangladesh
		Prov("ban").Conn("sha", cla.Land).Conn("weh", cla.Land).Conn("ind", cla.Coast...).Conn("bay", cla.Sea).Conn("sta", cla.Coast...).Conn("nov", cla.Land).Flag(cla.Coast...).
		// Sea of Japan
		Prov("soj").Conn("nok", cla.Sea).Conn("seo", cla.Sea).Conn("yel", cla.Sea).Conn("jap", cla.Sea).Conn("ber", cla.Sea).Conn("vla", cla.Sea).Flag(cla.Sea).
		// East Pacific
		Prov("eap").Conn("col", cla.Sea).Conn("pan", cla.Sea).Conn("cen", cla.Sea).Conn("mex", cla.Sea).Conn("los", cla.Sea).Conn("gus", cla.Sea).Conn("ber", cla.Sea).Conn("wep", cla.Sea).Flag(cla.Sea).
		// Spain
		Prov("spa").Conn("wtm", cla.Sea).Conn("par", cla.Coast...).Conn("eal", cla.Sea).Flag(cla.Coast...).
		// Indian Ocean
		Prov("inc").Conn("aus", cla.Sea).Conn("ins", cla.Sea).Conn("bay", cla.Sea).Conn("bay", cla.Sea).Conn("ind", cla.Sea).Conn("arb", cla.Sea).Conn("red", cla.Sea).Conn("etf", cla.Sea).Flag(cla.Sea).
		// Norwegian Sea
		Prov("nwe").Conn("ura", cla.Sea).Conn("arc", cla.Sea).Conn("grd", cla.Sea).Conn("wel", cla.Sea).Conn("eal", cla.Sea).Conn("nts", cla.Sea).Conn("lon", cla.Sea).Conn("nts", cla.Sea).Conn("noy", cla.Sea).Conn("len", cla.Sea).Conn("fin", cla.Sea).Flag(cla.Sea).
		// Hudson Bay
		Prov("hud").Conn("arc", cla.Sea).Conn("wtn", cla.Sea).Conn("tor", cla.Sea).Conn("que", cla.Sea).Conn("wel", cla.Sea).Flag(cla.Sea).
		// Philippeans
		Prov("phi").Conn("yel", cla.Sea).Conn("soc", cla.Sea).Conn("ins", cla.Coast...).Conn("wep", cla.Sea).Flag(cla.Coast...).
		// Mongolia
		Prov("mon").Conn("weh", cla.Land).Conn("sha", cla.Land).Conn("man", cla.Land).Conn("sib", cla.Land).Flag(cla.Land).
		// Yellow Sea
		Prov("yel").Conn("wep", cla.Sea).Conn("jap", cla.Sea).Conn("soj", cla.Sea).Conn("seo", cla.Sea).Conn("nok", cla.Sea).Conn("man", cla.Sea).Conn("sha", cla.Sea).Conn("soc", cla.Sea).Conn("phi", cla.Sea).Flag(cla.Sea).
		// West Germany
		Prov("weg").Conn("eag", cla.Land).Conn("den", cla.Coast...).Conn("nts", cla.Sea).Conn("par", cla.Coast...).Conn("yug", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// Greece
		Prov("grc").Conn("ion", cla.Sea).Conn("etm", cla.Sea).Conn("ist", cla.Coast...).Conn("ukr", cla.Land).Conn("yug", cla.Coast...).Conn("alb", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Arctic Ocean
		Prov("arc").Conn("grd", cla.Sea).Conn("nwe", cla.Sea).Conn("ura", cla.Sea).Conn("sib", cla.Sea).Conn("kam", cla.Sea).Conn("ber", cla.Sea).Conn("ala", cla.Sea).Conn("wtn", cla.Sea).Conn("hud", cla.Sea).Conn("wel", cla.Sea).Flag(cla.Sea).
		// Sweden
		Prov("swe").Conn("nts", cla.Sea).Conn("bal", cla.Sea).Conn("len", cla.Coast...).Conn("noy", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Iraq
		Prov("irq").Conn("arb", cla.Sea).Conn("irn", cla.Coast...).Conn("arm", cla.Land).Conn("lev", cla.Land).Conn("ara", cla.Coast...).Flag(cla.Coast...).
		// Pakistan
		Prov("pak").Conn("arb", cla.Sea).Conn("ind", cla.Coast...).Conn("weh", cla.Land).Conn("sib", cla.Land).Conn("afg", cla.Land).Conn("irn", cla.Coast...).Flag(cla.Coast...).
		// Shanghai
		Prov("sha").Conn("ban", cla.Land).Conn("nov", cla.Coast...).Conn("soc", cla.Sea).Conn("yel", cla.Sea).Conn("man", cla.Coast...).Conn("mon", cla.Land).Conn("weh", cla.Land).Flag(cla.Coast...).SC(USSR).
		// Mexico
		Prov("mex").Conn("gum", cla.Sea).Conn("sow", cla.Coast...).Conn("los", cla.Coast...).Conn("eap", cla.Sea).Conn("cen", cla.Coast...).Conn("car", cla.Sea).Flag(cla.Coast...).
		// West Canada
		Prov("wtn").Conn("los", cla.Coast...).Conn("mid", cla.Land).Conn("tor", cla.Coast...).Conn("hud", cla.Sea).Conn("arc", cla.Sea).Conn("ala", cla.Coast...).Conn("gus", cla.Sea).Flag(cla.Coast...).
		// West Pacific
		Prov("wep").Conn("eap", cla.Sea).Conn("ber", cla.Sea).Conn("jap", cla.Sea).Conn("yel", cla.Sea).Conn("phi", cla.Sea).Conn("ins", cla.Sea).Conn("ins", cla.Sea).Conn("aus", cla.Sea).Flag(cla.Sea).
		// Black Sea
		Prov("bla").Conn("cau", cla.Sea).Conn("ukr", cla.Sea).Conn("ist", cla.Sea).Conn("arm", cla.Sea).Flag(cla.Sea).
		// Egypt
		Prov("egy").Conn("red", cla.Sea).Conn("ara", cla.Coast...).Conn("lev", cla.Coast...).Conn("etm", cla.Sea).Conn("lib", cla.Coast...).Conn("etf", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// Central America
		Prov("cen").Conn("eap", cla.Sea).Conn("pan", cla.Coast...).Conn("car", cla.Sea).Conn("mex", cla.Coast...).Flag(cla.Coast...).
		// Red Sea
		Prov("red").Conn("egy", cla.Sea).Conn("etf", cla.Sea).Conn("inc", cla.Sea).Conn("arb", cla.Sea).Conn("ara", cla.Sea).Flag(cla.Sea).
		// Australia
		Prov("aus").Conn("wep", cla.Sea).Conn("ins", cla.Coast...).Conn("inc", cla.Sea).Flag(cla.Coast...).SC(NATO).
		// Siberia
		Prov("sib").Conn("pak", cla.Land).Conn("weh", cla.Land).Conn("mon", cla.Land).Conn("man", cla.Land).Conn("vla", cla.Land).Conn("kam", cla.Coast...).Conn("arc", cla.Sea).Conn("ura", cla.Coast...).Conn("afg", cla.Land).Flag(cla.Coast...).
		// Kamchatka
		Prov("kam").Conn("arc", cla.Sea).Conn("sib", cla.Coast...).Conn("vla", cla.Coast...).Conn("ber", cla.Sea).Flag(cla.Coast...).
		// Indonesia
		Prov("ins").Conn("aus", cla.Coast...).Conn("wep", cla.Sea).Conn("wep", cla.Sea).Conn("phi", cla.Coast...).Conn("soc", cla.Sea).Conn("bay", cla.Sea).Conn("inc", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// West Mediterranean
		Prov("wtm").Conn("spa", cla.Sea).Conn("eal", cla.Sea).Conn("nof", cla.Sea).Conn("tun", cla.Sea).Conn("ion", cla.Sea).Conn("ita", cla.Sea).Conn("par", cla.Sea).Flag(cla.Sea).
		// Colombia
		Prov("col").Conn("bra", cla.Land).Conn("ven", cla.Coast...).Conn("car", cla.Sea).Conn("pan", cla.Coast...).Conn("eap", cla.Sea).Flag(cla.Coast...).
		// Quebec
		Prov("que").Conn("new", cla.Coast...).Conn("wel", cla.Sea).Conn("hud", cla.Sea).Conn("tor", cla.Coast...).Flag(cla.Coast...).
		// South East Asia
		Prov("sta").Conn("sai", cla.Coast...).Conn("nov", cla.Coast...).Conn("ban", cla.Coast...).Conn("bay", cla.Sea).Conn("soc", cla.Sea).Flag(cla.Coast...).
		// Italy
		Prov("ita").Conn("par", cla.Coast...).Conn("wtm", cla.Sea).Conn("ion", cla.Sea).Conn("yug", cla.Coast...).Flag(cla.Coast...).
		// Moscow
		Prov("mos").Conn("cau", cla.Land).Conn("ura", cla.Land).Conn("fin", cla.Land).Conn("eag", cla.Land).Conn("ukr", cla.Land).Flag(cla.Land).SC(USSR).
		// Gulf of Alaska
		Prov("gus").Conn("eap", cla.Sea).Conn("los", cla.Sea).Conn("wtn", cla.Sea).Conn("ala", cla.Sea).Conn("ber", cla.Sea).Flag(cla.Sea).
		Done()
}
