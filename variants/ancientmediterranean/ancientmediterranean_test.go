package ancientmediterranean

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"

	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
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
	judge.SetUnit("aeg", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("aeg", "bla"), "", godip.ErrIllegalMove)

	// Test can sail through it.
	tst.AssertOrderValidity(t, judge, orders.Move("aeg", "byz"), Greece, nil)
	judge.SetUnit("byz", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("byz", "bla"), Greece, nil)

	// Check can't convoy via Byzantium.
	tst.AssertOrderValidity(t, judge, orders.Move("ath", "bit"), "", godip.ErrMissingConvoyPath)
}

func TestHighSeas(t *testing.T) {
	judge := startState(t)

	judge.SetUnit("aus", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "mes"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "lib"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("aus", "got"), Greece, nil)
	judge.SetUnit("mes", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("mes", "lib"), Greece, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("mes", "got"), Greece, nil)
	judge.SetUnit("lib", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("lib", "got"), Greece, nil)
}

func TestDiolkos(t *testing.T) {
	judge := startState(t)

	// Test can't bypass Athens or Sparta.
	judge.SetUnit("ion", godip.Unit{godip.Fleet, Greece})
	tst.AssertOrderValidity(t, judge, orders.Move("ion", "aeg"), "", godip.ErrIllegalMove)

	// Test can walk from Athens to Sparta.
	tst.AssertOrderValidity(t, judge, orders.Move("ath", "spa"), Greece, nil)
}

func TestSicily(t *testing.T) {
	judge := startState(t)

	// Test can walk or sail between Sicily and Neapolis.
	judge.SetUnit("sic", godip.Unit{godip.Army, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sic", "nea"), Rome, nil)
	tst.AssertOrderValidity(t, judge, orders.Move("nea", "sic"), Rome, nil)

	// Test can sail through the 'Strait of Messina'.
	judge.SetUnit("tys", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("tys", "aus"), Rome, nil)
}

func TestCorsica(t *testing.T) {
	judge := startState(t)

	// Test can walk or sail between Corsica and Sardinia.
	judge.SetUnit("cor", godip.Unit{godip.Army, Rome})
	judge.SetUnit("sad", godip.Unit{godip.Fleet, Rome})
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
	judge.SetUnit("gop", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("the", "jer"), Egypt, nil)

	// Illegal moves near Nile Delta
	judge.SetUnit("ree", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("ree", "ale"), "", godip.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("ree", "gop"), "", godip.ErrIllegalMove)
	judge.RemoveUnit("mem")
	judge.SetUnit("mem", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("mem", "sii"), "", godip.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("mem", "gop"), "", godip.ErrIllegalMove)
}

func TestConvoyBaleares(t *testing.T) {
	judge := startState(t)

	// Test convoys through Baleares.
	judge.SetUnit("sag", godip.Unit{godip.Army, Rome})
	judge.SetUnit("bal", godip.Unit{godip.Fleet, Rome})
	judge.SetUnit("lig", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sag", "cor"), Rome, nil)
	tst.AssertOrderValidity(t, judge, orders.Convoy("bal", "sag", "cor"), Rome, nil)

	// Test an army in Baleares can't be part of a convoy chain.
	judge.RemoveUnit("bal")
	judge.SetUnit("bal", godip.Unit{godip.Army, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("sag", "cor"), "", godip.ErrMissingConvoyPath)
	tst.AssertOrderValidity(t, judge, orders.Convoy("bal", "sag", "cor"), "", godip.ErrIllegalConvoyer)
}

func TestAutomaticDisbands(t *testing.T) {
	judge := startState(t)
	judge.RemoveUnit("car")
	judge.RemoveUnit("cir")
	judge.RemoveUnit("tha")
	// Give original HCs to Rome
	judge.SetUnit("cir", godip.Unit{godip.Army, Rome})
	judge.SetUnit("tha", godip.Unit{godip.Army, Rome})

	// Set up Carthage position from https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAwI6gjQoM/Phase/35/Map
	judge.SetUnit("tar", godip.Unit{godip.Army, Carthage})
	judge.SetUnit("bal", godip.Unit{godip.Fleet, Carthage})
	judge.SetUnit("ber", godip.Unit{godip.Fleet, Carthage})
	judge.SetUnit("pun", godip.Unit{godip.Fleet, Carthage})
	judge.SetUnit("car", godip.Unit{godip.Fleet, Carthage})

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
	tst.AssertUnit(t, judge, "pun", godip.Unit{godip.Fleet, Carthage})
	tst.AssertUnit(t, judge, "car", godip.Unit{godip.Fleet, Carthage})
}

func TestSuggestedMoveBaleares(t *testing.T) {
	judge := startState(t)

	// In Spring 8AD of this game Godip suggested moving an army in Saguntum to Baleares:
	// https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAwI6gjQoM/Phase/36
	judge.SetUnit("sag", godip.Unit{godip.Army, Rome})

	// Test there's no suggestion of a move from Saguntum to Baleares when the destination is empty and there's
	// no fleet to convoy.
	tst.AssertNoOptionToMoveTo(t, judge, Rome, "sag", "bal")

	// Test there's no suggestion of a move from Saguntum to Baleares when the destination contains a fleet
	// but there's still no fleet to convoy.
	judge.SetUnit("bal", godip.Unit{godip.Fleet, Carthage})
	tst.AssertNoOptionToMoveTo(t, judge, Rome, "sag", "bal")

	// Test there IS a suggestion of a move from sag to bal when there is a fleet in ber.
	judge.SetUnit("ber", godip.Unit{godip.Fleet, Carthage})
	tst.AssertOptionToMove(t, judge, Rome, "sag", "bal")

}

func TestIsauriaToAegeanAndMinoan(t *testing.T) {
	judge := startState(t)

	// Test can't sail from Isauria to the Aegean or Minoan Seas.
	judge.SetUnit("isa", godip.Unit{godip.Fleet, Rome})
	tst.AssertOrderValidity(t, judge, orders.Move("isa", "aeg"), "", godip.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("isa", "min"), "", godip.ErrIllegalMove)
}
