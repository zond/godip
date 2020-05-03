package youngstownredux

import (
	"testing"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"

	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
}

func startState(t *testing.T) *state.State {
	judge, err := YoungstownReduxStart()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return judge
}

func blankState(t *testing.T) *state.State {
	startPhase := classical.NewPhase(1901, godip.Spring, godip.Movement)
	judge := YoungstownReduxBlank(startPhase)
	return judge
}

func TestHebei(t *testing.T) {
	judge := startState(t)

	// Test (and document) that there is no connection from Hebei South Coast to Yellow Sea.
	judge.SetUnit("heb/sc", godip.Unit{godip.Fleet, Japan})
	tst.AssertOrderValidity(t, judge, orders.Move("heb/sc", "yel"), "", godip.ErrIllegalMove)

	// Check that this is possible from the North Coast.
	judge.RemoveUnit("heb/sc")
	judge.SetUnit("heb/nc", godip.Unit{godip.Fleet, Japan})
	tst.AssertOrderValidity(t, judge, orders.Move("heb/nc", "yel"), Japan, nil)

	// Check the reverse direction.
	judge.RemoveUnit("heb/nc")
	judge.SetUnit("yel", godip.Unit{godip.Fleet, Japan})
	tst.AssertOrderValidity(t, judge, orders.Move("yel", "heb/sc"), "", godip.ErrIllegalMove)
	tst.AssertOrderValidity(t, judge, orders.Move("yel", "heb/nc"), Japan, nil)
}

func TestBoxes(t *testing.T) {
	judge := startState(t)

	// Test some of the connections between boxes.
	judge.SetUnit("bxa", godip.Unit{godip.Fleet, Britain})
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxb")
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxc")
	tst.AssertOptionToMove(t, judge, Britain, "bxa", "bxd")
	tst.AssertNoOptionToMoveTo(t, judge, Britain, "bxa", "npo")
	tst.AssertNoOptionToMoveTo(t, judge, Britain, "bxa", "bxe")

	judge.SetUnit("bxb", godip.Unit{godip.Fleet, France})
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxa")
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxc")
	tst.AssertOptionToMove(t, judge, France, "bxb", "bxe")
	tst.AssertNoOptionToMoveTo(t, judge, France, "bxb", "bxg")

	judge.SetUnit("bxc", godip.Unit{godip.Fleet, Italy})
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxa")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxb")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxf")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxg")
	tst.AssertOptionToMove(t, judge, Italy, "bxc", "bxh")
}

func TestMogadishu(t *testing.T) {
	judge := startState(t)

	// Test that there is no sea connection between Mogadishu and Ethiopia.
	tst.AssertOrderValidity(t, judge, orders.Move("mog", "eth"), "", godip.ErrIllegalMove)
	judge.RemoveUnit("mog")
	judge.SetUnit("eth", godip.Unit{godip.Fleet, Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("eth", "mog"), "", godip.ErrIllegalMove)

	// Test that there is a land connection between Mogadishu and Ethiopia.
	judge.RemoveUnit("eth")
	judge.SetUnit("mog", godip.Unit{godip.Army, Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("mog", "eth"), Italy, nil)
	judge.RemoveUnit("mog")
	judge.SetUnit("eth", godip.Unit{godip.Army, Italy})
	tst.AssertOrderValidity(t, judge, orders.Move("eth", "mog"), Italy, nil)
}

func TestSillyNilPointer(t *testing.T) {
	judge := YoungstownReduxBlank(classical.NewPhase(1901, godip.Spring, godip.Movement))
	judge.SetUnit("mid", godip.Unit{godip.Fleet, Italy})
	judge.SetUnit("por", godip.Unit{godip.Fleet, Italy})
	judge.SetUnit("spa/sc", godip.Unit{godip.Fleet, Italy})
	judge.SetUnit("wms", godip.Unit{godip.Fleet, France})
	judge.SetOrder("por", orders.Move("por", "mid"))
	judge.SetOrder("mid", orders.Move("mid", "wms"))
	judge.SetOrder("spa", orders.SupportMove("spa", "mid", "wms"))
	judge.Next()
	tst.AssertUnit(t, judge, "wms", godip.Unit{godip.Fleet, Italy})
}
