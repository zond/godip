package coldwar

import (
	"reflect"
	"testing"

	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

func init() {
	dip.Debug = true
}

func assertOrderValidity(t *testing.T, validator dip.Validator, order dip.Order, nat dip.Nation, err error) {
	if gotNat, e := order.Validate(validator); e != err {
		t.Errorf("%v should validate to %v, but got %v", order, err, e)
	} else if gotNat != nat {
		t.Errorf("%v should validate with %q as issuer, but got %q", order, nat, gotNat)
	}
}

func assertMove(t *testing.T, j *state.State, src, dst dip.Province, success bool) {
	if success {
		unit, _, ok := j.Unit(src)
		if !ok {
			t.Errorf("Should be a unit at %v", src)
		}
		j.SetOrder(src, orders.Move(src, dst))
		j.Next()
		if err, ok := j.Resolutions()[src]; ok && err != nil {
			t.Errorf("Move from %v to %v should have worked, got %v", src, dst, err)
		}
		if now, _, ok := j.Unit(src); ok && reflect.DeepEqual(now, unit) {
			t.Errorf("%v should have moved from %v", now, src)
		}
		if now, _, _ := j.Unit(dst); !reflect.DeepEqual(now, unit) {
			t.Errorf("%v should be at %v now", unit, dst)
		}
	} else {
		unit, _, _ := j.Unit(src)
		j.SetOrder(src, orders.Move(src, dst))
		j.Next()
		if err, ok := j.Resolutions()[src]; !ok || err == nil {
			t.Errorf("Move from %v to %v should not have worked", src, dst)
		}
		if now, _, _ := j.Unit(src); !reflect.DeepEqual(now, unit) {
			t.Errorf("%v should not have moved from %v", now, src)
		}
	}
}

func assertUnit(t *testing.T, j *state.State, province dip.Province, unit dip.Unit) {
	if found, _, _ := j.Unit(province); !reflect.DeepEqual(found, unit) {
		t.Errorf("%v should be at %v now", unit, province)
	}
}

func assertNoUnit(t *testing.T, j *state.State, province dip.Province) {
	_, _, ok := j.Unit(province)
	if ok {
		t.Errorf("There should be no unit at %v now", province)
	}
}

func assertNoOptionToMoveTo(t *testing.T, j *state.State, nat dip.Nation, src dip.Province, dst dip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[cla.Move][dip.SrcProvince(src)][dst]; ok {
		t.Errorf("There should be no option for %v to move %v to %v", nat, src, dst)
	}
}

func assertOptionToMove(t *testing.T, j *state.State, nat dip.Nation, src dip.Province, dst dip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[cla.Move][dip.SrcProvince(src)][dst]; !ok {
		t.Errorf("There should be an option for %v to move %v to %v", nat, src, dst)
	}
}

func startState(t *testing.T) *state.State {
	judge, err := ColdWarStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestPanama(t *testing.T) {
	judge := startState(t)

	// Test naval connections from Panama.
	judge.SetUnit("pan", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("pan", "col/nc"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("pan", "col/wc"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("pan", "cen/wc"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("pan", "cen/ec"), USSR, nil)

	// Test connections for armies.
	judge.RemoveUnit("pan")
	judge.SetUnit("pan", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("pan", "col"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("pan", "cen"), USSR, nil)
}

func TestDenmarkSweden(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("nts", dip.Unit{cla.Fleet, NATO})
	assertOrderValidity(t, judge, orders.Move("nts", "swe"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("nts", "den"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("nts", "bal"), "", cla.ErrIllegalMove)
	judge.RemoveUnit("nts")

	judge.SetUnit("bal", dip.Unit{cla.Fleet, NATO})
	assertOrderValidity(t, judge, orders.Move("bal", "swe"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("bal", "den"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("bal", "nts"), "", cla.ErrIllegalMove)
	judge.RemoveUnit("bal")

	judge.SetUnit("swe", dip.Unit{cla.Army, NATO})
	assertOrderValidity(t, judge, orders.Move("swe", "den"), NATO, nil)
}

func TestArmyOceana(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("aus")

	// Test the land connections between Australia, Indonesia and the Philippeans.
	judge.SetUnit("ins", dip.Unit{cla.Fleet, NATO})
	assertOrderValidity(t, judge, orders.Move("ins", "aus"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("ins", "phi"), NATO, nil)
	assertOrderValidity(t, judge, orders.Move("ins", "sta"), "", cla.ErrIllegalMove)
}

func TestEgypt(t *testing.T) {
	judge := startState(t)

	// Test can sail through Egypt.
	judge.SetUnit("egy", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("egy", "etm"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("egy", "red"), USSR, nil)
}

func TestKorea(t *testing.T) {
	// Test can convoy from North Korea.
	judge := startState(t)
	judge.SetUnit("nok", dip.Unit{cla.Army, USSR})
	judge.SetUnit("soj", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("nok", "jap"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("nok", "phi"), USSR, nil)

	// Test can convoy to North Korea.
	judge = startState(t)
	judge.SetUnit("jap", dip.Unit{cla.Army, USSR})
	judge.SetUnit("soj", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("jap", "nok"), USSR, nil)
	judge.RemoveUnit("soj")
	judge.SetUnit("phi", dip.Unit{cla.Army, USSR})
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("phi", "nok"), USSR, nil)
}

func TestStraits(t *testing.T) {
	judge := startState(t)

	// Test that several bodies of water require convoys.
	judge.SetUnit("wtn", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("wtn", "grd"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("arc", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("wtn", "grd"), USSR, nil)

	judge.SetUnit("flo", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("flo", "hav"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("gum", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("flo", "hav"), USSR, nil)

	// Test convoys to and from Paris - it's worth a bit more testing because it has two coasts.
	judge.RemoveUnit("lon")
	judge.RemoveUnit("par")
	judge.SetUnit("par", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("par", "lon"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("nts", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("par", "lon"), USSR, nil)
	judge.RemoveUnit("par")
	judge.RemoveUnit("nts")
	judge.SetUnit("lon", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("lon", "par"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("nts", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("lon", "par"), USSR, nil)

	judge.SetUnit("ins", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("ins", "sta"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("soc", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("ins", "sta"), USSR, nil)

	judge.SetUnit("jap", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("jap", "kam"), "", cla.ErrMissingConvoyPath)
	assertOrderValidity(t, judge, orders.Move("jap", "vla"), "", cla.ErrMissingConvoyPath)
	assertOrderValidity(t, judge, orders.Move("jap", "seo"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("ber", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("jap", "kam"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("jap", "vla"), USSR, nil)
	judge.SetUnit("yel", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("jap", "seo"), USSR, nil)

	judge.SetUnit("hav", dip.Unit{cla.Army, USSR})
	assertOrderValidity(t, judge, orders.Move("hav", "mex"), "", cla.ErrMissingConvoyPath)
	assertOrderValidity(t, judge, orders.Move("hav", "flo"), "", cla.ErrMissingConvoyPath)
	judge.SetUnit("gum", dip.Unit{cla.Fleet, USSR})
	assertOrderValidity(t, judge, orders.Move("hav", "mex"), USSR, nil)
	assertOrderValidity(t, judge, orders.Move("hav", "flo"), USSR, nil)
}

func TestIstanbul(t *testing.T) {
	judge := startState(t)

	// Test can sail through Istanbul
	assertOrderValidity(t, judge, orders.Move("ist", "bla"), NATO, nil)
}
