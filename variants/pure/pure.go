package pure

import (
	"github.com/zond/godip/classical"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/classical/start"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

var PureVariant = common.Variant{
	Name: "Pure",
	Graph: func() dip.Graph { return start.PureGraph() },
	Start: PureStart,
	Blank: PureBlank,
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
}

func PureBlank(phase dip.Phase) *state.State {
	return state.New(start.PureGraph(), phase, classical.BackupRule)
}

func PureStart() (result *state.State, err error) {
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
}
