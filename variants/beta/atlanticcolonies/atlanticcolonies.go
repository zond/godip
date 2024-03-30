package atlanticcolonies

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	England  godip.Nation = "England"
	Portugal godip.Nation = "Portugal"
	Spain    godip.Nation = "Spain"
	France   godip.Nation = "France"
)

var Nations = []godip.Nation{England, Portugal, Spain, France}

var AtlanticColoniesVariant = common.Variant{
	Name:              "Atlantic Colonies",
	Graph:             func() godip.Graph { return AtlanticColoniesGraph() },
	Start:             AtlanticColoniesStart,
	Blank:             AtlanticColoniesBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(25),
	SoloSCCount:       func(*state.State) int { return 25 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/atlanticcoloniesmap.svg")
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
	CreatedBy: "Safari",
	Version:   "0.1",
	Description: `Set in the late 17th century, and focusing on the colonial exploits of England, France, Portugal, and Spain in the New World and Africa, 
				 it requires players to engage in diplomacy on multiple fronts simultaneously to achieve victory.`,
	Rules: `The short stretches of water between Flordia and Cuba, as well as Montreal and Newfoundland, are traversable by armies as indicated by red connectors.
				 Adjacent islands, such as Jamacia and Cuba, or Canary Islands and Madiera, are considered "island chains" and are traversable by armies.
				 It is possible for a fleet to move to Hudson Bay from Northwestern Atlantic via an unshown border to the north of the map.`,
}

func AtlanticColoniesBlank(phase godip.Phase) *state.State {
	return state.New(AtlanticColoniesGraph(), phase, classical.BackupRule, nil, nil)
}

func AtlanticColoniesStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1673, godip.Spring, godip.Movement)
	result = AtlanticColoniesBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"lon": godip.Unit{godip.Fleet, England},
		"bri": godip.Unit{godip.Fleet, England},
		"nee": godip.Unit{godip.Fleet, England},
		"jam": godip.Unit{godip.Fleet, England},
		"cog": godip.Unit{godip.Fleet, England},
		"moo": godip.Unit{godip.Army, England},
		"geo": godip.Unit{godip.Army, England},
		"sao": godip.Unit{godip.Fleet, Portugal},
		"azo": godip.Unit{godip.Fleet, Portugal},
		"rio": godip.Unit{godip.Fleet, Portugal},
		"ang": godip.Unit{godip.Fleet, Portugal},
		"por": godip.Unit{godip.Fleet, Portugal},
		"maz": godip.Unit{godip.Army, Portugal},
		"mnh": godip.Unit{godip.Army, Portugal},
		"mer": godip.Unit{godip.Fleet, Spain},
		"gib": godip.Unit{godip.Fleet, Spain},
		"val": godip.Unit{godip.Fleet, Spain},
		"lim": godip.Unit{godip.Army, Spain},
		"mex": godip.Unit{godip.Army, Spain},
		"sat": godip.Unit{godip.Army, Spain},
		"cai": godip.Unit{godip.Army, Spain},
		"bre": godip.Unit{godip.Fleet, France},
		"mss": godip.Unit{godip.Fleet, France},
		"win": godip.Unit{godip.Fleet, France},
		"gor": godip.Unit{godip.Fleet, France},
		"stl": godip.Unit{godip.Army, France},
		"mon": godip.Unit{godip.Army, France},
		"cay": godip.Unit{godip.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"lon": England,
		"bri": England,
		"nee": England,
		"jam": England,
		"cog": England,
		"moo": England,
		"geo": England,
		"sao": Portugal,
		"azo": Portugal,
		"rio": Portugal,
		"ang": Portugal,
		"por": Portugal,
		"maz": Portugal,
		"mnh": Portugal,
		"mer": Spain,
		"gib": Spain,
		"val": Spain,
		"lim": Spain,
		"mex": Spain,
		"sat": Spain,
		"cai": Spain,
		"bre": France,
		"mss": France,
		"win": France,
		"gor": France,
		"stl": France,
		"mon": France,
		"cay": France,
	})
	return
}

