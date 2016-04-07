package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zond/godip/classical"
	"github.com/zond/godip/classical/orders"
	"github.com/zond/godip/state"

	"appengine"

	dip "github.com/zond/godip/common"
)

type phase struct {
	Season        dip.Season
	Year          int
	Type          dip.PhaseType
	Units         map[dip.Province]dip.Unit
	Orders        map[dip.Nation]map[dip.Province][]string
	SupplyCenters map[dip.Province]dip.Nation
	Dislodgeds    map[dip.Province]dip.Unit
	Dislodgers    map[dip.Province]dip.Province
	Bounces       map[dip.Province]map[dip.Province]bool
	Resolutions   map[dip.Province]string
}

func newPhase(state *state.State) *phase {
	currentPhase := state.Phase()
	p := &phase{
		Orders:      map[dip.Nation]map[dip.Province][]string{},
		Resolutions: map[dip.Province]string{},
		Season:      currentPhase.Season(),
		Year:        currentPhase.Year(),
		Type:        currentPhase.Type(),
	}
	var resolutions map[dip.Province]error
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

func (self *phase) state(c appengine.Context) (*state.State, error) {
	parsedOrders, err := orders.ParseAll(self.Orders)
	if err != nil {
		return nil, err
	}
	return classical.Blank(classical.Phase(
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

func resolve(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	p := &phase{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	state, err := p.state(c)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err = state.Next(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Load the new godip phase from the state
	nextPhase := newPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(nextPhase); err != nil {
		return
	}
	return
}

func start(w http.ResponseWriter, r *http.Request) {
	state, err := classical.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	phase := newPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(phase); err != nil {
		return
	}
	return
}

func listVariants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `	
/classical
`)
}

func init() {
	r := mux.NewRouter()
	classical := r.Path("/classical").Subrouter()
	classical.Methods("POST").HandlerFunc(resolve)
	classical.Methods("GET").HandlerFunc(start)
	r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
}
