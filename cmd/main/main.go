package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saiful-4321/crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStores(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
