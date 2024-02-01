package classical

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/datc"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical/start"

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
	tst.AssertOrderValidity(t, judge, orders.SupportMove("bre", "par", "gas"), godip.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportHold("par", "bre"), godip.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("par", "bre", "gas"), godip.France, nil)
	judge.SetUnit("spa/sc", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("por", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("gol", godip.Unit{godip.Fleet, godip.France})
	tst.AssertOrderValidity(t, judge, orders.SupportMove("spa/sc", "por", "mid"), godip.France, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("gol", "mar", "spa"), godip.France, nil)
	// Missing unit
	tst.AssertOrderValidity(t, judge, orders.SupportMove("ruh", "kie", "hol"), "", godip.ErrMissingUnit)
	// Missing supportee
	tst.AssertOrderValidity(t, judge, orders.SupportHold("ber", "sil"), "", godip.ErrMissingSupportUnit)
	// Illegal support
	tst.AssertOrderValidity(t, judge, orders.SupportHold("bre", "par"), "", godip.ErrIllegalSupportPosition)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "por"), "", godip.ErrIllegalSupportDestination)
	judge.RemoveUnit("spa/sc")
	judge.SetUnit("spa/nc", godip.Unit{godip.Fleet, godip.France})
	tst.AssertOrderValidity(t, judge, orders.SupportMove("spa/nc", "mar", "gol"), "", godip.ErrIllegalSupportDestination)
	// Illegal moves
	tst.AssertOrderValidity(t, judge, orders.SupportMove("mar", "spa/nc", "bur"), "", godip.ErrIllegalSupportMove)
}

func TestConvoy(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("bal", godip.Unit{godip.Fleet, godip.Germany})
	tst.AssertOrderValidity(t, judge, orders.Move("ber", "lvn"), godip.Germany, nil)

	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.Italy})
	judge.SetUnit("gol", godip.Unit{godip.Fleet, godip.Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("rom", "spa"), godip.Italy, nil)
}

func TestConvoyValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("nth", godip.Unit{godip.Fleet, godip.France})
	judge.RemoveUnit("lon")
	judge.SetUnit("lon", godip.Unit{godip.Army, godip.England})
	tst.AssertOrderValidity(t, judge, orders.Convoy("nth", "lon", "nwy"), godip.France, nil)

	// Check that we can't convoy via Constantinople (nb. all edges are sea).
	judge.RemoveUnit("sev")
	judge.SetUnit("sev", godip.Unit{godip.Army, godip.Russia})
	judge.SetUnit("bla", godip.Unit{godip.Fleet, godip.Russia})
	judge.SetUnit("con", godip.Unit{godip.Fleet, godip.Russia})
	judge.SetUnit("aeg", godip.Unit{godip.Fleet, godip.Russia})
	tst.AssertOrderValidity(t, judge, orders.Convoy("bla", "sev", "gre"), "", godip.ErrIllegalConvoyMove)
}

func TestHoldValidation(t *testing.T) {
	judge := startState(t)
	tst.AssertOrderValidity(t, judge, orders.Hold("par"), godip.France, nil)
}

func TestBuildValidation(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("par")
	judge.SetUnit("spa", godip.Unit{godip.Army, godip.France})
	judge.Next()
	judge.Next()
	judge.Next()
	judge.Next()
	tst.AssertOrderValidity(t, judge, orders.Build("par", godip.Army, time.Now()), godip.France, nil)
}

