package westernworld901

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
	UmayyadEmirate      godip.Nation = "Umayyad Emirate"
	PrincipalityofKiev  godip.Nation = "Principality of Kiev"
	KingdomofDenmark    godip.Nation = "Kingdom of Denmark"
	KhaganateofKhazaria godip.Nation = "Khaganate of Khazaria"
	WestFrankishKingdom godip.Nation = "West Frankish Kingdom"
	TulunidEmirate      godip.Nation = "Tulunid Emirate"
	AbbasidCaliphate    godip.Nation = "Abbasid Caliphate"
	EastFrankishKingdom godip.Nation = "East Frankish Kingdom"
	EasternRomanEmpire  godip.Nation = "Eastern Roman Empire"
)

var Nations = []godip.Nation{UmayyadEmirate, PrincipalityofKiev, KingdomofDenmark, KhaganateofKhazaria, WestFrankishKingdom, TulunidEmirate, AbbasidCaliphate, EastFrankishKingdom, EasternRomanEmpire}

var newPhase = phase.Generator(hundred.BuildAnywhereParser, classical.AdjustSCs)

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return newPhase(year, season, typ)
}

var WesternWorld901Variant = common.Variant{
	Name:       "Western World 901",
	Graph:      func() godip.Graph { return WesternWorld901Graph() },
	Start:      WesternWorld901Start,
	Blank:      WesternWorld901Blank,
	Phase:      newPhase,
	Parser:     hundred.BuildAnywhereParser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(33),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/westernworld901map.svg")
	},
	SVGVersion: "4",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "David Cohen",
	Version:     "4.0",
	Description: "Nine powers compete for the Western World circa 901.",
	Rules: `First to 33 Supply Centers (SC) wins.
Units may be built in any owned SC.
Neutral SCs get an army which always holds and disbands when dislodged. This will be rebuilt if the SC is unowned during adjustment.
Five provinces have dual coasts: Saamiland, Veletia, Jorvik, Rome and Pechenega.
Constantinople has a canal as in the standard map.
The Khazar Sea is not connected to other sea regions.`,
}

func NeutralOrders(state state.State) (ret map[godip.Province]godip.Adjudicator) {
	ret = map[godip.Province]godip.Adjudicator{}
	switch state.Phase().Type() {
	case godip.Movement:
		// Strictly this is unnecessary - because hold is the default order.
		for prov, unit := range state.Units() {
			if unit.Nation == godip.Neutral {
				ret[prov] = orders.Hold(prov)
			}
		}
	case godip.Adjustment:
		// Rebuild any missing units.
		for _, prov := range state.Graph().AllSCs() {
			if n, _, ok := state.SupplyCenter(prov); ok && n == godip.Neutral {
				if _, _, ok := state.Unit(prov); !ok {
					ret[prov] = orders.BuildAnywhere(prov, godip.Army, time.Now())
				}
			}
		}
	}
	return
}

func WesternWorld901Blank(phase godip.Phase) *state.State {
	return state.New(WesternWorld901Graph(), phase, classical.BackupRule, NeutralOrders)
}

func WesternWorld901Start() (result *state.State, err error) {
	startPhase := Phase(901, godip.Spring, godip.Movement)
	result = WesternWorld901Blank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"cad":    godip.Unit{godip.Fleet, UmayyadEmirate},
		"val":    godip.Unit{godip.Fleet, UmayyadEmirate},
		"cod":    godip.Unit{godip.Army, UmayyadEmirate},
		"snc":    godip.Unit{godip.Army, UmayyadEmirate},
		"nov":    godip.Unit{godip.Fleet, PrincipalityofKiev},
		"kie":    godip.Unit{godip.Army, PrincipalityofKiev},
		"ros":    godip.Unit{godip.Army, PrincipalityofKiev},
		"smo":    godip.Unit{godip.Army, PrincipalityofKiev},
		"jel":    godip.Unit{godip.Fleet, KingdomofDenmark},
		"jor/ec": godip.Unit{godip.Fleet, KingdomofDenmark},
		"sca":    godip.Unit{godip.Fleet, KingdomofDenmark},
		"vik":    godip.Unit{godip.Army, KingdomofDenmark},
		"ati":    godip.Unit{godip.Army, KhaganateofKhazaria},
		"bnj":    godip.Unit{godip.Army, KhaganateofKhazaria},
		"sak":    godip.Unit{godip.Army, KhaganateofKhazaria},
		"tam":    godip.Unit{godip.Army, KhaganateofKhazaria},
		"par":    godip.Unit{godip.Fleet, WestFrankishKingdom},
		"aqt":    godip.Unit{godip.Army, WestFrankishKingdom},
		"gas":    godip.Unit{godip.Army, WestFrankishKingdom},
		"nar":    godip.Unit{godip.Army, WestFrankishKingdom},
		"bar":    godip.Unit{godip.Fleet, TulunidEmirate},
		"jer":    godip.Unit{godip.Fleet, TulunidEmirate},
		"ale":    godip.Unit{godip.Army, TulunidEmirate},
		"dam":    godip.Unit{godip.Army, TulunidEmirate},
		"ard":    godip.Unit{godip.Army, AbbasidCaliphate},
		"bag":    godip.Unit{godip.Army, AbbasidCaliphate},
		"isf":    godip.Unit{godip.Army, AbbasidCaliphate},
		"ira":    godip.Unit{godip.Army, AbbasidCaliphate},
		"bre":    godip.Unit{godip.Fleet, EastFrankishKingdom},
		"bav":    godip.Unit{godip.Army, EastFrankishKingdom},
		"sax":    godip.Unit{godip.Army, EastFrankishKingdom},
		"swa":    godip.Unit{godip.Army, EastFrankishKingdom},
		"att":    godip.Unit{godip.Fleet, EasternRomanEmpire},
		"tar":    godip.Unit{godip.Fleet, EasternRomanEmpire},
		"crn":    godip.Unit{godip.Fleet, EasternRomanEmpire},
		"con":    godip.Unit{godip.Army, EasternRomanEmpire},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"cad": UmayyadEmirate,
		"val": UmayyadEmirate,
		"cod": UmayyadEmirate,
		"snc": UmayyadEmirate,
		"nov": PrincipalityofKiev,
		"kie": PrincipalityofKiev,
		"ros": PrincipalityofKiev,
		"smo": PrincipalityofKiev,
		"jel": KingdomofDenmark,
		"jor": KingdomofDenmark,
		"sca": KingdomofDenmark,
		"vik": KingdomofDenmark,
		"ati": KhaganateofKhazaria,
		"bnj": KhaganateofKhazaria,
		"sak": KhaganateofKhazaria,
		"tam": KhaganateofKhazaria,
		"par": WestFrankishKingdom,
		"aqt": WestFrankishKingdom,
		"gas": WestFrankishKingdom,
		"nar": WestFrankishKingdom,
		"bar": TulunidEmirate,
		"jer": TulunidEmirate,
		"ale": TulunidEmirate,
		"dam": TulunidEmirate,
		"ard": AbbasidCaliphate,
		"bag": AbbasidCaliphate,
		"isf": AbbasidCaliphate,
		"ira": AbbasidCaliphate,
		"bre": EastFrankishKingdom,
		"bav": EastFrankishKingdom,
		"sax": EastFrankishKingdom,
		"swa": EastFrankishKingdom,
		"att": EasternRomanEmpire,
		"tar": EasternRomanEmpire,
		"crn": EasternRomanEmpire,
		"con": EasternRomanEmpire,
		"bor": godip.Neutral,
		"pam": godip.Neutral,
		"cyp": godip.Neutral,
		"bas": godip.Neutral,
		"btt": godip.Neutral,
		"maz": godip.Neutral,
		"dub": godip.Neutral,
		"dal": godip.Neutral,
		"mav": godip.Neutral,
		"low": godip.Neutral,
		"sad": godip.Neutral,
		"rom": godip.Neutral,
		"geo": godip.Neutral,
		"aze": godip.Neutral,
		"lot": godip.Neutral,
		"pec": godip.Neutral,
		"arm": godip.Neutral,
		"ifr": godip.Neutral,
		"cos": godip.Neutral,
		"sic": godip.Neutral,
		"est": godip.Neutral,
		"thr": godip.Neutral,
		"bja": godip.Neutral,
		"urg": godip.Neutral,
		"wsx": godip.Neutral,
		"cre": godip.Neutral,
		"bul": godip.Neutral,
		"mau": godip.Neutral,
	})
	for _, sc := range WesternWorld901Graph().SCs(godip.Neutral) {
		if err = result.SetUnit(godip.Province(sc), godip.Unit{
			Type:   godip.Army,
			Nation: godip.Neutral,
		}); err != nil {
			return
		}
	}
	return
}

