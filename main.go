package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func getRandomQuote() string {
	quotes := [3]string{
		"First quote.",
		"Second quote.",
		"Third quote.",
	}

	min := 0
	max := len(quotes)
	index := rand.Intn(max-min) + min

	return quotes[index]
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]string)
	response["quote"] = getRandomQuote()

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResponse)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