func TestDisbandValidation(t *testing.T) {
	judge := startState(t)
	judge.SetUnit("pic", godip.Unit{godip.Army, godip.Germany})
	judge.SetUnit("bur", godip.Unit{godip.Army, godip.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	// Disband after dislodge
	tst.AssertOrderValidity(t, judge, orders.Disband("par", time.Now()), godip.France, nil)
	judge.Next()
	judge.SetUnit("bur", godip.Unit{godip.Army, godip.France})
	judge.Next()
	judge.Next()
	// Disband after SC deficit
	tst.AssertOrderValidity(t, judge, orders.Disband("bur", time.Now()), godip.France, nil)
}

func TestMoveValidation(t *testing.T) {
	judge := startState(t)
	// Happy path fleet
	tst.AssertOrderValidity(t, judge, orders.Move("bre", "mid"), godip.France, nil)
	// Happy path army
	tst.AssertOrderValidity(t, judge, orders.Move("mun", "ruh"), godip.Germany, nil)
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
	judge.SetUnit("eng", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("wal", godip.Unit{godip.Army, godip.England})
	tst.AssertOrderValidity(t, judge, orders.Move("wal", "bre"), godip.England, nil)
	// Missing convoy
	tst.AssertOrderValidity(t, judge, orders.Move("wal", "gas"), "", godip.ErrMissingConvoyPath)

	judge.SetUnit("pic", godip.Unit{godip.Army, godip.Germany})
	judge.SetUnit("bur", godip.Unit{godip.Army, godip.Germany})
	judge.SetOrder("bur", orders.Move("bur", "par"))
	judge.SetOrder("pic", orders.SupportMove("pic", "bur", "par"))
	judge.Next()
	tst.AssertOrderValidity(t, judge, orders.Move("par", "gas"), godip.France, nil)

	judge.Next()
	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.Italy})
	judge.SetUnit("gol", godip.Unit{godip.Fleet, godip.Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("rom", "spa/sc"), godip.Italy, nil)
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
		s = Blank(NewPhase(
			1901,
			godip.Spring,
			godip.Movement,
		))
	} else {
		s = Blank(statePair.Before.Phase)
	}
	s.SetUnits(statePair.Before.Units)
	s.SetDislodgeds(statePair.Before.Dislodgeds)
	s.SetSupplyCenters(statePair.Before.SCs)
	for prov, order := range statePair.Before.Orders {
		if s.Phase().Type() == godip.Movement {
			if u, _, ok := s.Unit(prov); ok && u.Nation == order.Nation {
				s.SetOrder(prov, order.Order)
			}
		} else if s.Phase().Type() == godip.Retreat {
			if u, _, ok := s.Dislodged(prov); ok && u.Nation == order.Nation {
				s.SetOrder(prov, order.Order)
			}
		} else if s.Phase().Type() == godip.Adjustment {
			if order.Order.Type() == godip.Build {
				if n, _, ok := s.SupplyCenter(prov); ok && n == order.Nation {
					s.SetOrder(prov, order.Order)
				}
			} else if order.Order.Type() == godip.Disband {
				if u, _, ok := s.Unit(prov); ok && u.Nation == order.Nation {
					s.SetOrder(prov, order.Order)
				}
			}
		} else {
			t.Fatalf("Unsupported phase type %v", s.Phase().Type())
		}
	}
	for _, order := range statePair.Before.FailedOrders {
		if order.Order.Type() == godip.Move && !order.Order.Flags()[godip.ViaConvoy] {
			s.AddBounce(order.Order.Targets()[0], order.Order.Targets()[1])
		}
	}
	for _, order := range statePair.Before.SuccessfulOrders {
		if order.Order.Type() == godip.Move && !order.Order.Flags()[godip.ViaConvoy] {
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
	opts := judge.Phase().Options(judge, godip.England)
	tst.AssertOpt(t, opts, []string{"yor", "Move", "yor", "nwy"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "ber", "kie"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "con", "smy"})
	tst.AssertOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "nwy"})
	tst.AssertNoOpt(t, opts, []string{"nth", "Convoy", "nth", "yor", "yor"})
	opts = judge.Phase().Options(judge, godip.Russia)
	tst.AssertOpt(t, opts, []string{"bot", "Convoy", "bot", "stp", "swe"})
	tst.AssertOpt(t, opts, []string{"stp", "Move", "stp", "swe"})

	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.Italy})
	judge.SetUnit("gol", godip.Unit{godip.Fleet, godip.Italy})
	opts = judge.Phase().Options(judge, godip.Italy)
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
	opts := judge.Phase().Options(judge, godip.England)
	tst.AssertOpt(t, opts, []string{"eng", "Support", "eng", "wal", "lon"})
	tst.AssertOpt(t, opts, []string{"nth", "Support", "nth", "wal", "bel"})
	tst.AssertNoOpt(t, opts, []string{"eng", "Support", "eng", "wal", "bel"})
	opts = judge.Phase().Options(judge, godip.France)
	tst.AssertOpt(t, opts, []string{"par", "Support", "par", "bre", "pic"})
	tst.AssertNoOpt(t, opts, []string{"par", "Support", "par", "par", "pic"})
	tst.AssertNoOpt(t, opts, []string{"par", "Support", "par", "pic", "par"})
}

