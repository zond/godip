package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zond/godip"
	"github.com/zond/godip/variants"
	"google.golang.org/appengine"
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

	if _, ok := r.URL.Query()["options"]; ok {
		http.Error(w, "Query parameter 'options' is not allowed", 400)
		return
	}

	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	p := &Phase{}
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
	nextPhase := NewPhase(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(nextPhase); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func resolveWithOptions(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)

	variantName := mux.Vars(r)["variant"]
	variant, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	p := &Phase{}
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

	nextPhase := NewPhase(state)

	options := map[godip.Nation]godip.Options{}
	for _, nation := range state.Graph().Nations() {
		options[nation] = state.Phase().Options(state, nation)
	}

	response := struct {
		Phase   *Phase                         `json:"phase"`
		Options map[godip.Nation]godip.Options `json:"options"`
	}{
		Phase:   nextPhase,
		Options: options,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
	phase := NewPhase(state)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(phase); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func startWithOptions(w http.ResponseWriter, r *http.Request) {
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
	phase := NewPhase(state)

	options := map[godip.Nation]godip.Options{}
	for _, nation := range state.Graph().Nations() {
		options[nation] = state.Phase().Options(state, nation)
	}
	response := struct {
		Phase   *Phase                         `json:"phase"`
		Options map[godip.Nation]godip.Options `json:"options"`
	}{
		Phase:   phase,
		Options: options,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func listVariants(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(variants.Variants); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(preflight)
	variants := r.Path("/{variant}").Subrouter()
	variants.Methods("POST").HandlerFunc(resolve)
	variants.Path("/resolve-with-options").Methods("POST").HandlerFunc(resolveWithOptions)
	variants.Methods("GET").HandlerFunc(start)
	r.Path("/start-with-options/{variant}").Methods("GET").HandlerFunc(startWithOptions)
	r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
	appengine.Main()
}
