package europe1939

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Turkey  godip.Nation = "Turkey"
	Italy   godip.Nation = "Italy"
	Poland  godip.Nation = "Poland"
	France  godip.Nation = "France"
	Britain godip.Nation = "Britain"
	USSR    godip.Nation = "USSR"
	Germany godip.Nation = "Germany"
	Spain   godip.Nation = "Spain"
)

var Nations = []godip.Nation{Turkey, Italy, Poland, France, Britain, USSR, Germany, Spain}

var Europe1939Variant = common.Variant{
	Name:       "Europe 1939",
	Graph:      func() godip.Graph { return Europe1939Graph() },
	Start:      Europe1939Start,
	Blank:      Europe1939Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(28),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/europe1939map.svg")
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
	CreatedBy:   "Mikalis Kamaritis",
	Version:     "I",
	Description: "",
	Rules: "",
}

func Europe1939Blank(phase godip.Phase) *state.State {
	return state.New(Europe1939Graph(), phase, classical.BackupRule, nil)
}

func Europe1939Start() (result *state.State, err error) {
	startPhase := classical.NewPhase(1939, godip.Spring, godip.Movement)
	result = Europe1939Blank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"izm": godip.Unit{godip.Fleet, Turkey},
		"ank": godip.Unit{godip.Fleet, Turkey},
		"ist": godip.Unit{godip.Army, Turkey},
		"ada": godip.Unit{godip.Army, Turkey},
		"nap": godip.Unit{godip.Fleet, Italy},
		"tri": godip.Unit{godip.Fleet, Italy},
		"rom": godip.Unit{godip.Army, Italy},
		"mil": godip.Unit{godip.Army, Italy},
		"alb": godip.Unit{godip.Army, Italy},
		"dan": godip.Unit{godip.Fleet, Poland},
		"war": godip.Unit{godip.Army, Poland},
		"kra": godip.Unit{godip.Army, Poland},
		"wro": godip.Unit{godip.Army, Poland},
		"alg": godip.Unit{godip.Fleet, France},
		"bre": godip.Unit{godip.Army, France},
		"par": godip.Unit{godip.Army, France},
		"lyo": godip.Unit{godip.Army, France},
		"mar": godip.Unit{godip.Army, France},
		"edi": godip.Unit{godip.Fleet, Britain},
		"noi": godip.Unit{godip.Fleet, Britain},
		"lon": godip.Unit{godip.Fleet, Britain},
		"cai": godip.Unit{godip.Fleet, Britain},
		"bag": godip.Unit{godip.Army, Britain},
		"len": godip.Unit{godip.Fleet, USSR},
		"sea": godip.Unit{godip.Fleet, USSR},
		"ark": godip.Unit{godip.Army, USSR},
		"mos": godip.Unit{godip.Army, USSR},
		"sta": godip.Unit{godip.Army, USSR},
		"kie": godip.Unit{godip.Fleet, Germany},
		"col": godip.Unit{godip.Army, Germany},
		"ber": godip.Unit{godip.Army, Germany},
		"mun": godip.Unit{godip.Army, Germany},
		"vie": godip.Unit{godip.Army, Germany},
		"sei": godip.Unit{godip.Fleet, Spain},
		"tan": godip.Unit{godip.Fleet, Spain},
		"mad": godip.Unit{godip.Army, Spain},
		"bac": godip.Unit{godip.Army, Spain},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"izm": Turkey,
		"ank": Turkey,
		"ist": Turkey,
		"ada": Turkey,
		"nap": Italy,
		"tri": Italy,
		"rom": Italy,
		"mil": Italy,
		"alb": Italy,
		"dan": Poland,
		"war": Poland,
		"kra": Poland,
		"wro": Poland,
		"alg": France,
		"bre": France,
		"par": France,
		"lyo": France,
		"mar": France,
		"edi": Britain,
		"noi": Britain,
		"lon": Britain,
		"cai": Britain,
		"bag": Britain,
		"len": USSR,
		"sea": USSR,
		"ark": USSR,
		"mos": USSR,
		"sta": USSR,
		"kie": Germany,
		"col": Germany,
		"ber": Germany,
		"mun": Germany,
		"vie": Germany,
		"sei": Spain,
		"tan": Spain,
		"mad": Spain,
		"bac": Spain,
	})
	if err = result.SetUnit(godip.Province("ser"), godip.Unit{
		Type:   godip.Army,
		Nation: godip.Neutral,
	}); err != nil {
		return
	}
	return
}