func TestRetreatOpts(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("nap", orders.Move("nap", "ion"))
	judge.SetOrder("mun", orders.Move("mun", "tyr"))
	judge.SetOrder("tri", orders.Move("tri", "adr"))
	judge.SetOrder("bud", orders.Move("bud", "ser"))
	judge.SetOrder("vie", orders.Move("vie", "tri"))
	judge.Next()
	judge.Next()
	// Venice invades Trieste supported by Tyrolia.
	judge.SetOrder("ven", orders.Move("ven", "tri"))
	judge.SetOrder("tyr", orders.SupportMove("tyr", "ven", "tri"))
	// Serbia and Ionian Sea bounce in Albania.
	judge.SetOrder("ser", orders.Move("ser", "alb"))
	judge.SetOrder("ion", orders.Move("ion", "alb"))
	judge.Next()

	// Check the options for the displaced unit in Trieste.
	opts := judge.Phase().Options(judge, godip.Austria)
	tst.AssertOpt(t, opts, []string{"tri", "Move", "tri", "vie"})
	tst.AssertOpt(t, opts, []string{"tri", "Move", "tri", "bud"})
	// Can't retreat to where the attack came from.
	tst.AssertNoOpt(t, opts, []string{"tri", "Move", "tri", "ven"})
	// Can't retreat to an occupied region.
	tst.AssertNoOpt(t, opts, []string{"tri", "Move", "tri", "tyr"})
	tst.AssertNoOpt(t, opts, []string{"tri", "Move", "tri", "ser"})
	// There was a bounce in Albania.
	tst.AssertNoOpt(t, opts, []string{"tri", "Move", "tri", "alb"})
	// Can't retreat via convoy.
	tst.AssertNoOpt(t, opts, []string{"tri", "Move", "tri", "apu"})
}

func TestSTPOptionsAtStart(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, godip.Russia)
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
	opts := judge.Phase().Options(judge, godip.Russia)
	tst.AssertNoOpt(t, opts, []string{"stp"})
	filter := "MAX:Build:0"
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/nc", "Build", "Fleet", "stp/nc"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/sc", "Build", "Fleet", "stp/sc"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Fleet", "stp"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/nc", "Build", "Army", "stp"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/sc", "Build", "Army", "stp"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/nc"})
	tst.AssertNoOpt(t, opts, []string{"stp/sc", "Build", "Army", "stp/sc"})
}

func assertCorroborateErrors(t *testing.T, inconsistencies []godip.Inconsistency, truth map[godip.Province]string) {
	homelessErrors := []string{}
	for _, inc := range inconsistencies {
		foundStrings := []string{}
		for _, err := range inc.Errors {
			foundStrings = append(foundStrings, err.Error())
		}
		foundString := strings.Join(foundStrings, ",")
		if inc.Province == "" {
			homelessErrors = append(homelessErrors, foundString)
		} else {
			if trueString, found := truth[inc.Province]; found {
				delete(truth, inc.Province)
				if foundString != trueString {
					t.Errorf("Got %q for %q, wanted %q", foundString, inc.Province, trueString)
				}
			} else {
				t.Errorf("Got %q for %q, wanted nothing", foundString, inc.Province)
			}
		}
	}
	homelessErrString := strings.Join(homelessErrors, ",")
	if homelessErrString != truth[""] {
		t.Errorf("Got homeless errors %v, wanted %v", homelessErrString, truth[""])
	}
	delete(truth, "")
	if len(truth) > 0 {
		t.Errorf("Missing some corroborate errors: %+v", truth)
	}
}

