package classical

import (
	"github.com/zond/godip/classical"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/classical/start"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

var ClassicalVariant = common.Variant{
	Name:  "Classical",
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
}
