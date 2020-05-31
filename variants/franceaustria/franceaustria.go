package franceaustria

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var FranceAustriaVariant = common.Variant{
	Name: "France vs Austria",
	Graph: func() godip.Graph {
		okNations := map[godip.Nation]bool{
			godip.France:  true,
			godip.Austria: true,
			godip.Neutral: true,
		}
		neutral := godip.Neutral
		result := start.Graph()
		for _, node := range result.Nodes {
			if node.SC != nil && !okNations[*node.SC] {
				node.SC = &neutral
			}
		}
		return result
	},
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		if err = result.SetUnits(map[godip.Province]godip.Unit{
			"bre": godip.Unit{godip.Fleet, godip.France},
			"par": godip.Unit{godip.Army, godip.France},
			"mar": godip.Unit{godip.Army, godip.France},
			"tri": godip.Unit{godip.Fleet, godip.Austria},
			"vie": godip.Unit{godip.Army, godip.Austria},
			"bud": godip.Unit{godip.Army, godip.Austria},
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[godip.Province]godip.Nation{
			"bre": godip.France,
			"par": godip.France,
			"mar": godip.France,
			"tri": godip.Austria,
			"vie": godip.Austria,
			"bud": godip.Austria,
		})
		return
	},
	Blank:      classical.Blank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    []godip.Nation{godip.Austria, godip.France},
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion:  "1",
	SVGUnits:    classical.SVGUnits,
	CreatedBy:   "",
	Version:     "",
	Description: "A two player variant on the classical map.",
	SoloSCCount: func(*state.State) int { return 18 },
	Rules: `The first to 18 supply centers is the winner. 
	The game only has two nations: France and Austria.`,
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
