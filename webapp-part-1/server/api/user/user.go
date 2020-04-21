package user

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UserAPI struct{}
type User struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db = []*User{}
var nextUserID uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodDelete:
		doDelete(w, r)
	case http.MethodPut:
		doPut(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsuported methed '%v' to %v\n", r.Method, r.URL)
		log.Printf("Unsuported method '%v' to %v\n", r.Method, r.URL)
	}
}