func TestSupportCorroborate(t *testing.T) {
	judge := Blank(NewPhase(1903, godip.Fall, godip.Movement))
	judge.SetUnit("con", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("ser", godip.Unit{godip.Army, godip.England})
	judge.SetOrder("con", orders.Move("con", "bul/sc"))
	judge.SetOrder("ser", orders.SupportMove("ser", "con", "bul"))
	if incons := judge.Corroborate(godip.England); len(incons) != 0 {
		t.Errorf("Got %+v, wanted []", incons)
	}
}

func TestConvoyCorroborate(t *testing.T) {
	judge := Blank(NewPhase(1903, godip.Fall, godip.Movement))
	judge.SetUnit("edi", godip.Unit{godip.Army, godip.England})
	judge.SetUnit("nrg", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("nth", godip.Unit{godip.Fleet, godip.England})
	judge.SetOrder("edi", orders.Move("edi", "nwy"))
	judge.SetOrder("nth", orders.Convoy("nth", "edi", "nwy"))
	judge.SetOrder("nrg", orders.Move("nrg", "nat"))
	if incons := judge.Corroborate(godip.England); len(incons) != 0 {
		t.Errorf("Got %+v, wanted []", incons)
	}
}

func TestCorroborate(t *testing.T) {
	judge := startState(t)
	// Should be mismatched support due to wrong dest.
	judge.SetOrder("mos", orders.Move("mos", "lvn"))
	judge.SetOrder("war", orders.SupportMove("war", "mos", "ukr"))
	// Should be mismatched support due to non moving supportee.
	judge.SetOrder("bud", orders.SupportMove("bud", "vie", "gal"))
	// Should be mismatched support due to moving supportee.
	judge.SetOrder("ven", orders.SupportHold("ven", "rom"))
	judge.SetOrder("rom", orders.Move("rom", "tus"))
	// Prepare for next phase testing.
	judge.SetOrder("bre", orders.Move("bre", "mid"))
	judge.SetOrder("par", orders.Move("par", "bre"))
	judge.SetOrder("lvp", orders.Move("lvp", "wal"))
	judge.SetOrder("lon", orders.Move("lon", "eng"))
	judge.SetOrder("nap", orders.Move("nap", "tys"))
	assertCorroborateErrors(t, judge.Corroborate(godip.Russia), map[godip.Province]string{
		"sev": "InconsistencyMissingOrder",
		"stp": "InconsistencyMissingOrder",
		"war": "InconsistencyMismatchedSupporter:mos",
	})
	assertCorroborateErrors(t, judge.Corroborate(godip.Italy), map[godip.Province]string{
		"ven": "InconsistencyMismatchedSupporter:rom",
	})
	assertCorroborateErrors(t, judge.Corroborate(godip.Austria), map[godip.Province]string{
		"bud": "InconsistencyMismatchedSupporter:vie",
		"vie": "InconsistencyMissingOrder",
		"tri": "InconsistencyMissingOrder",
	})
	judge.Next()
	judge.Next()
	// Should not be mismatched because different nations.
	judge.SetOrder("ven", orders.SupportMove("ven", "tri", "adr"))
	// Should be mismatched convoy due to wrong dest.
	judge.SetOrder("mid", orders.Convoy("mid", "bre", "por"))
	judge.SetOrder("bre", orders.Move("bre", "spa"))
	// Should be mismatched convoy due to non moving convoyee.
	judge.SetOrder("eng", orders.Convoy("eng", "wal", "pic"))
	// Should be mismatched convoy due to missing convoy order.
	judge.SetOrder("tus", orders.Move("tus", "tun"))
	// Prepare for next phase.
	judge.SetOrder("mar", orders.Move("mar", "spa"))
	assertCorroborateErrors(t, judge.Corroborate(godip.Italy), map[godip.Province]string{
		"tus": "InconsistencyMismatchedConvoyee:tys",
		"tys": "InconsistencyMissingOrder",
	})
	assertCorroborateErrors(t, judge.Corroborate(godip.France), map[godip.Province]string{
		"mid": "InconsistencyMismatchedConvoyer:bre",
		"bre": "InconsistencyMismatchedConvoyee:mid",
	})
	assertCorroborateErrors(t, judge.Corroborate(godip.England), map[godip.Province]string{
		"eng": "InconsistencyMismatchedConvoyer:wal",
		"wal": "InconsistencyMissingOrder",
		"edi": "InconsistencyMissingOrder",
	})
	judge.Next()
	judge.Next()
	assertCorroborateErrors(t, judge.Corroborate(godip.France), map[godip.Province]string{
		"": "InconsistencyOrderTypeCount:Build:Found:0:Want:1",
	})
	judge.Next()
	// Should not be mismatched because different nations.
	judge.SetOrder("eng", orders.Convoy("eng", "bre", "bel"))
	// Should not be mismatched becaues different nations.
	judge.SetOrder("bre", orders.Move("bre", "wal"))
	// Prepare for next.
	judge.SetOrder("vie", orders.Move("vie", "boh"))
	judge.SetOrder("ven", orders.Move("ven", "tyr"))
	assertCorroborateErrors(t, judge.Corroborate(godip.England), map[godip.Province]string{
		"wal": "InconsistencyMissingOrder",
		"edi": "InconsistencyMissingOrder",
	})
	assertCorroborateErrors(t, judge.Corroborate(godip.France), map[godip.Province]string{
		"mid": "InconsistencyMissingOrder",
		"spa": "InconsistencyMissingOrder",
	})
	judge.Next()
	judge.Next()
	// Eject mun.
	judge.SetOrder("boh", orders.Move("boh", "mun"))
	judge.SetOrder("tyr", orders.SupportMove("tyr", "boh", "mun"))
	judge.Next()
	judge.SetOrder("mun", orders.Move("mun", "ruh"))
	judge.Next()
	assertCorroborateErrors(t, judge.Corroborate(godip.Germany), map[godip.Province]string{
		"": "InconsistencyOrderTypeCount:Disband:Found:0:Want:1",
	})
}

func TestFilteredOptions(t *testing.T) {
	judge := startState(t)
	judge.SetOrder("stp", orders.Move("stp/sc", "fin"))
	judge.SetOrder("sev", orders.Move("sev", "rum"))
	judge.SetOrder("war", orders.Move("war", "sil"))
	judge.SetOrder("vie", orders.Move("vie", "boh"))
	judge.Next()
	judge.Next()
	judge.SetOrder("fin", orders.Move("fin", "swe"))
	judge.SetOrder("boh", orders.Move("boh", "mun"))
	judge.SetOrder("sil", orders.SupportMove("sil", "boh", "mun"))
	judge.Next()
	judge.SetOrder("mun", orders.Move("mun", "ruh"))
	judge.Next()
	opts := judge.Phase().Options(judge, godip.Russia)
	filter := "MAX:Build:1"
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/nc", "Build", "Fleet", "stp/nc"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/sc", "Build", "Fleet", "stp/sc"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/nc", "Build", "Army", "stp"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"stp/sc", "Build", "Army", "stp"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"sev", "Build", "Fleet", "sev"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"sev", "Build", "Army", "sev"})
	opts = judge.Phase().Options(judge, godip.Germany)
	filter = "MAX:Disband:0"
	tst.AssertFilteredOpt(t, opts, filter, []string{"kie", "Disband", "kie"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"ber", "Disband", "ber"})
	tst.AssertFilteredOpt(t, opts, filter, []string{"ruh", "Disband", "ruh"})
}

func TestSupportSTPOpts(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, godip.Russia)
	// Check that initially Moscow can support St Petersburg South Coast to Livonia
	tst.AssertOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})
	// Check that the south coast is not mentioned in the suggestion list.
	tst.AssertNoOpt(t, opts, []string{"mos", "Support", "mos", "stp/sc", "lvn"})

	// Swap St Petersburg to North Coast and check there's no support option to Livonia
	judge.RemoveUnit("stp/sc")
	judge.SetUnit("stp/nc", godip.Unit{godip.Fleet, godip.Russia})
	opts = judge.Phase().Options(judge, godip.Russia)
	tst.AssertNoOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})

	// Swap St Petersburg to contain an army instead and check the support option is back.
	judge.RemoveUnit("stp/nc")
	judge.SetUnit("stp", godip.Unit{godip.Army, godip.Russia})
	opts = judge.Phase().Options(judge, godip.Russia)
	tst.AssertOpt(t, opts, []string{"mos", "Support", "mos", "stp", "lvn"})
}

