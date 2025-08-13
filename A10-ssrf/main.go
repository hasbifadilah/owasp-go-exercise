package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// Proxies any URL parameter without checks (SSRF)
func fetchHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/fetch", fetchHandler)
	log.Println("Listening on :8089 (A10 SSRF)")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
