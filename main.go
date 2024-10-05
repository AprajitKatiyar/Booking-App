package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets available for the conference and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	var bookings []string
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets would you like to book? ")
		fmt.Scan(&userTickets)

		isNameValid := len(firstName) > 2 && len(lastName) > 2
		isEmailValid := strings.Contains(email, "@")

		if !isNameValid || !isEmailValid {
			fmt.Println("Your input data is invalid.Please try again.")
			continue
		}

		//Check if the user is trying to book more tickets than available
		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. Please try again.\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Hello %v %v, you have booked %v tickets for the conference.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive a confirmation email at %v\n", email)

		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("The first names of the bookings are %v\n", firstNames)

		noTickets := remainingTickets == 0
		if noTickets {
			fmt.Println("All tickets have been booked!")
			break
		}
	}

	//Print the data type of the variables
	//fmt.Printf("Data type of conferenceName is %T\n", conferenceName)
	//fmt.Printf("Data type of conferenceTickets is %T\n", conferenceTickets)
	//fmt.Printf("Data type of remainingTickets is %T\n", remainingTickets)

	//Can also specify the types yourself
	// var conferenceName string = "Go Conference"
	// var conferenceTickets int = 50
	//var ticketPrice float64 = 100.50

	//Another way to declare variables
	//myName := "John Doe" //cannot do this for a constant though

}
