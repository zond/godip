package hundred

import (
	"fmt"

	"github.com/zond/godip"
	"github.com/zond/godip/phase"
)

const (
	YearSeason godip.Season = "Year"
)

var newPhase = phase.Generator(BuildAnywhereParser, func(phase *phase.Phase) bool {
	return phase.Ty == godip.Retreat && phase.Yr%10 == 0
})

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	if season != YearSeason {
		fmt.Errorf("Warning - Hundred only supports YearSeason, but got {}", season)
	}
	return &hundredPhase{newPhase(year, season, typ)}
}

type hundredPhase struct {
	godip.Phase
}

func (self *hundredPhase) Season() godip.Season {
	return YearSeason
}

func (self *hundredPhase) Next() godip.Phase {
	if self.Type() == godip.Movement {
		return Phase(self.Year(), YearSeason, godip.Retreat)
	} else if self.Type() == godip.Retreat {
		if self.Year()%10 == 5 {
			return Phase(self.Year()+5, YearSeason, godip.Movement)
		} else {
			return Phase(self.Year(), YearSeason, godip.Adjustment)
		}
	}
	return Phase(self.Year()+5, YearSeason, godip.Movement)
}
