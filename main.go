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
	router.HandleFunc("/", GetMetricsEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetMetricsEndpoint(w http.ResponseWriter, r *http.Request) {
	metrics := GetMetrics()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(metrics)
}