func TestBULOptions(t *testing.T) {
	judge := startState(t)
	opts := judge.Phase().Options(judge, godip.Turkey)
	tst.AssertNoOpt(t, opts, []string{"con", "Move", "con", "bul/sc"})
	tst.AssertNoOpt(t, opts, []string{"con", "Move", "con", "bul/ec"})
	tst.AssertOpt(t, opts, []string{"con", "Move", "con", "bul"})
	judge.SetOrder("con", orders.Move("con", "bul"))
	judge.Next()
	judge.Next()
	opts = judge.Phase().Options(judge, godip.Turkey)
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
	judge := Blank(NewPhase(1903, godip.Fall, godip.Movement))
	if err := judge.SetUnits(start.Units()); err != nil {
		t.Fatal(err)
	}
	scs := start.SupplyCenters()
	scs["por"] = godip.France
	scs["spa"] = godip.France
	scs["mar"] = godip.France
	judge.SetSupplyCenters(scs)
	judge.SetUnit("mid", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("por", godip.Unit{godip.Army, godip.France})
	judge.RemoveUnit("mar")
	judge.SetUnit("mar", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("gol", godip.Unit{godip.Fleet, godip.Italy})
	judge.SetUnit("wes", godip.Unit{godip.Fleet, godip.Italy})
	opts := judge.Phase().Options(judge, godip.France)
	tst.AssertOpt(t, opts, []string{"por", "Move", "por", "spa"})
	tst.AssertOpt(t, opts, []string{"mid", "Support", "mid", "por", "spa"})
}

// Test that we can't build in a captured home center.
func TestCantBuildInCapturedHomeCenter(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Fall, godip.Adjustment))
	judge.SetSC("mun", godip.France)
	// Check the option to build is not available.
	opts := judge.Phase().Options(judge, godip.France)
	tst.AssertNoOpt(t, opts, []string{"mun", "Build", "Army", "mun"})
	// Issue the order anyway.
	judge.SetOrder("mun", orders.Build("mun", godip.Army, time.Now()))
	judge.Next()
	// Check that it was not successful.
	tst.AssertNoUnit(t, judge, "mun")
}

