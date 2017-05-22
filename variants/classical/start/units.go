package start

import (
  . "github.com/zond/godip/classical/common"
  "github.com/zond/godip/common"
)

func Units() map[common.Province]common.Unit {
  return map[common.Province]common.Unit{
    "edi":    common.Unit{Fleet, England},
    "lvp":    common.Unit{Army, England},
    "lon":    common.Unit{Fleet, England},
    "bre":    common.Unit{Fleet, France},
    "par":    common.Unit{Army, France},
    "mar":    common.Unit{Army, France},
    "kie":    common.Unit{Fleet, Germany},
    "ber":    common.Unit{Army, Germany},
    "mun":    common.Unit{Army, Germany},
    "ven":    common.Unit{Army, Italy},
    "rom":    common.Unit{Army, Italy},
    "nap":    common.Unit{Fleet, Italy},
    "tri":    common.Unit{Fleet, Austria},
    "vie":    common.Unit{Army, Austria},
    "bud":    common.Unit{Army, Austria},
    "stp/sc": common.Unit{Fleet, Russia},
    "mos":    common.Unit{Army, Russia},
    "war":    common.Unit{Army, Russia},
    "sev":    common.Unit{Fleet, Russia},
    "con":    common.Unit{Army, Turkey},
    "smy":    common.Unit{Army, Turkey},
    "ank":    common.Unit{Fleet, Turkey},
  }
}
