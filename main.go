package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("W3W_API_KEY")
	apiUrl := "https://api.what3words.com/v3/autosuggest?input=film.crunchy.spiri&key=" + apiKey

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("[w3w] Error connecting to API: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[w3w] Error reading response body: %s", err)
	}

	fmt.Printf("%s", body)
}
