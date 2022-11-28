package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	var quotes [2]string
	quotes[0] = "First quote."
	quotes[1] = "Second quote."

	min := 0
	max := 2
	fmt.Println(rand.Intn(max-min) + min)

	randomIndex := rand.Intn(max-min) + min

	fmt.Fprintf(w, quotes[randomIndex])
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
