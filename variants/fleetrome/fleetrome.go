package fleetrome

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var FleetRomeVariant = common.Variant{
	Name:  "Fleet Rome",
	Graph: func() godip.Graph { return start.Graph() },
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		result.RemoveUnit(godip.Province("rom"))
		if err = result.SetUnit(godip.Province("rom"), godip.Unit{
			Type:   godip.Fleet,
			Nation: godip.Italy,
		}); err != nil {
			return
		}
		return
	},
	Blank:      classical.Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    classical.Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion:  "1",
	SVGUnits:    classical.SVGUnits,
	SVGFlags:    classical.SVGFlags,
	CreatedBy:   "Richard Sharp",
	Version:     "",
	Description: "Classical Diplomacy, but Italy starts with a fleet in Rome.",
	SoloSCCount: func(*state.State) int { return 18 },
	Rules: `The first to 18 supply centers is the winner.  
	Italy starts with a fleet in Rome rather than an army.`,
}

var provinceLongNames = map[godip.Province]string{
	"bul/ec": "Bulgaria (EC)",
	"bul/sc": "Bulgaria (SC)",
	"stp/sc": "Sevastopol (SC)",
	"stp/nc": "Sevastopol (NC)",
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
	"tys":    "Tyrhennian Sea",
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
