package northseawars

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Britons  godip.Nation = "Britons"
	Romans   godip.Nation = "Romans"
	Frysians godip.Nation = "Frysians"
	Norse    godip.Nation = "Norse"
)

var Nations = []godip.Nation{Britons, Romans, Frysians, Norse}

var NorthSeaWarsVariant = common.Variant{
	Name:   "North Sea Wars",
	Graph:  func() godip.Graph { return NorthSeaWarsGraph() },
	Start:  NorthSeaWarsStart,
	Blank:  NorthSeaWarsBlank,
	Phase:  classical.NewPhase,
	Parser: classical.Parser,
	ExtraDominanceRules: map[godip.Province]common.DominanceRule{
		"nbr": common.DominanceRule{
			Nation: Britons,
			Dependencies: map[godip.Province]godip.Nation{
				"cym": Britons,
				"sbr": Britons,
			},
		},
	},
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(8),
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/northseawarsmap.svg")
	},
	SVGVersion: "3",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "sqrg",
	Version:     "1",
	Description: "A battle for trade routes in the North Sea.",
	SoloSCCount: func(*state.State) int { return 8 },
	Rules: `First to 8 Supply Centers (SC) is the winner.
	Units can move from Central North Sea to three trade provinces containing SCs – Wood, Iron and Grain. Units in the trade provinces can move freely between them, but can’t return back to Central North Sea.
	Jutland has a dual coast.
	Sealand has land access to all neighbouring spaces (including Limfjorden) and naval access to Jutland (East Coast), but not Amsivaria.`,
}

func NorthSeaWarsBlank(phase godip.Phase) *state.State {
	return state.New(NorthSeaWarsGraph(), phase, classical.BackupRule, nil, nil)
}

func NorthSeaWarsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(0, godip.Spring, godip.Movement)
	result = NorthSeaWarsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"sbr": godip.Unit{godip.Fleet, Britons},
		"cym": godip.Unit{godip.Army, Britons},
		"men": godip.Unit{godip.Fleet, Romans},
		"ges": godip.Unit{godip.Army, Romans},
		"fri": godip.Unit{godip.Fleet, Frysians},
		"ams": godip.Unit{godip.Army, Frysians},
		"sor": godip.Unit{godip.Fleet, Norse},
		"got": godip.Unit{godip.Army, Norse},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"sbr": Britons,
		"cym": Britons,
		"men": Romans,
		"ges": Romans,
		"fri": Frysians,
		"ams": Frysians,
		"sor": Norse,
		"got": Norse,
	})
	return
}

