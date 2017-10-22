package testing

import (
	"reflect"
	"testing"

	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

func AssertOrderValidity(t *testing.T, validator dip.Validator, order dip.Order, nat dip.Nation, err error) {
	if gotNat, e := order.Validate(validator); e != err {
		t.Errorf("%v should validate to %v, but got %v", order, err, e)
	} else if gotNat != nat {
		t.Errorf("%v should validate with %q as issuer, but got %q", order, nat, gotNat)
	}
}

func AssertMove(t *testing.T, j *state.State, src, dst dip.Province, success bool) {
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

func AssertUnit(t *testing.T, j *state.State, province dip.Province, unit dip.Unit) {
	if found, _, _ := j.Unit(province); !reflect.DeepEqual(found, unit) {
		t.Errorf("%v should be at %v now", unit, province)
	}
}

func AssertNoUnit(t *testing.T, j *state.State, province dip.Province) {
	_, _, ok := j.Unit(province)
	if ok {
		t.Errorf("There should be no unit at %v now", province)
	}
}

func AssertNoOptionToMoveTo(t *testing.T, j *state.State, nat dip.Nation, src dip.Province, dst dip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[cla.Move][dip.SrcProvince(src)][dst]; ok {
		t.Errorf("There should be no option for %v to move %v to %v", nat, src, dst)
	}
}

func AssertOptionToMove(t *testing.T, j *state.State, nat dip.Nation, src dip.Province, dst dip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[cla.Move][dip.SrcProvince(src)][dst]; !ok {
		t.Errorf("There should be an option for %v to move %v to %v", nat, src, dst)
	}
}
