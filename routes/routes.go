package routes

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
	a "github.com/JasurbekUz/Library/controllers/authors"
	c "github.com/JasurbekUz/Library/controllers/categories"
	. "github.com/JasurbekUz/Library/controllers/books"
)

const PORT = ":8080"

func StartService () {

	router := mux.NewRouter()

	router.HandleFunc("/author", a.AddAuthor).Methods("POST")
	router.HandleFunc("/authors", a.GetAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", a.PutAuthor).Methods("PUT")

	//router.HandleFunc("/category", AddCategories).Methods("POST")
	router.HandleFunc("/categories", c.GetCategory).Methods("GET")
	//router.HandleFunc("/categories/{id}", PutCategory).Methods("PUT")

	router.HandleFunc("/book", AddBook).Methods("POST")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", PutBookById).Methods("PUT")
	router.HandleFunc("/book/{id}", DeleteBookbyId).Methods("DELETE")

	log.Printf("server is redy port %v", PORT)

	if err := http.ListenAndServe(PORT, router); err != nil {

		log.Fatalf("server listenning error: %v", err)
	}
}