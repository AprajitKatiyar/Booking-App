package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Decalring global variables/package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = conferenceTickets

var bookings = []UserData{}

var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

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
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		fmt.Printf("The bookings are %v\n", bookings)

		noTickets := remainingTickets == 0
		if noTickets {
			fmt.Println("All tickets have been booked!")
			break
		}
		fmt.Println("Do you want to book more tickets? (yes/no): ")
		var choice string
		fmt.Scan(&choice)
		if choice == "no" {
			break
		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets available for the conference and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var firstName = booking.firstName
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
	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Hello %v %v, you have booked %v tickets for the conference.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)

	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

	firstNames := getFirstNames()
	fmt.Printf("The first names of the bookings are %v\n", firstNames)
}
func sendTickets(tickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)
	fmt.Println("-------------------------------")
	fmt.Printf("Sending ticket:\n%v\nTo email address: %v\n", ticket, email)
	fmt.Println("-------------------------------")
	wg.Done()

}
