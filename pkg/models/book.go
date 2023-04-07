package models

import (
	"log"

	"github.com/aalekh12/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db gorm.DB

func init() {
	config.Connect()
	db = *config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	//db.NewRecord(b)
	db.Create(&b)
	log.Println(b)
	return b
}

func GetAllBooks() []Book {
	var getallbooks []Book
	db.Find(&getallbooks)
	return getallbooks
}

func GetBookById(Id int16) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("id=?", Id).Find(&getbook)
	return &getbook, db
}

func DeleteBook(Id int16) Book {
	var book Book
	db.Where("id=?", Id).Delete(book)
	return book
}
