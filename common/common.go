package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

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

type OrderType string

type PhaseType string

type Province string

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
	Prev() Phase
	PostProcess(State) error
	DefaultOrder(Province) Adjudicator
	Options(Validator, Nation) (result Options)
	Winner(Validator) *Nation
}

type PathFilter func(n Province, edgeFlags, provFlags map[Flag]bool, sc *Nation) bool

type Flag string

type Graph interface {
	Has(Province) bool
	Flags(Province) map[Flag]bool
	AllFlags(Province) map[Flag]bool
	SC(Province) *Nation
	Path(src, dst Province, filter PathFilter) []Province
	Coasts(Province) []Province
	Edges(src Province) map[Province]map[Flag]bool
	SCs(Nation) []Province
	Provinces() []Province
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

type Order interface {
	Type() OrderType
	DisplayType() OrderType
	Targets() []Province
	Validate(Validator) error
	Options(Validator, Nation, Province) Options
	At() time.Time
	Flags() map[Flag]bool
}

type Adjudicator interface {
	Order
	Adjudicate(Resolver) error
	Execute(State)
}

type BackupRule func(State, []Province) error

type StateFilter func(n Province, o Order, u *Unit) bool

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
	GetProfile() map[string]time.Duration
}

type Resolver interface {
	Validator

	AddBounce(src, dst Province)
	Resolve(Province) error
}

type State interface {
	Resolver

	Move(src, dst Province, preventRetreat bool)
	Retreat(src, dst Province) error

	RemoveDislodged(Province)
	RemoveUnit(Province)

	SetResolution(Province, error)
	SetSC(Province, Nation)
	SetUnit(Province, Unit) error

	ClearDislodgers()
	ClearBounces()
}
