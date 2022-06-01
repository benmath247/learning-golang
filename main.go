package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	// var conferenceName string = "Go Conference"

	var conferenceTickets = 50
	var remainingTickets = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")

	// var conferenceA int = 1
	// var conferenceA = -1 not possible because uint

	for {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "are still available")
		fmt.Println("Get your tickets here to attend")

		// Get type
		// fmt.Printf("ConferenceTickets: %T\n", conferenceTickets)

		var userName string
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// var bookings = [50]string{"Ben", "Yoni", "Shira"} // defining an array
		var bookings = []string{}

		fmt.Println("Choose a username: ")
		fmt.Scan(&userName)
		fmt.Println(Hello(userName))

		fmt.Println("What is your first name?")
		fmt.Scan(&firstName)
		fmt.Println("What is your last name?")
		fmt.Scan(&lastName)

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("What is your email?")
		fmt.Scan(&email)

		fmt.Println("How many tickets?")
		fmt.Scan(&userTickets)

		fmt.Printf("Thank you, %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		// userName = "Tom"
		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		remainingTickets = remainingTickets - userTickets
		// fmt.Printf("The whole array: %v\n", bookings)
		// fmt.Printf("The first values: %v\n", bookings[0])
		// fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array length: %v\n", len(bookings))

		// slice is an abstraction of an array
	}

}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
