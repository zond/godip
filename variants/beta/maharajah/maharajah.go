package maharajah

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"	
)

const (
	Gondwana    godip.Nation = "Gondwana"
	Mughalistan godip.Nation = "Mughalistan"
	Persia      godip.Nation = "Persia"
	Rajputana   godip.Nation = "Rajputana"
	Delhi       godip.Nation = "Delhi"
	Vijayanagar godip.Nation = "Vijayanagar"
	Bahmana     godip.Nation = "Bahmana"
)

var Nations = []godip.Nation{Gondwana, Mughalistan, Persia, Rajputana, Delhi, Vijayanagar, Bahmana}

var newPhase = phase.Generator(hundred.BuildAnywhereParser, classical.AdjustSCs)

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return newPhase(year, season, typ)
}

var MaharajahVariant = common.Variant{
	Name:              "Maharajah",
	Graph:             func() godip.Graph { return MaharajahGraph() },
	Start:             MaharajahStart,
	Blank:             MaharajahBlank,
	Phase:             Phase,
	Parser:            hundred.BuildAnywhereParser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(19),
	SoloSCCount:       func(*state.State) int { return 19 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/maharajahmap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "David E. Cohen",
	Version:     "2.0",
	Description: "THIS IS A BETA MAP. IT WILL ONLY BE VISIBLE TO PEOPLE USING BETA.DIPLICITY.COM OR THE BETA ANDROID APP. This map is, together with the Spice Islands map, part of the larger East Indies Diplomacy variant. Fight for dominance of the Indus region.",
	Rules:       `The first to 25 Supply Centers (SC) is the winner. 
Units may be built at any owned SC.
Provinces adjacent to a river are coastal and can be occupied by fleets. These can move to other provinces connected by the same river.`,
}

func MaharajahBlank(phase godip.Phase) *state.State {
	return state.New(MaharajahGraph(), phase, classical.BackupRule, nil, nil)
}

func MaharajahStart() (result *state.State, err error) {
	startPhase := Phase(1501, godip.Spring, godip.Movement)
	result = MaharajahBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"sab": godip.Unit{godip.Fleet, Gondwana},
		"jab": godip.Unit{godip.Army, Gondwana},
		"rai": godip.Unit{godip.Army, Gondwana},
		"kab": godip.Unit{godip.Army, Mughalistan},
		"bad": godip.Unit{godip.Army, Mughalistan},
		"bal": godip.Unit{godip.Army, Mughalistan},
		"hor": godip.Unit{godip.Fleet, Persia},
		"mes": godip.Unit{godip.Army, Persia},
		"isf": godip.Unit{godip.Army, Persia},
		"jas": godip.Unit{godip.Fleet, Rajputana},
		"mul": godip.Unit{godip.Army, Rajputana},
		"jod": godip.Unit{godip.Army, Rajputana},
		"awa": godip.Unit{godip.Army, Delhi},
		"agr": godip.Unit{godip.Army, Delhi},
		"muz": godip.Unit{godip.Army, Delhi},
		"pul": godip.Unit{godip.Fleet, Vijayanagar},
		"ban": godip.Unit{godip.Army, Vijayanagar},
		"cal": godip.Unit{godip.Army, Vijayanagar},
		"goa": godip.Unit{godip.Fleet, Bahmana},
		"ahn": godip.Unit{godip.Army, Bahmana},
		"bij": godip.Unit{godip.Army, Bahmana},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"sab": Gondwana,
		"jab": Gondwana,
		"rai": Gondwana,
		"kab": Mughalistan,
		"bad": Mughalistan,
		"bal": Mughalistan,
		"hor": Persia,
		"mes": Persia,
		"isf": Persia,
		"jas": Rajputana,
		"mul": Rajputana,
		"jod": Rajputana,
		"awa": Delhi,
		"agr": Delhi,
		"muz": Delhi,
		"pul": Vijayanagar,
		"ban": Vijayanagar,
		"cal": Vijayanagar,
		"goa": Bahmana,
		"ahn": Bahmana,
		"bij": Bahmana,
	})
	return
}

