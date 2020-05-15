package testing

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
)

func PP(i interface{}) string {
	b, err := json.MarshalIndent(i, "  ", "  ")
	if err != nil {
		return spew.Sdump(i)
	}
	return string(b)
}

func AssertOrderValidity(t *testing.T, validator godip.Validator, order godip.Order, nat godip.Nation, err error) {
	if gotNat, e := order.Validate(validator); e != err {
		t.Errorf("%v should validate to %v, but got %v", order, err, e)
	} else if gotNat != nat {
		t.Errorf("%v should validate with %q as issuer, but got %q", order, nat, gotNat)
	}
}

func AssertMove(t *testing.T, j *state.State, src, dst godip.Province, success bool) {
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

func AssertUnit(t *testing.T, j *state.State, province godip.Province, unit godip.Unit) {
	if found, _, _ := j.Unit(province); !reflect.DeepEqual(found, unit) {
		t.Errorf("%v should be at %v now", unit, province)
	}
}

func AssertNoUnit(t *testing.T, j *state.State, province godip.Province) {
	_, _, ok := j.Unit(province)
	if ok {
		t.Errorf("There should be no unit at %v now", province)
	}
}

func AssertNoOptionToMoveTo(t *testing.T, j *state.State, nat godip.Nation, src godip.Province, dst godip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[godip.Move][godip.SrcProvince(src)][dst]; ok {
		t.Errorf("There should be no option for %v to move %v to %v", nat, src, dst)
	}
}

func AssertOptionToMove(t *testing.T, j *state.State, nat godip.Nation, src godip.Province, dst godip.Province) {
	options := j.Phase().Options(j, nat)[src]
	if _, ok := options[godip.Move][godip.SrcProvince(src)][dst]; !ok {
		t.Errorf("There should be an option for %v to move %v to %v", nat, src, dst)
	}
}

func hasOptHelper(opts map[string]interface{}, filter string, order []string, originalOpts map[string]interface{}, originalOrder []string) error {
	if len(order) == 0 {
		return nil
	}
	foundInter, found := opts[order[0]]
	if !found {
		return fmt.Errorf("Got no option for %+v in %v, failed at %+v in %v, wanted it!", originalOrder, PP(originalOpts), order, PP(opts))
	}
	foundMap := foundInter.(map[string]interface{})
	if filter != "" && foundMap["Filter"] != filter {
		return fmt.Errorf("Found %+v in %v, but didn't get the filter we wanted (%q)", originalOrder, PP(originalOpts), filter)
	}
	return hasOptHelper(foundMap["Next"].(map[string]interface{}), "", order[1:], originalOpts, originalOrder)
}

func hasOpt(opts godip.Options, order []string) error {
	return hasFilteredOpt(opts, "", order)
}

func hasFilteredOpt(opts godip.Options, filter string, order []string) error {
	b, err := json.MarshalIndent(opts, "  ", "  ")
	if err != nil {
		return err
	}
	converted := map[string]interface{}{}
	if err := json.Unmarshal(b, &converted); err != nil {
		return err
	}
	return hasOptHelper(converted, filter, order, converted, order)
}

func AssertOpt(t *testing.T, opts godip.Options, order []string) {
	AssertFilteredOpt(t, opts, "", order)
}

func AssertFilteredOpt(t *testing.T, opts godip.Options, filter string, order []string) {
	t.Run(strings.Join(order, "_")+"/"+filter, func(t *testing.T) {
		err := hasFilteredOpt(opts, filter, order)
		if err != nil {
			t.Error(err)
		}
	})
}

func AssertNoOpt(t *testing.T, opts godip.Options, order []string) {
	t.Run(strings.Join(order, "_"), func(t *testing.T) {
		err := hasOpt(opts, order)
		if err == nil {
			b, err := json.MarshalIndent(opts, "  ", "  ")
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("Found option for %+v in %s, didn't want it", order, b)
		}
	})
}

func AssertOwner(t *testing.T, j *state.State, supplyCenter string, owner godip.Nation) {
	nation, _, ok := j.SupplyCenter(godip.Province(supplyCenter))
	if !ok {
		t.Errorf("Province %s was not owned", supplyCenter)
	}
	if nation != owner {
		t.Errorf("Province %s was owned by %s, but expected %s", supplyCenter, nation, owner)
	}
}

func AssertNoOwner(t *testing.T, j *state.State, supplyCenter string) {
	nation, _, ok := j.SupplyCenter(godip.Province(supplyCenter))
	if ok {
		t.Errorf("Province %s was owned by %s, but expected no owner", supplyCenter, nation)
	}
}

// Wait for the given number of phases.
func WaitForPhases(judge *state.State, phases int) {
	for phase := 0; phase < phases; phase++ {
		judge.Next()
	}
}

// Increase the current phase by the given number of years.
func WaitForYears(judge *state.State, years int) {
	WaitForPhases(judge, 5*years)
}
