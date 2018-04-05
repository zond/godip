package common

import (
	"testing"

	"github.com/zond/godip/state"

	dip "github.com/zond/godip"
)

const (
	Austria dip.Nation = "Austria"
	England dip.Nation = "England"

	Ankara         dip.Province = "Ankara"
	Belgium        dip.Province = "Belgium"
	Constantinople dip.Province = "Constantinople"
	Denmark        dip.Province = "Denmark"
	Edinbugh       dip.Province = "Edinburgh"
)

func init() {
	dip.Debug = true
}

func TestSCCountWinner_NothingOwned(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))

	winner := soloFunction(s)

	if winner != "" {
		t.Errorf("Expected no winner, but got %v", winner)
	}
}

func TestSCCountWinner_LeaderHasntWon(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))
	s.SetSC(Ankara, Austria)

	winner := soloFunction(s)

	if winner != "" {
		t.Errorf("Expected no winner, but got %v", winner)
	}
}

func TestSCCountWinner_ClearWinner(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))
	s.SetSC(Ankara, Austria)
	s.SetSC(Belgium, Austria)
	s.SetSC(Constantinople, England)

	winner := soloFunction(s)

	if winner != Austria {
		t.Errorf("Expected Austria to win, but got %v", winner)
	}
}

func TestSCCountWinner_JointLeader(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))
	s.SetSC(Ankara, Austria)
	s.SetSC(Belgium, Austria)
	s.SetSC(Constantinople, England)
	s.SetSC(Denmark, England)

	winner := soloFunction(s)

	if winner != "" {
		t.Errorf("Expected no winner (since joint leader), but got %v", winner)
	}
}

func TestSCCountWinner_SecondPlaceHasAlsoPassedTarget(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[dip.Province]dip.Nation))
	s.SetSC(Ankara, Austria)
	s.SetSC(Belgium, Austria)
	s.SetSC(Constantinople, England)
	s.SetSC(Denmark, England)
	s.SetSC(Edinbugh, England)

	winner := soloFunction(s)

	if winner != England {
		t.Errorf("Expected England to win, but got %v", winner)
	}
}
