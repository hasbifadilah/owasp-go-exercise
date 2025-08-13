package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

// BAD: MD5 used as password hash (fast, unsalted)
func md5Hash(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := r.FormValue("user")
	pass := r.FormValue("pass")
	stored := md5Hash(pass) // intentionally insecure
	fmt.Fprintf(w, "stored password hash for %s: %s\n", user, stored)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	log.Println("Listening on :8081 (A02 Cryptographic Failures)")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
