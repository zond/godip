package godip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	Sea        Flag = "Sea"
	Land       Flag = "Land"
	Convoyable Flag = "Convoyable"

	Army  UnitType = "Army"
	Fleet UnitType = "Fleet"

	England Nation = "England"
	France  Nation = "France"
	Germany Nation = "Germany"
	Russia  Nation = "Russia"
	Austria Nation = "Austria"
	Italy   Nation = "Italy"
	Turkey  Nation = "Turkey"
	Neutral Nation = "Neutral"

	Spring Season = "Spring"
	Fall   Season = "Fall"

	Movement   PhaseType = "Movement"
	Retreat    PhaseType = "Retreat"
	Adjustment PhaseType = "Adjustment"

	Build         OrderType = "Build"
	Move          OrderType = "Move"
	MoveViaConvoy OrderType = "MoveViaConvoy"
	Hold          OrderType = "Hold"
	Convoy        OrderType = "Convoy"
	Support       OrderType = "Support"
	Disband       OrderType = "Disband"

	ViaConvoy     Flag = "C"
	Anywhere      Flag = "A"
	AnyHomeCenter Flag = "H"
)

var Coast = []Flag{Sea, Land}
var Archipelago = []Flag{Sea, Land, Convoyable}

// Invalid is not understood
// Illegal is understood but not allowed
var ErrInvalidSource = fmt.Errorf("ErrInvalidSource")
var ErrInvalidDestination = fmt.Errorf("ErrInvalidDestination")
var ErrInvalidTarget = fmt.Errorf("ErrInvalidTarget")
var ErrInvalidPhase = fmt.Errorf("ErrInvalidPhase")
var ErrMissingUnit = fmt.Errorf("ErrMissingUnit")
var ErrIllegalDestination = fmt.Errorf("ErrIllegalDestination")
var ErrMissingConvoyPath = fmt.Errorf("ErrMissingConvoyPath")
var ErrIllegalMove = fmt.Errorf("ErrIllegalMove")
var ErrConvoyParadox = fmt.Errorf("ErrConvoyParadox")
var ErrIllegalSupportPosition = fmt.Errorf("ErrIllegalSupportPosition")
var ErrIllegalSupportDestination = fmt.Errorf("ErrIllegalSupportDestination")
var ErrIllegalSupportDestinationNation = fmt.Errorf("ErrIllegalSupportDestinationNation")
var ErrMissingSupportUnit = fmt.Errorf("ErrMissingSupportUnit")
var ErrIllegalSupportMove = fmt.Errorf("ErrIllegalSupportMove")
var ErrIllegalConvoyUnit = fmt.Errorf("ErrIllegalConvoyUnit")
var ErrIllegalConvoyPath = fmt.Errorf("ErrIllegalConvoyPath")
var ErrIllegalConvoyMove = fmt.Errorf("ErrIllegalConvoyMove")
var ErrMissingConvoyee = fmt.Errorf("ErrMissingConvoyee")
var ErrIllegalConvoyer = fmt.Errorf("ErrIllegalConvoyer")
var ErrIllegalConvoyee = fmt.Errorf("ErrIllegalConvoyee")
var ErrIllegalBuild = fmt.Errorf("ErrIllegalBuild")
var ErrIllegalDisband = fmt.Errorf("ErrIllegalDisband")
var ErrOccupiedSupplyCenter = fmt.Errorf("ErrOccupiedSupplyCenter")
var ErrMissingSupplyCenter = fmt.Errorf("ErrMissingSupplyCenter")
var ErrMissingSurplus = fmt.Errorf("ErrMissingSurplus")
var ErrIllegalUnitType = fmt.Errorf("ErrIllegalUnitType")
var ErrMissingDeficit = fmt.Errorf("ErrMissingDeficit")
var ErrOccupiedDestination = fmt.Errorf("ErrOccupiedDestination")
var ErrIllegalRetreat = fmt.Errorf("ErrIllegalRetreat")
var ErrForcedDisband = fmt.Errorf("ErrForcedDisband")
var ErrHostileSupplyCenter = fmt.Errorf("ErrHostileSupplyCenter")

type ErrDoubleBuild struct {
	Provinces []Province
}

func (self ErrDoubleBuild) Error() string {
	return fmt.Sprintf("ErrDoubleBuild:%v", self.Provinces)
}

type ErrConvoyDislodged struct {
	Province Province
}

func (self ErrConvoyDislodged) Error() string {
	return fmt.Sprintf("ErrConvoyDislodged:%v", self.Province)
}

type ErrSupportBroken struct {
	Province Province
}

func (self ErrSupportBroken) Error() string {
	return fmt.Sprintf("ErrSupportBroken:%v", self.Province)
}

type ErrBounce struct {
	Province Province
}

func (self ErrBounce) Error() string {
	return fmt.Sprintf("ErrBounce:%v", self.Province)
}

var Debug = false
var LogIndent = []string{}
var logBuffer = new(bytes.Buffer)

func Indent(s string) {
	if Debug {
		LogIndent = append(LogIndent, s)
	}
}

func DeIndent() {
	if Debug {
		LogIndent = LogIndent[:len(LogIndent)-1]
	}
}

func Logf(s string, o ...interface{}) {
	if Debug {
		fmt.Fprintf(logBuffer, fmt.Sprintf("%v%v\n", strings.Join(LogIndent, ""), s), o...)
	}
}

func ClearLog() {
	if Debug {
		logBuffer = new(bytes.Buffer)
	}
}

