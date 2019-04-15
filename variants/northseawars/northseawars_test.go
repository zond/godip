package northseawars

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
	judge, err := NorthSeaWarsStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(0, godip.Spring, godip.Movement)
	judge := NorthSeaWarsBlank(startPhase)
	return judge
}

func TestCentralNorthSea(t *testing.T) {
	judge := blankState(t)

	// Test resource connections from Central North Sea.
	judge.SetUnit("cns", godip.Unit{godip.Fleet, Norse})
	tst.AssertOrderValidity(t, judge, orders.Move("cns", "woo"), Norse, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("cns", "iro"), Norse, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("cns", "gra"), Norse, nil)
	opts := judge.Phase().Options(judge, Norse)
	tst.AssertOpt(t, opts, []string{"cns", "Move", "cns", "woo"})
	tst.AssertOpt(t, opts, []string{"cns", "Move", "cns", "iro"})
	tst.AssertOpt(t, opts, []string{"cns", "Move", "cns", "gra"})

	// Test no connections in reverse direction.
	judge.SetUnit("woo", godip.Unit{godip.Fleet, Norse})
	tst.AssertOrderValidity(t, judge, orders.Move("woo", "cns"), "", godip.ErrIllegalMove)
	judge.SetUnit("iro", godip.Unit{godip.Fleet, Norse})
	tst.AssertOrderValidity(t, judge, orders.Move("iro", "cns"), "", godip.ErrIllegalMove)
	judge.SetUnit("gra", godip.Unit{godip.Fleet, Norse})
	tst.AssertOrderValidity(t, judge, orders.Move("gra", "cns"), "", godip.ErrIllegalMove)
	opts = judge.Phase().Options(judge, Norse)
	tst.AssertNoOpt(t, opts, []string{"woo", "Move", "woo", "cns"})
	tst.AssertNoOpt(t, opts, []string{"iro", "Move", "iro", "cns"})
	tst.AssertNoOpt(t, opts, []string{"gra", "Move", "gra", "cns"})
}

func TestConvoyToResource(t *testing.T) {
	judge := startState(t)

	// Add fleets to allow convoy.
	judge.SetUnit("ens", godip.Unit{godip.Fleet, Frysians})
	judge.SetUnit("cns", godip.Unit{godip.Fleet, Frysians})

	tst.AssertOrderValidity(t, judge, orders.Move("ams", "woo"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ams", "iro"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ams", "gra"), Frysians, nil)
	opts := judge.Phase().Options(judge, Frysians)
	tst.AssertOpt(t, opts, []string{"ams", "MoveViaConvoy", "ams", "woo"})
	tst.AssertOpt(t, opts, []string{"ams", "MoveViaConvoy", "ams", "iro"})
	tst.AssertOpt(t, opts, []string{"ams", "MoveViaConvoy", "ams", "gra"})
}

func TestSupportCentralNorthSea(t *testing.T) {
	judge := blankState(t)

	// Test resource connections from Central North Sea.
	judge.SetUnit("cns", godip.Unit{godip.Fleet, Norse})
	judge.SetUnit("gra", godip.Unit{godip.Fleet, Norse})
	tst.AssertOrderValidity(t, judge, orders.SupportMove("cns", "gra", "woo"), Norse, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("gra", "cns", "woo"), Norse, nil)
	opts := judge.Phase().Options(judge, Norse)
	tst.AssertOpt(t, opts, []string{"cns", "Support", "gra", "woo"})
	tst.AssertOpt(t, opts, []string{"gra", "Support", "cns", "woo"})
}

func TestSealandArmy(t *testing.T) {
	judge := blankState(t)

	// Add fleets to allow convoy.
	judge.SetUnit("sea", godip.Unit{godip.Army, Frysians})

	tst.AssertOrderValidity(t, judge, orders.Move("sea", "got"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "mag"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "ams"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "jut"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "lim"), Frysians, nil)
	opts := judge.Phase().Options(judge, Frysians)
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "got"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "mag"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "ams"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "jut"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "lim"})

	// Check can convoy via Skagerrak.
	judge.SetUnit("ska", godip.Unit{godip.Fleet, Frysians})
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "ost"), Frysians, nil)
	opts = judge.Phase().Options(judge, Frysians)
	tst.AssertOpt(t, opts, []string{"sea", "MoveViaConvoy", "sea", "ost"})
}

func TestSealandFleet(t *testing.T) {
	judge := blankState(t)

	// Add fleets to allow convoy.
	judge.SetUnit("sea", godip.Unit{godip.Fleet, Frysians})

	tst.AssertOrderValidity(t, judge, orders.Move("sea", "lim"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "got"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "mag"), "", godip.ErrIllegalDestination)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "ams"), "", godip.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "jut/ec"), Frysians, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "jut/wc"), "", godip.ErrIllegalMove)
	// A fleet moving Sealand to Jutland is automatically interpretted as "jut/ec".
	tst.AssertOrderValidity(t, judge, orders.Move("sea", "jut"), Frysians, nil)
	opts := judge.Phase().Options(judge, Frysians)
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "lim"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "got"})
	tst.AssertNoOpt(t, opts, []string{"sea", "Move", "sea", "mag"})
	tst.AssertNoOpt(t, opts, []string{"sea", "Move", "sea", "ams"})
	tst.AssertOpt(t, opts, []string{"sea", "Move", "sea", "jut/ec"})
	tst.AssertNoOpt(t, opts, []string{"sea", "Move", "sea", "jut/wc"})
	tst.AssertNoOpt(t, opts, []string{"sea", "Move", "sea", "jut"})
}
