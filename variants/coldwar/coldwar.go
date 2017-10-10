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
		"Leningrad": dip.Unit{cla.Fleet, USSR},
		"Albania": dip.Unit{cla.Fleet, USSR},
		"Havana": dip.Unit{cla.Fleet, USSR},
		"Moscow": dip.Unit{cla.Army, USSR},
		"Shanghai": dip.Unit{cla.Army, USSR},
		"Vladivostok": dip.Unit{cla.Army, USSR},
		"London": dip.Unit{cla.Fleet, NATO},
		"Istanbul": dip.Unit{cla.Fleet, NATO},
		"Australia": dip.Unit{cla.Fleet, NATO},
		"New York": dip.Unit{cla.Army, NATO},
		"Los Angeles": dip.Unit{cla.Army, NATO},
		"Paris": dip.Unit{cla.Army, NATO},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
	})
	return
}

func ColdWarGraph() *graph.Graph {
	return graph.New().
		// yel
		Prov("yel").Conn("wep", cla.Sea).Conn("jap", cla.Sea).Conn("soj", cla.Sea).Conn("seo", cla.Sea).Conn("nok", cla.Sea).Conn("man", cla.Sea).Conn("sha", cla.Sea).Conn("sai", cla.Sea).Conn("phi", cla.Sea).Flag(cla.Sea).
		// alb
		Prov("alb").Conn("yug", cla.Land).Conn("ion", cla.Sea).Conn("grc", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// ita
		Prov("ita").Conn("par", cla.Land).Conn("wtm", cla.Sea).Conn("ion", cla.Sea).Conn("yug", cla.Land).Flag(cla.Coast...).
		// nwe
		Prov("nwe").Conn("nts", cla.Sea).Conn("lon", cla.Sea).Conn("nts", cla.Sea).Conn("noy", cla.Sea).Conn("len", cla.Sea).Conn("fin", cla.Sea).Conn("ura", cla.Sea).Conn("arc", cla.Sea).Conn("grd", cla.Sea).Conn("wel", cla.Sea).Conn("eal", cla.Sea).Flag(cla.Sea).
		// weg
		Prov("weg").Conn("eag", cla.Land).Conn("den", cla.Land).Conn("nts", cla.Sea).Conn("par", cla.Land).Conn("yug", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// weh
		Prov("weh").Conn("ban", cla.Land).Conn("sha", cla.Land).Conn("mon", cla.Land).Conn("sib", cla.Land).Conn("pak", cla.Land).Flag(cla.Land).
		// jap
		Prov("jap").Conn("soj", cla.Sea).Conn("yel", cla.Sea).Conn("wep", cla.Sea).Conn("ber", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// wel
		Prov("wel").Conn("eal", cla.Sea).Conn("nwe", cla.Sea).Conn("grd", cla.Sea).Conn("arc", cla.Sea).Conn("hud", cla.Sea).Conn("que", cla.Sea).Conn("new", cla.Sea).Conn("flo", cla.Sea).Conn("hav", cla.Sea).Conn("car", cla.Sea).Conn("bra", cla.Sea).Flag(cla.Sea).
		// aus
		Prov("aus").Conn("wep", cla.Sea).Conn("ins", cla.Land).Conn("arb", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// wep
		Prov("wep").Conn("eap", cla.Sea).Conn("ber", cla.Sea).Conn("jap", cla.Sea).Conn("yel", cla.Sea).Conn("phi", cla.Sea).Conn("ins", cla.Sea).Conn("ins", cla.Sea).Conn("aus", cla.Sea).Flag(cla.Sea).
		// ven
		Prov("ven").Conn("bra", cla.Land).Conn("car", cla.Sea).Conn("hav", cla.Land).Conn("pan", cla.Land).Conn("eap", cla.Sea).Flag(cla.Coast...).
		// fin
		Prov("fin").Conn("mos", cla.Land).Conn("ura", cla.Land).Conn("nwe", cla.Sea).Conn("len", cla.Land).Conn("bal", cla.Sea).Conn("eag", cla.Land).Flag(cla.Coast...).
		// hav
		Prov("hav").Conn("wel", cla.Sea).Conn("flo", cla.Land).Conn("gum", cla.Sea).Conn("col", cla.Land).Conn("col", cla.Land).Conn("gum", cla.Sea).Conn("sow", cla.Land).Conn("cen", cla.Land).Conn("pan", cla.Land).Conn("ven", cla.Land).Conn("car", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// nok
		Prov("nok").Conn("soj", cla.Sea).Conn("vla", cla.Land).Conn("man", cla.Land).Conn("yel", cla.Sea).Conn("seo", cla.Land).Flag(cla.Coast...).
		// mos
		Prov("mos").Conn("cau", cla.Land).Conn("ura", cla.Land).Conn("fin", cla.Land).Conn("eag", cla.Land).Conn("ukr", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// cen
		Prov("cen").Conn("pan", cla.Land).Conn("hav", cla.Land).Conn("sow", cla.Land).Conn("eap", cla.Sea).Flag(cla.Coast...).
		// nof
		Prov("nof").Conn("lib", cla.Land).Conn("tun", cla.Land).Conn("wtm", cla.Sea).Conn("eal", cla.Sea).Flag(cla.Coast...).
		// mon
		Prov("mon").Conn("weh", cla.Land).Conn("sha", cla.Land).Conn("man", cla.Land).Conn("sib", cla.Land).Flag(cla.Land).
		// nov
		Prov("nov").Conn("sai", cla.Land).Conn("sha", cla.Land).Conn("ban", cla.Land).Conn("sta", cla.Land).Conn("soc", cla.Sea).Flag(cla.Coast...).
		// ala
		Prov("ala").Conn("ber", cla.Sea).Conn("gus", cla.Sea).Conn("wtn", cla.Land).Conn("arc", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// gum
		Prov("gum").Conn("sow", cla.Sea).Conn("hav", cla.Sea).Conn("col", cla.Sea).Conn("hav", cla.Sea).Conn("flo", cla.Sea).Conn("mex", cla.Sea).Flag(cla.Sea).
		// bay
		Prov("bay").Conn("ind", cla.Sea).Conn("arb", cla.Sea).Conn("arb", cla.Sea).Conn("ins", cla.Sea).Conn("sai", cla.Sea).Conn("sta", cla.Sea).Conn("ban", cla.Sea).Flag(cla.Sea).
		// gus
		Prov("gus").Conn("ber", cla.Sea).Conn("eap", cla.Sea).Conn("los", cla.Sea).Conn("wtn", cla.Sea).Conn("ala", cla.Sea).Flag(cla.Sea).
		// sha
		Prov("sha").Conn("man", cla.Land).Conn("mon", cla.Land).Conn("weh", cla.Land).Conn("ban", cla.Land).Conn("nov", cla.Land).Conn("sai", cla.Land).Conn("yel", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// ban
		Prov("ban").Conn("ind", cla.Land).Conn("bay", cla.Sea).Conn("sta", cla.Land).Conn("nov", cla.Land).Conn("sha", cla.Land).Conn("weh", cla.Land).Flag(cla.Coast...).
		// bal
		Prov("bal").Conn("len", cla.Sea).Conn("swe", cla.Sea).Conn("den", cla.Sea).Conn("eag", cla.Sea).Conn("fin", cla.Sea).Flag(cla.Sea).
		// ara
		Prov("ara").Conn("red", cla.Sea).Conn("inc", cla.Sea).Conn("irq", cla.Land).Conn("lev", cla.Land).Conn("egy", cla.Land).Flag(cla.Coast...).
		// arb
		Prov("arb").Conn("aus", cla.Sea).Conn("ins", cla.Sea).Conn("bay", cla.Sea).Conn("bay", cla.Sea).Conn("ind", cla.Sea).Conn("inc", cla.Sea).Conn("red", cla.Sea).Conn("etf", cla.Sea).Flag(cla.Sea).
		// arc
		Prov("arc").Conn("sib", cla.Sea).Conn("kam", cla.Sea).Conn("ber", cla.Sea).Conn("ala", cla.Sea).Conn("wtn", cla.Sea).Conn("hud", cla.Sea).Conn("wel", cla.Sea).Conn("grd", cla.Sea).Conn("nwe", cla.Sea).Conn("ura", cla.Sea).Flag(cla.Sea).
		// seo
		Prov("seo").Conn("yel", cla.Sea).Conn("soj", cla.Sea).Conn("nok", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// arm
		Prov("arm").Conn("cau", cla.Land).Conn("bla", cla.Sea).Conn("ist", cla.Land).Conn("lev", cla.Land).Conn("irq", cla.Land).Conn("irn", cla.Land).Flag(cla.Coast...).
		// irq
		Prov("irq").Conn("irn", cla.Land).Conn("arm", cla.Land).Conn("lev", cla.Land).Conn("ara", cla.Land).Conn("inc", cla.Sea).Flag(cla.Coast...).
		// irn
		Prov("irn").Conn("cau", cla.Land).Conn("cau", cla.Land).Conn("arm", cla.Land).Conn("irq", cla.Land).Conn("inc", cla.Sea).Conn("pak", cla.Land).Conn("afg", cla.Land).Conn("ura", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// new
		Prov("new").Conn("mid", cla.Land).Conn("flo", cla.Land).Conn("wel", cla.Sea).Conn("que", cla.Land).Conn("tor", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// bla
		Prov("bla").Conn("ist", cla.Sea).Conn("arm", cla.Sea).Conn("cau", cla.Sea).Conn("ukr", cla.Sea).Flag(cla.Sea).
		// red
		Prov("red").Conn("ara", cla.Sea).Conn("egy", cla.Sea).Conn("etf", cla.Sea).Conn("arb", cla.Sea).Conn("inc", cla.Sea).Flag(cla.Sea).
		// flo
		Prov("flo").Conn("mid", cla.Land).Conn("mex", cla.Land).Conn("gum", cla.Sea).Conn("hav", cla.Land).Conn("wel", cla.Sea).Conn("new", cla.Land).Flag(cla.Coast...).
		// nts
		Prov("nts").Conn("par", cla.Sea).Conn("weg", cla.Sea).Conn("den", cla.Sea).Conn("swe", cla.Sea).Conn("noy", cla.Sea).Conn("nwe", cla.Sea).Conn("lon", cla.Sea).Conn("lon", cla.Sea).Conn("nwe", cla.Sea).Conn("eal", cla.Sea).Flag(cla.Sea).
		// len
		Prov("len").Conn("bal", cla.Sea).Conn("fin", cla.Land).Conn("nwe", cla.Sea).Conn("noy", cla.Land).Conn("swe", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// que
		Prov("que").Conn("new", cla.Land).Conn("wel", cla.Sea).Conn("hud", cla.Sea).Conn("tor", cla.Land).Flag(cla.Coast...).
		// lev
		Prov("lev").Conn("ist", cla.Land).Conn("etm", cla.Sea).Conn("egy", cla.Land).Conn("ara", cla.Land).Conn("irq", cla.Land).Conn("arm", cla.Land).Flag(cla.Coast...).
		// wtm
		Prov("wtm").Conn("par", cla.Sea).Conn("spa", cla.Sea).Conn("eal", cla.Sea).Conn("nof", cla.Sea).Conn("tun", cla.Sea).Conn("ion", cla.Sea).Conn("ita", cla.Sea).Flag(cla.Sea).
		// wtn
		Prov("wtn").Conn("hud", cla.Sea).Conn("arc", cla.Sea).Conn("ala", cla.Land).Conn("gus", cla.Sea).Conn("los", cla.Land).Conn("mid", cla.Land).Conn("tor", cla.Land).Flag(cla.Coast...).
		// ist
		Prov("ist").Conn("lev", cla.Land).Conn("arm", cla.Land).Conn("bla", cla.Sea).Conn("ukr", cla.Land).Conn("grc", cla.Land).Conn("etm", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// etf
		Prov("etf").Conn("arb", cla.Sea).Conn("red", cla.Sea).Conn("egy", cla.Land).Conn("lib", cla.Land).Flag(cla.Coast...).
		// col
		Prov("col").Conn("hav", cla.Land).Conn("gum", cla.Sea).Conn("hav", cla.Land).Flag(cla.Coast...).
		// yug
		Prov("yug").Conn("alb", cla.Land).Conn("grc", cla.Land).Conn("ukr", cla.Land).Conn("eag", cla.Land).Conn("weg", cla.Land).Conn("ita", cla.Land).Conn("ion", cla.Sea).Flag(cla.Coast...).
		// etm
		Prov("etm").Conn("grc", cla.Sea).Conn("ion", cla.Sea).Conn("lib", cla.Sea).Conn("egy", cla.Sea).Conn("lev", cla.Sea).Conn("ist", cla.Sea).Flag(cla.Sea).
		// soc
		Prov("soc").Conn("sta", cla.Sea).Conn("sai", cla.Sea).Conn("nov", cla.Sea).Flag(cla.Sea).
		// tun
		Prov("tun").Conn("nof", cla.Land).Conn("lib", cla.Land).Conn("ion", cla.Sea).Conn("wtm", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// lon
		Prov("lon").Conn("nts", cla.Sea).Conn("nwe", cla.Sea).Conn("nts", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// soj
		Prov("soj").Conn("jap", cla.Sea).Conn("ber", cla.Sea).Conn("vla", cla.Sea).Conn("nok", cla.Sea).Conn("seo", cla.Sea).Conn("yel", cla.Sea).Flag(cla.Sea).
		// sow
		Prov("sow").Conn("gum", cla.Sea).Conn("mex", cla.Land).Conn("los", cla.Land).Conn("eap", cla.Sea).Conn("cen", cla.Land).Conn("hav", cla.Land).Flag(cla.Coast...).
		// bra
		Prov("bra").Conn("wel", cla.Sea).Conn("car", cla.Sea).Conn("ven", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// tor
		Prov("tor").Conn("hud", cla.Sea).Conn("wtn", cla.Land).Conn("mid", cla.Land).Conn("new", cla.Land).Conn("que", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// swe
		Prov("swe").Conn("noy", cla.Land).Conn("nts", cla.Sea).Conn("bal", cla.Sea).Conn("len", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// ukr
		Prov("ukr").Conn("ist", cla.Land).Conn("bla", cla.Sea).Conn("cau", cla.Land).Conn("mos", cla.Land).Conn("eag", cla.Land).Conn("yug", cla.Land).Conn("grc", cla.Land).Flag(cla.Coast...).
		// mex
		Prov("mex").Conn("gum", cla.Sea).Conn("flo", cla.Land).Conn("mid", cla.Land).Conn("los", cla.Land).Conn("sow", cla.Land).Flag(cla.Coast...).
		// los
		Prov("los").Conn("eap", cla.Sea).Conn("sow", cla.Land).Conn("mex", cla.Land).Conn("mid", cla.Land).Conn("wtn", cla.Land).Conn("gus", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// vla
		Prov("vla").Conn("ber", cla.Sea).Conn("kam", cla.Land).Conn("sib", cla.Land).Conn("man", cla.Land).Conn("nok", cla.Land).Conn("soj", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// hud
		Prov("hud").Conn("wtn", cla.Sea).Conn("tor", cla.Sea).Conn("que", cla.Sea).Conn("wel", cla.Sea).Conn("arc", cla.Sea).Flag(cla.Sea).
		// eal
		Prov("eal").Conn("nof", cla.Sea).Conn("wtm", cla.Sea).Conn("spa", cla.Sea).Conn("par", cla.Sea).Conn("nts", cla.Sea).Conn("nwe", cla.Sea).Conn("wel", cla.Sea).Flag(cla.Sea).
		// eag
		Prov("eag").Conn("weg", cla.Land).Conn("yug", cla.Land).Conn("ukr", cla.Land).Conn("mos", cla.Land).Conn("fin", cla.Land).Conn("bal", cla.Sea).Conn("den", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// eap
		Prov("eap").Conn("ven", cla.Sea).Conn("pan", cla.Sea).Conn("cen", cla.Sea).Conn("sow", cla.Sea).Conn("los", cla.Sea).Conn("gus", cla.Sea).Conn("ber", cla.Sea).Conn("wep", cla.Sea).Flag(cla.Sea).
		// sta
		Prov("sta").Conn("ban", cla.Land).Conn("bay", cla.Sea).Conn("sai", cla.Land).Conn("soc", cla.Sea).Conn("nov", cla.Land).Flag(cla.Coast...).
		// car
		Prov("car").Conn("wel", cla.Sea).Conn("hav", cla.Sea).Conn("ven", cla.Sea).Conn("bra", cla.Sea).Flag(cla.Sea).
		// cau
		Prov("cau").Conn("irn", cla.Land).Conn("ura", cla.Land).Conn("mos", cla.Land).Conn("ukr", cla.Land).Conn("bla", cla.Sea).Conn("arm", cla.Land).Conn("irn", cla.Land).Flag(cla.Coast...).
		// den
		Prov("den").Conn("nts", cla.Sea).Conn("weg", cla.Land).Conn("eag", cla.Land).Conn("bal", cla.Sea).Flag(cla.Coast...).
		// ber
		Prov("ber").Conn("eap", cla.Sea).Conn("gus", cla.Sea).Conn("ala", cla.Sea).Conn("arc", cla.Sea).Conn("kam", cla.Sea).Conn("vla", cla.Sea).Conn("soj", cla.Sea).Conn("jap", cla.Sea).Conn("wep", cla.Sea).Flag(cla.Sea).
		// sai
		Prov("sai").Conn("nov", cla.Land).Conn("soc", cla.Sea).Conn("sta", cla.Land).Conn("bay", cla.Sea).Conn("ins", cla.Land).Conn("phi", cla.Land).Conn("yel", cla.Sea).Conn("sha", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// ins
		Prov("ins").Conn("phi", cla.Land).Conn("sai", cla.Land).Conn("bay", cla.Sea).Conn("arb", cla.Sea).Conn("aus", cla.Land).Conn("wep", cla.Sea).Conn("wep", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// ind
		Prov("ind").Conn("pak", cla.Land).Conn("inc", cla.Sea).Conn("arb", cla.Sea).Conn("bay", cla.Sea).Conn("ban", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// spa
		Prov("spa").Conn("par", cla.Land).Conn("eal", cla.Sea).Conn("wtm", cla.Sea).Flag(cla.Coast...).
		// inc
		Prov("inc").Conn("irq", cla.Sea).Conn("ara", cla.Sea).Conn("red", cla.Sea).Conn("arb", cla.Sea).Conn("ind", cla.Sea).Conn("pak", cla.Sea).Conn("irn", cla.Sea).Flag(cla.Sea).
		// par
		Prov("par").Conn("nts", cla.Sea).Conn("eal", cla.Sea).Conn("spa", cla.Land).Conn("wtm", cla.Sea).Conn("ita", cla.Land).Conn("weg", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// lib
		Prov("lib").Conn("etm", cla.Sea).Conn("ion", cla.Sea).Conn("tun", cla.Land).Conn("nof", cla.Land).Conn("etf", cla.Land).Conn("egy", cla.Land).Flag(cla.Coast...).
		// afg
		Prov("afg").Conn("pak", cla.Land).Conn("sib", cla.Land).Conn("ura", cla.Land).Conn("irn", cla.Land).Flag(cla.Land).
		// mid
		Prov("mid").Conn("flo", cla.Land).Conn("new", cla.Land).Conn("tor", cla.Land).Conn("wtn", cla.Land).Conn("los", cla.Land).Conn("mex", cla.Land).Flag(cla.Land).
		// sib
		Prov("sib").Conn("arc", cla.Sea).Conn("ura", cla.Land).Conn("afg", cla.Land).Conn("pak", cla.Land).Conn("weh", cla.Land).Conn("mon", cla.Land).Conn("man", cla.Land).Conn("vla", cla.Land).Conn("kam", cla.Land).Flag(cla.Coast...).
		// grd
		Prov("grd").Conn("nwe", cla.Sea).Conn("arc", cla.Sea).Conn("wel", cla.Sea).Flag(cla.Coast...).
		// grc
		Prov("grc").Conn("etm", cla.Sea).Conn("ist", cla.Land).Conn("ukr", cla.Land).Conn("yug", cla.Land).Conn("alb", cla.Land).Conn("ion", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// pak
		Prov("pak").Conn("afg", cla.Land).Conn("irn", cla.Land).Conn("inc", cla.Sea).Conn("ind", cla.Land).Conn("weh", cla.Land).Conn("sib", cla.Land).Flag(cla.Coast...).
		// pan
		Prov("pan").Conn("cen", cla.Land).Conn("eap", cla.Sea).Conn("ven", cla.Land).Conn("hav", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// phi
		Prov("phi").Conn("yel", cla.Sea).Conn("sai", cla.Land).Conn("ins", cla.Land).Conn("wep", cla.Sea).Flag(cla.Coast...).
		// kam
		Prov("kam").Conn("sib", cla.Land).Conn("vla", cla.Land).Conn("ber", cla.Sea).Conn("arc", cla.Sea).Flag(cla.Coast...).
		// ion
		Prov("ion").Conn("lib", cla.Sea).Conn("etm", cla.Sea).Conn("grc", cla.Sea).Conn("alb", cla.Sea).Conn("yug", cla.Sea).Conn("ita", cla.Sea).Conn("wtm", cla.Sea).Conn("tun", cla.Sea).Flag(cla.Sea).
		// man
		Prov("man").Conn("yel", cla.Sea).Conn("nok", cla.Land).Conn("vla", cla.Land).Conn("sib", cla.Land).Conn("mon", cla.Land).Conn("sha", cla.Land).Flag(cla.Coast...).
		// egy
		Prov("egy").Conn("lev", cla.Land).Conn("etm", cla.Sea).Conn("lib", cla.Land).Conn("etf", cla.Land).Conn("red", cla.Sea).Conn("ara", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// ura
		Prov("ura").Conn("irn", cla.Land).Conn("afg", cla.Land).Conn("sib", cla.Land).Conn("arc", cla.Sea).Conn("nwe", cla.Sea).Conn("fin", cla.Land).Conn("mos", cla.Land).Conn("cau", cla.Land).Flag(cla.Coast...).
		// noy
		Prov("noy").Conn("nts", cla.Sea).Conn("swe", cla.Land).Conn("len", cla.Land).Conn("nwe", cla.Sea).Flag(cla.Coast...).
		Done()
}
