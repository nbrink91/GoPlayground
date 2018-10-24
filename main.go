package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetMetrics).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Metrics struct {
	UtteranceCount int16 `json:"utteranceCount,omitempty"`
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metric := Metrics{UtteranceCount: 10}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(metric)
}
