package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zond/godip/variants"

	gae "common"
)

func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func preflight(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
}

func resolve(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	p := &gae.Phase{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	state, err := p.State(variant)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err = state.Next(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Load the new godip phase from the state
	nextPhase := gae.NewPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(nextPhase); err != nil {
		return
	}
	return
}

func start(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	state, err := variant.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	phase := gae.NewPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(phase); err != nil {
		return
	}
	return
}

func listVariants(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(variants.Variants); err != nil {
		return
	}
	return
}

func init() {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(preflight)
	variants := r.Path("/{variant}").Subrouter()
	variants.Methods("POST").HandlerFunc(resolve)
	variants.Methods("GET").HandlerFunc(start)
	r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
}
