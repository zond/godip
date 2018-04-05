package start

import (
	"github.com/zond/godip"
)

func Units() map[godip.Province]godip.Unit {
	return map[godip.Province]godip.Unit{
		"edi":    godip.Unit{godip.Fleet, godip.England},
		"lvp":    godip.Unit{godip.Army, godip.England},
		"lon":    godip.Unit{godip.Fleet, godip.England},
		"bre":    godip.Unit{godip.Fleet, godip.France},
		"par":    godip.Unit{godip.Army, godip.France},
		"mar":    godip.Unit{godip.Army, godip.France},
		"kie":    godip.Unit{godip.Fleet, godip.Germany},
		"ber":    godip.Unit{godip.Army, godip.Germany},
		"mun":    godip.Unit{godip.Army, godip.Germany},
		"ven":    godip.Unit{godip.Army, godip.Italy},
		"rom":    godip.Unit{godip.Army, godip.Italy},
		"nap":    godip.Unit{godip.Fleet, godip.Italy},
		"tri":    godip.Unit{godip.Fleet, godip.Austria},
		"vie":    godip.Unit{godip.Army, godip.Austria},
		"bud":    godip.Unit{godip.Army, godip.Austria},
		"stp/sc": godip.Unit{godip.Fleet, godip.Russia},
		"mos":    godip.Unit{godip.Army, godip.Russia},
		"war":    godip.Unit{godip.Army, godip.Russia},
		"sev":    godip.Unit{godip.Fleet, godip.Russia},
		"con":    godip.Unit{godip.Army, godip.Turkey},
		"smy":    godip.Unit{godip.Army, godip.Turkey},
		"ank":    godip.Unit{godip.Fleet, godip.Turkey},
	}
}
