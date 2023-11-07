package southofsahara

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
	judge, err := SouthofSaharaStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

 func TestSouthofSaharaBuildAnywhere(t *testing.T) {
	judge := startState(t)

	// Give Benin an extra SC in Kongo.
	judge.SetSC("kon", Benin)

	// Spring movement
	judge.SetOrder("ife", orders.Move("ife", "oyo"))
	judge.Next()
	// Spring retreat
	judge.Next()
	// Fall movement
	judge.SetOrder("edo", orders.Move("edo", "ije"))
	judge.Next()
	// Fall retreat
	judge.Next()

	// Fall adjustment - Try to build a new Army in Kongo.
	judge.SetOrder("kon", orders.BuildAnywhere("kon", godip.Army, time.Now()))
	judge.Next()
	// Check that it was successful.
	tst.AssertUnit(t, judge, "kon", godip.Unit{godip.Army, Benin})
}
