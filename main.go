package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Init books
var books []Book

func main() {
	//Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Title: "Breakfast Book", Author: &Author{FirstName: "Connor", LastName: "Williams"}})

	books = append(books, Book{ID: "2", Title: "Breakfast Book", Author: &Author{FirstName: "Connor", LastName: "Williams"}})

	books = append(books, Book{ID: "3", Title: "Breakfast Book", Author: &Author{FirstName: "Connor", LastName: "Williams"}})

	//endpoints defined
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(respWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("getBooks hit on port 8080...")
	respWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respWriter).Encode(books)
}

func getBook(respWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("getBook hit on port 8080...")
	respWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) //Get params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(respWriter).Encode(item)
			return
		}
	}
	json.NewEncoder(respWriter).Encode(&Book{})
}

func createBook(respWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Creating book...")
	respWriter.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))

	fmt.Println("Book created ... ", book.ID)
	books = append(books, book)
	json.NewEncoder(respWriter).Encode(book)
}

func updateBook(respWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Updating book...")
	respWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(respWriter).Encode(book)
			break
		}
	}
	json.NewEncoder(respWriter).Encode(books)
}

func deleteBook(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(respWriter).Encode(books)
}
