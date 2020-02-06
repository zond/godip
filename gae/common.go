package main

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

type Phase struct {
	Season        godip.Season
	Year          int
	Type          godip.PhaseType
	Units         map[godip.Province]godip.Unit
	Orders        map[godip.Nation]map[godip.Province][]string
	SupplyCenters map[godip.Province]godip.Nation
	Dislodgeds    map[godip.Province]godip.Unit
	Dislodgers    map[godip.Province]godip.Province
	Bounces       map[godip.Province]map[godip.Province]bool
	Resolutions   map[godip.Province]string
}

func NewPhase(state *state.State) *Phase {
	currentPhase := state.Phase()
	p := &Phase{
		Orders:      map[godip.Nation]map[godip.Province][]string{},
		Resolutions: map[godip.Province]string{},
		Season:      currentPhase.Season(),
		Year:        currentPhase.Year(),
		Type:        currentPhase.Type(),
	}
	var resolutions map[godip.Province]error
	p.Units, p.SupplyCenters, p.Dislodgeds, p.Dislodgers, p.Bounces, resolutions = state.Dump()
	for prov, err := range resolutions {
		if err == nil {
			p.Resolutions[prov] = "OK"
		} else {
			p.Resolutions[prov] = err.Error()
		}
	}
	return p
}

func (self *Phase) State(variant common.Variant) (*state.State, error) {
	parsedOrders, err := classical.Parser.ParseAll(self.Orders)
	if err != nil {
		return nil, err
	}
	return classical.Blank(variant.Phase(
		self.Year,
		self.Season,
		self.Type,
	)).Load(
		self.Units,
		self.SupplyCenters,
		self.Dislodgeds,
		self.Dislodgers,
		self.Bounces,
		parsedOrders,
	), nil
}
