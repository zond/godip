package classical

import (
	"fmt"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var (
	Nations    = []godip.Nation{godip.Austria, godip.England, godip.France, godip.Germany, godip.Italy, godip.Turkey, godip.Russia}
	PhaseTypes = []godip.PhaseType{godip.Movement, godip.Retreat, godip.Adjustment}
	Seasons    = []godip.Season{godip.Spring, godip.Fall}
	UnitTypes  = []godip.UnitType{godip.Army, godip.Fleet}
	SVGUnits   = map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	}
	SVGFlags = map[godip.Nation]func() ([]byte, error){
		godip.Austria: func() ([]byte, error) {
			return Asset("svg/austria.svg")
		},
		godip.England: func() ([]byte, error) {
			return Asset("svg/england.svg")
		},
		godip.France: func() ([]byte, error) {
			return Asset("svg/france.svg")
		},
		godip.Germany: func() ([]byte, error) {
			return Asset("svg/germany.svg")
		},
		godip.Italy: func() ([]byte, error) {
			return Asset("svg/italy.svg")
		},
		godip.Russia: func() ([]byte, error) {
			return Asset("svg/russia.svg")
		},
		godip.Turkey: func() ([]byte, error) {
			return Asset("svg/turkey.svg")
		},
	}
	Parser = orders.NewParser([]godip.Order{
		orders.BuildOrder,
		orders.ConvoyOrder,
		orders.DisbandOrder,
		orders.HoldOrder,
		orders.MoveOrder,
		orders.MoveViaConvoyOrder,
		orders.SupportOrder,
	})
)

func AdjustSCs(phase *phase.Phase) bool {
	return phase.Ty == godip.Retreat && phase.Se == godip.Fall
}

func NewPhase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return phase.Generator(Parser, AdjustSCs)(year, season, typ)
}

var ClassicalVariant = common.Variant{
	Name:  "Classical",
	Start: Start,
	Blank: Blank,
	BlankStart: func() (result *state.State, err error) {
		result = Blank(NewPhase(1900, godip.Fall, godip.Adjustment))
		return
	},
	Parser:     Parser,
	Graph:      func() godip.Graph { return start.Graph() },
	Phase:      NewPhase,
	Nations:    Nations,
	PhaseTypes: PhaseTypes,
	Seasons:    Seasons,
	UnitTypes:  UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/map.svg")
	},
	SVGVersion:  "5",
	SVGUnits:    SVGUnits,
	SVGFlags:    SVGFlags,
	CreatedBy:   "Allan B. Calhamer",
	Version:     "",
	Description: "The original game of Diplomacy.",
	Rules:       "The first to 18 supply centers is the winner. See the Wikibooks article for how to play: https://en.wikibooks.org/wiki/Diplomacy/Rules",
}

func Blank(phase godip.Phase) *state.State {
	return state.New(start.Graph(), phase, BackupRule, nil)
}

func Start() (result *state.State, err error) {
	result = Blank(NewPhase(1901, godip.Spring, godip.Movement))
	if err = result.SetUnits(start.Units()); err != nil {
		return
	}
	result.SetSupplyCenters(start.SupplyCenters())
	return
}

func BackupRule(state godip.State, deps []godip.Province) (err error) {
	only_moves := true
	convoys := false
	for _, prov := range deps {
		if order, _, ok := state.Order(prov); ok {
			if order.Type() != godip.Move {
				only_moves = false
			}
			if order.Type() == godip.Convoy {
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
			if order, _, ok := state.Order(prov); ok && order.Type() == godip.Convoy {
				state.SetResolution(prov, godip.ErrConvoyParadox)
			}
		}
		return
	}

	err = fmt.Errorf("Unknown circular dependency between %v", deps)
	return
}
