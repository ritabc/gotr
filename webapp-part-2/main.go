package main

import (
	"log"
	"net/http"

	"github.com/ritabc/gotr/webapp-part-2/server/api/user"
)

func main() {
	// Register static files handle '/index.html -> client/index.html
	http.Handle("/", http.FileServer(http.Dir("client")))
	// Register RESTful endpoint handler for '/users/
	http.Handle("/users/", &user.UserAPI{})
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
