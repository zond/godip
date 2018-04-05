package start

import (
	"github.com/zond/godip"

	. "github.com/zond/godip/variants/classical/common"
)

func Units() map[godip.Province]godip.Unit {
	return map[godip.Province]godip.Unit{
		"edi":    godip.Unit{Fleet, England},
		"lvp":    godip.Unit{Army, England},
		"lon":    godip.Unit{Fleet, England},
		"bre":    godip.Unit{Fleet, France},
		"par":    godip.Unit{Army, France},
		"mar":    godip.Unit{Army, France},
		"kie":    godip.Unit{Fleet, Germany},
		"ber":    godip.Unit{Army, Germany},
		"mun":    godip.Unit{Army, Germany},
		"ven":    godip.Unit{Army, Italy},
		"rom":    godip.Unit{Army, Italy},
		"nap":    godip.Unit{Fleet, Italy},
		"tri":    godip.Unit{Fleet, Austria},
		"vie":    godip.Unit{Army, Austria},
		"bud":    godip.Unit{Army, Austria},
		"stp/sc": godip.Unit{Fleet, Russia},
		"mos":    godip.Unit{Army, Russia},
		"war":    godip.Unit{Army, Russia},
		"sev":    godip.Unit{Fleet, Russia},
		"con":    godip.Unit{Army, Turkey},
		"smy":    godip.Unit{Army, Turkey},
		"ank":    godip.Unit{Fleet, Turkey},
	}
}
