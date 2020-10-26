package year1908

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
	judge, err := Year1908Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestMidAtlanticCairo(t *testing.T) {
	judge := startState(t)

	// Test there's a connection from Cairo to Mid-Atlantic Ocean.
	tst.AssertOrderValidity(t, judge, orders.Move("cai", "mid"), "Britain", nil)
	judge.SetOrder("cai", orders.Move("cai", "mid"))
	judge.Next()
	judge.Next()
	tst.AssertUnit(t, judge, godip.Province("mid"), godip.Unit{godip.Fleet, Britain})

	// Test there's a reverse connection too (and that it can be used to convoy the French army in Casablanca).
	tst.AssertOrderValidity(t, judge, orders.Convoy("mid", "cas", "cai"), "Britain", nil)
	judge.SetOrder("mid", orders.Convoy("mid", "cas", "cai"))
	tst.AssertOrderValidity(t, judge, orders.Move("cas", "cai"), "France", nil)
	judge.SetOrder("cas", orders.Move("cas", "cai"))
	judge.Next()
	tst.AssertUnit(t, judge, godip.Province("cai"), godip.Unit{godip.Army, France})
}
