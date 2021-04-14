package godip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
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

	ViaConvoy     Flag = "ViaConvoy"
	Anywhere      Flag = "Anywhere"
	AnyHomeCenter Flag = "AnyHomeCenter"
)

var (
	Coast       = []Flag{Sea, Land}
	Archipelago = []Flag{Sea, Land, Convoyable}

	// Invalid is not understood
	// Illegal is understood but not allowed
	ErrInvalidSource                   = fmt.Errorf("ErrInvalidSource")
	ErrInvalidDestination              = fmt.Errorf("ErrInvalidDestination")
	ErrInvalidTarget                   = fmt.Errorf("ErrInvalidTarget")
	ErrInvalidPhase                    = fmt.Errorf("ErrInvalidPhase")
	ErrMissingUnit                     = fmt.Errorf("ErrMissingUnit")
	ErrIllegalDestination              = fmt.Errorf("ErrIllegalDestination")
	ErrMissingConvoyPath               = fmt.Errorf("ErrMissingConvoyPath")
	ErrIllegalMove                     = fmt.Errorf("ErrIllegalMove")
	ErrConvoyParadox                   = fmt.Errorf("ErrConvoyParadox")
	ErrIllegalSupportPosition          = fmt.Errorf("ErrIllegalSupportPosition")
	ErrIllegalSupportDestination       = fmt.Errorf("ErrIllegalSupportDestination")
	ErrIllegalSupportDestinationNation = fmt.Errorf("ErrIllegalSupportDestinationNation")
	ErrMissingSupportUnit              = fmt.Errorf("ErrMissingSupportUnit")
	ErrIllegalSupportMove              = fmt.Errorf("ErrIllegalSupportMove")
	ErrInvalidSupporteeOrder           = fmt.Errorf("ErrInvalidSupporteeOrder")
	ErrIllegalConvoyUnit               = fmt.Errorf("ErrIllegalConvoyUnit")
	ErrIllegalConvoyPath               = fmt.Errorf("ErrIllegalConvoyPath")
	ErrIllegalConvoyMove               = fmt.Errorf("ErrIllegalConvoyMove")
	ErrMissingConvoyee                 = fmt.Errorf("ErrMissingConvoyee")
	ErrIllegalConvoyer                 = fmt.Errorf("ErrIllegalConvoyer")
	ErrIllegalConvoyee                 = fmt.Errorf("ErrIllegalConvoyee")
	ErrIllegalBuild                    = fmt.Errorf("ErrIllegalBuild")
	ErrIllegalDisband                  = fmt.Errorf("ErrIllegalDisband")
	ErrOccupiedSupplyCenter            = fmt.Errorf("ErrOccupiedSupplyCenter")
	ErrMissingSupplyCenter             = fmt.Errorf("ErrMissingSupplyCenter")
	ErrMissingSurplus                  = fmt.Errorf("ErrMissingSurplus")
	ErrIllegalUnitType                 = fmt.Errorf("ErrIllegalUnitType")
	ErrMissingDeficit                  = fmt.Errorf("ErrMissingDeficit")
	ErrOccupiedDestination             = fmt.Errorf("ErrOccupiedDestination")
	ErrIllegalRetreat                  = fmt.Errorf("ErrIllegalRetreat")
	ErrHostileSupplyCenter             = fmt.Errorf("ErrHostileSupplyCenter")
	InconsistencyMissingOrder          = fmt.Errorf("InconsistencyMissingOrder")
)

type InconsistencyMismatchedSupporter struct {
	Supportee Province
}

func (self InconsistencyMismatchedSupporter) Error() string {
	return fmt.Sprintf("InconsistencyMismatchedSupporter:%v", self.Supportee)
}

type InconsistencyMismatchedConvoyee struct {
	Convoyer Province
}

func (self InconsistencyMismatchedConvoyee) Error() string {
	return fmt.Sprintf("InconsistencyMismatchedConvoyee:%v", self.Convoyer)
}

type InconsistencyMismatchedConvoyer struct {
	Convoyee Province
}

func (self InconsistencyMismatchedConvoyer) Error() string {
	return fmt.Sprintf("InconsistencyMismatchedConvoyer:%v", self.Convoyee)
}

type InconsistencyOrderTypeCount struct {
	OrderType OrderType
	Found     int
	Want      int
}

func (self InconsistencyOrderTypeCount) Error() string {
	return fmt.Sprintf("InconsistencyOrderTypeCount:%v:Found:%v:Want:%v", self.OrderType, self.Found, self.Want)
}

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
var logMutex = sync.RWMutex{}

func Indent(s string) {
	if Debug {
		logMutex.Lock()
		defer logMutex.Unlock()
		LogIndent = append(LogIndent, s)
	}
}

func DeIndent() {
	if Debug {
		logMutex.Lock()
		defer logMutex.Unlock()
		LogIndent = LogIndent[:len(LogIndent)-1]
	}
}

