package neutralarmies

import (
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"
)

func NeutralOrders(state state.State) (ret map[godip.Province]godip.Adjudicator) {
	ret = map[godip.Province]godip.Adjudicator{}
	if state.Phase().Type() == godip.Movement {

	}
	switch state.Phase().Type() {
	case godip.Movement:
		// Strictly this is unnecessary - because hold is the default order.
		for prov, unit := range state.Units() {
			if unit.Nation == godip.Neutral {
				ret[prov] = orders.Hold(prov)
			}
		}
	case godip.Adjustment:
		// Rebuild any missing units.
		for _, prov := range state.Graph().SCs(godip.Neutral) {
			if _, _, ok := state.SupplyCenter(prov); !ok {
				if _, _, ok := state.Unit(prov); !ok {
					ret[prov] = orders.Build(prov, godip.Army, time.Now())
				}
			}
		}
	}
	return
}

func Blank(phase godip.Phase) *state.State {
	return state.New(start.Graph(), phase, classical.BackupRule, NeutralOrders)
}

func Start() (result *state.State, err error) {
	result = Blank(classical.NewPhase(1901, godip.Spring, godip.Movement))
	if err = result.SetUnits(start.Units()); err != nil {
		return
	}
	result.SetSupplyCenters(start.SupplyCenters())
	for _, sc := range start.Graph().SCs(godip.Neutral) {
		if err = result.SetUnit(godip.Province(sc), godip.Unit{
			Type:   godip.Army,
			Nation: godip.Neutral,
		}); err != nil {
			return
		}
	}
	return
}

var NeutralArmiesVariant = common.Variant{
	Name:       "Neutral Armies",
	Graph:      func() godip.Graph { return start.Graph() },
	Start:      Start,
	Blank:      classical.Blank,
	Phase:      classical.NewPhase,
	Parser:     hundred.BuildAnywhereParser,
	Nations:    []godip.Nation{godip.Austria, godip.England, godip.France, godip.Germany, godip.Italy, godip.Turkey, godip.Russia},
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion: "1482957154",
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
