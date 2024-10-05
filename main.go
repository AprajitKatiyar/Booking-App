package main

import (
	"booking-app/helper"
	"fmt"
)

// Decalring global variables/package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = conferenceTickets

var bookings = []map[string]string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isNameValid, isEmailValid := helper.IsValidUserInput(firstName, lastName, email)
		if !isNameValid || !isEmailValid {
			fmt.Println("Your input data is invalid.Please try again.")
			continue
		}

		//Check if the user is trying to book more tickets than available
		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. Please try again.\n", remainingTickets)
			continue
		}

		bookTickets(firstName, lastName, email, userTickets)
		fmt.Printf("The bookings are %v\n", bookings)

		noTickets := remainingTickets == 0
		if noTickets {
			fmt.Println("All tickets have been booked!")
			break
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets available for the conference and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var firstName = booking["firstName"]
		firstNames = append(firstNames, firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets

}
func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	//Declare a map for storing user data
	//var userData map[string]string
	//Alternate way to declare a empty map
	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["tickets"] = fmt.Sprint(userTickets)
	bookings = append(bookings, userData)

	fmt.Printf("Hello %v %v, you have booked %v tickets for the conference.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)

	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

	firstNames := getFirstNames()
	fmt.Printf("The first names of the bookings are %v\n", firstNames)
}
