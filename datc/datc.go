package datc

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/zond/godip/common"
)

var clearCommentReg = regexp.MustCompile("(?m)^\\s*([^#\n\t]+?)\\s*(#.*)?$")

var variantReg = regexp.MustCompile("^VARIANT_ALL\\s+(\\S*)\\s*$")
var caseReg = regexp.MustCompile("^CASE\\s+(.*)$")

var prestateSetPhaseReg = regexp.MustCompile("^PRESTATE_SETPHASE\\s+(\\S+)\\s+(\\d+),\\s+(\\S+)\\s*$")

var stateReg = regexp.MustCompile("^([^:\\s]+):?\\s+(\\S+)\\s+(\\S+)\\s*$")

var ordersReg = regexp.MustCompile("^([^:]+):\\s+(.*)$")

var preOrderReg = regexp.MustCompile("^(SUCCESS|FAILURE):\\s+([^:]+):\\s+(.*)$")

const (
	prestate                   = "PRESTATE"
	orders                     = "ORDERS"
	poststateSame              = "POSTSTATE_SAME"
	end                        = "END"
	poststate                  = "POSTSTATE"
	poststateDislodged         = "POSTSTATE_DISLODGED"
	prestateSupplycenterOwners = "PRESTATE_SUPPLYCENTER_OWNERS"
	prestateDislodged          = "PRESTATE_DISLODGED"
	prestateResults            = "PRESTATE_RESULTS"
	success                    = "SUCCESS"
	failure                    = "FAILURE"
)

func newState() *State {
	return &State{
		SCs:              make(map[common.Province]common.Nation),
		Units:            make(map[common.Province]common.Unit),
		Dislodgeds:       make(map[common.Province]common.Unit),
		Orders:           make(map[common.Province]NationalizedOrder),
		FailedOrders:     make(map[common.Province]NationalizedOrder),
		SuccessfulOrders: make(map[common.Province]NationalizedOrder),
	}
}

type NationalizedOrder struct {
	Order  common.Adjudicator
	Nation common.Nation
}

func (self NationalizedOrder) String() string {
	return fmt.Sprintf("'%v: %v'", self.Nation, self.Order)
}

type State struct {
	SCs              map[common.Province]common.Nation
	Units            map[common.Province]common.Unit
	Dislodgeds       map[common.Province]common.Unit
	Orders           map[common.Province]NationalizedOrder
	FailedOrders     map[common.Province]NationalizedOrder
	SuccessfulOrders map[common.Province]NationalizedOrder
	Phase            common.Phase
}

func (self *State) copyFrom(o *State) {
	for prov, unit := range o.Units {
		self.Units[prov] = unit
	}
	for prov, dislodged := range o.Dislodgeds {
		self.Dislodgeds[prov] = dislodged
	}
	for prov, nation := range o.SCs {
		self.SCs[prov] = nation
	}
}

func newStatePair() *StatePair {
	return &StatePair{
		Before: newState(),
		After:  newState(),
	}
}

type StatePair struct {
	Case   string
	Before *State
	After  *State
}

func (self *StatePair) copyBeforeToAfter() {
	self.After.copyFrom(self.Before)
}

type StatePairHandler func(states *StatePair)

type OrderParser func(text string) (common.Province, common.Adjudicator, error)

type PhaseParser func(season string, year int, typ string) (common.Phase, error)

type NationParser func(nation string) (common.Nation, error)

type UnitTypeParser func(typ string) (common.UnitType, error)

type ProvinceParser func(prov string) (common.Province, error)

type Parser struct {
	Variant        string
	OrderParser    OrderParser
	PhaseParser    PhaseParser
	NationParser   NationParser
	UnitTypeParser UnitTypeParser
	ProvinceParser ProvinceParser
}

const (
	waiting = iota
	inCase
	inPrestate
	inOrders
	inPoststate
	inPoststateDislodged
	inPrestateSupplycenterOwners
	inPrestateDislodged
	inPrestateResults
)

