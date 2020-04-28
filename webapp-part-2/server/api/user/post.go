package user

import (
	"encoding/json"
	"net/http"
)

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jd := json.NewDecoder(r.Body)

	aUser := &User{}
	err := jd.Decode(aUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lock.Lock()
	// start of code protected from multiple calls
	nextUserID++
	aUser.ID = nextUserID
	db = append(db, aUser)
	// end of protected code
	lock.Unlock()

	respUser := User{ID: aUser.ID, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}
