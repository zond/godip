package classicalcrowded

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Balkans       godip.Nation = "Balkans"
	Benelux       godip.Nation = "Benelux"
	Iberia        godip.Nation = "Iberia"
	Scandinavia   godip.Nation = "Scandinavia"
)


var (
	Nations    = []godip.Nation{godip.Austria, Balkans, Benelux, godip.England, godip.France, godip.Germany, Iberia, godip.Italy, godip.Turkey, godip.Russia, Scandinavia}
	SVGFlags = map[godip.Nation]func() ([]byte, error){
		Balkans: func() ([]byte, error) {
			return Asset("svg/balkans.svg")
		},
		Benelux: func() ([]byte, error) {
			return Asset("svg/benelux.svg")
		},
		Iberia: func() ([]byte, error) {
			return Asset("svg/iberia.svg")
		},
		Scandinavia: func() ([]byte, error) {
			return Asset("svg/scandinavia.svg")
		},
		godip.Austria: func() ([]byte, error) {
			return Asset("../classical/svg/austria.svg")
		},
		godip.England: func() ([]byte, error) {
			return Asset("../classical/svg/england.svg")
		},
		godip.France: func() ([]byte, error) {
			return Asset("../classical/svg/france.svg")
		},
		godip.Germany: func() ([]byte, error) {
			return Asset("../classical/svg/germany.svg")
		},
		godip.Italy: func() ([]byte, error) {
			return Asset("../classical/svg/italy.svg")
		},
		godip.Russia: func() ([]byte, error) {
			return Asset("../classical/svg/russia.svg")
		},
		godip.Turkey: func() ([]byte, error) {
			return Asset("../classical/svg/turkey.svg")
		},
	}
	Parser = classical.Parser
)

var ClassicalCrowdedVariant = common.Variant{
	Name:              "Classical - Crowded",
	Graph:             func() godip.Graph { return ClassicalCrowdedGraph() },
	Start:             ClassicalCrowdedStart,
	Blank:             ClassicalCrowdedBlank,
	Phase:             classical.NewPhase,
	Parser:            Parser,
	Nations:           Nations,
	NationColors:      map[godip.Nation]string{
		Balkans:       "#CD926A",
		Benelux:       "#F79D10",
		Iberia:        "#A6517B",
		Scandinavia:   "#416285",
	},
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(18),
	SoloSCCount:       func(*state.State) int { return 18 },
	ProvinceLongNames: classical.ClassicalVariant.ProvinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/classicalcrowdedmap.svg")
	},
	SVGVersion: "1",
	SVGUnits:    classical.SVGUnits,
	SVGFlags:    SVGFlags,
	CreatedBy:   "Unknown",
	Version:     "1",
	Description: "An expanded version of the original Diplomacy containing new nations of the Balkans, Benelux, Iberia, and Scandinavia. The standard map is used with the exception of having a supply center in Ruhr.",
	Rules:       `The first to 18 Supply Centers (SC) is the winner. 
Kiel and Constantinople have a canal, so fleets can exit on either side. 
Armies can move from Denmark to Kiel.`,
	ExtraDominanceRules: map[godip.Province]common.DominanceRule{
		"gas": common.DominanceRule{
			Priority: 0,
			Nation:   godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"bre": godip.France,
				"par": godip.France,
				"mar": godip.France,
				"spa": Iberia,
			},
		},
		"bur": common.DominanceRule{
			Nation: godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"mar": godip.France,
				"par": godip.France,
				"mun": godip.Germany,
				"bel": Benelux,
			},
		},
		"pie": common.DominanceRule{
			Nation: godip.Italy,
			Dependencies: map[godip.Province]godip.Nation{
				"mar": godip.France,
				"ven": godip.Italy,
				"rom": godip.Italy,
			},
		},
		"pic": common.DominanceRule{
			Nation: godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"par": godip.France,
				"bre": godip.France,
				"bel": Benelux,
			},
		},
		"tyr": common.DominanceRule{
			Nation: godip.Austria,
			Dependencies: map[godip.Province]godip.Nation{
				"mun": godip.Germany,
				"ven": godip.Italy,
				"tri": godip.Austria,
				"vie": godip.Austria,
			},
		},
		"boh": common.DominanceRule{
			Nation: godip.Austria,
			Dependencies: map[godip.Province]godip.Nation{
				"mun": godip.Germany,
				"vie": godip.Austria,
			},
		},
		"gal": common.DominanceRule{
			Nation: godip.Austria,
			Dependencies: map[godip.Province]godip.Nation{
				"bud": godip.Austria,
				"vie": godip.Austria,
				"war": godip.Russia,
				"rum": Balkans,
			},
		},
		"ukr": common.DominanceRule{
			Nation: godip.Russia,
			Dependencies: map[godip.Province]godip.Nation{
				"war": godip.Russia,
				"mos": godip.Russia,
				"stp": godip.Russia,
				"rum": Balkans,
			},
		},
		"fin": common.DominanceRule{
			Nation: godip.Russia,
			Dependencies: map[godip.Province]godip.Nation{
				"stp": godip.Russia,
				"swe": Scandinavia,
				"nwy": Scandinavia,
			},
		},
		"sil": common.DominanceRule{
			Nation: godip.Germany,
			Dependencies: map[godip.Province]godip.Nation{
				"ber": godip.Germany,
				"mun": godip.Germany,
				"war": godip.Russia,
			},
		},
		"pru": common.DominanceRule{
			Nation: godip.Germany,
			Dependencies: map[godip.Province]godip.Nation{
				"ber": godip.Germany,
				"war": godip.Russia,
			},
		},
		"arm": common.DominanceRule{
			Nation: godip.Turkey,
			Dependencies: map[godip.Province]godip.Nation{
				"ank": godip.Turkey,
				"smy": godip.Turkey,
				"sev": godip.Russia,
			},
		},
		"alb": common.DominanceRule{
			Nation: Balkans,
			Dependencies: map[godip.Province]godip.Nation{
				"ser": Balkans,
				"gre": Balkans,
				"tri": godip.Austria,
			},
		},
	},
}

