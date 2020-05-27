// Using the GORM ORM in Go

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string
	Username string
	Password string
}

type Profile struct {
	TZ *time.Location
}
type MyORM struct {
	db *sql.DB
}

func main() {
	log.Info("Connecting to SQL DB...")
	connStr := "user=rbennett dbname=gotr_sql host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user = &User{"Peter Jones", "ps@email.com", "new-pswd"}

	var orm = &MyORM{db}
	if err := orm.Create(user); err != nil {
		log.Fatal(err)
	}

	users, _ := orm.Query()
	for _, u := range users {

		fmt.Printf("Got: Name: %v, Username: %v\n", u.Name, u.Username)
	}
}

func (o *MyORM) Query() (users []User, err error) {
	if o == nil {
		return nil, errors.New("Invalid parameter")
	}
	rows, err := o.db.Query("SELECT name, username FROM users")
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Name, &u.Username)
		users = append(users, u)
	}
	return
}

func (o *MyORM) Create(u *User) (err error) {
	if o == nil {
		return errors.New("Invalid parameter")
	}
	_, err = o.db.Exec(`INSERT INTO users(name, username, password) VALUES($1, $2, $3);`, u.Name, u.Username, u.Password)
	return
}
