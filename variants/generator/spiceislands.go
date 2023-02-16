package spiceislands

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Brunei    godip.Nation = "Brunei"
	Majapahit godip.Nation = "Majapahit"
)

var Nations = []godip.Nation{Brunei, Majapahit}

var SpiceIslandsVariant = common.Variant{
	Name:              "SpiceIslands",
	Graph:             func() godip.Graph { return SpiceIslandsGraph() },
	Start:             SpiceIslandsStart,
	Blank:             SpiceIslandsBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(6),
	SoloSCCount:       func(*state.State) int { return 6 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/spiceislandsmap.svg")
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

func SpiceIslandsBlank(phase godip.Phase) *state.State {
	return state.New(SpiceIslandsGraph(), phase, classical.BackupRule, nil, nil)
}

func SpiceIslandsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1570, godip.Spring, godip.Movement)
	result = SpiceIslandsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"lum": godip.Unit{godip.Army, Brunei},
		"weh": godip.Unit{godip.Fleet, Majapahit},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"weh": Majapahit,
	})
	return
}

func SpiceIslandsGraph() *graph.Graph {
	return graph.New().
		// Halmahera
		Prov("hal").Conn("gul", godip.Sea).Conn("mol", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Gulf of Tomini
		Prov("gul").Conn("mol", godip.Sea).Conn("hal", godip.Sea).Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("mih", godip.Sea).Conn("luw", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Sea).
		// Wehali
		Prov("weh").Conn("sto", godip.Sea).Conn("tim", godip.Sea).Conn("ban", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).SC(Majapahit).
		// Aceh
		Prov("ace").Conn("wes", godip.Sea).Conn("ped", godip.Coast...).Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("mig", godip.Coast...).Flag(godip.Coast...).
		// Timor Sea
		Prov("tim").Conn("big", godip.Sea).Conn("ser", godip.Sea).Conn("mol", godip.Sea).Conn("bur", godip.Sea).Conn("ban", godip.Sea).Conn("weh", godip.Sea).Conn("sto", godip.Sea).Flag(godip.Sea).
		// Makassar
		Prov("mak").Conn("ban", godip.Sea).Conn("but", godip.Coast...).Conn("luw", godip.Coast...).Conn("mih", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Jambi
		Prov("jam").Conn("ria", godip.Coast...).Conn("wes", godip.Sea).Conn("pal", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).
		// Big Ocean
		Prov("big").Conn("mig", godip.Sea).Conn("ace", godip.Sea).Conn("ace", godip.Sea).Conn("ped", godip.Sea).Conn("ria", godip.Sea).Conn("ria", godip.Sea).Conn("jam", godip.Sea).Conn("pal", godip.Sea).Conn("pal", godip.Sea).Conn("mig", godip.Sea).Conn("paj", godip.Sea).Conn("jav", godip.Sea).Conn("tro", godip.Sea).Conn("tro", godip.Sea).Conn("lum", godip.Sea).Conn("sto", godip.Sea).Conn("weh", godip.Sea).Conn("ban", godip.Sea).Conn("mak", godip.Sea).Conn("mih", godip.Sea).Conn("mih", godip.Sea).Conn("gul", godip.Sea).Conn("gul", godip.Sea).Conn("hal", godip.Sea).Conn("mol", godip.Sea).Conn("ser", godip.Sea).Conn("tim", godip.Sea).Flag(godip.Sea).
		// Riau
		Prov("ria").Conn("ped", godip.Coast...).Conn("wes", godip.Sea).Conn("jam", godip.Coast...).Conn("big", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Seram
		Prov("ser").Conn("tim", godip.Sea).Conn("big", godip.Sea).Conn("mol", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Palembang
		Prov("pal").Conn("wes", godip.Sea).Conn("mig", godip.Coast...).Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("jam", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Lumajang
		Prov("lum").Conn("big", godip.Sea).Conn("tro", godip.Coast...).Conn("jav", godip.Coast...).Conn("sto", godip.Sea).Flag(godip.Coast...).
		// Western Ocean
		Prov("wes").Conn("ace", godip.Sea).Conn("mig", godip.Sea).Conn("pal", godip.Sea).Conn("jam", godip.Sea).Conn("ria", godip.Sea).Conn("ped", godip.Sea).Flag(godip.Sea).
		// Trowulan
		Prov("tro").Conn("jav", godip.Coast...).Conn("lum", godip.Coast...).Conn("big", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sunda
		Prov("sun").Conn("paj", godip.Land).Conn("mig", godip.Land).Conn("jav", godip.Land).Flag(godip.Land).
		// Buru
		Prov("bur").Conn("tim", godip.Sea).Conn("mol", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Moluccan Sea
		Prov("mol").Conn("tim", godip.Sea).Conn("ser", godip.Sea).Conn("big", godip.Sea).Conn("hal", godip.Sea).Conn("gul", godip.Sea).Conn("ban", godip.Sea).Conn("bur", godip.Sea).Flag(godip.Sea).
		// Javadvipa
		Prov("jav").Conn("paj", godip.Coast...).Conn("sun", godip.Land).Conn("mig", godip.Coast...).Conn("sto", godip.Sea).Conn("lum", godip.Coast...).Conn("tro", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Pajajaran
		Prov("paj").Conn("jav", godip.Coast...).Conn("big", godip.Sea).Conn("mig", godip.Coast...).Conn("sun", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Southern Ocean
		Prov("sto").Conn("tim", godip.Sea).Conn("weh", godip.Sea).Conn("big", godip.Sea).Conn("lum", godip.Sea).Conn("jav", godip.Sea).Conn("mig", godip.Sea).Flag(godip.Sea).
		// Banda Sea
		Prov("ban").Conn("big", godip.Sea).Conn("weh", godip.Sea).Conn("tim", godip.Sea).Conn("bur", godip.Sea).Conn("mol", godip.Sea).Conn("gul", godip.Sea).Conn("luw", godip.Sea).Conn("but", godip.Sea).Conn("mak", godip.Sea).Flag(godip.Sea).
		// Minangkabau
		Prov("mig").Conn("sto", godip.Sea).Conn("jav", godip.Coast...).Conn("sun", godip.Land).Conn("paj", godip.Coast...).Conn("big", godip.Sea).Conn("pal", godip.Coast...).Conn("wes", godip.Sea).Conn("ace", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).
		// Pedir
		Prov("ped").Conn("ria", godip.Coast...).Conn("big", godip.Sea).Conn("ace", godip.Coast...).Conn("wes", godip.Sea).Flag(godip.Coast...).
		// Minahassa
		Prov("mih").Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("mak", godip.Coast...).Conn("luw", godip.Coast...).Conn("gul", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Luwuk
		Prov("luw").Conn("but", godip.Coast...).Conn("ban", godip.Sea).Conn("gul", godip.Sea).Conn("mih", godip.Coast...).Conn("mak", godip.Coast...).Flag(godip.Coast...).
		// Buton
		Prov("but").Conn("luw", godip.Coast...).Conn("mak", godip.Coast...).Conn("ban", godip.Sea).Flag(godip.Coast...).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"hal": "Halmahera",
	"gul": "Gulf of Tomini",
	"weh": "Wehali",
	"ace": "Aceh",
	"tim": "Timor Sea",
	"mak": "Makassar",
	"jam": "Jambi",
	"big": "Big Ocean",
	"ria": "Riau",
	"ser": "Seram",
	"pal": "Palembang",
	"lum": "Lumajang",
	"wes": "Western Ocean",
	"tro": "Trowulan",
	"sun": "Sunda",
	"bur": "Buru",
	"mol": "Moluccan Sea",
	"jav": "Javadvipa",
	"paj": "Pajajaran",
	"sto": "Southern Ocean",
	"ban": "Banda Sea",
	"mig": "Minangkabau",
	"ped": "Pedir",
	"mih": "Minahassa",
	"luw": "Luwuk",
	"but": "Buton",
}
