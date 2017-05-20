package variants

import (
	"github.com/zond/godip/classical"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/classical/start"
	"github.com/zond/godip/state"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

const (
	Classical     = "Classical"
	FleetRome     = "Fleet Rome"
	FranceAustria = "France vs Austria"
	Pure          = "Pure"
)

// Variant defines a dippy variant supported by godip.
type Variant struct {
	// Name is the display name and key for this variant.
	Name string
	// Start returns a state with the correct graph, units, phase and supply centers for starting this variant.
	Start func() (*state.State, error) `json:"-"`
	// BlankStart returns a state with the correct graph, phase and supply centers for starting this variant.
	BlankStart func() (*state.State, error) `json:"-"`
	// Blank returns a state with the correct graph and the provided phase for this variant.
	Blank func(dip.Phase) *state.State `json:"-"`
	// Phase returns a phase with the provided year, season and phase type for this variant.
	Phase func(int, dip.Season, dip.PhaseType) dip.Phase `json:"-"`
	// ParserOrders parses a map of orders.
	ParseOrders func(map[dip.Nation]map[dip.Province][]string) (map[dip.Province]dip.Adjudicator, error) `json:"-"`
	// ParseOrder parses a single tokenized order.
	ParseOrder func([]string) (dip.Adjudicator, error) `json:"-"`
	// Graph is the graph for this variant.
	Graph func() dip.Graph
	// Nations are the nations playing this variant.
	Nations []dip.Nation
	// PhaseTypes are the phase types the phases of this variant have.
	PhaseTypes []dip.PhaseType
	// Seasons are the seasons the phases of this variant have.
	Seasons []dip.Season
	// UnitTypes are the types the units of this variant have.
	UnitTypes []dip.UnitType
	// OrderTypes are the types the orders of this variant have.
	OrderTypes []dip.OrderType
	// Number of SCs required to solo.
	SoloSupplyCenters int
	// SVG representing the variant map graphics.
	SVGMap func() ([]byte, error) `json:"-"`
	// A version for the vector graphics (for use in caching mechanisms).
	SVGVersion string
	// SVG representing the variant units.
	SVGUnits map[dip.UnitType]func() ([]byte, error) `json:"-"`
}

func init() {
	for _, variant := range OrderedVariants {
		Variants[variant.Name] = variant
	}
}

var Variants = map[string]Variant{}

var OrderedVariants = []Variant{
	Variant{
		Name:  Classical,
		Start: classical.Start,
		Blank: classical.Blank,
		BlankStart: func() (result *state.State, err error) {
			result = classical.Blank(classical.Phase(1900, cla.Fall, cla.Adjustment))
			return
		},
		ParseOrders:       orders.ParseAll,
		ParseOrder:        orders.Parse,
		Graph:             func() dip.Graph { return start.Graph() },
		Phase:             classical.Phase,
		OrderTypes:        orders.OrderTypes(),
		Nations:           cla.Nations,
		PhaseTypes:        cla.PhaseTypes,
		Seasons:           cla.Seasons,
		UnitTypes:         cla.UnitTypes,
		SoloSupplyCenters: 18,
		SVGMap: func() ([]byte, error) {
			return classical.Asset("svg/map.svg")
		},
		SVGVersion: "1482957154",
		SVGUnits: map[dip.UnitType]func() ([]byte, error){
			cla.Army: func() ([]byte, error) {
				return classical.Asset("svg/army.svg")
			},
			cla.Fleet: func() ([]byte, error) {
				return classical.Asset("svg/fleet.svg")
			},
		},
	},
	Variant{
		Name:  FleetRome,
		Graph: func() dip.Graph { return start.Graph() },
		Start: func() (result *state.State, err error) {
			if result, err = classical.Start(); err != nil {
				return
			}
			result.RemoveUnit(dip.Province("rom"))
			if err = result.SetUnit(dip.Province("rom"), dip.Unit{
				Type:   cla.Fleet,
				Nation: cla.Italy,
			}); err != nil {
				return
			}
			return
		},
		Blank:             classical.Blank,
		Phase:             classical.Phase,
		ParseOrders:       orders.ParseAll,
		ParseOrder:        orders.Parse,
		OrderTypes:        orders.OrderTypes(),
		Nations:           cla.Nations,
		PhaseTypes:        cla.PhaseTypes,
		Seasons:           cla.Seasons,
		UnitTypes:         cla.UnitTypes,
		SoloSupplyCenters: 18,
		SVGMap: func() ([]byte, error) {
			return classical.Asset("svg/map.svg")
		},
		SVGVersion: "1482957154",
		SVGUnits: map[dip.UnitType]func() ([]byte, error){
			cla.Army: func() ([]byte, error) {
				return classical.Asset("svg/army.svg")
			},
			cla.Fleet: func() ([]byte, error) {
				return classical.Asset("svg/fleet.svg")
			},
		},
	},
	Variant{
		Name: FranceAustria,
		Graph: func() dip.Graph {
			okNations := map[dip.Nation]bool{
				cla.France:  true,
				cla.Austria: true,
				cla.Neutral: true,
			}
			neutral := cla.Neutral
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
			if err = result.SetUnits(map[dip.Province]dip.Unit{
				"bre": dip.Unit{cla.Fleet, cla.France},
				"par": dip.Unit{cla.Army, cla.France},
				"mar": dip.Unit{cla.Army, cla.France},
				"tri": dip.Unit{cla.Fleet, cla.Austria},
				"vie": dip.Unit{cla.Army, cla.Austria},
				"bud": dip.Unit{cla.Army, cla.Austria},
			}); err != nil {
				return
			}
			result.SetSupplyCenters(map[dip.Province]dip.Nation{
				"bre": cla.France,
				"par": cla.France,
				"mar": cla.France,
				"tri": cla.Austria,
				"vie": cla.Austria,
				"bud": cla.Austria,
			})
			return
		},
		Blank:             classical.Blank,
		Phase:             classical.Phase,
		ParseOrders:       orders.ParseAll,
		ParseOrder:        orders.Parse,
		OrderTypes:        orders.OrderTypes(),
		Nations:           []dip.Nation{cla.Austria, cla.France},
		PhaseTypes:        cla.PhaseTypes,
		Seasons:           cla.Seasons,
		UnitTypes:         cla.UnitTypes,
		SoloSupplyCenters: 18,
		SVGMap: func() ([]byte, error) {
			return classical.Asset("svg/map.svg")
		},
		SVGVersion: "1482957154",
		SVGUnits: map[dip.UnitType]func() ([]byte, error){
			cla.Army: func() ([]byte, error) {
				return classical.Asset("svg/army.svg")
			},
			cla.Fleet: func() ([]byte, error) {
				return classical.Asset("svg/fleet.svg")
			},
		},
	},
	Variant{
		Name: Pure,
		Graph: func() dip.Graph { return start.PureGraph() },
		Start: func() (result *state.State, err error) {
			if result, err = classical.Start(); err != nil {
				return
			}
			if err = result.SetUnits(map[dip.Province]dip.Unit{
				"ber": dip.Unit{cla.Army, cla.Germany},
				"lon": dip.Unit{cla.Army, cla.England},
				"par": dip.Unit{cla.Army, cla.France},
				"rom": dip.Unit{cla.Army, cla.Italy},
				"con": dip.Unit{cla.Army, cla.Turkey},
				"vie": dip.Unit{cla.Army, cla.Austria},
				"mos": dip.Unit{cla.Army, cla.Russia},
			}); err != nil {
				return
			}
			result.SetSupplyCenters(map[dip.Province]dip.Nation{
				"ber": cla.Germany,
				"lon": cla.England,
				"par": cla.France,
				"rom": cla.Italy,
				"con": cla.Turkey,
				"vie": cla.Austria,
				"mos": cla.Russia,
			})
			return
		},
		Blank:             classical.Blank,
		Phase:             classical.Phase,
		ParseOrders:       orders.ParseAll,
		ParseOrder:        orders.Parse,
		OrderTypes:        []dip.OrderType{
			cla.Build,
			cla.Move,
			cla.Hold,
			cla.Support,
			cla.Disband,
		},
		Nations:           cla.Nations,
		PhaseTypes:        cla.PhaseTypes,
		Seasons:           cla.Seasons,
		UnitTypes:         []dip.UnitType{cla.Army},
		SoloSupplyCenters: 4,
		SVGMap: func() ([]byte, error) {
			return Asset("svg/puremap.svg")
		},
		SVGVersion: "1",
		SVGUnits: map[dip.UnitType]func() ([]byte, error){
			cla.Army: func() ([]byte, error) {
				return classical.Asset("svg/army.svg")
			},
		},
	},
}
