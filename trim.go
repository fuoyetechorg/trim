package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var memoryDatabase = &memDatabase{data: make(map[string]string)}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{hash}", RedirectHandler(memoryDatabase)).Methods(http.MethodGet)
	// Serve the home page
	r.Handle("/", http.FileServer(http.Dir(".")))
	// Serve static files
	statics := r.PathPrefix("/static/")
	statics.Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8080", r)
}
