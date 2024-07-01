package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       uint
}

func main() {
	users := make([]User, 0)

	var user1 = User{
		firstName: "Nam",
		lastName:  "Bui",
		age:       10,
	}

	var admin = User{
		firstName: "Vinh",
		lastName:  "Bui",
		age:       22,
	}

	users = append(users, user1)
	users = append(users, admin)

	for index, user := range users {
		firstName := user.firstName
		fmt.Println(index, firstName)
	}
}
