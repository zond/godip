package variants

import (
	"github.com/zond/godip/classical"
	cla "github.com/zond/godip/classical/common"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/classical/start"
	dip "github.com/zond/godip/common"
	"github.com/zond/godip/state"
)

const (
	Classical = "Classical"
	FleetRome = "Fleet Rome"
)

type Variant struct {
	Name        string
	Start       func() (*state.State, error)                                                             `json:"-"`
	BlankStart  func() (*state.State, error)                                                             `json:"-"`
	Blank       func(dip.Phase) *state.State                                                             `json:"-"`
	Phase       func(int, dip.Season, dip.PhaseType) dip.Phase                                           `json:"-"`
	ParseOrders func(map[dip.Nation]map[dip.Province][]string) (map[dip.Province]dip.Adjudicator, error) `json:"-"`
	ParseOrder  func([]string) (dip.Adjudicator, error)                                                  `json:"-"`
	Graph       dip.Graph
	Nations     []dip.Nation
	PhaseTypes  []dip.PhaseType
	Seasons     []dip.Season
	UnitTypes   []dip.UnitType
	OrderTypes  []dip.OrderType
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
		ParseOrders: orders.ParseAll,
		ParseOrder:  orders.Parse,
		Graph:       start.Graph(),
		Phase:       classical.Phase,
		OrderTypes:  orders.OrderTypes(),
		Nations:     cla.Nations,
		PhaseTypes:  cla.PhaseTypes,
		Seasons:     cla.Seasons,
		UnitTypes:   cla.UnitTypes,
	},
	Variant{
		Name:  FleetRome,
		Graph: start.Graph(),
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
		Blank:       classical.Blank,
		Phase:       classical.Phase,
		ParseOrders: orders.ParseAll,
		ParseOrder:  orders.Parse,
		OrderTypes:  orders.OrderTypes(),
		Nations:     cla.Nations,
		PhaseTypes:  cla.PhaseTypes,
		Seasons:     cla.Seasons,
		UnitTypes:   cla.UnitTypes,
	},
}
