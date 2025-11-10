package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type hello struct {
	Hello string `json:"hello"`
}

func getResponse(w http.ResponseWriter, r *http.Request) {
	var data = hello{
		Hello: "world",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {
	http.Handle("/", http.HandlerFunc(getResponse))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
