package classical

import (
	"fmt"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var (
	Nations    = []godip.Nation{godip.Austria, godip.England, godip.France, godip.Germany, godip.Italy, godip.Turkey, godip.Russia}
	PhaseTypes = []godip.PhaseType{godip.Movement, godip.Retreat, godip.Adjustment}
	Seasons    = []godip.Season{godip.Spring, godip.Fall}
	UnitTypes  = []godip.UnitType{godip.Army, godip.Fleet}
	SVGUnits   = map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	}
	SVGFlags = map[godip.Nation]func() ([]byte, error){
		godip.Austria: func() ([]byte, error) {
			return Asset("svg/austria.svg")
		},
		godip.England: func() ([]byte, error) {
			return Asset("svg/england.svg")
		},
		godip.France: func() ([]byte, error) {
			return Asset("svg/france.svg")
		},
		godip.Germany: func() ([]byte, error) {
			return Asset("svg/germany.svg")
		},
		godip.Italy: func() ([]byte, error) {
			return Asset("svg/italy.svg")
		},
		godip.Russia: func() ([]byte, error) {
			return Asset("svg/russia.svg")
		},
		godip.Turkey: func() ([]byte, error) {
			return Asset("svg/turkey.svg")
		},
	}
	Parser = orders.NewParser([]godip.Order{
		orders.BuildOrder,
		orders.ConvoyOrder,
		orders.DisbandOrder,
		orders.HoldOrder,
		orders.MoveOrder,
		orders.MoveViaConvoyOrder,
		orders.SupportOrder,
	})
)

func AdjustSCs(phase *phase.Phase) bool {
	return phase.Ty == godip.Retreat && phase.Se == godip.Fall
}

func NewPhase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return phase.Generator(Parser, AdjustSCs)(year, season, typ)
}

var ClassicalVariant = common.Variant{
	Name:  "Classical",
	Start: Start,
	Blank: Blank,
	BlankStart: func() (result *state.State, err error) {
		result = Blank(NewPhase(1900, godip.Fall, godip.Adjustment))
		return
	},
	Parser: Parser,
	Graph:  func() godip.Graph { return start.Graph() },
	ExtraDominanceRules: map[godip.Province]common.DominanceRule{
		"gas": common.DominanceRule{
			Priority: 0,
			Nation:   godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"bre": godip.France,
				"par": godip.France,
				"mar": godip.France,
				"spa": godip.Neutral,
			},
		},
		"bur": common.DominanceRule{
			Nation: godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"mar": godip.France,
				"par": godip.France,
				"mun": godip.Germany,
				"bel": godip.Neutral,
			},
		},
		"pic": common.DominanceRule{
			Nation: godip.France,
			Dependencies: map[godip.Province]godip.Nation{
				"par": godip.France,
				"bre": godip.France,
				"bel": godip.Neutral,
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
				"rum": godip.Neutral,
			},
		},
		"ukr": common.DominanceRule{
			Nation: godip.Russia,
			Dependencies: map[godip.Province]godip.Nation{
				"war": godip.Russia,
				"mos": godip.Russia,
				"stp": godip.Russia,
				"rum": godip.Neutral,
			},
		},
		"fin": common.DominanceRule{
			Nation: godip.Russia,
			Dependencies: map[godip.Province]godip.Nation{
				"stp": godip.Russia,
				"swe": godip.Neutral,
				"nwy": godip.Neutral,
			},
		},
		"ruh": common.DominanceRule{
			Nation: godip.Germany,
			Dependencies: map[godip.Province]godip.Nation{
				"kie": godip.Germany,
				"mun": godip.Germany,
				"bel": godip.Neutral,
				"hol": godip.Neutral,
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
	},
	Phase:      NewPhase,
	Nations:    Nations,
	PhaseTypes: PhaseTypes,
	Seasons:    Seasons,
	UnitTypes:  UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/map.svg")
	},
	ProvinceLongNames: provinceLongNames,
	SVGVersion:        "9",
	SVGUnits:          SVGUnits,
	SVGFlags:          SVGFlags,
	CreatedBy:         "Allan B. Calhamer",
	Version:           "",
	Description:       "The original Diplomacy.",
	SoloSCCount:       func(*state.State) int { return 18 },
	Rules: `The first to 18 Supply Centers (SC) is the winner. 
Kiel and Constantinople have a canal, fleets can move through it. 
Armies can move from Denmark to Kiel.`,
}

