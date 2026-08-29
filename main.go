package main

import (
	"net/http"
)

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app/",http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", handleHealthz)

	http.ListenAndServe("localhost:8080", mux)
}
