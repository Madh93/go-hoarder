package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Madh93/go-hoarder"
)

func main() {
	// Basic configuration
	apiUrl := "https://<YOUR_HOARDER_HOSTNAME>/api/v1" // Replace this with your API URL
	apiKey := "<YOUR_HOARDER_API_KEY>"                 // Replace this with your actual token

	// Set up Bearer authentication
	auth := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		return nil
	}

	// Create the Hoarder client
	client, err := hoarder.NewClient(apiUrl, hoarder.WithRequestEditorFn(auth))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	log.Printf("Hello world from %s", client.Server)
}
