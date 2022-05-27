package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	// var conferenceName string = "Go Conference"
	var conferenceTickets = 50
	var remainingTickets = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	// Get type
	fmt.Printf("ConferenceTickets: %T\n", conferenceTickets)
	// var conferenceA uint = 1
	// var conferenceA = -1 not possible because uint

	conferenceTickets = 30

	var userName string
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Choose a username: ")
	fmt.Scan(&userName)

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How many tickets?")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you, %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// userName = "Tom"
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}

func main2() {
	fmt.Println("hi")
}
