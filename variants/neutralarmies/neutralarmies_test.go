package neutralarmies

import (
	"testing"

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
	judge, err := Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.Phase(1901, godip.Spring, godip.Movement)
	judge := classical.Blank(startPhase)
	return judge
}

func TestBounceWithNeutralArmy(t *testing.T) {
	judge := startState(t)
	// Spring movement: Austria tries to take Rumania.
	judge.SetOrder("bud", orders.Move("bud", "rum"))
	judge.Next()
	tst.AssertUnit(t, judge, "rum", godip.Unit{godip.Army, godip.Neutral})
	// Sprint retreat
	judge.Next()
	// Fall movement: Russia supports Austria
	judge.SetOrder("bud", orders.Move("bud", "rum"))
	judge.SetOrder("sev", orders.SupportMove("sev", "bud", "rum"))
	judge.Next()
	tst.AssertUnit(t, judge, "rum", godip.Unit{godip.Army, godip.Austria})
}

func TestNeutralArmyRebuilt(t *testing.T) {
	judge := startState(t)
	// Remove the units from Portugal and Spain.
	judge.RemoveUnit("por")
	judge.RemoveUnit("spa")
	// Give Spain to France (but leave it vacant).
	judge.SetSC("spa", godip.France)

	// Spring movement
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.Next()
	// Fall retreat
	judge.Next()

	// Check that both SCs are still vacant.
	tst.AssertNoUnit(t, judge, "por")
	tst.AssertNoUnit(t, judge, "spa")

	// Fall adjustment - Check that only the Neutral army in Portugal is rebuilt.
	judge.Next()
	tst.AssertUnit(t, judge, "por", godip.Unit{godip.Army, godip.Neutral})
	tst.AssertNoUnit(t, judge, "spa")
}
