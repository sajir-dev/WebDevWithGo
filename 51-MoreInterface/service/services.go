package service

import (
	"time"
)

type Credentials struct {
	Username string
	Password string
}

type User struct {
	Username string
	Password string
	First    string
	Parent   string
	DOB      time.Time
	Married  bool
}

func (u User) FullName() string {
	return u.First + " " + u.Parent
}

func (u User) Age() int {
	return time.Now().Year() - u.DOB.Year()
}

type UserInterface interface {
	FullName() string
	Age() int
}

func UserAge(u UserInterface) int {
	return u.Age()
}
