package classical

import (
	"fmt"

	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/classical/start"

	cla "github.com/zond/godip/variants/classical/common"
	dip "github.com/zond/godip/common"
)

var ClassicalVariant = common.Variant{
	Name:  "Classical",
	Start: Start,
	Blank: Blank,
	BlankStart: func() (result *state.State, err error) {
		result = Blank(Phase(1900, cla.Fall, cla.Adjustment))
		return
	},
	ParseOrders:       orders.ParseAll,
	ParseOrder:        orders.Parse,
	Graph:             func() dip.Graph { return start.Graph() },
	Phase:             Phase,
	OrderTypes:        orders.OrderTypes(),
	Nations:           cla.Nations,
	PhaseTypes:        cla.PhaseTypes,
	Seasons:           cla.Seasons,
	UnitTypes:         cla.UnitTypes,
	SoloSupplyCenters: 18,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/map.svg")
	},
	SVGVersion: "3",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	},
	CreatedBy: "Allan B. Calhamer",
	Version: "",
	Description: "The original game of Diplomacy.",
	Rules: "The first to 18 supply centers is the winner. See the Wikibooks article for how to play: https://en.wikibooks.org/wiki/Diplomacy/Rules",
}
func Blank(phase dip.Phase) *state.State {
	return state.New(start.Graph(), phase, BackupRule)
}

func Start() (result *state.State, err error) {
	result = state.New(start.Graph(), &phase{1901, cla.Spring, cla.Movement}, BackupRule)
	if err = result.SetUnits(start.Units()); err != nil {
		return
	}
	result.SetSupplyCenters(start.SupplyCenters())
	return
}

func BackupRule(state dip.State, deps []dip.Province) (err error) {
	only_moves := true
	convoys := false
	for _, prov := range deps {
		if order, _, ok := state.Order(prov); ok {
			if order.Type() != cla.Move {
				only_moves = false
			}
			if order.Type() == cla.Convoy {
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
			if order, _, ok := state.Order(prov); ok && order.Type() == cla.Convoy {
				state.SetResolution(prov, cla.ErrConvoyParadox)
			}
		}
		return
	}

	err = fmt.Errorf("Unknown circular dependency between %v", deps)
	return
}
