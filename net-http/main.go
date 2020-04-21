package main

import (
	"fmt"
	"log"
	"net/http"
)

const usersAPIResp = `
<html>
<body>
<p>Hi, thanks for calling my /users API with HTTP Method '%v'</p>
<p>This is the %v call to this API</p>
</body>
</html>`

var userCounter int

type reportCounter struct {
	counter int
}

func main() {
	http.HandleFunc("/users", usersHandleFunc)
	var rc reportCounter
	http.Handle("/reports/", &rc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /reports")
	userCounter++
	s := fmt.Sprintf(usersAPIResp, r.Method, userCounter)
	fmt.Fprint(w, s)
}

func (rc *reportCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /reports")
	rc.counter++
	s := fmt.Sprintf("/reports API call count: %v", rc.counter)
	fmt.Fprint(w, s)
}
