package coldwar

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	USSR godip.Nation = "USSR"
	NATO godip.Nation = "NATO"
)

var Nations = []godip.Nation{USSR, NATO}

var ColdWarVariant = common.Variant{
	Name:       "Cold War",
	Graph:      func() godip.Graph { return ColdWarGraph() },
	Start:      ColdWarStart,
	Blank:      ColdWarBlank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(17),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/coldwarmap.svg")
	},
	SVGVersion: "2",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Firehawk & Safari",
	Version:     "2",
	Description: "NATO and the USSR fight each other to see which will be the dominant superpower.",
	Rules: "Rules are as per classical Diplomacy, but with a different map. The winner " +
		"is the first to seventeen supply centers, which is slightly more than half. " +
		"Indonesia is connected to Australia and the Phillipines by bridges which " +
		"allow armies and fleets to travel between them. Panama, Egypt and Istanbul " +
		"contain canals, which allows fleets to enter and exit from either side. " +
		"Denmark and Sweden are single coast provinces which fleets and armies can " +
		"move between. Fleets in the North Sea must move to one of these provinces " +
		"to get to the Baltic. Fleets may only convoy if they are in all-sea provinces.",
}

func ColdWarBlank(phase godip.Phase) *state.State {
	return state.New(ColdWarGraph(), phase, classical.BackupRule)
}

func ColdWarStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1960, godip.Spring, godip.Movement)
	result = state.New(ColdWarGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"len/sc": godip.Unit{godip.Fleet, USSR},
		"alb":    godip.Unit{godip.Fleet, USSR},
		"hav":    godip.Unit{godip.Fleet, USSR},
		"mos":    godip.Unit{godip.Army, USSR},
		"sha":    godip.Unit{godip.Army, USSR},
		"vla":    godip.Unit{godip.Army, USSR},
		"lon":    godip.Unit{godip.Fleet, NATO},
		"ist":    godip.Unit{godip.Fleet, NATO},
		"aus":    godip.Unit{godip.Fleet, NATO},
		"nyk":    godip.Unit{godip.Army, NATO},
		"los":    godip.Unit{godip.Army, NATO},
		"par":    godip.Unit{godip.Army, NATO},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"len": USSR,
		"alb": USSR,
		"hav": USSR,
		"mos": USSR,
		"sha": USSR,
		"vla": USSR,
		"lon": NATO,
		"ist": NATO,
		"aus": NATO,
		"nyk": NATO,
		"los": NATO,
		"par": NATO,
	})
	return
}