func TestConvoySupportBreaking(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("eng", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("lon", godip.Unit{godip.Army, godip.England})
	judge.SetUnit("bel", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("nth", godip.Unit{godip.Fleet, godip.France})
	judge.SetOrder("eng", orders.Convoy("eng", "lon", "bel"))
	judge.SetOrder("lon", orders.Move("lon", "bel"))
	judge.SetOrder("bel", orders.SupportMove("bel", "nth", "eng"))
	judge.SetOrder("nth", orders.Move("nth", "eng"))
	judge.Next()
	if found, ok := judge.Resolutions()["eng"].(godip.ErrConvoyDislodged); !ok {
		t.Errorf("Wanted eng to have ErrConvoyDislodged, got %v", found)
	}
}

func TestForceDisbandTracking(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("pie", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("mar", godip.Unit{godip.Army, godip.France})
	judge.SetUnit("tyr", godip.Unit{godip.Army, godip.France})
	judge.SetUnit("ven", godip.Unit{godip.Army, godip.France})
	judge.SetUnit("tus", godip.Unit{godip.Army, godip.France})
	judge.SetOrder("tyr", orders.Move("tyr", "pie"))
	judge.SetOrder("ven", orders.SupportMove("ven", "tyr", "pie"))
	judge.Next()
	if len(judge.ForceDisbands()) != 1 {
		t.Errorf("Wanted 1 forced disband, got %v", judge.ForceDisbands())
	}
	if !judge.ForceDisbands()["pie"] {
		t.Errorf("Wanted pie to be force disbanded, but it wasn't?")
	}
}

func TestForceDisbandAdjustment(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Fall, godip.Adjustment))
	judge.SetUnit("pie", godip.Unit{godip.Army, godip.Italy})
	judge.Next()
	if res, found := judge.Resolutions()["pie"]; found {
		t.Errorf("Wanted no resolution for pie, got %v", res)
	}
	if !judge.ForceDisbands()["pie"] {
		t.Errorf("Wanted pie to be force disbanded, but it wasn't?")
	}
}

func TestForceDisbandRetreat(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Fall, godip.Retreat))
	judge.SetDislodged("pie", godip.Unit{godip.Army, godip.Italy})
	judge.Next()
	if res, found := judge.Resolutions()["pie"]; found {
		t.Errorf("Wanted no resolution for pie, got %v", res)
	}
	if !judge.ForceDisbands()["pie"] {
		t.Errorf("Wanted pie to be force disbanded, but it wasn't?")
	}
}

