package spiceislands

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
	judge, err := SpiceIslandsStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

 func TestSpiceIslandsBuildAnywhere(t *testing.T) {
	judge := startState(t)

	// Give Brunei an extra SC in Sambas.
	judge.SetSC("sab", Brunei)

	// Spring movement
	judge.SetOrder("bru", orders.Move("bru", "neg"))
	judge.Next()
	// Spring retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("tun", orders.Move("tun", "kut"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Sambas.
	judge.SetOrder("sab", orders.BuildAnywhere("sab", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "sab", godip.Unit{godip.Army, Brunei})
}
