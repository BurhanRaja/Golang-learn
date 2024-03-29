package models

import (
	"github.com/burhan/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBook() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (Book, *gorm.DB) {
	var book Book
	db.First(&book, Id)
	return book, db
}

func (b *Book) CreateBook() Book {
	db.Create(&b)
	return *b
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
