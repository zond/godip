package classical

import (
	"os"
	"reflect"
	"testing"

	cla "github.com/zond/godip/classical/common"
	"github.com/zond/godip/classical/orders"
	dip "github.com/zond/godip/common"
	"github.com/zond/godip/datc"
	"github.com/zond/godip/state"
)

func init() {
	dip.Debug = true
}

func assertOrderValidity(t *testing.T, validator dip.Validator, order dip.Order, err error) {
	if e := order.Validate(validator); e != err {
		t.Errorf("%v should validate to %v, but got %v", order, err, e)
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

func startState(t *testing.T) *state.State {
	judge, err := Start()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestSupportValidation(t *testing.T) {
	judge := startState(t)
	// Happy paths
	assertOrderValidity(t, judge, orders.SupportMove("bre", "par", "gas"), nil)
	assertOrderValidity(t, judge, orders.SupportHold("par", "bre"), nil)
	assertOrderValidity(t, judge, orders.SupportMove("par", "bre", "gas"), nil)
	judge.SetUnit("spa/sc", dip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("por", dip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("gol", dip.Unit{cla.Fleet, cla.France})
	assertOrderValidity(t, judge, orders.SupportMove("spa/sc", "por", "mid"), nil)
	assertOrderValidity(t, judge, orders.SupportMove("gol", "mar", "spa"), nil)
	// Missing unit
	assertOrderValidity(t, judge, orders.SupportMove("ruh", "kie", "hol"), cla.ErrMissingUnit)
	// Missing supportee
	assertOrderValidity(t, judge, orders.SupportHold("ber", "sil"), cla.ErrMissingSupportUnit)
	// Illegal support
	assertOrderValidity(t, judge, orders.SupportHold("bre", "par"), cla.ErrIllegalSupportPosition)
	assertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "por"), cla.ErrIllegalSupportDestination)
	judge.RemoveUnit("spa/sc")
	judge.SetUnit("spa/nc", dip.Unit{cla.Fleet, cla.France})
	assertOrderValidity(t, judge, orders.SupportMove("spa/nc", "mar", "gol"), cla.ErrIllegalSupportDestination)
	// Illegal moves
	assertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "bur"), cla.ErrIllegalSupportMove)
}

func TestMoveValidation(t *testing.T) {
	judge := startState(t)
	// Happy path fleet
	assertOrderValidity(t, judge, orders.Move("bre", "mid"), nil)
	// Happy path army
	assertOrderValidity(t, judge, orders.Move("mun", "ruh"), nil)
	// Too far
	assertOrderValidity(t, judge, orders.Move("bre", "wes"), cla.ErrIllegalMove)
	// Fleet on land
	assertOrderValidity(t, judge, orders.Move("bre", "par"), cla.ErrIllegalDestination)
	// Army at sea
	assertOrderValidity(t, judge, orders.Move("smy", "eas"), cla.ErrIllegalDestination)
	// Unknown source
	assertOrderValidity(t, judge, orders.Move("a", "mid"), cla.ErrInvalidSource)
	// Unknown destination
	assertOrderValidity(t, judge, orders.Move("bre", "a"), cla.ErrInvalidDestination)
	// Missing sea path
	assertOrderValidity(t, judge, orders.Move("par", "mos"), cla.ErrMissingConvoyPath)
	// No unit
	assertOrderValidity(t, judge, orders.Move("spa", "por"), cla.ErrMissingUnit)
	// Working convoy
	judge.SetUnit("eng", dip.Unit{cla.Fleet, cla.England})
	judge.SetUnit("wal", dip.Unit{cla.Army, cla.England})
	assertOrderValidity(t, judge, orders.Move("wal", "bre"), nil)
	// Missing convoy
	assertOrderValidity(t, judge, orders.Move("wal", "gas"), cla.ErrMissingConvoyPath)
}

func TestMoveAdjudication(t *testing.T) {
	assertMove(t, startState(t), "bre", "mid", true)
	assertMove(t, startState(t), "stp/sc", "bot", true)
	assertMove(t, startState(t), "vie", "bud", false)
	assertMove(t, startState(t), "mid", "nat", false)
}

func testDATC(t *testing.T, statePair *datc.StatePair) {
	var s *state.State
	if statePair.Before.Phase == nil {
		s = Blank(&phase{
			year:   1901,
			season: cla.Spring,
			typ:    cla.Movement,
		})
	} else {
		s = Blank(statePair.Before.Phase)
	}
	s.SetUnits(statePair.Before.Units)
	s.SetDislodgeds(statePair.Before.Dislodgeds)
	s.SetSupplyCenters(statePair.Before.SCs)
	for prov, order := range statePair.Before.Orders {
		if s.Phase().Type() == cla.Movement {
			if u, _, ok := s.Unit(prov); ok && u.Nation == order.Nation {
				s.SetOrder(prov, order.Order)
			}
		} else if s.Phase().Type() == cla.Retreat {
			if u, _, ok := s.Dislodged(prov); ok && u.Nation == order.Nation {
				s.SetOrder(prov, order.Order)
			}
		} else if s.Phase().Type() == cla.Adjustment {
			if order.Order.Type() == cla.Build {
				if n, _, ok := s.SupplyCenter(prov); ok && n == order.Nation {
					s.SetOrder(prov, order.Order)
				}
			} else if order.Order.Type() == cla.Disband {
				if u, _, ok := s.Unit(prov); ok && u.Nation == order.Nation {
					s.SetOrder(prov, order.Order)
				}
			}
		} else {
			t.Fatalf("Unsupported phase type %v", s.Phase().Type())
		}
	}
	for _, order := range statePair.Before.FailedOrders {
		if order.Order.Type() == cla.Move && !order.Order.Flags()[cla.ViaConvoy] {
			s.AddBounce(order.Order.Targets()[0], order.Order.Targets()[1])
		}
	}
	for _, order := range statePair.Before.SuccessfulOrders {
		if order.Order.Type() == cla.Move && !order.Order.Flags()[cla.ViaConvoy] {
			s.SetDislodger(order.Order.Targets()[0], order.Order.Targets()[1])
		}
	}
	s.Next()
	err := false
	for prov, unit := range statePair.After.Units {
		if found, ok := s.Units()[prov]; ok {
			if !found.Equal(unit) {
				err = true
				t.Errorf("%v: Expected %v in %v, but found %v", statePair.Case, unit, prov, found)
			}
		} else {
			err = true
			t.Errorf("%v: Expected %v in %v, but found nothing", statePair.Case, unit, prov)
		}
	}
	for prov, unit := range statePair.After.Dislodgeds {
		if found, ok := s.Dislodgeds()[prov]; ok {
			if !found.Equal(unit) {
				err = true
				t.Errorf("%v: Expected %v dislodged in %v, but found %v", statePair.Case, unit, prov, found)
			}
		} else {
			err = true
			t.Errorf("%v: Expected %v dislodged in %v, but found nothing", statePair.Case, unit, prov)
		}
	}
	for prov, unit := range s.Units() {
		if _, ok := statePair.After.Units[prov]; !ok {
			err = true
			t.Errorf("%v: Expected %v to be empty, but found %v", statePair.Case, prov, unit)
		}
	}
	for prov, unit := range s.Dislodgeds() {
		if _, ok := statePair.After.Dislodgeds[prov]; !ok {
			err = true
			t.Errorf("%v: Expected %v to be empty of dislodged units, but found %v", statePair.Case, prov, unit)
		}
	}
	if err {
		dip.DumpLog()
		t.Errorf("%v: ### Units:", statePair.Case)
		for prov, unit := range statePair.Before.Units {
			t.Errorf("%v: %v %v", statePair.Case, prov, unit)
		}
		t.Errorf("%v: ### Dislodged before:", statePair.Case)
		for prov, disl := range statePair.Before.Dislodgeds {
			t.Errorf("%v: %v %v", statePair.Case, prov, disl)
		}
		t.Errorf("%v: ### Orders:", statePair.Case)
		for _, order := range statePair.Before.Orders {
			t.Errorf("%v: %v", statePair.Case, order.Order)
		}
		t.Errorf("%v: ### Units after:", statePair.Case)
		for prov, unit := range s.Units() {
			t.Errorf("%v: %v %v", statePair.Case, prov, unit)
		}
		t.Errorf("%v: ### Dislodged after:", statePair.Case)
		for prov, unit := range s.Dislodgeds() {
			t.Errorf("%v: %v %v", statePair.Case, prov, unit)
		}
		t.Errorf("%v: ### Errors:", statePair.Case)
		for prov, err := range s.Resolutions() {
			t.Errorf("%v: %v %v", statePair.Case, prov, err)
		}
		t.Fatalf("%v failed", statePair.Case)
	}
}

func assertDATC(t *testing.T, file string) {
	in, err := os.Open(file)
	if err != nil {
		t.Fatalf("%v", err)
	}
	parser := datc.Parser{
		Variant:        "Standard",
		OrderParser:    DATCOrder,
		PhaseParser:    DATCPhase,
		NationParser:   DATCNation,
		UnitTypeParser: DATCUnitType,
		ProvinceParser: DATCProvince,
	}
	if err := parser.Parse(in, func(statePair *datc.StatePair) {
		dip.ClearLog()
		dip.Logf("Running %v", statePair.Case)
		testDATC(t, statePair)
	}); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDATC(t *testing.T) {
	assertDATC(t, "datc/datc_v2.4_06.txt")
	assertDATC(t, "datc/diplicity_errors.txt")
	assertDATC(t, "datc/droidippy_errors.txt")
	assertDATC(t, "datc/dipai.txt")
	assertDATC(t, "datc/real.txt")
}
