package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Downloads a script and executes it without verifying signature or checksum.
func installHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	f, _ := os.CreateTemp("", "installer-*.sh")
	defer os.Remove(f.Name())
	io.Copy(f, resp.Body)
	f.Chmod(0o700)
	// WARNING: executing remote content is dangerous (intentional for exercise)
	cmd := exec.Command("/bin/sh", f.Name())
	cmd.Stdout = w
	cmd.Stderr = w
	_ = cmd.Run()
}

func main() {
	http.HandleFunc("/install", installHandler)
	log.Println("Listening on :8087 (A08 Software & Data Integrity Failures)")
	log.Fatal(http.ListenAndServe(":8087", nil))
}
