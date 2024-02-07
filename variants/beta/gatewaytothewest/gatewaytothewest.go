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
	SoloWinner:        common.SCCountWinner(21),
	SoloSCCount:       func(*state.State) int { return 21 },
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
		"mis": godip.Unit{godip.Army, Chickasaw},
		"new": godip.Unit{godip.Army, Chickasaw},
		"pem": godip.Unit{godip.Army, Chickasaw},
		"cap": godip.Unit{godip.Army, Chickasaw},
		"osa": godip.Unit{godip.Army, Osage},
		"col": godip.Unit{godip.Army, Osage},
		"jac": godip.Unit{godip.Army, Missouria},
		"cas": godip.Unit{godip.Army, Missouria},
		"laf": godip.Unit{godip.Army, Missouria},
		"stc": godip.Unit{godip.Army, Illini},
		"stl": godip.Unit{godip.Army, Illini},
		"slc": godip.Unit{godip.Army, Illini},
		"pla": godip.Unit{godip.Army, Oto},
		"cly": godip.Unit{godip.Army, Oto},
		"cli": godip.Unit{godip.Army, Oto},
		"mer": godip.Unit{godip.Army, Ioway},
		"put": godip.Unit{godip.Army, Ioway},
		"ada": godip.Unit{godip.Army, Ioway},
		"way": godip.Unit{godip.Army, Quapaw},
		"but": godip.Unit{godip.Army, Quapaw},
		"rip": godip.Unit{godip.Army, Quapaw},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"mis": Chickasaw,
		"new": Chickasaw,
		"pem": Chickasaw,
		"cap": Chickasaw,
		"osa": Osage,
		"col": Osage,
		"jac": Missouria,
		"cas": Missouria,
		"laf": Missouria,
		"stc": Illini,
		"stl": Illini,
		"slc": Illini,
		"pla": Oto,
		"cly": Oto,
		"cli": Oto,
		"mer": Ioway,
		"put": Ioway,
		"ada": Ioway,
		"way": Quapaw,
		"but": Quapaw,
		"rip": Quapaw,
	})
	return
}

