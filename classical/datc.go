package classical

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	cla "github.com/zond/godip/classical/common"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/classical/start"
	dip "github.com/zond/godip/common"
)

func init() {
	for _, prov := range start.Graph().Provinces() {
		datcProvinces[string(prov)] = prov
	}
}

var datcPhaseTypes = map[string]dip.PhaseType{
	"movement":   cla.Movement,
	"adjustment": cla.Adjustment,
	"retreat":    cla.Retreat,
}

var datcSeasons = map[string]dip.Season{
	"spring": cla.Spring,
	"fall":   cla.Fall,
}

var datcNationalities = map[string]dip.Nation{
	"england": cla.England,
	"france":  cla.France,
	"germany": cla.Germany,
	"russia":  cla.Russia,
	"austria": cla.Austria,
	"italy":   cla.Italy,
	"turkey":  cla.Turkey,
	"germnay": cla.Germany,
}

var datcUnitTypes = map[string]dip.UnitType{
	"a": cla.Army,
	"f": cla.Fleet,
}

var datcProvinces = map[string]dip.Province{}

func DATCPhase(season string, year int, typ string) (result dip.Phase, err error) {
	phaseType, ok := datcPhaseTypes[strings.ToLower(typ)]
	if !ok {
		err = fmt.Errorf("Unknown phase type %#v", typ)
		return
	}
	phaseSeason, ok := datcSeasons[strings.ToLower(season)]
	if !ok {
		err = fmt.Errorf("Unknown season %#v", season)
		return
	}
	result = &phase{
		season: phaseSeason,
		typ:    phaseType,
		year:   year,
	}
	return
}

func DATCProvince(n string) (result dip.Province, err error) {
	var ok bool
	result, ok = datcProvinces[strings.ToLower(n)]
	if !ok {
		err = fmt.Errorf("Unknown province %#v", n)
		return
	}
	return
}

var datcOrderTypes = map[*regexp.Regexp]func([]string) (dip.Province, dip.Adjudicator, error){
	regexp.MustCompile("(?i)^(A|F)\\s+(\\S+)\\s*-\\s*(\\S+)(\\s+via\\s+convoy)?$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		dst, err := DATCProvince(m[3])
		if err != nil {
			return
		}
		if m[4] == "" {
			order = orders.Move(prov, dst)
		} else {
			order = orders.Move(prov, dst).ViaConvoy()
		}
		return
	},
	regexp.MustCompile("^(?i)remove\\s+((A|F)\\s+)?(\\S+)$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[3]); err != nil {
			return
		}
		order = orders.Disband(prov, time.Now())
		return
	},
	regexp.MustCompile("^(?i)(A|F)\\s+(\\S+)\\s+disband$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		order = orders.Disband(prov, time.Now())
		return
	},
	regexp.MustCompile("^(?i)(A|F)\\s+(\\S+)\\s+S(UPP\\S*)?\\s+(A|F)\\s+([^-\\s]+)$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		target, err := DATCProvince(m[5])
		if err != nil {
			return
		}
		order = orders.SupportHold(prov, target)
		return
	},
	regexp.MustCompile("^(?i)(A|F)\\s+(\\S+)\\s+c(onv\\S*)?\\s+(A|F)\\s+(\\S+)\\s*-\\s*(\\S+)$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		from, err := DATCProvince(m[5])
		if err != nil {
			return
		}
		to, err := DATCProvince(m[6])
		if err != nil {
			return
		}
		order = orders.Convoy(prov, from, to)
		return
	},
	regexp.MustCompile("^(?i)(A|F)\\s+(\\S+)\\s+S(UPP\\S*)?\\s+((A|F)\\s+)?(\\S+)\\s*-\\s*(\\S+)$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		from, err := DATCProvince(m[6])
		if err != nil {
			return
		}
		to, err := DATCProvince(m[7])
		if err != nil {
			return
		}
		order = orders.SupportMove(prov, from, to)
		return
	},
	regexp.MustCompile("^(?i)(A|F)\\s+(\\S+)\\s+H(OLD)?$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		order = orders.Hold(prov)
		return
	},
	regexp.MustCompile("^(?i)build\\s+(A|F)\\s+(\\S+)\\s*$"): func(m []string) (prov dip.Province, order dip.Adjudicator, err error) {
		if prov, err = DATCProvince(m[2]); err != nil {
			return
		}
		unitType, err := DATCUnitType(m[1])
		if err != nil {
			return
		}
		order = orders.Build(prov, unitType, time.Now())
		return
	},
}

func DATCOrder(text string) (province dip.Province, order dip.Adjudicator, err error) {
	var match []string
	for reg, gen := range datcOrderTypes {
		if match = reg.FindStringSubmatch(text); match != nil {
			return gen(match)
		}
	}
	err = fmt.Errorf("Unknown order text: %#v", text)
	return
}

func DATCNation(typ string) (result dip.Nation, err error) {
	var ok bool
	result, ok = datcNationalities[strings.ToLower(typ)]
	if !ok {
		err = fmt.Errorf("Unknown nationality: %#v", typ)
		return
	}
	return
}

func DATCUnitType(typ string) (result dip.UnitType, err error) {
	var ok bool
	result, ok = datcUnitTypes[strings.ToLower(typ)]
	if !ok {
		err = fmt.Errorf("Unknown unit type: %#v", typ)
		return
	}
	return
}
