package router

import (
	"encoding/json"
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
	nextDipPhase := state.Phase()
	// Create a diplicity phase for the new phase
	nextPhase := &phase{
		Orders:      map[dip.Nation]map[dip.Province][]string{},
		Resolutions: map[dip.Province]string{},
		Season:      nextDipPhase.Season(),
		Year:        nextDipPhase.Year(),
		Type:        nextDipPhase.Type(),
	}
	// Set the new phase positions
	var resolutions map[dip.Province]error
	nextPhase.Units, nextPhase.SupplyCenters, nextPhase.Dislodgeds, nextPhase.Dislodgers, nextPhase.Bounces, resolutions = state.Dump()
	for prov, err := range resolutions {
		if err == nil {
			nextPhase.Resolutions[prov] = "OK"
		} else {
			nextPhase.Resolutions[prov] = err.Error()
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(nextPhase); err != nil {
		return
	}
	return
}

func init() {
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(resolve)
	http.Handle("/", r)
}
