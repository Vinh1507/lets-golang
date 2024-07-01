package main

import (
	"fmt"
	"lets-golang/loop"
)

func main() {
	fmt.Println("Hello World")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Println("Hello", conferenceName)

	remainingTickets--
	fmt.Println("remainingTickets", remainingTickets)

	fmt.Printf("%v are still available\n", remainingTickets)

	var userName string
	var userTickets int

	fmt.Scan(&userName)

	fmt.Print(userName, userTickets)

	var bookings = [50]string{"Nam", "Bac"}

	var userNames [50]string

	userNames[10] = "Tuan"

	fmt.Print(bookings)
	fmt.Print(userNames)

	var users []string

	users = append(users, "admin")

	fmt.Print(users)

	loop.LoopHandler()
}
