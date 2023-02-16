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
	SoloWinner:        common.SCCountWinner(2),
	SoloSCCount:       func(*state.State) int { return 2 },
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
		"lum": Brunei,
		"weh": Majapahit,
	})
	return
}

func SpiceIslandsGraph() *graph.Graph {
	return graph.New().
		// Wehali
		Prov("weh").Conn("big", godip.Sea).Conn("sto", godip.Sea).Conn("tim", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).SC(Majapahit).
		// Lumajang
		Prov("lum").Conn("big", godip.Sea).Conn("sto", godip.Sea).Conn("big", godip.Sea).Conn("tro", godip.Coast...).Flag(godip.Coast...).SC(Brunei).
		// Trowulan
		Prov("tro").Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("lum", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Big Ocean
		Prov("big").Conn("sto", godip.Sea).Conn("sto", godip.Sea).Conn("lum", godip.Sea).Conn("tro", godip.Sea).Conn("tro", godip.Sea).Conn("tro", godip.Sea).Conn("lum", godip.Sea).Conn("sto", godip.Sea).Conn("weh", godip.Sea).Conn("weh", godip.Sea).Conn("tim", godip.Sea).Flag(godip.Sea).
		// Southern Ocean
		Prov("sto").Conn("tim", godip.Sea).Conn("weh", godip.Sea).Conn("big", godip.Sea).Conn("lum", godip.Sea).Conn("big", godip.Sea).Conn("big", godip.Sea).Flag(godip.Sea).
		// Timor Sea
		Prov("tim").Conn("big", godip.Sea).Conn("weh", godip.Sea).Conn("sto", godip.Sea).Flag(godip.Sea).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"weh": "Wehali",
	"lum": "Lumajang",
	"tro": "Trowulan",
	"big": "Big Ocean",
	"sto": "Southern Ocean",
	"tim": "Timor Sea",
}
