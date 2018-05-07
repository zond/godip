package chaos

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	tst "github.com/zond/godip/variants/testing"
)

func TestDefaultBuild(t *testing.T) {
	j, err := Start()
	if err != nil {
		t.Fatal(err)
	}
	j.SetOrders(map[godip.Province]godip.Adjudicator{
		"tri": orders.BuildAnywhere("tri", godip.Fleet, time.Now()),
	})
	j.Next()
	tst.AssertUnit(t, j, "tri", godip.Unit{godip.Fleet, "Trieste"})
	tst.AssertUnit(t, j, "vie", godip.Unit{godip.Army, "Vienna"})
}
