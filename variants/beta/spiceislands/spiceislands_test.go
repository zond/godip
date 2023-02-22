package sengoku

import (
	"testing"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"

	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := SengokuStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestDislodgeNeutralArmy(t *testing.T) {
	judge := startState(t)

	// Set orders for Shimazu to dislodge the neutral army in Hizen.
	// Spring movement
	judge.SetOrder("osu", orders.Move("osu", "eas"))
	judge.SetOrder("sat", orders.Move("sat", "hig"))
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("eas", orders.Move("eas", "hiz"))
	judge.SetOrder("hig", orders.SupportMove("hig", "eas", "hiz"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Check that all SCs are still vacant.
	tst.AssertUnit(t, judge, "hiz", godip.Unit{godip.Fleet, Shimazu})
}

func TestNeutralArmyRebuilt(t *testing.T) {
	judge := startState(t)

	// Remove neutral army from Hizen.
	judge.RemoveUnit("hiz")

	// Spring movement
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.Next()
	// Fall retreat
	judge.Next()
	// Winter adjustment
	judge.Next()

	// Check rebuilt.
	tst.AssertUnit(t, judge, "hiz", godip.Unit{godip.Army, godip.Neutral})
}
