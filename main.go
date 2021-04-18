package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://opsdata.io", 301)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "OpsData.io")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/healthz", healthz)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
