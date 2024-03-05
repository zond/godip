package gatewaywest

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
	judge, err := GatewayWestStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

 func TestGatewayWestBuildAnywhere(t *testing.T) {
	judge := startState(t)

	// Give Illini an extra SC.
	judge.SetSC("mar", Illini)

	// Spring movement
	judge.SetOrder("lin", orders.Move("lin", "aud"))
	judge.Next()
	// Spring retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("aud", orders.Move("aud", "cal"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Scotland.
	judge.SetOrder("mar", orders.BuildAnywhere("mar", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "mar", godip.Unit{godip.Army, Illini})
}
