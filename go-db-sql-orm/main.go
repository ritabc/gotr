// Using the GORM ORM in Go

package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"not null"`
	Password string
	Messages []Message
}

// Say users have many messages
type Message struct {
	Body   string `gorm:"not null"`
	User   User
	UserID uint
	gorm.Model
}

const (
	connStr = "user=rbennett dbname=gotr_sql host=localhost sslmode=disable"
)

func main() {
	log.Info("Connecting to SQL DB...")
	db, err := gorm.Open("postgres", connStr)
	// db, err := gorm.Open("sqlite3", "test.db") // Easy to use another db for testing purposes
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user = &User{Name: "Peter Jones", Username: "ps@email.com", Password: "new-pswd"}

	// If we don't already have a table users, ask orm to do it for us
	db.AutoMigrate(&User{}) // table will be created with plural version
	db.AutoMigrate(&Message{})

	// If we already have a table users:
	if err := db.Model(&User{}).Create(user).Error; err != nil {
		log.Warn(err)
	}

	users, _ := db.Model(&User{}).Find(user).Rows()
	for users.Next() {
		u := new(User)
		db.ScanRows(users, u)
		fmt.Printf("Got: Name: %v, Username: %v\n", u.Name, u.Username)
	}

	// mesgs := []*Message{
	// 	&Message{Body: "My message 1"},
	// 	&Message{Body: "My message 2"},
	// 	&Message{Body: "My message 3"},
	// }

	// // Find first user
	// db.Model(user).Find(user)
	// for _, m := range mesgs {
	// 	db.Model(user).Association("Messages").Append(m)
	// 	// This would work as well
	// 	// m.User = user
	// 	// db.Model(&{Message}).Create(m)
	// }

	// Find all of user 2's messages:
	var user_two User
	db.First(&user_two, 2)
	var messages []Message
	db.Model(&user_two).Related(&messages)
	fmt.Println(&messages)
}
