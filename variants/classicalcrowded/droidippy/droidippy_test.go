package droidippy

import (
	"testing"

	"github.com/zond/godip/variants/classicalcrowded"

	dip "github.com/zond/godip"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func TestDroidippyGames(t *testing.T) {
	tst.TestGames(t, classicalcrowded.ClassicalCrowdedVariant)
}
