package main

import(
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)
 
//Book Struct (Model)
type Book struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Author *Author `json:"author"`
}

// Author Struct
type Author struct {
  Firstname string `json:"firstname"`
  Lastname string `json:"lastname`
}

// Init books var as a slice book struct
var books []book

// Get All Book
func getBooks(w http.ResponseWriter, r *http.Request){
  w.Header().Set("contect-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook(w http.ResposeWriter, r *http.Request){

}

// Create a New Book
func createBook(w http.ResposeWriter, r *http.Request){

}

func updateBook(w http.ResposeWriter, r *http.Request){

}

func deleteBook(w http.ResposeWriter, r *http.Request){

}
func main() {
  //Init router
  r:=mux.NewRouter()

  //Mock Data -@todo - implement DB
  books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firsrname: "John", Lastname:"Doe"}})
  books = append(books, Book{ID: "2", Isbn: "843737", Title: "Book Two", Author: &Author{Firsrname: "Steve", Lastname:"Smith"}})

  // Route Handlers / Endpoints
  r.HandleFunc("/api/books", getbooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getbooks).Methods("GET")
  r.HandleFunc("/api/books", createbooks).Methods("POST")
  r.HandleFunc("/api/books/{id}", updatebooks).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deletebooks).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000, r"))
}
