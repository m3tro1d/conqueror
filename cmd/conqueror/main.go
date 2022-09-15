package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	s := router.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("", handler).Methods(http.MethodGet)

	server := &http.Server{Addr: "localhost:8000", Handler: router}
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
