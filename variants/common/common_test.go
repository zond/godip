package common

import (
	"testing"

	"github.com/zond/godip"
	"github.com/zond/godip/state"
)

const (
	Austria godip.Nation = "Austria"
	England godip.Nation = "England"

	Ankara         godip.Province = "Ankara"
	Belgium        godip.Province = "Belgium"
	Constantinople godip.Province = "Constantinople"
	Denmark        godip.Province = "Denmark"
	Edinbugh       godip.Province = "Edinburgh"
)

func init() {
	godip.Debug = true
}

func TestSCCountWinner_NothingOwned(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))

	winner := soloFunction(s)

	if winner != "" {
		t.Errorf("Expected no winner, but got %v", winner)
	}
}

func TestSCCountWinner_LeaderHasntWon(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))
	s.SetSC(Ankara, Austria)

	winner := soloFunction(s)

	if winner != "" {
		t.Errorf("Expected no winner, but got %v", winner)
	}
}

func TestSCCountWinner_ClearWinner(t *testing.T) {
	soloFunction := SCCountWinner(2)
	s := new(state.State)
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))
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
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))
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
	s.SetSupplyCenters(make(map[godip.Province]godip.Nation))
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
