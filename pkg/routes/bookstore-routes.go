package routes

import (
	"github.com/aalekh12/go-bookstore/pkg/cantroller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoures = func(router *mux.Router) {
	router.HandleFunc("/book/", cantroller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", cantroller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", cantroller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", cantroller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", cantroller.DeleteBook).Methods("DELETE")

}
