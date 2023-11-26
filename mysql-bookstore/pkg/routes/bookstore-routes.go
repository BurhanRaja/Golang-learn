package routes

import (
	"github.com/burhan/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoute = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetAllBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}