package coldwar

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
	judge, err := ColdWarStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.ClassicalPhase(1960, cla.Spring, cla.Movement)
	judge := ColdWarBlank(startPhase)
	return judge
}

func TestPanama(t *testing.T) {
	judge := startState(t)

	// Test naval connections from Panama.
	judge.SetUnit("pan", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col/nc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col/wc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen/wc"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen/ec"), USSR, nil)

	// Test connections for armies.
	judge.RemoveUnit("pan")
	judge.SetUnit("pan", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "col"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("pan", "cen"), USSR, nil)
}

func TestDenmarkSweden(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("nts", dip.Unit{cla.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "swe"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "den"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nts", "bal"), "", cla.ErrIllegalMove)
	judge.RemoveUnit("nts")

	judge.SetUnit("bal", dip.Unit{cla.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "swe"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "den"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("bal", "nts"), "", cla.ErrIllegalMove)
	judge.RemoveUnit("bal")

	judge.SetUnit("swe", dip.Unit{cla.Army, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("swe", "den"), NATO, nil)
}

func TestArmyOceana(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("aus")

	// Test the land connections between Australia, Indonesia and the Philippeans.
	judge.SetUnit("ins", dip.Unit{cla.Fleet, NATO})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "aus"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "phi"), NATO, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), "", cla.ErrIllegalMove)
}

func TestEgypt(t *testing.T) {
	judge := startState(t)

	// Test can sail through Egypt.
	judge.SetUnit("egy", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("egy", "eme"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("egy", "red"), USSR, nil)
}

func TestKorea(t *testing.T) {
	// Test can convoy from North Korea.
	judge := startState(t)
	judge.SetUnit("nko", dip.Unit{cla.Army, USSR})
	judge.SetUnit("soj", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("nko", "jap"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("nko", "phi"), USSR, nil)

	// Test can convoy to North Korea.
	judge = startState(t)
	judge.SetUnit("jap", dip.Unit{cla.Army, USSR})
	judge.SetUnit("soj", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "nko"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("phi", dip.Unit{cla.Army, USSR})
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("phi", "nko"), USSR, nil)
}

func TestStraits(t *testing.T) {
	judge := blankState(t)

	// Test that several bodies of water require convoys.
	judge.SetUnit("wca", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("wca", "grd"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("arc", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("wca", "grd"), USSR, nil)

	judge.SetUnit("hav", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "mex"), "", cla.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "flo"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("gom", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "mex"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("hav", "flo"), USSR, nil)

	// Test convoys to and from Paris - it's worth a bit more testing because it has two coasts.
	judge.RemoveUnit("lon")
	judge.RemoveUnit("par")
	judge.SetUnit("par", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("par", "lon"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("nts", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("par", "lon"), USSR, nil)
	judge.RemoveUnit("par")
	judge.RemoveUnit("nts")
	judge.SetUnit("lon", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "par"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("nts", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("lon", "par"), USSR, nil)

	judge.SetUnit("ins", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("scs", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("ins", "sea"), USSR, nil)

	judge.SetUnit("jap", dip.Unit{cla.Army, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "kam"), "", cla.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "vla"), "", cla.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "seo"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("ber", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "kam"), USSR, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "vla"), USSR, nil)
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	tst.AssertOrderValidity(t, judge, orders.Move("jap", "seo"), USSR, nil)
}

func TestIstanbul(t *testing.T) {
	judge := startState(t)

	// Test can sail through Istanbul
	tst.AssertOrderValidity(t, judge, orders.Move("ist", "bla"), NATO, nil)
}