func MaharajahGraph() *graph.Graph {
	return graph.New().
		// Sambalpur
		Prov("sab").Conn("ori", godip.Coast...).Conn("beg", godip.Coast...).Conn("bee", godip.Coast...).Conn("jab", godip.Land).Conn("rai", godip.Land).Conn("war", godip.Land).Flag(godip.Coast...).SC(Gondwana).
		// Kara Kum
		Prov("kar").Conn("elb", godip.Land).Conn("mes", godip.Land).Conn("buk", godip.Land).Flag(godip.Land).
		// Mewar
		Prov("mew").Conn("maw", godip.Land).Conn("jap", godip.Land).Conn("bik", godip.Land).Conn("jod", godip.Land).Conn("ahb", godip.Land).Flag(godip.Land).
		// Tinnevelly
		Prov("tin").Conn("cal", godip.Coast...).Conn("lac", godip.Sea).Conn("nic", godip.Sea).Conn("pul", godip.Coast...).Flag(godip.Coast...).
		// Lahore
		Prov("lah").Conn("jap", godip.Land).Conn("agr", godip.Land).Conn("kam", godip.Land).Conn("kab", godip.Coast...).Conn("pes", godip.Coast...).Conn("mul", godip.Coast...).Flag(godip.Coast...).
		// Bengal
		Prov("beg").Conn("bee", godip.Coast...).Conn("sab", godip.Coast...).Conn("ori", godip.Coast...).Conn("bay", godip.Sea).Conn("peg", godip.Coast...).Conn("ass", godip.Coast...).Conn("muz", godip.Coast...).Flag(godip.Coast...).
		// Aceh
		Prov("ace").Conn("and", godip.Sea).Conn("nic", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Jaisalmer
		Prov("jas").Conn("sin", godip.Coast...).Conn("guj", godip.Land).Conn("jod", godip.Land).Conn("bik", godip.Coast...).Conn("pes", godip.Coast...).Flag(godip.Coast...).SC(Rajputana).
		// Warangal
		Prov("war").Conn("bid", godip.Land).Conn("ban", godip.Land).Conn("ori", godip.Land).Conn("sab", godip.Land).Conn("rai", godip.Land).Flag(godip.Land).
		// Nepal
		Prov("nep").Conn("muz", godip.Land).Conn("ava", godip.Land).Conn("gar", godip.Land).Conn("awa", godip.Land).Flag(godip.Land).
		// Muzaffarpur
		Prov("muz").Conn("awa", godip.Coast...).Conn("agr", godip.Coast...).Conn("bee", godip.Coast...).Conn("beg", godip.Coast...).Conn("ass", godip.Coast...).Conn("ava", godip.Land).Conn("nep", godip.Land).Flag(godip.Coast...).SC(Delhi).
		// Kashmir
		Prov("kam").Conn("kab", godip.Land).Conn("lah", godip.Land).Conn("agr", godip.Land).Conn("awa", godip.Land).Conn("gar", godip.Land).Conn("kag", godip.Land).Conn("fer", godip.Land).Conn("bad", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Berar
		Prov("ber").Conn("bee", godip.Land).Conn("maw", godip.Land).Conn("bid", godip.Land).Conn("rai", godip.Land).Conn("jab", godip.Land).Flag(godip.Land).
		// Balkh
		Prov("bal").Conn("sad", godip.Land).Conn("buk", godip.Land).Conn("her", godip.Land).Conn("kna", godip.Land).Conn("kab", godip.Land).Conn("bad", godip.Land).Flag(godip.Land).SC(Mughalistan).
		// Kra
		Prov("kra").Conn("ayu", godip.Coast...).Conn("and", godip.Sea).Flag(godip.Coast...).
		// Dasht-I-Kavir
		Prov("das").Conn("elb", godip.Land).Conn("qom", godip.Land).Conn("isf", godip.Land).Conn("yez", godip.Land).Conn("mes", godip.Land).Flag(godip.Land).
		// Kandahar
		Prov("kna").Conn("bal", godip.Land).Conn("her", godip.Land).Conn("yez", godip.Land).Conn("shi", godip.Land).Conn("pes", godip.Land).Conn("kab", godip.Land).Flag(godip.Land).
		// Bukhara
		Prov("buk").Conn("kar", godip.Land).Conn("mes", godip.Land).Conn("her", godip.Land).Conn("bal", godip.Land).Conn("sad", godip.Land).Flag(godip.Land).
		// Samarkand
		Prov("sad").Conn("buk", godip.Land).Conn("bal", godip.Land).Conn("bad", godip.Land).Conn("fer", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Yemen
		Prov("yem").Conn("gul", godip.Sea).Conn("had", godip.Coast...).Conn("rub", godip.Land).Conn("arb", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Peshawar
		Prov("pes").Conn("lah", godip.Coast...).Conn("kab", godip.Coast...).Conn("kna", godip.Land).Conn("shi", godip.Land).Conn("sin", godip.Coast...).Conn("jas", godip.Coast...).Conn("bik", godip.Coast...).Conn("mul", godip.Coast...).Flag(godip.Coast...).
		// Awadh
		Prov("awa").Conn("muz", godip.Coast...).Conn("nep", godip.Land).Conn("gar", godip.Land).Conn("kam", godip.Land).Conn("agr", godip.Coast...).Flag(godip.Coast...).SC(Delhi).
		// PERSIAN GULF
		Prov("per").Conn("arb", godip.Sea).Conn("oma", godip.Sea).Conn("ars", godip.Sea).Conn("hor", godip.Sea).Conn("isf", godip.Sea).Conn("qom", godip.Sea).Flag(godip.Sea).
		// Kashgar
		Prov("kag").Conn("fer", godip.Land).Conn("kam", godip.Land).Conn("gar", godip.Land).Flag(godip.Land).
		// Ayutthaya
		Prov("ayu").Conn("ava", godip.Land).Conn("peg", godip.Coast...).Conn("bay", godip.Sea).Conn("and", godip.Sea).Conn("kra", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shiraz
		Prov("shi").Conn("hor", godip.Land).Conn("sin", godip.Land).Conn("pes", godip.Land).Conn("kna", godip.Land).Conn("yez", godip.Land).Conn("isf", godip.Land).Flag(godip.Land).
		// Seylac
		Prov("sey").Conn("mas", godip.Sea).Conn("lac", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Isfahan
		Prov("isf").Conn("qom", godip.Coast...).Conn("per", godip.Sea).Conn("hor", godip.Coast...).Conn("shi", godip.Land).Conn("yez", godip.Land).Conn("das", godip.Land).Flag(godip.Coast...).SC(Persia).
		// Rub Al Khali
		Prov("rub").Conn("had", godip.Land).Conn("oma", godip.Land).Conn("arb", godip.Land).Conn("yem", godip.Land).Flag(godip.Land).
		// Kabul
		Prov("kab").Conn("kam", godip.Land).Conn("bad", godip.Land).Conn("bal", godip.Land).Conn("kna", godip.Land).Conn("pes", godip.Coast...).Conn("lah", godip.Coast...).Flag(godip.Coast...).SC(Mughalistan).
		// Pulicat
		Prov("pul").Conn("ban", godip.Coast...).Conn("cal", godip.Land).Conn("tin", godip.Coast...).Conn("nic", godip.Sea).Conn("bay", godip.Sea).Flag(godip.Coast...).SC(Vijayanagar).
		// Jaffna
		Prov("jaf").Conn("kay", godip.Coast...).Conn("nic", godip.Sea).Conn("lac", godip.Sea).Flag(godip.Coast...).
		// ANDAMAN SEA
		Prov("and").Conn("kra", godip.Sea).Conn("ayu", godip.Sea).Conn("bay", godip.Sea).Conn("nic", godip.Sea).Conn("ace", godip.Sea).Flag(godip.Sea).
		// Jaipur
		Prov("jap").Conn("lah", godip.Land).Conn("mul", godip.Land).Conn("bik", godip.Land).Conn("mew", godip.Land).Conn("maw", godip.Land).Conn("bee", godip.Land).Conn("agr", godip.Land).Flag(godip.Land).
		// Gartok
		Prov("gar").Conn("kag", godip.Land).Conn("kam", godip.Land).Conn("awa", godip.Land).Conn("nep", godip.Land).Flag(godip.Land).
		// Jabalpur
		Prov("jab").Conn("rai", godip.Land).Conn("sab", godip.Land).Conn("bee", godip.Land).Conn("ber", godip.Land).Flag(godip.Land).SC(Gondwana).
		// Raipur
		Prov("rai").Conn("ber", godip.Land).Conn("bid", godip.Land).Conn("war", godip.Land).Conn("sab", godip.Land).Conn("jab", godip.Land).Flag(godip.Land).SC(Gondwana).
		// Bidar
		Prov("bid").Conn("war", godip.Land).Conn("rai", godip.Land).Conn("ber", godip.Land).Conn("maw", godip.Land).Conn("ahb", godip.Land).Conn("kha", godip.Land).Conn("bij", godip.Land).Conn("hon", godip.Land).Conn("ban", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bijapur
		Prov("bij").Conn("goa", godip.Land).Conn("hon", godip.Land).Conn("bid", godip.Land).Conn("kha", godip.Land).Conn("ahn", godip.Land).Flag(godip.Land).SC(Bahmana).
		// Bangalore
		Prov("ban").Conn("pul", godip.Coast...).Conn("bay", godip.Sea).Conn("ori", godip.Coast...).Conn("war", godip.Land).Conn("bid", godip.Land).Conn("hon", godip.Land).Conn("cal", godip.Land).Flag(godip.Coast...).SC(Vijayanagar).
		// Bikaner
		Prov("bik").Conn("jap", godip.Land).Conn("mul", godip.Coast...).Conn("pes", godip.Coast...).Conn("jas", godip.Coast...).Conn("jod", godip.Land).Conn("mew", godip.Land).Flag(godip.Coast...).
		// Herat
		Prov("her").Conn("mes", godip.Land).Conn("yez", godip.Land).Conn("kna", godip.Land).Conn("bal", godip.Land).Conn("buk", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Elburz
		Prov("elb").Conn("qom", godip.Land).Conn("das", godip.Land).Conn("mes", godip.Land).Conn("kar", godip.Land).Flag(godip.Land).
		// NICOBAR SEA
		Prov("nic").Conn("bay", godip.Sea).Conn("pul", godip.Sea).Conn("tin", godip.Sea).Conn("lac", godip.Sea).Conn("jaf", godip.Sea).Conn("kay", godip.Sea).Conn("sea", godip.Sea).Conn("ace", godip.Sea).Conn("and", godip.Sea).Flag(godip.Sea).
		// Honavar
		Prov("hon").Conn("cal", godip.Coast...).Conn("ban", godip.Land).Conn("bid", godip.Land).Conn("bij", godip.Land).Conn("goa", godip.Coast...).Conn("lac", godip.Sea).Flag(godip.Coast...).
		// Gujarat
		Prov("guj").Conn("jod", godip.Land).Conn("jas", godip.Land).Conn("sin", godip.Coast...).Conn("ars", godip.Sea).Conn("ahb", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Meshed
		Prov("mes").Conn("her", godip.Land).Conn("buk", godip.Land).Conn("kar", godip.Land).Conn("elb", godip.Land).Conn("das", godip.Land).Conn("yez", godip.Land).Flag(godip.Land).SC(Persia).
		// Ava
		Prov("ava").Conn("nep", godip.Land).Conn("muz", godip.Land).Conn("ass", godip.Land).Conn("peg", godip.Coast...).Conn("ayu", godip.Land).Flag(godip.Coast...).
		// Pegu
		Prov("peg").Conn("bay", godip.Sea).Conn("ayu", godip.Coast...).Conn("ava", godip.Coast...).Conn("ass", godip.Land).Conn("beg", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Oman
		Prov("oma").Conn("had", godip.Coast...).Conn("ars", godip.Sea).Conn("per", godip.Sea).Conn("arb", godip.Coast...).Conn("rub", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Agra
		Prov("agr").Conn("jap", godip.Land).Conn("bee", godip.Coast...).Conn("muz", godip.Coast...).Conn("awa", godip.Coast...).Conn("kam", godip.Land).Conn("lah", godip.Land).Flag(godip.Coast...).SC(Delhi).
		// Jodhpur
		Prov("jod").Conn("ahb", godip.Land).Conn("mew", godip.Land).Conn("bik", godip.Land).Conn("jas", godip.Land).Conn("guj", godip.Land).Flag(godip.Land).SC(Rajputana).
		// Assam
		Prov("ass").Conn("ava", godip.Land).Conn("muz", godip.Coast...).Conn("beg", godip.Coast...).Conn("peg", godip.Land).Flag(godip.Coast...).
		// Sind
		Prov("sin").Conn("ars", godip.Sea).Conn("guj", godip.Coast...).Conn("jas", godip.Coast...).Conn("pes", godip.Coast...).Conn("shi", godip.Land).Conn("hor", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Goa
		Prov("goa").Conn("bij", godip.Land).Conn("ahn", godip.Coast...).Conn("ars", godip.Sea).Conn("lac", godip.Sea).Conn("hon", godip.Coast...).Flag(godip.Coast...).SC(Bahmana).
		// MALDIVE SEA
		Prov("mas").Conn("sea", godip.Sea).Conn("kay", godip.Sea).Conn("lac", godip.Sea).Conn("sey", godip.Sea).Flag(godip.Sea).
		// Hormuz
		Prov("hor").Conn("shi", godip.Land).Conn("isf", godip.Coast...).Conn("per", godip.Sea).Conn("ars", godip.Sea).Conn("sin", godip.Coast...).Flag(godip.Coast...).SC(Persia).
		// GULF OF ADEN
		Prov("gul").Conn("sey", godip.Sea).Conn("lac", godip.Sea).Conn("ars", godip.Sea).Conn("had", godip.Sea).Conn("yem", godip.Sea).Flag(godip.Sea).
		// Khandesh
		Prov("kha").Conn("ahb", godip.Land).Conn("ahn", godip.Land).Conn("bij", godip.Land).Conn("bid", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Qom
		Prov("qom").Conn("per", godip.Sea).Conn("isf", godip.Coast...).Conn("das", godip.Land).Conn("elb", godip.Land).Flag(godip.Coast...).
		// Ahmadabad
		Prov("ahb").Conn("jod", godip.Land).Conn("guj", godip.Coast...).Conn("ars", godip.Sea).Conn("ahn", godip.Coast...).Conn("kha", godip.Land).Conn("bid", godip.Land).Conn("maw", godip.Land).Conn("mew", godip.Land).Flag(godip.Coast...).
		// Ahmadnagar
		Prov("ahn").Conn("ars", godip.Sea).Conn("goa", godip.Coast...).Conn("bij", godip.Land).Conn("kha", godip.Land).Conn("ahb", godip.Coast...).Flag(godip.Coast...).SC(Bahmana).
		// Kandy
		Prov("kay").Conn("jaf", godip.Coast...).Conn("lac", godip.Sea).Conn("mas", godip.Sea).Conn("sea", godip.Sea).Conn("nic", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Hadramaut
		Prov("had").Conn("oma", godip.Coast...).Conn("rub", godip.Land).Conn("yem", godip.Coast...).Conn("gul", godip.Sea).Conn("ars", godip.Sea).Flag(godip.Coast...).
		// Badakhshan
		Prov("bad").Conn("sad", godip.Land).Conn("bal", godip.Land).Conn("kab", godip.Land).Conn("kam", godip.Land).Conn("fer", godip.Land).Flag(godip.Land).SC(Mughalistan).
		// Orissa
		Prov("ori").Conn("sab", godip.Coast...).Conn("war", godip.Land).Conn("ban", godip.Coast...).Conn("bay", godip.Sea).Conn("beg", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// SEA OF CEYLON
		Prov("sea").Conn("ace", godip.Sea).Conn("nic", godip.Sea).Conn("kay", godip.Sea).Conn("mas", godip.Sea).Flag(godip.Sea).
		// Malwa
		Prov("maw").Conn("mew", godip.Land).Conn("ahb", godip.Land).Conn("bid", godip.Land).Conn("ber", godip.Land).Conn("bee", godip.Land).Conn("jap", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ferghana
		Prov("fer").Conn("sad", godip.Land).Conn("bad", godip.Land).Conn("kam", godip.Land).Conn("kag", godip.Land).Flag(godip.Land).
		// Multan
		Prov("mul").Conn("lah", godip.Coast...).Conn("pes", godip.Coast...).Conn("bik", godip.Coast...).Conn("jap", godip.Land).Flag(godip.Coast...).SC(Rajputana).
		// Calicut
		Prov("cal").Conn("hon", godip.Coast...).Conn("lac", godip.Sea).Conn("tin", godip.Coast...).Conn("pul", godip.Land).Conn("ban", godip.Land).Flag(godip.Coast...).SC(Vijayanagar).
		// Benares
		Prov("bee").Conn("ber", godip.Land).Conn("jab", godip.Land).Conn("sab", godip.Coast...).Conn("beg", godip.Coast...).Conn("muz", godip.Coast...).Conn("agr", godip.Coast...).Conn("jap", godip.Land).Conn("maw", godip.Land).Flag(godip.Coast...).
		// ARABIAN SEA
		Prov("ars").Conn("hor", godip.Sea).Conn("per", godip.Sea).Conn("oma", godip.Sea).Conn("had", godip.Sea).Conn("gul", godip.Sea).Conn("lac", godip.Sea).Conn("goa", godip.Sea).Conn("ahn", godip.Sea).Conn("ahb", godip.Sea).Conn("guj", godip.Sea).Conn("sin", godip.Sea).Flag(godip.Sea).
		// Arabia
		Prov("arb").Conn("yem", godip.Land).Conn("rub", godip.Land).Conn("oma", godip.Coast...).Conn("per", godip.Sea).Flag(godip.Coast...).
		// LACCADIVE SEA
		Prov("lac").Conn("sey", godip.Sea).Conn("mas", godip.Sea).Conn("kay", godip.Sea).Conn("jaf", godip.Sea).Conn("nic", godip.Sea).Conn("tin", godip.Sea).Conn("cal", godip.Sea).Conn("hon", godip.Sea).Conn("goa", godip.Sea).Conn("ars", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// BAY OF BENGAL
		Prov("bay").Conn("nic", godip.Sea).Conn("and", godip.Sea).Conn("ayu", godip.Sea).Conn("peg", godip.Sea).Conn("beg", godip.Sea).Conn("ori", godip.Sea).Conn("ban", godip.Sea).Conn("pul", godip.Sea).Flag(godip.Sea).
		// Yezd
		Prov("yez").Conn("isf", godip.Land).Conn("shi", godip.Land).Conn("kna", godip.Land).Conn("her", godip.Land).Conn("mes", godip.Land).Conn("das", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"sab": "Sambalpur",
	"kar": "Kara Kum",
	"mew": "Mewar",
	"tin": "Tinnevelly",
	"lah": "Lahore",
	"beg": "Bengal",
	"ace": "Aceh",
	"jas": "Jaisalmer",
	"war": "Warangal",
	"nep": "Nepal",
	"muz": "Muzaffarpur",
	"kam": "Kashmir",
	"ber": "Berar",
	"bal": "Balkh",
	"kra": "Kra",
	"das": "Dasht-I-Kavir",
	"kna": "Kandahar",
	"buk": "Bukhara",
	"sad": "Samarkand",
	"yem": "Yemen",
	"pes": "Peshawar",
	"awa": "Awadh",
	"per": "PERSIAN GULF",
	"kag": "Kashgar",
	"ayu": "Ayutthaya",
	"shi": "Shiraz",
	"sey": "Seylac",
	"isf": "Isfahan",
	"rub": "Rub Al Khali",
	"kab": "Kabul",
	"pul": "Pulicat",
	"jaf": "Jaffna",
	"and": "ANDAMAN SEA",
	"jap": "Jaipur",
	"gar": "Gartok",
	"jab": "Jabalpur",
	"rai": "Raipur",
	"bid": "Bidar",
	"bij": "Bijapur",
	"ban": "Bangalore",
	"bik": "Bikaner",
	"her": "Herat",
	"elb": "Elburz",
	"nic": "NICOBAR SEA",
	"hon": "Honavar",
	"guj": "Gujarat",
	"mes": "Meshed",
	"ava": "Ava",
	"peg": "Pegu",
	"oma": "Oman",
	"agr": "Agra",
	"jod": "Jodhpur",
	"ass": "Assam",
	"sin": "Sind",
	"goa": "Goa",
	"mas": "MALDIVE SEA",
	"hor": "Hormuz",
	"gul": "GULF OF ADEN",
	"kha": "Khandesh",
	"qom": "Qom",
	"ahb": "Ahmadabad",
	"ahn": "Ahmadnagar",
	"kay": "Kandy",
	"had": "Hadramaut",
	"bad": "Badakhshan",
	"ori": "Orissa",
	"sea": "SEA OF CEYLON",
	"maw": "Malwa",
	"fer": "Ferghana",
	"mul": "Multan",
	"cal": "Calicut",
	"bee": "Benares",
	"ars": "ARABIAN SEA",
	"arb": "Arabia",
	"lac": "LACCADIVE SEA",
	"bay": "BAY OF BENGAL",
	"yez": "Yezd",
}
