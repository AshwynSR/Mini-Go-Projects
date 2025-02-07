package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	age       int
	CreatedAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

// constructor/creation function of a Struct
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("FirstName, LastName and BirthDate are mandatory fields!")
	}
	return &User{FirstName: firstName, LastName: lastName, BirthDate: birthdate, CreatedAt: time.Now()}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{FirstName: "Admin", LastName: "Admin", BirthDate: "____", CreatedAt: time.Now()},
	}
}

// Methods of a Struct
func (u *User) OutputUserData() {
	u.age = 30
	// (*u).age = 30
	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.CreatedAt, u.age)
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt, (*u).age)
}

// Mutation Method of a Struct
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
