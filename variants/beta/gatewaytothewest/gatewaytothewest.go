package gatewaytothewest

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
	Oto       godip.Nation = "Oto"
	Ioway     godip.Nation = "Ioway"
	Quapaw    godip.Nation = "Quapaw"
)

var Nations = []godip.Nation{Chickasaw, Osage, Missouria, Illini, Oto, Ioway, Quapaw}

var GatewayToTheWestVariant = common.Variant{
	Name:              "GatewayToTheWest",
	Graph:             func() godip.Graph { return GatewayToTheWestGraph() },
	Start:             GatewayToTheWestStart,
	Blank:             GatewayToTheWestBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(17),
	SoloSCCount:       func(*state.State) int { return 17 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/gatewaytothewestmap.svg")
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
	Rules:       "",
}

func GatewayToTheWestBlank(phase godip.Phase) *state.State {
	return state.New(GatewayToTheWestGraph(), phase, classical.BackupRule, nil, nil)
}

func GatewayToTheWestStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(2024, godip.Spring, godip.Movement)
	result = GatewayToTheWestBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"sco": godip.Unit{godip.Army, Chickasaw},
		"new": godip.Unit{godip.Army, Chickasaw},
		"sai": godip.Unit{godip.Army, Chickasaw},
		"osa": godip.Unit{godip.Army, Osage},
		"col": godip.Unit{godip.Army, Osage},
		"oza": godip.Unit{godip.Army, Osage},
		"jac": godip.Unit{godip.Army, Missouria},
		"laf": godip.Unit{godip.Army, Missouria},
		"sal": godip.Unit{godip.Army, Missouria},
		"stc": godip.Unit{godip.Army, Illini},
		"stl": godip.Unit{godip.Army, Illini},
		"lin": godip.Unit{godip.Army, Illini},
		"pla": godip.Unit{godip.Army, Oto},
		"nis": godip.Unit{godip.Army, Oto},
		"nod": godip.Unit{godip.Army, Oto},
		"gra": godip.Unit{godip.Army, Ioway},
		"ngr": godip.Unit{godip.Army, Ioway},
		"cli": godip.Unit{godip.Army, Ioway},
		"way": godip.Unit{godip.Army, Quapaw},
		"but": godip.Unit{godip.Army, Quapaw},
		"bla": godip.Unit{godip.Army, Quapaw},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"sco": Chickasaw,
		"new": Chickasaw,
		"sai": Chickasaw,
		"osa": Osage,
		"col": Osage,
		"oza": Osage,
		"jac": Missouria,
		"laf": Missouria,
		"sal": Missouria,
		"stc": Illini,
		"stl": Illini,
		"lin": Illini,
		"pla": Oto,
		"nis": Oto,
		"nod": Oto,
		"gra": Ioway,
		"ngr": Ioway,
		"cli": Ioway,
		"way": Quapaw,
		"but": Quapaw,
		"bla": Quapaw,
	})
	return
}

