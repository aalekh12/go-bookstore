package main

import (
	"log"
	"net/http"

	"github.com/aalekh12/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoures(r)
	http.Handle("/", r)

	log.Println("HTTP Server started on 9090")
	http.ListenAndServe(":9090", r)
}
