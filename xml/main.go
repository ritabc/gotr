package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/ritabc/gotr/types"
)

const (
	xmlFile  = "../data/user.db.xml"
	jsonFile = "../data/user.db.json"
)

func main() {
	createXMLFile()
	f, err := os.Open(xmlFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	db := types.UserDB{}
	dec.Decode(&db)
	fmt.Println(db)
}

func createXMLFile() {
	db, err := readJSONFile(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create(xmlFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	xmlEnc := xml.NewEncoder(f)
	xmlEnc.Encode(db)
}

func readJSONFile(s string) (db *types.UserDB, err error) {
	f, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)

	db = new(types.UserDB)
	dec.Decode(db)

	return
}
