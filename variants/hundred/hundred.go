package hundred

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	England  godip.Nation = "England"
	Burgundy godip.Nation = "Burgundy"
	France   godip.Nation = "France"
)

var Nations = []godip.Nation{England, Burgundy, France}

var BuildAnywhereParser = orders.NewParser([]godip.Order{
	orders.MoveOrder,
	orders.MoveViaConvoyOrder,
	orders.HoldOrder,
	orders.SupportOrder,
	orders.BuildAnywhereOrder,
	orders.DisbandOrder,
	orders.ConvoyOrder,
})

var HundredVariant = common.Variant{
	Name:   "Hundred",
	Graph:  func() godip.Graph { return HundredGraph() },
	Start:  HundredStart,
	Blank:  HundredBlank,
	Phase:  Phase,
	Parser: BuildAnywhereParser,
	ExtraDominanceRules: map[godip.Province]common.DominanceRule{
		"not": common.DominanceRule{
			Nation: England,
			Dependencies: map[godip.Province]godip.Nation{
				"dev": England,
				"sco": godip.Neutral,
			},
		},
		"anj": common.DominanceRule{
			Nation: France,
			Dependencies: map[godip.Province]godip.Nation{
				"orl": France,
				"nom": England,
				"brt": godip.Neutral,
			},
		},
		"poi": common.DominanceRule{
			Nation: France,
			Dependencies: map[godip.Province]godip.Nation{
				"tou": France,
				"orl": France,
				"guy": England,
			},
		},
		"cha": common.DominanceRule{
			Nation: Burgundy,
			Dependencies: map[godip.Province]godip.Nation{
				"dij": Burgundy,
				"par": France,
				"dau": France,
			},
		},
	},
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           []godip.Season{YearSeason},
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(9),
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/hundredmap.svg")
	},
	SVGVersion: "6",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	SVGFlags: map[godip.Nation]func() ([]byte, error){
		England: func() ([]byte, error) {
			return Asset("svg/England.svg")
		},
		Burgundy: func() ([]byte, error) {
			return Asset("svg/Burgundy.svg")
		},
		France: func() ([]byte, error) {
			return Asset("svg/France.svg")
		},
	},
	CreatedBy:   "Andy Schwarz",
	Version:     "3",
	Description: "The Hundred Years' War was a series of conflicts from 1337 to 1453, waged between the English House of Plantagenet and the French House of Valois, over the Kingdom of France.",
	SoloSCCount: func(*state.State) int { return 9 },
	Rules: `First to 9 Supply Centers (SC) is the winner.
	Units may be built on any owned SC.
	France starts with 5 units and 4 SC, so needs to disband unless they gain a center before 1430. 
	Yearly cycles of Spring & Fall are renamed to decade cycles with years ending with 5 and 0 respectively.
	Armies & fleets can move between London & Calais.
	Two provinces have dual coasts: Northumbria and Aragon.`,
}

func HundredBlank(phase godip.Phase) *state.State {
	return state.New(HundredGraph(), phase, classical.BackupRule, map[godip.Flag]bool{godip.Anywhere: true}, nil)
}

func HundredStart() (result *state.State, err error) {
	startPhase := Phase(1425, YearSeason, godip.Movement)
	result = HundredBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"lon": godip.Unit{godip.Fleet, England},
		"dev": godip.Unit{godip.Fleet, England},
		"cal": godip.Unit{godip.Army, England},
		"guy": godip.Unit{godip.Army, England},
		"nom": godip.Unit{godip.Army, England},
		"hol": godip.Unit{godip.Fleet, Burgundy},
		"dij": godip.Unit{godip.Army, Burgundy},
		"lux": godip.Unit{godip.Army, Burgundy},
		"fla": godip.Unit{godip.Army, Burgundy},
		"dau": godip.Unit{godip.Army, France},
		"orl": godip.Unit{godip.Army, France},
		"par": godip.Unit{godip.Army, France},
		"tou": godip.Unit{godip.Army, France},
		"pro": godip.Unit{godip.Army, France},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"lon": England,
		"dev": England,
		"cal": England,
		"guy": England,
		"nom": England,
		"hol": Burgundy,
		"dij": Burgundy,
		"lux": Burgundy,
		"fla": Burgundy,
		"dau": France,
		"orl": France,
		"par": France,
		"tou": France,
	})
	return
}

