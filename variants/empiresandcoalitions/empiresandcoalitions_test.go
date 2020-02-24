package empiresandcoalitions

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
	judge, err := EmpiresAndCoalitionsStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(1800, godip.Spring, godip.Movement)
	judge := EmpiresAndCoalitionsBlank(startPhase)
	return judge
}

func TestDenmarkExtraHomeCenter(t *testing.T) {
	judge := startState(t)
	// Try capturing Sweden.
	judge.SetOrder("cop", orders.Move("cop", "swe"))
	tst.WaitForYears(judge, 1)
	// Move back.
	judge.SetOrder("swe", orders.Move("swe", "cop"))
	tst.WaitForPhases(judge, 4)

	// Try building in Sweden.
	judge.SetOrder("swe", orders.Build("swe", godip.Army, time.Now()))
	tst.WaitForPhases(judge, 1)

	// Check successful.
	tst.AssertUnit(t, judge, "swe", godip.Unit{godip.Army, Denmark})
}

func TestSicilyExtraHomeCenter(t *testing.T) {
	judge := startState(t)
	// Try capturing Papal States.
	judge.SetOrder("nap", orders.Move("nap", "pap"))
	tst.WaitForYears(judge, 1)
	// Move back.
	judge.SetOrder("pap", orders.Move("pap", "nap"))
	tst.WaitForPhases(judge, 4)

	// Try building in Papal States.
	judge.SetOrder("pap", orders.Build("pap", godip.Army, time.Now()))
	tst.WaitForPhases(judge, 1)

	// Check successful.
	tst.AssertUnit(t, judge, "pap", godip.Unit{godip.Army, Sicily})
}

func TestSpainExtraHomeCenter(t *testing.T) {
	judge := startState(t)
	// Try capturing Portugal.
	judge.SetOrder("mad", orders.Move("mad", "por"))
	tst.WaitForYears(judge, 1)
	// Move back.
	judge.SetOrder("por", orders.Move("por", "mad"))
	tst.WaitForPhases(judge, 4)

	// Try building in Portugal.
	judge.SetOrder("por", orders.Build("por", godip.Army, time.Now()))
	tst.WaitForPhases(judge, 1)

	// Check successful.
	tst.AssertUnit(t, judge, "por", godip.Unit{godip.Army, Spain})
}

func TestOttomanEmpireExtraHomeCenter(t *testing.T) {
	judge := startState(t)
	// Try capturing Egypt.
	judge.SetOrder("ang", orders.Move("ang", "eas"))
	tst.WaitForPhases(judge, 2)
	judge.SetOrder("eas", orders.Move("eas", "egy"))
	tst.WaitForPhases(judge, 3)
	// Move out.
	judge.SetOrder("egy", orders.Move("egy", "tri"))
	tst.WaitForPhases(judge, 4)

	// Try building in Egypt.
	judge.SetOrder("egy", orders.Build("egy", godip.Army, time.Now()))
	tst.WaitForPhases(judge, 1)

	// Check successful.
	tst.AssertUnit(t, judge, "egy", godip.Unit{godip.Army, OttomanEmpire})
}

func TestNotBuildAnywhere(t *testing.T) {
	judge := startState(t)
	// Try capturing Finland.
	judge.SetOrder("stp", orders.Move("stp", "fil"))
	tst.WaitForYears(judge, 1)
	// Move back.
	judge.SetOrder("por", orders.Move("fil", "stp"))
	tst.WaitForPhases(judge, 4)

	// Try building in Finland.
	judge.SetOrder("fil", orders.Build("fil", godip.Army, time.Now()))
	tst.WaitForPhases(judge, 1)

	// Check not successful.
	tst.AssertNoUnit(t, judge, "fil")
}

func TestGibraltarMovement(t *testing.T) {
	judge := startState(t)

	// Check British fleet starts in Gibratar, not in its supply center (Liverpool).
	tst.AssertUnit(t, judge, "gib", godip.Unit{godip.Fleet, Britain})
	tst.AssertNoUnit(t, judge, "lie")
	tst.AssertOwner(t, judge, "lie", Britain)

	// Test naval connections from Gibraltar.
	tst.AssertOrderValidity(t, judge, orders.Move("gib", "mor"), Britain, nil)
	tst.AssertOptionToMove(t, judge, Britain, "gib", "mor")
	tst.AssertOrderValidity(t, judge, orders.Move("gib", "and/wc"), Britain, nil)
	tst.AssertOptionToMove(t, judge, Britain, "gib", "and/wc")
	tst.AssertOrderValidity(t, judge, orders.Move("gib", "and/ec"), Britain, nil)
	tst.AssertOptionToMove(t, judge, Britain, "gib", "and/ec")

	// Test connections for armies.
	judge.RemoveUnit("gib")
	judge.SetUnit("gib", godip.Unit{godip.Army, Britain})
	tst.AssertOrderValidity(t, judge, orders.Move("gib", "mor"), "", godip.ErrMissingConvoyPath)
	tst.AssertNoOptionToMoveTo(t, judge, Britain, "gib", "mor")
	tst.AssertOrderValidity(t, judge, orders.Move("gib", "and"), Britain, nil)
	tst.AssertOptionToMove(t, judge, Britain, "gib", "and")
}

func TestGibraltarConvoy(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("and", godip.Unit{godip.Army, Britain})

	// Test convoy via Gibraltar.
	opts := judge.Phase().Options(judge, Britain)
	tst.AssertOrderValidity(t, judge, orders.Move("and", "mor"), Britain, nil)
	tst.AssertOpt(t, opts, []string{"and", "MoveViaConvoy", "and", "mor"})
	tst.AssertOrderValidity(t, judge, orders.Convoy("gib", "and", "mor"), Britain, nil)
	tst.AssertOpt(t, opts, []string{"gib", "Convoy", "gib", "and", "mor"})
}
