package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	log.Info("Connecting to SQL DB...")
	connStr := "user=rbennett dbname=gotr_sql host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var name, username, pswd string = "Peter Jones", "ps@email.com", "new-pswd"

	qr, err := db.Exec(`INSERT INTO users(name, username, password) VALUES($1, $2, $3);`, name, username, pswd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Query Result: %v", qr)

	rows, err := db.Query("SELECT id, name, username FROM users")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&id, &name, &username)
		fmt.Printf("Got: ID: %v, Name: %v, Username: %v\n", id, name, username)
	}
}