func DumpLog() {
	if Debug {
		fmt.Print(string(logBuffer.Bytes()))
		ClearLog()
	}
}

func Max(is ...int) (result int) {
	for index, i := range is {
		if index == 0 || i > result {
			result = i
		}
	}
	return
}

func Min(is ...int) (result int) {
	for index, i := range is {
		if index == 0 || i < result {
			result = i
		}
	}
	return
}

type UnitType string

type Nation string

func (n *Nation) String() string {
	if n == nil {
		return ""
	}
	return string(*n)
}

type Nations []Nation

func (n Nations) Len() int {
	return len(n)
}

func (n Nations) Less(i, j int) bool {
	return bytes.Compare([]byte(n[i]), []byte(n[j])) < 0
}

func (n Nations) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type OrderType string

type PhaseType string

type Province string

func (p *Province) String() string {
	if p == nil {
		return ""
	}
	return string(*p)
}

type Season string

func (self Province) Split() (sup Province, sub Province) {
	split := strings.Split(string(self), "/")
	if len(split) > 0 {
		sup = Province(split[0])
	}
	if len(split) > 1 {
		sub = Province(split[1])
	}
	return
}

func (self Province) Join(n Province) (result Province) {
	if n != "" {
		result = Province(fmt.Sprintf("%v/%v", self, n))
	} else {
		result = self
	}
	return
}

func (self Province) Super() (result Province) {
	result, _ = self.Split()
	return
}

func (self Province) Sub() (result Province) {
	_, result = self.Split()
	return
}

func (self Province) Contains(p Province) bool {
	return self == p || (self.Super() == self && self == p.Super())
}

type Unit struct {
	Type   UnitType
	Nation Nation
}

func (self Unit) Equal(o Unit) bool {
	return self.Type == o.Type && self.Nation == o.Nation
}

func (self *Unit) String() string {
	return fmt.Sprint(*self)
}

type Phase interface {
	Year() int
	Season() Season
	Type() PhaseType
	Next() Phase
	PreProcess(State) error
	PostProcess(State) error
	DefaultOrder(Province) Adjudicator
	Options(Validator, Nation) (result Options)
}

type PathFilter func(n Province, edgeFlags, provFlags map[Flag]bool, sc *Nation, trace []Province) bool

type Flag string

type Graph interface {
	Has(Province) bool
	Flags(Province) map[Flag]bool
	AllFlags(Province) map[Flag]bool
	SC(Province) *Nation
	Path(src, dst Province, reverse bool, filter PathFilter) []Province
	Coasts(Province) []Province
	Edges(src Province) map[Province]map[Flag]bool
	ReverseEdges(src Province) map[Province]map[Flag]bool
	SCs(Nation) []Province
	AllSCs() []Province
	Provinces() []Province
	Nations() []Nation
}

type Orders []Order

func (self Orders) Less(a, b int) bool {
	return self[a].At().Before(self[b].At())
}

func (self Orders) Swap(a, b int) {
	self[a], self[b] = self[b], self[a]
}

func (self Orders) Len() int {
	return len(self)
}

type SrcProvince Province

type OptionValue interface{}

/*
Options defines a tree of valid orders for a given situation
*/
type Options map[OptionValue]Options

func (self Options) MarshalJSON() ([]byte, error) {
	repl := map[string]interface{}{}
	for k, v := range self {
		repl[fmt.Sprint(k)] = map[string]interface{}{
			"Type": reflect.ValueOf(k).Type().Name(),
			"Next": v,
		}
	}
	return json.Marshal(repl)
}

// Order is a basic order, but unable to adjudicate itself.
type Order interface {
	Type() OrderType
	DisplayType() OrderType
	Targets() []Province
	Parse([]string) (Adjudicator, error)
	Validate(Validator) (Nation, error)
	Options(Validator, Nation, Province) Options
	At() time.Time
	Flags() map[Flag]bool
}

// Adjudicator is what orders turn into when adjudication has started.
type Adjudicator interface {
	Order
	Adjudicate(Resolver) error
	Execute(State)
}

type BackupRule func(State, []Province) error

type StateFilter func(n Province, o Order, u *Unit) bool

// Validator is a game state able to validate orders, but not adjudicate them.
type Validator interface {
	Order(Province) (Order, Province, bool)
	Unit(Province) (Unit, Province, bool)
	Dislodged(Province) (Unit, Province, bool)
	SupplyCenter(Province) (Nation, Province, bool)
	Bounce(src, dst Province) bool

	Orders() map[Province]Adjudicator
	Units() map[Province]Unit
	Dislodgeds() map[Province]Unit
	SupplyCenters() map[Province]Nation

	Graph() Graph
	Phase() Phase
	Find(StateFilter) (provinces []Province, orders []Order, units []*Unit)

	Options([]Order, Nation) (result Options)

	Profile(string, time.Time)
	GetProfile() (map[string]time.Duration, map[string]int)

	MemoizeProvSlice(string, func() []Province) []Province
}

// Resolver is what validators turn into when adjudication has started.
type Resolver interface {
	Validator

	AddBounce(src, dst Province)
	Resolve(Province) error
}

// State is the super-user access to the entire game state.
type State interface {
	Resolver

	Move(src, dst Province, preventRetreat bool)
	Retreat(src, dst Province) error

	RemoveDislodged(Province)
	RemoveUnit(Province)

	SetResolution(Province, error)
	SetSC(Province, Nation)
	SetUnit(Province, Unit) error
	SetOrders(map[Province]Adjudicator)

	ClearDislodgers()
	ClearBounces()
}
