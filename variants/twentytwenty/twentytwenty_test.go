package twentytwenty

import (
	"testing"

	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
)

func init() {
	godip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := TwentyTwentyStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(2001, godip.Spring, godip.Movement)
	judge := TwentyTwentyBlank(startPhase)
	return judge
}

func TestNoWinnerAtStart(t *testing.T) {
	judge := startState(t)

	assertNoWinner(t, judge)
}

func TestTwentyLeadAtStartWins(t *testing.T) {
	judge := startState(t)
	// USA and e.g. Russia start with 4 SCs each, so USA needs 20 more to win in the first year.
	for _, province := range []godip.Province{
		"grd", "mex", "pan", "col", "dom", "alg", "mau", "sen", "gin", "cot",
		"cha", "car", "drc", "ang", "zam", "pai", "now", "swe", "fin", "hun"} {
		judge.SetSC(province, USA)
	}

	assertWinner(t, judge, "USA")
}

func TestNineteenLeadAtStartDoesntWin(t *testing.T) {
	judge := startState(t)
	// If USA has a lead of 19 SCs in 2001 then they don't win.
	for _, province := range []godip.Province{
		"grd", "mex", "pan", "col", "dom", "alg", "mau", "sen", "gin", "cot",
		"cha", "car", "drc", "ang", "zam", "pai", "now", "swe", "fin"} {
		judge.SetSC(province, USA)
	}

	assertNoWinner(t, judge)
}

func TestNineteenLeadInSecondYearWins(t *testing.T) {
	judge := startState(t)
	// Set year to 2002.
	waitForYears(judge, 1)
	// If USA has a lead of 19 SCs in 2002 then they do win.
	for _, province := range []godip.Province{
		"grd", "mex", "pan", "col", "dom", "alg", "mau", "sen", "gin", "cot",
		"cha", "car", "drc", "ang", "zam", "pai", "now", "swe", "fin"} {
		judge.SetSC(province, USA)
	}

	assertWinner(t, judge, "USA")
}

func TestMoreThanHalfWins(t *testing.T) {
	judge := blankState(t)
	// USA has 49, Russia has 48 (and the year is 2001).
	for _, province := range []godip.Province{
		"cho", "cap", "uzb", "ant", "phi", "lon", "mom", "yem", "lah", "man",
		"roe", "atl", "mac", "alm", "ale", "pai", "anc", "ank", "cha", "lag",
		"rec", "mau", "bhu", "tok", "alg", "tai", "bei", "hat", "grd", "dub",
		"now", "vla", "ser", "mum", "abu", "hun", "sap", "isl", "bue", "oms",
		"pam", "mad", "zam", "ist", "kan", "asw", "cad", "yum", "sen"} {
		judge.SetSC(province, USA)
	}
	for _, province := range []godip.Province{
		"rio", "kar", "bri", "dom", "bad", "men", "drc", "iqa", "pet", "syd",
		"rma", "grc", "nag", "cai", "chm", "swe", "dur", "was", "gin", "mex",
		"riy", "mil", "diy", "bar", "ned", "mos", "ham", "car", "nai", "tur",
		"cot", "mar", "van", "bna", "edi", "com", "oma", "ang", "okl", "bao",
		"fin", "ben", "col", "pre", "nap", "mot", "irk", "bnk"} {
		judge.SetSC(province, Russia)
	}

	assertWinner(t, judge, "USA")
}

func TestLessThanHalfDoesntWin(t *testing.T) {
	judge := blankState(t)
	// USA has 48, Russia has 47 (and the year is 2001).
	for _, province := range []godip.Province{
		"cho", "cap", "uzb", "ant", "phi", "lon", "mom", "yem", "lah", "man",
		"roe", "atl", "mac", "alm", "ale", "pai", "anc", "ank", "cha", "lag",
		"rec", "mau", "bhu", "tok", "alg", "tai", "bei", "hat", "grd", "dub",
		"now", "vla", "ser", "mum", "abu", "hun", "sap", "isl", "bue", "oms",
		"pam", "mad", "zam", "ist", "kan", "asw", "cad", "yum"} {
		judge.SetSC(province, USA)
	}
	for _, province := range []godip.Province{
		"rio", "kar", "bri", "dom", "bad", "men", "drc", "iqa", "pet", "syd",
		"rma", "grc", "nag", "cai", "chm", "swe", "dur", "was", "gin", "mex",
		"riy", "mil", "diy", "bar", "ned", "mos", "ham", "car", "nai", "tur",
		"cot", "mar", "van", "bna", "edi", "com", "oma", "ang", "okl", "bao",
		"fin", "ben", "col", "pre", "nap", "mot", "irk"} {
		judge.SetSC(province, Russia)
	}

	assertNoWinner(t, judge)
}

func TestOneCenterLeadWinsIn2020(t *testing.T) {
	judge := startState(t)
	judge.SetSC("grd", USA)

	// No winning in 2019 with a lead of one.
	waitForYears(judge, 18)
	assertNoWinner(t, judge)

	// USA wins in 2020 with a lead of one.
	waitForYears(judge, 1)
	assertWinner(t, judge, USA)
}

func TestOneCenterLeadWinsIn2021(t *testing.T) {
	judge := startState(t)

	// No winner if tied in 2020
	waitForYears(judge, 19)
	assertNoWinner(t, judge)

	// Still no winner if tied in 2021
	waitForYears(judge, 1)
	assertNoWinner(t, judge)

	// Winner with lead of one in 2021
	judge.SetSC("grd", USA)
	assertWinner(t, judge, USA)
}

func assertNoWinner(t *testing.T, judge *state.State) {
	winner := TwentyTwentyWinner(judge)
	if winner != "" {
		t.Errorf("Expected no winner but got %v", winner)
	}
}

func assertWinner(t *testing.T, judge *state.State, expected godip.Nation) {
	winner := TwentyTwentyWinner(judge)
	if winner != expected {
		t.Errorf("Expected %v to win but got %v", expected, winner)
	}
}

// Increase the current phase by the given number of years.
func waitForYears(judge *state.State, years int) {
	for year := 0; year < years; year++ {
		for phase := 0; phase < 5; phase++ {
			judge.Next()
		}
	}
}
