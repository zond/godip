package testing

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
)

func init() {
	godip.Debug = true
}

var (
	gameFileReg = regexp.MustCompile("^game_\\d+\\.txt$")

	phaseReg = regexp.MustCompile("^PHASE (\\d+) (\\S+) (\\S+)$")
	posReg   = regexp.MustCompile("^(.+): (fleet|army|supply|fleet/dislodged|army/dislodged) (\\S+)$")

	moveReg          = regexp.MustCompile("^(\\S+)\\s+move\\s+(\\S+)$")
	moveViaConvoyReg = regexp.MustCompile("^(\\S+)\\s+move\\s+(\\S+)\\s+via\\s+convoy$")
	supportMoveReg   = regexp.MustCompile("^(\\S+)\\s+support\\s+(\\S+)\\s+move\\s+(\\S+)$")
	supportHoldReg   = regexp.MustCompile("^(\\S+)\\s+support\\s+(\\S+)$")
	holdReg          = regexp.MustCompile("^(\\S+)\\s+hold$")
	convoyReg        = regexp.MustCompile("^(\\S+)\\s+convoy\\s+(\\S+)\\s+move\\s+(\\S+)$")
	buildReg         = regexp.MustCompile("^build\\s+(Army|Fleet)\\s+(\\S+)$")
	buildAnywhereReg = regexp.MustCompile("^build\\s+anywhere\\s+(Army|Fleet)\\s+(\\S+)$")
	removeReg        = regexp.MustCompile("^remove\\s+(\\S+)$")
	disbandReg       = regexp.MustCompile("^(\\S+)\\s+disband$")

	optionsCalculated           int64
	timeSpentCalculatingOptions time.Duration
	worstOptionsCalculation     time.Duration
)

const (
	positionsTag = "POSITIONS"
	ordersTag    = "ORDERS"
)

const (
	inNothing = iota
	inPositions
	inOrders
)

func verifyReversePositions(t *testing.T, s *state.State, scCollector map[godip.Province]godip.Nation, unitCollector, dislodgedCollector map[godip.Province]godip.Unit, fails *int) {
	for prov, nation1 := range s.SupplyCenters() {
		if nation2, ok := scCollector[prov]; !ok || nation1 != nation2 {
			t.Errorf("%v: Found %v in %v, expected %v, %v", s.Phase(), nation1, prov, nation2, ok)
			*fails += 1
		}
	}
	for prov, unit1 := range s.Units() {
		if unit2, ok := unitCollector[prov]; !ok || unit2.Nation != unit1.Nation || unit1.Type != unit2.Type {
			t.Errorf("%v: Found %v in %v, expected %v, %v", s.Phase(), unit1, prov, unit2, ok)
			*fails += 1
		}
	}
	for prov, unit1 := range s.Dislodgeds() {
		if unit2, ok := dislodgedCollector[prov]; !ok || unit2.Nation != unit1.Nation || unit1.Type != unit2.Type {
			t.Errorf("%v: Found %v dislodged in %v, expected %v, %v", s.Phase(), unit1, prov, unit2, ok)
			*fails += 1
		}
	}
}

