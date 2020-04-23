package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	rest()
	//cmd()
}

func rest() {
	router := mux.NewRouter()
	router.HandleFunc("/search/{hint}", searchText).Methods("GET")
	http.ListenAndServe("localhost:8090", router)
}

func cmd() {
	hint := os.Args[1]
	fmt.Println("Searching:", hint)
	naiveSearch(hint)
	jsonParserLibrary(hint)
	goJSONQ(hint)
}

func searchText(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Println("\nSearching:", vars["hint"])
	//naiveSearch(vars["hint"])
	//jsonParserLibrary(vars["hint"])
	goJSONQ(vars["hint"])
}
