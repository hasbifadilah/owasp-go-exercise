package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

// Weak token: only 4 random bytes (too short), no expiry tracking
func weakResetToken() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func resetRequestHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	token := weakResetToken()
	// No rate limit, no expiry, token too short
	fmt.Fprintf(w, "Sent reset token %s to %s (pretend)\n", token, email)
}

func main() {
	http.HandleFunc("/password/reset", resetRequestHandler)
	log.Println("Listening on :8083 (A04 Insecure Design)")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
