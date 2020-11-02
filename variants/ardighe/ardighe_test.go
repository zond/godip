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
		judge.SetSC(sc, Mumhan)
	}
	tst.WaitForYears(judge, 1)
	winner := ArdigheVariant.SoloWinner(judge)
	if winner != Mumhan {
		t.Errorf("Expected Mumhan to have won")
	}
}

// Rule 6a. During Spring orders a fleet in one of the four outermost sea areas may give the Raiding order.
// 6b. During the resolution of Spring orders move any unit ordered to Raid from the sea area to the border of the map.
// 6c. During Fall orders a "Raiding Return" order most be written for each fleet that went "Raiding".
//     This order denotes which of the four outermost sea areas the "Raiding" Fleet will return to.
// 6d. A fleet following a “Raiding Return” order may not retreat, and thus is removed from the game as the result of a bounce.
// 6e. During the next adjustment phase the controller of any units that successfully executed the "Raiding Return" order
//     is considered to have an extra Supply Center for each fleet that successfully returned from the raid.
func TestBasicRaid(t *testing.T) {
}

// Check that a raiding fleet has no options except to return.
func TestRaiderCanOnlyReturn(t *testing.T) {
}

// Check that both bouncing fleets are removed.
func TestBounceReturningFromRaid(t *testing.T) {
}

// Check that a raiding fleet is removed if it is not given a return order.
func TestNoOrderReturningFromRaid(t *testing.T) {
}

// Check that builds gained by raiding do not count towards victory.
func TestRaidingDoesntCountForWin(t *testing.T) {
}
