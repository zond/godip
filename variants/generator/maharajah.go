package maharajah

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Oda     godip.Nation = "Oda"
	Shimazu godip.Nation = "Shimazu"
)

var Nations = []godip.Nation{Oda, Shimazu}

var MaharajahVariant = common.Variant{
	Name:              "Maharajah",
	Graph:             func() godip.Graph { return MaharajahGraph() },
	Start:             MaharajahStart,
	Blank:             MaharajahBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(7),
	SoloSCCount:       func(*state.State) int { return 7 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/maharajahmap.svg")
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

func MaharajahBlank(phase godip.Phase) *state.State {
	return state.New(MaharajahGraph(), phase, classical.BackupRule, nil, nil)
}

func MaharajahStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1501, godip.Spring, godip.Movement)
	result = MaharajahBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"kra": godip.Unit{godip.Army, Oda},
		"ayu": godip.Unit{godip.Army, Shimazu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"ayu": Shimazu,
	})
	return
}

func MaharajahGraph() *graph.Graph {
	return graph.New().
		// Sambalpur
		Prov("sam").Conn("jab", godip.Land).Conn("rai", godip.Land).Conn("war", godip.Land).Conn("ori", godip.Land).Conn("beg", godip.Land).Conn("bea", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Tinnevelly
		Prov("tin").Conn("pul", godip.Coast...).Conn("cal", godip.Coast...).Conn("are", godip.Sea).Conn("nic", godip.Sea).Flag(godip.Coast...).
		// Aceh
		Prov("ace").Conn("and", godip.Sea).Conn("nic", godip.Sea).Conn("are", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Warangal
		Prov("war").Conn("ori", godip.Land).Conn("sam", godip.Land).Conn("rai", godip.Coast...).Conn("are", godip.Sea).Conn("ban", godip.Coast...).Flag(godip.Coast...).
		// Nepal
		Prov("nep").Conn("muz", godip.Land).Conn("ava", godip.Land).Conn("gar", godip.Land).Conn("awa", godip.Land).Flag(godip.Land).
		// Muzaffarpur
		Prov("muz").Conn("awa", godip.Coast...).Conn("are", godip.Sea).Conn("bea", godip.Coast...).Conn("beg", godip.Land).Conn("ass", godip.Land).Conn("ava", godip.Land).Conn("nep", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Berar
		Prov("ber").Conn("jab", godip.Land).Conn("bea", godip.Coast...).Conn("are", godip.Sea).Conn("are", godip.Sea).Conn("rai", godip.Coast...).Flag(godip.Coast...).
		// Raipur
		Prov("rai").Conn("are", godip.Sea).Conn("war", godip.Coast...).Conn("sam", godip.Land).Conn("jab", godip.Land).Conn("ber", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Kra
		Prov("kra").Conn("ayu", godip.Coast...).Conn("and", godip.Sea).Flag(godip.Coast...).
		// Kashgar
		Prov("kas").Conn("are", godip.Sea).Conn("are", godip.Sea).Conn("gar", godip.Coast...).Flag(godip.Coast...).
		// Ayutthaya
		Prov("ayu").Conn("ava", godip.Land).Conn("peg", godip.Coast...).Conn("bay", godip.Sea).Conn("and", godip.Sea).Conn("kra", godip.Coast...).Flag(godip.Coast...).SC(Shimazu).
		// Pulicat
		Prov("pul").Conn("tin", godip.Coast...).Conn("nic", godip.Sea).Conn("bay", godip.Sea).Conn("ban", godip.Coast...).Conn("cal", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Jaffna
		Prov("jaf").Conn("kan", godip.Coast...).Conn("nic", godip.Sea).Conn("are", godip.Sea).Flag(godip.Coast...).
		// ANDAMAN SEA
		Prov("and").Conn("kra", godip.Sea).Conn("ayu", godip.Sea).Conn("bay", godip.Sea).Conn("nic", godip.Sea).Conn("ace", godip.Sea).Flag(godip.Sea).
		// Area
		Prov("are").Conn("ace", godip.Sea).Conn("nic", godip.Sea).Conn("kan", godip.Sea).Conn("kan", godip.Sea).Conn("kan", godip.Sea).Conn("jaf", godip.Sea).Conn("nic", godip.Sea).Conn("tin", godip.Sea).Conn("cal", godip.Sea).Conn("cal", godip.Sea).Conn("ban", godip.Sea).Conn("ban", godip.Sea).Conn("war", godip.Sea).Conn("rai", godip.Sea).Conn("ber", godip.Sea).Conn("ber", godip.Sea).Conn("bea", godip.Sea).Conn("bea", godip.Sea).Conn("bea", godip.Sea).Conn("muz", godip.Sea).Conn("awa", godip.Sea).Conn("awa", godip.Sea).Conn("gar", godip.Sea).Conn("kas", godip.Sea).Conn("kas", godip.Sea).Flag(godip.Sea).
		// Gartok
		Prov("gar").Conn("awa", godip.Coast...).Conn("nep", godip.Land).Conn("kas", godip.Coast...).Conn("are", godip.Sea).Flag(godip.Coast...).
		// Jabalpur
		Prov("jab").Conn("ber", godip.Land).Conn("rai", godip.Land).Conn("sam", godip.Land).Conn("bea", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Benares
		Prov("bea").Conn("are", godip.Sea).Conn("are", godip.Sea).Conn("ber", godip.Coast...).Conn("jab", godip.Land).Conn("sam", godip.Land).Conn("beg", godip.Land).Conn("muz", godip.Coast...).Conn("are", godip.Sea).Flag(godip.Coast...).
		// Awadh
		Prov("awa").Conn("gar", godip.Coast...).Conn("are", godip.Sea).Conn("are", godip.Sea).Conn("muz", godip.Coast...).Conn("nep", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Bangalore
		Prov("ban").Conn("pul", godip.Coast...).Conn("bay", godip.Sea).Conn("ori", godip.Coast...).Conn("war", godip.Coast...).Conn("are", godip.Sea).Conn("are", godip.Sea).Conn("cal", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// NICOBAR SEA
		Prov("nic").Conn("are", godip.Sea).Conn("ace", godip.Sea).Conn("and", godip.Sea).Conn("bay", godip.Sea).Conn("pul", godip.Sea).Conn("tin", godip.Sea).Conn("are", godip.Sea).Conn("jaf", godip.Sea).Conn("kan", godip.Sea).Flag(godip.Sea).
		// Ava
		Prov("ava").Conn("nep", godip.Land).Conn("muz", godip.Land).Conn("ass", godip.Land).Conn("peg", godip.Land).Conn("ayu", godip.Land).Flag(godip.Land).
		// Assam
		Prov("ass").Conn("ava", godip.Land).Conn("muz", godip.Land).Conn("beg", godip.Land).Conn("peg", godip.Land).Flag(godip.Land).
		// Kandy
		Prov("kan").Conn("are", godip.Sea).Conn("nic", godip.Sea).Conn("jaf", godip.Coast...).Conn("are", godip.Sea).Conn("are", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Bengal
		Prov("beg").Conn("peg", godip.Coast...).Conn("ass", godip.Land).Conn("muz", godip.Land).Conn("bea", godip.Land).Conn("sam", godip.Land).Conn("ori", godip.Coast...).Conn("bay", godip.Sea).Flag(godip.Coast...).
		// Orissa
		Prov("ori").Conn("war", godip.Land).Conn("ban", godip.Coast...).Conn("bay", godip.Sea).Conn("beg", godip.Coast...).Conn("sam", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Calicut
		Prov("cal").Conn("tin", godip.Coast...).Conn("pul", godip.Land).Conn("ban", godip.Coast...).Conn("are", godip.Sea).Conn("are", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Pegu
		Prov("peg").Conn("beg", godip.Coast...).Conn("bay", godip.Sea).Conn("ayu", godip.Coast...).Conn("ava", godip.Land).Conn("ass", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// BAY OF BENGAL
		Prov("bay").Conn("and", godip.Sea).Conn("ayu", godip.Sea).Conn("peg", godip.Sea).Conn("beg", godip.Sea).Conn("ori", godip.Sea).Conn("ban", godip.Sea).Conn("pul", godip.Sea).Conn("nic", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"sam": "Sambalpur",
	"tin": "Tinnevelly",
	"ace": "Aceh",
	"war": "Warangal",
	"nep": "Nepal",
	"muz": "Muzaffarpur",
	"ber": "Berar",
	"rai": "Raipur",
	"kra": "Kra",
	"kas": "Kashgar",
	"ayu": "Ayutthaya",
	"pul": "Pulicat",
	"jaf": "Jaffna",
	"and": "ANDAMAN SEA",
	"are": "Area",
	"gar": "Gartok",
	"jab": "Jabalpur",
	"bea": "Benares",
	"awa": "Awadh",
	"ban": "Bangalore",
	"nic": "NICOBAR SEA",
	"ava": "Ava",
	"ass": "Assam",
	"kan": "Kandy",
	"beg": "Bengal",
	"ori": "Orissa",
	"cal": "Calicut",
	"peg": "Pegu",
	"bay": "BAY OF BENGAL",
}
