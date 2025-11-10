package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	apiKey := os.Getenv("W3W_API_KEY")
	apiUrl := "https://api.what3words.com/v3/autosuggest?input=film.crunchy.spiri&key=" + apiKey

	fmt.Println("API KEY: ", apiKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal("[w3w] Error connecting to API: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatal("[w3w] Error retrieving data from API: ", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("[w3w] Error reading response body: ", err)
	}

	fmt.Printf("%s", body)
}