func verifyPosition(t *testing.T, s *state.State, match []string, scCollector map[godip.Province]godip.Nation, unitCollector, dislodgedCollector map[godip.Province]godip.Unit, fails *int) {
	if match[2] == "supply" {
		if nation, _, ok := s.SupplyCenter(godip.Province(match[3])); ok && nation == godip.Nation(match[1]) {
			scCollector[godip.Province(match[3])] = nation
		} else {
			t.Errorf("%v: Expected %v to own SC in %v, but found %v, %v", s.Phase(), match[1], match[3], nation, ok)
			*fails += 1
		}
	} else if match[2] == "army" {
		if unit, _, ok := s.Unit(godip.Province(match[3])); ok && unit.Nation == godip.Nation(match[1]) && unit.Type == godip.Army {
			unitCollector[godip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v in %v, but found %v, %v", s.Phase(), match[1], godip.Army, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "fleet" {
		if unit, _, ok := s.Unit(godip.Province(match[3])); ok && unit.Nation == godip.Nation(match[1]) && unit.Type == godip.Fleet {
			unitCollector[godip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v in %v, but found %v, %v", s.Phase(), match[1], godip.Fleet, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "fleet/dislodged" {
		if unit, _, ok := s.Dislodged(godip.Province(match[3])); ok && unit.Nation == godip.Nation(match[1]) && unit.Type == godip.Fleet {
			dislodgedCollector[godip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v dislodged in %v, but found %v, %v", s.Phase(), match[1], godip.Army, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "army/dislodged" {
		if unit, _, ok := s.Dislodged(godip.Province(match[3])); ok && unit.Nation == godip.Nation(match[1]) && unit.Type == godip.Army {
			dislodgedCollector[godip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v dislodged in %v, but found %v, %v", s.Phase(), match[1], godip.Army, match[3], unit, ok)
			*fails += 1
		}
	} else {
		t.Fatalf("Unknown position description %v", match)
	}
}

func setPhase(t *testing.T, sp **state.State, match []string, blankFn func(godip.Phase) *state.State) {
	year, err := strconv.Atoi(match[1])
	if err != nil {
		t.Fatalf("%v", err)
	}
	season := match[2]
	typ := match[3]
	s := *sp
	for (s.Phase().Year() <= year && (string(s.Phase().Season()) != season || string(s.Phase().Type()) != typ)) || s.Phase().Year() != year {
		s.Next()
		newS := blankFn(s.Phase())
		a, b, c, d, e, _ := s.Dump()
		newS.Load(a, b, c, d, e, map[godip.Province]godip.Adjudicator{})
		*sp = newS
	}
	if s.Phase().Year() > year {
		t.Fatalf("What the, we wanted %v but ended up with %v", match, s.Phase())
	}
}

func verifyValidOrder(t *testing.T, nat godip.Nation, v godip.Validator, order []string, parse func(bits []string) (result godip.Adjudicator, err error)) {
	order[0], order[1] = order[1], order[0]
	parsed, err := parse(order)
	if err != nil {
		t.Errorf("got unparseable order %+v: %v", order, err)
	}
	foundNat, err := parsed.Validate(v)
	if foundNat != nat {
		t.Errorf("Wanted %q, got %q", nat, foundNat)
	}
	if err != nil {
		t.Errorf("got invalid order %v: %v", parsed, err)
	}
}

func verifyValidOptions(t *testing.T, nat godip.Nation, v godip.Validator, opts godip.Options, stack []string, parse func(bits []string) (result godip.Adjudicator, err error)) {
	if len(opts) == 0 {
		verifyValidOrder(t, nat, v, stack, parse)
	}
	for nextPart, nextOptions := range opts {
		verifyValidOptions(t, nat, v, nextOptions, append(append([]string{}, stack...), fmt.Sprint(nextPart)), parse)
	}
}

func assertGame(t *testing.T, name string, nations []godip.Nation,
	startFn func() (*state.State, error), blankFn func(godip.Phase) *state.State,
	parse func(bits []string) (result godip.Adjudicator, err error)) (phases, ords, positions, fails int, s *state.State) {

	worstOptionsCalculation = 0
	file, err := os.Open(fmt.Sprintf("games/%v", name))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if s, err = startFn(); err != nil {
		t.Fatalf("%v", err)
	}
	lines := bufio.NewReader(file)
	var match []string
	state := inNothing
	scCollector, unitCollector, dislodgedCollector := make(map[godip.Province]godip.Nation), make(map[godip.Province]godip.Unit), make(map[godip.Province]godip.Unit)
	for line, err := lines.ReadString('\n'); err == nil; line, err = lines.ReadString('\n') {
		line = strings.TrimSpace(line)
		switch state {
		case inNothing:
			if os.Getenv("BENCHMARK_OPTIONS") == "true" {
				for _, nat := range nations {
					t1 := time.Now()
					options := s.Phase().Options(s, nat)
					spent := time.Now().Sub(t1)
					timeSpentCalculatingOptions += spent
					if spent > worstOptionsCalculation {
						worstOptionsCalculation = spent
					}
					optionsCalculated++
					for _, opts := range options {
						verifyValidOptions(t, nat, s, opts, nil, parse)
					}
				}
			}
			if match = phaseReg.FindStringSubmatch(line); match != nil {
				phases += 1
				setPhase(t, &s, match, blankFn)
			} else if line == positionsTag {
				state = inPositions
			} else {
				t.Fatalf("Unknown line for state inNothing: %v", line)
			}
		case inPositions:
			if match = posReg.FindStringSubmatch(line); match != nil {
				positions += 1
				verifyPosition(t, s, match, scCollector, unitCollector, dislodgedCollector, &fails)
			} else if line == ordersTag {
				verifyReversePositions(t, s, scCollector, unitCollector, dislodgedCollector, &fails)
				if fails > 0 {
					return
				}
				godip.ClearLog()
				scCollector, unitCollector, dislodgedCollector = make(map[godip.Province]godip.Nation), make(map[godip.Province]godip.Unit), make(map[godip.Province]godip.Unit)
				state = inOrders
			} else {
				t.Fatalf("Unknown line for state inPositions: %v", line)
			}
		case inOrders:
			ords += 1
			if match = moveReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Move(godip.Province(match[1]), godip.Province(match[2])))
			} else if match = moveViaConvoyReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Move(godip.Province(match[1]), godip.Province(match[2])).ViaConvoy())
			} else if match = supportMoveReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.SupportMove(godip.Province(match[1]), godip.Province(match[2]), godip.Province(match[3])))
			} else if match = supportHoldReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.SupportHold(godip.Province(match[1]), godip.Province(match[2])))
			} else if match = holdReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Hold(godip.Province(match[1])))
			} else if match = convoyReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Convoy(godip.Province(match[1]), godip.Province(match[2]), godip.Province(match[3])))
			} else if match = buildReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[2]), orders.Build(godip.Province(match[2]), godip.UnitType(match[1]), time.Now()))
			} else if match = buildAnywhereReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[2]), orders.BuildAnywhere(godip.Province(match[2]), godip.UnitType(match[1]), time.Now()))
			} else if match = removeReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Disband(godip.Province(match[1]), time.Now()))
			} else if match = disbandReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(godip.Province(match[1]), orders.Disband(godip.Province(match[1]), time.Now()))
			} else if match = phaseReg.FindStringSubmatch(line); match != nil {
				ords -= 1
				phases += 1
				setPhase(t, &s, match, blankFn)
				state = inNothing
			} else {
				t.Fatalf("Unknown line for state inOrders: %v", line)
			}
		default:
			t.Fatalf("Unknown state %v", state)
		}
	}
	return
}

