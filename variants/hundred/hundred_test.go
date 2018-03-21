package hundred

import (
	"testing"

	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := HundredStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.Phase(1901, cla.Spring, cla.Movement)
	judge := HundredBlank(startPhase)
	return judge
}

func TestLondonCalais(t *testing.T) {
	judge := blankState(t)

	// Test the connection from London to Calais.
	judge.SetUnit("lon", dip.Unit{cla.Army, England})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "cal"), England, nil)

	// Check the reverse too.
	judge.RemoveUnit("lon")
	judge.SetUnit("cal", dip.Unit{cla.Army, England})
	tst.AssertOrderValidity(t, judge, orders.Move("cal", "lon"), England, nil)
}
