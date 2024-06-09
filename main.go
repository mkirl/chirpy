package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	mux.Handle("/logo.png", http.FileServer(http.Dir("./assets/")))

	http.ListenAndServe(":8080", mux)
}
