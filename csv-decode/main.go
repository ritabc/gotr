package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ritabc/gotr/types"

	log "github.com/SpirentOrion/logrus"
)

const (
	csvDB = "../data/user.db.csv"
)

func main() {
	log.Info("CSV Decoding")

	f, err := os.Open(csvDB)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read() // Ignore header
	for {
		csvRecord, err := r.Read()
		if err == nil {
			process(csvRecord)
		} else if err == io.EOF {
			break
		} else {
			log.Fatal(err)
		}
	}
}

func process(ss []string) {
	u := &types.User{}
	u.FromCSV(ss)
	fmt.Println(u.Firstname, u.Email)
}
