package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Debug endpoint that leaks environment and request headers
func debugHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ENV:")
	for _, e := range os.Environ() {
		fmt.Fprintln(w, e)
	}
	fmt.Fprintln(w, "\nHEADERS:")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", k, v)
	}
}

// Overly permissive CORS behavior (example)
func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/debug", debugHandler)
	http.HandleFunc("/cors", corsHandler)
	log.Println("Listening on :8084 (A05 Security Misconfiguration)")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
