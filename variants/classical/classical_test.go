package classical

import (
	"os"
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/datc"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/start"

	cla "github.com/zond/godip/variants/classical/common"
	ord "github.com/zond/godip/variants/classical/orders"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
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
	tst.AssertOrderValidity(t, judge, orders.SupportMove("bre", "par", "gas"), cla.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportHold("par", "bre"), cla.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("par", "bre", "gas"), cla.France, nil)
	judge.SetUnit("spa/sc", godip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("por", godip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("gol", godip.Unit{cla.Fleet, cla.France})
	tst.AssertOrderValidity(t, judge, orders.SupportMove("spa/sc", "por", "mid"), cla.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("gol", "mar", "spa"), cla.France, nil)
	// Missing unit
	tst.AssertOrderValidity(t, judge, orders.SupportMove("ruh", "kie", "hol"), "", godip.ErrMissingUnit)
	// Missing supportee
	tst.AssertOrderValidity(t, judge, orders.SupportHold("ber", "sil"), "", godip.ErrMissingSupportUnit)
	// Illegal support
	tst.AssertOrderValidity(t, judge, orders.SupportHold("bre", "par"), "", godip.ErrIllegalSupportPosition)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "por"), "", godip.ErrIllegalSupportDestination)
	judge.RemoveUnit("spa/sc")
	judge.SetUnit("spa/nc", godip.Unit{cla.Fleet, cla.France})
	tst.AssertOrderValidity(t, judge, orders.SupportMove("spa/nc", "mar", "gol"), "", godip.ErrIllegalSupportDestination)
	// Illegal moves
	tst.AssertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "bur"), "", godip.ErrIllegalSupportMove)
}

func TestConvoy(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("bal", godip.Unit{cla.Fleet, cla.Germany})
	tst.AssertOrderValidity(t, judge, orders.Move("ber", "lvn"), cla.Germany, nil)

	judge.SetUnit("tys", godip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", godip.Unit{cla.Fleet, cla.Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("rom", "spa"), cla.Italy, nil)
}

func TestConvoyValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("nth", godip.Unit{cla.Fleet, cla.France})
	judge.RemoveUnit("lon")
	judge.SetUnit("lon", godip.Unit{cla.Army, cla.England})
	tst.AssertOrderValidity(t, judge, orders.Convoy("nth", "lon", "nwy"), cla.France, nil)

	// Check that we can't convoy via Constantinople (nb. all edges are sea).
	judge.RemoveUnit("sev")
	judge.SetUnit("sev", godip.Unit{cla.Army, cla.Russia})
	judge.SetUnit("bla", godip.Unit{cla.Fleet, cla.Russia})
	judge.SetUnit("con", godip.Unit{cla.Fleet, cla.Russia})
	judge.SetUnit("aeg", godip.Unit{cla.Fleet, cla.Russia})
	tst.AssertOrderValidity(t, judge, orders.Convoy("bla", "sev", "gre"), "", godip.ErrIllegalConvoyMove)
}

func TestHoldValidation(t *testing.T) {
	judge := startState(t)
	tst.AssertOrderValidity(t, judge, orders.Hold("par"), cla.France, nil)
}

func TestBuildValidation(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("par")
	judge.SetUnit("spa", godip.Unit{cla.Army, cla.France})
	judge.Next()
	judge.Next()
	judge.Next()
	judge.Next()
	tst.AssertOrderValidity(t, judge, orders.Build("par", cla.Army, time.Now()), cla.France, nil)
}

func TestDisbandValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("pic", godip.Unit{cla.Army, cla.Germany})
	judge.SetUnit("bur", godip.Unit{cla.Army, cla.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	// Disband after dislodge
	tst.AssertOrderValidity(t, judge, orders.Disband("par", time.Now()), cla.France, nil)
	judge.Next()
	judge.SetUnit("bur", godip.Unit{cla.Army, cla.France})
	judge.Next()
	judge.Next()
	// Disband after SC deficit
	tst.AssertOrderValidity(t, judge, orders.Disband("bur", time.Now()), cla.France, nil)
}

func TestMoveValidation(t *testing.T) {
	judge := startState(t)
	// Happy path fleet
	tst.AssertOrderValidity(t, judge, orders.Move("bre", "mid"), cla.France, nil)
	// Happy path army
	tst.AssertOrderValidity(t, judge, orders.Move("mun", "ruh"), cla.Germany, nil)
	// Too far
	tst.AssertOrderValidity(t, judge, orders.Move("bre", "wes"), "", godip.ErrIllegalMove)
	// Fleet on land
	tst.AssertOrderValidity(t, judge, orders.Move("bre", "par"), "", godip.ErrIllegalDestination)
	// Army at sea
	tst.AssertOrderValidity(t, judge, orders.Move("smy", "eas"), "", godip.ErrIllegalDestination)
	// Unknown source
	tst.AssertOrderValidity(t, judge, orders.Move("a", "mid"), "", godip.ErrInvalidSource)
	// Unknown destination
	tst.AssertOrderValidity(t, judge, orders.Move("bre", "a"), "", godip.ErrInvalidDestination)
	// Missing sea path
	tst.AssertOrderValidity(t, judge, orders.Move("par", "mos"), "", godip.ErrMissingConvoyPath)
	// No unit
	tst.AssertOrderValidity(t, judge, orders.Move("spa", "por"), "", godip.ErrMissingUnit)
	// Working convoy
	judge.SetUnit("eng", godip.Unit{cla.Fleet, cla.England})
	judge.SetUnit("wal", godip.Unit{cla.Army, cla.England})
	tst.AssertOrderValidity(t, judge, orders.Move("wal", "bre"), cla.England, nil)
	// Missing convoy
	tst.AssertOrderValidity(t, judge, orders.Move("wal", "gas"), "", godip.ErrMissingConvoyPath)

	judge.SetUnit("pic", godip.Unit{cla.Army, cla.Germany})
	judge.SetUnit("bur", godip.Unit{cla.Army, cla.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	tst.AssertOrderValidity(t, judge, orders.Move("par", "gas"), cla.France, nil)

	judge.Next()
	judge.SetUnit("tys", godip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", godip.Unit{cla.Fleet, cla.Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("rom", "spa/sc"), cla.Italy, nil)
}

func TestMoveAdjudication(t *testing.T) {
	tst.AssertMove(t, startState(t), "bre", "mid", true)
	tst.AssertMove(t, startState(t), "stp/sc", "bot", true)
	tst.AssertMove(t, startState(t), "vie", "bud", false)
	tst.AssertMove(t, startState(t), "mid", "nat", false)
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
		godip.DumpLog()
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
		godip.ClearLog()
		godip.Logf("Running %v", statePair.Case)
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

func TestConvoyOpts(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("lon", orders.Move("lon", "nth"))
	judge.SetOrder("lvp", orders.Move("lvp", "yor"))
	judge.SetOrder("stp", orders.Move("stp", "bot"))
	judge.SetOrder("mos", orders.Move("mos", "stp"))
	judge.Next()
	judge.Next()
	opts := judge.Phase().Options(judge, cla.England)
	tst.AssertOpt(t, opts, []string{"yor", "Move", "yor", "nwy"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "ber", "kie"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "con", "smy"})
	tst.AssertOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "nwy"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "yor"})
	opts = judge.Phase().Options(judge, cla.Russia)
	tst.AssertOpt(t, opts, []string{"bot", "Convoy", "bot", "stp", "swe"})
	tst.AssertOpt(t, opts, []string{"stp", "Move", "stp", "swe"})

	judge.SetUnit("tys", godip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("gol", godip.Unit{cla.Fleet, cla.Italy})
	opts = judge.Phase().Options(judge, cla.Italy)
	tst.AssertOpt(t, opts, []string{"rom", "Move", "rom", "spa"})
	tst.AssertOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa"})
	tst.AssertNoOpt(t, opts, []string{"rom", "Move", "rom", "spa/sc"})
	tst.AssertNoOpt(t, opts, []string{"rom", "Move", "rom", "spa/nc"})
	tst.AssertNoOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa/sc"})
	tst.AssertNoOpt(t, opts, []string{"tys", "Convoy", "tys", "rom", "spa/nc"})
}

func TestSupportOpts(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("lon", orders.Move("lon", "eng"))
	judge.SetOrder("lvp", orders.Move("lvp", "wal"))
	judge.SetOrder("edi", orders.Move("edi", "nth"))
	judge.Next()
	judge.Next()
	opts := judge.Phase().Options(judge, cla.England)
	tst.AssertOpt(t, opts, []string{"eng", "Support", "eng", "wal", "lon"})
	tst.AssertOpt(t, opts, []string{"nth", "Support", "nth", "wal", "bel"})
	tst.AssertNoOpt(t, opts, []string{"eng", "Support", "eng", "wal", "bel"})
	opts = judge.Phase().Options(judge, cla.France)
	tst.AssertOpt(t, opts, []string{"par", "Support", "par", "bre", "pic"})
	tst.AssertNoOpt(t, opts, []string{"par", "Support", "par", "par", "pic"})
	tst.AssertNoOpt(t, opts, []string{"par", "Support", "par", "pic", "par"})
}

func TestSTPOptionsAtStart(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, cla.Russia)
	tst.AssertNoOpt(t, opts, []string{"stp/nc"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc"})
	tst.AssertOpt(t, opts, []string{"stp", "Move", "stp/sc", "lvn"})
	tst.AssertOpt(t, opts, []string{"stp", "Move", "stp/sc", "bot"})
	tst.AssertOpt(t, opts, []string{"stp", "Move", "stp/sc", "fin"})
	tst.AssertNoOpt(t, opts, []string{"stp", "Convoy"})
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
	tst.AssertNoOpt(t, opts, []string{"stp"})
	tst.AssertOpt(t, opts, []string{"stp/nc", "Build", "Fleet", "stp/nc"})
	tst.AssertOpt(t, opts, []string{"stp/sc", "Build", "Fleet", "stp/sc"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Fleet", "stp"})
	tst.AssertOpt(t, opts, []string{"stp/nc", "Build", "Army", "stp"})
	tst.AssertOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/nc"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/sc"})
}

func TestSupportSTPOpts(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, cla.Russia)
	// Check that initially Moscow can support St Petersburg South Coast to Livonia
	tst.AssertOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})
	// Check that the south coast is not mentioned in the suggestion list.
	tst.AssertNoOpt(t, opts, []string{"mos", "Support", "mos", "stp/sc", "lvn"})

	// Swap St Petersburg to North Coast and check there's no support option to Livonia
	judge.RemoveUnit("stp/sc")
	judge.SetUnit("stp/nc", godip.Unit{cla.Fleet, cla.Russia})
	opts = judge.Phase().Options(judge, cla.Russia)
	tst.AssertNoOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})

	// Swap St Petersburg to contain an army instead and check the support option is back.
	judge.RemoveUnit("stp/nc")
	judge.SetUnit("stp", godip.Unit{cla.Army, cla.Russia})
	opts = judge.Phase().Options(judge, cla.Russia)
	tst.AssertOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})
}

