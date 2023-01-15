package unconstitutional

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	tst "github.com/zond/godip/variants/testing"
)

func TestBuildAnywhere(t *testing.T) {
	j, err := UnconstitutionalStart()
	if err != nil {
		t.Fatal(err)
	}
	// Remove New York's unit from New York City and replace the neutral unit in Turks and Caicos with a New York fleet.
	j.RemoveUnit("nyc")
	j.RemoveUnit("tur")
	j.SetUnit("tur", godip.Unit{godip.Fleet, NewYork})
	// Wait for New York to capture the SC.
	tst.WaitForPhases(j, 4)
	// Remove the fleet from Turks and Caicos.
	j.RemoveUnit("tur")
	tst.AssertNoUnit(t, j, "tur")

	tst.WaitForPhases(j, 5)

	// Test that New York has the option to build in Turks and Caicos.
	tst.AssertOpt(t, j.Phase().Options(j, NewYork), []string{"tur", "Build", "Fleet", "tur"})
	tst.AssertOrderValidity(t, j, orders.BuildAnywhere("tur", godip.Fleet, time.Now()), NewYork, nil)

	// Test that build anywhere works.
	j.SetOrders(map[godip.Province]godip.Adjudicator{
		"nyc": orders.BuildAnywhere("tur", godip.Fleet, time.Now()),
	})
	j.Next()
	tst.AssertUnit(t, j, "tur", godip.Unit{godip.Fleet, "New York"})
}
