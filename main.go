package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func getRandomQuote() string {
	var quotes [3]string
	quotes[0] = "First quote."
	quotes[1] = "Second quote."
	quotes[2] = "Third quote."

	min := 0
	max := len(quotes)
	index := rand.Intn(max-min) + min

	return quotes[index]
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["quote"] = getRandomQuote()

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