func Europe1939Graph() *graph.Graph {
	return graph.New().
		// Silesia
		Prov("sil").Conn("mun", godip.Land).Conn("boh", godip.Land).Conn("wro", godip.Land).Conn("poz", godip.Land).Conn("pru", godip.Land).Conn("ber", godip.Land).Flag(godip.Land).
		// Sevastopol
		Prov("sea").Conn("ros", godip.Coast...).Conn("sta", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("mol", godip.Coast...).Conn("wbs", godip.Sea).Conn("ebs", godip.Sea).Flag(godip.Coast...).SC(USSR).
		// Albania
		Prov("alb").Conn("mon", godip.Coast...).Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("gre", godip.Coast...).Conn("mac", godip.Land).Flag(godip.Coast...).SC(Italy).
		// Rafha
		Prov("raf").Conn("per", godip.Sea).Conn("bag", godip.Coast...).Conn("pal", godip.Land).Conn("ara", godip.Land).Flag(godip.Coast...).
		// Slovenia
		Prov("sle").Conn("vie", godip.Land).Conn("tyo", godip.Land).Conn("ven", godip.Land).Conn("cro", godip.Land).Conn("hun", godip.Land).Flag(godip.Land).
		// Palestine
		Prov("pal").Conn("eam", godip.Sea).Conn("egy", godip.Sea).Conn("cai", godip.Coast...).Conn("hed", godip.Land).Conn("ara", godip.Land).Conn("raf", godip.Land).Conn("bag", godip.Land).Conn("syr", godip.Coast...).Flag(godip.Coast...).
		// Bielorussia
		Prov("bie").Conn("len", godip.Land).Conn("lat", godip.Land).Conn("lub", godip.Land).Conn("ukr", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).
		// Syria
		Prov("syr").Conn("arm", godip.Land).Conn("ada", godip.Coast...).Conn("eam", godip.Sea).Conn("pal", godip.Coast...).Conn("bag", godip.Land).Flag(godip.Coast...).
		// London
		Prov("lon").Conn("eng", godip.Sea).Conn("not", godip.Sea).Conn("yor", godip.Coast...).Conn("liv", godip.Land).Conn("wal", godip.Coast...).Flag(godip.Coast...).SC(Britain).
		// Yorkshire
		Prov("yor").Conn("edi", godip.Coast...).Conn("liv", godip.Land).Conn("lon", godip.Coast...).Conn("not", godip.Sea).Flag(godip.Coast...).
		// Red Sea
		Prov("red").Conn("hed", godip.Sea).Conn("cai", godip.Sea).Flag(godip.Sea).
		// Rostov
		Prov("ros").Conn("cau", godip.Coast...).Conn("sta", godip.Land).Conn("sea", godip.Coast...).Conn("ebs", godip.Sea).Flag(godip.Coast...).
		// East Black Sea
		Prov("ebs").Conn("ank", godip.Sea).Conn("cau", godip.Sea).Conn("ros", godip.Sea).Conn("sea", godip.Sea).Conn("wbs", godip.Sea).Flag(godip.Sea).
		// West Black Sea
		Prov("wbs").Conn("bul", godip.Sea).Conn("ist", godip.Sea).Conn("ank", godip.Sea).Conn("ebs", godip.Sea).Conn("sea", godip.Sea).Conn("mol", godip.Sea).Conn("rum", godip.Sea).Flag(godip.Sea).
		// Tyrolia
		Prov("tyo").Conn("vie", godip.Land).Conn("boh", godip.Land).Conn("mun", godip.Land).Conn("swi", godip.Land).Conn("mil", godip.Land).Conn("ven", godip.Land).Conn("sle", godip.Land).Flag(godip.Land).
		// Tuscany
		Prov("tus").Conn("lig", godip.Sea).Conn("rom", godip.Coast...).Conn("ven", godip.Land).Conn("mil", godip.Land).Conn("pie", godip.Coast...).Flag(godip.Coast...).
		// Danzig
		Prov("dan").Conn("bal", godip.Sea).Conn("pru", godip.Coast...).Conn("poz", godip.Land).Conn("war", godip.Land).Conn("eap", godip.Coast...).Flag(godip.Coast...).SC(Poland).
		// Lublin
		Prov("lub").Conn("war", godip.Land).Conn("kra", godip.Land).Conn("ukr", godip.Land).Conn("bie", godip.Land).Conn("lat", godip.Land).Conn("lit", godip.Land).Conn("eap", godip.Land).Flag(godip.Land).
		// Rome
		Prov("rom").Conn("tys", godip.Sea).Conn("nap", godip.Coast...).Conn("apu", godip.Land).Conn("ven", godip.Land).Conn("tus", godip.Coast...).Conn("lig", godip.Sea).Flag(godip.Coast...).SC(Italy).
		// Brest
		Prov("bre").Conn("pic", godip.Coast...).Conn("eng", godip.Sea).Conn("bay", godip.Sea).Conn("gas", godip.Coast...).Conn("par", godip.Land).Flag(godip.Coast...).SC(France).
		// Coast of the Azores
		Prov("coa").Conn("mid", godip.Sea).Conn("azo", godip.Sea).Conn("sao", godip.Sea).Conn("por", godip.Sea).Conn("nav", godip.Sea).Conn("bay", godip.Sea).Conn("eng", godip.Sea).Conn("iri", godip.Sea).Flag(godip.Sea).
		// Wroclaw
		Prov("wro").Conn("poz", godip.Land).Conn("sil", godip.Land).Conn("boh", godip.Land).Conn("sla", godip.Land).Conn("kra", godip.Land).Conn("war", godip.Land).Flag(godip.Land).SC(Poland).
		// Paris
		Prov("par").Conn("gas", godip.Land).Conn("lyo", godip.Land).Conn("bur", godip.Land).Conn("pic", godip.Land).Conn("bre", godip.Land).Flag(godip.Land).SC(France).
		// Ionian Sea
		Prov("ion").Conn("aeg", godip.Sea).Conn("gre", godip.Sea).Conn("alb", godip.Sea).Conn("adr", godip.Sea).Conn("apu", godip.Sea).Conn("nap", godip.Sea).Conn("mal", godip.Sea).Conn("gos", godip.Sea).Conn("cre", godip.Sea).Flag(godip.Sea).
		// Ankara
		Prov("ank").Conn("ist", godip.Coast...).Conn("izm", godip.Land).Conn("ada", godip.Land).Conn("arm", godip.Land).Conn("cau", godip.Coast...).Conn("ebs", godip.Sea).Conn("wbs", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Portugal
		Prov("por").Conn("sao", godip.Sea).Conn("sei", godip.Coast...).Conn("mad", godip.Land).Conn("nav", godip.Coast...).Conn("coa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Adana
		Prov("ada").Conn("arm", godip.Land).Conn("ank", godip.Land).Conn("izm", godip.Coast...).Conn("eam", godip.Sea).Conn("syr", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Baghdad
		Prov("bag").Conn("khu", godip.Coast...).Conn("arm", godip.Land).Conn("syr", godip.Land).Conn("pal", godip.Land).Conn("raf", godip.Coast...).Conn("per", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Switzerland
		Prov("swi").Conn("tyo", godip.Land).Conn("mun", godip.Land).Conn("als", godip.Land).Conn("bur", godip.Land).Conn("lyo", godip.Land).Conn("mar", godip.Land).Conn("pie", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Moldova
		Prov("mol").Conn("rum", godip.Coast...).Conn("wbs", godip.Sea).Conn("sea", godip.Coast...).Conn("ukr", godip.Land).Flag(godip.Coast...).
		// Poznan
		Prov("poz").Conn("war", godip.Land).Conn("dan", godip.Land).Conn("pru", godip.Land).Conn("sil", godip.Land).Conn("wro", godip.Land).Flag(godip.Land).
		// Norway
		Prov("noa").Conn("noe", godip.Sea).Conn("not", godip.Sea).Conn("ska", godip.Sea).Conn("swe", godip.Coast...).Conn("nar", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Montenegro
		Prov("mon").Conn("cro", godip.Coast...).Conn("adr", godip.Sea).Conn("alb", godip.Coast...).Conn("mac", godip.Land).Conn("ser", godip.Land).Conn("bos", godip.Land).Flag(godip.Coast...).
		// Aegean Sea
		Prov("aeg").Conn("ion", godip.Sea).Conn("cre", godip.Sea).Conn("eam", godip.Sea).Conn("izm", godip.Sea).Conn("ist", godip.Sea).Conn("gre", godip.Sea).Flag(godip.Sea).
		// Arabia
		Prov("ara").Conn("raf", godip.Land).Conn("pal", godip.Land).Conn("hed", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Stalingrad
		Prov("sta").Conn("ros", godip.Land).Conn("cau", godip.Coast...).Conn("upp", godip.Sea).Conn("kaz", godip.Coast...).Conn("niz", godip.Land).Conn("mos", godip.Land).Conn("sea", godip.Land).Flag(godip.Coast...).SC(USSR).
		// Eastern Mediterranean
		Prov("eam").Conn("syr", godip.Sea).Conn("ada", godip.Sea).Conn("izm", godip.Sea).Conn("aeg", godip.Sea).Conn("cre", godip.Sea).Conn("egy", godip.Sea).Conn("pal", godip.Sea).Flag(godip.Sea).
		// Leningrad
		Prov("len").Conn("mos", godip.Land).Conn("niz", godip.Land).Conn("ark", godip.Land).Conn("fin", godip.Coast...).Conn("gob", godip.Sea).Conn("est", godip.Coast...).Conn("lat", godip.Coast...).Conn("bie", godip.Land).Flag(godip.Coast...).SC(USSR).
		// Baltic Sea
		Prov("bal").Conn("lit", godip.Sea).Conn("lat", godip.Sea).Conn("gob", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Conn("kie", godip.Sea).Conn("ber", godip.Sea).Conn("pru", godip.Sea).Conn("dan", godip.Sea).Conn("eap", godip.Sea).Flag(godip.Sea).
		// Atlas
		Prov("atl").Conn("ora", godip.Land).Conn("soa", godip.Land).Conn("mor", godip.Coast...).Conn("sao", godip.Sea).Flag(godip.Coast...).
		// Barents Sea
		Prov("bas").Conn("noe", godip.Sea).Conn("nar", godip.Sea).Conn("ark", godip.Sea).Flag(godip.Sea).
		// Piedmont
		Prov("pie").Conn("mar", godip.Coast...).Conn("lig", godip.Sea).Conn("tus", godip.Coast...).Conn("mil", godip.Land).Conn("swi", godip.Land).Flag(godip.Coast...).
		// Arkhangelsk
		Prov("ark").Conn("sib", godip.Land).Conn("bas", godip.Sea).Conn("nar", godip.Coast...).Conn("fin", godip.Land).Conn("len", godip.Land).Conn("niz", godip.Land).Flag(godip.Coast...).SC(USSR).
		// Khorasan
		Prov("kho").Conn("tur", godip.Land).Conn("teh", godip.Land).Conn("khu", godip.Land).Flag(godip.Land).
		// Iceland
		Prov("ice").Conn("nao", godip.Sea).Conn("noe", godip.Sea).Flag(godip.Coast...).
		// Bulgaria
		Prov("bul").Conn("wbs", godip.Sea).Conn("rum", godip.Coast...).Conn("ser", godip.Land).Conn("mac", godip.Land).Conn("gre", godip.Land).Conn("ist", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Sirte
		Prov("gos").Conn("tri", godip.Sea).Conn("tob", godip.Sea).Conn("egy", godip.Sea).Conn("cre", godip.Sea).Conn("ion", godip.Sea).Conn("mal", godip.Sea).Flag(godip.Sea).
		// Morocco
		Prov("mor").Conn("tan", godip.Coast...).Conn("sao", godip.Sea).Conn("atl", godip.Coast...).Conn("soa", godip.Coast...).Conn("gov", godip.Sea).Conn("str", godip.Sea).Flag(godip.Coast...).
		// Serbia
		Prov("ser").Conn("cro", godip.Land).Conn("bos", godip.Land).Conn("mon", godip.Land).Conn("mac", godip.Land).Conn("bul", godip.Land).Conn("rum", godip.Land).Conn("tra", godip.Land).Conn("hun", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Belgium
		Prov("bel").Conn("hol", godip.Coast...).Conn("not", godip.Sea).Conn("eng", godip.Sea).Conn("pic", godip.Coast...).Conn("bur", godip.Land).Conn("als", godip.Land).Conn("col", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Latvia
		Prov("lat").Conn("gob", godip.Sea).Conn("bal", godip.Sea).Conn("lit", godip.Coast...).Conn("lub", godip.Land).Conn("bie", godip.Land).Conn("len", godip.Coast...).Conn("est", godip.Coast...).Flag(godip.Coast...).
		// Hungary
		Prov("hun").Conn("tra", godip.Land).Conn("sla", godip.Land).Conn("vie", godip.Land).Conn("sle", godip.Land).Conn("cro", godip.Land).Conn("ser", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Slovakia
		Prov("sla").Conn("tra", godip.Land).Conn("kra", godip.Land).Conn("wro", godip.Land).Conn("boh", godip.Land).Conn("vie", godip.Land).Conn("hun", godip.Land).Flag(godip.Land).
		// Ireland
		Prov("ire").Conn("mid", godip.Sea).Conn("iri", godip.Sea).Conn("noi", godip.Coast...).Flag(godip.Coast...).
		// Skaggerack
		Prov("ska").Conn("swe", godip.Sea).Conn("noa", godip.Sea).Conn("not", godip.Sea).Conn("den", godip.Sea).Flag(godip.Sea).
		// Upper Caspian Sea
		Prov("upp").Conn("tur", godip.Sea).Conn("kaz", godip.Sea).Conn("sta", godip.Sea).Conn("cau", godip.Sea).Conn("low", godip.Sea).Flag(godip.Sea).
		// Liverpool
		Prov("liv").Conn("nao", godip.Sea).Conn("iri", godip.Sea).Conn("wal", godip.Coast...).Conn("lon", godip.Land).Conn("yor", godip.Land).Conn("edi", godip.Land).Conn("cly", godip.Coast...).Flag(godip.Coast...).
		// Nizhniy Novgorod
		Prov("niz").Conn("mos", godip.Land).Conn("sta", godip.Land).Conn("kaz", godip.Land).Conn("sib", godip.Land).Conn("ark", godip.Land).Conn("len", godip.Land).Flag(godip.Land).
		// Denmark
		Prov("den").Conn("ska", godip.Sea).Conn("not", godip.Sea).Conn("hel", godip.Sea).Conn("kie", godip.Coast...).Conn("bal", godip.Sea).Conn("swe", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Adriatic Sea
		Prov("adr").Conn("cro", godip.Sea).Conn("ven", godip.Sea).Conn("apu", godip.Sea).Conn("ion", godip.Sea).Conn("alb", godip.Sea).Conn("mon", godip.Sea).Flag(godip.Sea).
		// Tobruk
		Prov("tob").Conn("fez", godip.Land).Conn("cyr", godip.Land).Conn("ela", godip.Coast...).Conn("egy", godip.Sea).Conn("gos", godip.Sea).Conn("tri", godip.Coast...).Flag(godip.Coast...).
		// Heligoland Bight
		Prov("hel").Conn("hol", godip.Sea).Conn("kie", godip.Sea).Conn("den", godip.Sea).Conn("not", godip.Sea).Flag(godip.Sea).
		// English Channel
		Prov("eng").Conn("lon", godip.Sea).Conn("wal", godip.Sea).Conn("iri", godip.Sea).Conn("coa", godip.Sea).Conn("bay", godip.Sea).Conn("bre", godip.Sea).Conn("pic", godip.Sea).Conn("bel", godip.Sea).Conn("not", godip.Sea).Flag(godip.Sea).
		// Caucasus
		Prov("cau").Conn("ros", godip.Coast...).Conn("ebs", godip.Sea).Conn("ank", godip.Coast...).Conn("arm", godip.Land).Conn("teh", godip.Coast...).Conn("low", godip.Sea).Conn("upp", godip.Sea).Conn("sta", godip.Coast...).Flag(godip.Coast...).
		// Armenia
		Prov("arm").Conn("ada", godip.Land).Conn("syr", godip.Land).Conn("bag", godip.Land).Conn("khu", godip.Land).Conn("teh", godip.Land).Conn("cau", godip.Land).Conn("ank", godip.Land).Flag(godip.Land).
		// Navarra
		Prov("nav").Conn("mad", godip.Land).Conn("bac", godip.Land).Conn("auv", godip.Land).Conn("gas", godip.Coast...).Conn("bay", godip.Sea).Conn("coa", godip.Sea).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// Madrid
		Prov("mad").Conn("and", godip.Land).Conn("bac", godip.Land).Conn("nav", godip.Land).Conn("por", godip.Land).Conn("sei", godip.Land).Flag(godip.Land).SC(Spain).
		// Holland
		Prov("hol").Conn("col", godip.Land).Conn("kie", godip.Coast...).Conn("hel", godip.Sea).Conn("not", godip.Sea).Conn("bel", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Sicily
		Prov("sic").Conn("tys", godip.Sea).Conn("mal", godip.Sea).Conn("tys", godip.Sea).Flag(godip.Coast...).
		// Istanbul
		Prov("ist").Conn("izm", godip.Coast...).Conn("ank", godip.Coast...).Conn("wbs", godip.Sea).Conn("bul", godip.Coast...).Conn("gre", godip.Coast...).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Western Mediterranean
		Prov("wem").Conn("alg", godip.Sea).Conn("tun", godip.Sea).Conn("tys", godip.Sea).Conn("lig", godip.Sea).Conn("gol", godip.Sea).Conn("gov", godip.Sea).Flag(godip.Sea).
		// Lower Caspian Sea
		Prov("low").Conn("cau", godip.Sea).Conn("teh", godip.Sea).Conn("tur", godip.Sea).Conn("upp", godip.Sea).Flag(godip.Sea).
		// Tyrrhenian Sea
		Prov("tys").Conn("sic", godip.Sea).Conn("sic", godip.Sea).Conn("mal", godip.Sea).Conn("nap", godip.Sea).Conn("rom", godip.Sea).Conn("lig", godip.Sea).Conn("wem", godip.Sea).Conn("tun", godip.Sea).Conn("mal", godip.Sea).Flag(godip.Sea).
		// Fezzan
		Prov("fez").Conn("cyr", godip.Land).Conn("tob", godip.Land).Conn("tri", godip.Land).Conn("tun", godip.Land).Conn("ora", godip.Land).Flag(godip.Land).
		// Seville
		Prov("sei").Conn("str", godip.Sea).Conn("and", godip.Coast...).Conn("mad", godip.Land).Conn("por", godip.Coast...).Conn("sao", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Narvik
		Prov("nar").Conn("fin", godip.Land).Conn("ark", godip.Coast...).Conn("bas", godip.Sea).Conn("noe", godip.Sea).Conn("noa", godip.Coast...).Conn("swe", godip.Land).Flag(godip.Coast...).
		// Gulf of Bothnia
		Prov("gob").Conn("fin", godip.Sea).Conn("swe", godip.Sea).Conn("bal", godip.Sea).Conn("lat", godip.Sea).Conn("est", godip.Sea).Conn("len", godip.Sea).Flag(godip.Sea).
		// Gascony
		Prov("gas").Conn("lyo", godip.Land).Conn("par", godip.Land).Conn("bre", godip.Coast...).Conn("bay", godip.Sea).Conn("nav", godip.Coast...).Conn("auv", godip.Land).Flag(godip.Coast...).
		// Izmir
		Prov("izm").Conn("ist", godip.Coast...).Conn("aeg", godip.Sea).Conn("eam", godip.Sea).Conn("ada", godip.Coast...).Conn("ank", godip.Land).Flag(godip.Coast...).SC(Turkey).
		// Cologne
		Prov("col").Conn("kie", godip.Land).Conn("hol", godip.Land).Conn("bel", godip.Land).Conn("als", godip.Land).Conn("mun", godip.Land).Flag(godip.Land).SC(Germany).
		// Andalucia
		Prov("and").Conn("mad", godip.Land).Conn("sei", godip.Coast...).Conn("str", godip.Sea).Conn("gov", godip.Sea).Conn("gol", godip.Sea).Conn("bac", godip.Coast...).Flag(godip.Coast...).
		// North Sea
		Prov("not").Conn("noe", godip.Sea).Conn("edi", godip.Sea).Conn("yor", godip.Sea).Conn("lon", godip.Sea).Conn("eng", godip.Sea).Conn("bel", godip.Sea).Conn("hol", godip.Sea).Conn("hel", godip.Sea).Conn("den", godip.Sea).Conn("ska", godip.Sea).Conn("noa", godip.Sea).Flag(godip.Sea).
		// Auvergne
		Prov("auv").Conn("mar", godip.Coast...).Conn("lyo", godip.Land).Conn("gas", godip.Land).Conn("nav", godip.Land).Conn("bac", godip.Coast...).Conn("gol", godip.Sea).Conn("lig", godip.Sea).Flag(godip.Coast...).
		// Marseilles
		Prov("mar").Conn("auv", godip.Coast...).Conn("lig", godip.Sea).Conn("pie", godip.Coast...).Conn("swi", godip.Land).Conn("lyo", godip.Land).Flag(godip.Coast...).SC(France).
		// Ukraine
		Prov("ukr").Conn("mol", godip.Land).Conn("sea", godip.Land).Conn("mos", godip.Land).Conn("bie", godip.Land).Conn("lub", godip.Land).Conn("kra", godip.Land).Conn("rum", godip.Land).Flag(godip.Land).
		// Mid Atlantic Ocean
		Prov("mid").Conn("azo", godip.Sea).Conn("coa", godip.Sea).Conn("iri", godip.Sea).Conn("ire", godip.Sea).Conn("noi", godip.Sea).Conn("nao", godip.Sea).Flag(godip.Sea).
		// Persian Gulf
		Prov("per").Conn("khu", godip.Sea).Conn("bag", godip.Sea).Conn("raf", godip.Sea).Flag(godip.Sea).
		// Algiers
		Prov("alg").Conn("tun", godip.Coast...).Conn("wem", godip.Sea).Conn("gov", godip.Sea).Conn("soa", godip.Coast...).Conn("ora", godip.Land).Flag(godip.Coast...).SC(France).
		// Gulf of Valencia
		Prov("gov").Conn("str", godip.Sea).Conn("mor", godip.Sea).Conn("soa", godip.Sea).Conn("alg", godip.Sea).Conn("wem", godip.Sea).Conn("gol", godip.Sea).Conn("and", godip.Sea).Flag(godip.Sea).
		// South Atlantic Ocean
		Prov("sao").Conn("atl", godip.Sea).Conn("mor", godip.Sea).Conn("tan", godip.Sea).Conn("str", godip.Sea).Conn("sei", godip.Sea).Conn("por", godip.Sea).Conn("coa", godip.Sea).Conn("azo", godip.Sea).Flag(godip.Sea).
		// North Atlantic Ocean
		Prov("nao").Conn("mid", godip.Sea).Conn("noi", godip.Sea).Conn("iri", godip.Sea).Conn("liv", godip.Sea).Conn("cly", godip.Sea).Conn("noe", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Sea).
		// Northern Ireland
		Prov("noi").Conn("nao", godip.Sea).Conn("mid", godip.Sea).Conn("ire", godip.Coast...).Conn("iri", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Tehran
		Prov("teh").Conn("kho", godip.Land).Conn("tur", godip.Coast...).Conn("low", godip.Sea).Conn("cau", godip.Coast...).Conn("arm", godip.Land).Conn("khu", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Malta Sea
		Prov("mal").Conn("tun", godip.Sea).Conn("tri", godip.Sea).Conn("gos", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Sea).Conn("tys", godip.Sea).Conn("sic", godip.Sea).Conn("tys", godip.Sea).Flag(godip.Sea).
		// Kazakhstan
		Prov("kaz").Conn("sib", godip.Land).Conn("niz", godip.Land).Conn("sta", godip.Coast...).Conn("upp", godip.Sea).Conn("tur", godip.Coast...).Flag(godip.Coast...).
		// Warsaw
		Prov("war").Conn("lub", godip.Land).Conn("eap", godip.Land).Conn("dan", godip.Land).Conn("poz", godip.Land).Conn("wro", godip.Land).Conn("kra", godip.Land).Flag(godip.Land).SC(Poland).
		// Norwegian Sea
		Prov("noe").Conn("ice", godip.Sea).Conn("nao", godip.Sea).Conn("cly", godip.Sea).Conn("edi", godip.Sea).Conn("not", godip.Sea).Conn("noa", godip.Sea).Conn("nar", godip.Sea).Conn("bas", godip.Sea).Flag(godip.Sea).
		// Lithuania
		Prov("lit").Conn("bal", godip.Sea).Conn("eap", godip.Coast...).Conn("lub", godip.Land).Conn("lat", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Tangiers
		Prov("tan").Conn("mor", godip.Coast...).Conn("str", godip.Sea).Conn("sao", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Egyptian Sea
		Prov("egy").Conn("eam", godip.Sea).Conn("cre", godip.Sea).Conn("gos", godip.Sea).Conn("tob", godip.Sea).Conn("ela", godip.Sea).Conn("cai", godip.Sea).Conn("pal", godip.Sea).Flag(godip.Sea).
		// Tripoli
		Prov("tri").Conn("gos", godip.Sea).Conn("mal", godip.Sea).Conn("tun", godip.Coast...).Conn("fez", godip.Land).Conn("tob", godip.Coast...).Flag(godip.Coast...).SC(Italy).
		// Wales
		Prov("wal").Conn("eng", godip.Sea).Conn("lon", godip.Coast...).Conn("liv", godip.Coast...).Conn("iri", godip.Sea).Flag(godip.Coast...).
		// Cyrenacia
		Prov("cyr").Conn("cai", godip.Land).Conn("ela", godip.Land).Conn("tob", godip.Land).Conn("fez", godip.Land).Flag(godip.Land).
		// Greece
		Prov("gre").Conn("aeg", godip.Sea).Conn("ist", godip.Coast...).Conn("bul", godip.Land).Conn("mac", godip.Land).Conn("alb", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Venice
		Prov("ven").Conn("sle", godip.Land).Conn("tyo", godip.Land).Conn("mil", godip.Land).Conn("tus", godip.Land).Conn("rom", godip.Land).Conn("apu", godip.Coast...).Conn("adr", godip.Sea).Conn("cro", godip.Coast...).Flag(godip.Coast...).
		// East Prussia
		Prov("eap").Conn("war", godip.Land).Conn("lub", godip.Land).Conn("lit", godip.Coast...).Conn("bal", godip.Sea).Conn("dan", godip.Coast...).Flag(godip.Coast...).
		// Cairo
		Prov("cai").Conn("red", godip.Sea).Conn("hed", godip.Coast...).Conn("pal", godip.Coast...).Conn("egy", godip.Sea).Conn("ela", godip.Coast...).Conn("cyr", godip.Land).Flag(godip.Coast...).SC(Britain).
		// Vienna
		Prov("vie").Conn("hun", godip.Land).Conn("sla", godip.Land).Conn("boh", godip.Land).Conn("tyo", godip.Land).Conn("sle", godip.Land).Flag(godip.Land).SC(Germany).
		// Transylvania
		Prov("tra").Conn("sla", godip.Land).Conn("hun", godip.Land).Conn("ser", godip.Land).Conn("rum", godip.Land).Conn("kra", godip.Land).Flag(godip.Land).
		// Sweden
		Prov("swe").Conn("ska", godip.Sea).Conn("den", godip.Coast...).Conn("bal", godip.Sea).Conn("gob", godip.Sea).Conn("fin", godip.Coast...).Conn("nar", godip.Land).Conn("noa", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Crete
		Prov("cre").Conn("ion", godip.Sea).Conn("gos", godip.Sea).Conn("egy", godip.Sea).Conn("eam", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Croatia
		Prov("cro").Conn("mon", godip.Coast...).Conn("bos", godip.Land).Conn("ser", godip.Land).Conn("hun", godip.Land).Conn("sle", godip.Land).Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Bohemia
		Prov("boh").Conn("vie", godip.Land).Conn("sla", godip.Land).Conn("wro", godip.Land).Conn("sil", godip.Land).Conn("mun", godip.Land).Conn("tyo", godip.Land).Flag(godip.Land).
		// Krakow
		Prov("kra").Conn("rum", godip.Land).Conn("ukr", godip.Land).Conn("lub", godip.Land).Conn("war", godip.Land).Conn("wro", godip.Land).Conn("sla", godip.Land).Conn("tra", godip.Land).Flag(godip.Land).SC(Poland).
		// Rumania
		Prov("rum").Conn("kra", godip.Land).Conn("tra", godip.Land).Conn("ser", godip.Land).Conn("bul", godip.Coast...).Conn("wbs", godip.Sea).Conn("mol", godip.Coast...).Conn("ukr", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Milan
		Prov("mil").Conn("pie", godip.Land).Conn("tus", godip.Land).Conn("ven", godip.Land).Conn("tyo", godip.Land).Conn("swi", godip.Land).Flag(godip.Land).SC(Italy).
		// Macedonia
		Prov("mac").Conn("ser", godip.Land).Conn("mon", godip.Land).Conn("alb", godip.Land).Conn("gre", godip.Land).Conn("bul", godip.Land).Flag(godip.Land).
		// Khuzestan
		Prov("khu").Conn("kho", godip.Land).Conn("teh", godip.Land).Conn("arm", godip.Land).Conn("bag", godip.Coast...).Conn("per", godip.Sea).Flag(godip.Coast...).
		// Alsace
		Prov("als").Conn("col", godip.Land).Conn("bel", godip.Land).Conn("bur", godip.Land).Conn("swi", godip.Land).Conn("mun", godip.Land).Flag(godip.Land).
		// Tunisia
		Prov("tun").Conn("mal", godip.Sea).Conn("tys", godip.Sea).Conn("wem", godip.Sea).Conn("alg", godip.Coast...).Conn("ora", godip.Land).Conn("fez", godip.Land).Conn("tri", godip.Coast...).Flag(godip.Coast...).
		// Southern Algeria
		Prov("soa").Conn("atl", godip.Land).Conn("ora", godip.Land).Conn("alg", godip.Coast...).Conn("gov", godip.Sea).Conn("mor", godip.Coast...).Flag(godip.Coast...).
		// Barcelona
		Prov("bac").Conn("gol", godip.Sea).Conn("auv", godip.Coast...).Conn("nav", godip.Land).Conn("mad", godip.Land).Conn("and", godip.Coast...).Flag(godip.Coast...).SC(Spain).
		// Turkmenistan
		Prov("tur").Conn("kaz", godip.Coast...).Conn("upp", godip.Sea).Conn("low", godip.Sea).Conn("teh", godip.Coast...).Conn("kho", godip.Land).Flag(godip.Coast...).
		// Bosnia
		Prov("bos").Conn("mon", godip.Land).Conn("ser", godip.Land).Conn("cro", godip.Land).Flag(godip.Land).
		// Clyde
		Prov("cly").Conn("edi", godip.Coast...).Conn("noe", godip.Sea).Conn("nao", godip.Sea).Conn("liv", godip.Coast...).Flag(godip.Coast...).
		// Apulia
		Prov("apu").Conn("nap", godip.Coast...).Conn("ion", godip.Sea).Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("rom", godip.Land).Flag(godip.Coast...).
		// Edinburgh
		Prov("edi").Conn("cly", godip.Coast...).Conn("liv", godip.Land).Conn("yor", godip.Coast...).Conn("not", godip.Sea).Conn("noe", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Bay of Biscay
		Prov("bay").Conn("coa", godip.Sea).Conn("nav", godip.Sea).Conn("gas", godip.Sea).Conn("bre", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Sea).
		// Siberia
		Prov("sib").Conn("ark", godip.Land).Conn("niz", godip.Land).Conn("kaz", godip.Land).Flag(godip.Land).
		// Napels
		Prov("nap").Conn("ion", godip.Sea).Conn("apu", godip.Coast...).Conn("rom", godip.Coast...).Conn("tys", godip.Sea).Conn("mal", godip.Sea).Flag(godip.Coast...).SC(Italy).
		// Gulf of Lyon
		Prov("gol").Conn("auv", godip.Sea).Conn("bac", godip.Sea).Conn("and", godip.Sea).Conn("gov", godip.Sea).Conn("wem", godip.Sea).Conn("lig", godip.Sea).Flag(godip.Sea).
		// Irish Sea
		Prov("iri").Conn("mid", godip.Sea).Conn("coa", godip.Sea).Conn("eng", godip.Sea).Conn("wal", godip.Sea).Conn("liv", godip.Sea).Conn("nao", godip.Sea).Conn("noi", godip.Sea).Conn("ire", godip.Sea).Flag(godip.Sea).
		// Finland
		Prov("fin").Conn("gob", godip.Sea).Conn("len", godip.Coast...).Conn("ark", godip.Land).Conn("nar", godip.Land).Conn("swe", godip.Coast...).Flag(godip.Coast...).
		// Estonia
		Prov("est").Conn("len", godip.Coast...).Conn("gob", godip.Sea).Conn("lat", godip.Coast...).Flag(godip.Coast...).
		// Prussia
		Prov("pru").Conn("sil", godip.Land).Conn("poz", godip.Land).Conn("dan", godip.Coast...).Conn("bal", godip.Sea).Conn("ber", godip.Coast...).Flag(godip.Coast...).
		// Berlin
		Prov("ber").Conn("bal", godip.Sea).Conn("kie", godip.Coast...).Conn("mun", godip.Land).Conn("sil", godip.Land).Conn("pru", godip.Coast...).Flag(godip.Coast...).SC(Germany).
		// Ligurian Sea
		Prov("lig").Conn("tys", godip.Sea).Conn("rom", godip.Sea).Conn("tus", godip.Sea).Conn("pie", godip.Sea).Conn("mar", godip.Sea).Conn("auv", godip.Sea).Conn("gol", godip.Sea).Conn("wem", godip.Sea).Flag(godip.Sea).
		// Burgundy
		Prov("bur").Conn("lyo", godip.Land).Conn("swi", godip.Land).Conn("als", godip.Land).Conn("bel", godip.Land).Conn("pic", godip.Land).Conn("par", godip.Land).Flag(godip.Land).
		// Hedjaz
		Prov("hed").Conn("ara", godip.Land).Conn("pal", godip.Land).Conn("cai", godip.Coast...).Conn("red", godip.Sea).Flag(godip.Coast...).
		// Lyon
		Prov("lyo").Conn("bur", godip.Land).Conn("par", godip.Land).Conn("gas", godip.Land).Conn("auv", godip.Land).Conn("mar", godip.Land).Conn("swi", godip.Land).Flag(godip.Land).SC(France).
		// El Alamein
		Prov("ela").Conn("cyr", godip.Land).Conn("cai", godip.Coast...).Conn("egy", godip.Sea).Conn("tob", godip.Coast...).Flag(godip.Coast...).
		// Azores
		Prov("azo").Conn("sao", godip.Sea).Conn("coa", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Oran
		Prov("ora").Conn("fez", godip.Land).Conn("tun", godip.Land).Conn("alg", godip.Land).Conn("soa", godip.Land).Conn("atl", godip.Land).Flag(godip.Land).
		// Kiel
		Prov("kie").Conn("col", godip.Land).Conn("mun", godip.Land).Conn("ber", godip.Coast...).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Conn("hel", godip.Sea).Conn("hol", godip.Coast...).Flag(godip.Coast...).SC(Germany).
		// Moscow
		Prov("mos").Conn("len", godip.Land).Conn("bie", godip.Land).Conn("ukr", godip.Land).Conn("sea", godip.Land).Conn("sta", godip.Land).Conn("niz", godip.Land).Flag(godip.Land).SC(USSR).
		// Straights of Gibraltar
		Prov("str").Conn("gov", godip.Sea).Conn("and", godip.Sea).Conn("sei", godip.Sea).Conn("sao", godip.Sea).Conn("tan", godip.Sea).Conn("mor", godip.Sea).Flag(godip.Sea).
		// Picardy
		Prov("pic").Conn("bre", godip.Coast...).Conn("par", godip.Land).Conn("bur", godip.Land).Conn("bel", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).
		// Munich
		Prov("mun").Conn("sil", godip.Land).Conn("ber", godip.Land).Conn("kie", godip.Land).Conn("col", godip.Land).Conn("als", godip.Land).Conn("swi", godip.Land).Conn("tyo", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).SC(Germany).
		Done()
}
