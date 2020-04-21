package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// doPut updates a user from the db using the path '/users/id' a JSON request body, eg: /users/2
func doPut(w http.ResponseWriter, r *http.Request) {
	// Decode input JSON body
	jd := json.NewDecoder(r.Body)
	inputUserData := &User{}
	err := jd.Decode(inputUserData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the ID of the user to edit from the path
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lock.Lock()
	// start of code protected from multiple calls
	// store updated user information in tmpUser
	var tmpUser User
	for _, u := range db {
		if id == u.ID {
			tmpUser = *u
			break
		}
	}
	inputUsername := inputUserData.Username
	inputPassword := inputUserData.Password
	var respUser User
	switch {
	case inputUsername != "" && inputPassword != "":
		respUser = User{ID: id, Username: inputUsername, Password: inputPassword}
	case inputUsername != "":
		respUser = User{ID: id, Username: inputUsername, Password: tmpUser.Password}
	case inputPassword != "":
		respUser = User{ID: id, Username: tmpUser.Username, Password: inputPassword}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	// delete tmpUser from db
	var tmpDB = []*User{}
	for _, u := range db {
		if id == u.ID {
			continue
		}
		tmpDB = append(tmpDB, u)
	}
	db = tmpDB
	// append respUser
	db = append(db, &respUser)
	// end of protected code
	lock.Unlock()

	je := json.NewEncoder(w)
	je.Encode(respUser)
	w.Header().Set("Content-Type", "application/json")
}