func ColdWarGraph() *graph.Graph {
	return graph.New().
		// Tunisia
		Prov("tun").Conn("naf", godip.Coast...).Conn("lib", godip.Coast...).Conn("ion", godip.Sea).Conn("wme", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// North Vietnam
		Prov("nvi").Conn("sai", godip.Coast...).Conn("scs", godip.Sea).Conn("sha", godip.Coast...).Conn("sea", godip.Coast...).Flag(godip.Coast...).
		// Albania
		Prov("alb").Conn("ion", godip.Sea).Conn("grc", godip.Coast...).Conn("yug", godip.Coast...).Flag(godip.Coast...).SC(USSR).
		// Iran
		Prov("irn").Conn("arm", godip.Land).Conn("irq", godip.Coast...).Conn("arb", godip.Sea).Conn("pak", godip.Coast...).Conn("afg", godip.Land).Conn("ura", godip.Land).Conn("cau", godip.Land).Conn("cau", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Florida
		Prov("flo").Conn("wat", godip.Sea).Conn("nyk", godip.Coast...).Conn("mid", godip.Land).Conn("sow", godip.Coast...).Conn("gom", godip.Sea).Conn("car", godip.Sea).Flag(godip.Coast...).
		// London
		Prov("lon").Conn("nts", godip.Sea).Conn("nws", godip.Sea).Conn("nts", godip.Sea).Conn("eat", godip.Sea).Flag(godip.Coast...).SC(NATO).
		// Afghanistan
		Prov("afg").Conn("pak", godip.Land).Conn("sib", godip.Land).Conn("ura", godip.Land).Conn("irn", godip.Land).Flag(godip.Land).
		// Midwest
		Prov("mid").Conn("nyk", godip.Land).Conn("tor", godip.Land).Conn("wca", godip.Land).Conn("los", godip.Land).Conn("sow", godip.Land).Conn("flo", godip.Land).Flag(godip.Land).
		// Levant
		Prov("lev").Conn("eme", godip.Sea).Conn("egy", godip.Coast...).Conn("ara", godip.Land).Conn("irq", godip.Land).Conn("arm", godip.Land).Conn("ist", godip.Coast...).Flag(godip.Coast...).
		// North Korea
		Prov("nko").Conn("seo", godip.Land).Conn("vla", godip.Land).Conn("man", godip.Land).Flag(godip.Land).
		// North Korea (East Coast)
		Prov("nko/ec").Conn("seo", godip.Sea).Conn("soj", godip.Sea).Conn("vla", godip.Sea).Flag(godip.Sea).
		// North Korea (West Coast)
		Prov("nko/wc").Conn("yel", godip.Sea).Conn("seo", godip.Sea).Conn("man", godip.Sea).Flag(godip.Sea).
		// India
		Prov("ind").Conn("ban", godip.Coast...).Conn("pak", godip.Coast...).Conn("arb", godip.Sea).Conn("inc", godip.Sea).Conn("bay", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// New York
		Prov("nyk").Conn("que", godip.Coast...).Conn("tor", godip.Land).Conn("mid", godip.Land).Conn("flo", godip.Coast...).Conn("wat", godip.Sea).Flag(godip.Coast...).SC(NATO).
		// Venezuela
		Prov("ven").Conn("col", godip.Land).Conn("col/nc", godip.Sea).Conn("bra", godip.Coast...).Conn("wat", godip.Sea).Conn("car", godip.Sea).Flag(godip.Coast...).
		// Caribbean Sea
		Prov("car").Conn("hav", godip.Sea).Conn("gom", godip.Sea).Conn("mex", godip.Sea).Conn("mex/ec", godip.Sea).Conn("cen", godip.Sea).Conn("cen/ec", godip.Sea).Conn("pan", godip.Sea).Conn("col", godip.Sea).Conn("col/nc", godip.Sea).Conn("ven", godip.Sea).Conn("wat", godip.Sea).Conn("flo", godip.Sea).Flag(godip.Sea).
		// Greenland
		Prov("grd").Conn("arc", godip.Sea).Conn("wat", godip.Sea).Conn("nws", godip.Sea).Flag(godip.Coast...).
		// Paris
		Prov("par").Conn("ita", godip.Land).Conn("wge", godip.Land).Conn("spa", godip.Land).Flag(godip.Land).SC(NATO).
		// Paris (North Coast)
		Prov("par/nc").Conn("wge", godip.Sea).Conn("nts", godip.Sea).Conn("eat", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// Paris (South Coast)
		Prov("par/sc").Conn("wme", godip.Sea).Conn("ita", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// Ionian Sea
		Prov("ion").Conn("grc", godip.Sea).Conn("alb", godip.Sea).Conn("yug", godip.Sea).Conn("ita", godip.Sea).Conn("wme", godip.Sea).Conn("tun", godip.Sea).Conn("lib", godip.Sea).Conn("eme", godip.Sea).Flag(godip.Sea).
		// Brazil
		Prov("bra").Conn("wat", godip.Sea).Conn("ven", godip.Coast...).Conn("col", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Mexico
		Prov("gom").Conn("mex", godip.Sea).Conn("mex/ec", godip.Sea).Conn("car", godip.Sea).Conn("hav", godip.Sea).Conn("car", godip.Sea).Conn("flo", godip.Sea).Conn("sow", godip.Sea).Flag(godip.Sea).
		// West Atlantic
		Prov("wat").Conn("eat", godip.Sea).Conn("nws", godip.Sea).Conn("grd", godip.Sea).Conn("arc", godip.Sea).Conn("hud", godip.Sea).Conn("que", godip.Sea).Conn("nyk", godip.Sea).Conn("flo", godip.Sea).Conn("car", godip.Sea).Conn("ven", godip.Sea).Conn("bra", godip.Sea).Flag(godip.Sea).
		// West China
		Prov("wch").Conn("mon", godip.Land).Conn("sib", godip.Land).Conn("pak", godip.Land).Conn("ban", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).
		// Havana
		Prov("hav").Conn("car", godip.Sea).Conn("gom", godip.Sea).Flag(godip.Coast...).SC(USSR).
		// Arabia
		Prov("ara").Conn("egy", godip.Coast...).Conn("red", godip.Sea).Conn("arb", godip.Sea).Conn("irq", godip.Coast...).Conn("lev", godip.Land).Flag(godip.Coast...).
		// East Germany
		Prov("ege").Conn("wge", godip.Land).Conn("yug", godip.Land).Conn("ukr", godip.Land).Conn("mos", godip.Land).Conn("len", godip.Land).Conn("len/sc", godip.Sea).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Leningrad
		Prov("len").Conn("noy", godip.Land).Conn("fin", godip.Land).Conn("ege", godip.Land).Conn("mos", godip.Land).Conn("ura", godip.Land).Flag(godip.Land).SC(USSR).
		// Leningrad (North Coast)
		Prov("len/nc").Conn("noy", godip.Sea).Conn("ura", godip.Sea).Conn("nws", godip.Sea).Flag(godip.Sea).
		// Leningrad (South Coast)
		Prov("len/sc").Conn("fin", godip.Sea).Conn("bal", godip.Sea).Conn("ege", godip.Sea).Flag(godip.Sea).
		// North Africa
		Prov("naf").Conn("lib", godip.Land).Conn("tun", godip.Coast...).Conn("wme", godip.Sea).Conn("eat", godip.Sea).Flag(godip.Coast...).
		// Baltic Sea
		Prov("bal").Conn("fin", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Conn("ege", godip.Sea).Conn("len", godip.Sea).Conn("len/sc", godip.Sea).Flag(godip.Sea).
		// Yugoslavia
		Prov("yug").Conn("wge", godip.Land).Conn("ita", godip.Coast...).Conn("ion", godip.Sea).Conn("alb", godip.Coast...).Conn("grc", godip.Coast...).Conn("ukr", godip.Land).Conn("ege", godip.Land).Flag(godip.Coast...).
		// Toronto
		Prov("tor").Conn("nyk", godip.Land).Conn("que", godip.Coast...).Conn("hud", godip.Sea).Conn("wca", godip.Land).Conn("wca/nc", godip.Sea).Conn("mid", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Norway
		Prov("noy").Conn("nws", godip.Sea).Conn("nts", godip.Sea).Conn("swe", godip.Coast...).Conn("fin", godip.Coast...).Conn("len", godip.Land).Conn("len/nc", godip.Sea).Flag(godip.Coast...).
		// Vladivostok
		Prov("vla").Conn("man", godip.Land).Conn("nko", godip.Land).Conn("nko/ec", godip.Sea).Conn("soj", godip.Sea).Conn("ber", godip.Sea).Conn("kam", godip.Coast...).Conn("sib", godip.Land).Flag(godip.Coast...).SC(USSR).
		// East Africa
		Prov("eaf").Conn("inc", godip.Sea).Conn("red", godip.Sea).Conn("egy", godip.Coast...).Conn("lib", godip.Land).Flag(godip.Coast...).
		// Libya
		Prov("lib").Conn("eaf", godip.Land).Conn("egy", godip.Coast...).Conn("eme", godip.Sea).Conn("ion", godip.Sea).Conn("tun", godip.Coast...).Conn("naf", godip.Land).Flag(godip.Coast...).
		// Japan
		Prov("jap").Conn("soj", godip.Sea).Conn("yel", godip.Sea).Conn("wpa", godip.Sea).Conn("ber", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Denmark
		Prov("den").Conn("bal", godip.Sea).Conn("nts", godip.Sea).Conn("wge", godip.Coast...).Conn("ege", godip.Coast...).Conn("swe", godip.Coast...).Flag(godip.Coast...).
		// Seoul
		Prov("seo").Conn("nko", godip.Land).Conn("nko/wc", godip.Sea).Conn("nko/ec", godip.Sea).Conn("yel", godip.Sea).Conn("soj", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Bering Sea
		Prov("ber").Conn("epa", godip.Sea).Conn("goa", godip.Sea).Conn("ala", godip.Sea).Conn("arc", godip.Sea).Conn("kam", godip.Sea).Conn("vla", godip.Sea).Conn("soj", godip.Sea).Conn("jap", godip.Sea).Conn("wpa", godip.Sea).Flag(godip.Sea).
		// Los Angeles
		Prov("los").Conn("wca", godip.Land).Conn("wca/wc", godip.Sea).Conn("goa", godip.Sea).Conn("epa", godip.Sea).Conn("mex", godip.Land).Conn("mex/wc", godip.Sea).Conn("sow", godip.Land).Conn("mid", godip.Land).Flag(godip.Coast...).SC(NATO).
		// Caucasus
		Prov("cau").Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Conn("irn", godip.Land).Conn("irn", godip.Land).Conn("ura", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Coast...).Flag(godip.Coast...).
		// Armenia
		Prov("arm").Conn("irq", godip.Land).Conn("irn", godip.Land).Conn("cau", godip.Coast...).Conn("bla", godip.Sea).Conn("ist", godip.Coast...).Conn("lev", godip.Land).Flag(godip.Coast...).
		// Panama
		Prov("pan").Conn("col", godip.Land).Conn("col/nc", godip.Sea).Conn("col/wc", godip.Sea).Conn("car", godip.Sea).Conn("cen", godip.Land).Conn("cen/ec", godip.Sea).Conn("cen/wc", godip.Sea).Conn("epa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Southwest
		Prov("sow").Conn("mid", godip.Land).Conn("los", godip.Land).Conn("mex", godip.Land).Conn("mex/ec", godip.Sea).Conn("gom", godip.Sea).Conn("flo", godip.Coast...).Flag(godip.Coast...).
		// South China Sea
		Prov("scs").Conn("sai", godip.Sea).Conn("sea", godip.Sea).Conn("bay", godip.Sea).Conn("ins", godip.Sea).Conn("phi", godip.Sea).Conn("yel", godip.Sea).Conn("sha", godip.Sea).Conn("nvi", godip.Sea).Flag(godip.Sea).
		// Istanbul
		Prov("ist").Conn("grc", godip.Coast...).Conn("eme", godip.Sea).Conn("lev", godip.Coast...).Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Conn("ukr", godip.Coast...).Flag(godip.Coast...).SC(NATO).
		// Arabian Sea
		Prov("arb").Conn("irq", godip.Sea).Conn("ara", godip.Sea).Conn("red", godip.Sea).Conn("inc", godip.Sea).Conn("ind", godip.Sea).Conn("pak", godip.Sea).Conn("irn", godip.Sea).Flag(godip.Sea).
		// Finland
		Prov("fin").Conn("bal", godip.Sea).Conn("len", godip.Land).Conn("len/sc", godip.Sea).Conn("noy", godip.Coast...).Conn("swe", godip.Coast...).Flag(godip.Coast...).
		// East Mediterranean
		Prov("eme").Conn("lev", godip.Sea).Conn("ist", godip.Sea).Conn("grc", godip.Sea).Conn("ion", godip.Sea).Conn("lib", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Sea).
		// North Sea
		Prov("nts").Conn("swe", godip.Sea).Conn("noy", godip.Sea).Conn("nws", godip.Sea).Conn("lon", godip.Sea).Conn("eat", godip.Sea).Conn("par", godip.Sea).Conn("par/nc", godip.Sea).Conn("wge", godip.Sea).Conn("den", godip.Sea).Flag(godip.Sea).
		// Urals
		Prov("ura").Conn("nws", godip.Sea).Conn("len", godip.Land).Conn("len/nc", godip.Sea).Conn("mos", godip.Land).Conn("cau", godip.Land).Conn("irn", godip.Land).Conn("afg", godip.Land).Conn("sib", godip.Coast...).Conn("arc", godip.Sea).Flag(godip.Coast...).
		// Manchuria
		Prov("man").Conn("vla", godip.Land).Conn("sib", godip.Land).Conn("mon", godip.Land).Conn("sha", godip.Coast...).Conn("yel", godip.Sea).Conn("nko", godip.Land).Conn("nko/wc", godip.Sea).Flag(godip.Coast...).
		// East Atlantic
		Prov("eat").Conn("naf", godip.Sea).Conn("wme", godip.Sea).Conn("spa", godip.Sea).Conn("par", godip.Sea).Conn("par/nc", godip.Sea).Conn("nts", godip.Sea).Conn("nws", godip.Sea).Conn("wat", godip.Sea).Conn("lon", godip.Sea).Flag(godip.Sea).
		// Alaska
		Prov("ala").Conn("arc", godip.Sea).Conn("ber", godip.Sea).Conn("goa", godip.Sea).Conn("wca", godip.Land).Conn("wca/nc", godip.Sea).Conn("wca/wc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Bay of Bengal
		Prov("bay").Conn("ins", godip.Sea).Conn("scs", godip.Sea).Conn("sea", godip.Sea).Conn("ban", godip.Sea).Conn("ind", godip.Sea).Conn("inc", godip.Sea).Conn("inc", godip.Sea).Flag(godip.Sea).
		// Ukraine
		Prov("ukr").Conn("cau", godip.Coast...).Conn("mos", godip.Land).Conn("ege", godip.Land).Conn("yug", godip.Land).Conn("grc", godip.Land).Conn("ist", godip.Coast...).Conn("bla", godip.Sea).Flag(godip.Coast...).
		// Saigon
		Prov("sai").Conn("scs", godip.Sea).Conn("nvi", godip.Coast...).Conn("sea", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Bangladesh
		Prov("ban").Conn("sha", godip.Land).Conn("wch", godip.Land).Conn("ind", godip.Coast...).Conn("bay", godip.Sea).Conn("sea", godip.Coast...).Flag(godip.Coast...).
		// Sea of Japan
		Prov("soj").Conn("nko", godip.Sea).Conn("nko/ec", godip.Sea).Conn("seo", godip.Sea).Conn("yel", godip.Sea).Conn("jap", godip.Sea).Conn("ber", godip.Sea).Conn("vla", godip.Sea).Flag(godip.Sea).
		// East Pacific
		Prov("epa").Conn("col", godip.Sea).Conn("col/wc", godip.Sea).Conn("pan", godip.Sea).Conn("cen", godip.Sea).Conn("cen/wc", godip.Sea).Conn("mex", godip.Sea).Conn("mex/wc", godip.Sea).Conn("los", godip.Sea).Conn("goa", godip.Sea).Conn("ber", godip.Sea).Conn("wpa", godip.Sea).Flag(godip.Sea).
		// Spain
		Prov("spa").Conn("wme", godip.Sea).Conn("par", godip.Land).Conn("par/nc", godip.Sea).Conn("par/sc", godip.Sea).Conn("eat", godip.Sea).Flag(godip.Coast...).
		// Indian Ocean
		Prov("inc").Conn("aus", godip.Sea).Conn("ins", godip.Sea).Conn("bay", godip.Sea).Conn("bay", godip.Sea).Conn("ind", godip.Sea).Conn("arb", godip.Sea).Conn("red", godip.Sea).Conn("eaf", godip.Sea).Flag(godip.Sea).
		// Norwegian Sea
		Prov("nws").Conn("ura", godip.Sea).Conn("arc", godip.Sea).Conn("grd", godip.Sea).Conn("wat", godip.Sea).Conn("eat", godip.Sea).Conn("nts", godip.Sea).Conn("lon", godip.Sea).Conn("nts", godip.Sea).Conn("noy", godip.Sea).Conn("len", godip.Sea).Conn("len/nc", godip.Sea).Flag(godip.Sea).
		// Hudson Bay
		Prov("hud").Conn("arc", godip.Sea).Conn("wca", godip.Sea).Conn("wca/nc", godip.Sea).Conn("tor", godip.Sea).Conn("que", godip.Sea).Conn("wat", godip.Sea).Flag(godip.Sea).
		// Philippines
		Prov("phi").Conn("yel", godip.Sea).Conn("scs", godip.Sea).Conn("ins", godip.Coast...).Conn("wpa", godip.Sea).Flag(godip.Coast...).
		// Mongolia
		Prov("mon").Conn("wch", godip.Land).Conn("sha", godip.Land).Conn("man", godip.Land).Conn("sib", godip.Land).Flag(godip.Land).
		// Yellow Sea
		Prov("yel").Conn("wpa", godip.Sea).Conn("jap", godip.Sea).Conn("soj", godip.Sea).Conn("seo", godip.Sea).Conn("nko", godip.Sea).Conn("nko/wc", godip.Sea).Conn("man", godip.Sea).Conn("sha", godip.Sea).Conn("scs", godip.Sea).Conn("phi", godip.Sea).Flag(godip.Sea).
		// West Germany
		Prov("wge").Conn("ege", godip.Land).Conn("den", godip.Coast...).Conn("nts", godip.Sea).Conn("par", godip.Land).Conn("par/nc", godip.Sea).Conn("yug", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Greece
		Prov("grc").Conn("ion", godip.Sea).Conn("eme", godip.Sea).Conn("ist", godip.Coast...).Conn("ukr", godip.Land).Conn("yug", godip.Coast...).Conn("alb", godip.Coast...).Flag(godip.Coast...).
		// Arctic Ocean
		Prov("arc").Conn("grd", godip.Sea).Conn("nws", godip.Sea).Conn("ura", godip.Sea).Conn("sib", godip.Sea).Conn("kam", godip.Sea).Conn("ber", godip.Sea).Conn("ala", godip.Sea).Conn("wca", godip.Sea).Conn("wca/nc", godip.Sea).Conn("hud", godip.Sea).Conn("wat", godip.Sea).Flag(godip.Sea).
		// Sweden
		Prov("swe").Conn("bal", godip.Sea).Conn("nts", godip.Sea).Conn("fin", godip.Coast...).Conn("noy", godip.Coast...).Conn("den", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Iraq
		Prov("irq").Conn("arb", godip.Sea).Conn("irn", godip.Coast...).Conn("arm", godip.Land).Conn("lev", godip.Land).Conn("ara", godip.Coast...).Flag(godip.Coast...).
		// Pakistan
		Prov("pak").Conn("arb", godip.Sea).Conn("ind", godip.Coast...).Conn("wch", godip.Land).Conn("sib", godip.Land).Conn("afg", godip.Land).Conn("irn", godip.Coast...).Flag(godip.Coast...).
		// Shanghai
		Prov("sha").Conn("ban", godip.Land).Conn("nvi", godip.Coast...).Conn("scs", godip.Sea).Conn("yel", godip.Sea).Conn("man", godip.Coast...).Conn("mon", godip.Land).Conn("wch", godip.Land).Conn("sea", godip.Land).Flag(godip.Coast...).SC(USSR).
		// Mexico
		Prov("mex").Conn("sow", godip.Land).Conn("los", godip.Land).Conn("cen", godip.Land).Flag(godip.Land).
		// Mexico (East Coast)
		Prov("mex/ec").Conn("gom", godip.Sea).Conn("sow", godip.Sea).Conn("cen/ec", godip.Sea).Conn("car", godip.Sea).Flag(godip.Sea).
		// Mexico (West Coast)
		Prov("mex/wc").Conn("los", godip.Sea).Conn("epa", godip.Sea).Conn("cen/wc", godip.Sea).Flag(godip.Sea).
		// West Canada
		Prov("wca").Conn("los", godip.Land).Conn("mid", godip.Land).Conn("tor", godip.Land).Conn("ala", godip.Land).Flag(godip.Land).
		// West Canada (North Coast)
		Prov("wca/nc").Conn("tor", godip.Sea).Conn("hud", godip.Sea).Conn("arc", godip.Sea).Conn("ala", godip.Sea).Flag(godip.Sea).
		// West Canada (West Coast)
		Prov("wca/wc").Conn("los", godip.Sea).Conn("ala", godip.Sea).Conn("goa", godip.Sea).Flag(godip.Sea).
		// West Pacific
		Prov("wpa").Conn("epa", godip.Sea).Conn("ber", godip.Sea).Conn("jap", godip.Sea).Conn("yel", godip.Sea).Conn("phi", godip.Sea).Conn("ins", godip.Sea).Conn("ins", godip.Sea).Conn("aus", godip.Sea).Flag(godip.Sea).
		// Black Sea
		Prov("bla").Conn("cau", godip.Sea).Conn("ukr", godip.Sea).Conn("ist", godip.Sea).Conn("arm", godip.Sea).Flag(godip.Sea).
		// Egypt
		Prov("egy").Conn("red", godip.Sea).Conn("ara", godip.Coast...).Conn("lev", godip.Coast...).Conn("eme", godip.Sea).Conn("lib", godip.Coast...).Conn("eaf", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Central America
		Prov("cen").Conn("pan", godip.Land).Conn("mex", godip.Land).Flag(godip.Land).
		// Central America (East Coast)
		Prov("cen/ec").Conn("pan", godip.Sea).Conn("car", godip.Sea).Conn("mex/ec", godip.Sea).Flag(godip.Sea).
		// Central America (West Coast)
		Prov("cen/wc").Conn("epa", godip.Sea).Conn("pan", godip.Sea).Conn("mex/wc", godip.Sea).Flag(godip.Sea).
		// Red Sea
		Prov("red").Conn("egy", godip.Sea).Conn("eaf", godip.Sea).Conn("inc", godip.Sea).Conn("arb", godip.Sea).Conn("ara", godip.Sea).Flag(godip.Sea).
		// Australia
		Prov("aus").Conn("wpa", godip.Sea).Conn("ins", godip.Coast...).Conn("inc", godip.Sea).Flag(godip.Coast...).SC(NATO).
		// Siberia
		Prov("sib").Conn("pak", godip.Land).Conn("wch", godip.Land).Conn("mon", godip.Land).Conn("man", godip.Land).Conn("vla", godip.Land).Conn("kam", godip.Coast...).Conn("arc", godip.Sea).Conn("ura", godip.Coast...).Conn("afg", godip.Land).Flag(godip.Coast...).
		// Kamchatka
		Prov("kam").Conn("arc", godip.Sea).Conn("sib", godip.Coast...).Conn("vla", godip.Coast...).Conn("ber", godip.Sea).Flag(godip.Coast...).
		// Indonesia
		Prov("ins").Conn("aus", godip.Coast...).Conn("wpa", godip.Sea).Conn("wpa", godip.Sea).Conn("phi", godip.Coast...).Conn("scs", godip.Sea).Conn("bay", godip.Sea).Conn("inc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// West Mediterranean
		Prov("wme").Conn("spa", godip.Sea).Conn("eat", godip.Sea).Conn("naf", godip.Sea).Conn("tun", godip.Sea).Conn("ion", godip.Sea).Conn("ita", godip.Sea).Conn("par", godip.Sea).Conn("par/sc", godip.Sea).Flag(godip.Sea).
		// Colombia
		Prov("col").Conn("bra", godip.Land).Conn("ven", godip.Land).Conn("pan", godip.Land).Flag(godip.Land).
		// Colombia
		Prov("col/nc").Conn("ven", godip.Sea).Conn("car", godip.Sea).Conn("pan", godip.Sea).Flag(godip.Sea).
		// Colombia
		Prov("col/wc").Conn("pan", godip.Sea).Conn("epa", godip.Sea).Flag(godip.Sea).
		// Quebec
		Prov("que").Conn("nyk", godip.Coast...).Conn("wat", godip.Sea).Conn("hud", godip.Sea).Conn("tor", godip.Coast...).Flag(godip.Coast...).
		// South East Asia
		Prov("sea").Conn("sai", godip.Coast...).Conn("nvi", godip.Coast...).Conn("ban", godip.Coast...).Conn("bay", godip.Sea).Conn("scs", godip.Sea).Conn("sha", godip.Land).Flag(godip.Coast...).
		// Italy
		Prov("ita").Conn("par", godip.Land).Conn("par/sc", godip.Sea).Conn("wme", godip.Sea).Conn("ion", godip.Sea).Conn("yug", godip.Coast...).Flag(godip.Coast...).
		// Moscow
		Prov("mos").Conn("cau", godip.Land).Conn("ura", godip.Land).Conn("len", godip.Land).Conn("ege", godip.Land).Conn("ukr", godip.Land).Flag(godip.Land).SC(USSR).
		// Gulf of Alaska
		Prov("goa").Conn("epa", godip.Sea).Conn("los", godip.Sea).Conn("wca", godip.Sea).Conn("wca/wc", godip.Sea).Conn("ala", godip.Sea).Conn("ber", godip.Sea).Flag(godip.Sea).
		Done()
}
