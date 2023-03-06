package maharajah

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"

	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := MaharajahStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

 func TestMaharajahBuildAnywhere(t *testing.T) {
	judge := startState(t)

	// Give Delhi an extra SC in Ayu.
	judge.SetSC("ayu", Delhi)

	// Spring movement
	judge.SetOrder("muz", orders.Move("muz", "beg"))
	judge.Next()
	// Spring retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("awa", orders.Move("awa", "kam"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Ayu.
	judge.SetOrder("ayu", orders.BuildAnywhere("ayu", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "ayu", godip.Unit{godip.Army, Delhi})
}
