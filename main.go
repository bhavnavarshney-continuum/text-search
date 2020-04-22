package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/search/{hint}", searchText).Methods("GET")
	http.ListenAndServe("localhost:8090", router)
}

func searchText(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Searching:", vars["hint"])
	naiveSearch(vars["hint"])
	//improvedUnmarshalling()
}
