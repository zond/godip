package empiresandcoalitions

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	OttomanEmpire godip.Nation = "Ottoman Empire"
	Denmark       godip.Nation = "Denmark"
	Sicily        godip.Nation = "Sicily"
	Prussia       godip.Nation = "Prussia"
	Austria       godip.Nation = "Austria"
	France        godip.Nation = "France"
	Britain       godip.Nation = "Britain"
	Russia        godip.Nation = "Russia"
	Spain         godip.Nation = "Spain"
)

var Nations = []godip.Nation{OttomanEmpire, Denmark, Sicily, Prussia, Austria, France, Britain, Russia, Spain}

var EmpiresAndCoalitionsVariant = common.Variant{
	Name:       "1800: Empires And Coalitions",
	Graph:      func() godip.Graph { return EmpiresAndCoalitionsGraph() },
	Start:      EmpiresAndCoalitionsStart,
	Blank:      EmpiresAndCoalitionsBlank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(23),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/empiresandcoalitionsmap.svg")
	},
	SVGVersion: "8",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "VaeVictis",
	Version:     "1",
	Description: "Major and minor powers battle for control of Europe during the Napoleonic Wars.",
	Rules: "The four smallest nations are minor powers and start with only two units. They each have " +
		"an extra center which can be built in once captured. These centers are Sweden (Denmark), " +
		"Papal States (Sicily), Portugal (Spain) and Egypt (Ottoman Empire). The British fleet from " +
		"Liverpool starts in Gibraltar, but note that Gibratar is not a supply center. Armies can " +
		"move between Gibraltar and Andalusia, but may not move direcly between Gibraltar and " +
		"Morocco. Fleets in Gibraltar are considered to be at sea, and so can take part in convoys. " +
		"There are three bridges marked on the map, which connect regions (for both armies and " +
		"fleets). Four regions have dual coasts; these are St. Petersburg, Schleswig, Andalusia and " +
		"Papal States. The winner is the first nation to own 23 of the 44 centers.",
}

func EmpiresAndCoalitionsBlank(phase godip.Phase) *state.State {
	return state.New(EmpiresAndCoalitionsGraph(), phase, classical.BackupRule, nil)
}

func EmpiresAndCoalitionsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1800, godip.Spring, godip.Movement)
	result = EmpiresAndCoalitionsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"ang":    godip.Unit{godip.Fleet, OttomanEmpire},
		"con":    godip.Unit{godip.Army, OttomanEmpire},
		"cop":    godip.Unit{godip.Fleet, Denmark},
		"chr":    godip.Unit{godip.Fleet, Denmark},
		"pal":    godip.Unit{godip.Fleet, Sicily},
		"nap":    godip.Unit{godip.Army, Sicily},
		"ber":    godip.Unit{godip.Army, Prussia},
		"kon":    godip.Unit{godip.Army, Prussia},
		"brl":    godip.Unit{godip.Army, Prussia},
		"vie":    godip.Unit{godip.Army, Austria},
		"ven":    godip.Unit{godip.Army, Austria},
		"bud":    godip.Unit{godip.Army, Austria},
		"mar":    godip.Unit{godip.Fleet, France},
		"par":    godip.Unit{godip.Army, France},
		"brt":    godip.Unit{godip.Army, France},
		"lyo":    godip.Unit{godip.Army, France},
		"edi":    godip.Unit{godip.Fleet, Britain},
		"lon":    godip.Unit{godip.Fleet, Britain},
		"gib":    godip.Unit{godip.Fleet, Britain},
		"han":    godip.Unit{godip.Army, Britain},
		"stp/sc": godip.Unit{godip.Fleet, Russia},
		"sev":    godip.Unit{godip.Fleet, Russia},
		"mos":    godip.Unit{godip.Army, Russia},
		"kie":    godip.Unit{godip.Army, Russia},
		"mad":    godip.Unit{godip.Fleet, Spain},
		"val":    godip.Unit{godip.Fleet, Spain},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"ang": OttomanEmpire,
		"con": OttomanEmpire,
		"cop": Denmark,
		"chr": Denmark,
		"pal": Sicily,
		"nap": Sicily,
		"ber": Prussia,
		"kon": Prussia,
		"brl": Prussia,
		"vie": Austria,
		"ven": Austria,
		"bud": Austria,
		"mar": France,
		"par": France,
		"brt": France,
		"lyo": France,
		"edi": Britain,
		"lon": Britain,
		"han": Britain,
		"lie": Britain,
		"stp": Russia,
		"sev": Russia,
		"mos": Russia,
		"kie": Russia,
		"mad": Spain,
		"val": Spain,
	})
	return
}