func NorthSeaWarsGraph() *graph.Graph {
	return graph.New().
		// Iron
		Prov("iro").Conn("woo", godip.Coast...).Conn("gra", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Albion
		Prov("ali").Conn("cha", godip.Sea).Conn("sbr", godip.Coast...).Conn("cym", godip.Land).Flag(godip.Coast...).
		// Upper North Sea
		Prov("uns").Conn("ala", godip.Sea).Conn("wns", godip.Sea).Conn("cns", godip.Sea).Conn("ens", godip.Sea).Conn("sor", godip.Sea).Conn("ves", godip.Sea).Flag(godip.Sea).
		// West Belgica
		Prov("wbe").Conn("ebe", godip.Land).Conn("men", godip.Coast...).Conn("cha", godip.Sea).Flag(godip.Coast...).
		// Cymru
		Prov("cym").Conn("nbr", godip.Land).Conn("ali", godip.Land).Conn("sbr", godip.Land).Flag(godip.Land).SC(Britons).
		// Jutland
		Prov("jut").Conn("lim", godip.Land).Conn("ams", godip.Land).Conn("sea", godip.Land).Flag(godip.Land).
		// Jutland (West Coast)
		Prov("jut/wc").Conn("lim", godip.Sea).Conn("ens", godip.Sea).Conn("ams", godip.Sea).Flag(godip.Sea).
		// Jutland (East Coast)
		Prov("jut/ec").Conn("lim", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Sea).
		// Germania Inferior
		Prov("gei").Conn("fri", godip.Land).Conn("bat", godip.Land).Conn("men", godip.Land).Conn("ebe", godip.Land).Conn("ges", godip.Land).Conn("mag", godip.Land).Conn("ams", godip.Land).Flag(godip.Land).
		// Germania Superior
		Prov("ges").Conn("mag", godip.Land).Conn("gei", godip.Land).Conn("ebe", godip.Land).Flag(godip.Land).SC(Romans).
		// Alba
		Prov("ala").Conn("nbr", godip.Coast...).Conn("wns", godip.Sea).Conn("uns", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Channel
		Prov("cha").Conn("wbe", godip.Sea).Conn("men", godip.Sea).Conn("lns", godip.Sea).Conn("sbr", godip.Sea).Conn("ali", godip.Sea).Flag(godip.Sea).
		// South Britanny
		Prov("sbr").Conn("lns", godip.Sea).Conn("wns", godip.Sea).Conn("nbr", godip.Coast...).Conn("cym", godip.Land).Conn("ali", godip.Coast...).Conn("cha", godip.Sea).Flag(godip.Coast...).SC(Britons).
		// Grains
		Prov("gra").Conn("woo", godip.Coast...).Conn("iro", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Lower North Sea
		Prov("lns").Conn("bat", godip.Sea).Conn("fri", godip.Sea).Conn("ens", godip.Sea).Conn("cns", godip.Sea).Conn("wns", godip.Sea).Conn("sbr", godip.Sea).Conn("cha", godip.Sea).Conn("men", godip.Sea).Flag(godip.Sea).
		// Menapia
		Prov("men").Conn("cha", godip.Sea).Conn("wbe", godip.Coast...).Conn("ebe", godip.Land).Conn("gei", godip.Land).Conn("bat", godip.Coast...).Conn("lns", godip.Sea).Flag(godip.Coast...).SC(Romans).
		// Frisia
		Prov("fri").Conn("lns", godip.Sea).Conn("bat", godip.Coast...).Conn("gei", godip.Land).Conn("ams", godip.Coast...).Conn("ens", godip.Sea).Flag(godip.Coast...).SC(Frysians).
		// Central North Sea
		Prov("cns").Conn("lns", godip.Sea).Conn("ens", godip.Sea).Conn("uns", godip.Sea).Conn("wns", godip.Sea).Conn("woo", godip.Sea).Conn("iro", godip.Sea).Conn("gra", godip.Sea).Flag(godip.Sea).
		// Vestland
		Prov("ves").Conn("uns", godip.Sea).Conn("sor", godip.Coast...).Conn("ost", godip.Land).Flag(godip.Coast...).
		// Magna Germania
		Prov("mag").Conn("sea", godip.Land).Conn("ams", godip.Land).Conn("gei", godip.Land).Conn("ges", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ostland
		Prov("ost").Conn("ves", godip.Land).Conn("sor", godip.Coast...).Conn("ska", godip.Sea).Conn("got", godip.Coast...).Flag(godip.Coast...).
		// Sealand
		Prov("sea").Conn("got", godip.Coast...).Conn("ska", godip.Sea).Conn("lim", godip.Coast...).Conn("jut", godip.Land).Conn("jut/ec", godip.Sea).Conn("ams", godip.Land).Conn("mag", godip.Land).Flag(godip.Coast...).
		// East Belgica
		Prov("ebe").Conn("ges", godip.Land).Conn("gei", godip.Land).Conn("men", godip.Land).Conn("wbe", godip.Land).Flag(godip.Land).
		// North Britanny
		Prov("nbr").Conn("cym", godip.Land).Conn("sbr", godip.Coast...).Conn("wns", godip.Sea).Conn("ala", godip.Coast...).Flag(godip.Coast...).
		// Skagerrak
		Prov("ska").Conn("got", godip.Sea).Conn("ost", godip.Sea).Conn("sor", godip.Sea).Conn("ens", godip.Sea).Conn("lim", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Sea).
		// West North Sea
		Prov("wns").Conn("uns", godip.Sea).Conn("ala", godip.Sea).Conn("nbr", godip.Sea).Conn("sbr", godip.Sea).Conn("lns", godip.Sea).Conn("cns", godip.Sea).Flag(godip.Sea).
		// Amsivaria
		Prov("ams").Conn("jut", godip.Land).Conn("jut/wc", godip.Sea).Conn("ens", godip.Sea).Conn("fri", godip.Coast...).Conn("gei", godip.Land).Conn("mag", godip.Land).Conn("sea", godip.Land).Flag(godip.Coast...).SC(Frysians).
		// Sorland
		Prov("sor").Conn("uns", godip.Sea).Conn("ens", godip.Sea).Conn("ska", godip.Sea).Conn("ost", godip.Coast...).Conn("ves", godip.Coast...).Flag(godip.Coast...).SC(Norse).
		// Limfjorden
		Prov("lim").Conn("sea", godip.Coast...).Conn("ska", godip.Sea).Conn("ens", godip.Sea).Conn("jut", godip.Land).Conn("jut/wc", godip.Sea).Conn("jut/ec", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// East North Sea
		Prov("ens").Conn("sor", godip.Sea).Conn("uns", godip.Sea).Conn("cns", godip.Sea).Conn("lns", godip.Sea).Conn("fri", godip.Sea).Conn("ams", godip.Sea).Conn("jut", godip.Sea).Conn("jut/wc", godip.Sea).Conn("lim", godip.Sea).Conn("ska", godip.Sea).Flag(godip.Sea).
		// Batavia
		Prov("bat").Conn("lns", godip.Sea).Conn("men", godip.Coast...).Conn("gei", godip.Land).Conn("fri", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Gotaland
		Prov("got").Conn("ost", godip.Coast...).Conn("ska", godip.Sea).Conn("sea", godip.Coast...).Flag(godip.Coast...).SC(Norse).
		// Wood
		Prov("woo").Conn("gra", godip.Coast...).Conn("iro", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"ebe":    "East Belgica",
	"jut":    "Jutland",
	"ali":    "Albion",
	"nbr":    "North Britanny",
	"ost":    "Ostland",
	"lns":    "Lower North Sea",
	"ala":    "Alba",
	"sor":    "Sorland",
	"sea":    "Sealand",
	"sbr":    "South Britanny",
	"ams":    "Amsivaria",
	"cha":    "Channel",
	"cns":    "Central North Sea",
	"lim":    "Limfjorden",
	"fri":    "Frisia",
	"woo":    "Wood (Trade)",
	"ens":    "East North Sea",
	"ska":    "Skagerrak",
	"gra":    "Grain (Trade)",
	"got":    "Gotaland",
	"ges":    "Germania Superior",
	"men":    "Menapia",
	"iro":    "Iron (Trade)",
	"wns":    "West North Sea",
	"mag":    "Magna Germania",
	"gei":    "Germania Inferior",
	"cym":    "Cymru",
	"ves":    "Vestland",
	"bat":    "Batavia",
	"uns":    "Upper North Sea",
	"wbe":    "West Belgica",
	"jut/ec": "Jutland (EC)",
	"jut/wc": "Jutland (WC)",
}
