package main

import (
	"fmt"
	"strings"
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
		var noTicketsRemaining bool = remainingTickets == 0

		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is sold out. Come back next year.")
			break
		}

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
		var firstNames = []string{}

		fmt.Println("Choose a username: ")
		fmt.Scan(&userName)
		fmt.Println(Hello(userName))

		fmt.Println("What is your first name?")
		fmt.Scan(&firstName)
		firstNames = append(firstNames, firstName)
		fmt.Println("What is your last name?")
		fmt.Scan(&lastName)

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("What is your email?")
		fmt.Scan(&email)

		fmt.Println("How many tickets?")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			fmt.Printf("Thank you, %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			// userName = "Tom"
			fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

			remainingTickets = remainingTickets - userTickets
			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first values: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// slice is an abstraction of an array
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("You bought the last tickets! Enjoy the show.")
		} else {
			fmt.Printf("We only have %v tickets left. Please book fewer tickets.\n", remainingTickets)
		}
	}
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