func ClassicalCrowdedBlank(phase godip.Phase) *state.State {
	return state.New(ClassicalCrowdedGraph(), phase, classical.BackupRule, nil, nil)
}

func AdjustSCs(phase *phase.Phase) bool {
	return phase.Ty == godip.Retreat && phase.Se == godip.Fall
}

func NewPhase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return phase.Generator(Parser, AdjustSCs)(year, season, typ)
}

func ClassicalCrowdedStart() (result *state.State, err error) {
	result = ClassicalCrowdedBlank(NewPhase(1901, godip.Spring, godip.Movement))
	if err = result.SetUnits(ClassicalCrowdedUnits()); err != nil {
		return
	}
	result.SetSupplyCenters(ClassicalCrowdedSupplyCenters())
	return
}

func ClassicalCrowdedUnits() map[godip.Province]godip.Unit {
	return map[godip.Province]godip.Unit{
		"edi":    godip.Unit{godip.Fleet, godip.England},
		"lvp":    godip.Unit{godip.Army, godip.England},
		"lon":    godip.Unit{godip.Fleet, godip.England},
		"bre":    godip.Unit{godip.Fleet, godip.France},
		"par":    godip.Unit{godip.Army, godip.France},
		"mar":    godip.Unit{godip.Army, godip.France},
		"kie":    godip.Unit{godip.Fleet, godip.Germany},
		"ber":    godip.Unit{godip.Army, godip.Germany},
		"mun":    godip.Unit{godip.Army, godip.Germany},
		"ven":    godip.Unit{godip.Army, godip.Italy},
		"rom":    godip.Unit{godip.Army, godip.Italy},
		"nap":    godip.Unit{godip.Fleet, godip.Italy},
		"tri":    godip.Unit{godip.Fleet, godip.Austria},
		"vie":    godip.Unit{godip.Army, godip.Austria},
		"bud":    godip.Unit{godip.Army, godip.Austria},
		"stp/sc": godip.Unit{godip.Fleet, godip.Russia},
		"mos":    godip.Unit{godip.Army, godip.Russia},
		"war":    godip.Unit{godip.Army, godip.Russia},
		"sev":    godip.Unit{godip.Fleet, godip.Russia},
		"con":    godip.Unit{godip.Army, godip.Turkey},
		"smy":    godip.Unit{godip.Army, godip.Turkey},
		"ank":    godip.Unit{godip.Fleet, godip.Turkey},
		"nwy":    godip.Unit{godip.Fleet, Scandinavia},
		"swe":    godip.Unit{godip.Army, Scandinavia},
		"den":    godip.Unit{godip.Fleet, Scandinavia},
		"hol":    godip.Unit{godip.Army, Benelux},
		"bel":    godip.Unit{godip.Fleet, Benelux},
		"ruh":    godip.Unit{godip.Army, Benelux},
		"spa":    godip.Unit{godip.Army, Iberia},
		"por":    godip.Unit{godip.Fleet, Iberia},
		"tun":    godip.Unit{godip.Fleet, Iberia},
		"rum":    godip.Unit{godip.Fleet, Balkans},
		"ser":    godip.Unit{godip.Army, Balkans},
		"bul":    godip.Unit{godip.Army, Balkans},
		"gre":    godip.Unit{godip.Fleet, Balkans},
	}
}