func HundredGraph() *graph.Graph {
	return graph.New().
		// Atlantic Sea
		Prov("atl").Conn("med", godip.Sea).Conn("cas", godip.Sea).Conn("bis", godip.Sea).Conn("brs", godip.Sea).Conn("iri", godip.Sea).Conn("thp", godip.Sea).Flag(godip.Sea).
		// Normandy
		Prov("nom").Conn("str", godip.Sea).Conn("eng", godip.Sea).Conn("brt", godip.Coast...).Conn("anj", godip.Land).Conn("orl", godip.Land).Conn("par", godip.Land).Conn("cal", godip.Coast...).Flag(godip.Coast...).SC(England).
		// Dauphine
		Prov("dau").Conn("orl", godip.Land).Conn("lim", godip.Land).Conn("pro", godip.Land).Conn("sav", godip.Land).Conn("can", godip.Land).Conn("dij", godip.Land).Conn("cha", godip.Land).Conn("par", godip.Land).Flag(godip.Land).SC(France).
		// Anjou
		Prov("anj").Conn("orl", godip.Land).Conn("nom", godip.Land).Conn("brt", godip.Land).Flag(godip.Land).
		// Guyenne
		Prov("guy").Conn("ara", godip.Land).Conn("ara/nc", godip.Sea).Conn("tou", godip.Land).Conn("poi", godip.Land).Conn("brt", godip.Coast...).Conn("bis", godip.Sea).Flag(godip.Coast...).SC(England).
		// Friesland
		Prov("fri").Conn("lux", godip.Land).Conn("thw", godip.Sea).Conn("hol", godip.Coast...).Flag(godip.Coast...).
		// North Sea
		Prov("nos").Conn("sco", godip.Sea).Conn("not", godip.Sea).Conn("not/ec", godip.Sea).Conn("ang", godip.Sea).Conn("thw", godip.Sea).Conn("min", godip.Sea).Flag(godip.Sea).
		// The Wash
		Prov("thw").Conn("nos", godip.Sea).Conn("ang", godip.Sea).Conn("str", godip.Sea).Conn("hol", godip.Sea).Conn("fri", godip.Sea).Flag(godip.Sea).
		// Devon
		Prov("dev").Conn("wal", godip.Coast...).Conn("brs", godip.Sea).Conn("eng", godip.Sea).Conn("lon", godip.Coast...).Conn("ang", godip.Land).Conn("not", godip.Land).Flag(godip.Coast...).SC(England).
		// London
		Prov("lon").Conn("cal", godip.Coast...).Conn("dev", godip.Coast...).Conn("eng", godip.Sea).Conn("str", godip.Sea).Conn("ang", godip.Coast...).Flag(godip.Coast...).SC(England).
		// Calais
		Prov("cal").Conn("lon", godip.Coast...).Conn("str", godip.Sea).Conn("nom", godip.Coast...).Conn("par", godip.Land).Conn("dij", godip.Land).Conn("fla", godip.Coast...).Flag(godip.Coast...).SC(England).
		// Alsace
		Prov("als").Conn("lor", godip.Land).Conn("can", godip.Land).Flag(godip.Land).
		// Poitou
		Prov("poi").Conn("lim", godip.Land).Conn("orl", godip.Land).Conn("brt", godip.Land).Conn("guy", godip.Land).Conn("tou", godip.Land).Flag(godip.Land).
		// Biscay
		Prov("bis").Conn("cas", godip.Sea).Conn("ara", godip.Sea).Conn("ara/nc", godip.Sea).Conn("guy", godip.Sea).Conn("brt", godip.Sea).Conn("brs", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Savoy
		Prov("sav").Conn("can", godip.Land).Conn("dau", godip.Land).Conn("pro", godip.Coast...).Conn("med", godip.Sea).Flag(godip.Coast...).
		// Orleanais
		Prov("orl").Conn("dau", godip.Land).Conn("par", godip.Land).Conn("nom", godip.Land).Conn("anj", godip.Land).Conn("brt", godip.Land).Conn("poi", godip.Land).Conn("lim", godip.Land).Flag(godip.Land).SC(France).
		// Strait of Dover
		Prov("str").Conn("cal", godip.Sea).Conn("fla", godip.Sea).Conn("hol", godip.Sea).Conn("thw", godip.Sea).Conn("ang", godip.Sea).Conn("lon", godip.Sea).Conn("eng", godip.Sea).Conn("nom", godip.Sea).Flag(godip.Sea).
		// Mediterranean
		Prov("med").Conn("sav", godip.Sea).Conn("pro", godip.Sea).Conn("tou", godip.Sea).Conn("ara", godip.Sea).Conn("ara/sc", godip.Sea).Conn("cas", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Lorraine
		Prov("lor").Conn("lux", godip.Land).Conn("dij", godip.Land).Conn("can", godip.Land).Conn("als", godip.Land).Flag(godip.Land).
		// Flanders
		Prov("fla").Conn("str", godip.Sea).Conn("cal", godip.Coast...).Conn("dij", godip.Land).Conn("lux", godip.Land).Conn("hol", godip.Coast...).Flag(godip.Coast...).SC(Burgundy).
		// Bristol Channel
		Prov("brs").Conn("wal", godip.Sea).Conn("iri", godip.Sea).Conn("atl", godip.Sea).Conn("bis", godip.Sea).Conn("brt", godip.Sea).Conn("eng", godip.Sea).Conn("dev", godip.Sea).Flag(godip.Sea).
		// Cantons
		Prov("can").Conn("als", godip.Land).Conn("lor", godip.Land).Conn("dij", godip.Land).Conn("dau", godip.Land).Conn("sav", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Northumbria
		Prov("not").Conn("wal", godip.Land).Conn("dev", godip.Land).Conn("ang", godip.Land).Conn("sco", godip.Land).Flag(godip.Land).
		// Northumbria (West Coast)
		Prov("not/wc").Conn("wal", godip.Sea).Conn("sco", godip.Sea).Conn("iri", godip.Sea).Flag(godip.Sea).
		// Northumbria (East Coast)
		Prov("not/ec").Conn("ang", godip.Sea).Conn("nos", godip.Sea).Conn("sco", godip.Sea).Flag(godip.Sea).
		// Provence
		Prov("pro").Conn("lim", godip.Land).Conn("tou", godip.Coast...).Conn("med", godip.Sea).Conn("sav", godip.Coast...).Conn("dau", godip.Land).Flag(godip.Coast...).
		// Paris
		Prov("par").Conn("orl", godip.Land).Conn("dau", godip.Land).Conn("cha", godip.Land).Conn("dij", godip.Land).Conn("cal", godip.Land).Conn("nom", godip.Land).Flag(godip.Land).SC(France).
		// Toulouse
		Prov("tou").Conn("med", godip.Sea).Conn("pro", godip.Coast...).Conn("lim", godip.Land).Conn("poi", godip.Land).Conn("guy", godip.Land).Conn("ara", godip.Land).Conn("ara/sc", godip.Sea).Flag(godip.Coast...).SC(France).
		// Irish Sea
		Prov("iri").Conn("thp", godip.Sea).Conn("atl", godip.Sea).Conn("brs", godip.Sea).Conn("wal", godip.Sea).Conn("not", godip.Sea).Conn("not/wc", godip.Sea).Conn("sco", godip.Sea).Conn("min", godip.Sea).Flag(godip.Sea).
		// Dijon
		Prov("dij").Conn("lux", godip.Land).Conn("fla", godip.Land).Conn("cal", godip.Land).Conn("par", godip.Land).Conn("cha", godip.Land).Conn("dau", godip.Land).Conn("can", godip.Land).Conn("lor", godip.Land).Flag(godip.Land).SC(Burgundy).
		// Scotland
		Prov("sco").Conn("nos", godip.Sea).Conn("iri", godip.Sea).Conn("not", godip.Land).Conn("not/ec", godip.Sea).Conn("not/wc", godip.Sea).Conn("min", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Brittany
		Prov("brt").Conn("bis", godip.Sea).Conn("guy", godip.Coast...).Conn("poi", godip.Land).Conn("orl", godip.Land).Conn("anj", godip.Land).Conn("nom", godip.Coast...).Conn("eng", godip.Sea).Conn("brs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Limousin
		Prov("lim").Conn("poi", godip.Land).Conn("tou", godip.Land).Conn("pro", godip.Land).Conn("dau", godip.Land).Conn("orl", godip.Land).Flag(godip.Land).
		// Luxembourg
		Prov("lux").Conn("fri", godip.Land).Conn("hol", godip.Land).Conn("fla", godip.Land).Conn("dij", godip.Land).Conn("lor", godip.Land).Flag(godip.Land).SC(Burgundy).
		// Wales
		Prov("wal").Conn("brs", godip.Sea).Conn("dev", godip.Coast...).Conn("not", godip.Land).Conn("not/wc", godip.Sea).Conn("iri", godip.Sea).Flag(godip.Coast...).
		// English Channel
		Prov("eng").Conn("str", godip.Sea).Conn("lon", godip.Sea).Conn("dev", godip.Sea).Conn("brs", godip.Sea).Conn("brt", godip.Sea).Conn("nom", godip.Sea).Flag(godip.Sea).
		// Anglia
		Prov("ang").Conn("thw", godip.Sea).Conn("nos", godip.Sea).Conn("not", godip.Land).Conn("not/ec", godip.Sea).Conn("dev", godip.Land).Conn("lon", godip.Coast...).Conn("str", godip.Sea).Flag(godip.Coast...).
		// Aragon
		Prov("ara").Conn("tou", godip.Land).Conn("guy", godip.Land).Conn("cas", godip.Land).Flag(godip.Land).
		// Aragon (North Coast)
		Prov("ara/nc").Conn("guy", godip.Sea).Conn("bis", godip.Sea).Conn("cas", godip.Sea).Flag(godip.Sea).
		// Aragon (South Coast)
		Prov("ara/sc").Conn("med", godip.Sea).Conn("tou", godip.Sea).Conn("cas", godip.Sea).Flag(godip.Sea).
		// Castile
		Prov("cas").Conn("bis", godip.Sea).Conn("atl", godip.Sea).Conn("med", godip.Sea).Conn("ara", godip.Land).Conn("ara/nc", godip.Sea).Conn("ara/sc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Charolais
		Prov("cha").Conn("dau", godip.Land).Conn("dij", godip.Land).Conn("par", godip.Land).Flag(godip.Land).
		// The Pale
		Prov("thp").Conn("atl", godip.Sea).Conn("iri", godip.Sea).Conn("min", godip.Sea).Flag(godip.Coast...).
		// Holland
		Prov("hol").Conn("str", godip.Sea).Conn("fla", godip.Coast...).Conn("lux", godip.Land).Conn("fri", godip.Coast...).Conn("thw", godip.Sea).Flag(godip.Coast...).SC(Burgundy).
		// Minch
		Prov("min").Conn("thp", godip.Sea).Conn("iri", godip.Sea).Conn("sco", godip.Sea).Conn("nos", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"min":    "Minch",
	"not/wc": "Northumbria (WC)",
	"not/ec": "Northumbria (EC)",
	"ara/nc": "Aragon (NC)",
	"ara/sc": "Aragon (SC)",
	"med":    "Mediterranean",
	"str":    "Strait of Dover",
	"atl":    "Atlantic Sea",
	"nos":    "North Sea",
	"thw":    "The Wash",
	"brs":    "Bristol Channel",
	"eng":    "English Channel",
	"iri":    "Irish Sea",
	"bis":    "Biscay",
	"cas":    "Castile",
	"ara":    "Aragon",
	"tou":    "Toulouse",
	"guy":    "Guyenne",
	"brt":    "Brittany",
	"anj":    "Anjou",
	"poi":    "Poitou",
	"pro":    "Provence",
	"sav":    "Savoy",
	"lim":    "Limousin",
	"dau":    "Dauphine",
	"nom":    "Normandy",
	"can":    "Cantons",
	"orl":    "Orleanais",
	"cha":    "Charlerois",
	"par":    "Paris",
	"dij":    "Dijon",
	"cal":    "Calais",
	"als":    "Alsace",
	"lor":    "Lorraine",
	"fla":    "Flanders",
	"hol":    "Holland",
	"fri":    "Friesland",
	"lux":    "Luxembourg",
	"lon":    "London",
	"dev":    "Devon",
	"wal":    "Wales",
	"ang":    "Anglia",
	"not":    "Northumbria",
	"sco":    "Scotland",
	"thp":    "The Pale",
}
