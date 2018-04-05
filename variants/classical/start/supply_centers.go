package start

import (
	"github.com/zond/godip"
)

func SupplyCenters() map[godip.Province]godip.Nation {
	return map[godip.Province]godip.Nation{
		"edi": godip.England,
		"lvp": godip.England,
		"lon": godip.England,
		"bre": godip.France,
		"par": godip.France,
		"mar": godip.France,
		"kie": godip.Germany,
		"ber": godip.Germany,
		"mun": godip.Germany,
		"ven": godip.Italy,
		"rom": godip.Italy,
		"nap": godip.Italy,
		"tri": godip.Austria,
		"vie": godip.Austria,
		"bud": godip.Austria,
		"con": godip.Turkey,
		"ank": godip.Turkey,
		"smy": godip.Turkey,
		"sev": godip.Russia,
		"mos": godip.Russia,
		"stp": godip.Russia,
		"war": godip.Russia,
	}
}
