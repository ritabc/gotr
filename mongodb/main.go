// using mgo package

package main

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo/bson"

	log "github.com/SpirentOrion/logrus"
	"github.com/globalsign/mgo"
)

const (
	url = "localhost"
)

func main() {
	dbName := "test"
	if 1 == len(os.Args) {
		log.Warnf("No db specified, using '%v'", dbName)
	} else {
		dbName = os.Args[1]
	}
	// list colllections in selected db
	// connecting to mongodb server
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)

	db := session.DB(dbName)
	if db == nil {
		log.Errorf("db '%v' not found, exiting... ", dbName)
		return
	}

	// iterate collections
	cols, err := db.CollectionNames()
	if err != nil {
		log.Warnf("No collections in db '%v'", dbName)
	}
	fmt.Printf("Collections in db '%v':\n", dbName)
	for _, c := range cols {
		fmt.Printf("[%v]\n", c)
		listDocs(db, c)
	}
}

func listDocs(db *mgo.Database, col string) {
	coll := db.C(col)
	if coll == nil {
		return
	}

	// Option 1
	// Use interface as way of discovering the data, what type of documents are in the collection
	// var result []interface{}

	// Option 2
	// Add more structure now that we know what type of data
	// var result []map[string]interface{} // -> []bson.M

	// Option 3
	type Document struct {
		ID   int    `json:"_id, omitempty"`
		Desc string `json: "desc, omitempty"`
		Done bool   `json:"done, omitempty"` // we're using 'done' in the db, but 'Done' in our code
	}
	var result []map[string]string

	// pass nil to find if we want everything
	coll.Find(nil).All(&result)
	for i, d := range result {
		fmt.Printf("\tDoc%3v - %#v\n", i+1, d)
		obj := bson.ObjectId(d["_id"]) // Get MongoDB's version of the ID
		fmt.Printf("\t\t Hex: %v, String: %v, Time: %v\n", obj.Hex(), obj.String(), obj.Time())
	}
}
