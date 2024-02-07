package missouri

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Missouria godip.Nation = "Missouria"
	Oto       godip.Nation = "Oto"
)

var Nations = []godip.Nation{Missouria, Oto}

var MissouriVariant = common.Variant{
	Name:              "Missouri",
	Graph:             func() godip.Graph { return MissouriGraph() },
	Start:             MissouriStart,
	Blank:             MissouriBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(2),
	SoloSCCount:       func(*state.State) int { return 2 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/missourimap.svg")
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

func MissouriBlank(phase godip.Phase) *state.State {
	return state.New(MissouriGraph(), phase, classical.BackupRule, nil, nil)
}

func MissouriStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(2024, godip.Spring, godip.Movement)
	result = MissouriBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"and": godip.Unit{godip.Army, Missouria},
		"atc": godip.Unit{godip.Army, Oto},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"and": Missouria,
		"atc": Oto,
	})
	return
}

func MissouriGraph() *graph.Graph {
	return graph.New().
		// Polk
		Prov("pol").Conn("CLR", godip.Land).Conn("ced", godip.Land).Conn("dad", godip.Land).Conn("gre", godip.Land).Conn("dal", godip.Land).Conn("hic", godip.Land).Flag(godip.Land).
		// Phelps
		Prov("phe").Conn("den", godip.Land).Conn("cra", godip.Land).Conn("gas", godip.Land).Conn("mae", godip.Land).Conn("pul", godip.Land).Conn("tex", godip.Land).Flag(godip.Land).
		// Barton
		Prov("bao").Conn("jas", godip.Land).Conn("dad", godip.Land).Conn("ced", godip.Land).Conn("ver", godip.Land).Flag(godip.Land).
		// Camden
		Prov("cam").Conn("ben", godip.Land).Conn("hic", godip.Land).Conn("dal", godip.Land).Conn("lac", godip.Land).Conn("pul", godip.Land).Conn("mil", godip.Land).Conn("mor", godip.Land).Flag(godip.Land).
		// Montgomery
		Prov("mog").Conn("lic", godip.Land).Conn("pik", godip.Land).Conn("aud", godip.Land).Conn("cll", godip.Land).Conn("gas", godip.Land).Conn("war", godip.Land).Flag(godip.Land).
		// Gasconade
		Prov("gas").Conn("fra", godip.Land).Conn("war", godip.Land).Conn("mog", godip.Land).Conn("osa", godip.Land).Conn("mae", godip.Land).Conn("phe", godip.Land).Conn("cra", godip.Land).Flag(godip.Land).
		// Holt
		Prov("hol").Conn("atc", godip.Land).Conn("and", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).
		// Shelby
		Prov("she").Conn("mao", godip.Land).Conn("lew", godip.Land).Conn("kno", godip.Land).Conn("mac", godip.Land).Conn("mnr", godip.Land).Flag(godip.Land).
		// St. Louis City
		Prov("slc").Conn("slo", godip.Land).Conn("slo", godip.Land).Flag(godip.Land).
		// Audrain
		Prov("aud").Conn("boo", godip.Land).Conn("cll", godip.Land).Conn("mog", godip.Land).Conn("pik", godip.Land).Conn("ral", godip.Land).Conn("mnr", godip.Land).Conn("ran", godip.Land).Flag(godip.Land).
		// Clark
		Prov("clr").Conn("sco", godip.Land).Conn("kno", godip.Land).Conn("lew", godip.Land).Flag(godip.Land).
		// Scotland
		Prov("sco").Conn("clr", godip.Land).Conn("sch", godip.Land).Conn("ada", godip.Land).Conn("kno", godip.Land).Flag(godip.Land).
		// Douglas
		Prov("dou").Conn("oza", godip.Land).Conn("hoe", godip.Land).Conn("tex", godip.Land).Conn("wri", godip.Land).Conn("web", godip.Land).Conn("chr", godip.Land).Conn("tan", godip.Land).Flag(godip.Land).
		// Dent
		Prov("den").Conn("phe", godip.Land).Conn("tex", godip.Land).Conn("sha", godip.Land).Conn("jef", godip.Land).Conn("jef", godip.Land).Conn("cra", godip.Land).Flag(godip.Land).
		// Jasper
		Prov("jas").Conn("dad", godip.Land).Conn("bao", godip.Land).Conn("new", godip.Land).Conn("law", godip.Land).Flag(godip.Land).
		// Cooper
		Prov("coo").Conn("sal", godip.Land).Conn("pet", godip.Land).Conn("mor", godip.Land).Conn("moi", godip.Land).Conn("boo", godip.Land).Conn("hoa", godip.Land).Flag(godip.Land).
		// Jefferson
		Prov("jef").Conn("ore", godip.Land).Conn("ore", godip.Land).Conn("slo", godip.Land).Conn("fra", godip.Land).Conn("fra", godip.Land).Conn("cra", godip.Land).Conn("cra", godip.Land).Conn("den", godip.Land).Conn("den", godip.Land).Conn("sha", godip.Land).Conn("sha", godip.Land).Flag(godip.Land).
		// St. Louis
		Prov("slo").Conn("slc", godip.Land).Conn("stc", godip.Land).Conn("fra", godip.Land).Conn("jef", godip.Land).Conn("slc", godip.Land).Flag(godip.Land).
		// Laclede
		Prov("lac").Conn("web", godip.Land).Conn("wri", godip.Land).Conn("tex", godip.Land).Conn("pul", godip.Land).Conn("cam", godip.Land).Conn("dal", godip.Land).Flag(godip.Land).
		// Randolph
		Prov("ran").Conn("cha", godip.Land).Conn("hoa", godip.Land).Conn("boo", godip.Land).Conn("aud", godip.Land).Conn("mnr", godip.Land).Conn("mac", godip.Land).Flag(godip.Land).
		// Dade
		Prov("dad").Conn("gre", godip.Land).Conn("pol", godip.Land).Conn("ced", godip.Land).Conn("bao", godip.Land).Conn("jas", godip.Land).Conn("law", godip.Land).Flag(godip.Land).
		// Barry
		Prov("bay").Conn("sto", godip.Land).Conn("law", godip.Land).Conn("new", godip.Land).Conn("mcd", godip.Land).Flag(godip.Land).
		// Atchison
		Prov("atc").Conn("hol", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).SC(Oto).
		// Gentry
		Prov("gen").Conn("and", godip.Land).Conn("dek", godip.Land).Conn("dav", godip.Land).Conn("har", godip.Land).Conn("wor", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).
		// Macon
		Prov("mac").Conn("cha", godip.Land).Conn("ran", godip.Land).Conn("she", godip.Land).Conn("kno", godip.Land).Conn("ada", godip.Land).Conn("lnn", godip.Land).Flag(godip.Land).
		// Pulaski
		Prov("pul").Conn("mil", godip.Land).Conn("cam", godip.Land).Conn("lac", godip.Land).Conn("tex", godip.Land).Conn("phe", godip.Land).Conn("mae", godip.Land).Flag(godip.Land).
		// Johnson
		Prov("joh").Conn("pet", godip.Land).Conn("laf", godip.Land).Conn("jac", godip.Land).Conn("cas", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// Marion
		Prov("mao").Conn("she", godip.Land).Conn("mnr", godip.Land).Conn("ral", godip.Land).Conn("lew", godip.Land).Flag(godip.Land).
		// Knox
		Prov("kno").Conn("clr", godip.Land).Conn("sco", godip.Land).Conn("ada", godip.Land).Conn("mac", godip.Land).Conn("she", godip.Land).Conn("lew", godip.Land).Flag(godip.Land).
		// Cole
		Prov("col").Conn("osa", godip.Land).Conn("cll", godip.Land).Conn("boo", godip.Land).Conn("moi", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).
		// Clay
		Prov("cly").Conn("jac", godip.Land).Conn("ray", godip.Land).Conn("cli", godip.Land).Conn("pla", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("mae", godip.Land).Conn("gas", godip.Land).Conn("cll", godip.Land).Conn("col", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).
		// De Kalb
		Prov("dek").Conn("cad", godip.Land).Conn("dav", godip.Land).Conn("gen", godip.Land).Conn("and", godip.Land).Conn("buc", godip.Land).Conn("cli", godip.Land).Flag(godip.Land).
		// Mcdonald
		Prov("mcd").Conn("new", godip.Land).Conn("bay", godip.Land).Flag(godip.Land).
		// Ray
		Prov("ray").Conn("laf", godip.Land).Conn("car", godip.Land).Conn("cad", godip.Land).Conn("cli", godip.Land).Conn("cly", godip.Land).Conn("jac", godip.Land).Flag(godip.Land).
		// Lawrence
		Prov("law").Conn("dad", godip.Land).Conn("jas", godip.Land).Conn("new", godip.Land).Conn("bay", godip.Land).Conn("sto", godip.Land).Conn("chr", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).
		// Ralls
		Prov("ral").Conn("mao", godip.Land).Conn("mnr", godip.Land).Conn("aud", godip.Land).Conn("pik", godip.Land).Flag(godip.Land).
		// Maries
		Prov("mae").Conn("osa", godip.Land).Conn("mil", godip.Land).Conn("pul", godip.Land).Conn("phe", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).
		// Warren
		Prov("war").Conn("lic", godip.Land).Conn("mog", godip.Land).Conn("gas", godip.Land).Conn("fra", godip.Land).Conn("stc", godip.Land).Flag(godip.Land).
		// Callaway
		Prov("cll").Conn("osa", godip.Land).Conn("mog", godip.Land).Conn("aud", godip.Land).Conn("boo", godip.Land).Conn("col", godip.Land).Flag(godip.Land).
		// Pike
		Prov("pik").Conn("aud", godip.Land).Conn("mog", godip.Land).Conn("lic", godip.Land).Conn("ral", godip.Land).Flag(godip.Land).
		// Daviess
		Prov("dav").Conn("gen", godip.Land).Conn("dek", godip.Land).Conn("cad", godip.Land).Conn("liv", godip.Land).Conn("gru", godip.Land).Conn("har", godip.Land).Flag(godip.Land).
		// Jackson
		Prov("jac").Conn("cly", godip.Land).Conn("cas", godip.Land).Conn("joh", godip.Land).Conn("laf", godip.Land).Conn("ray", godip.Land).Flag(godip.Land).
		// Greene
		Prov("gre").Conn("dad", godip.Land).Conn("law", godip.Land).Conn("chr", godip.Land).Conn("web", godip.Land).Conn("dal", godip.Land).Conn("pol", godip.Land).Flag(godip.Land).
		// Nodaway
		Prov("nod").Conn("atc", godip.Land).Conn("hol", godip.Land).Conn("and", godip.Land).Conn("gen", godip.Land).Conn("wor", godip.Land).Flag(godip.Land).
		// Oregon
		Prov("ore").Conn("jef", godip.Land).Conn("sha", godip.Land).Conn("hoe", godip.Land).Conn("jef", godip.Land).Flag(godip.Land).
		// Lafayette
		Prov("laf").Conn("ray", godip.Land).Conn("jac", godip.Land).Conn("joh", godip.Land).Conn("pet", godip.Land).Conn("sal", godip.Land).Conn("car", godip.Land).Flag(godip.Land).
		// St. Clair
		Prov("CLR").Conn("pol", godip.Land).Conn("hic", godip.Land).Conn("ben", godip.Land).Conn("hen", godip.Land).Conn("bat", godip.Land).Conn("ver", godip.Land).Conn("ced", godip.Land).Flag(godip.Land).
		// Mercer
		Prov("mer").Conn("gru", godip.Land).Conn("sul", godip.Land).Conn("put", godip.Land).Conn("har", godip.Land).Flag(godip.Land).
		// Monroe
		Prov("mnr").Conn("aud", godip.Land).Conn("ral", godip.Land).Conn("mao", godip.Land).Conn("she", godip.Land).Conn("ran", godip.Land).Flag(godip.Land).
		// Moniteau
		Prov("moi").Conn("coo", godip.Land).Conn("mor", godip.Land).Conn("mil", godip.Land).Conn("col", godip.Land).Conn("boo", godip.Land).Flag(godip.Land).
		// Benton
		Prov("ben").Conn("cam", godip.Land).Conn("mor", godip.Land).Conn("pet", godip.Land).Conn("hen", godip.Land).Conn("CLR", godip.Land).Conn("hic", godip.Land).Flag(godip.Land).
		// Boone
		Prov("boo").Conn("aud", godip.Land).Conn("ran", godip.Land).Conn("hoa", godip.Land).Conn("coo", godip.Land).Conn("moi", godip.Land).Conn("col", godip.Land).Conn("cll", godip.Land).Flag(godip.Land).
		// Pettis
		Prov("pet").Conn("ben", godip.Land).Conn("mor", godip.Land).Conn("coo", godip.Land).Conn("sal", godip.Land).Conn("laf", godip.Land).Conn("joh", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// Lincoln
		Prov("lic").Conn("stc", godip.Land).Conn("pik", godip.Land).Conn("mog", godip.Land).Conn("war", godip.Land).Flag(godip.Land).
		// Sullivan
		Prov("sul").Conn("put", godip.Land).Conn("mer", godip.Land).Conn("gru", godip.Land).Conn("lnn", godip.Land).Conn("ada", godip.Land).Flag(godip.Land).
		// Bates
		Prov("bat").Conn("cas", godip.Land).Conn("ver", godip.Land).Conn("CLR", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// Hickory
		Prov("hic").Conn("dal", godip.Land).Conn("cam", godip.Land).Conn("ben", godip.Land).Conn("CLR", godip.Land).Conn("pol", godip.Land).Flag(godip.Land).
		// Chariton
		Prov("cha").Conn("ran", godip.Land).Conn("mac", godip.Land).Conn("lnn", godip.Land).Conn("liv", godip.Land).Conn("car", godip.Land).Conn("sal", godip.Land).Conn("hoa", godip.Land).Flag(godip.Land).
		// Dallas
		Prov("dal").Conn("hic", godip.Land).Conn("pol", godip.Land).Conn("gre", godip.Land).Conn("web", godip.Land).Conn("lac", godip.Land).Conn("cam", godip.Land).Flag(godip.Land).
		// Ozark
		Prov("oza").Conn("dou", godip.Land).Conn("tan", godip.Land).Conn("hoe", godip.Land).Flag(godip.Land).
		// Livingston
		Prov("liv").Conn("lnn", godip.Land).Conn("gru", godip.Land).Conn("dav", godip.Land).Conn("cad", godip.Land).Conn("car", godip.Land).Conn("cha", godip.Land).Flag(godip.Land).
		// Howell
		Prov("hoe").Conn("ore", godip.Land).Conn("sha", godip.Land).Conn("tex", godip.Land).Conn("dou", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// Taney
		Prov("tan").Conn("dou", godip.Land).Conn("chr", godip.Land).Conn("sto", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).
		// Platte
		Prov("pla").Conn("cly", godip.Land).Conn("cli", godip.Land).Conn("buc", godip.Land).Flag(godip.Land).
		// Henry
		Prov("hen").Conn("joh", godip.Land).Conn("cas", godip.Land).Conn("bat", godip.Land).Conn("CLR", godip.Land).Conn("ben", godip.Land).Conn("pet", godip.Land).Flag(godip.Land).
		// Carroll
		Prov("car").Conn("sal", godip.Land).Conn("cha", godip.Land).Conn("liv", godip.Land).Conn("cad", godip.Land).Conn("ray", godip.Land).Conn("laf", godip.Land).Flag(godip.Land).
		// Vernon
		Prov("ver").Conn("bao", godip.Land).Conn("ced", godip.Land).Conn("CLR", godip.Land).Conn("bat", godip.Land).Flag(godip.Land).
		// Grundy
		Prov("gru").Conn("liv", godip.Land).Conn("lnn", godip.Land).Conn("sul", godip.Land).Conn("mer", godip.Land).Conn("har", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).
		// Saline
		Prov("sal").Conn("car", godip.Land).Conn("laf", godip.Land).Conn("pet", godip.Land).Conn("coo", godip.Land).Conn("hoa", godip.Land).Conn("cha", godip.Land).Flag(godip.Land).
		// St. Charles
		Prov("stc").Conn("lic", godip.Land).Conn("war", godip.Land).Conn("fra", godip.Land).Conn("slo", godip.Land).Flag(godip.Land).
		// Worth
		Prov("wor").Conn("har", godip.Land).Conn("nod", godip.Land).Conn("gen", godip.Land).Flag(godip.Land).
		// Schuyler
		Prov("sch").Conn("put", godip.Land).Conn("ada", godip.Land).Conn("sco", godip.Land).Flag(godip.Land).
		// Cass
		Prov("cas").Conn("bat", godip.Land).Conn("hen", godip.Land).Conn("joh", godip.Land).Conn("jac", godip.Land).Flag(godip.Land).
		// Adair
		Prov("ada").Conn("sul", godip.Land).Conn("mac", godip.Land).Conn("kno", godip.Land).Conn("sco", godip.Land).Conn("sch", godip.Land).Conn("put", godip.Land).Flag(godip.Land).
		// Shannon
		Prov("sha").Conn("tex", godip.Land).Conn("hoe", godip.Land).Conn("ore", godip.Land).Conn("jef", godip.Land).Conn("jef", godip.Land).Conn("den", godip.Land).Flag(godip.Land).
		// Christian
		Prov("chr").Conn("tan", godip.Land).Conn("dou", godip.Land).Conn("web", godip.Land).Conn("gre", godip.Land).Conn("law", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).
		// Miller
		Prov("mil").Conn("pul", godip.Land).Conn("mae", godip.Land).Conn("osa", godip.Land).Conn("col", godip.Land).Conn("moi", godip.Land).Conn("mor", godip.Land).Conn("cam", godip.Land).Flag(godip.Land).
		// Linn
		Prov("lnn").Conn("liv", godip.Land).Conn("cha", godip.Land).Conn("mac", godip.Land).Conn("sul", godip.Land).Conn("gru", godip.Land).Flag(godip.Land).
		// Webster
		Prov("web").Conn("lac", godip.Land).Conn("dal", godip.Land).Conn("gre", godip.Land).Conn("chr", godip.Land).Conn("dou", godip.Land).Conn("wri", godip.Land).Flag(godip.Land).
		// Harrison
		Prov("har").Conn("wor", godip.Land).Conn("gen", godip.Land).Conn("dav", godip.Land).Conn("gru", godip.Land).Conn("mer", godip.Land).Flag(godip.Land).
		// Crawford
		Prov("cra").Conn("jef", godip.Land).Conn("fra", godip.Land).Conn("gas", godip.Land).Conn("phe", godip.Land).Conn("den", godip.Land).Conn("jef", godip.Land).Flag(godip.Land).
		// Morgan
		Prov("mor").Conn("mil", godip.Land).Conn("moi", godip.Land).Conn("coo", godip.Land).Conn("pet", godip.Land).Conn("ben", godip.Land).Conn("cam", godip.Land).Flag(godip.Land).
		// Stone
		Prov("sto").Conn("bay", godip.Land).Conn("tan", godip.Land).Conn("chr", godip.Land).Conn("law", godip.Land).Flag(godip.Land).
		// Texas
		Prov("tex").Conn("sha", godip.Land).Conn("den", godip.Land).Conn("phe", godip.Land).Conn("pul", godip.Land).Conn("lac", godip.Land).Conn("wri", godip.Land).Conn("dou", godip.Land).Conn("hoe", godip.Land).Flag(godip.Land).
		// Buchanan
		Prov("buc").Conn("cli", godip.Land).Conn("dek", godip.Land).Conn("and", godip.Land).Conn("pla", godip.Land).Flag(godip.Land).
		// Howard
		Prov("hoa").Conn("boo", godip.Land).Conn("ran", godip.Land).Conn("cha", godip.Land).Conn("sal", godip.Land).Conn("coo", godip.Land).Flag(godip.Land).
		// Andrew
		Prov("and").Conn("buc", godip.Land).Conn("dek", godip.Land).Conn("gen", godip.Land).Conn("nod", godip.Land).Conn("hol", godip.Land).Flag(godip.Land).SC(Missouria).
		// Franklin
		Prov("fra").Conn("stc", godip.Land).Conn("war", godip.Land).Conn("gas", godip.Land).Conn("cra", godip.Land).Conn("jef", godip.Land).Conn("jef", godip.Land).Conn("slo", godip.Land).Flag(godip.Land).
		// Caldwell
		Prov("cad").Conn("dek", godip.Land).Conn("cli", godip.Land).Conn("ray", godip.Land).Conn("car", godip.Land).Conn("liv", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).
		// Clinton
		Prov("cli").Conn("buc", godip.Land).Conn("pla", godip.Land).Conn("cly", godip.Land).Conn("ray", godip.Land).Conn("cad", godip.Land).Conn("dek", godip.Land).Flag(godip.Land).
		// Cedar
		Prov("ced").Conn("dad", godip.Land).Conn("pol", godip.Land).Conn("CLR", godip.Land).Conn("ver", godip.Land).Conn("bao", godip.Land).Flag(godip.Land).
		// Lewis
		Prov("lew").Conn("clr", godip.Land).Conn("kno", godip.Land).Conn("she", godip.Land).Conn("mao", godip.Land).Flag(godip.Land).
		// Putman
		Prov("put").Conn("sch", godip.Land).Conn("mer", godip.Land).Conn("sul", godip.Land).Conn("ada", godip.Land).Flag(godip.Land).
		// Newton
		Prov("new").Conn("jas", godip.Land).Conn("mcd", godip.Land).Conn("bay", godip.Land).Conn("law", godip.Land).Flag(godip.Land).
		// Wright
		Prov("wri").Conn("tex", godip.Land).Conn("lac", godip.Land).Conn("web", godip.Land).Conn("dou", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"pol": "Polk",
	"phe": "Phelps",
	"bao": "Barton",
	"cam": "Camden",
	"mog": "Montgomery",
	"gas": "Gasconade",
	"hol": "Holt",
	"she": "Shelby",
	"slc": "St. Louis City",
	"aud": "Audrain",
	"clr": "Clark",
	"sco": "Scotland",
	"dou": "Douglas",
	"den": "Dent",
	"jas": "Jasper",
	"coo": "Cooper",
	"jef": "Jefferson",
	"slo": "St. Louis",
	"lac": "Laclede",
	"ran": "Randolph",
	"dad": "Dade",
	"bay": "Barry",
	"atc": "Atchison",
	"gen": "Gentry",
	"mac": "Macon",
	"pul": "Pulaski",
	"joh": "Johnson",
	"mao": "Marion",
	"kno": "Knox",
	"col": "Cole",
	"cly": "Clay",
	"osa": "Osage",
	"dek": "De Kalb",
	"mcd": "Mcdonald",
	"ray": "Ray",
	"law": "Lawrence",
	"ral": "Ralls",
	"mae": "Maries",
	"war": "Warren",
	"cll": "Callaway",
	"pik": "Pike",
	"dav": "Daviess",
	"jac": "Jackson",
	"gre": "Greene",
	"nod": "Nodaway",
	"ore": "Oregon",
	"laf": "Lafayette",
	"CLR": "St. Clair",
	"mer": "Mercer",
	"mnr": "Monroe",
	"moi": "Moniteau",
	"ben": "Benton",
	"boo": "Boone",
	"pet": "Pettis",
	"lic": "Lincoln",
	"sul": "Sullivan",
	"bat": "Bates",
	"hic": "Hickory",
	"cha": "Chariton",
	"dal": "Dallas",
	"oza": "Ozark",
	"liv": "Livingston",
	"hoe": "Howell",
	"tan": "Taney",
	"pla": "Platte",
	"hen": "Henry",
	"car": "Carroll",
	"ver": "Vernon",
	"gru": "Grundy",
	"sal": "Saline",
	"stc": "St. Charles",
	"wor": "Worth",
	"sch": "Schuyler",
	"cas": "Cass",
	"ada": "Adair",
	"sha": "Shannon",
	"chr": "Christian",
	"mil": "Miller",
	"lnn": "Linn",
	"web": "Webster",
	"har": "Harrison",
	"cra": "Crawford",
	"mor": "Morgan",
	"sto": "Stone",
	"tex": "Texas",
	"buc": "Buchanan",
	"hoa": "Howard",
	"and": "Andrew",
	"fra": "Franklin",
	"cad": "Caldwell",
	"cli": "Clinton",
	"ced": "Cedar",
	"lew": "Lewis",
	"put": "Putman",
	"new": "Newton",
	"wri": "Wright",
}