func (self Parser) Parse(r io.Reader, handler StatePairHandler) (err error) {
	lr := bufio.NewReader(r)
	var match []string
	state := waiting
	statePair := newStatePair()
	var line string
	for line, err = lr.ReadString('\n'); err == nil; line, err = lr.ReadString('\n') {
		if match = clearCommentReg.FindStringSubmatch(line); match != nil {
			line = strings.TrimSpace(match[1])
			switch state {
			case waiting:
				if match = variantReg.FindStringSubmatch(line); match != nil {
					if match[1] != self.Variant {
						err = fmt.Errorf("%+v only supports DATC files for %v variant", self, self.Variant)
					}
				} else if match = caseReg.FindStringSubmatch(line); match != nil {
					state = inCase
					statePair.Case = match[1]
				} else {
					err = fmt.Errorf("Unrecognized line for state waiting: %#v", line)
				}
			case inPrestateSupplycenterOwners:
				if match = stateReg.FindStringSubmatch(line); match != nil {
					var owner common.Nation
					if owner, err = self.NationParser(match[1]); err != nil {
						return
					}
					var prov common.Province
					if prov, err = self.ProvinceParser(match[3]); err != nil {
						return
					}
					statePair.Before.SCs[prov] = owner
				} else if line == prestate {
					state = inPrestate
				} else if line == orders {
					state = inOrders
				} else {
					err = fmt.Errorf("Unrecognized line for state inPrestateSupplycenterOwners: %#v", line)
				}
			case inCase:
				if match = prestateSetPhaseReg.FindStringSubmatch(line); match != nil {
					year := 0
					if year, err = strconv.Atoi(match[2]); err != nil {
						return
					}
					if statePair.Before.Phase, err = self.PhaseParser(match[1], year, match[3]); err != nil {
						return
					}
				} else if line == prestate {
					state = inPrestate
				} else if line == prestateSupplycenterOwners {
					state = inPrestateSupplycenterOwners
				} else {
					err = fmt.Errorf("Unrecognized line for state inCase: %#v", line)
					return
				}
			case inPrestate:
				if match = stateReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					if prov, err = self.ProvinceParser(match[3]); err != nil {
						return
					}
					var unit common.UnitType
					if unit, err = self.UnitTypeParser(match[2]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[1]); err != nil {
						return
					}
					statePair.Before.Units[prov] = common.Unit{
						unit,
						nation,
					}
				} else if line == prestateResults {
					state = inPrestateResults
				} else if line == orders {
					state = inOrders
				} else if line == prestateSupplycenterOwners {
					state = inPrestateSupplycenterOwners
				} else if line == prestateDislodged {
					state = inPrestateDislodged
				} else {
					err = fmt.Errorf("Unrecognized line for state inPrestate: %#v", line)
					return
				}
			case inPoststate:
				if match = stateReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					if prov, err = self.ProvinceParser(match[3]); err != nil {
						return
					}
					var unit common.UnitType
					if unit, err = self.UnitTypeParser(match[2]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[1]); err != nil {
						return
					}
					statePair.After.Units[prov] = common.Unit{
						unit,
						nation,
					}
				} else if line == end {
					handler(statePair)
					statePair = newStatePair()
					state = waiting
				} else if line == poststateDislodged {
					state = inPoststateDislodged
				} else {
					err = fmt.Errorf("Unrecognized line for state inPoststate: %#v", line)
					return
				}
			case inPrestateDislodged:
				if match = stateReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					if prov, err = self.ProvinceParser(match[3]); err != nil {
						return
					}
					var unit common.UnitType
					if unit, err = self.UnitTypeParser(match[2]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[1]); err != nil {
						return
					}
					statePair.Before.Dislodgeds[prov] = common.Unit{
						unit,
						nation,
					}
				} else if line == orders {
					state = inOrders
				} else if line == prestateResults {
					state = inPrestateResults
				} else {
					err = fmt.Errorf("Unrecognized line for state inPrestateDislodged: %#v", line)
					return
				}
			case inPrestateResults:
				if match = preOrderReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					var order common.Adjudicator
					if prov, order, err = self.OrderParser(match[3]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[2]); err != nil {
						return
					}
					nOrder := NationalizedOrder{
						Order:  order,
						Nation: nation,
					}
					if match[1] == success {
						statePair.Before.SuccessfulOrders[prov] = nOrder
					} else if match[1] == failure {
						statePair.Before.FailedOrders[prov] = nOrder
					} else {
						err = fmt.Errorf("Unrecognized state for pre order: %#v", match[1])
						return
					}
				} else if line == orders {
					state = inOrders
				} else {
					err = fmt.Errorf("Unrecognized line for state inPrestateResult: %#v", line)
					return
				}
			case inPoststateDislodged:
				if match = stateReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					if prov, err = self.ProvinceParser(match[3]); err != nil {
						return
					}
					var unit common.UnitType
					if unit, err = self.UnitTypeParser(match[2]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[1]); err != nil {
						return
					}
					statePair.After.Dislodgeds[prov] = common.Unit{
						unit,
						nation,
					}
				} else if line == end {
					handler(statePair)
					statePair = newStatePair()
					state = waiting
				} else {
					err = fmt.Errorf("Unrecognized line for state inPoststateDislodged: %#v", line)
					return
				}
			case inOrders:
				if match = ordersReg.FindStringSubmatch(line); match != nil {
					var prov common.Province
					var order common.Adjudicator
					if prov, order, err = self.OrderParser(match[2]); err != nil {
						return
					}
					var nation common.Nation
					if nation, err = self.NationParser(match[1]); err != nil {
						return
					}
					statePair.Before.Orders[prov] = NationalizedOrder{
						Order:  order,
						Nation: nation,
					}
				} else if line == poststateSame {
					statePair.copyBeforeToAfter()
				} else if line == poststate {
					state = inPoststate
				} else if line == end {
					handler(statePair)
					statePair = newStatePair()
					state = waiting
				} else {
					err = fmt.Errorf("Unrecognized line for state inOrders: %#v", line)
					return
				}
			default:
				err = fmt.Errorf("Unknown state %v", state)
				return
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
