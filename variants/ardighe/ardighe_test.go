package ardighe

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := ArdigheVariant.Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(379, godip.Fall, godip.Adjustment)
	judge := ArdigheVariant.Blank(startPhase)
	return judge
}

// Rule 4. The game starts with the adjustment phase, in effect you may choose your starting units.
func TestBuildFirstPhase(t *testing.T) {
	judge := startState(t)
	// Test option to build army in Tuathal.
	tst.AssertOrderValidity(t, judge, orders.Build("tua", godip.Army, time.Now()), Connacht, nil)
	tst.AssertOrderValidity(t, judge, orders.Build("tua", godip.Fleet, time.Now()), "", godip.ErrIllegalUnitType)
	// Test option to build army or fleet in Guara.
	tst.AssertOrderValidity(t, judge, orders.Build("gua", godip.Army, time.Now()), Midhe, nil)
	tst.AssertOrderValidity(t, judge, orders.Build("gua", godip.Fleet, time.Now()), Midhe, nil)
}

// Rule 5. The winner is the power whom holds all 15 supply centers during an adjustment phase.
func TestWinnerNeedsFifteen(t *testing.T) {
	judge := blankState(t)
	for _, sc := range ArdigheVariant.Graph().AllSCs() {
		judge.SetSC(sc, Midhe)
	}
	tst.WaitForYears(judge, 1)
	winner := ArdigheVariant.SoloWinner(judge)
	if winner != Midhe {
		t.Errorf("Expected Midhe to have won")
	}
}