func TestGames(t *testing.T, variant common.Variant) {
	gamedir, err := os.Open("games")
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer gamedir.Close()
	gamefiles, err := gamedir.Readdirnames(0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	sort.Sort(sort.StringSlice(gamefiles))
	for _, name := range gamefiles {
		if skip := os.Getenv("SKIP"); skip == "" || bytes.Compare([]byte(skip), []byte(name)) < 1 {
			if gameFileReg.MatchString(name) {
				fmt.Printf("Testing %v %v\n", variant.Name, name)
				phases, orders, positions, fails, s := assertGame(t, name, variant.Nations, variant.Start, variant.Blank, variant.Parser.Parse)
				if os.Getenv("DEBUG") == "true" {
					fmt.Printf("Checked %v phases, executed %v orders and asserted %v positions in %v, found %v failures.\n", phases, orders, positions, name, fails)
				}
				if os.Getenv("BENCHMARK_OPTIONS") == "true" {
					fmt.Printf("Spent on average %v calculating options, never more than %v.", timeSpentCalculatingOptions/time.Duration(optionsCalculated), worstOptionsCalculation)
				}
				if fails > 0 {
					godip.DumpLog()
					for prov, err := range s.Resolutions() {
						t.Errorf("%v: %v", prov, err)
					}
					t.Fatalf("%v failed!", name)
				}
			}
		}
	}
}
