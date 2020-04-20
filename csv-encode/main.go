package main

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/ritabc/gotr/types"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := readJSONFile("../data/user.db.json")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/user.db.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, user := range db.Users {
		ss := user.EncodeAsStrings()
		w.Write(ss)
	}
	w.Flush()

	err = w.Error()
	if err != nil {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *types.UserDB, err error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)

	db = new(types.UserDB)
	dec.Decode(db)

	return
}
