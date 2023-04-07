package main

import (
	"log"
	"net/http"

	"github.com/aalekh12/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoures(r)
	http.Handle("/", r)

	log.Println("HTTP Server started on 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
