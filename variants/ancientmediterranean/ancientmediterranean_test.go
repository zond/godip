package ancientmediterranean

import (
	"testing"
	"time"

	"github.com/zond/godip/state"
	"github.com/zond/godip/orders"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := AncientMediterraneanStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func TestByzantium(t *testing.T) {
	judge := startState(t)

	// Test can't bypass Byzantium.
	judge.SetUnit("aeg", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("aeg", "bla"), "", cla.ErrIllegalMove)

	// Test can sail through it.
	tst.AssertOrderValidity(t, judge, orders.Move("aeg", "byz"), Greece, nil)
	judge.SetUnit("byz", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("byz", "bla"), Greece, nil)

	// Check can't convoy via Byzantium.
	tst.AssertOrderValidity(t, judge, orders.Move("ath", "bit"), "", cla.ErrMissingConvoyPath)
}

func TestHighSeas(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("aus", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "mes"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "lib"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "got"), Greece, nil)
	judge.SetUnit("mes", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("mes", "lib"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("mes", "got"), Greece, nil)
	judge.SetUnit("lib", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("lib", "got"), Greece, nil)
}

func TestDiolkos(t *testing.T) {
	judge := startState(t)

	// Test can't bypass Athens or Sparta.
	judge.SetUnit("ion", dip.Unit{cla.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("ion", "aeg"), "", cla.ErrIllegalMove)

	// Test can walk from Athens to Sparta.
	tst.AssertOrderValidity(t, judge, orders.Move("ath", "spa"), Greece, nil)
}

func TestSicily(t *testing.T) {
	judge := startState(t)

	// Test can walk or sail between Sicily and Neapolis.
	judge.SetUnit("sic", dip.Unit{cla.Army, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sic", "nea"), Rome, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nea", "sic"), Rome, nil)

	// Test can sail through the 'Strait of Messina'.
	judge.SetUnit("tys", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("tys", "aus"), Rome, nil)
}

func TestCorsica(t *testing.T) {
	judge := startState(t)

	// Test can walk or sail between Corsica and Sardinia.
	judge.SetUnit("cor", dip.Unit{cla.Army, Rome})
	judge.SetUnit("sad", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("cor", "sad"), Rome, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("sad", "cor"), Rome, nil)
}

func TestNileDelta(t *testing.T) {
	judge := startState(t)

	// Happy paths near Nile Delta
	tst.AssertOrderValidity(t, judge, orders.Move("the", "sii"), Egypt, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportHold("the", "mem"), Egypt, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportHold("the", "ale"), Egypt, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("mem", "the", "ale"), Egypt, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("ale", "sii"), Egypt, nil)
	tst.AssertOrderValidity(t, judge, orders.SupportMove("the", "ale", "sii"), Egypt, nil)
	judge.SetUnit("gop", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("the", "jer"), Egypt, nil)

	// Illegal moves near Nile Delta
	judge.SetUnit("ree", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("ree", "ale"), "", cla.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("ree", "gop"), "", cla.ErrIllegalMove)
	judge.RemoveUnit("mem")
	judge.SetUnit("mem", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("mem", "sii"), "", cla.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("mem", "gop"), "", cla.ErrIllegalMove)
}

func TestConvoyBaleares(t *testing.T) {
	judge := startState(t)

	// Test convoys through Baleares.
	judge.SetUnit("sag", dip.Unit{cla.Army, Rome})
	judge.SetUnit("bal", dip.Unit{cla.Fleet, Rome})
	judge.SetUnit("lig", dip.Unit{cla.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sag", "cor"), Rome, nil)
	tst.AssertOrderValidity(t, judge, orders.Convoy("bal", "sag", "cor"), Rome, nil)

	// Test an army in Baleares can't be part of a convoy chain.
	judge.RemoveUnit("bal")
	judge.SetUnit("bal", dip.Unit{cla.Army, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sag", "cor"), "", cla.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Convoy("bal", "sag", "cor"), "", cla.ErrIllegalConvoyer)
}

func TestAutomaticDisbands(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("car")
	judge.RemoveUnit("cir")
	judge.RemoveUnit("tha")
	// Give original HCs to Rome
	judge.SetUnit("cir", dip.Unit{cla.Army, Rome})
	judge.SetUnit("tha", dip.Unit{cla.Army, Rome})

	// Set up Carthage position from https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAwI6gjQoM/Phase/35/Map
	judge.SetUnit("tar", dip.Unit{cla.Army, Carthage})
	judge.SetUnit("bal", dip.Unit{cla.Fleet, Carthage})
	judge.SetUnit("ber", dip.Unit{cla.Fleet, Carthage})
	judge.SetUnit("pun", dip.Unit{cla.Fleet, Carthage})
	judge.SetUnit("car", dip.Unit{cla.Fleet, Carthage})

	// Spring movement
	judge.Next()
	// Sprint retreat
	judge.Next()
	// Fall movement
	judge.Next()
	// Fall retreat
	judge.Next()
	// Order contains one disband but should have three.
	judge.SetOrder("ber", orders.Disband("ber", time.Now()))
	judge.Next()
	// Check that automatic disbands worked.
	tst.AssertNoUnit(t, judge, "tar")
	tst.AssertNoUnit(t, judge, "bal")
	tst.AssertNoUnit(t, judge, "ber")
	tst.AssertUnit(t, judge, "pun", dip.Unit{cla.Fleet, Carthage})
	tst.AssertUnit(t, judge, "car", dip.Unit{cla.Fleet, Carthage})
}

func TestSuggestedMoveBaleares(t *testing.T) {
	judge := startState(t)

	// In Spring 8AD of this game Godip suggested moving an army in Saguntum to Baleares:
	// https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAwI6gjQoM/Phase/36
	judge.SetUnit("sag", dip.Unit{cla.Army, Rome})

	// Test there's no suggestion of a move from Saguntum to Baleares when the destination is empty and there's
	// no fleet to convoy.
	tst.AssertNoOptionToMoveTo(t, judge, Rome, "sag", "bal")

	// Test there's no suggestion of a move from Saguntum to Baleares when the destination contains a fleet
	// but there's still no fleet to convoy.
	judge.SetUnit("bal", dip.Unit{cla.Fleet, Carthage})
	tst.AssertNoOptionToMoveTo(t, judge, Rome, "sag", "bal")

	// Test there IS a suggestion of a move from sag to bal when there is a fleet in ber.
	judge.SetUnit("ber", dip.Unit{cla.Fleet, Carthage})
	tst.AssertOptionToMove(t, judge, Rome, "sag", "bal")

}
