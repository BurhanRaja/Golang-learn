package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/burhan/bookstore/pkg/models"
	"github.com/burhan/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("%v is not a valid integer.", ID)
		fmt.Print(msg)
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Handle Body
	updateBook := models.Book{}
	utils.ParseBody(r, updateBook)
	// ID
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("%v is not a valid integer.", ID)
		fmt.Print(msg)
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)                     // Getting ID from Parameter
	bookId := params["bookId"]                // Accessing ID
	ID, err := strconv.ParseInt(bookId, 0, 0) // converting ID String to Integer
	if err != nil {
		msg := fmt.Sprintf("%v is not a valid integer.", ID)
		fmt.Print(msg)
	}
	book := models.DeleteBook(ID) // Delete book
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
