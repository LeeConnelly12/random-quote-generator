package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var quotes [2]string
	quotes[0] = "First quote."
	quotes[1] = "Second quote."

	min := 0
	max := 2
	randomIndex := rand.Intn(max-min) + min

	resp := make(map[string]string)
	resp["quote"] = quotes[randomIndex]

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
