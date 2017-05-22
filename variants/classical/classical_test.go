package classical

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/zond/godip/datc"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/orders"

	cla "github.com/zond/godip/variants/classical/common"
	dip "github.com/zond/godip/common"
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
	assertOrderValidity(t, judge, orders.SupportMove("bre", "par", "gas"), cla.France, nil)
	assertOrderValidity(t, judge, orders.SupportHold("par", "bre"), cla.France, nil)
	assertOrderValidity(t, judge, orders.SupportMove("par", "bre", "gas"), cla.France, nil)
	judge.SetUnit("spa/sc", dip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("por", dip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("gol", dip.Unit{cla.Fleet, cla.France})
	assertOrderValidity(t, judge, orders.SupportMove("spa/sc", "por", "mid"), cla.France, nil)
	assertOrderValidity(t, judge, orders.SupportMove("gol", "mar", "spa"), cla.France, nil)
	// Missing unit
	assertOrderValidity(t, judge, orders.SupportMove("ruh", "kie", "hol"), "", cla.ErrMissingUnit)
	// Missing supportee
	assertOrderValidity(t, judge, orders.SupportHold("ber", "sil"), "", cla.ErrMissingSupportUnit)
	// Illegal support
	assertOrderValidity(t, judge, orders.SupportHold("bre", "par"), "", cla.ErrIllegalSupportPosition)
	assertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "por"), "", cla.ErrIllegalSupportDestination)
	judge.RemoveUnit("spa/sc")
	judge.SetUnit("spa/nc", dip.Unit{cla.Fleet, cla.France})
	assertOrderValidity(t, judge, orders.SupportMove("spa/nc", "mar", "gol"), "", cla.ErrIllegalSupportDestination)
	// Illegal moves
	assertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "bur"), "", cla.ErrIllegalSupportMove)
}

func TestConvoy(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("bal", dip.Unit{cla.Fleet, cla.Germany})
	assertOrderValidity(t, judge, orders.Move("ber", "lvn"), cla.Germany, nil)

	judge.SetUnit("tys", dip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", dip.Unit{cla.Fleet, cla.Italy})
	assertOrderValidity(t, judge, orders.Move("rom", "spa"), cla.Italy, nil)
}

func TestConvoyValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("nth", dip.Unit{cla.Fleet, cla.France})
	judge.RemoveUnit("lon")
	judge.SetUnit("lon", dip.Unit{cla.Army, cla.England})
	assertOrderValidity(t, judge, orders.Convoy("nth", "lon", "nwy"), cla.France, nil)
}

func TestHoldValidation(t *testing.T) {
	judge := startState(t)
	assertOrderValidity(t, judge, orders.Hold("par"), cla.France, nil)
}

func TestBuildValidation(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("par")
	judge.SetUnit("spa", dip.Unit{cla.Army, cla.France})
	judge.Next()
	judge.Next()
	judge.Next()
	judge.Next()
	assertOrderValidity(t, judge, orders.Build("par", cla.Army, time.Now()), cla.France, nil)
}

func TestDisbandValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("pic", dip.Unit{cla.Army, cla.Germany})
	judge.SetUnit("bur", dip.Unit{cla.Army, cla.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	// Disband after dislodge
	assertOrderValidity(t, judge, orders.Disband("par", time.Now()), cla.France, nil)
	judge.Next()
	judge.SetUnit("bur", dip.Unit{cla.Army, cla.France})
	judge.Next()
	judge.Next()
	// Disband after SC deficit
	assertOrderValidity(t, judge, orders.Disband("bur", time.Now()), cla.France, nil)
}

func TestMoveValidation(t *testing.T) {
	judge := startState(t)
	// Happy path fleet
	assertOrderValidity(t, judge, orders.Move("bre", "mid"), cla.France, nil)
	// Happy path army
	assertOrderValidity(t, judge, orders.Move("mun", "ruh"), cla.Germany, nil)
	// Too far
	assertOrderValidity(t, judge, orders.Move("bre", "wes"), "", cla.ErrIllegalMove)
	// Fleet on land
	assertOrderValidity(t, judge, orders.Move("bre", "par"), "", cla.ErrIllegalDestination)
	// Army at sea
	assertOrderValidity(t, judge, orders.Move("smy", "eas"), "", cla.ErrIllegalDestination)
	// Unknown source
	assertOrderValidity(t, judge, orders.Move("a", "mid"), "", cla.ErrInvalidSource)
	// Unknown destination
	assertOrderValidity(t, judge, orders.Move("bre", "a"), "", cla.ErrInvalidDestination)
	// Missing sea path
	assertOrderValidity(t, judge, orders.Move("par", "mos"), "", cla.ErrMissingConvoyPath)
	// No unit
	assertOrderValidity(t, judge, orders.Move("spa", "por"), "", cla.ErrMissingUnit)
	// Working convoy
	judge.SetUnit("eng", dip.Unit{cla.Fleet, cla.England})
	judge.SetUnit("wal", dip.Unit{cla.Army, cla.England})
	assertOrderValidity(t, judge, orders.Move("wal", "bre"), cla.England, nil)
	// Missing convoy
	assertOrderValidity(t, judge, orders.Move("wal", "gas"), "", cla.ErrMissingConvoyPath)

	judge.SetUnit("pic", dip.Unit{cla.Army, cla.Germany})
	judge.SetUnit("bur", dip.Unit{cla.Army, cla.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	assertOrderValidity(t, judge, orders.Move("par", "gas"), cla.France, nil)

	judge.Next()
	judge.SetUnit("tys", dip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", dip.Unit{cla.Fleet, cla.Italy})
	assertOrderValidity(t, judge, orders.Move("rom", "spa/sc"), cla.Italy, nil)
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

func hasOptHelper(opts map[string]interface{}, order []string, originalOpts map[string]interface{}, originalOrder []string) error {
	if len(order) == 0 {
		return nil
	}
	if _, found := opts[order[0]]; !found {
		b, err := json.MarshalIndent(originalOpts, "  ", "  ")
		if err != nil {
			return err
		}
		b2, err := json.MarshalIndent(opts, "  ", "  ")
		if err != nil {
			return err
		}
		return fmt.Errorf("Got no option for %+v in %s, failed at %+v in %s, wanted it!", originalOrder, b, order, b2)
	}
	return hasOptHelper(opts[order[0]].(map[string]interface{})["Next"].(map[string]interface{}), order[1:], originalOpts, originalOrder)
}

func hasOpt(opts dip.Options, order []string) error {
	b, err := json.MarshalIndent(opts, "  ", "  ")
	if err != nil {
		return err
	}
	converted := map[string]interface{}{}
	if err := json.Unmarshal(b, &converted); err != nil {
		return err
	}
	return hasOptHelper(converted, order, converted, order)
}

func assertOpt(t *testing.T, opts dip.Options, order []string) {
	err := hasOpt(opts, order)
	if err != nil {
		t.Error(err)
	}
}

func assertNoOpt(t *testing.T, opts dip.Options, order []string) {
	err := hasOpt(opts, order)
	if err == nil {
		b, err := json.MarshalIndent(opts, "  ", "  ")
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("Found option for %+v in %s, didn't want it", order, b)
	}
}

func TestConvoyOpts(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("lon", orders.Move("lon", "nth"))
	judge.SetOrder("lvp", orders.Move("lvp", "yor"))
	judge.SetOrder("stp", orders.Move("stp", "bot"))
	judge.SetOrder("mos", orders.Move("mos", "stp"))
	judge.Next()
	judge.Next()
	opts := judge.Phase().Options(judge, cla.England)
	assertOpt(t, opts, []string{"yor", "Move", "yor", "nwy"})
	assertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "ber", "kie"})
	assertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "con", "smy"})
	assertOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "nwy"})
	assertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "yor"})
	opts = judge.Phase().Options(judge, cla.Russia)
	assertOpt(t, opts, []string{"bot", "Convoy", "bot", "stp", "swe"})
	assertOpt(t, opts, []string{"stp", "Move", "stp", "swe"})

	judge.SetUnit("tys", dip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", dip.Unit{cla.Fleet, cla.Italy})
	opts = judge.Phase().Options(judge, cla.Italy)
	assertOpt(t, opts, []string{"rom", "Move", "rom", "spa"})
	assertOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa"})
	assertNoOpt(t, opts, []string{"rom", "Move", "rom", "spa/sc"})
	assertNoOpt(t, opts, []string{"rom", "Move", "rom", "spa/nc"})
	assertNoOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa/sc"})
	assertNoOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa/nc"})
}

