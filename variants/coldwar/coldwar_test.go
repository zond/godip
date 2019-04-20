package coldwar

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
	judge, err := ColdWarStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(1960, godip.Spring, godip.Movement)
	judge := ColdWarBlank(startPhase)
	return judge
}

func TestPanama(t *testing.T) {
	judge := startState(t)

	// Test naval connections from Panama.
	judge.SetUnit("pan", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col/nc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col/wc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen/wc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen/ec"), USSR, nil)

	// Test connections for armies.
	judge.RemoveUnit("pan")
	judge.SetUnit("pan", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen"), USSR, nil)
}

func TestDenmarkSweden(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("nts", godip.Unit{godip.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "swe"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "den"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "bal"), "", godip.ErrIllegalMove)
	judge.RemoveUnit("nts")

	judge.SetUnit("bal", godip.Unit{godip.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "swe"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "den"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "nts"), "", godip.ErrIllegalMove)
	judge.RemoveUnit("bal")

	judge.SetUnit("swe", godip.Unit{godip.Army, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("swe", "den"), NATO, nil)
}

func TestArmyOceana(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("aus")

	// Test the land connections between Australia, Indonesia and the Philippeans.
	judge.SetUnit("ins", godip.Unit{godip.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "aus"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "phi"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), "", godip.ErrIllegalMove)
}

func TestEgypt(t *testing.T) {
	judge := startState(t)

	// Test can sail through Egypt.
	judge.SetUnit("egy", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("egy", "eme"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("egy", "red"), USSR, nil)
}

func TestKorea(t *testing.T) {
	// Test can convoy from North Korea.
	judge := startState(t)
	judge.SetUnit("nko", godip.Unit{godip.Army, USSR})
	judge.SetUnit("soj", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("nko", "jap"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("yel", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("nko", "phi"), USSR, nil)

	// Test can convoy to North Korea.
	judge = startState(t)
	judge.SetUnit("jap", godip.Unit{godip.Army, USSR})
	judge.SetUnit("soj", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "nko"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("phi", godip.Unit{godip.Army, USSR})
	judge.SetUnit("yel", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("phi", "nko"), USSR, nil)
}

func TestStraits(t *testing.T) {
	judge := blankState(t)

	// Test that several bodies of water require convoys.
	judge.SetUnit("wca", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("wca", "grd"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("arc", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("wca", "grd"), USSR, nil)

	judge.SetUnit("hav", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "mex"), "", godip.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "flo"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("gom", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "mex"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "flo"), USSR, nil)

	// Test convoys to and from Paris - it's worth a bit more testing because it has two coasts.
	judge.RemoveUnit("lon")
	judge.RemoveUnit("par")
	judge.SetUnit("par", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("par", "lon"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("nts", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("par", "lon"), USSR, nil)
	judge.RemoveUnit("par")
	judge.RemoveUnit("nts")
	judge.SetUnit("lon", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "par"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("nts", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "par"), USSR, nil)

	judge.SetUnit("ins", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("scs", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), USSR, nil)

	judge.SetUnit("jap", godip.Unit{godip.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "kam"), "", godip.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "vla"), "", godip.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "seo"), "", godip.ErrMissingConvoyPath)
	judge.SetUnit("ber", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "kam"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "vla"), USSR, nil)
	judge.SetUnit("yel", godip.Unit{godip.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "seo"), USSR, nil)
}

func TestIstanbul(t *testing.T) {
	judge := startState(t)

	// Test can sail through Istanbul
	tst.AssertOrderValidity(t, judge, orders.Move("ist", "bla"), NATO, nil)
}

func TestGreeceYugoslavia(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("yug", godip.Unit{godip.Army, USSR})

	// Test can drive to Greece
	tst.AssertOrderValidity(t, judge, orders.Move("yug", "grc"), USSR, nil)

	judge.RemoveUnit("yug")
	judge.SetUnit("yug", godip.Unit{godip.Fleet, USSR})

	// Test cannot sail to Greece
	tst.AssertOrderValidity(t, judge, orders.Move("yug", "grc"), "", godip.ErrIllegalMove)

	judge.RemoveUnit("yug")
	judge.SetUnit("grc", godip.Unit{godip.Army, USSR})

	// Test can drive to Yugoslavia
	tst.AssertOrderValidity(t, judge, orders.Move("grc", "yug"), USSR, nil)

	judge.RemoveUnit("grc")
	judge.SetUnit("grc", godip.Unit{godip.Fleet, USSR})

	// Test cannot sail to Yugoslavia
	tst.AssertOrderValidity(t, judge, orders.Move("grc", "yug"), "", godip.ErrIllegalMove)
}

func TestNewYorkToronto(t *testing.T) {
	judge := startState(t)

	// Test Great Lake prevents movement.
	tst.AssertOrderValidity(t, judge, orders.Move("nyk", "tor"), "", godip.ErrMissingConvoyPath)
    
	// ...and back again.
	judge.RemoveUnit("nyk")
	judge.SetUnit("tor", godip.Unit{godip.Army, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("tor", "nyk"), "", godip.ErrMissingConvoyPath)
}
