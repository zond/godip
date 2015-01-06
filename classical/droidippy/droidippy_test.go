package droidippy

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

	"github.com/zond/godip/classical"
	cla "github.com/zond/godip/classical/common"
	"github.com/zond/godip/classical/orders"
	dip "github.com/zond/godip/common"
	"github.com/zond/godip/state"
)

func init() {
	dip.Debug = true
}

var gameFileReg = regexp.MustCompile("^game_\\d+\\.txt$")

var phaseReg = regexp.MustCompile("^PHASE (\\d+) (\\S+) (\\S+)$")
var posReg = regexp.MustCompile("^(\\S+): (fleet|army|supply|fleet/dislodged|army/dislodged) (\\S+)$")

var moveReg = regexp.MustCompile("^(\\S+)\\s+move\\s+(\\S+)$")
var moveViaConvoyReg = regexp.MustCompile("^(\\S+)\\s+move\\s+(\\S+)\\s+via\\s+convoy$")
var supportMoveReg = regexp.MustCompile("^(\\S+)\\s+support\\s+(\\S+)\\s+move\\s+(\\S+)$")
var supportHoldReg = regexp.MustCompile("^(\\S+)\\s+support\\s+(\\S+)$")
var holdReg = regexp.MustCompile("^(\\S+)\\s+hold$")
var convoyReg = regexp.MustCompile("^(\\S+)\\s+convoy\\s+(\\S+)\\s+move\\s+(\\S+)$")
var buildReg = regexp.MustCompile("^build\\s+(Army|Fleet)\\s+(\\S+)$")
var removeReg = regexp.MustCompile("^remove\\s+(\\S+)$")
var disbandReg = regexp.MustCompile("^(\\S+)\\s+disband$")

const (
	positionsTag = "POSITIONS"
	ordersTag    = "ORDERS"
)

const (
	inNothing = iota
	inPositions
	inOrders
)

