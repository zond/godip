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
		"nwy":    godip.Unit{godip.Fleet, godip.Scandinavia},
		"swe":    godip.Unit{godip.Army, godip.Scandinavia},
		"den":    godip.Unit{godip.Fleet, godip.Scandinavia},
		"hol":    godip.Unit{godip.Army, godip.Benelux},
		"bel":    godip.Unit{godip.Fleet, godip.Benelux},
		"ruh":    godip.Unit{godip.Army, godip.Benelux},
		"spa":    godip.Unit{godip.Army, godip.Iberia},
		"por":    godip.Unit{godip.Fleet, godip.Iberia},
		"tun":    godip.Unit{godip.Fleet, godip.Iberia},
		"rum":    godip.Unit{godip.Fleet, godip.Balkans},
		"ser":    godip.Unit{godip.Army, godip.Balkans},
		"bul":    godip.Unit{godip.Army, godip.Balkans},
		"gre":    godip.Unit{godip.Fleet, godip.Balkans},
	}
}