func ClassicalCrowdedSupplyCenters() map[godip.Province]godip.Nation {
	return map[godip.Province]godip.Nation{
		"edi": godip.England,
		"lvp": godip.England,
		"lon": godip.England,
		"bre": godip.France,
		"par": godip.France,
		"mar": godip.France,
		"kie": godip.Germany,
		"ber": godip.Germany,
		"mun": godip.Germany,
		"ven": godip.Italy,
		"rom": godip.Italy,
		"nap": godip.Italy,
		"tri": godip.Austria,
		"vie": godip.Austria,
		"bud": godip.Austria,
		"con": godip.Turkey,
		"ank": godip.Turkey,
		"smy": godip.Turkey,
		"sev": godip.Russia,
		"mos": godip.Russia,
		"stp": godip.Russia,
		"war": godip.Russia,
		"nwy": Scandinavia,
		"swe": Scandinavia,
		"den": Scandinavia,
		"hol": Benelux,
		"bel": Benelux,
		"ruh": Benelux,
		"spa": Iberia,
		"por": Iberia,
		"tun": Iberia,
		"rum": Balkans,
		"ser": Balkans,
		"bul": Balkans,
		"gre": Balkans,
	}
}

func ClassicalCrowdedGraph() *graph.Graph {
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
		Prov("stp").Conn("fin", godip.Land).Conn("nwy", godip.Land).Conn("mos", godip.Land).Conn("lvn", godip.Land).Flag(godip.Land).Conn("bar", godip.Sea).Conn("bot", godip.Sea).SC(godip.Russia).
		// mos
		Prov("mos").Conn("stp", godip.Land).Conn("sev", godip.Land).Conn("ukr", godip.Land).Conn("war", godip.Land).Conn("lvn", godip.Land).Flag(godip.Land).SC(godip.Russia).
		// sev
		Prov("sev").Conn("ukr", godip.Land).Conn("mos", godip.Land).Conn("arm", godip.Coast...).Conn("bla", godip.Sea).Conn("rum", godip.Coast...).Flag(godip.Coast...).SC(godip.Russia).
		// arm
		Prov("arm").Conn("ank", godip.Coast...).Conn("bla", godip.Sea).Conn("sev", godip.Coast...).Conn("syr", godip.Land).Conn("smy", godip.Land).Flag(godip.Coast...).
		// syr
		Prov("syr").Conn("eas", godip.Sea).Conn("smy", godip.Coast...).Conn("arm", godip.Land).Flag(godip.Coast...).
		// eas
		Prov("eas").Conn("ion", godip.Sea).Conn("aeg", godip.Sea).Conn("smy", godip.Sea).Conn("syr", godip.Sea).Flag(godip.Sea).
		// ion
		Prov("ion").Conn("apu", godip.Sea).Conn("adr", godip.Sea).Conn("tun", godip.Sea).Conn("tys", godip.Sea).Conn("nap", godip.Sea).Conn("alb", godip.Sea).Conn("gre", godip.Sea).Conn("aeg", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// tun
		Prov("tun").Conn("naf", godip.Coast...).Conn("wes", godip.Sea).Conn("tys", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Iberia).
		// naf
		Prov("naf").Conn("mid", godip.Sea).Conn("wes", godip.Sea).Conn("tun", godip.Coast...).Flag(godip.Coast...).
		// mid
		Prov("mid").Conn("wes", godip.Sea).Conn("nat", godip.Sea).Conn("iri", godip.Sea).Conn("eng", godip.Sea).Conn("bre", godip.Sea).Conn("gas", godip.Sea).Conn("spa/nc", godip.Sea).Conn("por", godip.Sea).Conn("spa/sc", godip.Sea).Conn("naf", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// iri
		Prov("iri").Conn("nat", godip.Sea).Conn("lvp", godip.Sea).Conn("wal", godip.Sea).Conn("eng", godip.Sea).Conn("mid", godip.Sea).Flag(godip.Sea).
		// lvp
		Prov("lvp").Conn("iri", godip.Sea).Conn("nat", godip.Sea).Conn("cly", godip.Coast...).Conn("edi", godip.Land).Conn("yor", godip.Land).Conn("wal", godip.Coast...).Flag(godip.Coast...).SC(godip.England).
		// cly
		Prov("cly").Conn("nat", godip.Sea).Conn("nrg", godip.Sea).Conn("edi", godip.Coast...).Conn("lvp", godip.Coast...).Flag(godip.Coast...).
		// edi
		Prov("edi").Conn("cly", godip.Coast...).Conn("nrg", godip.Sea).Conn("nth", godip.Sea).Conn("yor", godip.Coast...).Conn("lvp", godip.Land).Flag(godip.Coast...).SC(godip.England).
		// nth
		Prov("nth").Conn("eng", godip.Sea).Conn("edi", godip.Sea).Conn("nrg", godip.Sea).Conn("nwy", godip.Sea).Conn("ska", godip.Sea).Conn("den", godip.Sea).Conn("hel", godip.Sea).Conn("hol", godip.Sea).Conn("bel", godip.Sea).Conn("lon", godip.Sea).Conn("yor", godip.Sea).Flag(godip.Sea).
		// nwy
		Prov("nwy").Conn("nth", godip.Sea).Conn("nrg", godip.Sea).Conn("bar", godip.Sea).Conn("stp/nc", godip.Sea).Conn("stp", godip.Land).Conn("fin", godip.Land).Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(Scandinavia).
		// stp/sc
		Prov("stp/sc").Conn("bot", godip.Sea).Conn("fin", godip.Sea).Conn("lvn", godip.Sea).Flag(godip.Sea).
		// lvn
		Prov("lvn").Conn("stp", godip.Land).Conn("bal", godip.Sea).Conn("bot", godip.Sea).Conn("stp/sc", godip.Sea).Conn("mos", godip.Land).Conn("war", godip.Land).Conn("pru", godip.Coast...).Flag(godip.Coast...).
		// war
		Prov("war").Conn("sil", godip.Land).Conn("pru", godip.Land).Conn("lvn", godip.Land).Conn("mos", godip.Land).Conn("ukr", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).SC(godip.Russia).
		// ukr
		Prov("ukr").Conn("war", godip.Land).Conn("mos", godip.Land).Conn("sev", godip.Land).Conn("rum", godip.Land).Conn("gal", godip.Land).Flag(godip.Land).
		// bla
		Prov("bla").Conn("bul/ec", godip.Sea).Conn("rum", godip.Sea).Conn("sev", godip.Sea).Conn("arm", godip.Sea).Conn("ank", godip.Sea).Conn("con", godip.Sea).Conn("bul", godip.Sea).Flag(godip.Sea).
		// ank
		Prov("ank").Conn("con", godip.Coast...).Conn("bla", godip.Sea).Conn("arm", godip.Coast...).Conn("smy", godip.Land).Flag(godip.Coast...).SC(godip.Turkey).
		// smy
		Prov("smy").Conn("aeg", godip.Sea).Conn("con", godip.Coast...).Conn("ank", godip.Land).Conn("arm", godip.Land).Conn("syr", godip.Coast...).Conn("eas", godip.Sea).Flag(godip.Coast...).SC(godip.Turkey).
		// aeg
		Prov("aeg").Conn("eas", godip.Sea).Conn("ion", godip.Sea).Conn("gre", godip.Sea).Conn("bul/sc", godip.Sea).Conn("con", godip.Sea).Conn("smy", godip.Sea).Conn("bul", godip.Sea).Flag(godip.Sea).
		// gre
		Prov("gre").Conn("ion", godip.Sea).Conn("alb", godip.Coast...).Conn("ser", godip.Land).Conn("bul", godip.Land).Conn("bul/sc", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(Balkans).
		// nap
		Prov("nap").Conn("tys", godip.Sea).Conn("rom", godip.Coast...).Conn("apu", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(godip.Italy).
		// tys
		Prov("tys").Conn("wes", godip.Sea).Conn("gol", godip.Sea).Conn("tus", godip.Sea).Conn("rom", godip.Sea).Conn("nap", godip.Sea).Conn("ion", godip.Sea).Conn("tun", godip.Sea).Flag(godip.Sea).
		// wes
		Prov("wes").Conn("mid", godip.Sea).Conn("spa/sc", godip.Sea).Conn("gol", godip.Sea).Conn("tys", godip.Sea).Conn("tun", godip.Sea).Conn("naf", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// spa/sc
		Prov("spa/sc").Conn("mid", godip.Sea).Conn("por", godip.Sea).Conn("mar", godip.Sea).Conn("gol", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// spa
		Prov("spa").Conn("por", godip.Land).Conn("gas", godip.Land).Conn("mar", godip.Land).Conn("mid", godip.Sea).Conn("gol", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Land).SC(Iberia).
		// spa/nc
		Prov("spa/nc").Conn("por", godip.Sea).Conn("mid", godip.Sea).Conn("gas", godip.Sea).Flag(godip.Sea).
		// por
		Prov("por").Conn("mid", godip.Sea).Conn("spa/nc", godip.Sea).Conn("spa", godip.Land).Conn("spa/sc", godip.Sea).Flag(godip.Coast...).SC(Iberia).
		// gas
		Prov("gas").Conn("mid", godip.Sea).Conn("bre", godip.Coast...).Conn("par", godip.Land).Conn("bur", godip.Land).Conn("mar", godip.Land).Conn("spa", godip.Land).Conn("spa/nc", godip.Sea).Flag(godip.Coast...).
		// bre
		Prov("bre").Conn("mid", godip.Sea).Conn("eng", godip.Sea).Conn("pic", godip.Coast...).Conn("par", godip.Land).Conn("gas", godip.Coast...).Flag(godip.Coast...).SC(godip.France).
		// eng
		Prov("eng").Conn("mid", godip.Sea).Conn("iri", godip.Sea).Conn("wal", godip.Sea).Conn("lon", godip.Sea).Conn("nth", godip.Sea).Conn("bel", godip.Sea).Conn("pic", godip.Sea).Conn("bre", godip.Sea).Flag(godip.Sea).
		// wal
		Prov("wal").Conn("iri", godip.Sea).Conn("lvp", godip.Coast...).Conn("yor", godip.Land).Conn("lon", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).
		// yor
		Prov("yor").Conn("lvp", godip.Land).Conn("edi", godip.Coast...).Conn("nth", godip.Sea).Conn("lon", godip.Coast...).Conn("wal", godip.Land).Flag(godip.Coast...).
		// ska
		Prov("ska").Conn("nth", godip.Sea).Conn("nwy", godip.Sea).Conn("swe", godip.Sea).Conn("den", godip.Sea).Flag(godip.Sea).
		// swe
		Prov("swe").Conn("ska", godip.Sea).Conn("nwy", godip.Coast...).Conn("fin", godip.Coast...).Conn("bot", godip.Sea).Conn("bal", godip.Sea).Conn("den", godip.Coast...).Flag(godip.Coast...).SC(Scandinavia).
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
		Prov("rum").Conn("bla", godip.Sea).Conn("bud", godip.Land).Conn("gal", godip.Land).Conn("ukr", godip.Land).Conn("sev", godip.Coast...).Conn("bul/ec", godip.Sea).Conn("bul", godip.Land).Conn("ser", godip.Land).Flag(godip.Coast...).SC(Balkans).
		// bul/ec
		Prov("bul/ec").Conn("rum", godip.Sea).Conn("bla", godip.Sea).Conn("con", godip.Sea).Flag(godip.Sea).
		// bul
		Prov("bul").Conn("ser", godip.Land).Conn("rum", godip.Land).Conn("con", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).Conn("aeg", godip.Sea).Conn("bla", godip.Sea).SC(Balkans).
		// con
		Prov("con").Conn("bul/sc", godip.Sea).Conn("bul", godip.Land).Conn("bul/ec", godip.Sea).Conn("bla", godip.Sea).Conn("ank", godip.Coast...).Conn("smy", godip.Coast...).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(godip.Turkey).
		// bul/sc
		Prov("bul/sc").Conn("gre", godip.Sea).Conn("con", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Sea).
		// ser
		Prov("ser").Conn("tri", godip.Land).Conn("bud", godip.Land).Conn("rum", godip.Land).Conn("bul", godip.Land).Conn("gre", godip.Land).Conn("alb", godip.Land).Flag(godip.Land).SC(Balkans).
		// alb
		Prov("alb").Conn("adr", godip.Sea).Conn("tri", godip.Coast...).Conn("ser", godip.Land).Conn("gre", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).
		// adr
		Prov("adr").Conn("ven", godip.Sea).Conn("tri", godip.Sea).Conn("alb", godip.Sea).Conn("ion", godip.Sea).Conn("apu", godip.Sea).Flag(godip.Sea).
		// apu
		Prov("apu").Conn("rom", godip.Land).Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("nap", godip.Coast...).Flag(godip.Coast...).
		// rom
		Prov("rom").Conn("tys", godip.Sea).Conn("tus", godip.Coast...).Conn("ven", godip.Land).Conn("apu", godip.Land).Conn("nap", godip.Coast...).Flag(godip.Coast...).SC(godip.Italy).
		// tus
		Prov("tus").Conn("gol", godip.Sea).Conn("pie", godip.Coast...).Conn("ven", godip.Land).Conn("rom", godip.Coast...).Conn("tys", godip.Sea).Flag(godip.Coast...).
		// gol
		Prov("gol").Conn("spa/sc", godip.Sea).Conn("mar", godip.Sea).Conn("pie", godip.Sea).Conn("tus", godip.Sea).Conn("tys", godip.Sea).Conn("wes", godip.Sea).Conn("spa", godip.Sea).Flag(godip.Sea).
		// mar
		Prov("mar").Conn("spa", godip.Land).Conn("gas", godip.Land).Conn("bur", godip.Land).Conn("pie", godip.Coast...).Conn("gol", godip.Sea).Conn("spa/sc", godip.Sea).Flag(godip.Coast...).SC(godip.France).
		// bur
		Prov("bur").Conn("par", godip.Land).Conn("pic", godip.Land).Conn("bel", godip.Land).Conn("ruh", godip.Land).Conn("mun", godip.Land).Conn("mar", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).
		// par
		Prov("par").Conn("bre", godip.Land).Conn("pic", godip.Land).Conn("bur", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).SC(godip.France).
		// pic
		Prov("pic").Conn("bre", godip.Coast...).Conn("eng", godip.Sea).Conn("bel", godip.Coast...).Conn("bur", godip.Land).Conn("par", godip.Land).Flag(godip.Coast...).
		// lon
		Prov("lon").Conn("wal", godip.Coast...).Conn("yor", godip.Coast...).Conn("nth", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(godip.England).
		// bel
		Prov("bel").Conn("pic", godip.Coast...).Conn("eng", godip.Sea).Conn("nth", godip.Sea).Conn("hol", godip.Coast...).Conn("ruh", godip.Land).Conn("bur", godip.Land).Flag(godip.Coast...).SC(Benelux).
		// hol
		Prov("hol").Conn("nth", godip.Sea).Conn("hel", godip.Sea).Conn("kie", godip.Coast...).Conn("ruh", godip.Land).Conn("bel", godip.Coast...).Flag(godip.Coast...).SC(Benelux).
		// hel
		Prov("hel").Conn("nth", godip.Sea).Conn("den", godip.Sea).Conn("kie", godip.Sea).Conn("hol", godip.Sea).Flag(godip.Sea).
		// den
		Prov("den").Conn("hel", godip.Sea).Conn("nth", godip.Sea).Conn("ska", godip.Sea).Conn("swe", godip.Coast...).Conn("bal", godip.Sea).Conn("kie", godip.Coast...).Flag(godip.Coast...).SC(Scandinavia).
		// ber
		Prov("ber").Conn("kie", godip.Coast...).Conn("bal", godip.Sea).Conn("pru", godip.Coast...).Conn("sil", godip.Land).Conn("mun", godip.Land).Flag(godip.Coast...).SC(godip.Germany).
		// mun
		Prov("mun").Conn("bur", godip.Land).Conn("ruh", godip.Land).Conn("kie", godip.Land).Conn("ber", godip.Land).Conn("sil", godip.Land).Conn("boh", godip.Land).Conn("tyr", godip.Land).Flag(godip.Land).SC(godip.Germany).
		// boh
		Prov("boh").Conn("mun", godip.Land).Conn("sil", godip.Land).Conn("gal", godip.Land).Conn("vie", godip.Land).Conn("tyr", godip.Land).Flag(godip.Land).
		// vie
		Prov("vie").Conn("tyr", godip.Land).Conn("boh", godip.Land).Conn("gal", godip.Land).Conn("bud", godip.Land).Conn("tri", godip.Land).Flag(godip.Land).SC(godip.Austria).
		// bud
		Prov("bud").Conn("tri", godip.Land).Conn("vie", godip.Land).Conn("gal", godip.Land).Conn("rum", godip.Land).Conn("ser", godip.Land).Flag(godip.Land).SC(godip.Austria).
		// tri
		Prov("tri").Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("tyr", godip.Land).Conn("vie", godip.Land).Conn("bud", godip.Land).Conn("ser", godip.Land).Conn("alb", godip.Coast...).Flag(godip.Coast...).SC(godip.Austria).
		// ven
		Prov("ven").Conn("tus", godip.Land).Conn("pie", godip.Land).Conn("tyr", godip.Land).Conn("tri", godip.Coast...).Conn("adr", godip.Sea).Conn("apu", godip.Coast...).Conn("rom", godip.Land).Flag(godip.Coast...).SC(godip.Italy).
		// pie
		Prov("pie").Conn("mar", godip.Coast...).Conn("tyr", godip.Land).Conn("ven", godip.Land).Conn("tus", godip.Coast...).Conn("gol", godip.Sea).Flag(godip.Coast...).
		// ruh
		Prov("ruh").Conn("bel", godip.Land).Conn("hol", godip.Land).Conn("kie", godip.Land).Conn("mun", godip.Land).Conn("bur", godip.Land).Flag(godip.Land).SC(Benelux).
		// tyr
		Prov("tyr").Conn("mun", godip.Land).Conn("boh", godip.Land).Conn("vie", godip.Land).Conn("tri", godip.Land).Conn("ven", godip.Land).Conn("pie", godip.Land).Flag(godip.Land).
		// kie
		Prov("kie").Conn("hol", godip.Coast...).Conn("hel", godip.Sea).Conn("den", godip.Coast...).Conn("bal", godip.Sea).Conn("ber", godip.Coast...).Conn("mun", godip.Land).Conn("ruh", godip.Land).Flag(godip.Coast...).SC(godip.Germany).
		Done()
}