func AtlanticColoniesGraph() *graph.Graph {
	return graph.New().
		// Upper Mississippi
		Prov("upp").Conn("ont", godip.Land).Conn("wrl", godip.Land).Conn("lou", godip.Land).Conn("stl", godip.Land).Conn("det", godip.Land).Flag(godip.Land).
		// Bogota
		Prov("bog").Conn("sur", godip.Land).Conn("cca", godip.Land).Conn("neg", godip.Land).Conn("lim", godip.Land).Flag(godip.Land).
		// Bermuda Triangle
		Prov("bet").Conn("mia", godip.Sea).Conn("beu", godip.Sea).Conn("ese", godip.Sea).Conn("geo", godip.Sea).Conn("geo/ec", godip.Sea).Conn("flo", godip.Sea).Conn("bah", godip.Sea).Conn("cub", godip.Sea).Conn("jam", godip.Sea).Conn("his", godip.Sea).Conn("his", godip.Sea).Conn("lee", godip.Sea).Flag(godip.Sea).
		// Scotland
		Prov("sco").Conn("ote", godip.Sea).Conn("iri", godip.Sea).Conn("bri", godip.Coast...).Conn("lon", godip.Coast...).Flag(godip.Coast...).
		// East Africa
		Prov("eaa").Conn("sud", godip.Land).Conn("kon", godip.Land).Conn("nam", godip.Land).Conn("cog", godip.Land).Flag(godip.Land).
		// Cuba
		Prov("cub").Conn("flo", godip.Coast...).Conn("bah", godip.Coast...).Conn("gom", godip.Sea).Conn("wec", godip.Sea).Conn("jam", godip.Coast...).Conn("bet", godip.Sea).Flag(godip.Coast...).
		// Florida
		Prov("flo").Conn("cub", godip.Coast...).Conn("geo", godip.Land).Conn("geo/ec", godip.Sea).Conn("geo/sc", godip.Sea).Conn("gom", godip.Sea).Conn("bah", godip.Coast...).Conn("bet", godip.Sea).Flag(godip.Coast...).
		// Detroit
		Prov("det").Conn("upp", godip.Land).Conn("stl", godip.Land).Conn("app", godip.Land).Conn("iro", godip.Land).Conn("ont", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// London
		Prov("lon").Conn("ote", godip.Sea).Conn("sco", godip.Coast...).Conn("bri", godip.Coast...).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(England).
		// Azores Islands
		Prov("azo").Conn("cse", godip.Sea).Conn("eur", godip.Sea).Conn("wea", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// Western Atlantic
		Prov("wea").Conn("beu", godip.Sea).Conn("mia", godip.Sea).Conn("cse", godip.Sea).Conn("azo", godip.Sea).Conn("eur", godip.Sea).Conn("ote", godip.Sea).Conn("now", godip.Sea).Conn("nef", godip.Sea).Conn("ese", godip.Sea).Flag(godip.Sea).
		// Southern Atlantic
		Prov("soa").Conn("coa", godip.Sea).Conn("soo", godip.Sea).Conn("sea", godip.Sea).Conn("gre", godip.Sea).Conn("swa", godip.Sea).Conn("sao", godip.Sea).Flag(godip.Sea).
		// Newfoundland
		Prov("nef").Conn("mon", godip.Coast...).Conn("ese", godip.Sea).Conn("wea", godip.Sea).Conn("now", godip.Sea).Conn("now", godip.Sea).Conn("now", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Ontario
		Prov("ont").Conn("iro", godip.Land).Conn("mon", godip.Land).Conn("erl", godip.Land).Conn("moo", godip.Land).Conn("wrl", godip.Land).Conn("upp", godip.Land).Conn("det", godip.Land).Flag(godip.Land).
		// Hudson Bay
		Prov("hud").Conn("the", godip.Sea).Conn("moo", godip.Sea).Conn("erl", godip.Sea).Conn("mon", godip.Sea).Conn("now", godip.Sea).Flag(godip.Sea).
		// Caribbean Atlantic
		Prov("caa").Conn("trn", godip.Sea).Conn("sur", godip.Sea).Conn("cay", godip.Sea).Conn("eem", godip.Sea).Conn("swa", godip.Sea).Conn("cav", godip.Sea).Conn("mia", godip.Sea).Conn("ant", godip.Sea).Conn("win", godip.Sea).Conn("bar", godip.Sea).Flag(godip.Sea).
		// Gold Coast
		Prov("gol").Conn("gre", godip.Sea).Conn("sla", godip.Coast...).Conn("nig", godip.Land).Conn("gra", godip.Coast...).Flag(godip.Coast...).
		// Brest
		Prov("bre").Conn("eng", godip.Sea).Conn("eur", godip.Sea).Conn("bor", godip.Coast...).Conn("par", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Sao Salvador
		Prov("sao").Conn("rio", godip.Coast...).Conn("coa", godip.Sea).Conn("soa", godip.Sea).Conn("swa", godip.Sea).Conn("mnh", godip.Coast...).Conn("goy", godip.Land).Conn("saa", godip.Land).Flag(godip.Coast...).SC(Portugal).
		// Paris
		Prov("par").Conn("mss", godip.Land).Conn("eng", godip.Sea).Conn("bre", godip.Coast...).Conn("bor", godip.Land).Flag(godip.Coast...).
		// Ionian Sea
		Prov("ion").Conn("guo", godip.Sea).Conn("trp", godip.Sea).Flag(godip.Sea).
		// Gulf of Mexico
		Prov("gom").Conn("wec", godip.Sea).Conn("cub", godip.Sea).Conn("bah", godip.Sea).Conn("flo", godip.Sea).Conn("geo", godip.Sea).Conn("geo/sc", godip.Sea).Conn("neo", godip.Sea).Conn("saf", godip.Sea).Conn("mex", godip.Sea).Conn("mer", godip.Sea).Flag(godip.Sea).
		// Cayenne
		Prov("cay").Conn("goy", godip.Land).Conn("eem", godip.Coast...).Conn("caa", godip.Sea).Conn("sur", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Grain Coast
		Prov("gra").Conn("gor", godip.Coast...).Conn("gre", godip.Sea).Conn("gol", godip.Coast...).Conn("nig", godip.Land).Conn("sen", godip.Coast...).Flag(godip.Coast...).
		// Windward Islands
		Prov("win").Conn("bar", godip.Coast...).Conn("caa", godip.Sea).Conn("ant", godip.Coast...).Conn("lee", godip.Coast...).Conn("eac", godip.Sea).Conn("cca", godip.Coast...).Conn("trn", godip.Coast...).Flag(godip.Coast...).SC(France).
		// Northwestern Atlantic
		Prov("now").Conn("hud", godip.Sea).Conn("mon", godip.Sea).Conn("mon", godip.Sea).Conn("iro", godip.Sea).Conn("nov", godip.Sea).Conn("nov", godip.Sea).Conn("ese", godip.Sea).Conn("nef", godip.Sea).Conn("nef", godip.Sea).Conn("nef", godip.Sea).Conn("wea", godip.Sea).Conn("ote", godip.Sea).Flag(godip.Sea).
		// Bermuda
		Prov("beu").Conn("wea", godip.Sea).Conn("ese", godip.Sea).Conn("bet", godip.Sea).Conn("mia", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Eastern Seaboard
		Prov("ese").Conn("nef", godip.Sea).Conn("now", godip.Sea).Conn("nov", godip.Sea).Conn("nee", godip.Sea).Conn("pen", godip.Sea).Conn("geo", godip.Sea).Conn("geo/ec", godip.Sea).Conn("bet", godip.Sea).Conn("beu", godip.Sea).Conn("wea", godip.Sea).Flag(godip.Sea).
		// St. Louis
		Prov("stl").Conn("neo", godip.Land).Conn("app", godip.Land).Conn("det", godip.Land).Conn("upp", godip.Land).Conn("lou", godip.Land).Flag(godip.Land).SC(France).
		// Gibraltar
		Prov("gib").Conn("val", godip.Coast...).Conn("and", godip.Land).Conn("str", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Merida
		Prov("mer").Conn("mex", godip.Coast...).Conn("sie", godip.Land).Conn("bei", godip.Coast...).Conn("wec", godip.Sea).Conn("gom", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// New Granada
		Prov("neg").Conn("lim", godip.Land).Conn("bog", godip.Land).Conn("cca", godip.Coast...).Conn("eac", godip.Sea).Conn("wec", godip.Sea).Conn("bei", godip.Coast...).Flag(godip.Coast...).
		// Canary Islands
		Prov("cai").Conn("wsa", godip.Coast...).Conn("som", godip.Coast...).Conn("eur", godip.Sea).Conn("mad", godip.Coast...).Conn("cse", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Pennsylvania
		Prov("pen").Conn("ese", godip.Sea).Conn("nee", godip.Coast...).Conn("iro", godip.Land).Conn("app", godip.Land).Conn("geo", godip.Land).Conn("geo/ec", godip.Sea).Flag(godip.Coast...).
		// Maranhao
		Prov("mnh").Conn("eem", godip.Coast...).Conn("goy", godip.Land).Conn("sao", godip.Coast...).Conn("swa", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// Eastern Caribbean
		Prov("eac").Conn("lee", godip.Sea).Conn("his", godip.Sea).Conn("jam", godip.Sea).Conn("wec", godip.Sea).Conn("neg", godip.Sea).Conn("cca", godip.Sea).Conn("win", godip.Sea).Flag(godip.Sea).
		// Algeria
		Prov("alg").Conn("nom", godip.Coast...).Conn("mis", godip.Land).Conn("tun", godip.Coast...).Conn("bae", godip.Sea).Conn("str", godip.Sea).Flag(godip.Coast...).
		// Drake's Passage
		Prov("dra").Conn("fal", godip.Sea).Conn("soo", godip.Sea).Conn("coa", godip.Sea).Conn("bue", godip.Sea).Conn("arg", godip.Sea).Flag(godip.Sea).
		// Madeira
		Prov("mad").Conn("eur", godip.Sea).Conn("cse", godip.Sea).Conn("cai", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Argentina
		Prov("arg").Conn("sat", godip.Land).Conn("dra", godip.Sea).Conn("bue", godip.Coast...).Flag(godip.Coast...).
		// Greater Gulf of Guinea
		Prov("gre").Conn("gol", godip.Sea).Conn("gra", godip.Sea).Conn("gor", godip.Sea).Conn("cav", godip.Sea).Conn("swa", godip.Sea).Conn("soa", godip.Sea).Conn("sea", godip.Sea).Conn("sla", godip.Sea).Flag(godip.Sea).
		// West Sahara
		Prov("wsa").Conn("som", godip.Coast...).Conn("cai", godip.Coast...).Conn("cse", godip.Sea).Conn("cav", godip.Coast...).Conn("sen", godip.Coast...).Conn("nig", godip.Land).Conn("mis", godip.Land).Flag(godip.Coast...).
		// Nigrita
		Prov("nig").Conn("gra", godip.Land).Conn("gol", godip.Land).Conn("sla", godip.Land).Conn("sud", godip.Land).Conn("mis", godip.Land).Conn("wsa", godip.Land).Conn("sen", godip.Land).Flag(godip.Land).
		// The North West
		Prov("the").Conn("wrl", godip.Land).Conn("moo", godip.Coast...).Conn("hud", godip.Sea).Flag(godip.Coast...).
		// Goyaz
		Prov("goy").Conn("sao", godip.Land).Conn("mnh", godip.Land).Conn("eem", godip.Land).Conn("cay", godip.Land).Conn("saa", godip.Land).Flag(godip.Land).
		// Santa Fe
		Prov("saf").Conn("sie", godip.Land).Conn("mex", godip.Coast...).Conn("gom", godip.Sea).Conn("neo", godip.Coast...).Conn("lou", godip.Land).Conn("cal", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Lima
		Prov("lim").Conn("neg", godip.Land).Conn("sat", godip.Land).Conn("saa", godip.Land).Conn("bog", godip.Land).Flag(godip.Land).SC(Spain).
		// Ireland
		Prov("ire").Conn("ote", godip.Sea).Conn("iri", godip.Sea).Conn("iri", godip.Sea).Flag(godip.Coast...).
		// Cape of Good Hope
		Prov("cog").Conn("eaa", godip.Land).Conn("nam", godip.Coast...).Conn("soo", godip.Sea).Flag(godip.Coast...).SC(England).
		// Sudan
		Prov("sud").Conn("esa", godip.Land).Conn("mis", godip.Land).Conn("nig", godip.Land).Conn("sla", godip.Land).Conn("kon", godip.Land).Conn("eaa", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Middle Sahara
		Prov("mis").Conn("tun", godip.Land).Conn("alg", godip.Land).Conn("nom", godip.Land).Conn("som", godip.Land).Conn("wsa", godip.Land).Conn("nig", godip.Land).Conn("sud", godip.Land).Conn("esa", godip.Land).Flag(godip.Land).
		// Namib Desert
		Prov("nam").Conn("cog", godip.Coast...).Conn("eaa", godip.Land).Conn("kon", godip.Land).Conn("ang", godip.Coast...).Conn("sea", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Coast...).
		// Valencia
		Prov("val").Conn("gib", godip.Coast...).Conn("str", godip.Sea).Conn("bae", godip.Sea).Conn("cat", godip.Land).Conn("cas", godip.Land).Conn("and", godip.Land).Flag(godip.Coast...).SC(Spain).
		// Central Brazil
		Prov("cen").Conn("sat", godip.Land).Conn("bue", godip.Land).Conn("rio", godip.Land).Conn("saa", godip.Land).Flag(godip.Land).
		// Bordeaux
		Prov("bor").Conn("eur", godip.Sea).Conn("cas", godip.Coast...).Conn("cat", godip.Land).Conn("mss", godip.Land).Conn("par", godip.Land).Conn("bre", godip.Coast...).Flag(godip.Coast...).
		// English Channel
		Prov("eng").Conn("ote", godip.Sea).Conn("lon", godip.Sea).Conn("bri", godip.Sea).Conn("iri", godip.Sea).Conn("eur", godip.Sea).Conn("bre", godip.Sea).Conn("par", godip.Sea).Flag(godip.Sea).
		// Cape Verde
		Prov("cav").Conn("cse", godip.Sea).Conn("mia", godip.Sea).Conn("caa", godip.Sea).Conn("swa", godip.Sea).Conn("gre", godip.Sea).Conn("gor", godip.Coast...).Conn("sen", godip.Coast...).Conn("wsa", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Western Caribbean
		Prov("wec").Conn("gom", godip.Sea).Conn("mer", godip.Sea).Conn("bei", godip.Sea).Conn("neg", godip.Sea).Conn("eac", godip.Sea).Conn("jam", godip.Sea).Conn("cub", godip.Sea).Flag(godip.Sea).
		// Oregon
		Prov("ore").Conn("wrl", godip.Land).Conn("cal", godip.Land).Conn("lou", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Antigua
		Prov("ant").Conn("lee", godip.Coast...).Conn("win", godip.Coast...).Conn("caa", godip.Sea).Conn("mia", godip.Sea).Flag(godip.Coast...).
		// Andalucia
		Prov("and").Conn("por", godip.Coast...).Conn("eur", godip.Sea).Conn("gib", godip.Land).Conn("val", godip.Land).Conn("cas", godip.Coast...).Flag(godip.Coast...).
		// Catalan
		Prov("cat").Conn("guo", godip.Sea).Conn("mss", godip.Coast...).Conn("bor", godip.Land).Conn("cas", godip.Land).Conn("val", godip.Land).Flag(godip.Coast...).
		// Senegal
		Prov("sen").Conn("nig", godip.Land).Conn("wsa", godip.Land).Conn("cav", godip.Land).Conn("gor", godip.Land).Conn("gra", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Rio de Janeiro
		Prov("rio").Conn("sao", godip.Coast...).Conn("saa", godip.Land).Conn("cen", godip.Land).Conn("bue", godip.Coast...).Conn("coa", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// Santa Anna
		Prov("saa").Conn("sat", godip.Land).Conn("cen", godip.Land).Conn("rio", godip.Land).Conn("sao", godip.Land).Conn("goy", godip.Land).Conn("lim", godip.Land).Flag(godip.Land).
		// Belem
		Prov("eem").Conn("caa", godip.Sea).Conn("cay", godip.Coast...).Conn("goy", godip.Land).Conn("mnh", godip.Coast...).Conn("swa", godip.Sea).Flag(godip.Coast...).
		// Marseilles
		Prov("mss").Conn("par", godip.Land).Conn("bor", godip.Land).Conn("cat", godip.Coast...).Conn("guo", godip.Sea).Flag(godip.Coast...).SC(France).
		// Louisiana
		Prov("lou").Conn("neo", godip.Land).Conn("stl", godip.Land).Conn("upp", godip.Land).Conn("wrl", godip.Land).Conn("ore", godip.Land).Conn("cal", godip.Land).Conn("saf", godip.Land).Flag(godip.Land).
		// Baeleric Sea
		Prov("bae").Conn("val", godip.Sea).Conn("str", godip.Sea).Conn("alg", godip.Sea).Conn("tun", godip.Sea).Conn("guo", godip.Sea).Flag(godip.Sea).
		// Canaries Sea
		Prov("cse").Conn("azo", godip.Sea).Conn("wea", godip.Sea).Conn("mia", godip.Sea).Conn("cav", godip.Sea).Conn("wsa", godip.Sea).Conn("cai", godip.Sea).Conn("mad", godip.Sea).Conn("eur", godip.Sea).Flag(godip.Sea).
		// Caracas
		Prov("cca").Conn("bog", godip.Land).Conn("sur", godip.Land).Conn("trn", godip.Land).Conn("win", godip.Coast...).Conn("eac", godip.Sea).Conn("neg", godip.Coast...).Flag(godip.Coast...).
		// Bahamas
		Prov("bah").Conn("cub", godip.Coast...).Conn("bet", godip.Sea).Conn("flo", godip.Coast...).Conn("gom", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Buenos Aires
		Prov("bue").Conn("rio", godip.Coast...).Conn("cen", godip.Land).Conn("sat", godip.Land).Conn("arg", godip.Coast...).Conn("dra", godip.Sea).Conn("coa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Santiago
		Prov("sat").Conn("saa", godip.Land).Conn("lim", godip.Land).Conn("arg", godip.Land).Conn("bue", godip.Land).Conn("cen", godip.Land).Flag(godip.Land).SC(Spain).
		// Strait of Gibraltar
		Prov("str").Conn("bae", godip.Sea).Conn("val", godip.Sea).Conn("gib", godip.Sea).Conn("eur", godip.Sea).Conn("maz", godip.Sea).Conn("nom", godip.Sea).Conn("alg", godip.Sea).Flag(godip.Sea).
		// Iroquois Territory
		Prov("iro").Conn("nee", godip.Land).Conn("nov", godip.Coast...).Conn("now", godip.Sea).Conn("mon", godip.Coast...).Conn("ont", godip.Land).Conn("det", godip.Land).Conn("app", godip.Land).Conn("pen", godip.Land).Flag(godip.Coast...).
		// Surinam
		Prov("sur").Conn("bog", godip.Land).Conn("cay", godip.Coast...).Conn("caa", godip.Sea).Conn("trn", godip.Coast...).Conn("cca", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// East Rupert's Land
		Prov("erl").Conn("hud", godip.Sea).Conn("moo", godip.Coast...).Conn("ont", godip.Land).Conn("mon", godip.Coast...).Flag(godip.Coast...).
		// Southern Ocean
		Prov("soo").Conn("cog", godip.Sea).Conn("nam", godip.Sea).Conn("sea", godip.Sea).Conn("soa", godip.Sea).Conn("coa", godip.Sea).Conn("dra", godip.Sea).Conn("fal", godip.Sea).Flag(godip.Sea).
		// East Sahara
		Prov("esa").Conn("trp", godip.Land).Conn("tun", godip.Land).Conn("mis", godip.Land).Conn("sud", godip.Land).Flag(godip.Land).
		// Appalachian Mountains
		Prov("app").Conn("stl", godip.Land).Conn("neo", godip.Land).Conn("geo", godip.Land).Conn("pen", godip.Land).Conn("iro", godip.Land).Conn("det", godip.Land).Flag(godip.Land).
		// North Morocco
		Prov("nom").Conn("maz", godip.Coast...).Conn("som", godip.Land).Conn("mis", godip.Land).Conn("alg", godip.Coast...).Conn("str", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Northern Atlantic
		Prov("ote").Conn("now", godip.Sea).Conn("wea", godip.Sea).Conn("eur", godip.Sea).Conn("iri", godip.Sea).Conn("ire", godip.Sea).Conn("iri", godip.Sea).Conn("sco", godip.Sea).Conn("lon", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Sea).
		// Nova Scotia
		Prov("nov").Conn("now", godip.Sea).Conn("iro", godip.Coast...).Conn("nee", godip.Coast...).Conn("ese", godip.Sea).Conn("now", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Tripoli
		Prov("trp").Conn("ion", godip.Sea).Conn("guo", godip.Sea).Conn("tun", godip.Coast...).Conn("esa", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Soyo
		Prov("soy").Conn("ang", godip.Coast...).Conn("kon", godip.Coast...).Conn("sea", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Falkland Islands
		Prov("fal").Conn("soo", godip.Sea).Conn("dra", godip.Sea).Flag(godip.Coast...).
		// South Morocco
		Prov("som").Conn("wsa", godip.Coast...).Conn("mis", godip.Land).Conn("nom", godip.Land).Conn("maz", godip.Coast...).Conn("eur", godip.Sea).Conn("cai", godip.Coast...).Flag(godip.Coast...).
		// Hispaniola
		Prov("his").Conn("eac", godip.Sea).Conn("lee", godip.Coast...).Conn("bet", godip.Sea).Conn("bet", godip.Sea).Conn("jam", godip.Coast...).Flag(godip.Coast...).
		// New Orleans
		Prov("neo").Conn("stl", godip.Land).Conn("lou", godip.Land).Conn("saf", godip.Coast...).Conn("gom", godip.Sea).Conn("geo", godip.Land).Conn("geo/sc", godip.Sea).Conn("app", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Belize
		Prov("bei").Conn("sie", godip.Land).Conn("neg", godip.Coast...).Conn("wec", godip.Sea).Conn("mer", godip.Coast...).Flag(godip.Coast...).
		// Bristol
		Prov("bri").Conn("lon", godip.Coast...).Conn("sco", godip.Coast...).Conn("iri", godip.Sea).Conn("eng", godip.Sea).Flag(godip.Coast...).SC(England).
		// Jamaica
		Prov("jam").Conn("cub", godip.Coast...).Conn("wec", godip.Sea).Conn("eac", godip.Sea).Conn("his", godip.Coast...).Conn("bet", godip.Sea).Flag(godip.Coast...).SC(England).
		// Moose Fort
		Prov("moo").Conn("the", godip.Coast...).Conn("wrl", godip.Land).Conn("ont", godip.Land).Conn("erl", godip.Coast...).Conn("hud", godip.Sea).Flag(godip.Coast...).SC(England).
		// Barbados
		Prov("bar").Conn("trn", godip.Coast...).Conn("caa", godip.Sea).Conn("win", godip.Coast...).Flag(godip.Coast...).
		// Southwestern Atlantic
		Prov("swa").Conn("cav", godip.Sea).Conn("caa", godip.Sea).Conn("eem", godip.Sea).Conn("mnh", godip.Sea).Conn("sao", godip.Sea).Conn("soa", godip.Sea).Conn("gre", godip.Sea).Flag(godip.Sea).
		// Goree
		Prov("gor").Conn("gra", godip.Coast...).Conn("sen", godip.Coast...).Conn("cav", godip.Coast...).Conn("gre", godip.Sea).Flag(godip.Coast...).SC(France).
		// West Rupert's Land
		Prov("wrl").Conn("ore", godip.Land).Conn("lou", godip.Land).Conn("upp", godip.Land).Conn("ont", godip.Land).Conn("moo", godip.Land).Conn("the", godip.Land).Flag(godip.Land).
		// Slave Coast
		Prov("sla").Conn("sud", godip.Land).Conn("nig", godip.Land).Conn("gol", godip.Coast...).Conn("gre", godip.Sea).Conn("sea", godip.Sea).Conn("kon", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Tunisia
		Prov("tun").Conn("mis", godip.Land).Conn("esa", godip.Land).Conn("trp", godip.Coast...).Conn("bae", godip.Sea).Conn("guo", godip.Sea).Conn("guo", godip.Sea).Conn("alg", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Sea of Angola
		Prov("sea").Conn("gre", godip.Sea).Conn("soa", godip.Sea).Conn("soo", godip.Sea).Conn("nam", godip.Sea).Conn("ang", godip.Sea).Conn("soy", godip.Sea).Conn("kon", godip.Sea).Conn("sla", godip.Sea).Flag(godip.Sea).
		// Coast of Brazil
		Prov("coa").Conn("soa", godip.Sea).Conn("sao", godip.Sea).Conn("rio", godip.Sea).Conn("bue", godip.Sea).Conn("dra", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Sea).
		// Mexico City
		Prov("mex").Conn("mer", godip.Coast...).Conn("gom", godip.Sea).Conn("saf", godip.Coast...).Conn("sie", godip.Land).Flag(godip.Coast...).SC(Spain).
		// Mazagan
		Prov("maz").Conn("som", godip.Coast...).Conn("nom", godip.Coast...).Conn("str", godip.Sea).Conn("eur", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// Leeward Islands
		Prov("lee").Conn("eac", godip.Sea).Conn("win", godip.Coast...).Conn("ant", godip.Coast...).Conn("mia", godip.Sea).Conn("bet", godip.Sea).Conn("his", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Lion
		Prov("guo").Conn("cat", godip.Sea).Conn("bae", godip.Sea).Conn("tun", godip.Sea).Conn("tun", godip.Sea).Conn("trp", godip.Sea).Conn("ion", godip.Sea).Conn("mss", godip.Sea).Flag(godip.Sea).
		// New England
		Prov("nee").Conn("iro", godip.Land).Conn("pen", godip.Coast...).Conn("ese", godip.Sea).Conn("nov", godip.Coast...).Flag(godip.Coast...).SC(England).
		// Angola
		Prov("ang").Conn("soy", godip.Coast...).Conn("sea", godip.Sea).Conn("nam", godip.Coast...).Conn("kon", godip.Coast...).Flag(godip.Coast...).SC(Portugal).
		// Trinidad
		Prov("trn").Conn("bar", godip.Coast...).Conn("win", godip.Coast...).Conn("cca", godip.Land).Conn("sur", godip.Coast...).Conn("caa", godip.Sea).Flag(godip.Coast...).
		// Irish Sea
		Prov("iri").Conn("ote", godip.Sea).Conn("eur", godip.Sea).Conn("eng", godip.Sea).Conn("bri", godip.Sea).Conn("sco", godip.Sea).Conn("ote", godip.Sea).Conn("ire", godip.Sea).Conn("ire", godip.Sea).Flag(godip.Sea).
		// Castilla y Leon
		Prov("cas").Conn("eur", godip.Sea).Conn("por", godip.Coast...).Conn("and", godip.Coast...).Conn("val", godip.Land).Conn("cat", godip.Land).Conn("bor", godip.Coast...).Flag(godip.Coast...).
		// European Atlantic
		Prov("eur").Conn("cse", godip.Sea).Conn("mad", godip.Sea).Conn("cai", godip.Sea).Conn("som", godip.Sea).Conn("maz", godip.Sea).Conn("str", godip.Sea).Conn("and", godip.Sea).Conn("por", godip.Sea).Conn("cas", godip.Sea).Conn("bor", godip.Sea).Conn("bre", godip.Sea).Conn("eng", godip.Sea).Conn("iri", godip.Sea).Conn("ote", godip.Sea).Conn("wea", godip.Sea).Conn("azo", godip.Sea).Flag(godip.Sea).
		// Mid Atlantic
		Prov("mia").Conn("bet", godip.Sea).Conn("lee", godip.Sea).Conn("ant", godip.Sea).Conn("caa", godip.Sea).Conn("cav", godip.Sea).Conn("cse", godip.Sea).Conn("wea", godip.Sea).Conn("beu", godip.Sea).Flag(godip.Sea).
		// California
		Prov("cal").Conn("saf", godip.Land).Conn("lou", godip.Land).Conn("ore", godip.Land).Conn("sie", godip.Land).Flag(godip.Land).
		// Sierra Madre
		Prov("sie").Conn("saf", godip.Land).Conn("cal", godip.Land).Conn("bei", godip.Land).Conn("mer", godip.Land).Conn("mex", godip.Land).Flag(godip.Land).
		// Montreal
		Prov("mon").Conn("nef", godip.Coast...).Conn("now", godip.Sea).Conn("hud", godip.Sea).Conn("erl", godip.Coast...).Conn("ont", godip.Land).Conn("iro", godip.Coast...).Conn("now", godip.Sea).Flag(godip.Coast...).SC(France).
		// Kongo
		Prov("kon").Conn("sla", godip.Coast...).Conn("sea", godip.Sea).Conn("soy", godip.Coast...).Conn("ang", godip.Coast...).Conn("nam", godip.Land).Conn("eaa", godip.Land).Conn("sud", godip.Land).Flag(godip.Coast...).
		// Portugal
		Prov("por").Conn("and", godip.Coast...).Conn("cas", godip.Coast...).Conn("eur", godip.Sea).Flag(godip.Coast...).SC(Portugal).
		// Georgia
		Prov("geo").Conn("pen", godip.Land).Conn("app", godip.Land).Conn("neo", godip.Land).Conn("flo", godip.Land).Flag(godip.Land).SC(England).
		// Georgia (EC)
		Prov("geo/ec").Conn("ese", godip.Sea).Conn("pen", godip.Sea).Conn("flo", godip.Sea).Conn("bet", godip.Sea).Flag(godip.Sea).
		// Georgia (SC)
		Prov("geo/sc").Conn("neo", godip.Sea).Conn("gom", godip.Sea).Conn("flo", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"upp":    "Upper Mississippi",
	"bog":    "Bogota",
	"bet":    "Bermuda Triangle",
	"sco":    "Scotland",
	"eaa":    "East Africa",
	"cub":    "Cuba",
	"flo":    "Florida",
	"det":    "Detroit",
	"lon":    "London",
	"azo":    "Azores Islands",
	"wea":    "Western Atlantic",
	"soa":    "Southern Atlantic",
	"nef":    "Newfoundland",
	"ont":    "Ontario",
	"hud":    "Hudson Bay",
	"caa":    "Caribbean Atlantic",
	"gol":    "Gold Coast",
	"bre":    "Brest",
	"sao":    "Sao Salvador",
	"par":    "Paris",
	"ion":    "Ionian Sea",
	"gom":    "Gulf of Mexico",
	"cay":    "Cayenne",
	"gra":    "Grain Coast",
	"win":    "Windward Islands",
	"now":    "Northwestern Atlantic",
	"beu":    "Bermuda",
	"ese":    "Eastern Seaboard",
	"stl":    "St. Louis",
	"gib":    "Gibraltar",
	"mer":    "Merida",
	"neg":    "New Granada",
	"cai":    "Canary Islands",
	"pen":    "Pennsylvania",
	"mnh":    "Maranhao",
	"eac":    "Eastern Caribbean",
	"alg":    "Algeria",
	"dra":    "Drake's Passage",
	"mad":    "Madeira",
	"arg":    "Argentina",
	"gre":    "Greater Gulf of Guinea",
	"wsa":    "West Sahara",
	"nig":    "Nigrita",
	"the":    "The North West",
	"goy":    "Goyaz",
	"saf":    "Santa Fe",
	"lim":    "Lima",
	"ire":    "Ireland",
	"cog":    "Cape of Good Hope",
	"sud":    "Sudan",
	"mis":    "Middle Sahara",
	"nam":    "Namib Desert",
	"val":    "Valencia",
	"cen":    "Central Brazil",
	"bor":    "Bordeaux",
	"eng":    "English Channel",
	"cav":    "Cape Verde",
	"wec":    "Western Caribbean",
	"ore":    "Oregon",
	"ant":    "Antigua",
	"and":    "Andalucia",
	"cat":    "Catalan",
	"sen":    "Senegal",
	"rio":    "Rio de Janeiro",
	"saa":    "Santa Anna",
	"eem":    "Belem",
	"mss":    "Marseilles",
	"lou":    "Louisiana",
	"bae":    "Baeleric Sea",
	"cse":    "Canaries Sea",
	"cca":    "Caracas",
	"bah":    "Bahamas",
	"bue":    "Buenos Aires",
	"sat":    "Santiago",
	"str":    "Strait of Gibraltar",
	"iro":    "Iroquois Territory",
	"sur":    "Surinam",
	"erl":    "East Rupert's Land",
	"soo":    "Southern Ocean",
	"esa":    "East Sahara",
	"app":    "Appalachian Mountains",
	"nom":    "North Morocco",
	"ote":    "Northern Atlantic",
	"nov":    "Nova Scotia",
	"trp":    "Tripoli",
	"soy":    "Soyo",
	"fal":    "Falkland Islands",
	"geo":    "Georgia",
	"geo/ec": "Georgia (EC)",
	"geo/sc": "Georgia (SC)",
	"som":    "South Morocco",
	"his":    "Hispaniola",
	"neo":    "New Orleans",
	"bei":    "Belize",
	"bri":    "Bristol",
	"jam":    "Jamaica",
	"moo":    "Moose Fort",
	"bar":    "Barbados",
	"swa":    "Southwestern Atlantic",
	"gor":    "Goree",
	"wrl":    "West Rupert's Land",
	"sla":    "Slave Coast",
	"tun":    "Tunisia",
	"sea":    "Sea of Angola",
	"coa":    "Coast of Brazil",
	"mex":    "Mexico City",
	"maz":    "Mazagan",
	"lee":    "Leeward Islands",
	"guo":    "Gulf of Lion",
	"nee":    "New England",
	"ang":    "Angola",
	"trn":    "Trinidad",
	"iri":    "Irish Sea",
	"cas":    "Castilla y Leon",
	"eur":    "European Atlantic",
	"mia":    "Mid Atlantic",
	"cal":    "California",
	"sie":    "Sierra Madre",
	"mon":    "Montreal",
	"kon":    "Kongo",
	"por":    "Portugal",
}
