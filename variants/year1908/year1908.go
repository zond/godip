package year1908

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Turkey  godip.Nation = "Turkey"
	Austria godip.Nation = "Austria"
	Italy   godip.Nation = "Italy"
	Russia  godip.Nation = "Russia"
	Germany godip.Nation = "Germany"
	Britain godip.Nation = "Britain"
	France  godip.Nation = "France"
)

var Nations = []godip.Nation{Turkey, Austria, Italy, Russia, Germany, Britain, France}

var Year1908Variant = common.Variant{
	Name:       "Year1908",
	Graph:      func() godip.Graph { return Year1908Graph() },
	Start:      Year1908Start,
	Blank:      Year1908Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(21),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/year1908map.svg")
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
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules: "",
}

func Year1908Blank(phase godip.Phase) *state.State {
	return state.New(Year1908Graph(), phase, classical.BackupRule, nil)
}

func Year1908Start() (result *state.State, err error) {
	startPhase := classical.NewPhase(1908, godip.Spring, godip.Movement)
	result = Year1908Blank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"ang": godip.Unit{godip.Fleet, Turkey},
		"con": godip.Unit{godip.Army, Turkey},
		"smy": godip.Unit{godip.Army, Turkey},
		"vie": godip.Unit{godip.Army, Austria},
		"bud": godip.Unit{godip.Army, Austria},
		"tri": godip.Unit{godip.Army, Austria},
		"nap": godip.Unit{godip.Fleet, Italy},
		"rom": godip.Unit{godip.Army, Italy},
		"mil": godip.Unit{godip.Army, Italy},
		"stp": godip.Unit{godip.Fleet, Russia},
		"sev": godip.Unit{godip.Fleet, Russia},
		"war": godip.Unit{godip.Army, Russia},
		"mos": godip.Unit{godip.Army, Russia},
		"kie": godip.Unit{godip.Fleet, Germany},
		"fra": godip.Unit{godip.Army, Germany},
		"ber": godip.Unit{godip.Army, Germany},
		"mun": godip.Unit{godip.Army, Germany},
		"edi": godip.Unit{godip.Fleet, Britain},
		"lon": godip.Unit{godip.Fleet, Britain},
		"cai": godip.Unit{godip.Fleet, Britain},
		"lvp": godip.Unit{godip.Army, Britain},
		"bre": godip.Unit{godip.Fleet, France},
		"cas": godip.Unit{godip.Army, France},
		"par": godip.Unit{godip.Army, France},
		"mar": godip.Unit{godip.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"ang": Turkey,
		"con": Turkey,
		"smy": Turkey,
		"vie": Austria,
		"bud": Austria,
		"tri": Austria,
		"nap": Italy,
		"rom": Italy,
		"mil": Italy,
		"stp": Russia,
		"sev": Russia,
		"war": Russia,
		"mos": Russia,
		"kie": Germany,
		"fra": Germany,
		"ber": Germany,
		"mun": Germany,
		"edi": Britain,
		"lon": Britain,
		"cai": Britain,
		"lvp": Britain,
		"bre": France,
		"cas": France,
		"par": France,
		"mar": France,
	})
	return
}

