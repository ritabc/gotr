package types

import (
	"encoding/xml"
	"strconv"
)

var header = []string{`id`, `fname`, `lname`, `username`, `password`, `email`}

type User struct {
	ID        int    `json:"id" xml:"id,attr"`
	Firstname string `json:"firstname" xml:"name>first"`
	Lastname  string `json:"lastname" xml:"name>last"`
	Username  string `json:"username,omitempty" xml:"secret>username"`
	Password  string `json:"password,omitempty" xml:"secret>password"`
	Email     string `json:"email,omitempty" xml:"email"`
}

type UserDB struct {
	XMLName xml.Name `json:"-" xml:"users"`
	Type    string   `json:"type,omitempty" xml:"type"`
	Users   []User   `json:"users,omitempty" xml:"user"`
}

func (u User) EncodeAsStrings() (ss []string) {
	ss = make([]string, 6)
	ss[0] = strconv.Itoa(u.ID)
	ss[1] = u.Firstname
	ss[2] = u.Lastname
	ss[3] = u.Username
	ss[4] = u.Password
	ss[5] = u.Email
	return ss
}

func GetHeader() []string {
	return header
}

func (user *User) FromCSV(ss []string) {
	if user == nil {
		return
	}
	if ss == nil {
		return
	}
	user.ID, _ = strconv.Atoi(ss[0])
	user.Firstname = ss[1]
	user.Lastname = ss[2]
	user.Username = ss[3]
	user.Password = ss[4]
	user.Email = ss[5]
}
