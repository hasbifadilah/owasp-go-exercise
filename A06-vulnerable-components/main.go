package main

import (
	"fmt"
	"log"
	"net/http"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This exercise shows a project that may pin outdated deps.")
	fmt.Fprintln(w, "Run tools like 'go list -m -u all' and 'govulncheck ./...' in your environment.")
}

func main() {
	http.HandleFunc("/", versionHandler)
	log.Println("Listening on :8085 (A06 Vulnerable Components)")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