func Year1908Graph() *graph.Graph {
	return graph.New().
		// Silesia
		Prov("sil").Conn("gal", godip.Land).Conn("war", godip.Land).Conn("pru", godip.Land).Conn("ber", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).
		// Sevastopol
		Prov("sev").Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("rum", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Flag(godip.Coast...).SC(Russia).
		// Mid-Atlantic Ocean
		Prov("mid").Conn("cas", godip.Sea).Conn("wes", godip.Sea).Conn("spa", godip.Sea).Conn("por", godip.Sea).Conn("bob", godip.Sea).Conn("bob", godip.Sea).Conn("eng", godip.Sea).Conn("iri", godip.Sea).Conn("nat", godip.Sea).Flag(godip.Sea).
		// St. Petersburg
		Prov("stp").Conn("bar", godip.Sea).Conn("nwy", godip.Coast...).Conn("fin", godip.Coast...).Conn("gob", godip.Sea).Conn("lvn", godip.Coast...).Conn("bye", godip.Land).Conn("mos", godip.Land).Flag(godip.Coast...).SC(Russia).
		// London
		Prov("lon").Conn("wal", godip.Coast...).Conn("eng", godip.Sea).Conn("eng", godip.Sea).Conn("nth", godip.Sea).Conn("yor", godip.Coast...).Flag(godip.Coast...).SC(Britain).
		// Yorkshire
		Prov("yor").Conn("edi", godip.Coast...).Conn("lvp", godip.Land).Conn("wal", godip.Land).Conn("lon", godip.Coast...).Conn("nth", godip.Sea).Flag(godip.Coast...).
		// Galicia
		Prov("gal").Conn("bud", godip.Land).Conn("tra", godip.Land).Conn("rum", godip.Land).Conn("ukr", godip.Land).Conn("war", godip.Land).Conn("sil", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).
		// Tuscany
		Prov("tus").Conn("mil", godip.Land).Conn("pie", godip.Coast...).Conn("gol", godip.Sea).Conn("tyn", godip.Sea).Conn("rom", godip.Coast...).Flag(godip.Coast...).
		// Rome
		Prov("rom").Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("mil", godip.Land).Conn("tus", godip.Coast...).Conn("tyn", godip.Sea).Conn("nap", godip.Coast...).Conn("apu", godip.Coast...).Flag(godip.Coast...).SC(Italy).
		// Brest
		Prov("bre").Conn("gas", godip.Coast...).Conn("par", godip.Land).Conn("pic", godip.Coast...).Conn("eng", godip.Sea).Conn("bob", godip.Sea).Flag(godip.Coast...).SC(France).
		// Angora
		Prov("ang").Conn("con", godip.Coast...).Conn("smy", godip.Land).Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Tyrol
		Prov("tyr").Conn("vie", godip.Land).Conn("boh", godip.Land).Conn("mun", godip.Land).Conn("swi", godip.Land).Conn("mil", godip.Land).Conn("ven", godip.Land).Conn("tri", godip.Land).Flag(godip.Land).
		// Paris
		Prov("par").Conn("gas", godip.Land).Conn("bur", godip.Land).Conn("pic", godip.Land).Conn("bre", godip.Land).Flag(godip.Land).SC(France).
		// Ionian Sea
		Prov("ion").Conn("apu", godip.Sea).Conn("nap", godip.Sea).Conn("tyn", godip.Sea).Conn("tyn", godip.Sea).Conn("tun", godip.Sea).Conn("trp", godip.Sea).Conn("cyr", godip.Sea).Conn("eas", godip.Sea).Conn("aeg", godip.Sea).Conn("gre", godip.Sea).Conn("mac", godip.Sea).Conn("adr", godip.Sea).Flag(godip.Sea).
		// Portugal
		Prov("por").Conn("spa", godip.Coast...).Conn("bob", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Switzerland
		Prov("swi").Conn("swa", godip.Land).Conn("bur", godip.Land).Conn("mar", godip.Land).Conn("pie", godip.Land).Conn("mil", godip.Land).Conn("tyr", godip.Land).Conn("mun", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Skagerrak
		Prov("ska").Conn("nwy", godip.Sea).Conn("nth", godip.Sea).Conn("den", godip.Sea).Conn("swe", godip.Sea).Flag(godip.Sea).
		// Netherlands
		Prov("net").Conn("ruh", godip.Land).Conn("kie", godip.Coast...).Conn("hel", godip.Sea).Conn("nth", godip.Sea).Conn("bel", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Aegean Sea
		Prov("aeg").Conn("mac", godip.Sea).Conn("gre", godip.Sea).Conn("ion", godip.Sea).Conn("eas", godip.Sea).Conn("smy", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
		// Eastern Mediterranean
		Prov("eas").Conn("cyr", godip.Sea).Conn("cai", godip.Sea).Conn("lev", godip.Sea).Conn("smy", godip.Sea).Conn("aeg", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// Casablanca
		Prov("cas").Conn("alg", godip.Coast...).Conn("wes", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(France).
		// Algeria
		Prov("alg").Conn("trp", godip.Land).Conn("tun", godip.Coast...).Conn("wes", godip.Sea).Conn("cas", godip.Coast...).Flag(godip.Coast...).
		// Baltic Sea
		Prov("bal").Conn("ber", godip.Sea).Conn("pru", godip.Sea).Conn("lvn", godip.Sea).Conn("gob", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Conn("kie", godip.Sea).Flag(godip.Sea).
		// Piedmont
		Prov("pie").Conn("tus", godip.Coast...).Conn("mil", godip.Land).Conn("swi", godip.Land).Conn("mar", godip.Coast...).Conn("gol", godip.Sea).Flag(godip.Coast...).
		// Budapest
		Prov("bud").Conn("gal", godip.Land).Conn("boh", godip.Land).Conn("vie", godip.Land).Conn("tri", godip.Land).Conn("tra", godip.Land).Flag(godip.Land).SC(Austria).
		// Bulgaria
		Prov("bul").Conn("bla", godip.Sea).Conn("rum", godip.Coast...).Conn("ser", godip.Land).Conn("mac", godip.Land).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Serbia
		Prov("ser").Conn("mac", godip.Land).Conn("bul", godip.Land).Conn("rum", godip.Land).Conn("tra", godip.Land).Conn("tri", godip.Land).Conn("bos", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Belgium
		Prov("bel").Conn("ruh", godip.Land).Conn("net", godip.Coast...).Conn("nth", godip.Sea).Conn("eng", godip.Sea).Conn("pic", godip.Coast...).Conn("bur", godip.Land).Conn("swa", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Ruhr
		Prov("ruh").Conn("bel", godip.Land).Conn("swa", godip.Land).Conn("fra", godip.Land).Conn("kie", godip.Land).Conn("net", godip.Land).Flag(godip.Land).
		// Byelorussia
		Prov("bye").Conn("war", godip.Land).Conn("ukr", godip.Land).Conn("mos", godip.Land).Conn("stp", godip.Land).Conn("lvn", godip.Land).Flag(godip.Land).
		// Liverpool
		Prov("lvp").Conn("iri", godip.Sea).Conn("wal", godip.Coast...).Conn("yor", godip.Land).Conn("edi", godip.Land).Conn("cly", godip.Coast...).Conn("nat", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Denmark
		Prov("den").Conn("nth", godip.Sea).Conn("hel", godip.Sea).Conn("kie", godip.Coast...).Conn("bal", godip.Sea).Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Adriatic Sea
		Prov("adr").Conn("ven", godip.Sea).Conn("rom", godip.Sea).Conn("apu", godip.Sea).Conn("ion", godip.Sea).Conn("mac", godip.Sea).Conn("bos", godip.Sea).Conn("tri", godip.Sea).Flag(godip.Sea).
		// English Channel
		Prov("eng").Conn("nth", godip.Sea).Conn("lon", godip.Sea).Conn("lon", godip.Sea).Conn("wal", godip.Sea).Conn("iri", godip.Sea).Conn("mid", godip.Sea).Conn("bob", godip.Sea).Conn("bre", godip.Sea).Conn("pic", godip.Sea).Conn("bel", godip.Sea).Flag(godip.Sea).
		// Armenia
		Prov("arm").Conn("sev", godip.Coast...).Conn("bla", godip.Sea).Conn("ang", godip.Coast...).Conn("smy", godip.Land).Conn("lev", godip.Land).Flag(godip.Coast...).
		// Norway
		Prov("nwy").Conn("stp", godip.Coast...).Conn("bar", godip.Sea).Conn("nrg", godip.Sea).Conn("nth", godip.Sea).Conn("ska", godip.Sea).Conn("swe", godip.Coast...).Conn("fin", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Western Mediterranean
		Prov("wes").Conn("tun", godip.Sea).Conn("tyn", godip.Sea).Conn("gol", godip.Sea).Conn("spa", godip.Sea).Conn("mid", godip.Sea).Conn("cas", godip.Sea).Conn("alg", godip.Sea).Flag(godip.Sea).
		// Tyrrhenian Sea
		Prov("tyn").Conn("nap", godip.Sea).Conn("rom", godip.Sea).Conn("tus", godip.Sea).Conn("gol", godip.Sea).Conn("gol", godip.Sea).Conn("wes", godip.Sea).Conn("tun", godip.Sea).Conn("ion", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// Gulf of Bothnia
		Prov("gob").Conn("stp", godip.Sea).Conn("fin", godip.Sea).Conn("swe", godip.Sea).Conn("bal", godip.Sea).Conn("lvn", godip.Sea).Flag(godip.Sea).
		// Gascony
		Prov("gas").Conn("mar", godip.Land).Conn("bur", godip.Land).Conn("par", godip.Land).Conn("bre", godip.Coast...).Conn("bob", godip.Sea).Conn("spa", godip.Coast...).Flag(godip.Coast...).
		// North Sea
		Prov("nth").Conn("den", godip.Sea).Conn("ska", godip.Sea).Conn("nwy", godip.Sea).Conn("nrg", godip.Sea).Conn("edi", godip.Sea).Conn("yor", godip.Sea).Conn("lon", godip.Sea).Conn("eng", godip.Sea).Conn("bel", godip.Sea).Conn("net", godip.Sea).Conn("hel", godip.Sea).Flag(godip.Sea).
		// Constantinople
		Prov("con").Conn("smy", godip.Coast...).Conn("ang", godip.Coast...).Conn("bla", godip.Sea).Conn("bul", godip.Coast...).Conn("mac", godip.Coast...).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Smyrna
		Prov("smy").Conn("con", godip.Coast...).Conn("aeg", godip.Sea).Conn("eas", godip.Sea).Conn("lev", godip.Coast...).Conn("arm", godip.Land).Conn("ang", godip.Land).Flag(godip.Coast...).SC(Turkey).
		// Marseilles
		Prov("mar").Conn("swi", godip.Land).Conn("bur", godip.Land).Conn("gas", godip.Land).Conn("spa", godip.Coast...).Conn("gol", godip.Sea).Conn("pie", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Ukraine
		Prov("ukr").Conn("sev", godip.Land).Conn("mos", godip.Land).Conn("bye", godip.Land).Conn("war", godip.Land).Conn("gal", godip.Land).Conn("rum", godip.Land).Flag(godip.Land).
		// North Atlantic Ocean
		Prov("nat").Conn("mid", godip.Sea).Conn("iri", godip.Sea).Conn("lvp", godip.Sea).Conn("cly", godip.Sea).Conn("nrg", godip.Sea).Conn("nrg", godip.Sea).Flag(godip.Sea).
		// Swabia
		Prov("swa").Conn("swi", godip.Land).Conn("mun", godip.Land).Conn("fra", godip.Land).Conn("ruh", godip.Land).Conn("bel", godip.Land).Conn("bur", godip.Land).Flag(godip.Land).
		// Cyrenaica
		Prov("cyr").Conn("cai", godip.Coast...).Conn("eas", godip.Sea).Conn("ion", godip.Sea).Conn("trp", godip.Coast...).Flag(godip.Coast...).
		// Spain
		Prov("spa").Conn("por", godip.Coast...).Conn("mid", godip.Sea).Conn("wes", godip.Sea).Conn("gol", godip.Sea).Conn("mar", godip.Coast...).Conn("gas", godip.Coast...).Conn("bob", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Warsaw
		Prov("war").Conn("bye", godip.Land).Conn("lvn", godip.Land).Conn("pru", godip.Land).Conn("sil", godip.Land).Conn("gal", godip.Land).Conn("ukr", godip.Land).Flag(godip.Land).SC(Russia).
		// Norwegian Sea
		Prov("nrg").Conn("nat", godip.Sea).Conn("nat", godip.Sea).Conn("cly", godip.Sea).Conn("edi", godip.Sea).Conn("nth", godip.Sea).Conn("nwy", godip.Sea).Conn("bar", godip.Sea).Flag(godip.Sea).
		// Levant
		Prov("lev").Conn("arm", godip.Land).Conn("smy", godip.Coast...).Conn("eas", godip.Sea).Conn("cai", godip.Coast...).Flag(godip.Coast...).
		// Wales
		Prov("wal").Conn("lvp", godip.Coast...).Conn("iri", godip.Sea).Conn("eng", godip.Sea).Conn("lon", godip.Coast...).Conn("yor", godip.Land).Flag(godip.Coast...).
		// Greece
		Prov("gre").Conn("mac", godip.Coast...).Conn("ion", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Venice
		Prov("ven").Conn("adr", godip.Sea).Conn("tri", godip.Coast...).Conn("tyr", godip.Land).Conn("mil", godip.Land).Conn("rom", godip.Coast...).Flag(godip.Coast...).
		// Cairo
		Prov("cai").Conn("lev", godip.Coast...).Conn("eas", godip.Sea).Conn("cyr", godip.Coast...).Flag(godip.Coast...).SC(Britain).
		// Vienna
		Prov("vie").Conn("boh", godip.Land).Conn("tyr", godip.Land).Conn("tri", godip.Land).Conn("bud", godip.Land).Flag(godip.Land).SC(Austria).
		// Helgoland Bight
		Prov("hel").Conn("den", godip.Sea).Conn("nth", godip.Sea).Conn("net", godip.Sea).Conn("kie", godip.Sea).Flag(godip.Sea).
		// Transylvania
		Prov("tra").Conn("tri", godip.Land).Conn("ser", godip.Land).Conn("rum", godip.Land).Conn("gal", godip.Land).Conn("bud", godip.Land).Flag(godip.Land).
		// Sweden
		Prov("swe").Conn("gob", godip.Sea).Conn("fin", godip.Coast...).Conn("nwy", godip.Coast...).Conn("ska", godip.Sea).Conn("den", godip.Coast...).Conn("bal", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Bohemia
		Prov("boh").Conn("vie", godip.Land).Conn("bud", godip.Land).Conn("gal", godip.Land).Conn("sil", godip.Land).Conn("ber", godip.Land).Conn("mun", godip.Land).Conn("tyr", godip.Land).Flag(godip.Land).
		// Rumania
		Prov("rum").Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Conn("ukr", godip.Land).Conn("gal", godip.Land).Conn("tra", godip.Land).Conn("ser", godip.Land).Conn("bul", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Milan
		Prov("mil").Conn("tus", godip.Land).Conn("rom", godip.Land).Conn("ven", godip.Land).Conn("tyr", godip.Land).Conn("swi", godip.Land).Conn("pie", godip.Land).Flag(godip.Land).SC(Italy).
		// Macedonia
		Prov("mac").Conn("ser", godip.Land).Conn("bos", godip.Coast...).Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("gre", godip.Coast...).Conn("aeg", godip.Sea).Conn("con", godip.Coast...).Conn("bul", godip.Land).Flag(godip.Coast...).
		// Black Sea
		Prov("bla").Conn("rum", godip.Sea).Conn("bul", godip.Sea).Conn("con", godip.Sea).Conn("ang", godip.Sea).Conn("arm", godip.Sea).Conn("sev", godip.Sea).Flag(godip.Sea).
		// Tunisia
		Prov("tun").Conn("wes", godip.Sea).Conn("alg", godip.Coast...).Conn("trp", godip.Coast...).Conn("ion", godip.Sea).Conn("tyn", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Barents Sea
		Prov("bar").Conn("nrg", godip.Sea).Conn("nwy", godip.Sea).Conn("stp", godip.Sea).Flag(godip.Sea).
		// Bosnia
		Prov("bos").Conn("tri", godip.Coast...).Conn("adr", godip.Sea).Conn("mac", godip.Coast...).Conn("ser", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Clyde
		Prov("cly").Conn("edi", godip.Coast...).Conn("nrg", godip.Sea).Conn("nat", godip.Sea).Conn("lvp", godip.Coast...).Flag(godip.Coast...).
		// Apulia
		Prov("apu").Conn("ion", godip.Sea).Conn("adr", godip.Sea).Conn("rom", godip.Coast...).Conn("nap", godip.Coast...).Flag(godip.Coast...).
		// Edinburgh
		Prov("edi").Conn("cly", godip.Coast...).Conn("lvp", godip.Land).Conn("yor", godip.Coast...).Conn("nth", godip.Sea).Conn("nrg", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Bay of Biscay
		Prov("bob").Conn("mid", godip.Sea).Conn("por", godip.Sea).Conn("spa", godip.Sea).Conn("gas", godip.Sea).Conn("bre", godip.Sea).Conn("eng", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Sea).
		// Gulf of Lyon
		Prov("gol").Conn("tyn", godip.Sea).Conn("tus", godip.Sea).Conn("pie", godip.Sea).Conn("mar", godip.Sea).Conn("spa", godip.Sea).Conn("wes", godip.Sea).Conn("tyn", godip.Sea).Flag(godip.Sea).
		// Irish Sea
		Prov("iri").Conn("lvp", godip.Sea).Conn("nat", godip.Sea).Conn("mid", godip.Sea).Conn("eng", godip.Sea).Conn("wal", godip.Sea).Flag(godip.Sea).
		// Finland
		Prov("fin").Conn("nwy", godip.Land).Conn("swe", godip.Coast...).Conn("gob", godip.Sea).Conn("stp", godip.Coast...).Flag(godip.Coast...).
		// Prussia
		Prov("pru").Conn("ber", godip.Coast...).Conn("sil", godip.Land).Conn("war", godip.Land).Conn("lvn", godip.Coast...).Conn("bal", godip.Sea).Flag(godip.Coast...).
		// Berlin
		Prov("ber").Conn("pru", godip.Coast...).Conn("bal", godip.Sea).Conn("kie", godip.Coast...).Conn("mun", godip.Land).Conn("boh", godip.Land).Conn("sil", godip.Land).Flag(godip.Coast...).SC(Germany).
		// Livonia
		Prov("lvn").Conn("pru", godip.Coast...).Conn("war", godip.Land).Conn("bye", godip.Land).Conn("stp", godip.Coast...).Conn("gob", godip.Sea).Conn("bal", godip.Sea).Flag(godip.Coast...).
		// Burgundy
		Prov("bur").Conn("swi", godip.Land).Conn("swa", godip.Land).Conn("bel", godip.Land).Conn("pic", godip.Land).Conn("par", godip.Land).Conn("gas", godip.Land).Conn("mar", godip.Land).Flag(godip.Land).
		// Frankfurt
		Prov("fra").Conn("ruh", godip.Land).Conn("swa", godip.Land).Conn("mun", godip.Land).Conn("kie", godip.Land).Flag(godip.Land).SC(Germany).
		// Naples
		Prov("nap").Conn("tyn", godip.Sea).Conn("ion", godip.Sea).Conn("apu", godip.Coast...).Conn("rom", godip.Coast...).Flag(godip.Coast...).SC(Italy).
		// Tripolitania
		Prov("trp").Conn("cyr", godip.Coast...).Conn("ion", godip.Sea).Conn("tun", godip.Coast...).Conn("alg", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Kiel
		Prov("kie").Conn("mun", godip.Land).Conn("ber", godip.Coast...).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Conn("hel", godip.Sea).Conn("net", godip.Coast...).Conn("ruh", godip.Land).Conn("fra", godip.Land).Flag(godip.Coast...).SC(Germany).
		// Moscow
		Prov("mos").Conn("stp", godip.Land).Conn("bye", godip.Land).Conn("ukr", godip.Land).Conn("sev", godip.Land).Flag(godip.Land).SC(Russia).
		// Trieste
		Prov("tri").Conn("bos", godip.Coast...).Conn("ser", godip.Land).Conn("tra", godip.Land).Conn("bud", godip.Land).Conn("vie", godip.Land).Conn("tyr", godip.Land).Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Flag(godip.Coast...).SC(Austria).
		// Picardy
		Prov("pic").Conn("bur", godip.Land).Conn("bel", godip.Coast...).Conn("eng", godip.Sea).Conn("bre", godip.Coast...).Conn("par", godip.Land).Flag(godip.Coast...).
		// Munich
		Prov("mun").Conn("kie", godip.Land).Conn("fra", godip.Land).Conn("swa", godip.Land).Conn("swi", godip.Land).Conn("tyr", godip.Land).Conn("boh", godip.Land).Conn("ber", godip.Land).Flag(godip.Land).SC(Germany).
		Done()
}