func Logf(s string, o ...interface{}) {
	if Debug {
		logMutex.Lock()
		defer logMutex.Unlock()
		fmt.Fprintf(logBuffer, fmt.Sprintf("%v%v\n", strings.Join(LogIndent, ""), s), o...)
	}
}

func ClearLog() {
	if Debug {
		logMutex.Lock()
		defer logMutex.Unlock()
		logBuffer = new(bytes.Buffer)
	}
}

func DumpLog() {
	if Debug {
		logMutex.RLock()
		fmt.Print(string(logBuffer.Bytes()))
		logMutex.RUnlock()
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

func (n Nations) Equal(o Nations) bool {
	if len(n) != len(o) {
		return false
	}
	sortedN := make(Nations, len(n))
	copy(sortedN, n)
	sort.Sort(sortedN)
	sortedO := make(Nations, len(o))
	copy(sortedO, o)
	sort.Sort(sortedO)
	for idx, nNat := range sortedN {
		if sortedO[idx] != nNat {
			return false
		}
	}
	return true
}

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
	Options(Validator, Nation) Options
	Messages(Validator, Nation) []string
	Corroborate(Validator, Nation) []Inconsistency
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
	Edges(src Province, reverse bool) map[Province]map[Flag]bool
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

type FilteredOptionValue struct {
	Filter string
	Value  OptionValue
}

/*
Options defines a tree of valid orders for a given situation
*/
type Options map[OptionValue]Options

func (self Options) BubbleFilters() Options {
	_, result := self.bubbleFiltersHelper(false)
	return result
}

func (self Options) bubbleFiltersHelper(bubbleSelf bool) (string, Options) {
	lastFilter := ""
	filters := map[string]bool{}
	bubbledChildren := Options{}
	for k, v := range self {
		if filtered, ok := k.(FilteredOptionValue); ok {
			lastFilter = filtered.Filter
			filters[lastFilter] = true
			bubbledChildren[k] = v
		} else {
			childCommonFilter, newChild := v.bubbleFiltersHelper(true)
			lastFilter = childCommonFilter
			filters[lastFilter] = true
			if childCommonFilter == "" {
				bubbledChildren[k] = v
			} else {
				bubbledChildren[FilteredOptionValue{
					Filter: childCommonFilter,
					Value:  k,
				}] = newChild
			}
		}
	}
	if !bubbleSelf || lastFilter == "" || len(filters) > 1 {
		return "", bubbledChildren
	}
	bubbledSelf := Options{}
	for k, v := range bubbledChildren {
		filtered := k.(FilteredOptionValue)
		bubbledSelf[filtered.Value] = v
	}
	return lastFilter, bubbledSelf
}

func (self Options) MarshalJSON() ([]byte, error) {
	repl := map[string]interface{}{}
	for k, v := range self.BubbleFilters() {
		kVal := reflect.ValueOf(k)
		filter := ""
		if kVal.Type() == reflect.TypeOf(FilteredOptionValue{}) {
			filter = kVal.FieldByName("Filter").String()
			kVal = reflect.ValueOf(kVal.FieldByName("Value").Interface())
		}
		val := map[string]interface{}{
			"Type": kVal.Type().Name(),
			"Next": v,
		}
		if filter != "" {
			val["Filter"] = filter
		}
		repl[fmt.Sprint(kVal.Interface())] = val
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
	Corroborate(Validator) []error
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
	ForceDisbands() map[Province]bool
	SupplyCenters() map[Province]Nation

	Graph() Graph
	Phase() Phase
	Find(StateFilter) (provinces []Province, orders []Order, units []*Unit)

	Options([]Order, Nation) (result Options)

	Profile(string, time.Time)
	GetProfile() (map[string]time.Duration, map[string]int)

	MemoizeProvSlice(string, func() []Province) []Province

	Flags() map[Flag]bool
}

// Resolver is what validators turn into when adjudication has started.
type Resolver interface {
	Validator

	AddBounce(src, dst Province)
	Resolve(Province) error
}

type Inconsistency struct {
	Province Province
	Errors   []error
}

func (i Inconsistency) MarshalJSON() ([]byte, error) {
	errStrings := make([]string, len(i.Errors))
	for idx := range i.Errors {
		errStrings[idx] = i.Errors[idx].Error()
	}
	return json.Marshal(map[string]interface{}{
		"Province":        i.Province,
		"Inconsistencies": errStrings,
	})
}

// State is the super-user access to the entire game state.
type State interface {
	Resolver

	Move(src, dst Province, preventRetreat bool)
	Retreat(src, dst Province) error

	RemoveDislodged(Province)
	RemoveUnit(Province)
	ForceDisband(Province)

	SetResolution(Province, error)
	SetSC(Province, Nation)
	SetUnit(Province, Unit) error
	SetOrders(map[Province]Adjudicator)

	ClearDislodgers()
	ClearBounces()

	Corroborate(Nation) []Inconsistency
}
