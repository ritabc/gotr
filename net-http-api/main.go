package main

import (
	"log"
	"net/http"

	"github.com/ritabc/gotr/net-http-api/api/user"
)

func main() {
	// Register RESTful endpoint handler for '/users/
	http.Handle("/users/", &user.UserAPI{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// start server
}
