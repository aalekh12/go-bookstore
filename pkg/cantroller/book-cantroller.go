package cantroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aalekh12/go-bookstore/pkg/models"
	"github.com/aalekh12/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookDetails, _ := models.GetBookById(int16(Id))
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookId"]
	Id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Error While Deleteing")
	}
	bookdetails := models.DeleteBook(int16(Id))
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatebook := &models.Book{}
	utils.ParseBody(r, updatebook)
	vars := mux.Vars(r)
	bookid := vars["bookId"]
	Id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Error While Updating")
	}
	bookdetails, db := models.GetBookById(int16(Id))
	if bookdetails.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if bookdetails.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if bookdetails.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