func WesternWorld901Graph() *graph.Graph {
	return graph.New().
		// Bulgar
		Prov("bul").Conn("vya", godip.Land).Conn("mod", godip.Land).Conn("udm", godip.Land).Conn("kom", godip.Land).Conn("chm", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Franconia
		Prov("fra").Conn("bre", godip.Land).Conn("fri", godip.Land).Conn("lot", godip.Land).Conn("swa", godip.Land).Conn("sax", godip.Land).Flag(godip.Land).
		// Toledo
		Prov("tol").Conn("zar", godip.Land).Conn("snc", godip.Land).Conn("cod", godip.Land).Conn("val", godip.Land).Flag(godip.Land).
		// Mordvinia
		Prov("mod").Conn("vya", godip.Land).Conn("sev", godip.Land).Conn("sak", godip.Land).Conn("ati", godip.Land).Conn("udm", godip.Land).Conn("bul", godip.Land).Flag(godip.Land).
		// Alexandria
		Prov("ale").Conn("egy", godip.Sea).Conn("bar", godip.Coast...).Conn("tri", godip.Land).Conn("faz", godip.Land).Conn("aqa", godip.Land).Conn("mec", godip.Land).Conn("jer", godip.Coast...).Flag(godip.Coast...).SC(TulunidEmirate).
		// Kakheti
		Prov("kak").Conn("bnj", godip.Land).Conn("geo", godip.Land).Conn("cap", godip.Land).Conn("arm", godip.Land).Conn("mos", godip.Land).Conn("aze", godip.Land).Conn("der", godip.Land).Flag(godip.Land).
		// Thrace
		Prov("thr").Conn("ono", godip.Land).Conn("mac", godip.Land).Conn("con", godip.Coast...).Conn("wes", godip.Sea).Conn("vla", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Munster
		Prov("mun").Conn("dub", godip.Coast...).Conn("ice", godip.Sea).Conn("oce", godip.Sea).Flag(godip.Coast...).
		// Krivichia
		Prov("kri").Conn("smo", godip.Land).Conn("vya", godip.Land).Conn("chm", godip.Land).Conn("ros", godip.Land).Conn("nov", godip.Land).Flag(godip.Land).
		// Kipchak
		Prov("kip").Conn("sto", godip.Land).Conn("bas", godip.Land).Conn("ati", godip.Coast...).Conn("nks", godip.Sea).Conn("urg", godip.Coast...).Conn("kyz", godip.Land).Flag(godip.Coast...).
		// Scania
		Prov("sca").Conn("abo", godip.Sea).Conn("let", godip.Sea).Conn("got", godip.Coast...).Conn("vik", godip.Coast...).Conn("kat", godip.Sea).Flag(godip.Coast...).SC(KingdomofDenmark).
		// Lower Burgundy
		Prov("low").Conn("aut", godip.Land).Conn("nar", godip.Coast...).Conn("lig", godip.Sea).Conn("lom", godip.Coast...).Conn("ueg", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Abkhazia
		Prov("abk").Conn("geo", godip.Coast...).Conn("bnj", godip.Land).Conn("tam", godip.Coast...).Conn("eas", godip.Sea).Flag(godip.Coast...).
		// Galicia
		Prov("gal").Conn("oce", godip.Sea).Conn("cad", godip.Coast...).Conn("snc", godip.Land).Conn("ast", godip.Coast...).Conn("can", godip.Sea).Flag(godip.Coast...).
		// Libyan Sea
		Prov("lib").Conn("sic", godip.Sea).Conn("tyr", godip.Sea).Conn("ifr", godip.Sea).Conn("tri", godip.Sea).Conn("bar", godip.Sea).Conn("egy", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// Bjarmaland
		Prov("bja").Conn("saa", godip.Land).Conn("saa/nc", godip.Sea).Conn("kar", godip.Land).Conn("chm", godip.Land).Conn("kom", godip.Coast...).Conn("whi", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Rostov
		Prov("ros").Conn("chm", godip.Land).Conn("kar", godip.Land).Conn("nov", godip.Land).Conn("kri", godip.Land).Flag(godip.Land).SC(PrincipalityofKiev).
		// Stone Belt
		Prov("sto").Conn("kom", godip.Land).Conn("udm", godip.Land).Conn("bas", godip.Land).Conn("kip", godip.Land).Flag(godip.Land).
		// Kutamia
		Prov("kut").Conn("tah", godip.Land).Conn("ifr", godip.Coast...).Conn("str", godip.Sea).Conn("mau", godip.Coast...).Flag(godip.Coast...).
		// Slavonia
		Prov("sla").Conn("mav", godip.Land).Conn("bav", godip.Land).Conn("aql", godip.Land).Conn("dal", godip.Land).Conn("ono", godip.Land).Flag(godip.Land).
		// Wales
		Prov("wal").Conn("wel", godip.Sea).Conn("wsx", godip.Coast...).Conn("jor", godip.Land).Conn("jor/wc", godip.Sea).Flag(godip.Coast...).
		// Lettish Sea
		Prov("let").Conn("bor", godip.Sea).Conn("liv", godip.Sea).Conn("fin", godip.Sea).Conn("ula", godip.Sea).Conn("got", godip.Sea).Conn("sca", godip.Sea).Conn("abo", godip.Sea).Flag(godip.Sea).
		// Rome
		Prov("rom").Conn("sle", godip.Land).Conn("spo", godip.Land).Conn("aql", godip.Land).Conn("lom", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Rome (West Coast)
		Prov("rom/wc").Conn("sle", godip.Sea).Conn("lom", godip.Sea).Conn("lig", godip.Sea).Conn("tyr", godip.Sea).Flag(godip.Sea).
		// Rome (East Coast)
		Prov("rom/ec").Conn("spo", godip.Sea).Conn("ill", godip.Sea).Conn("aql", godip.Sea).Flag(godip.Sea).
		// Sarkel
		Prov("sak").Conn("tam", godip.Land).Conn("bnj", godip.Land).Conn("ati", godip.Land).Conn("mod", godip.Land).Conn("sev", godip.Land).Flag(godip.Land).SC(KhaganateofKhazaria).
		// Jerusalem
		Prov("jer").Conn("egy", godip.Sea).Conn("ale", godip.Coast...).Conn("mec", godip.Land).Conn("jaz", godip.Land).Conn("dam", godip.Coast...).Conn("sey", godip.Sea).Flag(godip.Coast...).SC(TulunidEmirate).
		// Zaragoza
		Prov("zar").Conn("tol", godip.Land).Conn("val", godip.Land).Conn("spa", godip.Land).Conn("pam", godip.Land).Conn("ast", godip.Land).Conn("snc", godip.Land).Flag(godip.Land).
		// Mauretania
		Prov("mau").Conn("tah", godip.Land).Conn("kut", godip.Coast...).Conn("str", godip.Sea).Conn("soa", godip.Sea).Conn("bah", godip.Coast...).Conn("sss", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Paris
		Prov("par").Conn("lot", godip.Coast...).Conn("bsa", godip.Sea).Conn("btt", godip.Coast...).Conn("aqt", godip.Land).Conn("aut", godip.Land).Flag(godip.Coast...).SC(WestFrankishKingdom).
		// Ionian Sea
		Prov("ion").Conn("tar", godip.Sea).Conn("tyr", godip.Sea).Conn("sic", godip.Sea).Conn("lib", godip.Sea).Conn("egy", godip.Sea).Conn("cre", godip.Sea).Conn("aeg", godip.Sea).Conn("con", godip.Sea).Conn("epi", godip.Sea).Conn("ill", godip.Sea).Flag(godip.Sea).
		// Derbent
		Prov("der").Conn("kak", godip.Land).Conn("aze", godip.Coast...).Conn("sks", godip.Sea).Conn("nks", godip.Sea).Conn("bnj", godip.Coast...).Flag(godip.Coast...).
		// Lothairingia
		Prov("lot").Conn("sgs", godip.Sea).Conn("bsa", godip.Sea).Conn("par", godip.Coast...).Conn("aut", godip.Land).Conn("ueg", godip.Land).Conn("swa", godip.Land).Conn("fra", godip.Land).Conn("fri", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Baghdad
		Prov("bag").Conn("ard", godip.Land).Conn("mos", godip.Land).Conn("jaz", godip.Land).Conn("nef", godip.Land).Conn("ira", godip.Land).Flag(godip.Land).SC(AbbasidCaliphate).
		// Autun
		Prov("aut").Conn("aqt", godip.Land).Conn("tou", godip.Land).Conn("nar", godip.Land).Conn("low", godip.Land).Conn("ueg", godip.Land).Conn("lot", godip.Land).Conn("par", godip.Land).Flag(godip.Land).
		// Narbonne
		Prov("nar").Conn("aut", godip.Land).Conn("tou", godip.Land).Conn("spa", godip.Coast...).Conn("lig", godip.Sea).Conn("low", godip.Coast...).Flag(godip.Coast...).SC(WestFrankishKingdom).
		// Sea of Tangiers
		Prov("soa").Conn("bah", godip.Sea).Conn("mau", godip.Sea).Conn("str", godip.Sea).Conn("cad", godip.Sea).Conn("oce", godip.Sea).Flag(godip.Sea).
		// Moravia
		Prov("mav").Conn("pol", godip.Land).Conn("bav", godip.Land).Conn("sla", godip.Land).Conn("ono", godip.Land).Conn("vol", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// West Euxine Sea
		Prov("wes").Conn("con", godip.Sea).Conn("buc", godip.Sea).Conn("eas", godip.Sea).Conn("crn", godip.Sea).Conn("pec", godip.Sea).Conn("pec/wc", godip.Sea).Conn("vla", godip.Sea).Conn("thr", godip.Sea).Flag(godip.Sea).
		// Komia
		Prov("kom").Conn("whi", godip.Sea).Conn("bja", godip.Coast...).Conn("chm", godip.Land).Conn("bul", godip.Land).Conn("udm", godip.Land).Conn("sto", godip.Land).Flag(godip.Coast...).
		// Saxony
		Prov("sax").Conn("swa", godip.Land).Conn("bav", godip.Land).Conn("pol", godip.Land).Conn("pom", godip.Land).Conn("vel", godip.Land).Conn("bre", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).SC(EastFrankishKingdom).
		// Aegean Sea
		Prov("aeg").Conn("ion", godip.Sea).Conn("cre", godip.Sea).Conn("cil", godip.Sea).Conn("att", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
		// Azerbaijan
		Prov("aze").Conn("sks", godip.Sea).Conn("der", godip.Coast...).Conn("kak", godip.Land).Conn("mos", godip.Land).Conn("ard", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Jorvik
		Prov("jor").Conn("cal", godip.Land).Conn("wal", godip.Land).Conn("wsx", godip.Land).Flag(godip.Land).SC(KingdomofDenmark).
		// Jorvik (West Coast)
		Prov("jor/wc").Conn("cal", godip.Sea).Conn("wel", godip.Sea).Conn("wal", godip.Sea).Flag(godip.Sea).
		// Jorvik (East Coast)
		Prov("jor/ec").Conn("ngs", godip.Sea).Conn("ice", godip.Sea).Conn("cal", godip.Sea).Conn("wsx", godip.Sea).Flag(godip.Sea).
		// Icelandic Sea
		Prov("ice").Conn("oce", godip.Sea).Conn("mun", godip.Sea).Conn("dub", godip.Sea).Conn("wel", godip.Sea).Conn("cal", godip.Sea).Conn("jor", godip.Sea).Conn("jor/ec", godip.Sea).Conn("ngs", godip.Sea).Conn("now", godip.Sea).Conn("saa", godip.Sea).Conn("saa/nc", godip.Sea).Conn("bja", godip.Sea).Conn("whi", godip.Sea).Flag(godip.Sea).
		// Ifriqiya
		Prov("ifr").Conn("tah", godip.Land).Conn("faz", godip.Land).Conn("tri", godip.Coast...).Conn("lib", godip.Sea).Conn("tyr", godip.Sea).Conn("str", godip.Sea).Conn("kut", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Gottland
		Prov("got").Conn("ula", godip.Coast...).Conn("vik", godip.Land).Conn("sca", godip.Coast...).Conn("let", godip.Sea).Flag(godip.Coast...).
		// White Sea
		Prov("whi").Conn("ice", godip.Sea).Conn("bja", godip.Sea).Conn("kom", godip.Sea).Flag(godip.Sea).
		// Salamanca
		Prov("snc").Conn("ast", godip.Land).Conn("gal", godip.Land).Conn("cad", godip.Land).Conn("cod", godip.Land).Conn("tol", godip.Land).Conn("zar", godip.Land).Flag(godip.Land).SC(UmayyadEmirate).
		// Bashkortostan
		Prov("bas").Conn("kip", godip.Land).Conn("sto", godip.Land).Conn("udm", godip.Land).Conn("ati", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Vyatichia
		Prov("vya").Conn("mod", godip.Land).Conn("bul", godip.Land).Conn("chm", godip.Land).Conn("kri", godip.Land).Conn("smo", godip.Land).Conn("kie", godip.Land).Conn("sev", godip.Land).Flag(godip.Land).
		// Smolensk
		Prov("smo").Conn("kri", godip.Land).Conn("nov", godip.Land).Conn("liv", godip.Land).Conn("dre", godip.Land).Conn("kie", godip.Land).Conn("vya", godip.Land).Flag(godip.Land).SC(PrincipalityofKiev).
		// Norway
		Prov("now").Conn("kat", godip.Sea).Conn("vik", godip.Coast...).Conn("ula", godip.Land).Conn("saa", godip.Land).Conn("saa/nc", godip.Sea).Conn("ice", godip.Sea).Conn("ngs", godip.Sea).Flag(godip.Coast...).
		// Welsh Sea
		Prov("wel").Conn("dub", godip.Sea).Conn("oce", godip.Sea).Conn("wsx", godip.Sea).Conn("wal", godip.Sea).Conn("jor", godip.Sea).Conn("jor/wc", godip.Sea).Conn("cal", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Sea).
		// Bavaria
		Prov("bav").Conn("hel", godip.Land).Conn("lom", godip.Land).Conn("aql", godip.Land).Conn("sla", godip.Land).Conn("mav", godip.Land).Conn("pol", godip.Land).Conn("sax", godip.Land).Conn("swa", godip.Land).Flag(godip.Land).SC(EastFrankishKingdom).
		// Veletia
		Prov("vel").Conn("bre", godip.Land).Conn("sax", godip.Land).Conn("pom", godip.Land).Conn("jel", godip.Land).Flag(godip.Land).
		// Veletia (North Coast)
		Prov("vel/nc").Conn("pom", godip.Sea).Conn("abo", godip.Sea).Conn("jel", godip.Sea).Flag(godip.Sea).
		// Veletia (West Coast)
		Prov("vel/wc").Conn("sgs", godip.Sea).Conn("bre", godip.Sea).Conn("jel", godip.Sea).Flag(godip.Sea).
		// Al-Qatta'i
		Prov("aqa").Conn("mec", godip.Land).Conn("ale", godip.Land).Conn("faz", godip.Land).Flag(godip.Land).
		// Khorasan
		Prov("kho").Conn("kyz", godip.Land).Conn("urg", godip.Land).Conn("ali", godip.Land).Conn("sst", godip.Land).Flag(godip.Land).
		// Pamplona
		Prov("pam").Conn("ast", godip.Coast...).Conn("zar", godip.Land).Conn("spa", godip.Land).Conn("tou", godip.Land).Conn("gas", godip.Coast...).Conn("can", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Polania
		Prov("pol").Conn("mav", godip.Land).Conn("vol", godip.Land).Conn("maz", godip.Land).Conn("bor", godip.Land).Conn("pom", godip.Land).Conn("sax", godip.Land).Conn("bav", godip.Land).Flag(godip.Land).
		// Novgorod
		Prov("nov").Conn("ros", godip.Land).Conn("kar", godip.Coast...).Conn("fin", godip.Sea).Conn("est", godip.Coast...).Conn("liv", godip.Land).Conn("smo", godip.Land).Conn("kri", godip.Land).Flag(godip.Coast...).SC(PrincipalityofKiev).
		// Fazzan
		Prov("faz").Conn("aqa", godip.Land).Conn("ale", godip.Land).Conn("tri", godip.Land).Conn("ifr", godip.Land).Conn("tah", godip.Land).Flag(godip.Land).
		// Karelia
		Prov("kar").Conn("bja", godip.Land).Conn("saa", godip.Land).Conn("saa/sc", godip.Sea).Conn("fin", godip.Sea).Conn("nov", godip.Coast...).Conn("ros", godip.Land).Conn("chm", godip.Land).Flag(godip.Coast...).
		// South German Sea
		Prov("sgs").Conn("jel", godip.Sea).Conn("ngs", godip.Sea).Conn("wsx", godip.Sea).Conn("bsa", godip.Sea).Conn("lot", godip.Sea).Conn("fri", godip.Sea).Conn("bre", godip.Sea).Conn("vel", godip.Sea).Conn("vel/wc", godip.Sea).Flag(godip.Sea).
		// Ocean Sea
		Prov("oce").Conn("soa", godip.Sea).Conn("cad", godip.Sea).Conn("gal", godip.Sea).Conn("can", godip.Sea).Conn("btt", godip.Sea).Conn("bsa", godip.Sea).Conn("wsx", godip.Sea).Conn("wel", godip.Sea).Conn("dub", godip.Sea).Conn("mun", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Sea).
		// Straits of Jebel Tarik
		Prov("str").Conn("bcs", godip.Sea).Conn("val", godip.Sea).Conn("gra", godip.Sea).Conn("cad", godip.Sea).Conn("soa", godip.Sea).Conn("mau", godip.Sea).Conn("kut", godip.Sea).Conn("ifr", godip.Sea).Conn("tyr", godip.Sea).Conn("sad", godip.Sea).Flag(godip.Sea).
		// Brittany
		Prov("btt").Conn("can", godip.Sea).Conn("aqt", godip.Coast...).Conn("par", godip.Coast...).Conn("bsa", godip.Sea).Conn("oce", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Valencia
		Prov("val").Conn("bcs", godip.Sea).Conn("spa", godip.Coast...).Conn("zar", godip.Land).Conn("tol", godip.Land).Conn("cod", godip.Land).Conn("gra", godip.Coast...).Conn("str", godip.Sea).Flag(godip.Coast...).SC(UmayyadEmirate).
		// Urgench
		Prov("urg").Conn("kip", godip.Coast...).Conn("nks", godip.Sea).Conn("sks", godip.Sea).Conn("ali", godip.Coast...).Conn("kho", godip.Land).Conn("kyz", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// North Khazar Sea
		Prov("nks").Conn("kip", godip.Sea).Conn("ati", godip.Sea).Conn("bnj", godip.Sea).Conn("der", godip.Sea).Conn("sks", godip.Sea).Conn("urg", godip.Sea).Flag(godip.Sea).
		// Mazovia
		Prov("maz").Conn("bor", godip.Land).Conn("pol", godip.Land).Conn("vol", godip.Land).Conn("dre", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Balanjar
		Prov("bnj").Conn("kak", godip.Land).Conn("der", godip.Coast...).Conn("nks", godip.Sea).Conn("ati", godip.Coast...).Conn("sak", godip.Land).Conn("tam", godip.Land).Conn("abk", godip.Land).Conn("geo", godip.Land).Flag(godip.Coast...).SC(KhaganateofKhazaria).
		// South Khazar Sea
		Prov("sks").Conn("ard", godip.Sea).Conn("ali", godip.Sea).Conn("urg", godip.Sea).Conn("nks", godip.Sea).Conn("der", godip.Sea).Conn("aze", godip.Sea).Flag(godip.Sea).
		// Finnish Sea
		Prov("fin").Conn("ula", godip.Sea).Conn("let", godip.Sea).Conn("liv", godip.Sea).Conn("est", godip.Sea).Conn("nov", godip.Sea).Conn("kar", godip.Sea).Conn("saa", godip.Sea).Conn("saa/sc", godip.Sea).Flag(godip.Sea).
		// Sicily
		Prov("sic").Conn("lib", godip.Sea).Conn("ion", godip.Sea).Conn("tyr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Jelling
		Prov("jel").Conn("sgs", godip.Sea).Conn("vel", godip.Land).Conn("vel/wc", godip.Sea).Conn("vel/nc", godip.Sea).Conn("abo", godip.Sea).Conn("kat", godip.Sea).Conn("ngs", godip.Sea).Flag(godip.Coast...).SC(KingdomofDenmark).
		// Tyrrhenian Sea
		Prov("tyr").Conn("rom", godip.Sea).Conn("rom/wc", godip.Sea).Conn("lig", godip.Sea).Conn("cos", godip.Sea).Conn("bcs", godip.Sea).Conn("sad", godip.Sea).Conn("str", godip.Sea).Conn("ifr", godip.Sea).Conn("lib", godip.Sea).Conn("sic", godip.Sea).Conn("ion", godip.Sea).Conn("tar", godip.Sea).Conn("sle", godip.Sea).Flag(godip.Sea).
		// Kattegat
		Prov("kat").Conn("abo", godip.Sea).Conn("sca", godip.Sea).Conn("vik", godip.Sea).Conn("now", godip.Sea).Conn("ngs", godip.Sea).Conn("jel", godip.Sea).Flag(godip.Sea).
		// Caledonia
		Prov("cal").Conn("jor", godip.Land).Conn("jor/wc", godip.Sea).Conn("jor/ec", godip.Sea).Conn("ice", godip.Sea).Conn("wel", godip.Sea).Flag(godip.Coast...).
		// Cyprus
		Prov("cyp").Conn("sey", godip.Sea).Conn("cil", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Salerno
		Prov("sle").Conn("rom", godip.Land).Conn("rom/wc", godip.Sea).Conn("tyr", godip.Sea).Conn("tar", godip.Coast...).Conn("spo", godip.Land).Flag(godip.Coast...).
		// Gascony
		Prov("gas").Conn("aqt", godip.Coast...).Conn("can", godip.Sea).Conn("pam", godip.Coast...).Conn("tou", godip.Land).Flag(godip.Coast...).SC(WestFrankishKingdom).
		// Cadiz
		Prov("cad").Conn("soa", godip.Sea).Conn("str", godip.Sea).Conn("gra", godip.Coast...).Conn("cod", godip.Land).Conn("snc", godip.Land).Conn("gal", godip.Coast...).Conn("oce", godip.Sea).Flag(godip.Coast...).SC(UmayyadEmirate).
		// Illyrian Sea
		Prov("ill").Conn("aql", godip.Sea).Conn("rom", godip.Sea).Conn("rom/ec", godip.Sea).Conn("spo", godip.Sea).Conn("tar", godip.Sea).Conn("ion", godip.Sea).Conn("epi", godip.Sea).Conn("dal", godip.Sea).Flag(godip.Sea).
		// Borussia
		Prov("bor").Conn("maz", godip.Land).Conn("dre", godip.Land).Conn("liv", godip.Coast...).Conn("let", godip.Sea).Conn("abo", godip.Sea).Conn("pom", godip.Coast...).Conn("pol", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Friesland
		Prov("fri").Conn("fra", godip.Land).Conn("bre", godip.Coast...).Conn("sgs", godip.Sea).Conn("lot", godip.Coast...).Flag(godip.Coast...).
		// Sardinia
		Prov("sad").Conn("bcs", godip.Sea).Conn("str", godip.Sea).Conn("tyr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Constantinople
		Prov("con").Conn("epi", godip.Coast...).Conn("ion", godip.Sea).Conn("aeg", godip.Sea).Conn("att", godip.Coast...).Conn("buc", godip.Coast...).Conn("wes", godip.Sea).Conn("thr", godip.Coast...).Conn("mac", godip.Land).Flag(godip.Coast...).SC(EasternRomanEmpire).
		// Vlacha
		Prov("vla").Conn("kie", godip.Land).Conn("vol", godip.Land).Conn("ono", godip.Land).Conn("thr", godip.Coast...).Conn("wes", godip.Sea).Conn("pec", godip.Land).Conn("pec/wc", godip.Sea).Flag(godip.Coast...).
		// Jazira
		Prov("jaz").Conn("mec", godip.Land).Conn("nef", godip.Land).Conn("bag", godip.Land).Conn("mos", godip.Land).Conn("arm", godip.Land).Conn("cap", godip.Land).Conn("dam", godip.Land).Conn("jer", godip.Land).Flag(godip.Land).
		// Taranto
		Prov("tar").Conn("ion", godip.Sea).Conn("ill", godip.Sea).Conn("spo", godip.Coast...).Conn("sle", godip.Coast...).Conn("tyr", godip.Sea).Flag(godip.Coast...).SC(EasternRomanEmpire).
		// Dalmatia
		Prov("dal").Conn("epi", godip.Coast...).Conn("mac", godip.Land).Conn("ono", godip.Land).Conn("sla", godip.Land).Conn("aql", godip.Coast...).Conn("ill", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cantabric Sea
		Prov("can").Conn("btt", godip.Sea).Conn("oce", godip.Sea).Conn("gal", godip.Sea).Conn("ast", godip.Sea).Conn("pam", godip.Sea).Conn("gas", godip.Sea).Conn("aqt", godip.Sea).Flag(godip.Sea).
		// Granada
		Prov("gra").Conn("str", godip.Sea).Conn("val", godip.Coast...).Conn("cod", godip.Land).Conn("cad", godip.Coast...).Flag(godip.Coast...).
		// Sea of Tyre
		Prov("sey").Conn("dam", godip.Sea).Conn("cap", godip.Sea).Conn("cil", godip.Sea).Conn("cyp", godip.Sea).Conn("egy", godip.Sea).Conn("jer", godip.Sea).Flag(godip.Sea).
		// Sijilmassa
		Prov("sss").Conn("tah", godip.Land).Conn("mau", godip.Land).Conn("bah", godip.Land).Flag(godip.Land).
		// Swabia
		Prov("swa").Conn("sax", godip.Land).Conn("fra", godip.Land).Conn("lot", godip.Land).Conn("ueg", godip.Land).Conn("hel", godip.Land).Conn("bav", godip.Land).Flag(godip.Land).SC(EastFrankishKingdom).
		// Attalia
		Prov("att").Conn("con", godip.Coast...).Conn("aeg", godip.Sea).Conn("cil", godip.Sea).Conn("cap", godip.Coast...).Conn("buc", godip.Land).Flag(godip.Coast...).SC(EasternRomanEmpire).
		// Pechenega
		Prov("pec").Conn("tam", godip.Land).Conn("sev", godip.Land).Conn("kie", godip.Land).Conn("vla", godip.Land).Conn("crn", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Pechenega (West Coast)
		Prov("pec/wc").Conn("vla", godip.Sea).Conn("wes", godip.Sea).Conn("crn", godip.Sea).Flag(godip.Sea).
		// Pechenega (East Coast)
		Prov("pec/ec").Conn("eas", godip.Sea).Conn("tam", godip.Sea).Conn("crn", godip.Sea).Flag(godip.Sea).
		// Irak
		Prov("ira").Conn("bag", godip.Land).Conn("nef", godip.Land).Conn("bsr", godip.Land).Conn("isf", godip.Land).Conn("ard", godip.Land).Flag(godip.Land).SC(AbbasidCaliphate).
		// Toulouse
		Prov("tou").Conn("aut", godip.Land).Conn("aqt", godip.Land).Conn("gas", godip.Land).Conn("pam", godip.Land).Conn("spa", godip.Land).Conn("nar", godip.Land).Flag(godip.Land).
		// Corsica
		Prov("cos").Conn("bcs", godip.Sea).Conn("tyr", godip.Sea).Conn("lig", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Alidistan
		Prov("ali").Conn("ard", godip.Coast...).Conn("isf", godip.Land).Conn("sst", godip.Land).Conn("kho", godip.Land).Conn("urg", godip.Coast...).Conn("sks", godip.Sea).Flag(godip.Coast...).
		// Atil
		Prov("ati").Conn("bas", godip.Land).Conn("udm", godip.Land).Conn("mod", godip.Land).Conn("sak", godip.Land).Conn("bnj", godip.Coast...).Conn("nks", godip.Sea).Conn("kip", godip.Coast...).Flag(godip.Coast...).SC(KhaganateofKhazaria).
		// Egyptian Sea
		Prov("egy").Conn("jer", godip.Sea).Conn("sey", godip.Sea).Conn("cyp", godip.Sea).Conn("cil", godip.Sea).Conn("cre", godip.Sea).Conn("ion", godip.Sea).Conn("lib", godip.Sea).Conn("bar", godip.Sea).Conn("ale", godip.Sea).Flag(godip.Sea).
		// Balearic Sea
		Prov("bcs").Conn("str", godip.Sea).Conn("sad", godip.Sea).Conn("tyr", godip.Sea).Conn("cos", godip.Sea).Conn("lig", godip.Sea).Conn("spa", godip.Sea).Conn("val", godip.Sea).Flag(godip.Sea).
		// Armenia
		Prov("arm").Conn("jaz", godip.Land).Conn("mos", godip.Land).Conn("kak", godip.Land).Conn("cap", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wessex
		Prov("wsx").Conn("oce", godip.Sea).Conn("bsa", godip.Sea).Conn("sgs", godip.Sea).Conn("ngs", godip.Sea).Conn("jor", godip.Land).Conn("jor/ec", godip.Sea).Conn("wal", godip.Coast...).Conn("wel", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Volhynia
		Prov("vol").Conn("maz", godip.Land).Conn("pol", godip.Land).Conn("mav", godip.Land).Conn("ono", godip.Land).Conn("vla", godip.Land).Conn("kie", godip.Land).Conn("dre", godip.Land).Flag(godip.Land).
		// Cordova
		Prov("cod").Conn("tol", godip.Land).Conn("snc", godip.Land).Conn("cad", godip.Land).Conn("gra", godip.Land).Conn("val", godip.Land).Flag(godip.Land).SC(UmayyadEmirate).
		// Aquileia
		Prov("aql").Conn("sla", godip.Land).Conn("bav", godip.Land).Conn("lom", godip.Land).Conn("rom", godip.Land).Conn("rom/ec", godip.Sea).Conn("ill", godip.Sea).Conn("dal", godip.Coast...).Flag(godip.Coast...).
		// Cheremissia
		Prov("chm").Conn("ros", godip.Land).Conn("kri", godip.Land).Conn("vya", godip.Land).Conn("bul", godip.Land).Conn("kom", godip.Land).Conn("bja", godip.Land).Conn("kar", godip.Land).Flag(godip.Land).
		// Georgia
		Prov("geo").Conn("eas", godip.Sea).Conn("buc", godip.Coast...).Conn("cap", godip.Land).Conn("kak", godip.Land).Conn("bnj", godip.Land).Conn("abk", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Lombardy
		Prov("lom").Conn("bav", godip.Land).Conn("hel", godip.Land).Conn("ueg", godip.Land).Conn("low", godip.Coast...).Conn("lig", godip.Sea).Conn("rom", godip.Land).Conn("rom/wc", godip.Sea).Conn("aql", godip.Land).Flag(godip.Coast...).
		// Spoleto
		Prov("spo").Conn("rom", godip.Land).Conn("rom/ec", godip.Sea).Conn("sle", godip.Land).Conn("tar", godip.Coast...).Conn("ill", godip.Sea).Flag(godip.Coast...).
		// Basra
		Prov("bsr").Conn("sst", godip.Land).Conn("isf", godip.Land).Conn("ira", godip.Land).Conn("nef", godip.Land).Flag(godip.Land).
		// Pomerania
		Prov("pom").Conn("bor", godip.Coast...).Conn("abo", godip.Sea).Conn("vel", godip.Land).Conn("vel/nc", godip.Sea).Conn("sax", godip.Land).Conn("pol", godip.Land).Flag(godip.Coast...).
		// East Euxine Sea
		Prov("eas").Conn("geo", godip.Sea).Conn("abk", godip.Sea).Conn("tam", godip.Sea).Conn("pec", godip.Sea).Conn("pec/ec", godip.Sea).Conn("crn", godip.Sea).Conn("wes", godip.Sea).Conn("buc", godip.Sea).Flag(godip.Sea).
		// Cherson
		Prov("crn").Conn("eas", godip.Sea).Conn("pec", godip.Land).Conn("pec/wc", godip.Sea).Conn("pec/ec", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(EasternRomanEmpire).
		// Ardebil
		Prov("ard").Conn("bag", godip.Land).Conn("ira", godip.Land).Conn("isf", godip.Land).Conn("ali", godip.Coast...).Conn("sks", godip.Sea).Conn("aze", godip.Coast...).Conn("mos", godip.Land).Flag(godip.Coast...).SC(AbbasidCaliphate).
		// Crete
		Prov("cre").Conn("egy", godip.Sea).Conn("cil", godip.Sea).Conn("aeg", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cilician Sea
		Prov("cil").Conn("sey", godip.Sea).Conn("cap", godip.Sea).Conn("att", godip.Sea).Conn("aeg", godip.Sea).Conn("cre", godip.Sea).Conn("egy", godip.Sea).Conn("cyp", godip.Sea).Flag(godip.Sea).
		// Saamiland
		Prov("saa").Conn("bja", godip.Land).Conn("now", godip.Land).Conn("ula", godip.Land).Conn("kar", godip.Land).Flag(godip.Land).
		// Saamiland (North Coast)
		Prov("saa/nc").Conn("bja", godip.Sea).Conn("ice", godip.Sea).Conn("now", godip.Sea).Flag(godip.Sea).
		// Saamiland (South Coast)
		Prov("saa/sc").Conn("ula", godip.Sea).Conn("fin", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Epirus
		Prov("epi").Conn("con", godip.Coast...).Conn("mac", godip.Land).Conn("dal", godip.Coast...).Conn("ill", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).
		// Bucellaria
		Prov("buc").Conn("con", godip.Coast...).Conn("att", godip.Land).Conn("cap", godip.Land).Conn("geo", godip.Coast...).Conn("eas", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Coast...).
		// Isfahan
		Prov("isf").Conn("bsr", godip.Land).Conn("sst", godip.Land).Conn("ali", godip.Land).Conn("ard", godip.Land).Conn("ira", godip.Land).Flag(godip.Land).SC(AbbasidCaliphate).
		// Udmurtia
		Prov("udm").Conn("bul", godip.Land).Conn("mod", godip.Land).Conn("ati", godip.Land).Conn("bas", godip.Land).Conn("sto", godip.Land).Conn("kom", godip.Land).Flag(godip.Land).
		// Dublin
		Prov("dub").Conn("mun", godip.Coast...).Conn("oce", godip.Sea).Conn("wel", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Upper Burgundy
		Prov("ueg").Conn("swa", godip.Land).Conn("lot", godip.Land).Conn("aut", godip.Land).Conn("low", godip.Land).Conn("lom", godip.Land).Conn("hel", godip.Land).Flag(godip.Land).
		// Bremen
		Prov("bre").Conn("sax", godip.Land).Conn("vel", godip.Land).Conn("vel/wc", godip.Sea).Conn("sgs", godip.Sea).Conn("fri", godip.Coast...).Conn("fra", godip.Land).Flag(godip.Coast...).SC(EastFrankishKingdom).
		// Asturias
		Prov("ast").Conn("pam", godip.Coast...).Conn("can", godip.Sea).Conn("gal", godip.Coast...).Conn("snc", godip.Land).Conn("zar", godip.Land).Flag(godip.Coast...).
		// Cappadocia
		Prov("cap").Conn("dam", godip.Coast...).Conn("jaz", godip.Land).Conn("arm", godip.Land).Conn("kak", godip.Land).Conn("geo", godip.Land).Conn("buc", godip.Land).Conn("att", godip.Coast...).Conn("cil", godip.Sea).Conn("sey", godip.Sea).Flag(godip.Coast...).
		// Tahert
		Prov("tah").Conn("faz", godip.Land).Conn("ifr", godip.Land).Conn("kut", godip.Land).Conn("mau", godip.Land).Conn("sss", godip.Land).Flag(godip.Land).
		// Kyzyl Kum
		Prov("kyz").Conn("kip", godip.Land).Conn("urg", godip.Land).Conn("kho", godip.Land).Flag(godip.Land).
		// Nefud
		Prov("nef").Conn("bsr", godip.Land).Conn("ira", godip.Land).Conn("bag", godip.Land).Conn("jaz", godip.Land).Conn("mec", godip.Land).Flag(godip.Land).
		// Macedonia
		Prov("mac").Conn("thr", godip.Land).Conn("ono", godip.Land).Conn("dal", godip.Land).Conn("epi", godip.Land).Conn("con", godip.Land).Flag(godip.Land).
		// Kiev
		Prov("kie").Conn("pec", godip.Land).Conn("sev", godip.Land).Conn("vya", godip.Land).Conn("smo", godip.Land).Conn("dre", godip.Land).Conn("vol", godip.Land).Conn("vla", godip.Land).Flag(godip.Land).SC(PrincipalityofKiev).
		// Helvetia
		Prov("hel").Conn("bav", godip.Land).Conn("swa", godip.Land).Conn("ueg", godip.Land).Conn("lom", godip.Land).Flag(godip.Land).
		// Dregovichia
		Prov("dre").Conn("liv", godip.Land).Conn("bor", godip.Land).Conn("maz", godip.Land).Conn("vol", godip.Land).Conn("kie", godip.Land).Conn("smo", godip.Land).Flag(godip.Land).
		// Onoguria
		Prov("ono").Conn("thr", godip.Land).Conn("vla", godip.Land).Conn("vol", godip.Land).Conn("mav", godip.Land).Conn("sla", godip.Land).Conn("dal", godip.Land).Conn("mac", godip.Land).Flag(godip.Land).
		// Uppland
		Prov("ula").Conn("got", godip.Coast...).Conn("let", godip.Sea).Conn("fin", godip.Sea).Conn("saa", godip.Land).Conn("saa/sc", godip.Sea).Conn("now", godip.Land).Conn("vik", godip.Land).Flag(godip.Coast...).
		// Spanish March
		Prov("spa").Conn("val", godip.Coast...).Conn("bcs", godip.Sea).Conn("lig", godip.Sea).Conn("nar", godip.Coast...).Conn("tou", godip.Land).Conn("pam", godip.Land).Conn("zar", godip.Land).Flag(godip.Coast...).
		// Viken
		Prov("vik").Conn("ula", godip.Land).Conn("now", godip.Coast...).Conn("kat", godip.Sea).Conn("sca", godip.Coast...).Conn("got", godip.Land).Flag(godip.Coast...).SC(KingdomofDenmark).
		// British Channel
		Prov("bsa").Conn("oce", godip.Sea).Conn("btt", godip.Sea).Conn("par", godip.Sea).Conn("lot", godip.Sea).Conn("sgs", godip.Sea).Conn("wsx", godip.Sea).Flag(godip.Sea).
		// Aquitaine
		Prov("aqt").Conn("gas", godip.Coast...).Conn("tou", godip.Land).Conn("aut", godip.Land).Conn("par", godip.Land).Conn("btt", godip.Coast...).Conn("can", godip.Sea).Flag(godip.Coast...).SC(WestFrankishKingdom).
		// Barghawata
		Prov("bah").Conn("sss", godip.Land).Conn("mau", godip.Coast...).Conn("soa", godip.Sea).Flag(godip.Coast...).
		// Severyana
		Prov("sev").Conn("mod", godip.Land).Conn("vya", godip.Land).Conn("kie", godip.Land).Conn("pec", godip.Land).Conn("tam", godip.Land).Conn("sak", godip.Land).Flag(godip.Land).
		// Ligurian Sea
		Prov("lig").Conn("rom", godip.Sea).Conn("rom/wc", godip.Sea).Conn("lom", godip.Sea).Conn("low", godip.Sea).Conn("nar", godip.Sea).Conn("spa", godip.Sea).Conn("bcs", godip.Sea).Conn("cos", godip.Sea).Conn("tyr", godip.Sea).Flag(godip.Sea).
		// Livonia
		Prov("liv").Conn("est", godip.Coast...).Conn("fin", godip.Sea).Conn("let", godip.Sea).Conn("bor", godip.Coast...).Conn("dre", godip.Land).Conn("smo", godip.Land).Conn("nov", godip.Land).Flag(godip.Coast...).
		// North German Sea
		Prov("ngs").Conn("jor", godip.Sea).Conn("jor/ec", godip.Sea).Conn("wsx", godip.Sea).Conn("sgs", godip.Sea).Conn("jel", godip.Sea).Conn("kat", godip.Sea).Conn("now", godip.Sea).Conn("ice", godip.Sea).Flag(godip.Sea).
		// Tamantarka
		Prov("tam").Conn("sak", godip.Land).Conn("sev", godip.Land).Conn("pec", godip.Land).Conn("pec/ec", godip.Sea).Conn("eas", godip.Sea).Conn("abk", godip.Coast...).Conn("bnj", godip.Land).Flag(godip.Coast...).SC(KhaganateofKhazaria).
		// Mosul
		Prov("mos").Conn("bag", godip.Land).Conn("ard", godip.Land).Conn("aze", godip.Land).Conn("kak", godip.Land).Conn("arm", godip.Land).Conn("jaz", godip.Land).Flag(godip.Land).
		// Mecca
		Prov("mec").Conn("nef", godip.Land).Conn("jaz", godip.Land).Conn("jer", godip.Land).Conn("ale", godip.Land).Conn("aqa", godip.Land).Flag(godip.Land).
		// Abodrite Sea
		Prov("abo").Conn("sca", godip.Sea).Conn("kat", godip.Sea).Conn("jel", godip.Sea).Conn("vel", godip.Sea).Conn("vel/nc", godip.Sea).Conn("pom", godip.Sea).Conn("bor", godip.Sea).Conn("let", godip.Sea).Flag(godip.Sea).
		// Tripolitania
		Prov("tri").Conn("ifr", godip.Coast...).Conn("faz", godip.Land).Conn("ale", godip.Land).Conn("bar", godip.Coast...).Conn("lib", godip.Sea).Flag(godip.Coast...).
		// Esteland
		Prov("est").Conn("liv", godip.Coast...).Conn("nov", godip.Coast...).Conn("fin", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sijistan
		Prov("sst").Conn("kho", godip.Land).Conn("ali", godip.Land).Conn("isf", godip.Land).Conn("bsr", godip.Land).Flag(godip.Land).
		// Barca
		Prov("bar").Conn("tri", godip.Coast...).Conn("ale", godip.Coast...).Conn("egy", godip.Sea).Conn("lib", godip.Sea).Flag(godip.Coast...).SC(TulunidEmirate).
		// Damascus
		Prov("dam").Conn("sey", godip.Sea).Conn("jer", godip.Coast...).Conn("jaz", godip.Land).Conn("cap", godip.Coast...).Flag(godip.Coast...).SC(TulunidEmirate).
		Done()
}
