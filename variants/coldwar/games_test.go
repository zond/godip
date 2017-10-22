package coldwar

import (
	"testing"

	dip "github.com/zond/godip/common"
	tst "github.com/zond/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func TestGames(t *testing.T) {
	tst.TestGames(t, Nations, ColdWarStart, ColdWarBlank)
}
