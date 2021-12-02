package books

import (

	"log"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

func AddBook (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)
	}

	var newBook UpdateAndPostBook

	json.Unmarshal(body, &newBook)

	json.NewEncoder(w).Encode(InsertNewBook(newBook))
}

func GetBooks (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(SelectAllBooks())
}

func GetBookById (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)

	iid, _ := strconv.Atoi(sid["id"])

	value, err := SelectBookById(iid)

	if err != nil {

		json.NewEncoder(w).Encode("no data")

		return	
	}

	json.NewEncoder(w).Encode(value)
}

func PutBookById (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)

	iid, _ := strconv.Atoi(sid["id"])

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)

		return
	}

	var newBook UpdateAndPostBook

	json.Unmarshal(body, &newBook)

	value, err := UpdateBookById(iid, newBook)

	if err != nil {
		json.NewEncoder(w).Encode("no data")

		return
	}

	json.NewEncoder(w).Encode(value)
}

func DeleteBookbyId (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)

	iid, _ := strconv.Atoi(sid["id"])

	value, err := DbDeleteBookbyId(iid)

	if err != nil {
		json.NewEncoder(w).Encode("no data")

		return
	}

	json.NewEncoder(w).Encode(value)
}