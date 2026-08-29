package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe("localhost:8080", mux)
}
