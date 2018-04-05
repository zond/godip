package start

import (
	"github.com/zond/godip"

	. "github.com/zond/godip/variants/classical/common"
)

func SupplyCenters() map[godip.Province]godip.Nation {
	return map[godip.Province]godip.Nation{
		"edi": England,
		"lvp": England,
		"lon": England,
		"bre": France,
		"par": France,
		"mar": France,
		"kie": Germany,
		"ber": Germany,
		"mun": Germany,
		"ven": Italy,
		"rom": Italy,
		"nap": Italy,
		"tri": Austria,
		"vie": Austria,
		"bud": Austria,
		"con": Turkey,
		"ank": Turkey,
		"smy": Turkey,
		"sev": Russia,
		"mos": Russia,
		"stp": Russia,
		"war": Russia,
	}
}
