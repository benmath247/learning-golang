package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var conferenceTickets = 50
	var remainingTickets = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	conferenceTickets = 30

	var userName string
	var userTickets int
	// ask for their name
	userName = "Tom"

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}

func main2() {
	fmt.Println("hi")
}
