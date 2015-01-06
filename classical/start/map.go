package start

import (
	c "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
	"github.com/zond/godip/graph"
)

func SCs() (result map[dip.Province]dip.Nation) {
	result = map[dip.Province]dip.Nation{}
	g := Graph()
	for _, prov := range g.Provinces() {
		if nat := g.SC(prov); nat != nil {
			result[prov] = *nat
		}
	}
	return
}

func Graph() *graph.Graph {
	return graph.New().
		// nat
		Prov("nat").Conn("nrg", c.Sea).Conn("cly", c.Sea).Conn("lvp", c.Sea).Conn("iri", c.Sea).Conn("mid", c.Sea).Flag(c.Sea).
		// nrg
		Prov("nrg").Conn("nat", c.Sea).Conn("bar", c.Sea).Conn("nwy", c.Sea).Conn("nth", c.Sea).Conn("edi", c.Sea).Conn("cly", c.Sea).Flag(c.Sea).
		// bar
		Prov("bar").Conn("nrg", c.Sea).Conn("stp/nc", c.Sea).Conn("nwy", c.Sea).Flag(c.Sea).
		// stp/nc
		Prov("stp/nc").Conn("bar", c.Sea).Conn("nwy", c.Sea).Flag(c.Sea).
		// stp
		Prov("stp").Conn("fin", c.Land).Conn("nwy", c.Land).Conn("mos", c.Land).Conn("lvn", c.Land).Flag(c.Land).SC(c.Russia).
		// mos
		Prov("mos").Conn("stp", c.Land).Conn("sev", c.Land).Conn("ukr", c.Land).Conn("war", c.Land).Conn("lvn", c.Land).Flag(c.Land).SC(c.Russia).
		// sev
		Prov("sev").Conn("ukr", c.Land).Conn("mos", c.Land).Conn("arm", c.Coast...).Conn("bla", c.Sea).Conn("rum", c.Coast...).Flag(c.Coast...).SC(c.Russia).
		// arm
		Prov("arm").Conn("ank", c.Coast...).Conn("bla", c.Sea).Conn("sev", c.Coast...).Conn("syr", c.Land).Conn("smy", c.Land).Flag(c.Coast...).
		// syr
		Prov("syr").Conn("eas", c.Sea).Conn("smy", c.Coast...).Conn("arm", c.Land).Flag(c.Coast...).
		// eas
		Prov("eas").Conn("ion", c.Sea).Conn("aeg", c.Sea).Conn("smy", c.Sea).Conn("syr", c.Sea).Flag(c.Sea).
		// ion
		Prov("ion").Conn("apu", c.Sea).Conn("adr", c.Sea).Conn("tun", c.Sea).Conn("tys", c.Sea).Conn("nap", c.Sea).Conn("alb", c.Sea).Conn("gre", c.Sea).Conn("aeg", c.Sea).Conn("eas", c.Sea).Flag(c.Sea).
		// tun
		Prov("tun").Conn("naf", c.Coast...).Conn("wes", c.Sea).Conn("tys", c.Sea).Conn("ion", c.Sea).Flag(c.Coast...).SC(c.Neutral).
		// naf
		Prov("naf").Conn("mid", c.Sea).Conn("wes", c.Sea).Conn("tun", c.Coast...).Flag(c.Coast...).
		// mid
		Prov("mid").Conn("wes", c.Sea).Conn("nat", c.Sea).Conn("iri", c.Sea).Conn("eng", c.Sea).Conn("bre", c.Sea).Conn("gas", c.Sea).Conn("spa/nc", c.Sea).Conn("por", c.Sea).Conn("spa/sc", c.Sea).Conn("naf", c.Sea).Flag(c.Sea).
		// iri
		Prov("iri").Conn("nat", c.Sea).Conn("lvp", c.Sea).Conn("wal", c.Sea).Conn("eng", c.Sea).Conn("mid", c.Sea).Flag(c.Sea).
		// lvp
		Prov("lvp").Conn("iri", c.Sea).Conn("nat", c.Sea).Conn("cly", c.Coast...).Conn("edi", c.Land).Conn("yor", c.Land).Conn("wal", c.Coast...).Flag(c.Coast...).SC(c.England).
		// cly
		Prov("cly").Conn("nat", c.Sea).Conn("nrg", c.Sea).Conn("edi", c.Coast...).Conn("lvp", c.Coast...).Flag(c.Coast...).
		// edi
		Prov("edi").Conn("cly", c.Coast...).Conn("nrg", c.Sea).Conn("nth", c.Sea).Conn("yor", c.Coast...).Conn("lvp", c.Land).Flag(c.Coast...).SC(c.England).
		// nth
		Prov("nth").Conn("eng", c.Sea).Conn("edi", c.Sea).Conn("nrg", c.Sea).Conn("nwy", c.Sea).Conn("ska", c.Sea).Conn("den", c.Sea).Conn("hel", c.Sea).Conn("hol", c.Sea).Conn("bel", c.Sea).Conn("lon", c.Sea).Conn("yor", c.Sea).Flag(c.Sea).
		// nwy
		Prov("nwy").Conn("nth", c.Sea).Conn("nrg", c.Sea).Conn("bar", c.Sea).Conn("stp/nc", c.Sea).Conn("stp", c.Land).Conn("fin", c.Land).Conn("swe", c.Coast...).Conn("ska", c.Sea).Flag(c.Coast...).SC(c.Neutral).
		// stp/sc
		Prov("stp/sc").Conn("bot", c.Sea).Conn("fin", c.Sea).Conn("lvn", c.Sea).Flag(c.Sea).
		// lvn
		Prov("lvn").Conn("stp", c.Land).Conn("bal", c.Sea).Conn("bot", c.Sea).Conn("stp/sc", c.Sea).Conn("mos", c.Land).Conn("war", c.Land).Conn("pru", c.Coast...).Flag(c.Coast...).
		// war
		Prov("war").Conn("sil", c.Land).Conn("pru", c.Land).Conn("lvn", c.Land).Conn("mos", c.Land).Conn("ukr", c.Land).Conn("gal", c.Land).Flag(c.Land).SC(c.Russia).
		// ukr
		Prov("ukr").Conn("war", c.Land).Conn("mos", c.Land).Conn("sev", c.Land).Conn("rum", c.Land).Conn("gal", c.Land).Flag(c.Land).
		// bla
		Prov("bla").Conn("bul/ec", c.Sea).Conn("rum", c.Sea).Conn("sev", c.Sea).Conn("arm", c.Sea).Conn("ank", c.Sea).Conn("con", c.Sea).Flag(c.Sea).
		// ank
		Prov("ank").Conn("con", c.Coast...).Conn("bla", c.Sea).Conn("arm", c.Coast...).Conn("smy", c.Land).Flag(c.Coast...).SC(c.Turkey).
		// smy
		Prov("smy").Conn("aeg", c.Sea).Conn("con", c.Coast...).Conn("ank", c.Land).Conn("arm", c.Land).Conn("syr", c.Coast...).Conn("eas", c.Sea).Flag(c.Coast...).SC(c.Turkey).
		// aeg
		Prov("aeg").Conn("eas", c.Sea).Conn("ion", c.Sea).Conn("gre", c.Sea).Conn("bul/sc", c.Sea).Conn("con", c.Sea).Conn("smy", c.Sea).Flag(c.Sea).
		// gre
		Prov("gre").Conn("ion", c.Sea).Conn("alb", c.Coast...).Conn("ser", c.Land).Conn("bul", c.Land).Conn("bul/sc", c.Sea).Conn("aeg", c.Sea).Flag(c.Coast...).SC(c.Neutral).
		// nap
		Prov("nap").Conn("tys", c.Sea).Conn("rom", c.Coast...).Conn("apu", c.Coast...).Conn("ion", c.Sea).Flag(c.Coast...).SC(c.Italy).
		// tys
		Prov("tys").Conn("wes", c.Sea).Conn("gol", c.Sea).Conn("tus", c.Sea).Conn("rom", c.Sea).Conn("nap", c.Sea).Conn("ion", c.Sea).Conn("tun", c.Sea).Flag(c.Sea).
		// wes
		Prov("wes").Conn("mid", c.Sea).Conn("spa/sc", c.Sea).Conn("gol", c.Sea).Conn("tys", c.Sea).Conn("tun", c.Sea).Conn("naf", c.Sea).Flag(c.Sea).
		// spa/sc
		Prov("spa/sc").Conn("mid", c.Sea).Conn("por", c.Sea).Conn("mar", c.Sea).Conn("gol", c.Sea).Conn("wes", c.Sea).Flag(c.Sea).
		// spa
		Prov("spa").Conn("por", c.Land).Conn("gas", c.Land).Conn("mar", c.Land).Flag(c.Land).SC(c.Neutral).
		// spa/nc
		Prov("spa/nc").Conn("por", c.Sea).Conn("mid", c.Sea).Conn("gas", c.Sea).Flag(c.Sea).
		// por
		Prov("por").Conn("mid", c.Sea).Conn("spa/nc", c.Sea).Conn("spa", c.Land).Conn("spa/sc", c.Sea).Flag(c.Coast...).SC(c.Neutral).
		// gas
		Prov("gas").Conn("mid", c.Sea).Conn("bre", c.Coast...).Conn("par", c.Land).Conn("bur", c.Land).Conn("mar", c.Land).Conn("spa", c.Land).Conn("spa/nc", c.Sea).Flag(c.Coast...).
		// bre
		Prov("bre").Conn("mid", c.Sea).Conn("eng", c.Sea).Conn("pic", c.Coast...).Conn("par", c.Land).Conn("gas", c.Coast...).Flag(c.Coast...).SC(c.France).
		// eng
		Prov("eng").Conn("mid", c.Sea).Conn("iri", c.Sea).Conn("wal", c.Sea).Conn("lon", c.Sea).Conn("nth", c.Sea).Conn("bel", c.Sea).Conn("pic", c.Sea).Conn("bre", c.Sea).Flag(c.Sea).
		// wal
		Prov("wal").Conn("iri", c.Sea).Conn("lvp", c.Coast...).Conn("yor", c.Land).Conn("lon", c.Coast...).Conn("eng", c.Sea).Flag(c.Coast...).
		// yor
		Prov("yor").Conn("lvp", c.Land).Conn("edi", c.Coast...).Conn("nth", c.Sea).Conn("lon", c.Coast...).Conn("wal", c.Land).Flag(c.Coast...).
		// ska
		Prov("ska").Conn("nth", c.Sea).Conn("nwy", c.Sea).Conn("swe", c.Sea).Conn("den", c.Sea).Flag(c.Sea).
		// swe
		Prov("swe").Conn("ska", c.Sea).Conn("nwy", c.Coast...).Conn("fin", c.Coast...).Conn("bot", c.Sea).Conn("bal", c.Sea).Conn("den", c.Coast...).Flag(c.Coast...).SC(c.Neutral).
		// fin
		Prov("fin").Conn("nwy", c.Land).Conn("bot", c.Sea).Conn("swe", c.Coast...).Conn("stp", c.Land).Conn("stp/sc", c.Sea).Flag(c.Coast...).
		// bot
		Prov("bot").Conn("swe", c.Sea).Conn("fin", c.Sea).Conn("stp/sc", c.Sea).Conn("lvn", c.Sea).Conn("bal", c.Sea).Flag(c.Sea).
		// bal
		Prov("bal").Conn("den", c.Sea).Conn("swe", c.Sea).Conn("bot", c.Sea).Conn("lvn", c.Sea).Conn("pru", c.Sea).Conn("ber", c.Sea).Conn("kie", c.Sea).Flag(c.Sea).
		// pru
		Prov("pru").Conn("ber", c.Coast...).Conn("bal", c.Sea).Conn("lvn", c.Coast...).Conn("war", c.Land).Conn("sil", c.Land).Flag(c.Coast...).
		// sil
		Prov("sil").Conn("mun", c.Land).Conn("ber", c.Land).Conn("pru", c.Land).Conn("war", c.Land).Conn("gal", c.Land).Conn("boh", c.Land).Flag(c.Land).
		// gal
		Prov("gal").Conn("boh", c.Land).Conn("sil", c.Land).Conn("war", c.Land).Conn("ukr", c.Land).Conn("rum", c.Land).Conn("bud", c.Land).Conn("vie", c.Land).Flag(c.Land).
		// rum
		Prov("rum").Conn("bla", c.Sea).Conn("bud", c.Land).Conn("gal", c.Land).Conn("ukr", c.Land).Conn("sev", c.Coast...).Conn("bul/ec", c.Sea).Conn("bul", c.Land).Conn("ser", c.Land).Flag(c.Coast...).SC(c.Neutral).
		// bul/ec
		Prov("bul/ec").Conn("rum", c.Sea).Conn("bla", c.Sea).Conn("con", c.Sea).Flag(c.Sea).
		// bul
		Prov("bul").Conn("ser", c.Land).Conn("rum", c.Land).Conn("con", c.Land).Conn("gre", c.Land).Flag(c.Land).SC(c.Neutral).
		// con
		Prov("con").Conn("bul/sc", c.Sea).Conn("bul", c.Land).Conn("bul/ec", c.Sea).Conn("bla", c.Sea).Conn("ank", c.Coast...).Conn("smy", c.Coast...).Conn("aeg", c.Sea).Flag(c.Coast...).SC(c.Turkey).
		// bul/sc
		Prov("bul/sc").Conn("gre", c.Sea).Conn("con", c.Sea).Conn("aeg", c.Sea).Flag(c.Sea).
		// ser
		Prov("ser").Conn("tri", c.Land).Conn("bud", c.Land).Conn("rum", c.Land).Conn("bul", c.Land).Conn("gre", c.Land).Conn("alb", c.Land).Flag(c.Land).SC(c.Neutral).
		// alb
		Prov("alb").Conn("adr", c.Sea).Conn("tri", c.Coast...).Conn("ser", c.Land).Conn("gre", c.Coast...).Conn("ion", c.Sea).Flag(c.Coast...).
		// adr
		Prov("adr").Conn("ven", c.Sea).Conn("tri", c.Sea).Conn("alb", c.Sea).Conn("ion", c.Sea).Conn("apu", c.Sea).Flag(c.Sea).
		// apu
		Prov("apu").Conn("rom", c.Land).Conn("ven", c.Coast...).Conn("adr", c.Sea).Conn("ion", c.Sea).Conn("nap", c.Coast...).Flag(c.Coast...).
		// rom
		Prov("rom").Conn("tys", c.Sea).Conn("tus", c.Coast...).Conn("ven", c.Land).Conn("apu", c.Land).Conn("nap", c.Coast...).Flag(c.Coast...).SC(c.Italy).
		// tus
		Prov("tus").Conn("gol", c.Sea).Conn("pie", c.Coast...).Conn("ven", c.Land).Conn("rom", c.Coast...).Conn("tys", c.Sea).Flag(c.Coast...).
		// gol
		Prov("gol").Conn("spa/sc", c.Sea).Conn("mar", c.Sea).Conn("pie", c.Sea).Conn("tus", c.Sea).Conn("tys", c.Sea).Conn("wes", c.Sea).Flag(c.Sea).
		// mar
		Prov("mar").Conn("spa", c.Land).Conn("gas", c.Land).Conn("bur", c.Land).Conn("pie", c.Coast...).Conn("gol", c.Sea).Conn("spa/sc", c.Sea).Flag(c.Coast...).SC(c.France).
		// bur
		Prov("bur").Conn("par", c.Land).Conn("pic", c.Land).Conn("bel", c.Land).Conn("ruh", c.Land).Conn("mun", c.Land).Conn("mar", c.Land).Conn("gas", c.Land).Flag(c.Land).
		// par
		Prov("par").Conn("bre", c.Land).Conn("pic", c.Land).Conn("bur", c.Land).Conn("gas", c.Land).Flag(c.Land).SC(c.France).
		// pic
		Prov("pic").Conn("bre", c.Coast...).Conn("eng", c.Sea).Conn("bel", c.Coast...).Conn("bur", c.Land).Conn("par", c.Land).Flag(c.Coast...).
		// lon
		Prov("lon").Conn("wal", c.Coast...).Conn("yor", c.Coast...).Conn("nth", c.Sea).Conn("eng", c.Sea).Flag(c.Coast...).SC(c.England).
		// bel
		Prov("bel").Conn("pic", c.Coast...).Conn("eng", c.Sea).Conn("nth", c.Sea).Conn("hol", c.Coast...).Conn("ruh", c.Land).Conn("bur", c.Land).Flag(c.Coast...).SC(c.Neutral).
		// hol
		Prov("hol").Conn("nth", c.Sea).Conn("hel", c.Sea).Conn("kie", c.Coast...).Conn("ruh", c.Land).Conn("bel", c.Coast...).Flag(c.Coast...).SC(c.Neutral).
		// hel
		Prov("hel").Conn("nth", c.Sea).Conn("den", c.Sea).Conn("kie", c.Sea).Conn("hol", c.Sea).Flag(c.Sea).
		// den
		Prov("den").Conn("hel", c.Sea).Conn("nth", c.Sea).Conn("ska", c.Sea).Conn("swe", c.Coast...).Conn("bal", c.Sea).Conn("kie", c.Coast...).Flag(c.Coast...).SC(c.Neutral).
		// ber
		Prov("ber").Conn("kie", c.Coast...).Conn("bal", c.Sea).Conn("pru", c.Coast...).Conn("sil", c.Land).Conn("mun", c.Land).Flag(c.Coast...).SC(c.Germany).
		// mun
		Prov("mun").Conn("bur", c.Land).Conn("ruh", c.Land).Conn("kie", c.Land).Conn("ber", c.Land).Conn("sil", c.Land).Conn("boh", c.Land).Conn("tyr", c.Land).Flag(c.Land).SC(c.Germany).
		// boh
		Prov("boh").Conn("mun", c.Land).Conn("sil", c.Land).Conn("gal", c.Land).Conn("vie", c.Land).Conn("tyr", c.Land).Flag(c.Land).
		// vie
		Prov("vie").Conn("tyr", c.Land).Conn("boh", c.Land).Conn("gal", c.Land).Conn("bud", c.Land).Conn("tri", c.Land).Flag(c.Land).SC(c.Austria).
		// bud
		Prov("bud").Conn("tri", c.Land).Conn("vie", c.Land).Conn("gal", c.Land).Conn("rum", c.Land).Conn("ser", c.Land).Flag(c.Land).SC(c.Austria).
		// tri
		Prov("tri").Conn("adr", c.Sea).Conn("ven", c.Coast...).Conn("tyr", c.Land).Conn("vie", c.Land).Conn("bud", c.Land).Conn("ser", c.Land).Conn("alb", c.Coast...).Flag(c.Coast...).SC(c.Austria).
		// ven
		Prov("ven").Conn("tus", c.Land).Conn("pie", c.Land).Conn("tyr", c.Land).Conn("tri", c.Coast...).Conn("adr", c.Sea).Conn("apu", c.Coast...).Conn("rom", c.Land).Flag(c.Coast...).SC(c.Italy).
		// pie
		Prov("pie").Conn("mar", c.Coast...).Conn("tyr", c.Land).Conn("ven", c.Land).Conn("tus", c.Coast...).Conn("gol", c.Sea).Flag(c.Coast...).
		// ruh
		Prov("ruh").Conn("bel", c.Land).Conn("hol", c.Land).Conn("kie", c.Land).Conn("mun", c.Land).Conn("bur", c.Land).Flag(c.Land).
		// tyr
		Prov("tyr").Conn("mun", c.Land).Conn("boh", c.Land).Conn("vie", c.Land).Conn("tri", c.Land).Conn("ven", c.Land).Conn("pie", c.Land).Flag(c.Land).
		// kie
		Prov("kie").Conn("hol", c.Coast...).Conn("hel", c.Sea).Conn("den", c.Coast...).Conn("bal", c.Sea).Conn("ber", c.Coast...).Conn("mun", c.Land).Conn("ruh", c.Land).Flag(c.Coast...).SC(c.Germany).
		Done()
}
