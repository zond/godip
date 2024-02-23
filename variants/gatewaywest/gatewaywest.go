package gatewaywest

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Chickasaw godip.Nation = "Chickasaw"
	Osage     godip.Nation = "Osage"
	Missouria godip.Nation = "Missouria"
	Illini    godip.Nation = "Illini"
	Otoe      godip.Nation = "Otoe"
	Ioway     godip.Nation = "Ioway"
	Quapaw    godip.Nation = "Quapaw"
)

var Nations = []godip.Nation{Chickasaw, Osage, Missouria, Illini, Otoe, Ioway, Quapaw}

var GatewayWestVariant = common.Variant{
	Name:              "Gateway West",
	NationColors: map[godip.Nation]string{
		Chickasaw: "#1B316D",
		Osage:     "#97CBFF",
		Missouria: "#8D5117",
		Illini:    "#FF5F05",
		Otoe:      "#850405",
		Ioway:     "#009396",
		Quapaw:    "#F2CC00",
	},
	Graph:             func() godip.Graph { return GatewayWestGraph() },
	Start:             GatewayWestStart,
	Blank:             GatewayWestBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(18),
	SoloSCCount:       func(*state.State) int { return 18 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/gatewaywestmap.svg")
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
	SVGFlags: map[godip.Nation]func() ([]byte, error){
		Chickasaw: func() ([]byte, error) {
			return Asset("svg/Chickasaw.svg")
		},
		Osage: func() ([]byte, error) {
			return Asset("svg/Osage.svg")
		},
		Illini: func() ([]byte, error) {
			return Asset("svg/Illini.svg")
		},
		Missouria: func() ([]byte, error) {
			return Asset("svg/Missouria.svg")
		},
		Otoe: func() ([]byte, error) {
			return Asset("svg/Otoe.svg")
		},
		Ioway: func() ([]byte, error) {
			return Asset("svg/Ioway.svg")
		},
		Quapaw: func() ([]byte, error) {
			return Asset("svg/Quapaw.svg")
		},
	},
	CreatedBy:   "TJ Brooks",
	Version:     "1",
	Description: `Alternative history variant. Missouri has recently 
	achieved statehood, and the exploration of the western territories 
	is gaining momentum as settlers seek new opportunities. However, 
	indigenous tribes continue to inhabit these lands, and they aspire 
	to organize and reclaim the territories they once freely roamed.`,
	Rules:       `The first to 18 Supply Centers (SC) is the winner.`,
	ExtraDominanceRules: map[godip.Province]common.DominanceRule{
		"sog": common.DominanceRule{
			Priority: 0,
			Nation:   Missouria,
			Dependencies: map[godip.Province]godip.Nation{
				"jac": Missouria,
				"laf": Missouria,
			},
		},
		"joh": common.DominanceRule{
			Nation: Missouria,
			Dependencies: map[godip.Province]godip.Nation{
				"sal": Missouria,
				"laf": Missouria,
			},
		},
		"pet": common.DominanceRule{
			Nation: Missouria,
			Dependencies: map[godip.Province]godip.Nation{
				"sal": Missouria,
				"laf": Missouria,
				"col": Osage,
				"oza": Osage,
			},
		},
		"whi": common.DominanceRule{
			Nation: Osage,
			Dependencies: map[godip.Province]godip.Nation{
				"oza": Osage,
				"bla": Quapaw,
			},
		},
		"big": common.DominanceRule{
			Nation: Osage,
			Dependencies: map[godip.Province]godip.Nation{
				"oza": Osage,
				"osa": Osage,
				"bla": Quapaw,
			},
		},
		"gas": common.DominanceRule{
			Nation: Osage,
			Dependencies: map[godip.Province]godip.Nation{
				"osa": Osage,
				"fra": godip.Neutral,
			},
		},
		"gre": common.DominanceRule{
			Nation: Quapaw,
			Dependencies: map[godip.Province]godip.Nation{
				"way": Quapaw,
				"but": Quapaw,
				"stf": Chickasaw,
			},
		},
		"mad": common.DominanceRule{
			Nation: Quapaw,
			Dependencies: map[godip.Province]godip.Nation{
				"way": Quapaw,
				"but": Quapaw,
				"ste": godip.Neutral,
				"per": godip.Neutral,
			},
		},
		"cap": common.DominanceRule{
			Nation: Chickasaw,
			Dependencies: map[godip.Province]godip.Nation{
				"sco": Chickasaw,
				"stf": Chickasaw,
				"per": godip.Neutral,
			},
		},
		"mus": common.DominanceRule{
			Nation: Ioway,
			Dependencies: map[godip.Province]godip.Nation{
				"gra": Ioway,
			},
		},
		"pik": common.DominanceRule{
			Nation: Illini,
			Dependencies: map[godip.Province]godip.Nation{
				"lin": Illini,
				"stc": Illini,
			},
		},
		"aud": common.DominanceRule{
			Nation: Missouria,
			Dependencies: map[godip.Province]godip.Nation{
				"lin": Illini,
				"stc": Illini,
				"cal": godip.Neutral,
			},
		},
		"mot": common.DominanceRule{
			Nation: Missouria,
			Dependencies: map[godip.Province]godip.Nation{
				"lin": Illini,
				"stc": Illini,
				"cal": godip.Neutral,
			},
		},
	},
}

func GatewayWestBlank(phase godip.Phase) *state.State {
	return state.New(GatewayWestGraph(), phase, classical.BackupRule, nil, nil)
}

func GatewayWestStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1826, godip.Spring, godip.Movement)
	result = GatewayWestBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"new": godip.Unit{godip.Fleet, Chickasaw},
		"sco": godip.Unit{godip.Army, Chickasaw},
		"stf": godip.Unit{godip.Army, Chickasaw},
		"osa": godip.Unit{godip.Army, Osage},
		"oza": godip.Unit{godip.Army, Osage},
		"col": godip.Unit{godip.Army, Osage},
		"sal": godip.Unit{godip.Fleet, Missouria},
		"laf": godip.Unit{godip.Fleet, Missouria},
		"jac": godip.Unit{godip.Army, Missouria},
		"stc": godip.Unit{godip.Fleet, Illini},
		"stl": godip.Unit{godip.Army, Illini},
		"lin": godip.Unit{godip.Army, Illini},
		"cly": godip.Unit{godip.Fleet, Otoe},
		"pla": godip.Unit{godip.Army, Otoe},
		"nis": godip.Unit{godip.Army, Otoe},
		"gra": godip.Unit{godip.Army, Ioway},
		"nog": godip.Unit{godip.Army, Ioway},
		"cli": godip.Unit{godip.Army, Ioway},
		"way": godip.Unit{godip.Army, Quapaw},
		"but": godip.Unit{godip.Army, Quapaw},
		"bla": godip.Unit{godip.Army, Quapaw},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"new": Chickasaw,
		"sco": Chickasaw,
		"stf": Chickasaw,
		"osa": Osage,
		"oza": Osage,
		"col": Osage,
		"sal": Missouria,
		"laf": Missouria,
		"jac": Missouria,
		"stc": Illini,
		"stl": Illini,
		"lin": Illini,
		"cly": Otoe,
		"pla": Otoe,
		"nis": Otoe,
		"gra": Ioway,
		"nog": Ioway,
		"cli": Ioway,
		"way": Quapaw,
		"but": Quapaw,
		"bla": Quapaw,
	})
	return
}

