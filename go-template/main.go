package main

import (
	"fmt"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type BooksPage struct {
	BooksListing []Book
}

type Book struct {
	ISBN    string
	Author  string
	Title   string
	PubDate string
	Price   Currency
}

type Currency float64

func main() {

	http.HandleFunc("/books/", booksHandler)
	// Register static files handle '/index.html' ->serves-> client/index.html
	http.Handle("/", http.FileServer(http.Dir("client")))
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	books := &BooksPage{
		BooksListing: []Book{
			{"0458-345", "Author 2", "title 2", "2014-03-30", 9.43},
			{"0458-345", "Author 1", "title 1", "2014-03-30", 10.52},
			{"0458-345", "Author 3", "title 3", "2014-03-30", 15.97},
		},
	}
	tmpl := template.Must(template.ParseFiles("server/books/index.html"))
	err := tmpl.Execute(w, books)
	if err != nil {
		panic(err)
	}
}

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
