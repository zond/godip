package droidippy

import (
	"testing"

	"github.com/zond/godip/variants/classical"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func TestDroidippyGames(t *testing.T) {
	tst.TestGames(t, cla.Nations, classical.Start, classical.Blank)
}
