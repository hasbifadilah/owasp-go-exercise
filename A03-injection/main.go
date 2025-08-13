package main

import (
	"fmt"
	"log"
	"net/http"
)

// Simulate building SQL from user input (vulnerable to injection)
func searchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", q)
	fmt.Fprintf(w, "Would execute: %s\n", query)
}

func main() {
	http.HandleFunc("/search", searchHandler)
	log.Println("Listening on :8082 (A03 Injection)")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
