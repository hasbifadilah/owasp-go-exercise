package main

import (
	"fmt"
	"log"
	"net/http"
)

// currentUserID reads an X-User-ID header (UNTRUSTWORTHY)
func currentUserID(r *http.Request) string {
	return r.Header.Get("X-User-ID")
}

// /profile?user_id=123
// Vulnerable: trusts query param or header for identity.
func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = currentUserID(r)
	}

	if userID == "" {
		http.Error(w, "missing identity", http.StatusUnauthorized)
		return
	}
	// Fake data access
	fmt.Fprintf(w, "Profile for user %s\n", userID)
}

func main() {
	http.HandleFunc("/profile", profileHandler)
	log.Println("Listening on :8080 (A01 Broken Access Control)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
