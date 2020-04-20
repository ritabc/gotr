package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ritabc/gotr/types"
)

const jsonFile = "../data/user.db.json"

func main() {
	createJSONFile()
	f, err := os.Open(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	db := types.UserDB{}
	dec.Decode(&db)
	fmt.Println(db)
}

func createJSONFile() {
	users := []types.User{
		{ID: 1, Username: "John Doe", Password: "change me", Email: "johndoe@email.com"},
		{ID: 2, Username: "Jane Doe", Password: "please change me", Email: "janedoe@email.com"},
	}
	db := types.UserDB{Users: users, Type: "Simple"}
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)

	f, err := os.Create(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
