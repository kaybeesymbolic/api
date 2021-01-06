package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var books []*Book

//Book for api
type Book struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"desc"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, each := range books {
		if each.ID == id {
			json.NewEncoder(w).Encode(each)
			return
		}
	}
}
func insertBook(w http.ResponseWriter, r *http.Request) {
	var book *Book
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	books = append(books, book)
}
func main() {
	books = append(books, &Book{Name: "Physics", ID: "224453", Description: "quantum physics"},
		&Book{Name: "Chemistry", ID: "342453", Description: "Study of atomic energy"},
		&Book{Name: "Biology", ID: "9988766", Description: "Molecular life"})

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks)
	router.HandleFunc("/api/book/{id}", getBook)
	router.HandleFunc("/api/add", insertBook)
	fmt.Println("........ starting server ............")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//my repo url
//https://github.com/kaybeesymbolic/api.git
