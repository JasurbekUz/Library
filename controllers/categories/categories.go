package categories

import (

	"net/http"
	"encoding/json"
	//"github.com/gorilla/mux"
)

func GetCategory (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(SelectCategories())
}