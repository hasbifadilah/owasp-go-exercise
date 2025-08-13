package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// No logging of failures; success logs include secrets.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	if pass != "s3cr3t" {
		// Missing failure log
		http.Error(w, "invalid", http.StatusUnauthorized)
		return
	}
	// BAD: logging sensitive info
	log.Printf("user %s logged in at %s with pass=%s", user, time.Now().Format(time.RFC3339), pass)
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/login", loginHandler)
	log.Println("Listening on :8088 (A09 Security Logging & Monitoring Failures)")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