func TestSupportOpts(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("lon", orders.Move("lon", "eng"))
	judge.SetOrder("lvp", orders.Move("lvp", "wal"))
	judge.SetOrder("edi", orders.Move("edi", "nth"))
	judge.Next()
	judge.Next()
	opts := judge.Phase().Options(judge, cla.England)
	assertOpt(t, opts, []string{"eng", "Support", "eng", "wal", "lon"})
	assertOpt(t, opts, []string{"nth", "Support", "nth", "wal", "bel"})
	assertNoOpt(t, opts, []string{"eng", "Support", "eng", "wal", "bel"})
	opts = judge.Phase().Options(judge, cla.France)
	assertOpt(t, opts, []string{"par", "Support", "par", "bre", "pic"})
	assertNoOpt(t, opts, []string{"par", "Support", "par", "par", "pic"})
	assertNoOpt(t, opts, []string{"par", "Support", "par", "pic", "par"})
}

func TestSTPOptionsAtStart(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, cla.Russia)
	assertNoOpt(t, opts, []string{"stp/nc"})
	assertNoOpt(t, opts, []string{"stp/sc"})
	assertOpt(t, opts, []string{"stp", "Move", "stp/sc", "lvn"})
	assertOpt(t, opts, []string{"stp", "Move", "stp/sc", "bot"})
	assertOpt(t, opts, []string{"stp", "Move", "stp/sc", "fin"})
	assertNoOpt(t, opts, []string{"stp", "Convoy"})
}

func TestSTPBuildOptions(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("stp", orders.Move("stp/sc", "fin"))
	judge.Next()
	judge.Next()
	judge.SetOrder("fin", orders.Move("fin", "swe"))
	judge.Next()
	judge.Next()
	opts := judge.Phase().Options(judge, cla.Russia)
	assertNoOpt(t, opts, []string{"stp"})
	assertOpt(t, opts, []string{"stp/nc", "Build", "Fleet", "stp/nc"})
	assertOpt(t, opts, []string{"stp/sc", "Build", "Fleet", "stp/sc"})
	assertNoOpt(t, opts, []string{"stp/sc", "Build", "Fleet", "stp"})
	assertOpt(t, opts, []string{"stp/nc", "Build", "Army", "stp"})
	assertOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp"})
	assertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/nc"})
	assertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/sc"})
}

func TestBULOptions(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, cla.Turkey)
	assertNoOpt(t, opts, []string{"con", "Move", "con", "bul/sc"})
	assertNoOpt(t, opts, []string{"con", "Move", "con", "bul/ec"})
	assertOpt(t, opts, []string{"con", "Move", "con", "bul"})
	judge.SetOrder("con", orders.Move("con", "bul"))
	judge.Next()
	judge.Next()
	opts = judge.Phase().Options(judge, cla.Turkey)
	assertNoOpt(t, opts, []string{"bul/sc"})
	assertNoOpt(t, opts, []string{"bul/ec"})
	assertNoOpt(t, opts, []string{"bul", "Move", "bul/sc"})
	assertNoOpt(t, opts, []string{"bul", "Move", "bul/ec"})
	assertOpt(t, opts, []string{"bul", "Move", "bul", "rum"})
	assertOpt(t, opts, []string{"bul", "Move", "bul", "gre"})
}
