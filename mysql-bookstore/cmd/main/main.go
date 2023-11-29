package main

import (
	"log"
	"net/http"

	"github.com/burhan/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
