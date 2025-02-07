package main

import (
	"fmt"

	"example.com/structs/user"
)

// type User struct {
// 	firstName string
// 	lastName string
// 	birthDate string
// 	age int
// 	createdAt time.Time
// }

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	// newUser := User{firstName: firstName, lastName: lastName, birthDate: birthdate, age: 25, createdAt: time.Now()}
	newUser, err1 := user.New(firstName, lastName, birthdate)
	newUser2, err2 := user.New("John", "Doe", "10102008")

	admin := user.NewAdmin("text.example", "hello");

	// admin.User.OutputUserData()
	// admin.User.ClearUserName()
	admin.OutputUserData()
	admin.ClearUserName()
	admin.OutputUserData()

	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return
	}
	// fmt.Println(newUser.firstName, newUser.lastName, newUser.birthDate, newUser.createdAt)
	newUser.OutputUserData()
	newUser2.OutputUserData()
	newUser.ClearUserName()
	newUser.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
