package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// BAD: trusts a token provided in query param without verifying signature.
// Here token is expected to be base64(username) for demonstration.
func meHandler(w http.ResponseWriter, r *http.Request) {
	tok := r.URL.Query().Get("token")
	if tok == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(tok)
	if err != nil {
		http.Error(w, "invalid token format", http.StatusBadRequest)
		return
	}
	// Trusting unverified token content
	fmt.Fprintf(w, "Hello, %s (UNVERIFIED token)\n", string(decoded))
}

func main() {
	http.HandleFunc("/me", meHandler)
	log.Println("Listening on :8086 (A07 Identification & Authentication Failures)")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