func EmpiresAndCoalitionsGraph() *graph.Graph {
	return graph.New().
		// Konigsberg
		Prov("kon").Conn("brl", godip.Land).Conn("pol", godip.Land).Conn("lio", godip.Coast...).Conn("bai", godip.Sea).Conn("ber", godip.Coast...).Flag(godip.Coast...).SC(Prussia).
		// Copenhagen
		Prov("cop").Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Conn("sch", godip.Land).Conn("sch/wc", godip.Sea).Conn("sch/ec", godip.Sea).Conn("bai", godip.Sea).Flag(godip.Coast...).SC(Denmark).
		// Tunisia
		Prov("tun").Conn("tri", godip.Coast...).Conn("cen", godip.Sea).Conn("tys", godip.Sea).Conn("wem", godip.Sea).Conn("alg", godip.Coast...).Flag(godip.Coast...).
		// Sevastopol
		Prov("sev").Conn("ura", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("mol", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Flag(godip.Coast...).SC(Russia).
		// St. Petersburg
		Prov("stp").Conn("fim", godip.Land).Conn("fil", godip.Land).Conn("lio", godip.Land).Conn("mos", godip.Land).Conn("ura", godip.Land).Flag(godip.Land).SC(Russia).
		// St. Petersburg (South Coast)
		Prov("stp/sc").Conn("fil", godip.Sea).Conn("gob", godip.Sea).Conn("lio", godip.Sea).Flag(godip.Sea).
		// St. Petersburg (North Coast)
		Prov("stp/nc").Conn("bar", godip.Sea).Conn("fim", godip.Sea).Flag(godip.Sea).
		// Iceland
		Prov("ice").Conn("noa", godip.Sea).Conn("now", godip.Sea).Conn("arc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Syria
		Prov("syr").Conn("arm", godip.Land).Conn("ang", godip.Coast...).Conn("eas", godip.Sea).Conn("egy", godip.Coast...).Flag(godip.Coast...).
		// London
		Prov("lon").Conn("not", godip.Sea).Conn("yor", godip.Coast...).Conn("wae", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Hanover
		Prov("han").Conn("wet", godip.Coast...).Conn("bra", godip.Land).Conn("mec", godip.Land).Conn("sch", godip.Land).Conn("sch/wc", godip.Sea).Conn("not", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Wales
		Prov("wae").Conn("eng", godip.Sea).Conn("lon", godip.Coast...).Conn("yor", godip.Land).Conn("lie", godip.Coast...).Conn("cel", godip.Sea).Flag(godip.Coast...).
		// Brest
		Prov("brt").Conn("gas", godip.Coast...).Conn("lyo", godip.Land).Conn("par", godip.Land).Conn("bel", godip.Coast...).Conn("eng", godip.Sea).Conn("bay", godip.Sea).Flag(godip.Coast...).SC(France).
		// Bosnia
		Prov("bos").Conn("tra", godip.Land).Conn("bud", godip.Land).Conn("cro", godip.Coast...).Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("gre", godip.Coast...).Conn("con", godip.Land).Conn("waa", godip.Land).Flag(godip.Coast...).
		// Angora
		Prov("ang").Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Conn("con", godip.Coast...).Conn("aeg", godip.Sea).Conn("eas", godip.Sea).Conn("syr", godip.Coast...).Flag(godip.Coast...).SC(OttomanEmpire).
		// Tyrol
		Prov("tyo").Conn("cis", godip.Land).Conn("ven", godip.Land).Conn("vie", godip.Land).Conn("bav", godip.Land).Conn("rhi", godip.Land).Conn("hel", godip.Land).Flag(godip.Land).
		// Paris
		Prov("par").Conn("lyo", godip.Land).Conn("lor", godip.Land).Conn("bel", godip.Land).Conn("brt", godip.Land).Flag(godip.Land).SC(France).
		// Ionian Sea
		Prov("ion").Conn("cen", godip.Sea).Conn("eas", godip.Sea).Conn("aeg", godip.Sea).Conn("gre", godip.Sea).Conn("bos", godip.Sea).Conn("adr", godip.Sea).Conn("apu", godip.Sea).Conn("nap", godip.Sea).Conn("tys", godip.Sea).Conn("pal", godip.Sea).Flag(godip.Sea).
		// Portugal
		Prov("por").Conn("mad", godip.Coast...).Conn("atl", godip.Sea).Conn("and", godip.Land).Conn("and/wc", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Skagerrak
		Prov("ska").Conn("now", godip.Sea).Conn("not", godip.Sea).Conn("sch", godip.Sea).Conn("sch/wc", godip.Sea).Conn("cop", godip.Sea).Conn("swe", godip.Sea).Conn("chr", godip.Sea).Flag(godip.Sea).
		// Saxony
		Prov("sax").Conn("brl", godip.Land).Conn("ber", godip.Land).Conn("bra", godip.Land).Conn("wet", godip.Land).Conn("rhi", godip.Land).Conn("bav", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Aegean Sea
		Prov("aeg").Conn("eas", godip.Sea).Conn("ang", godip.Sea).Conn("con", godip.Sea).Conn("gre", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// Gibraltar
		Prov("gib").Conn("wem", godip.Sea).Conn("and", godip.Land).Conn("and/wc", godip.Sea).Conn("and/ec", godip.Sea).Conn("atl", godip.Sea).Conn("mor", godip.Sea).Flag(godip.Archipelago...).
		// Andalusia
		Prov("and").Conn("val", godip.Land).Conn("mad", godip.Land).Conn("por", godip.Land).Conn("gib", godip.Land).Flag(godip.Land).
		// Andalusia (West Coast)
		Prov("and/wc").Conn("por", godip.Sea).Conn("atl", godip.Sea).Conn("gib", godip.Sea).Flag(godip.Sea).
		// Andalusia (East Coast)
		Prov("and/ec").Conn("val", godip.Sea).Conn("gib", godip.Sea).Conn("wem", godip.Sea).Flag(godip.Sea).
		// Eastern Mediterranean
		Prov("eas").Conn("aeg", godip.Sea).Conn("ion", godip.Sea).Conn("cen", godip.Sea).Conn("tri", godip.Sea).Conn("egy", godip.Sea).Conn("syr", godip.Sea).Conn("ang", godip.Sea).Flag(godip.Sea).
		// Catalonia
		Prov("cat").Conn("gol", godip.Sea).Conn("mar", godip.Coast...).Conn("gas", godip.Land).Conn("nav", godip.Land).Conn("mad", godip.Land).Conn("val", godip.Coast...).Conn("bae", godip.Sea).Flag(godip.Coast...).
		// Brandenburg
		Prov("bra").Conn("wet", godip.Land).Conn("sax", godip.Land).Conn("ber", godip.Land).Conn("mec", godip.Land).Conn("han", godip.Land).Flag(godip.Land).
		// Algeria
		Prov("alg").Conn("mor", godip.Coast...).Conn("tun", godip.Coast...).Conn("wem", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Westphalia
		Prov("wet").Conn("bat", godip.Coast...).Conn("bel", godip.Land).Conn("lor", godip.Land).Conn("rhi", godip.Land).Conn("sax", godip.Land).Conn("bra", godip.Land).Conn("han", godip.Coast...).Conn("not", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Baltic Sea
		Prov("bai").Conn("sch", godip.Sea).Conn("sch/ec", godip.Sea).Conn("mec", godip.Sea).Conn("ber", godip.Sea).Conn("kon", godip.Sea).Conn("lio", godip.Sea).Conn("gob", godip.Sea).Conn("swe", godip.Sea).Conn("cop", godip.Sea).Flag(godip.Sea).
		// Arctic Ocean
		Prov("arc").Conn("noa", godip.Sea).Conn("ice", godip.Sea).Conn("now", godip.Sea).Flag(godip.Sea).
		// Piedmont
		Prov("pie").Conn("tys", godip.Sea).Conn("pap", godip.Land).Conn("pap/wc", godip.Sea).Conn("cis", godip.Land).Conn("hel", godip.Land).Conn("lyo", godip.Land).Conn("mar", godip.Coast...).Conn("gol", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Budapest
		Prov("bud").Conn("gal", godip.Land).Conn("boh", godip.Land).Conn("vie", godip.Land).Conn("cro", godip.Land).Conn("bos", godip.Land).Conn("tra", godip.Land).Flag(godip.Land).SC(Austria).
		// Bavaria
		Prov("bav").Conn("vie", godip.Land).Conn("boh", godip.Land).Conn("sax", godip.Land).Conn("rhi", godip.Land).Conn("tyo", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Moldovia
		Prov("mol").Conn("gal", godip.Land).Conn("tra", godip.Land).Conn("waa", godip.Coast...).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Conn("ukr", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Morocco
		Prov("mor").Conn("alg", godip.Coast...).Conn("wem", godip.Sea).Conn("gib", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Coast...).
		// Belgium
		Prov("bel").Conn("eng", godip.Sea).Conn("brt", godip.Coast...).Conn("par", godip.Land).Conn("lor", godip.Land).Conn("wet", godip.Land).Conn("bat", godip.Coast...).Conn("not", godip.Sea).Flag(godip.Coast...).
		// Rhineland
		Prov("rhi").Conn("lor", godip.Land).Conn("hel", godip.Land).Conn("tyo", godip.Land).Conn("bav", godip.Land).Conn("sax", godip.Land).Conn("wet", godip.Land).Flag(godip.Land).
		// Schleswig
		Prov("sch").Conn("cop", godip.Land).Conn("han", godip.Land).Conn("mec", godip.Land).Flag(godip.Land).
		// Schleswig (West Coast)
		Prov("sch/wc").Conn("cop", godip.Sea).Conn("ska", godip.Sea).Conn("not", godip.Sea).Conn("han", godip.Sea).Flag(godip.Sea).
		// Schleswig (East Coast)
		Prov("sch/ec").Conn("bai", godip.Sea).Conn("cop", godip.Sea).Conn("mec", godip.Sea).Flag(godip.Sea).
		// Ireland
		Prov("ire").Conn("cel", godip.Sea).Conn("cel", godip.Sea).Conn("noa", godip.Sea).Conn("edi", godip.Coast...).Flag(godip.Coast...).
		// Liverpool
		Prov("lie").Conn("wae", godip.Coast...).Conn("yor", godip.Land).Conn("edi", godip.Coast...).Conn("cel", godip.Sea).Flag(godip.Coast...).SC(Britain).
		// Adriatic Sea
		Prov("adr").Conn("cis", godip.Sea).Conn("pap", godip.Sea).Conn("pap/ec", godip.Sea).Conn("apu", godip.Sea).Conn("ion", godip.Sea).Conn("bos", godip.Sea).Conn("cro", godip.Sea).Conn("vie", godip.Sea).Conn("ven", godip.Sea).Flag(godip.Sea).
		// Valencia
		Prov("val").Conn("bae", godip.Sea).Conn("cat", godip.Coast...).Conn("mad", godip.Land).Conn("and", godip.Land).Conn("and/ec", godip.Sea).Conn("wem", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// English Channel
		Prov("eng").Conn("bay", godip.Sea).Conn("brt", godip.Sea).Conn("bel", godip.Sea).Conn("not", godip.Sea).Conn("lon", godip.Sea).Conn("wae", godip.Sea).Conn("cel", godip.Sea).Flag(godip.Sea).
		// Armenia
		Prov("arm").Conn("sev", godip.Coast...).Conn("bla", godip.Sea).Conn("ang", godip.Coast...).Conn("syr", godip.Land).Flag(godip.Coast...).
		// Palermo
		Prov("pal").Conn("tys", godip.Sea).Conn("cen", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Coast...).Flag(godip.Coast...).SC(Sicily).
		// Mecklenburg
		Prov("mec").Conn("bai", godip.Sea).Conn("sch", godip.Land).Conn("sch/ec", godip.Sea).Conn("han", godip.Land).Conn("bra", godip.Land).Conn("ber", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Madrid
		Prov("mad").Conn("nav", godip.Coast...).Conn("atl", godip.Sea).Conn("por", godip.Coast...).Conn("and", godip.Land).Conn("val", godip.Land).Conn("cat", godip.Land).Flag(godip.Coast...).SC(Spain).
		// Celtic Sea
		Prov("cel").Conn("ire", godip.Sea).Conn("noa", godip.Sea).Conn("atl", godip.Sea).Conn("bay", godip.Sea).Conn("eng", godip.Sea).Conn("wae", godip.Sea).Conn("lie", godip.Sea).Conn("edi", godip.Sea).Conn("noa", godip.Sea).Conn("ire", godip.Sea).Flag(godip.Sea).
		// Western Mediterranean
		Prov("wem").Conn("bae", godip.Sea).Conn("val", godip.Sea).Conn("and", godip.Sea).Conn("and/ec", godip.Sea).Conn("gib", godip.Sea).Conn("mor", godip.Sea).Conn("alg", godip.Sea).Conn("tun", godip.Sea).Conn("tys", godip.Sea).Flag(godip.Sea).
		// Tyrrhenian Sea
		Prov("tys").Conn("pal", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Sea).Conn("pap", godip.Sea).Conn("pap/wc", godip.Sea).Conn("pie", godip.Sea).Conn("gol", godip.Sea).Conn("bae", godip.Sea).Conn("wem", godip.Sea).Conn("tun", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Sea).
		// Gulf of Bothnia
		Prov("gob").Conn("fil", godip.Sea).Conn("swe", godip.Sea).Conn("bai", godip.Sea).Conn("lio", godip.Sea).Conn("stp", godip.Sea).Conn("stp/sc", godip.Sea).Flag(godip.Sea).
		// Gascony
		Prov("gas").Conn("lyo", godip.Land).Conn("brt", godip.Coast...).Conn("bay", godip.Sea).Conn("nav", godip.Coast...).Conn("cat", godip.Land).Conn("mar", godip.Land).Flag(godip.Coast...).
		// Poland
		Prov("pol").Conn("brl", godip.Land).Conn("boh", godip.Land).Conn("gal", godip.Land).Conn("kie", godip.Land).Conn("lio", godip.Land).Conn("kon", godip.Land).Flag(godip.Land).
		// North Sea
		Prov("not").Conn("lon", godip.Sea).Conn("eng", godip.Sea).Conn("bel", godip.Sea).Conn("bat", godip.Sea).Conn("wet", godip.Sea).Conn("han", godip.Sea).Conn("sch", godip.Sea).Conn("sch/wc", godip.Sea).Conn("ska", godip.Sea).Conn("now", godip.Sea).Conn("noa", godip.Sea).Conn("edi", godip.Sea).Conn("yor", godip.Sea).Flag(godip.Sea).
		// Urals
		Prov("ura").Conn("stp", godip.Land).Conn("mos", godip.Land).Conn("sev", godip.Land).Flag(godip.Land).
		// Galicia
		Prov("gal").Conn("bud", godip.Land).Conn("tra", godip.Land).Conn("mol", godip.Land).Conn("ukr", godip.Land).Conn("kie", godip.Land).Conn("pol", godip.Land).Conn("boh", godip.Land).Flag(godip.Land).
		// North Atlantic
		Prov("noa").Conn("atl", godip.Sea).Conn("cel", godip.Sea).Conn("ire", godip.Sea).Conn("cel", godip.Sea).Conn("edi", godip.Sea).Conn("not", godip.Sea).Conn("now", godip.Sea).Conn("ice", godip.Sea).Conn("arc", godip.Sea).Flag(godip.Sea).
		// Constantinople
		Prov("con").Conn("bla", godip.Sea).Conn("waa", godip.Coast...).Conn("bos", godip.Land).Conn("gre", godip.Coast...).Conn("aeg", godip.Sea).Conn("ang", godip.Coast...).Flag(godip.Coast...).SC(OttomanEmpire).
		// Marseilles
		Prov("mar").Conn("pie", godip.Coast...).Conn("lyo", godip.Land).Conn("gas", godip.Land).Conn("cat", godip.Coast...).Conn("gol", godip.Sea).Flag(godip.Coast...).SC(France).
		// York
		Prov("yor").Conn("lon", godip.Coast...).Conn("not", godip.Sea).Conn("edi", godip.Coast...).Conn("lie", godip.Land).Conn("wae", godip.Land).Flag(godip.Coast...).
		// Ukraine
		Prov("ukr").Conn("mos", godip.Land).Conn("kie", godip.Land).Conn("gal", godip.Land).Conn("mol", godip.Land).Conn("sev", godip.Land).Flag(godip.Land).
		// Papal States
		Prov("pap").Conn("apu", godip.Land).Conn("cis", godip.Land).Conn("pie", godip.Land).Conn("nap", godip.Land).Flag(godip.Land).SC(Sicily).
		// Papal States (West Coast)
		Prov("pap/wc").Conn("pie", godip.Sea).Conn("tys", godip.Sea).Conn("nap", godip.Sea).Flag(godip.Sea).
		// Papal States (East Coast)
		Prov("pap/ec").Conn("apu", godip.Sea).Conn("adr", godip.Sea).Conn("cis", godip.Sea).Flag(godip.Sea).
		// Christiania
		Prov("chr").Conn("fim", godip.Coast...).Conn("now", godip.Sea).Conn("ska", godip.Sea).Conn("swe", godip.Coast...).Flag(godip.Coast...).SC(Denmark).
		// Norwegian Sea
		Prov("now").Conn("arc", godip.Sea).Conn("ice", godip.Sea).Conn("noa", godip.Sea).Conn("not", godip.Sea).Conn("ska", godip.Sea).Conn("chr", godip.Sea).Conn("fim", godip.Sea).Conn("bar", godip.Sea).Flag(godip.Sea).
		// Balearic Sea
		Prov("bae").Conn("wem", godip.Sea).Conn("tys", godip.Sea).Conn("gol", godip.Sea).Conn("cat", godip.Sea).Conn("val", godip.Sea).Flag(godip.Sea).
		// Atlantic
		Prov("atl").Conn("mor", godip.Sea).Conn("gib", godip.Sea).Conn("and", godip.Sea).Conn("and/wc", godip.Sea).Conn("por", godip.Sea).Conn("mad", godip.Sea).Conn("nav", godip.Sea).Conn("bay", godip.Sea).Conn("cel", godip.Sea).Conn("noa", godip.Sea).Flag(godip.Sea).
		// Wallachia
		Prov("waa").Conn("bla", godip.Sea).Conn("mol", godip.Coast...).Conn("tra", godip.Land).Conn("bos", godip.Land).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Greece
		Prov("gre").Conn("ion", godip.Sea).Conn("aeg", godip.Sea).Conn("con", godip.Coast...).Conn("bos", godip.Coast...).Flag(godip.Coast...).
		// Venice
		Prov("ven").Conn("cis", godip.Coast...).Conn("adr", godip.Sea).Conn("vie", godip.Coast...).Conn("tyo", godip.Land).Flag(godip.Coast...).SC(Austria).
		// Vienna
		Prov("vie").Conn("boh", godip.Land).Conn("bav", godip.Land).Conn("tyo", godip.Land).Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Conn("cro", godip.Coast...).Conn("bud", godip.Land).Flag(godip.Coast...).SC(Austria).
		// Transylvania
		Prov("tra").Conn("bos", godip.Land).Conn("waa", godip.Land).Conn("mol", godip.Land).Conn("gal", godip.Land).Conn("bud", godip.Land).Flag(godip.Land).
		// Finnmark
		Prov("fim").Conn("fil", godip.Land).Conn("stp", godip.Land).Conn("stp/nc", godip.Sea).Conn("bar", godip.Sea).Conn("now", godip.Sea).Conn("chr", godip.Coast...).Conn("swe", godip.Land).Flag(godip.Coast...).
		// Sweden
		Prov("swe").Conn("cop", godip.Coast...).Conn("bai", godip.Sea).Conn("gob", godip.Sea).Conn("fil", godip.Coast...).Conn("fim", godip.Land).Conn("chr", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(Denmark).
		// Croatia
		Prov("cro").Conn("adr", godip.Sea).Conn("bos", godip.Coast...).Conn("bud", godip.Land).Conn("vie", godip.Coast...).Flag(godip.Coast...).
		// Bohemia
		Prov("boh").Conn("gal", godip.Land).Conn("pol", godip.Land).Conn("brl", godip.Land).Conn("sax", godip.Land).Conn("bav", godip.Land).Conn("vie", godip.Land).Conn("bud", godip.Land).Flag(godip.Land).
		// Black Sea
		Prov("bla").Conn("waa", godip.Sea).Conn("con", godip.Sea).Conn("ang", godip.Sea).Conn("arm", godip.Sea).Conn("sev", godip.Sea).Conn("mol", godip.Sea).Flag(godip.Sea).
		// Egypt
		Prov("egy").Conn("syr", godip.Coast...).Conn("eas", godip.Sea).Conn("tri", godip.Coast...).Flag(godip.Coast...).SC(OttomanEmpire).
		// Barents Sea
		Prov("bar").Conn("now", godip.Sea).Conn("fim", godip.Sea).Conn("stp", godip.Sea).Conn("stp/nc", godip.Sea).Flag(godip.Sea).
		// Kiev
		Prov("kie").Conn("lio", godip.Land).Conn("pol", godip.Land).Conn("gal", godip.Land).Conn("ukr", godip.Land).Conn("mos", godip.Land).Flag(godip.Land).SC(Russia).
		// Helvetia
		Prov("hel").Conn("lyo", godip.Land).Conn("pie", godip.Land).Conn("cis", godip.Land).Conn("tyo", godip.Land).Conn("rhi", godip.Land).Conn("lor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lorraine
		Prov("lor").Conn("rhi", godip.Land).Conn("wet", godip.Land).Conn("bel", godip.Land).Conn("par", godip.Land).Conn("lyo", godip.Land).Conn("hel", godip.Land).Flag(godip.Land).
		// Apulia
		Prov("apu").Conn("pap", godip.Land).Conn("pap/ec", godip.Sea).Conn("nap", godip.Coast...).Conn("ion", godip.Sea).Conn("adr", godip.Sea).Flag(godip.Coast...).
		// Breslau
		Prov("brl").Conn("sax", godip.Land).Conn("boh", godip.Land).Conn("pol", godip.Land).Conn("kon", godip.Land).Conn("ber", godip.Land).Flag(godip.Land).SC(Prussia).
		// Edinburgh
		Prov("edi").Conn("not", godip.Sea).Conn("noa", godip.Sea).Conn("cel", godip.Sea).Conn("lie", godip.Coast...).Conn("ire", godip.Coast...).Conn("yor", godip.Coast...).Flag(godip.Coast...).SC(Britain).
		// Bay of Biscay
		Prov("bay").Conn("eng", godip.Sea).Conn("cel", godip.Sea).Conn("atl", godip.Sea).Conn("nav", godip.Sea).Conn("gas", godip.Sea).Conn("brt", godip.Sea).Flag(godip.Sea).
		// Naples
		Prov("nap").Conn("ion", godip.Sea).Conn("apu", godip.Coast...).Conn("pal", godip.Coast...).Conn("pap", godip.Land).Conn("pap/wc", godip.Sea).Conn("tys", godip.Sea).Flag(godip.Coast...).SC(Sicily).
		// Gulf of Lyon
		Prov("gol").Conn("cat", godip.Sea).Conn("bae", godip.Sea).Conn("tys", godip.Sea).Conn("pie", godip.Sea).Conn("mar", godip.Sea).Flag(godip.Sea).
		// Finland
		Prov("fil").Conn("stp", godip.Land).Conn("stp/sc", godip.Sea).Conn("fim", godip.Land).Conn("swe", godip.Coast...).Conn("gob", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cisalpine
		Prov("cis").Conn("ven", godip.Coast...).Conn("tyo", godip.Land).Conn("hel", godip.Land).Conn("pie", godip.Land).Conn("pap", godip.Land).Conn("pap/ec", godip.Sea).Conn("adr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Berlin
		Prov("ber").Conn("bai", godip.Sea).Conn("mec", godip.Coast...).Conn("bra", godip.Land).Conn("sax", godip.Land).Conn("brl", godip.Land).Conn("kon", godip.Coast...).Flag(godip.Coast...).SC(Prussia).
		// Livonia
		Prov("lio").Conn("pol", godip.Land).Conn("kie", godip.Land).Conn("mos", godip.Land).Conn("stp", godip.Land).Conn("stp/sc", godip.Sea).Conn("gob", godip.Sea).Conn("bai", godip.Sea).Conn("kon", godip.Coast...).Flag(godip.Coast...).
		// Navarre
		Prov("nav").Conn("mad", godip.Coast...).Conn("cat", godip.Land).Conn("gas", godip.Coast...).Conn("bay", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Coast...).
		// Lyon
		Prov("lyo").Conn("gas", godip.Land).Conn("mar", godip.Land).Conn("pie", godip.Land).Conn("hel", godip.Land).Conn("lor", godip.Land).Conn("par", godip.Land).Conn("brt", godip.Land).Flag(godip.Land).SC(France).
		// Tripolitania
		Prov("tri").Conn("tun", godip.Coast...).Conn("egy", godip.Coast...).Conn("eas", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Moscow
		Prov("mos").Conn("ukr", godip.Land).Conn("sev", godip.Land).Conn("ura", godip.Land).Conn("stp", godip.Land).Conn("lio", godip.Land).Conn("kie", godip.Land).Flag(godip.Land).SC(Russia).
		// Central Mediterranean
		Prov("cen").Conn("ion", godip.Sea).Conn("pal", godip.Sea).Conn("tys", godip.Sea).Conn("tun", godip.Sea).Conn("tri", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Batavia
		Prov("bat").Conn("wet", godip.Coast...).Conn("not", godip.Sea).Conn("bel", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}