func verifyReversePositions(t *testing.T, s *state.State, scCollector map[dip.Province]dip.Nation, unitCollector, dislodgedCollector map[dip.Province]dip.Unit, fails *int) {
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

func verifyPosition(t *testing.T, s *state.State, match []string, scCollector map[dip.Province]dip.Nation, unitCollector, dislodgedCollector map[dip.Province]dip.Unit, fails *int) {
	if match[2] == "supply" {
		if nation, _, ok := s.SupplyCenter(dip.Province(match[3])); ok && nation == dip.Nation(match[1]) {
			scCollector[dip.Province(match[3])] = nation
		} else {
			t.Errorf("%v: Expected %v to own SC in %v, but found %v, %v", s.Phase(), match[1], match[3], nation, ok)
			*fails += 1
		}
	} else if match[2] == "army" {
		if unit, _, ok := s.Unit(dip.Province(match[3])); ok && unit.Nation == dip.Nation(match[1]) && unit.Type == cla.Army {
			unitCollector[dip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v in %v, but found %v, %v", s.Phase(), match[1], cla.Army, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "fleet" {
		if unit, _, ok := s.Unit(dip.Province(match[3])); ok && unit.Nation == dip.Nation(match[1]) && unit.Type == cla.Fleet {
			unitCollector[dip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v in %v, but found %v, %v", s.Phase(), match[1], cla.Fleet, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "fleet/dislodged" {
		if unit, _, ok := s.Dislodged(dip.Province(match[3])); ok && unit.Nation == dip.Nation(match[1]) && unit.Type == cla.Fleet {
			dislodgedCollector[dip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v dislodged in %v, but found %v, %v", s.Phase(), match[1], cla.Army, match[3], unit, ok)
			*fails += 1
		}
	} else if match[2] == "army/dislodged" {
		if unit, _, ok := s.Dislodged(dip.Province(match[3])); ok && unit.Nation == dip.Nation(match[1]) && unit.Type == cla.Army {
			dislodgedCollector[dip.Province(match[3])] = unit
		} else {
			t.Errorf("%v: Expected to find %v %v dislodged in %v, but found %v, %v", s.Phase(), match[1], cla.Army, match[3], unit, ok)
			*fails += 1
		}
	} else {
		t.Fatalf("Unknown position description %v", match)
	}
}

func setPhase(t *testing.T, sp **state.State, match []string) {
	year, err := strconv.Atoi(match[1])
	if err != nil {
		t.Fatalf("%v", err)
	}
	season := match[2]
	typ := match[3]
	s := *sp
	for (s.Phase().Year() <= year && (string(s.Phase().Season()) != season || string(s.Phase().Type()) != typ)) || s.Phase().Year() != year {
		s.Next()
		newS := classical.Blank(s.Phase())
		a, b, c, d, e, _ := s.Dump()
		newS.Load(a, b, c, d, e, map[dip.Province]dip.Adjudicator{})
		*sp = newS
	}
	if s.Phase().Year() > year {
		t.Fatalf("What the, we wanted %v but ended up with %v", match, s.Phase())
	}
}

func assertGame(t *testing.T, name string) (phases, ords, positions, fails int, s *state.State) {
	file, err := os.Open(fmt.Sprintf("games/%v", name))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if s, err = classical.Start(); err != nil {
		t.Fatalf("%v", err)
	}
	lines := bufio.NewReader(file)
	var match []string
	state := inNothing
	scCollector, unitCollector, dislodgedCollector := make(map[dip.Province]dip.Nation), make(map[dip.Province]dip.Unit), make(map[dip.Province]dip.Unit)
	for line, err := lines.ReadString('\n'); err == nil; line, err = lines.ReadString('\n') {
		line = strings.TrimSpace(line)
		switch state {
		case inNothing:
			if match = phaseReg.FindStringSubmatch(line); match != nil {
				phases += 1
				setPhase(t, &s, match)
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
				dip.ClearLog()
				scCollector, unitCollector, dislodgedCollector = make(map[dip.Province]dip.Nation), make(map[dip.Province]dip.Unit), make(map[dip.Province]dip.Unit)
				state = inOrders
			} else {
				t.Fatalf("Unknown line for state inPositions: %v", line)
			}
		case inOrders:
			ords += 1
			if match = moveReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Move(dip.Province(match[1]), dip.Province(match[2])))
			} else if match = moveViaConvoyReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Move(dip.Province(match[1]), dip.Province(match[2])).ViaConvoy())
			} else if match = supportMoveReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.SupportMove(dip.Province(match[1]), dip.Province(match[2]), dip.Province(match[3])))
			} else if match = supportHoldReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.SupportHold(dip.Province(match[1]), dip.Province(match[2])))
			} else if match = holdReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Hold(dip.Province(match[1])))
			} else if match = convoyReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Convoy(dip.Province(match[1]), dip.Province(match[2]), dip.Province(match[3])))
			} else if match = buildReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[2]), orders.Build(dip.Province(match[2]), dip.UnitType(match[1]), time.Now()))
			} else if match = removeReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Disband(dip.Province(match[1]), time.Now()))
			} else if match = disbandReg.FindStringSubmatch(line); match != nil {
				s.SetOrder(dip.Province(match[1]), orders.Disband(dip.Province(match[1]), time.Now()))
			} else if match = phaseReg.FindStringSubmatch(line); match != nil {
				ords -= 1
				phases += 1
				setPhase(t, &s, match)
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

func TestDroidippyGames(t *testing.T) {
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
				phases, orders, positions, fails, s := assertGame(t, name)
				fmt.Printf("Checked %v phases, executed %v orders and asserted %v positions in %v, found %v failures.\n", phases, orders, positions, name, fails)
				if fails > 0 {
					dip.DumpLog()
					for prov, err := range s.Resolutions() {
						t.Errorf("%v: %v", prov, err)
					}
					t.Fatalf("%v failed!", name)
				}
			}
		}
	}
}