func Blank(phase godip.Phase) *state.State {
	return state.New(start.Graph(), phase, BackupRule, nil, nil)
}

func Start() (result *state.State, err error) {
	result = Blank(NewPhase(1901, godip.Spring, godip.Movement))
	if err = result.SetUnits(start.Units()); err != nil {
		return
	}
	result.SetSupplyCenters(start.SupplyCenters())
	return
}

func BackupRule(state godip.State, deps []godip.Province) (err error) {
	only_moves := true
	convoys := false
	for _, prov := range deps {
		if order, _, ok := state.Order(prov); ok {
			if order.Type() != godip.Move {
				only_moves = false
			}
			if order.Type() == godip.Convoy {
				convoys = true
			}
		}
	}

	if only_moves {
		for _, prov := range deps {
			state.SetResolution(prov, nil)
		}
		return
	}
	if convoys {
		for _, prov := range deps {
			if order, _, ok := state.Order(prov); ok && order.Type() == godip.Convoy {
				state.SetResolution(prov, godip.ErrConvoyParadox)
			}
		}
		return
	}

	err = fmt.Errorf("Unknown circular dependency between %v", deps)
	return
}

var provinceLongNames = map[godip.Province]string{
	"bul/ec": "Bulgaria (EC)",
	"bul/sc": "Bulgaria (SC)",
	"stp/sc": "St. Petersburg (SC)",
	"stp/nc": "St. Petersburg (NC)",
	"spa/nc": "Spain (NC)",
	"spa/sc": "Spain (SC)",
	"con":    "Constantinople",
	"sil":    "Silesia",
	"bal":    "Baltic Sea",
	"ber":    "Berlin",
	"den":    "Denmark",
	"stp":    "St. Petersburg",
	"ion":    "Ionian Sea",
	"boh":    "Bohemia",
	"yor":    "Yorkshire",
	"hel":    "Heligoland Bight",
	"bot":    "Gulf of Bothnia",
	"iri":    "Irish Sea",
	"syr":    "Syria",
	"bel":    "Belgium",
	"lvp":    "Liverpool",
	"bar":    "Barents Sea",
	"lvn":    "Livonia",
	"tri":    "Trieste",
	"bud":    "Budapest",
	"ank":    "Ankara",
	"eas":    "East Med",
	"adr":    "Adriatic Sea",
	"ven":    "Venice",
	"bul":    "Bulgaria",
	"gal":    "Galicia",
	"nth":    "North Sea",
	"nwy":    "Norway",
	"gas":    "Gascony",
	"tus":    "Tuscany",
	"nrg":    "Norwegian Sea",
	"bur":    "Burgundy",
	"rum":    "Rumania",
	"aeg":    "Aegean Sea",
	"tys":    "Tyrrhenian Sea",
	"mar":    "Marseilles",
	"ruh":    "Ruhr",
	"cly":    "Clyde",
	"war":    "Warsaw",
	"bla":    "Black Sea",
	"mun":    "Munich",
	"kie":    "Kiel",
	"nat":    "North Atlantic",
	"tyr":    "Tyrolia",
	"ska":    "Skagerakk (SKA)",
	"gre":    "Greece",
	"nap":    "Naples",
	"mos":    "Moscow",
	"wes":    "West Mediterranean",
	"ukr":    "Ukraine",
	"lon":    "London",
	"hol":    "Holland",
	"mid":    "Mid-Atlantic",
	"eng":    "English Channel",
	"smy":    "Smyrna",
	"naf":    "North Africa",
	"wal":    "Wales",
	"par":    "Paris",
	"gol":    "Gulf of Lyon",
	"rom":    "Rome",
	"arm":    "Armenia",
	"fin":    "Finland",
	"bre":    "Brest",
	"spa":    "Spain",
	"pic":    "Picardy",
	"pru":    "Prussia",
	"apu":    "Apulia",
	"pie":    "Piedmont",
	"alb":    "Albania",
	"edi":    "Edinburgh",
	"por":    "Portugal",
	"swe":    "Sweden",
	"vie":    "Vienna",
	"ser":    "Serbia",
	"sev":    "Sevastopol",
	"tun":    "Tunis",
}