func TestInvalidSupportOrders(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("pic", godip.Unit{godip.Army, godip.France})
	judge.SetUnit("par", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("bur", godip.Unit{godip.Army, godip.England})
	judge.SetUnit("bre", godip.Unit{godip.Army, godip.Austria})
	judge.SetOrder("pic", orders.Move("pic", "bel"))
	judge.SetOrder("bur", orders.SupportMove("bur", "pic", "par"))
	judge.SetOrder("par", orders.SupportHold("par", "pic"))
	judge.Next()
	if found := judge.Resolutions()["bur"]; found != godip.ErrInvalidSupporteeOrder {
		t.Errorf("Wanted InvalidSUpporteeOrder, got %v", found)
	}
	if found := judge.Resolutions()["par"]; found != godip.ErrInvalidSupporteeOrder {
		t.Errorf("Wanted InvalidSUpporteeOrder, got %v", found)
	}
}

func TestMessages(t *testing.T) {
	judge := startState(t)
	tst.WaitForPhases(judge, 4)
	// Transfer two SCs from England to France.
	judge.SetSC("lon", godip.France)
	judge.SetSC("edi", godip.France)
	// Free up an SC and a HC for France, leaving two units for France and two for England.
	judge.RemoveUnit("par")
	judge.RemoveUnit("lon")

	messages := judge.Phase().Messages(judge, godip.France)

	// France can only build one unit because only one HC is free.
	for _, expected := range []string{"MayBuild:1", "OtherMustDisband:England:1"} {
		found := false
		for _, message := range messages {
			if message == expected {
				found = true
			}
		}
		if !found {
			t.Errorf("Expected to find message %v but got %v.", expected, messages)
		}
	}
}

func TestAdjacentConvoyOtherFleetViaConvoy(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("nap", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("wes", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("ion", godip.Unit{godip.Fleet, godip.France})
	judge.SetOrder("nap", orders.Move("nap", "rom").ViaConvoy())
	judge.SetOrder("tys", orders.Convoy("tys", "nap", "rom"))
	judge.SetOrder("wes", orders.Move("wes", "tys"))
	judge.SetOrder("ion", orders.SupportMove("wes", "wes", "tys"))
	judge.Next()
	if found := judge.Resolutions()["nap"]; found != godip.ErrMissingConvoyPath {
		t.Errorf("Wanted failure for nap, got %v", found)
	}
	if found, ok := judge.Resolutions()["tys"].(godip.ErrConvoyDislodged); !ok {
		t.Errorf("Wanted failure for tys, got %v", found)
	}
	if found := judge.Resolutions()["wes"]; found != nil {
		t.Errorf("Wanted success for wes, got %v", found)
	}
	if found := judge.Resolutions()["ion"]; found != nil {
		t.Errorf("Wanted success for ion, got %v", found)
	}
}

func TestAdjacentConvoyOtherFleet(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("nap", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("wes", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("ion", godip.Unit{godip.Fleet, godip.France})
	judge.SetOrder("nap", orders.Move("nap", "rom"))
	judge.SetOrder("tys", orders.Convoy("tys", "nap", "rom"))
	judge.SetOrder("wes", orders.Move("wes", "tys"))
	judge.SetOrder("ion", orders.SupportMove("wes", "wes", "tys"))
	judge.Next()
	if found := judge.Resolutions()["nap"]; found != nil {
		t.Errorf("Wanted success for nap, got %v", found)
	}
	if found, ok := judge.Resolutions()["tys"].(godip.ErrConvoyDislodged); !ok {
		t.Errorf("Wanted failure for tys, got %v", found)
	}
	if found := judge.Resolutions()["wes"]; found != nil {
		t.Errorf("Wanted success for wes, got %v", found)
	}
	if found := judge.Resolutions()["ion"]; found != nil {
		t.Errorf("Wanted success for ion, got %v", found)
	}
}

func TestAdjacentConvoyOtherFleetWithEnemy(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("nap", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("rom", godip.Unit{godip.Army, godip.Germany})
	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("wes", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("ion", godip.Unit{godip.Fleet, godip.France})
	judge.SetOrder("rom", orders.Move("rom", "nap"))
	judge.SetOrder("nap", orders.Move("nap", "rom"))
	judge.SetOrder("tys", orders.Convoy("tys", "nap", "rom"))
	judge.SetOrder("wes", orders.Move("wes", "tys"))
	judge.SetOrder("ion", orders.SupportMove("wes", "wes", "tys"))
	judge.Next()
	if found, ok := judge.Resolutions()["nap"].(godip.ErrBounce); !ok {
		t.Errorf("Wanted failure for nap, got %v", found)
	}
	if found, ok := judge.Resolutions()["tys"].(godip.ErrConvoyDislodged); !ok {
		t.Errorf("Wanted failure for tys, got %v", found)
	}
	if found := judge.Resolutions()["wes"]; found != nil {
		t.Errorf("Wanted success for wes, got %v", found)
	}
	if found := judge.Resolutions()["ion"]; found != nil {
		t.Errorf("Wanted success for ion, got %v", found)
	}
}

func TestAdjacentConvoyOwnFleet(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("nap", godip.Unit{godip.Army, godip.Italy})
	judge.SetUnit("tys", godip.Unit{godip.Fleet, godip.Italy})
	judge.SetUnit("wes", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("ion", godip.Unit{godip.Fleet, godip.France})
	judge.SetOrder("nap", orders.Move("nap", "rom"))
	judge.SetOrder("tys", orders.Convoy("tys", "nap", "rom"))
	judge.SetOrder("wes", orders.Move("wes", "tys"))
	judge.SetOrder("ion", orders.SupportMove("wes", "wes", "tys"))
	judge.Next()
	if found := judge.Resolutions()["nap"]; found != godip.ErrMissingConvoyPath {
		t.Errorf("Wanted failure for nap, got %v", found)
	}
	if found, ok := judge.Resolutions()["tys"].(godip.ErrConvoyDislodged); !ok {
		t.Errorf("Wanted failure for tys, got %v", found)
	}
	if found := judge.Resolutions()["wes"]; found != nil {
		t.Errorf("Wanted success for wes, got %v", found)
	}
	if found := judge.Resolutions()["ion"]; found != nil {
		t.Errorf("Wanted success for ion, got %v", found)
	}
}

func TestAdjacentConvoyOwnFleetUnnecessaryParticipant(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("wal", godip.Unit{godip.Army, godip.England})
	judge.SetUnit("eng", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("nrg", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("iri", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("nao", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("nth", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("ska", godip.Unit{godip.Fleet, godip.Germany})
	judge.SetUnit("hel", godip.Unit{godip.Fleet, godip.Germany})
	judge.SetOrder("wal", orders.Move("wal", "lon"))
	judge.SetOrder("nrg", orders.Convoy("nrg", "wal", "lon"))
	judge.SetOrder("ska", orders.Move("ska", "nth"))
	judge.SetOrder("hel", orders.SupportMove("ska", "ska", "nth"))
	judge.Next()
	if found := judge.Resolutions()["wal"]; found != godip.ErrMissingConvoyPath {
		t.Errorf("Wanted failure for wal, got %v", found)
	}
}

func TestAdjacentConvoyOwnFleetUnnecessaryParticipantDislodged(t *testing.T) {
	judge := Blank(NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("wal", godip.Unit{godip.Army, godip.England})
	judge.SetUnit("eng", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("nrg", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("iri", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("nao", godip.Unit{godip.Fleet, godip.France})
	judge.SetUnit("nth", godip.Unit{godip.Fleet, godip.England})
	judge.SetUnit("ska", godip.Unit{godip.Fleet, godip.Germany})
	judge.SetUnit("hel", godip.Unit{godip.Fleet, godip.Germany})
	judge.SetOrder("wal", orders.Move("wal", "lon"))
	judge.SetOrder("nth", orders.Convoy("nrg", "wal", "lon"))
	judge.SetOrder("ska", orders.Move("ska", "nth"))
	judge.SetOrder("hel", orders.SupportMove("ska", "ska", "nth"))
	judge.Next()
	if found := judge.Resolutions()["wal"]; found != godip.ErrMissingConvoyPath {
		t.Errorf("Wanted failure for wal, got %v", found)
	}
}
