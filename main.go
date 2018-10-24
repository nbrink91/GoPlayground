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
	router.HandleFunc("/", GetMetricsEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetMetricsEndpoint(w http.ResponseWriter, r *http.Request) {
	var transcript Transcript
	err := json.NewDecoder(r.Body).Decode(&transcript)

	if err != nil {
		panic(err)
	}

	metrics := GetMetrics(transcript)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(metrics)
}