func TestBULOptions(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, cla.Turkey)
	tst.AssertNoOpt(t, opts, []string{"con", "Move", "con", "bul/sc"})
	tst.AssertNoOpt(t, opts, []string{"con", "Move", "con", "bul/ec"})
	tst.AssertOpt(t, opts, []string{"con", "Move", "con", "bul"})
	judge.SetOrder("con", orders.Move("con", "bul"))
	judge.Next()
	judge.Next()
	opts = judge.Phase().Options(judge, cla.Turkey)
	tst.AssertNoOpt(t, opts, []string{"bul/sc"})
	tst.AssertNoOpt(t, opts, []string{"bul/ec"})
	tst.AssertNoOpt(t, opts, []string{"bul", "Move", "bul/sc"})
	tst.AssertNoOpt(t, opts, []string{"bul", "Move", "bul/ec"})
	tst.AssertOpt(t, opts, []string{"bul", "Move", "bul", "rum"})
	tst.AssertOpt(t, opts, []string{"bul", "Move", "bul", "gre"})
}

// Test that por M spa supported by mid works in
// https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAgOr0mgoM/Phase/12/Map
func TestMIDPORSPASupportOptions(t *testing.T) {
	judge := state.New(start.Graph(), &phase{1903, cla.Fall, cla.Movement, ord.ClassicalParser}, BackupRule)
	if err := judge.SetUnits(start.Units()); err != nil {
		t.Fatal(err)
	}
	scs := start.SupplyCenters()
	scs["por"] = cla.France
	scs["spa"] = cla.France
	scs["mar"] = cla.France
	judge.SetSupplyCenters(scs)
	judge.SetUnit("mid", godip.Unit{cla.Fleet, cla.France})
	judge.SetUnit("por", godip.Unit{cla.Army, cla.France})
	judge.RemoveUnit("mar")
	judge.SetUnit("mar", godip.Unit{cla.Army, cla.Italy})
	judge.SetUnit("gol", godip.Unit{cla.Fleet, cla.Italy})
	judge.SetUnit("wes", godip.Unit{cla.Fleet, cla.Italy})
	opts := judge.Phase().Options(judge, cla.France)
	tst.AssertOpt(t, opts, []string{"por", "Move", "por", "spa"})
	tst.AssertOpt(t, opts, []string{"mid", "Support", "mid", "por", "spa"})
}
