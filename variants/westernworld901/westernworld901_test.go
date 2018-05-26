package westernworld901

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
	judge, err := WesternWorld901Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestBounceWithNeutralArmy(t *testing.T) {
	judge := startState(t)
	// Spring movement: West Frankish Kingdom tries to take Lothairingia.
	judge.SetOrder("par", orders.Move("par", "lot"))
	judge.Next()
	tst.AssertUnit(t, judge, "lot", godip.Unit{godip.Army, godip.Neutral})
	// Sprint retreat
	judge.Next()
	// Fall movement: East Frankish Kingdom supports West Frankish Kingdom
	judge.SetOrder("par", orders.Move("par", "lot"))
	judge.SetOrder("swa", orders.SupportMove("swa", "par", "lot"))
	judge.Next()
	tst.AssertUnit(t, judge, "lot", godip.Unit{godip.Fleet, WestFrankishKingdom})
}

func TestNeutralArmyRebuilt(t *testing.T) {
	judge := startState(t)
	// Remove the units from Esteland (Neutral), Novgorod (Principality of Kiev) and Bulgar (Neutral).
	judge.RemoveUnit("est")
	judge.RemoveUnit("nov")
	judge.RemoveUnit("bul")
	// Give Esteland to Principality of Kiev (but leave it vacant).
	judge.SetSC("est", PrincipalityofKiev)

	// Spring movement
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.Next()
	// Fall retreat
	judge.Next()

	// Check that all SCs are still vacant.
	tst.AssertNoUnit(t, judge, "est")
	tst.AssertNoUnit(t, judge, "nov")
	tst.AssertNoUnit(t, judge, "bul")

	// Fall adjustment - Add explicit order from Principality of Kiev to rebuild Novgorod.
	judge.SetOrder("nov", orders.Build("nov", godip.Army, time.Now()))
	judge.Next()
	tst.AssertNoUnit(t, judge, "est")
	tst.AssertUnit(t, judge, "nov", godip.Unit{godip.Army, PrincipalityofKiev})
	tst.AssertUnit(t, judge, "bul", godip.Unit{godip.Army, godip.Neutral})
}
