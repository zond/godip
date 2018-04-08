package hundred

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
	judge, err := HundredStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.Phase(1901, godip.Spring, godip.Movement)
	judge := HundredBlank(startPhase)
	return judge
}

func TestLondonCalais(t *testing.T) {
	judge := blankState(t)

	// Test the connection from London to Calais.
	judge.SetUnit("lon", godip.Unit{godip.Army, England})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "cal"), England, nil)

	// Check the reverse too.
	judge.RemoveUnit("lon")
	judge.SetUnit("cal", godip.Unit{godip.Army, England})
	tst.AssertOrderValidity(t, judge, orders.Move("cal", "lon"), England, nil)
}

func TestBuildAnywhere(t *testing.T) {
	judge := startState(t)
	// Give England an extra SC in England.
	judge.SetSC("sco", England)

	// Spring movement
	judge.SetOrder("lon", orders.Move("lon", "str"))
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("str", orders.Move("str", "hol"))
	judge.SetOrder("hol", orders.Move("hol", "fri"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Scotland.
	judge.SetOrder("sco", orders.BuildAnywhere("sco", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "sco", godip.Unit{godip.Army, England})
}

func TestDisbandFrenchUnit(t *testing.T) {
	judge := startState(t)
	// France starts with one more unit than SC.

	// Spring movement
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Disband Army in Provence.
	judge.SetOrder("pro", orders.Disband("pro", time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertNoUnit(t, judge, "pro")
}
