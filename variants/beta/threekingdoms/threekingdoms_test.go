package threekingdoms

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
	judge, err := ThreeKingdomsStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

 func TestThreeKingdomBuildAnywhere(t *testing.T) {
	judge := startState(t)

	// Give Shu an extra SC in England.
	judge.SetSC("jig", Shu)

	// Spring movement
	judge.SetOrder("che", orders.Move("che", "sha"))
	judge.Next()
	// Spring retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("zha", orders.Move("zha", "hep"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Scotland.
	judge.SetOrder("jig", orders.BuildAnywhere("jig", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "jig", godip.Unit{godip.Army, Shu})
}
