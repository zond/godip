package chaos

import (
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"
)

const (
	Ankara         godip.Nation = "Ankara"
	Belgium        godip.Nation = "Belgium"
	Berlin         godip.Nation = "Berlin"
	Brest          godip.Nation = "Brest"
	Budapest       godip.Nation = "Budapest"
	Bulgaria       godip.Nation = "Bulgaria"
	Constantinople godip.Nation = "Constantinople"
	Denmark        godip.Nation = "Denmark"
	Edinburgh      godip.Nation = "Edinburgh"
	Greece         godip.Nation = "Greece"
	Holland        godip.Nation = "Holland"
	Kiel           godip.Nation = "Kiel"
	Liverpool      godip.Nation = "Liverpool"
	London         godip.Nation = "London"
	Marseilles     godip.Nation = "Marseilles"
	Moscow         godip.Nation = "Moscow"
	Munich         godip.Nation = "Munich"
	Naples         godip.Nation = "Naples"
	Norway         godip.Nation = "Norway"
	Paris          godip.Nation = "Paris"
	Portugal       godip.Nation = "Portugal"
	Rome           godip.Nation = "Rome"
	Rumania        godip.Nation = "Rumania"
	Serbia         godip.Nation = "Serbia"
	Sevastopol     godip.Nation = "Sevastopol"
	Smyrna         godip.Nation = "Smyrna"
	Spain          godip.Nation = "Spain"
	StPetersburg   godip.Nation = "StPetersburg"
	Sweden         godip.Nation = "Sweden"
	Trieste        godip.Nation = "Trieste"
	Tunis          godip.Nation = "Tunis"
	Venice         godip.Nation = "Venice"
	Vienna         godip.Nation = "Vienna"
	Warsaw         godip.Nation = "Warsaw"
)

var Nations = []godip.Nation{
	Ankara,
	Belgium,
	Berlin,
	Brest,
	Budapest,
	Bulgaria,
	Constantinople,
	Denmark,
	Edinburgh,
	Greece,
	Holland,
	Kiel,
	Liverpool,
	London,
	Marseilles,
	Moscow,
	Munich,
	Naples,
	Norway,
	Paris,
	Portugal,
	Rome,
	Rumania,
	Serbia,
	Sevastopol,
	Smyrna,
	Spain,
	StPetersburg,
	Sweden,
	Trieste,
	Tunis,
	Venice,
	Vienna,
	Warsaw,
}