func GatewayToTheWestGraph() *graph.Graph {
	return graph.New().
		// Table Rock
		Prov("tab").Conn("whi", godip.Land).Conn("sto", godip.Land).Conn("neo", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North Central Mississippi River
		Prov("ncm").Conn("cen", godip.Sea).Conn("nmr", godip.Sea).Conn("mar", godip.Sea).Conn("pik", godip.Sea).Conn("lin", godip.Sea).Flag(godip.Sea).
		// Boone
		Prov("boo").Conn("how", godip.Coast...).Conn("bnv", godip.Sea).Conn("cal", godip.Coast...).Conn("mor", godip.Land).Conn("ran", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Randolph
		Prov("ran").Conn("gla", godip.Sea).Conn("how", godip.Coast...).Conn("boo", godip.Land).Conn("mor", godip.Land).Conn("cui", godip.Land).Conn("cha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Grand
		Prov("gra").Conn("ngr", godip.Land).Conn("cli", godip.Land).Conn("ray", godip.Land).Conn("tes", godip.Land).Conn("car", godip.Land).Conn("mus", godip.Land).Flag(godip.Land).SC(Ioway).
		// South Grand
		Prov("sog").Conn("joh", godip.Land).Conn("laf", godip.Land).Conn("jac", godip.Land).Conn("ver", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// Herman
		Prov("her").Conn("fra", godip.Sea).Conn("eas", godip.Sea).Conn("stc", godip.Sea).Conn("mot", godip.Sea).Conn("cal", godip.Sea).Conn("jec", godip.Sea).Conn("gas", godip.Sea).Flag(godip.Sea).
		// Miami
		Prov("mia").Conn("car", godip.Sea).Conn("tes", godip.Sea).Conn("wes", godip.Sea).Conn("laf", godip.Sea).Conn("sal", godip.Sea).Conn("gla", godip.Sea).Conn("cha", godip.Sea).Flag(godip.Sea).
		// East Missouri River
		Prov("eas").Conn("stl", godip.Sea).Conn("scm", godip.Sea).Conn("cen", godip.Sea).Conn("stc", godip.Sea).Conn("her", godip.Sea).Conn("fra", godip.Sea).Conn("bou", godip.Sea).Flag(godip.Sea).
		// Monroe
		Prov("mor").Conn("aud", godip.Land).Conn("cui", godip.Land).Conn("ran", godip.Land).Conn("boo", godip.Land).Conn("cal", godip.Land).Flag(godip.Land).
		// Black
		Prov("bla").Conn("way", godip.Land).Conn("bou", godip.Land).Conn("mer", godip.Land).Conn("big", godip.Land).Conn("whi", godip.Land).Conn("but", godip.Land).Flag(godip.Land).SC(Quapaw).
		// Stockton
		Prov("sto").Conn("oza", godip.Land).Conn("ver", godip.Land).Conn("neo", godip.Land).Conn("tab", godip.Land).Conn("whi", godip.Land).Flag(godip.Land).
		// N. Grand
		Prov("ngr").Conn("gra", godip.Land).Conn("cli", godip.Land).Flag(godip.Land).SC(Ioway).
		// South Mississippi River
		Prov("smr").Conn("scm", godip.Sea).Conn("ste", godip.Sea).Conn("per", godip.Sea).Conn("cap", godip.Sea).Conn("sco", godip.Sea).Conn("sai", godip.Sea).Flag(godip.Sea).
		// Bourbeuse
		Prov("bou").Conn("bla", godip.Land).Conn("way", godip.Land).Conn("mad", godip.Land).Conn("ste", godip.Land).Conn("jee", godip.Land).Conn("stl", godip.Coast...).Conn("eas", godip.Sea).Conn("fra", godip.Coast...).Conn("mer", godip.Land).Flag(godip.Coast...).
		// Vernon
		Prov("ver").Conn("neo", godip.Land).Conn("sto", godip.Land).Conn("oza", godip.Land).Conn("hen", godip.Land).Conn("sog", godip.Land).Flag(godip.Land).
		// Marion
		Prov("mar").Conn("lew", godip.Coast...).Conn("cui", godip.Land).Conn("aud", godip.Land).Conn("pik", godip.Coast...).Conn("ncm", godip.Sea).Conn("nmr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cole
		Prov("col").Conn("bnv", godip.Sea).Conn("pet", godip.Coast...).Conn("oza", godip.Land).Conn("osa", godip.Coast...).Conn("jec", godip.Sea).Flag(godip.Coast...).SC(Osage).
		// Greenville
		Prov("gre").Conn("cap", godip.Land).Conn("mad", godip.Land).Conn("way", godip.Land).Conn("but", godip.Land).Conn("new", godip.Land).Flag(godip.Land).
		// Test
		Prov("tes").Conn("ray", godip.Coast...).Conn("wes", godip.Sea).Conn("mia", godip.Sea).Conn("car", godip.Coast...).Conn("gra", godip.Land).Flag(godip.Coast...).
		// New Madrid
		Prov("new").Conn("sai", godip.Land).Conn("sco", godip.Land).Conn("cap", godip.Land).Conn("gre", godip.Land).Conn("but", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Scott
		Prov("sco").Conn("cap", godip.Coast...).Conn("new", godip.Land).Conn("sai", godip.Coast...).Conn("smr", godip.Sea).Flag(godip.Coast...).SC(Chickasaw).
		// Montgomery
		Prov("mot").Conn("cal", godip.Coast...).Conn("her", godip.Sea).Conn("stc", godip.Coast...).Conn("lin", godip.Land).Conn("aud", godip.Land).Flag(godip.Coast...).
		// White
		Prov("whi").Conn("oza", godip.Land).Conn("sto", godip.Land).Conn("tab", godip.Land).Conn("bla", godip.Land).Conn("big", godip.Land).Flag(godip.Land).
		// Butler
		Prov("but").Conn("new", godip.Land).Conn("gre", godip.Land).Conn("way", godip.Land).Conn("bla", godip.Land).Flag(godip.Land).SC(Quapaw).
		// South Central Mississippi River
		Prov("scm").Conn("cen", godip.Sea).Conn("eas", godip.Sea).Conn("stl", godip.Sea).Conn("jee", godip.Sea).Conn("jee", godip.Sea).Conn("ste", godip.Sea).Conn("smr", godip.Sea).Flag(godip.Sea).
		// Gasconade
		Prov("gas").Conn("jec", godip.Sea).Conn("osa", godip.Coast...).Conn("big", godip.Land).Conn("mer", godip.Land).Conn("fra", godip.Coast...).Conn("her", godip.Sea).Flag(godip.Coast...).
		// Pettis
		Prov("pet").Conn("oza", godip.Land).Conn("col", godip.Coast...).Conn("bnv", godip.Sea).Conn("sal", godip.Coast...).Conn("laf", godip.Land).Conn("joh", godip.Land).Conn("hen", godip.Land).Flag(godip.Coast...).
		// Lincoln
		Prov("lin").Conn("aud", godip.Land).Conn("mot", godip.Land).Conn("stc", godip.Coast...).Conn("cen", godip.Sea).Conn("ncm", godip.Sea).Conn("pik", godip.Coast...).Flag(godip.Coast...).SC(Illini).
		// West Missouri River
		Prov("wes").Conn("cly", godip.Sea).Conn("nwm", godip.Sea).Conn("jac", godip.Sea).Conn("laf", godip.Sea).Conn("mia", godip.Sea).Conn("tes", godip.Sea).Conn("ray", godip.Sea).Flag(godip.Sea).
		// Saint Francis
		Prov("sai").Conn("smr", godip.Sea).Conn("sco", godip.Coast...).Conn("new", godip.Land).Flag(godip.Coast...).SC(Chickasaw).
		// Perry
		Prov("per").Conn("cap", godip.Coast...).Conn("smr", godip.Sea).Conn("ste", godip.Coast...).Conn("mad", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Audrain
		Prov("aud").Conn("pik", godip.Land).Conn("mar", godip.Land).Conn("cui", godip.Land).Conn("mor", godip.Land).Conn("cal", godip.Land).Conn("mot", godip.Land).Conn("lin", godip.Land).Flag(godip.Land).
		// Cuivre
		Prov("cui").Conn("clr", godip.Land).Conn("mus", godip.Land).Conn("cha", godip.Land).Conn("ran", godip.Land).Conn("mor", godip.Land).Conn("aud", godip.Land).Conn("mar", godip.Land).Conn("lew", godip.Land).Flag(godip.Land).
		// Johnson
		Prov("joh").Conn("sog", godip.Land).Conn("hen", godip.Land).Conn("pet", godip.Land).Conn("laf", godip.Land).Flag(godip.Land).
		// Carroll
		Prov("car").Conn("mia", godip.Sea).Conn("cha", godip.Coast...).Conn("mus", godip.Land).Conn("gra", godip.Land).Conn("tes", godip.Coast...).Flag(godip.Coast...).
		// Howard
		Prov("how").Conn("boo", godip.Coast...).Conn("ran", godip.Coast...).Conn("gla", godip.Sea).Conn("bnv", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Nishnabotna
		Prov("nis").Conn("nod", godip.Land).Conn("nwm", godip.Sea).Conn("pla", godip.Coast...).Flag(godip.Coast...).SC(Oto).
		// Ste. Genevieve
		Prov("ste").Conn("mad", godip.Land).Conn("per", godip.Coast...).Conn("smr", godip.Sea).Conn("scm", godip.Sea).Conn("jee", godip.Coast...).Conn("bou", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Central Mississippi River
		Prov("cen").Conn("stc", godip.Sea).Conn("stc", godip.Sea).Conn("eas", godip.Sea).Conn("scm", godip.Sea).Conn("ncm", godip.Sea).Conn("lin", godip.Sea).Flag(godip.Sea).
		// Platte
		Prov("pla").Conn("nwm", godip.Sea).Conn("nwm", godip.Sea).Conn("cly", godip.Coast...).Conn("cli", godip.Land).Conn("nod", godip.Land).Conn("nis", godip.Coast...).Flag(godip.Coast...).SC(Oto).
		// Henry
		Prov("hen").Conn("pet", godip.Land).Conn("joh", godip.Land).Conn("sog", godip.Land).Conn("ver", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("oza", godip.Land).Conn("big", godip.Land).Conn("gas", godip.Coast...).Conn("jec", godip.Sea).Conn("col", godip.Coast...).Flag(godip.Coast...).SC(Osage).
		// Chariton
		Prov("cha").Conn("mus", godip.Land).Conn("car", godip.Coast...).Conn("mia", godip.Sea).Conn("gla", godip.Sea).Conn("ran", godip.Coast...).Conn("cui", godip.Land).Flag(godip.Coast...).
		// Clinton
		Prov("cli").Conn("nod", godip.Land).Conn("pla", godip.Land).Conn("cly", godip.Land).Conn("ray", godip.Land).Conn("gra", godip.Land).Conn("ngr", godip.Land).Flag(godip.Land).SC(Ioway).
		// Ray
		Prov("ray").Conn("cly", godip.Coast...).Conn("wes", godip.Sea).Conn("tes", godip.Coast...).Conn("gra", godip.Land).Conn("cli", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Booneville
		Prov("bnv").Conn("sal", godip.Sea).Conn("pet", godip.Sea).Conn("col", godip.Sea).Conn("jec", godip.Sea).Conn("cal", godip.Sea).Conn("boo", godip.Sea).Conn("how", godip.Sea).Conn("gla", godip.Sea).Flag(godip.Sea).
		// Clark
		Prov("clr").Conn("cui", godip.Land).Conn("lew", godip.Coast...).Conn("nmr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Neosho
		Prov("neo").Conn("ver", godip.Land).Conn("tab", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North Mississippi River
		Prov("nmr").Conn("clr", godip.Sea).Conn("lew", godip.Sea).Conn("mar", godip.Sea).Conn("ncm", godip.Sea).Flag(godip.Sea).
		// Madison
		Prov("mad").Conn("cap", godip.Land).Conn("per", godip.Land).Conn("ste", godip.Land).Conn("bou", godip.Land).Conn("way", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).
		// Ozark
		Prov("oza").Conn("whi", godip.Land).Conn("big", godip.Land).Conn("osa", godip.Land).Conn("col", godip.Land).Conn("pet", godip.Land).Conn("hen", godip.Land).Conn("ver", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).SC(Osage).
		// Saline
		Prov("sal").Conn("bnv", godip.Sea).Conn("gla", godip.Sea).Conn("mia", godip.Sea).Conn("laf", godip.Coast...).Conn("pet", godip.Coast...).Flag(godip.Coast...).SC(Missouria).
		// St. Charles
		Prov("stc").Conn("cen", godip.Sea).Conn("lin", godip.Coast...).Conn("mot", godip.Coast...).Conn("her", godip.Sea).Conn("eas", godip.Sea).Conn("cen", godip.Sea).Flag(godip.Coast...).SC(Illini).
		// Callaway
		Prov("cal").Conn("her", godip.Sea).Conn("mot", godip.Coast...).Conn("aud", godip.Land).Conn("mor", godip.Land).Conn("boo", godip.Coast...).Conn("bnv", godip.Sea).Conn("jec", godip.Sea).Flag(godip.Coast...).
		// Pike
		Prov("pik").Conn("aud", godip.Land).Conn("lin", godip.Coast...).Conn("ncm", godip.Sea).Conn("mar", godip.Coast...).Flag(godip.Coast...).
		// Jefferson City
		Prov("jec").Conn("gas", godip.Sea).Conn("her", godip.Sea).Conn("cal", godip.Sea).Conn("bnv", godip.Sea).Conn("col", godip.Sea).Conn("osa", godip.Sea).Flag(godip.Sea).
		// Jefferson
		Prov("jee").Conn("scm", godip.Sea).Conn("scm", godip.Sea).Conn("stl", godip.Coast...).Conn("bou", godip.Land).Conn("ste", godip.Coast...).Flag(godip.Coast...).
		// Glasgow
		Prov("gla").Conn("mia", godip.Sea).Conn("sal", godip.Sea).Conn("bnv", godip.Sea).Conn("how", godip.Sea).Conn("ran", godip.Sea).Conn("cha", godip.Sea).Flag(godip.Sea).
		// Big Pinay
		Prov("big").Conn("oza", godip.Land).Conn("whi", godip.Land).Conn("bla", godip.Land).Conn("mer", godip.Land).Conn("gas", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Meramec
		Prov("mer").Conn("fra", godip.Land).Conn("gas", godip.Land).Conn("big", godip.Land).Conn("bla", godip.Land).Conn("bou", godip.Land).Flag(godip.Land).
		// Lewis
		Prov("lew").Conn("mar", godip.Coast...).Conn("nmr", godip.Sea).Conn("clr", godip.Coast...).Conn("cui", godip.Land).Flag(godip.Coast...).
		// Jackson
		Prov("jac").Conn("sog", godip.Land).Conn("laf", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(Missouria).
		// Franklin
		Prov("fra").Conn("mer", godip.Land).Conn("bou", godip.Coast...).Conn("eas", godip.Sea).Conn("her", godip.Sea).Conn("gas", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Clay
		Prov("cly").Conn("ray", godip.Coast...).Conn("cli", godip.Land).Conn("pla", godip.Coast...).Conn("nwm", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Cape Girardeau
		Prov("cap").Conn("mad", godip.Land).Conn("gre", godip.Land).Conn("new", godip.Land).Conn("sco", godip.Coast...).Conn("smr", godip.Sea).Conn("per", godip.Coast...).Flag(godip.Coast...).
		// Nodaway
		Prov("nod").Conn("nis", godip.Land).Conn("pla", godip.Land).Conn("cli", godip.Land).Flag(godip.Land).SC(Oto).
		// North West Missouri River
		Prov("nwm").Conn("pla", godip.Sea).Conn("nis", godip.Sea).Conn("wes", godip.Sea).Conn("cly", godip.Sea).Conn("pla", godip.Sea).Flag(godip.Sea).
		// Lafayette
		Prov("laf").Conn("mia", godip.Sea).Conn("wes", godip.Sea).Conn("jac", godip.Coast...).Conn("sog", godip.Land).Conn("joh", godip.Land).Conn("pet", godip.Land).Conn("sal", godip.Coast...).Flag(godip.Coast...).SC(Missouria).
		// Wayne
		Prov("way").Conn("bla", godip.Land).Conn("but", godip.Land).Conn("gre", godip.Land).Conn("mad", godip.Land).Conn("bou", godip.Land).Flag(godip.Land).SC(Quapaw).
		// St. Louis
		Prov("stl").Conn("eas", godip.Sea).Conn("bou", godip.Coast...).Conn("jee", godip.Coast...).Conn("scm", godip.Sea).Flag(godip.Coast...).SC(Illini).
		// Musse
		Prov("mus").Conn("cha", godip.Land).Conn("cui", godip.Land).Conn("gra", godip.Land).Conn("car", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"tab": "Table Rock",
	"ncm": "North Central Mississippi River",
	"boo": "Boone",
	"ran": "Randolph",
	"gra": "Grand",
	"sog": "South Grand",
	"her": "Herman",
	"mia": "Miami",
	"eas": "East Missouri River",
	"mor": "Monroe",
	"bla": "Black",
	"sto": "Stockton",
	"ngr": "N. Grand",
	"smr": "South Mississippi River",
	"bou": "Bourbeuse",
	"ver": "Vernon",
	"mar": "Marion",
	"col": "Cole",
	"gre": "Greenville",
	"tes": "Test",
	"new": "New Madrid",
	"sco": "Scott",
	"mot": "Montgomery",
	"whi": "White",
	"but": "Butler",
	"scm": "South Central Mississippi River",
	"gas": "Gasconade",
	"pet": "Pettis",
	"lin": "Lincoln",
	"wes": "West Missouri River",
	"sai": "Saint Francis",
	"per": "Perry",
	"aud": "Audrain",
	"cui": "Cuivre",
	"joh": "Johnson",
	"car": "Carroll",
	"how": "Howard",
	"nis": "Nishnabotna",
	"ste": "Ste. Genevieve",
	"cen": "Central Mississippi River",
	"pla": "Platte",
	"hen": "Henry",
	"osa": "Osage",
	"cha": "Chariton",
	"cli": "Clinton",
	"ray": "Ray",
	"bnv": "Booneville",
	"clr": "Clark",
	"neo": "Neosho",
	"nmr": "North Mississippi River",
	"mad": "Madison",
	"oza": "Ozark",
	"sal": "Saline",
	"stc": "St. Charles",
	"cal": "Callaway",
	"pik": "Pike",
	"jec": "Jefferson City",
	"jee": "Jefferson",
	"gla": "Glasgow",
	"big": "Big Pinay",
	"mer": "Meramec",
	"lew": "Lewis",
	"jac": "Jackson",
	"fra": "Franklin",
	"cly": "Clay",
	"cap": "Cape Girardeau",
	"nod": "Nodaway",
	"nwm": "North West Missouri River",
	"laf": "Lafayette",
	"way": "Wayne",
	"stl": "St. Louis",
	"mus": "Musse",
}
