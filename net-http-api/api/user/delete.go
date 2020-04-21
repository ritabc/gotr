package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// doDelete deletes a user from the db using the path '/users/id', eg: /users/2
func doDelete(w http.ResponseWriter, r *http.Request) {
	// get the user ID from the path
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to delete user %v", id)

	lock.Lock()
	// start of code protected from multiple calls
	var tmp = []*User{}
	for _, u := range db {
		if id == u.ID {
			continue
		}
		tmp = append(tmp, u)
	}
	db = tmp
	// end of protected code
	lock.Unlock()

	w.Header().Set("Content-Type", "application/json")
}
