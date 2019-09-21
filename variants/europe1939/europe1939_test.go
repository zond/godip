package europe1939

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
	judge, err := Europe1939Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(1901, godip.Spring, godip.Movement)
	judge := Europe1939Blank(startPhase)
	return judge
}

func TestLiverpoolNorthernIreland(t *testing.T) {
	judge := blankState(t)

	// Test the connection from Liverpool to Northern Ireland.
	judge.SetUnit("liv", godip.Unit{godip.Army, Britain})
	tst.AssertOrderValidity(t, judge, orders.Move("liv", "noi"), Britain, nil)

	// Check the reverse too.
	judge.RemoveUnit("liv")
	judge.SetUnit("noi", godip.Unit{godip.Army, Britain})
	tst.AssertOrderValidity(t, judge, orders.Move("noi", "liv"), Britain, nil)
}

func TestNapelsSicily(t *testing.T) {
	judge := blankState(t)

	// Test the connection from Napels to Sicily.
	judge.SetUnit("nap", godip.Unit{godip.Army, Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("nap", "sic"), Italy, nil)

	// Check the reverse too.
	judge.RemoveUnit("nap")
	judge.SetUnit("sic", godip.Unit{godip.Army, Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("sic", "nap"), Italy, nil)
}

func TestSerbianArmy(t *testing.T) {
	judge := startState(t)

	// Check a neutral army starts in Serbia.
	tst.AssertUnit(t, judge, godip.Province("ser"), godip.Unit{godip.Army, godip.Neutral})
	// Check no neutral army starts in e.g. Croatia.
	tst.AssertNoUnit(t, judge, godip.Province("cro"))

	// Try dislodging neutral army and check it is not rebuilt.
	judge.SetUnit("rum", godip.Unit{godip.Army, USSR})
	judge.SetUnit("tra", godip.Unit{godip.Army, USSR})
	judge.SetOrder("rum", orders.Move("rum", "ser"))
	judge.SetOrder("tra", orders.SupportMove("tra", "rum", "ser"))
	judge.Next()
	tst.AssertUnit(t, judge, godip.Province("ser"), godip.Unit{godip.Army, USSR})
	judge.Next()
	judge.SetOrder("ser", orders.Move("ser", "mac"))
	judge.Next()
	judge.Next()
	judge.Next()
	tst.AssertNoUnit(t, judge, godip.Province("ser"))
}

func TestAfricanSeaRoute(t *testing.T) {
	judge := blankState(t)

	// Test the connection from South Atlantic Ocean.
	judge.SetUnit("sao", godip.Unit{godip.Fleet, France})
	tst.AssertOrderValidity(t, judge, orders.Move("sao", "red"), France, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sao", "per"), France, nil)
	judge.RemoveUnit("sao")

	// Test the connection from Red Sea.
	judge.SetUnit("red", godip.Unit{godip.Fleet, France})
	tst.AssertOrderValidity(t, judge, orders.Move("red", "sao"), France, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("red", "per"), France, nil)
	judge.RemoveUnit("red")

	// Test the connection from Persian Gulf.
	judge.SetUnit("per", godip.Unit{godip.Fleet, France})
	tst.AssertOrderValidity(t, judge, orders.Move("per", "red"), France, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("per", "sao"), France, nil)
}