func GatewayWestGraph() *graph.Graph {
	return graph.New().
		// North Central Mississippi River
		Prov("npp").Conn("nmr", godip.Sea).Conn("mar", godip.Sea).Conn("pik", godip.Sea).Conn("lin", godip.Sea).Conn("cpp", godip.Sea).Flag(godip.Sea).
		// Boone
		Prov("boo").Conn("cal", godip.Coast...).Conn("mro", godip.Land).Conn("ran", godip.Land).Conn("col", godip.Land).Conn("how", godip.Coast...).Conn("ceo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Randolph
		Prov("ran").Conn("mro", godip.Land).Conn("cui", godip.Land).Conn("cha", godip.Coast...).Conn("noo", godip.Sea).Conn("how", godip.Coast...).Conn("boo", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Grand
		Prov("gra").Conn("nog", godip.Land).Conn("cli", godip.Land).Conn("ray", godip.Land).Conn("ric", godip.Land).Conn("car", godip.Land).Conn("mus", godip.Land).Flag(godip.Land).SC(Ioway).
		// South Grand
		Prov("sog").Conn("jac", godip.Land).Conn("ver", godip.Land).Conn("hen", godip.Land).Conn("joh", godip.Land).Conn("laf", godip.Land).Flag(godip.Land).
		// Audrain
		Prov("aud").Conn("cui", godip.Land).Conn("mro", godip.Land).Conn("cal", godip.Land).Conn("mot", godip.Land).Conn("lin", godip.Land).Conn("pik", godip.Land).Conn("mar", godip.Land).Flag(godip.Land).
		// East Missouri River
		Prov("emr").Conn("ecm", godip.Sea).Conn("gas", godip.Sea).Conn("fra", godip.Sea).Conn("sem", godip.Sea).Conn("mot", godip.Sea).Conn("cal", godip.Sea).Flag(godip.Sea).
		// Monroe
		Prov("mro").Conn("aud", godip.Land).Conn("cui", godip.Land).Conn("ran", godip.Land).Conn("boo", godip.Land).Conn("cal", godip.Land).Flag(godip.Land).
		// Black
		Prov("bla").Conn("bou", godip.Land).Conn("mer", godip.Land).Conn("big", godip.Land).Conn("whi", godip.Land).Conn("but", godip.Land).Conn("way", godip.Land).Flag(godip.Land).SC(Quapaw).
		// Stockton
		Prov("sto").Conn("ver", godip.Land).Conn("neo", godip.Land).Conn("tab", godip.Land).Conn("whi", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// South Mississippi River
		Prov("smr").Conn("scp", godip.Sea).Conn("ste", godip.Sea).Conn("per", godip.Sea).Conn("cap", godip.Sea).Conn("sco", godip.Sea).Conn("new", godip.Sea).Flag(godip.Sea).
		// Bourbeuse
		Prov("bou").Conn("bla", godip.Land).Conn("way", godip.Land).Conn("mad", godip.Land).Conn("ste", godip.Land).Conn("jef", godip.Land).Conn("stl", godip.Coast...).Conn("sem", godip.Sea).Conn("fra", godip.Coast...).Conn("mer", godip.Land).Flag(godip.Coast...).
		// Vernon
		Prov("ver").Conn("neo", godip.Land).Conn("sto", godip.Land).Conn("oza", godip.Land).Conn("hen", godip.Land).Conn("sog", godip.Land).Flag(godip.Land).
		// Cole
		Prov("col").Conn("pet", godip.Coast...).Conn("oza", godip.Land).Conn("boo", godip.Land).Conn("osa", godip.Coast...).Conn("ecm", godip.Sea).Conn("ceo", godip.Sea).Flag(godip.Coast...).SC(Osage).
		// Greenville
		Prov("gre").Conn("but", godip.Land).Conn("stf", godip.Land).Conn("cap", godip.Land).Conn("mad", godip.Land).Conn("way", godip.Land).Flag(godip.Land).
		// Nishnabotna
		Prov("nis").Conn("nwm", godip.Sea).Conn("pla", godip.Coast...).Conn("nod", godip.Land).Flag(godip.Coast...).SC(Otoe).
		// New Madrid
		Prov("new").Conn("smr", godip.Sea).Conn("sco", godip.Coast...).Conn("stf", godip.Land).Flag(godip.Coast...).SC(Chickasaw).
		// Scott
		Prov("sco").Conn("smr", godip.Sea).Conn("cap", godip.Coast...).Conn("stf", godip.Land).Conn("new", godip.Coast...).Flag(godip.Coast...).SC(Chickasaw).
		// Montgomery
		Prov("mot").Conn("sem", godip.Sea).Conn("stc", godip.Coast...).Conn("lin", godip.Land).Conn("aud", godip.Land).Conn("cal", godip.Coast...).Conn("emr", godip.Sea).Flag(godip.Coast...).
		// White
		Prov("whi").Conn("sto", godip.Land).Conn("tab", godip.Land).Conn("bla", godip.Land).Conn("big", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// Butler
		Prov("but").Conn("gre", godip.Land).Conn("way", godip.Land).Conn("bla", godip.Land).Conn("stf", godip.Land).Flag(godip.Land).SC(Quapaw).
		// South Central Mississippi River
		Prov("scp").Conn("sem", godip.Sea).Conn("stl", godip.Sea).Conn("jef", godip.Sea).Conn("jef", godip.Sea).Conn("ste", godip.Sea).Conn("smr", godip.Sea).Conn("cpp", godip.Sea).Flag(godip.Sea).
		// Gasconade
		Prov("gas").Conn("mer", godip.Land).Conn("fra", godip.Coast...).Conn("emr", godip.Sea).Conn("ecm", godip.Sea).Conn("osa", godip.Coast...).Conn("big", godip.Land).Flag(godip.Coast...).
		// Pettis
		Prov("pet").Conn("hen", godip.Land).Conn("oza", godip.Land).Conn("col", godip.Coast...).Conn("ceo", godip.Sea).Conn("sal", godip.Coast...).Conn("laf", godip.Land).Conn("joh", godip.Land).Flag(godip.Coast...).
		// Lincoln
		Prov("lin").Conn("npp", godip.Sea).Conn("pik", godip.Coast...).Conn("aud", godip.Land).Conn("mot", godip.Land).Conn("stc", godip.Coast...).Conn("cpp", godip.Sea).Flag(godip.Coast...).SC(Illini).
		// East Central Missouri River
		Prov("ecm").Conn("emr", godip.Sea).Conn("cal", godip.Sea).Conn("ceo", godip.Sea).Conn("col", godip.Sea).Conn("osa", godip.Sea).Conn("gas", godip.Sea).Flag(godip.Sea).
		// Perry
		Prov("per").Conn("smr", godip.Sea).Conn("ste", godip.Coast...).Conn("mad", godip.Land).Conn("cap", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Clay
		Prov("cly").Conn("nwm", godip.Sea).Conn("wes", godip.Sea).Conn("ray", godip.Coast...).Conn("jac", godip.Land).Conn("cli", godip.Land).Conn("pla", godip.Coast...).Flag(godip.Coast...).SC(Otoe).
		// Richmond
		Prov("ric").Conn("wes", godip.Sea).Conn("nwc", godip.Sea).Conn("car", godip.Coast...).Conn("gra", godip.Land).Conn("ray", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// North West Central Missouri River
		Prov("nwc").Conn("wes", godip.Sea).Conn("laf", godip.Sea).Conn("sal", godip.Sea).Conn("noo", godip.Sea).Conn("cha", godip.Sea).Conn("car", godip.Sea).Conn("ric", godip.Sea).Flag(godip.Sea).
		// Johnson
		Prov("joh").Conn("hen", godip.Land).Conn("pet", godip.Land).Conn("laf", godip.Land).Conn("sog", godip.Land).Flag(godip.Land).
		// South East Missouri River
		Prov("sem").Conn("scp", godip.Sea).Conn("cpp", godip.Sea).Conn("stc", godip.Sea).Conn("stc", godip.Sea).Conn("mot", godip.Sea).Conn("emr", godip.Sea).Conn("fra", godip.Sea).Conn("fra", godip.Sea).Conn("bou", godip.Sea).Conn("stl", godip.Sea).Flag(godip.Sea).
		// Carroll
		Prov("car").Conn("ric", godip.Coast...).Conn("nwc", godip.Sea).Conn("cha", godip.Coast...).Conn("mus", godip.Land).Conn("gra", godip.Land).Flag(godip.Coast...).
		// Howard
		Prov("how").Conn("noo", godip.Sea).Conn("ceo", godip.Sea).Conn("boo", godip.Coast...).Conn("ran", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Marion
		Prov("mar").Conn("npp", godip.Sea).Conn("nmr", godip.Sea).Conn("lew", godip.Coast...).Conn("cui", godip.Land).Conn("aud", godip.Land).Conn("pik", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Table Rock
		Prov("tab").Conn("neo", godip.Land).Conn("whi", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ste. Genevieve
		Prov("ste").Conn("bou", godip.Land).Conn("mad", godip.Land).Conn("per", godip.Coast...).Conn("smr", godip.Sea).Conn("scp", godip.Sea).Conn("jef", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// North Central Missouri River
		Prov("noo").Conn("ceo", godip.Sea).Conn("how", godip.Sea).Conn("ran", godip.Sea).Conn("cha", godip.Sea).Conn("nwc", godip.Sea).Conn("sal", godip.Sea).Flag(godip.Sea).
		// Central Mississippi River
		Prov("cpp").Conn("stc", godip.Sea).Conn("stc", godip.Sea).Conn("sem", godip.Sea).Conn("scp", godip.Sea).Conn("npp", godip.Sea).Conn("lin", godip.Sea).Flag(godip.Sea).
		// Platte
		Prov("pla").Conn("nis", godip.Coast...).Conn("nwm", godip.Sea).Conn("nwm", godip.Sea).Conn("cly", godip.Coast...).Conn("cli", godip.Land).Conn("nod", godip.Land).Flag(godip.Coast...).SC(Otoe).
		// Henry
		Prov("hen").Conn("pet", godip.Land).Conn("joh", godip.Land).Conn("sog", godip.Land).Conn("ver", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("gas", godip.Coast...).Conn("ecm", godip.Sea).Conn("col", godip.Coast...).Conn("oza", godip.Land).Conn("big", godip.Land).Flag(godip.Coast...).SC(Osage).
		// Chariton
		Prov("cha").Conn("mus", godip.Land).Conn("car", godip.Coast...).Conn("nwc", godip.Sea).Conn("noo", godip.Sea).Conn("ran", godip.Coast...).Conn("cui", godip.Land).Flag(godip.Coast...).
		// Clinton
		Prov("cli").Conn("nod", godip.Land).Conn("pla", godip.Land).Conn("cly", godip.Land).Conn("ray", godip.Land).Conn("gra", godip.Land).Conn("nog", godip.Land).Flag(godip.Land).SC(Ioway).
		// Ray
		Prov("ray").Conn("ric", godip.Coast...).Conn("gra", godip.Land).Conn("cli", godip.Land).Conn("cly", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// North Grand
		Prov("nog").Conn("cli", godip.Land).Conn("gra", godip.Land).Flag(godip.Land).SC(Ioway).
		// St. Francis
		Prov("stf").Conn("new", godip.Land).Conn("sco", godip.Land).Conn("cap", godip.Land).Conn("gre", godip.Land).Conn("but", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Clark
		Prov("lar").Conn("cui", godip.Land).Conn("lew", godip.Coast...).Conn("nmr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Neosho
		Prov("neo").Conn("ver", godip.Land).Conn("tab", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North Mississippi River
		Prov("nmr").Conn("lar", godip.Sea).Conn("lew", godip.Sea).Conn("mar", godip.Sea).Conn("npp", godip.Sea).Flag(godip.Sea).
		// Madison
		Prov("mad").Conn("way", godip.Land).Conn("gre", godip.Land).Conn("cap", godip.Land).Conn("per", godip.Land).Conn("ste", godip.Land).Conn("bou", godip.Land).Flag(godip.Land).
		// Ozark
		Prov("oza").Conn("osa", godip.Land).Conn("col", godip.Land).Conn("pet", godip.Land).Conn("hen", godip.Land).Conn("ver", godip.Land).Conn("sto", godip.Land).Conn("whi", godip.Land).Conn("big", godip.Land).Flag(godip.Land).SC(Osage).
		// Saline
		Prov("sal").Conn("laf", godip.Coast...).Conn("pet", godip.Coast...).Conn("ceo", godip.Sea).Conn("noo", godip.Sea).Conn("nwc", godip.Sea).Flag(godip.Coast...).SC(Missouria).
		// St. Charles
		Prov("stc").Conn("cpp", godip.Sea).Conn("stl", godip.Land).Conn("lin", godip.Coast...).Conn("mot", godip.Coast...).Conn("sem", godip.Sea).Conn("sem", godip.Sea).Conn("cpp", godip.Sea).Flag(godip.Coast...).SC(Illini).
		// Callaway
		Prov("cal").Conn("aud", godip.Land).Conn("mro", godip.Land).Conn("boo", godip.Coast...).Conn("ceo", godip.Sea).Conn("ecm", godip.Sea).Conn("emr", godip.Sea).Conn("mot", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Pike
		Prov("pik").Conn("lin", godip.Coast...).Conn("npp", godip.Sea).Conn("mar", godip.Coast...).Conn("aud", godip.Land).Flag(godip.Coast...).
		// Jefferson
		Prov("jef").Conn("scp", godip.Sea).Conn("stl", godip.Coast...).Conn("bou", godip.Land).Conn("ste", godip.Coast...).Conn("scp", godip.Sea).Flag(godip.Coast...).
		// Big Pinay
		Prov("big").Conn("bla", godip.Land).Conn("mer", godip.Land).Conn("gas", godip.Land).Conn("osa", godip.Land).Conn("oza", godip.Land).Conn("whi", godip.Land).Flag(godip.Land).
		// Meramec
		Prov("mer").Conn("gas", godip.Land).Conn("big", godip.Land).Conn("bla", godip.Land).Conn("bou", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).
		// Lewis
		Prov("lew").Conn("nmr", godip.Sea).Conn("lar", godip.Coast...).Conn("cui", godip.Land).Conn("mar", godip.Coast...).Flag(godip.Coast...).
		// Jackson
		Prov("jac").Conn("sog", godip.Land).Conn("cly", godip.Land).Conn("laf", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(Missouria).
		// Franklin
		Prov("fra").Conn("bou", godip.Coast...).Conn("sem", godip.Sea).Conn("sem", godip.Sea).Conn("emr", godip.Sea).Conn("gas", godip.Coast...).Conn("mer", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Central Missouri River
		Prov("ceo").Conn("noo", godip.Sea).Conn("sal", godip.Sea).Conn("pet", godip.Sea).Conn("col", godip.Sea).Conn("ecm", godip.Sea).Conn("cal", godip.Sea).Conn("boo", godip.Sea).Conn("how", godip.Sea).Flag(godip.Sea).
		// Cape Girardeau
		Prov("cap").Conn("stf", godip.Land).Conn("sco", godip.Coast...).Conn("smr", godip.Sea).Conn("per", godip.Coast...).Conn("mad", godip.Land).Conn("gre", godip.Land).Flag(godip.Coast...).
		// Cuivre
		Prov("cui").Conn("mus", godip.Land).Conn("cha", godip.Land).Conn("ran", godip.Land).Conn("mro", godip.Land).Conn("aud", godip.Land).Conn("mar", godip.Land).Conn("lew", godip.Land).Conn("lar", godip.Land).Flag(godip.Land).
		// Nodaway
		Prov("nod").Conn("nis", godip.Land).Conn("pla", godip.Land).Conn("cli", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North West Missouri River
		Prov("nwm").Conn("wes", godip.Sea).Conn("cly", godip.Sea).Conn("pla", godip.Sea).Conn("pla", godip.Sea).Conn("nis", godip.Sea).Flag(godip.Sea).
		// Lafayette
		Prov("laf").Conn("jac", godip.Coast...).Conn("sog", godip.Land).Conn("joh", godip.Land).Conn("pet", godip.Land).Conn("sal", godip.Coast...).Conn("nwc", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(Missouria).
		// Wayne
		Prov("way").Conn("mad", godip.Land).Conn("bou", godip.Land).Conn("bla", godip.Land).Conn("but", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).SC(Quapaw).
		// West Missouri River
		Prov("wes").Conn("nwm", godip.Sea).Conn("jac", godip.Sea).Conn("laf", godip.Sea).Conn("nwc", godip.Sea).Conn("ric", godip.Sea).Conn("ray", godip.Sea).Conn("cly", godip.Sea).Flag(godip.Sea).
		// St. Louis
		Prov("stl").Conn("scp", godip.Sea).Conn("stc", godip.Land).Conn("sem", godip.Sea).Conn("bou", godip.Coast...).Conn("jef", godip.Coast...).Flag(godip.Coast...).SC(Illini).
		// Musse
		Prov("mus").Conn("gra", godip.Land).Conn("car", godip.Land).Conn("cha", godip.Land).Conn("cui", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"npp": "North Central Mississippi River",
	"boo": "Boone",
	"ran": "Randolph",
	"gra": "Grand",
	"sog": "South Grand",
	"aud": "Audrain",
	"emr": "East Missouri River",
	"mro": "Monroe",
	"bla": "Black",
	"sto": "Stockton",
	"smr": "South Mississippi River",
	"bou": "Bourbeuse",
	"ver": "Vernon",
	"col": "Cole",
	"gre": "Greenville",
	"nis": "Nishnabotna",
	"new": "New Madrid",
	"sco": "Scott",
	"mot": "Montgomery",
	"whi": "White",
	"but": "Butler",
	"scp": "South Central Mississippi River",
	"gas": "Gasconade",
	"pet": "Pettis",
	"lin": "Lincoln",
	"ecm": "East Central Missouri River",
	"per": "Perry",
	"cly": "Clay",
	"ric": "Richmond",
	"nwc": "North West Central Missouri River",
	"joh": "Johnson",
	"sem": "South East Missouri River",
	"car": "Carroll",
	"how": "Howard",
	"mar": "Marion",
	"tab": "Table Rock",
	"ste": "Ste. Genevieve",
	"noo": "North Central Missouri River",
	"cpp": "Central Mississippi River",
	"pla": "Platte",
	"hen": "Henry",
	"osa": "Osage",
	"cha": "Chariton",
	"cli": "Clinton",
	"ray": "Ray",
	"nog": "North Grand",
	"stf": "St. Francis",
	"lar": "Clark",
	"neo": "Neosho",
	"nmr": "North Mississippi River",
	"mad": "Madison",
	"oza": "Ozark",
	"sal": "Saline",
	"stc": "St. Charles",
	"cal": "Callaway",
	"pik": "Pike",
	"jef": "Jefferson",
	"big": "Big Pinay",
	"mer": "Meramec",
	"lew": "Lewis",
	"jac": "Jackson",
	"fra": "Franklin",
	"ceo": "Central Missouri River",
	"cap": "Cape Girardeau",
	"cui": "Cuivre",
	"nod": "Nodaway",
	"nwm": "North West Missouri River",
	"laf": "Lafayette",
	"way": "Wayne",
	"wes": "West Missouri River",
	"stl": "St. Louis",
	"mus": "Musse",
}