func Graph() *graph.Graph {
	return graph.New().
		// nat
		Prov("nat").Conn("nrg", godip.Sea).Conn("cly", godip.Sea).Conn("lvp", godip.Sea).Conn("iri", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Sea).
		// nrg
		Prov("nrg").Conn("nat", godip.Sea).Conn("bar", godip.Sea).Conn("nwy", godip.Sea).Conn("nth", godip.Sea).Conn("edi", godip.Sea).Conn("cly", godip.Sea).Flag(godip.Sea).
		// bar
		Prov("bar").Conn("nrg", godip.Sea).Conn("stp/nc", godip.Sea).Conn("nwy", godip.Sea).Conn("stp", godip.Sea).Flag(godip.Sea).
		// stp/nc
		Prov("stp/nc").Conn("bar", godip.Sea).Conn("nwy", godip.Sea).Flag(godip.Sea).
		// stp
		Prov("stp").Conn("fin", godip.Land).Conn("nwy", godip.Land).Conn("mos", godip.Land).Conn("lvn", godip.Land).Flag(godip.Land).Conn("bar", godip.Sea).Conn("bot", godip.Sea).SC(StPetersburg).
		// mos
		Prov("mos").Conn("stp", godip.Land).Conn("sev", godip.Land).Conn("ukr", godip.Land).Conn("war", godip.Land).Conn("lvn", godip.Land).Flag(godip.Land).SC(Moscow).
		// sev
		Prov("sev").Conn("ukr", godip.Land).Conn("mos", godip.Land).Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Conn("rum", godip.Coast...).Flag(godip.Coast...).SC(Sevastopol).
		// arm
		Prov("arm").Conn("ank", godip.Coast...).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Conn("syr", godip.Land).Conn("smy", godip.Land).Flag(godip.Coast...).
		// syr
		Prov("syr").Conn("eas", godip.Sea).Conn("smy", godip.Coast...).Conn("arm", godip.Land).Flag(godip.Coast...).
		// eas
		Prov("eas").Conn("ion", godip.Sea).Conn("aeg", godip.Sea).Conn("smy", godip.Sea).Conn("syr", godip.Sea).Flag(godip.Sea).
		// ion
		Prov("ion").Conn("apu", godip.Sea).Conn("adr", godip.Sea).Conn("tun", godip.Sea).Conn("tys", godip.Sea).Conn("nap", godip.Sea).Conn("alb", godip.Sea).Conn("gre", godip.Sea).Conn("aeg", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// tun
		Prov("tun").Conn("naf", godip.Coast...).Conn("wes", godip.Sea).Conn("tys", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Tunis).
		// naf
		Prov("naf").Conn("mid", godip.Sea).Conn("wes", godip.Sea).Conn("tun", godip.Coast...).Flag(godip.Coast...).
		// mid
		Prov("mid").Conn("wes", godip.Sea).Conn("nat", godip.Sea).Conn("iri", godip.Sea).Conn("eng", godip.Sea).Conn("bre", godip.Sea).Conn("gas", godip.Sea).Conn("spa/nc", godip.Sea).Conn("por", godip.Sea).Conn("spa/sc", godip.Sea).Conn("naf", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// iri
		Prov("iri").Conn("nat", godip.Sea).Conn("lvp", godip.Sea).Conn("wal", godip.Sea).Conn("eng", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Sea).
		// lvp
		Prov("lvp").Conn("iri", godip.Sea).Conn("nat", godip.Sea).Conn("cly", godip.Coast...).Conn("edi", godip.Land).Conn("yor", godip.Land).Conn("wal", godip.Coast...).Flag(godip.Coast...).SC(Liverpool).
		// cly
		Prov("cly").Conn("nat", godip.Sea).Conn("nrg", godip.Sea).Conn("edi", godip.Coast...).Conn("lvp", godip.Coast...).Flag(godip.Coast...).
		// edi
		Prov("edi").Conn("cly", godip.Coast...).Conn("nrg", godip.Sea).Conn("nth", godip.Sea).Conn("yor", godip.Coast...).Conn("lvp", godip.Land).Flag(godip.Coast...).SC(Edinburgh).
		// nth
		Prov("nth").Conn("eng", godip.Sea).Conn("edi", godip.Sea).Conn("nrg", godip.Sea).Conn("nwy", godip.Sea).Conn("ska", godip.Sea).Conn("den", godip.Sea).Conn("hel", godip.Sea).Conn("hol", godip.Sea).Conn("bel", godip.Sea).Conn("lon", godip.Sea).Conn("yor", godip.Sea).Flag(godip.Sea).
		// nwy
		Prov("nwy").Conn("nth", godip.Sea).Conn("nrg", godip.Sea).Conn("bar", godip.Sea).Conn("stp/nc", godip.Sea).Conn("stp", godip.Land).Conn("fin", godip.Land).Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(Norway).
		// stp/sc
		Prov("stp/sc").Conn("bot", godip.Sea).Conn("fin", godip.Sea).Conn("lvn", godip.Sea).Flag(godip.Sea).
		// lvn
		Prov("lvn").Conn("stp", godip.Land).Conn("bal", godip.Sea).Conn("bot", godip.Sea).Conn("stp/sc", godip.Sea).Conn("mos", godip.Land).Conn("war", godip.Land).Conn("pru", godip.Coast...).Flag(godip.Coast...).
		// war
		Prov("war").Conn("sil", godip.Land).Conn("pru", godip.Land).Conn("lvn", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).SC(Warsaw).
		// ukr
		Prov("ukr").Conn("war", godip.Land).Conn("mos", godip.Land).Conn("sev", godip.Land).Conn("rum", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).
		// bla
		Prov("bla").Conn("bul/ec", godip.Sea).Conn("rum", godip.Sea).Conn("sev", godip.Sea).Conn("arm", godip.Sea).Conn("ank", godip.Sea).Conn("con", godip.Sea).Conn("bul", godip.Sea).Flag(godip.Sea).
		// ank
		Prov("ank").Conn("con", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Conn("smy", godip.Land).Flag(godip.Coast...).SC(Ankara).
		// smy
		Prov("smy").Conn("aeg", godip.Sea).Conn("con", godip.Coast...).Conn("ank", godip.Land).Conn("arm", godip.Land).Conn("syr", godip.Coast...).Conn("eas", godip.Sea).Flag(godip.Coast...).SC(Smyrna).
		// aeg
		Prov("aeg").Conn("eas", godip.Sea).Conn("ion", godip.Sea).Conn("gre", godip.Sea).Conn("bul/sc", godip.Sea).Conn("con", godip.Sea).Conn("smy", godip.Sea).Conn("bul", godip.Sea).Flag(godip.Sea).
		// gre
		Prov("gre").Conn("ion", godip.Sea).Conn("alb", godip.Coast...).Conn("ser", godip.Land).Conn("bul", godip.Land).Conn("bul/sc", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(Greece).
		// nap
		Prov("nap").Conn("tys", godip.Sea).Conn("rom", godip.Coast...).Conn("apu", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Naples).
		// tys
		Prov("tys").Conn("wes", godip.Sea).Conn("gol", godip.Sea).Conn("tus", godip.Sea).Conn("rom", godip.Sea).Conn("nap", godip.Sea).Conn("ion", godip.Sea).Conn("tun", godip.Sea).Flag(godip.Sea).
		// wes
		Prov("wes").Conn("mid", godip.Sea).Conn("spa/sc", godip.Sea).Conn("gol", godip.Sea).Conn("tys", godip.Sea).Conn("tun", godip.Sea).Conn("naf", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// spa/sc
		Prov("spa/sc").Conn("mid", godip.Sea).Conn("por", godip.Sea).Conn("mar", godip.Sea).Conn("gol", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// spa
		Prov("spa").Conn("por", godip.Land).Conn("gas", godip.Land).Conn("mar", godip.Land).Conn("mid", godip.Sea).Conn("gol", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Land).SC(Spain).
		// spa/nc
		Prov("spa/nc").Conn("por", godip.Sea).Conn("mid", godip.Sea).Conn("gas", godip.Sea).Flag(godip.Sea).
		// por
		Prov("por").Conn("mid", godip.Sea).Conn("spa/nc", godip.Sea).Conn("spa", godip.Land).Conn("spa/sc", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// gas
		Prov("gas").Conn("mid", godip.Sea).Conn("bre", godip.Coast...).Conn("par", godip.Land).Conn("bur", godip.Land).Conn("mar", godip.Land).Conn("spa", godip.Land).Conn("spa/nc", godip.Sea).Flag(godip.Coast...).
		// bre
		Prov("bre").Conn("mid", godip.Sea).Conn("eng", godip.Sea).Conn("pic", godip.Coast...).Conn("par", godip.Land).Conn("gas", godip.Coast...).Flag(godip.Coast...).SC(Brest).
		// eng
		Prov("eng").Conn("mid", godip.Sea).Conn("iri", godip.Sea).Conn("wal", godip.Sea).Conn("lon", godip.Sea).Conn("nth", godip.Sea).Conn("bel", godip.Sea).Conn("pic", godip.Sea).Conn("bre", godip.Sea).Flag(godip.Sea).
		// wal
		Prov("wal").Conn("iri", godip.Sea).Conn("lvp", godip.Coast...).Conn("yor", godip.Land).Conn("lon", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).
		// yor
		Prov("yor").Conn("lvp", godip.Land).Conn("edi", godip.Coast...).Conn("nth", godip.Sea).Conn("lon", godip.Coast...).Conn("wal", godip.Land).Flag(godip.Coast...).
		// ska
		Prov("ska").Conn("nth", godip.Sea).Conn("nwy", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Flag(godip.Sea).
		// swe
		Prov("swe").Conn("ska", godip.Sea).Conn("nwy", godip.Coast...).Conn("fin", godip.Coast...).Conn("bot", godip.Sea).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Flag(godip.Coast...).SC(Sweden).
		// fin
		Prov("fin").Conn("nwy", godip.Land).Conn("bot", godip.Sea).Conn("swe", godip.Coast...).Conn("stp", godip.Land).Conn("stp/sc", godip.Sea).Flag(godip.Coast...).
		// bot
		Prov("bot").Conn("swe", godip.Sea).Conn("fin", godip.Sea).Conn("stp/sc", godip.Sea).Conn("lvn", godip.Sea).Conn("bal", godip.Sea).Conn("stp", godip.Sea).Flag(godip.Sea).
		// bal
		Prov("bal").Conn("den", godip.Sea).Conn("swe", godip.Sea).Conn("bot", godip.Sea).Conn("lvn", godip.Sea).Conn("pru", godip.Sea).Conn("ber", godip.Sea).Conn("kie", godip.Sea).Flag(godip.Sea).
		// pru
		Prov("pru").Conn("ber", godip.Coast...).Conn("bal", godip.Sea).Conn("lvn", godip.Coast...).Conn("war", godip.Land).Conn("sil", godip.Land).Flag(godip.Coast...).
		// sil
		Prov("sil").Conn("mun", godip.Land).Conn("ber", godip.Land).Conn("pru", godip.Land).Conn("war", godip.Land).Conn("gal", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).
		// gal
		Prov("gal").Conn("boh", godip.Land).Conn("sil", godip.Land).Conn("war", godip.Land).Conn("ukr", godip.Land).Conn("rum", godip.Land).Conn("bud", godip.Land).Conn("vie", godip.Land).Flag(godip.Land).
		// rum
		Prov("rum").Conn("bla", godip.Sea).Conn("bud", godip.Land).Conn("gal", godip.Land).Conn("ukr", godip.Land).Conn("sev", godip.Coast...).Conn("bul/ec", godip.Sea).Conn("bul", godip.Land).Conn("ser", godip.Land).Flag(godip.Coast...).SC(Rumania).
		// bul/ec
		Prov("bul/ec").Conn("rum", godip.Sea).Conn("bla", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
		// bul
		Prov("bul").Conn("ser", godip.Land).Conn("rum", godip.Land).Conn("con", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).Conn("aeg", godip.Sea).Conn("bla", godip.Sea).SC(Bulgaria).
		// con
		Prov("con").Conn("bul/sc", godip.Sea).Conn("bul", godip.Land).Conn("bul/ec", godip.Sea).Conn("bla", godip.Sea).Conn("ank", godip.Coast...).Conn("smy", godip.Coast...).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(Constantinople).
		// bul/sc
		Prov("bul/sc").Conn("gre", godip.Sea).Conn("con", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Sea).
		// ser
		Prov("ser").Conn("tri", godip.Land).Conn("bud", godip.Land).Conn("rum", godip.Land).Conn("bul", godip.Land).Conn("gre", godip.Land).Conn("alb", godip.Land).Flag(godip.Land).SC(Serbia).
		// alb
		Prov("alb").Conn("adr", godip.Sea).Conn("tri", godip.Coast...).Conn("ser", godip.Land).Conn("gre", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).
		// adr
		Prov("adr").Conn("ven", godip.Sea).Conn("tri", godip.Sea).Conn("alb", godip.Sea).Conn("ion", godip.Sea).Conn("apu", godip.Sea).Flag(godip.Sea).
		// apu
		Prov("apu").Conn("rom", godip.Land).Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Coast...).Flag(godip.Coast...).
		// rom
		Prov("rom").Conn("tys", godip.Sea).Conn("tus", godip.Coast...).Conn("ven", godip.Land).Conn("apu", godip.Land).Conn("nap", godip.Coast...).Flag(godip.Coast...).SC(Rome).
		// tus
		Prov("tus").Conn("gol", godip.Sea).Conn("pie", godip.Coast...).Conn("ven", godip.Land).Conn("rom", godip.Coast...).Conn("tys", godip.Sea).Flag(godip.Coast...).
		// gol
		Prov("gol").Conn("spa/sc", godip.Sea).Conn("mar", godip.Sea).Conn("pie", godip.Sea).Conn("tus", godip.Sea).Conn("tys", godip.Sea).Conn("wes", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// mar
		Prov("mar").Conn("spa", godip.Land).Conn("gas", godip.Land).Conn("bur", godip.Land).Conn("pie", godip.Coast...).Conn("gol", godip.Sea).Conn("spa/sc", godip.Sea).Flag(godip.Coast...).SC(Marseilles).
		// bur
		Prov("bur").Conn("par", godip.Land).Conn("pic", godip.Land).Conn("bel", godip.Land).Conn("ruh", godip.Land).Conn("mun", godip.Land).Conn("mar", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).
		// par
		Prov("par").Conn("bre", godip.Land).Conn("pic", godip.Land).Conn("bur", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).SC(Paris).
		// pic
		Prov("pic").Conn("bre", godip.Coast...).Conn("eng", godip.Sea).Conn("bel", godip.Coast...).Conn("bur", godip.Land).Conn("par", godip.Land).Flag(godip.Coast...).
		// lon
		Prov("lon").Conn("wal", godip.Coast...).Conn("yor", godip.Coast...).Conn("nth", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(London).
		// bel
		Prov("bel").Conn("pic", godip.Coast...).Conn("eng", godip.Sea).Conn("nth", godip.Sea).Conn("hol", godip.Coast...).Conn("ruh", godip.Land).Conn("bur", godip.Land).Flag(godip.Coast...).SC(Belgium).
		// hol
		Prov("hol").Conn("nth", godip.Sea).Conn("hel", godip.Sea).Conn("kie", godip.Coast...).Conn("ruh", godip.Land).Conn("bel", godip.Coast...).Flag(godip.Coast...).SC(Holland).
		// hel
		Prov("hel").Conn("nth", godip.Sea).Conn("den", godip.Sea).Conn("kie", godip.Sea).Conn("hol", godip.Sea).Flag(godip.Sea).
		// den
		Prov("den").Conn("hel", godip.Sea).Conn("nth", godip.Sea).Conn("ska", godip.Sea).Conn("swe", godip.Coast...).Conn("bal", godip.Sea).Conn("kie", godip.Coast...).Flag(godip.Coast...).SC(Denmark).
		// ber
		Prov("ber").Conn("kie", godip.Coast...).Conn("bal", godip.Sea).Conn("pru", godip.Coast...).Conn("sil", godip.Land).Conn("mun", godip.Land).Flag(godip.Coast...).SC(Berlin).
		// mun
		Prov("mun").Conn("bur", godip.Land).Conn("ruh", godip.Land).Conn("kie", godip.Land).Conn("ber", godip.Land).Conn("sil", godip.Land).Conn("boh", godip.Land).Conn("tyr", godip.Land).Flag(godip.Land).SC(Munich).
		// boh
		Prov("boh").Conn("mun", godip.Land).Conn("sil", godip.Land).Conn("gal", godip.Land).Conn("vie", godip.Land).Conn("tyr", godip.Land).Flag(godip.Land).
		// vie
		Prov("vie").Conn("tyr", godip.Land).Conn("boh", godip.Land).Conn("gal", godip.Land).Conn("bud", godip.Land).Conn("tri", godip.Land).Flag(godip.Land).SC(Vienna).
		// bud
		Prov("bud").Conn("tri", godip.Land).Conn("vie", godip.Land).Conn("gal", godip.Land).Conn("rum", godip.Land).Conn("ser", godip.Land).Flag(godip.Land).SC(Budapest).
		// tri
		Prov("tri").Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("tyr", godip.Land).Conn("vie", godip.Land).Conn("bud", godip.Land).Conn("ser", godip.Land).Conn("alb", godip.Coast...).Flag(godip.Coast...).SC(Trieste).
		// ven
		Prov("ven").Conn("tus", godip.Land).Conn("pie", godip.Land).Conn("tyr", godip.Land).Conn("tri", godip.Coast...).Conn("adr", godip.Sea).Conn("apu", godip.Coast...).Conn("rom", godip.Land).Flag(godip.Coast...).SC(Venice).
		// pie
		Prov("pie").Conn("mar", godip.Coast...).Conn("tyr", godip.Land).Conn("ven", godip.Land).Conn("tus", godip.Coast...).Conn("gol", godip.Sea).Flag(godip.Coast...).
		// ruh
		Prov("ruh").Conn("bel", godip.Land).Conn("hol", godip.Land).Conn("kie", godip.Land).Conn("mun", godip.Land).Conn("bur", godip.Land).Flag(godip.Land).
		// tyr
		Prov("tyr").Conn("mun", godip.Land).Conn("boh", godip.Land).Conn("vie", godip.Land).Conn("tri", godip.Land).Conn("ven", godip.Land).Conn("pie", godip.Land).Flag(godip.Land).
		// kie
		Prov("kie").Conn("hol", godip.Coast...).Conn("hel", godip.Sea).Conn("den", godip.Coast...).Conn("bal", godip.Sea).Conn("ber", godip.Coast...).Conn("mun", godip.Land).Conn("ruh", godip.Land).Flag(godip.Coast...).SC(Kiel).
		Done()
}

func Blank(phase godip.Phase) *state.State {
	return state.New(Graph(), phase, classical.BackupRule)
}

func Start() (*state.State, error) {
	g := Graph()
	result := state.New(g, Phase(1900, godip.Fall, godip.Adjustment), classical.BackupRule)
	scMap := map[godip.Province]godip.Nation{}
	for _, prov := range g.Provinces() {
		if nat := g.SC(prov); nat != nil {
			scMap[prov.Super()] = *nat
		}
	}
	result.SetSupplyCenters(scMap)
	return result, nil
}

var newPhase = phase.Generator(hundred.BuildAnywhereParser, func(phase *phase.Phase) bool {
	return phase.Ty == godip.Retreat && phase.Se == godip.Fall
})

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return &chaosPhase{newPhase(year, season, typ)}
}

type chaosPhase struct {
	godip.Phase
}

func (self *chaosPhase) PreProcess(s godip.State) (err error) {
	if self.Type() == godip.Adjustment && self.Year() == 1900 {
		ords := s.Orders()
		for _, prov := range s.Graph().Provinces() {
			if nat := s.Graph().SC(prov); nat != nil && prov == prov.Super() {
				if _, _, found := s.Order(prov); !found {
					if _, _, found := s.Unit(prov); !found {
						ords[prov] = orders.BuildAnywhere(prov, godip.Army, time.Now())
					}
				}
			}
		}
		s.SetOrders(ords)
	}
	return nil
}

var ChaosVariant = common.Variant{
	Name:       "Chaos",
	Graph:      func() godip.Graph { return Graph() },
	Start:      Start,
	Blank:      Blank,
	Phase:      Phase,
	Parser:     hundred.BuildAnywhereParser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
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
	CreatedBy:   "Unknown",
	Version:     "",
	Description: "The chaos variant is played by 34 players, each owning a single supply center on the standard Diplomacy map. The game starts in Winter of 1900 at the adjustment phase. Each player starts out choosing which type of unit they wish to build and progressing from there. When building units, a unit may be built at any owned supply center. To mitigate absent players, all supply centers without orders during adjustment will build armies by default.",
	Rules:       "The first to 18 supply centers is the winner.",
}
