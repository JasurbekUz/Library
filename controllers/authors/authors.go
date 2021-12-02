package authors

import (
	"log"
	"strconv"	
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

func AddAuthor (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)
	}

	var newAuthor PostAndUpdateAuthor

	json.Unmarshal(body, &newAuthor)

	json.NewEncoder(w).Encode(InsertNewAuthor(newAuthor))
}

func GetAuthors (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(SelectAuthors())
}

func PutAuthor (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)

	iid, _ := strconv.Atoi(sid["id"])

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)

		return
	}

	var editedAuthor PostAndUpdateAuthor

	json.Unmarshal(body, &editedAuthor)

	value, err := UpdateAuthorById(iid, editedAuthor)

	if err != nil {
		json.NewEncoder(w).Encode("no data")

		return
	}

	json.NewEncoder(w).Encode(value)
}