func GatewayToTheWestGraph() *graph.Graph {
	return graph.New().
		// Polk
		Prov("pol").Conn("dal", godip.Land).Conn("hic", godip.Land).Conn("scr", godip.Land).Conn("ced", godip.Land).Conn("dad", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Phelps
		Prov("phe").Conn("mae", godip.Land).Conn("pul", godip.Land).Conn("tex", godip.Land).Conn("den", godip.Land).Conn("cra", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).
		// Neosho
		Prov("neo").Conn("bar", godip.Land).Conn("dad", godip.Land).Conn("ver", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Marion
		Prov("mao").Conn("moo", godip.Land).Conn("ral", godip.Land).Conn("lew", godip.Land).Conn("she", godip.Land).Flag(godip.Land).
		// Montgomery
		Prov("mog").Conn("caa", godip.Land).Conn("gas", godip.Land).Conn("war", godip.Land).Conn("lin", godip.Land).Conn("pik", godip.Land).Conn("aud", godip.Land).Flag(godip.Land).
		// Pemiscot
		Prov("pem").Conn("new", godip.Land).Conn("dun", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Washington
		Prov("was").Conn("jef", godip.Land).Conn("fra", godip.Land).Conn("cra", godip.Land).Conn("rey", godip.Land).Conn("stf", godip.Land).Flag(godip.Land).
		// Shelby
		Prov("she").Conn("lew", godip.Land).Conn("ada", godip.Land).Conn("moo", godip.Land).Conn("mao", godip.Land).Conn("lew", godip.Land).Flag(godip.Land).
		// St. Louis City
		Prov("slc").Conn("stl", godip.Land).Conn("stl", godip.Land).Flag(godip.Land).SC(Illini).
		// Audrain
		Prov("aud").Conn("caa", godip.Land).Conn("mog", godip.Land).Conn("pik", godip.Land).Conn("ral", godip.Land).Conn("moo", godip.Land).Conn("ran", godip.Land).Conn("boo", godip.Land).Flag(godip.Land).
		// Clark
		Prov("clr").Conn("lew", godip.Land).Conn("lew", godip.Land).Conn("lew", godip.Land).Conn("sch", godip.Land).Conn("ada", godip.Land).Flag(godip.Land).
		// Dent
		Prov("den").Conn("phe", godip.Land).Conn("tex", godip.Land).Conn("rip", godip.Land).Conn("rey", godip.Land).Conn("rey", godip.Land).Conn("cra", godip.Land).Flag(godip.Land).
		// Cooper
		Prov("coo").Conn("boo", godip.Land).Conn("how", godip.Land).Conn("sal", godip.Land).Conn("pet", godip.Land).Conn("mor", godip.Land).Conn("moi", godip.Land).Flag(godip.Land).
		// Jefferson
		Prov("jef").Conn("stf", godip.Land).Conn("ste", godip.Land).Conn("stl", godip.Land).Conn("fra", godip.Land).Conn("was", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// St. Louis
		Prov("stl").Conn("stc", godip.Land).Conn("fra", godip.Land).Conn("jef", godip.Land).Conn("slc", godip.Land).Conn("slc", godip.Land).Flag(godip.Land).SC(Illini).
		// Laclede
		Prov("lac").Conn("pul", godip.Land).Conn("cam", godip.Land).Conn("dal", godip.Land).Conn("web", godip.Land).Conn("wri", godip.Land).Conn("tex", godip.Land).Flag(godip.Land).
		// Randolph
		Prov("ran").Conn("moo", godip.Land).Conn("ada", godip.Land).Conn("cha", godip.Land).Conn("how", godip.Land).Conn("boo", godip.Land).Conn("aud", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Dade
		Prov("dad").Conn("chr", godip.Land).Conn("gre", godip.Land).Conn("gre", godip.Land).Conn("pol", godip.Land).Conn("ced", godip.Land).Conn("ver", godip.Land).Conn("neo", godip.Land).Conn("bar", godip.Land).Flag(godip.Land).
		// Barry
		Prov("bar").Conn("chr", godip.Land).Conn("chr", godip.Land).Conn("dad", godip.Land).Conn("neo", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Scott
		Prov("sco").Conn("new", godip.Land).Conn("mis", godip.Land).Conn("cap", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).
		// Putman
		Prov("put").Conn("mer", godip.Land).Conn("ada", godip.Land).Conn("ada", godip.Land).Conn("sch", godip.Land).Flag(godip.Land).SC(Ioway).
		// Perry
		Prov("per").Conn("cap", godip.Land).Conn("ste", godip.Land).Conn("stf", godip.Land).Conn("mad", godip.Land).Conn("bol", godip.Land).Flag(godip.Land).
		// Pulaski
		Prov("pul").Conn("lac", godip.Land).Conn("tex", godip.Land).Conn("phe", godip.Land).Conn("mae", godip.Land).Conn("mil", godip.Land).Conn("cam", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// St. Francois
		Prov("stf").Conn("jef", godip.Land).Conn("was", godip.Land).Conn("rey", godip.Land).Conn("mad", godip.Land).Conn("per", godip.Land).Conn("ste", godip.Land).Flag(godip.Land).
		// Johnson
		Prov("joh").Conn("pet", godip.Land).Conn("laf", godip.Land).Conn("jac", godip.Land).Conn("cas", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// Dunklin
		Prov("dun").Conn("sto", godip.Land).Conn("but", godip.Land).Conn("pem", godip.Land).Conn("new", godip.Land).Flag(godip.Land).
		// Camden
		Prov("cam").Conn("pul", godip.Land).Conn("mil", godip.Land).Conn("mor", godip.Land).Conn("ben", godip.Land).Conn("hic", godip.Land).Conn("dal", godip.Land).Conn("lac", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Cole
		Prov("col").Conn("osa", godip.Land).Conn("caa", godip.Land).Conn("boo", godip.Land).Conn("moi", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).SC(Osage).
		// Clay
		Prov("cly").Conn("jac", godip.Land).Conn("ray", godip.Land).Conn("cli", godip.Land).Conn("pla", godip.Land).Flag(godip.Land).SC(Oto).
		// Osage
		Prov("osa").Conn("col", godip.Land).Conn("mil", godip.Land).Conn("mae", godip.Land).Conn("gas", godip.Land).Conn("caa", godip.Land).Flag(godip.Land).SC(Osage).
		// Ray
		Prov("ray").Conn("laf", godip.Land).Conn("car", godip.Land).Conn("cad", godip.Land).Conn("cli", godip.Land).Conn("cly", godip.Land).Conn("jac", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ralls
		Prov("ral").Conn("mao", godip.Land).Conn("moo", godip.Land).Conn("aud", godip.Land).Conn("pik", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Maries
		Prov("mae").Conn("phe", godip.Land).Conn("gas", godip.Land).Conn("osa", godip.Land).Conn("mil", godip.Land).Conn("pul", godip.Land).Flag(godip.Land).
		// Warren
		Prov("war").Conn("stc", godip.Land).Conn("lin", godip.Land).Conn("mog", godip.Land).Conn("gas", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).
		// Callaway
		Prov("caa").Conn("mog", godip.Land).Conn("aud", godip.Land).Conn("boo", godip.Land).Conn("col", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Pike
		Prov("pik").Conn("lin", godip.Land).Conn("ral", godip.Land).Conn("aud", godip.Land).Conn("mog", godip.Land).Flag(godip.Land).
		// Daviess
		Prov("dav").Conn("cli", godip.Land).Conn("cli", godip.Land).Conn("cad", godip.Land).Conn("liv", godip.Land).Conn("gru", godip.Land).Conn("mer", godip.Land).Flag(godip.Land).
		// Jackson
		Prov("jac").Conn("cas", godip.Land).Conn("joh", godip.Land).Conn("laf", godip.Land).Conn("ray", godip.Land).Conn("cly", godip.Land).Flag(godip.Land).SC(Missouria).
		// Greene
		Prov("gre").Conn("pol", godip.Land).Conn("dad", godip.Land).Conn("dad", godip.Land).Conn("chr", godip.Land).Conn("web", godip.Land).Conn("dal", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Cape Girardeau
		Prov("cap").Conn("per", godip.Land).Conn("bol", godip.Land).Conn("sto", godip.Land).Conn("sco", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Wayne
		Prov("way").Conn("but", godip.Land).Conn("sto", godip.Land).Conn("bol", godip.Land).Conn("mad", godip.Land).Conn("rey", godip.Land).Conn("rey", godip.Land).Conn("rip", godip.Land).Flag(godip.Land).SC(Quapaw).
		// Lafayette
		Prov("laf").Conn("ray", godip.Land).Conn("jac", godip.Land).Conn("joh", godip.Land).Conn("pet", godip.Land).Conn("sal", godip.Land).Conn("car", godip.Land).Flag(godip.Land).SC(Missouria).
		// St. Clair
		Prov("scr").Conn("hic", godip.Land).Conn("ben", godip.Land).Conn("hen", godip.Land).Conn("bat", godip.Land).Conn("ver", godip.Land).Conn("ced", godip.Land).Conn("pol", godip.Land).Flag(godip.Land).
		// Mercer
		Prov("mer").Conn("gru", godip.Land).Conn("gru", godip.Land).Conn("ada", godip.Land).Conn("put", godip.Land).Conn("pla", godip.Land).Conn("cli", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).SC(Ioway).
		// Monroe
		Prov("moo").Conn("ran", godip.Land).Conn("aud", godip.Land).Conn("ral", godip.Land).Conn("mao", godip.Land).Conn("she", godip.Land).Flag(godip.Land).
		// Moniteau
		Prov("moi").Conn("coo", godip.Land).Conn("mor", godip.Land).Conn("mil", godip.Land).Conn("col", godip.Land).Conn("boo", godip.Land).Flag(godip.Land).
		// Benton
		Prov("ben").Conn("scr", godip.Land).Conn("hic", godip.Land).Conn("cam", godip.Land).Conn("mor", godip.Land).Conn("pet", godip.Land).Conn("hen", godip.Land).Flag(godip.Land).
		// New Madrid
		Prov("new").Conn("mis", godip.Land).Conn("sco", godip.Land).Conn("sto", godip.Land).Conn("dun", godip.Land).Conn("pem", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Boone
		Prov("boo").Conn("coo", godip.Land).Conn("moi", godip.Land).Conn("col", godip.Land).Conn("caa", godip.Land).Conn("aud", godip.Land).Conn("ran", godip.Land).Conn("how", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Pettis
		Prov("pet").Conn("joh", godip.Land).Conn("hen", godip.Land).Conn("ben", godip.Land).Conn("mor", godip.Land).Conn("coo", godip.Land).Conn("sal", godip.Land).Conn("laf", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lincoln
		Prov("lin").Conn("pik", godip.Land).Conn("mog", godip.Land).Conn("war", godip.Land).Conn("stc", godip.Land).Flag(godip.Land).
		// Reynolds
		Prov("rey").Conn("rip", godip.Land).Conn("rip", godip.Land).Conn("way", godip.Land).Conn("way", godip.Land).Conn("mad", godip.Land).Conn("stf", godip.Land).Conn("was", godip.Land).Conn("cra", godip.Land).Conn("den", godip.Land).Conn("den", godip.Land).Flag(godip.Land).
		// Bates
		Prov("bat").Conn("ver", godip.Land).Conn("scr", godip.Land).Conn("hen", godip.Land).Conn("cas", godip.Land).Flag(godip.Land).
		// Hickory
		Prov("hic").Conn("dal", godip.Land).Conn("cam", godip.Land).Conn("ben", godip.Land).Conn("scr", godip.Land).Conn("pol", godip.Land).Flag(godip.Land).
		// Chariton
		Prov("cha").Conn("ran", godip.Land).Conn("ada", godip.Land).Conn("ada", godip.Land).Conn("liv", godip.Land).Conn("car", godip.Land).Conn("sal", godip.Land).Conn("how", godip.Land).Flag(godip.Land).
		// Dallas
		Prov("dal").Conn("pol", godip.Land).Conn("gre", godip.Land).Conn("web", godip.Land).Conn("lac", godip.Land).Conn("cam", godip.Land).Conn("hic", godip.Land).Flag(godip.Land).
		// Cedar
		Prov("ced").Conn("ver", godip.Land).Conn("dad", godip.Land).Conn("pol", godip.Land).Conn("scr", godip.Land).Conn("ver", godip.Land).Flag(godip.Land).
		// Livingston
		Prov("liv").Conn("cad", godip.Land).Conn("car", godip.Land).Conn("cha", godip.Land).Conn("ada", godip.Land).Conn("gru", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).
		// Stoddard
		Prov("sto").Conn("new", godip.Land).Conn("sco", godip.Land).Conn("cap", godip.Land).Conn("bol", godip.Land).Conn("way", godip.Land).Conn("but", godip.Land).Conn("dun", godip.Land).Flag(godip.Land).
		// Ste. Genevieve
		Prov("ste").Conn("per", godip.Land).Conn("jef", godip.Land).Conn("stf", godip.Land).Flag(godip.Land).
		// Platte
		Prov("pla").Conn("cly", godip.Land).Conn("cli", godip.Land).Conn("mer", godip.Land).Flag(godip.Land).SC(Oto).
		// Henry
		Prov("hen").Conn("pet", godip.Land).Conn("joh", godip.Land).Conn("cas", godip.Land).Conn("bat", godip.Land).Conn("scr", godip.Land).Conn("ben", godip.Land).Flag(godip.Land).
		// Madison
		Prov("mad").Conn("bol", godip.Land).Conn("per", godip.Land).Conn("stf", godip.Land).Conn("rey", godip.Land).Conn("way", godip.Land).Flag(godip.Land).
		// Carroll
		Prov("car").Conn("cha", godip.Land).Conn("liv", godip.Land).Conn("cad", godip.Land).Conn("ray", godip.Land).Conn("laf", godip.Land).Conn("sal", godip.Land).Flag(godip.Land).
		// Vernon
		Prov("ver").Conn("neo", godip.Land).Conn("dad", godip.Land).Conn("ced", godip.Land).Conn("ced", godip.Land).Conn("scr", godip.Land).Conn("bat", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Grundy
		Prov("gru").Conn("mer", godip.Land).Conn("dav", godip.Land).Conn("liv", godip.Land).Conn("ada", godip.Land).Conn("ada", godip.Land).Conn("mer", godip.Land).Flag(godip.Land).
		// Saline
		Prov("sal").Conn("coo", godip.Land).Conn("how", godip.Land).Conn("cha", godip.Land).Conn("car", godip.Land).Conn("laf", godip.Land).Conn("pet", godip.Land).Flag(godip.Land).
		// St. Charles
		Prov("stc").Conn("war", godip.Land).Conn("fra", godip.Land).Conn("stl", godip.Land).Conn("lin", godip.Land).Flag(godip.Land).SC(Illini).
		// Mississippi
		Prov("mis").Conn("new", godip.Land).Conn("sco", godip.Land).Flag(godip.Land).SC(Chickasaw).
		// Schuyler
		Prov("sch").Conn("ada", godip.Land).Conn("clr", godip.Land).Conn("put", godip.Land).Flag(godip.Land).
		// Cass
		Prov("cas").Conn("hen", godip.Land).Conn("joh", godip.Land).Conn("jac", godip.Land).Conn("bat", godip.Land).Flag(godip.Land).SC(Missouria).
		// Adair
		Prov("ada").Conn("lew", godip.Land).Conn("lew", godip.Land).Conn("clr", godip.Land).Conn("sch", godip.Land).Conn("put", godip.Land).Conn("put", godip.Land).Conn("mer", godip.Land).Conn("gru", godip.Land).Conn("gru", godip.Land).Conn("liv", godip.Land).Conn("cha", godip.Land).Conn("cha", godip.Land).Conn("ran", godip.Land).Conn("she", godip.Land).Flag(godip.Land).SC(Ioway).
		// Christian
		Prov("chr").Conn("dad", godip.Land).Conn("bar", godip.Land).Conn("bar", godip.Land).Conn("oza", godip.Land).Conn("oza", godip.Land).Conn("oza", godip.Land).Conn("web", godip.Land).Conn("gre", godip.Land).Flag(godip.Land).
		// Miller
		Prov("mil").Conn("mor", godip.Land).Conn("cam", godip.Land).Conn("pul", godip.Land).Conn("mae", godip.Land).Conn("osa", godip.Land).Conn("col", godip.Land).Conn("moi", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Butler
		Prov("but").Conn("way", godip.Land).Conn("rip", godip.Land).Conn("rip", godip.Land).Conn("dun", godip.Land).Conn("sto", godip.Land).Flag(godip.Land).SC(Quapaw).
		// Webster
		Prov("web").Conn("oza", godip.Land).Conn("wri", godip.Land).Conn("lac", godip.Land).Conn("dal", godip.Land).Conn("gre", godip.Land).Conn("chr", godip.Land).Flag(godip.Land).
		// Crawford
		Prov("cra").Conn("phe", godip.Land).Conn("den", godip.Land).Conn("rey", godip.Land).Conn("was", godip.Land).Conn("fra", godip.Land).Conn("gas", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Morgan
		Prov("mor").Conn("mil", godip.Land).Conn("moi", godip.Land).Conn("coo", godip.Land).Conn("pet", godip.Land).Conn("ben", godip.Land).Conn("cam", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bollinger
		Prov("bol").Conn("mad", godip.Land).Conn("way", godip.Land).Conn("sto", godip.Land).Conn("cap", godip.Land).Conn("per", godip.Land).Flag(godip.Land).
		// Texas
		Prov("tex").Conn("lac", godip.Land).Conn("wri", godip.Land).Conn("oza", godip.Land).Conn("rip", godip.Land).Conn("rip", godip.Land).Conn("den", godip.Land).Conn("phe", godip.Land).Conn("pul", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Gasconade
		Prov("gas").Conn("fra", godip.Land).Conn("war", godip.Land).Conn("mog", godip.Land).Conn("osa", godip.Land).Conn("mae", godip.Land).Conn("phe", godip.Land).Conn("cra", godip.Land).Flag(godip.Land).
		// Howard
		Prov("how").Conn("ran", godip.Land).Conn("cha", godip.Land).Conn("sal", godip.Land).Conn("coo", godip.Land).Conn("boo", godip.Land).Flag(godip.Land).
		// Ripley
		Prov("rip").Conn("oza", godip.Land).Conn("but", godip.Land).Conn("but", godip.Land).Conn("way", godip.Land).Conn("rey", godip.Land).Conn("rey", godip.Land).Conn("den", godip.Land).Conn("tex", godip.Land).Conn("tex", godip.Land).Conn("oza", godip.Land).Flag(godip.Land).SC(Quapaw).
		// Franklin
		Prov("fra").Conn("gas", godip.Land).Conn("cra", godip.Land).Conn("was", godip.Land).Conn("jef", godip.Land).Conn("stl", godip.Land).Conn("stc", godip.Land).Conn("war", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Caldwell
		Prov("cad").Conn("liv", godip.Land).Conn("dav", godip.Land).Conn("cli", godip.Land).Conn("cli", godip.Land).Conn("ray", godip.Land).Conn("car", godip.Land).Flag(godip.Land).
		// Clinton
		Prov("cli").Conn("dav", godip.Land).Conn("mer", godip.Land).Conn("pla", godip.Land).Conn("cly", godip.Land).Conn("ray", godip.Land).Conn("cad", godip.Land).Conn("cad", godip.Land).Conn("dav", godip.Land).Flag(godip.Land).SC(Oto).
		// Ozark
		Prov("oza").Conn("web", godip.Land).Conn("chr", godip.Land).Conn("chr", godip.Land).Conn("chr", godip.Land).Conn("rip", godip.Land).Conn("rip", godip.Land).Conn("tex", godip.Land).Conn("wri", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lewis
		Prov("lew").Conn("she", godip.Land).Conn("she", godip.Land).Conn("mao", godip.Land).Conn("clr", godip.Land).Conn("clr", godip.Land).Conn("clr", godip.Land).Conn("ada", godip.Land).Conn("ada", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wright
		Prov("wri").Conn("web", godip.Land).Conn("oza", godip.Land).Conn("tex", godip.Land).Conn("lac", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"pol": "Polk",
	"phe": "Phelps",
	"neo": "Neosho",
	"mao": "Marion",
	"mog": "Montgomery",
	"pem": "Pemiscot",
	"was": "Washington",
	"she": "Shelby",
	"slc": "St. Louis City",
	"aud": "Audrain",
	"clr": "Clark",
	"den": "Dent",
	"coo": "Cooper",
	"jef": "Jefferson",
	"stl": "St. Louis",
	"lac": "Laclede",
	"ran": "Randolph",
	"dad": "Dade",
	"bar": "Barry",
	"sco": "Scott",
	"put": "Putman",
	"per": "Perry",
	"pul": "Pulaski",
	"stf": "St. Francois",
	"joh": "Johnson",
	"dun": "Dunklin",
	"cam": "Camden",
	"col": "Cole",
	"cly": "Clay",
	"osa": "Osage",
	"ray": "Ray",
	"ral": "Ralls",
	"mae": "Maries",
	"war": "Warren",
	"caa": "Callaway",
	"pik": "Pike",
	"dav": "Daviess",
	"jac": "Jackson",
	"gre": "Greene",
	"cap": "Cape Girardeau",
	"way": "Wayne",
	"laf": "Lafayette",
	"scr": "St. Clair",
	"mer": "Mercer",
	"moo": "Monroe",
	"moi": "Moniteau",
	"ben": "Benton",
	"new": "New Madrid",
	"boo": "Boone",
	"pet": "Pettis",
	"lin": "Lincoln",
	"rey": "Reynolds",
	"bat": "Bates",
	"hic": "Hickory",
	"cha": "Chariton",
	"dal": "Dallas",
	"ced": "Cedar",
	"liv": "Livingston",
	"sto": "Stoddard",
	"ste": "Ste. Genevieve",
	"pla": "Platte",
	"hen": "Henry",
	"mad": "Madison",
	"car": "Carroll",
	"ver": "Vernon",
	"gru": "Grundy",
	"sal": "Saline",
	"stc": "St. Charles",
	"mis": "Mississippi",
	"sch": "Schuyler",
	"cas": "Cass",
	"ada": "Adair",
	"chr": "Christian",
	"mil": "Miller",
	"but": "Butler",
	"web": "Webster",
	"cra": "Crawford",
	"mor": "Morgan",
	"bol": "Bollinger",
	"tex": "Texas",
	"gas": "Gasconade",
	"how": "Howard",
	"rip": "Ripley",
	"fra": "Franklin",
	"cad": "Caldwell",
	"cli": "Clinton",
	"oza": "Ozark",
	"lew": "Lewis",
	"wri": "Wright",
}
