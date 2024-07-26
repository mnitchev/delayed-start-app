package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Waiting 13 minutes before starting server...")
	time.Sleep(13 * time.Minute)

	fmt.Println("Starting server...")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Got request on %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
