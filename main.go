package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets available for the conference and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	var bookings [50]string

	//Alternate way to declare an array
	//var bookings = [50]string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//take user input
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array is %v\n", bookings)
	fmt.Printf("The first element is %v\n", bookings[0])
	fmt.Printf("Type of the array is %T\n", bookings)
	fmt.Printf("The length of the array is %v\n", len(bookings))

	//To print the memory address of the variable
	//fmt.Println(&remainingTickets)

	fmt.Printf("Hello %v %v, you have booked %v tickets for the conference.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)

	fmt.Printf("Remaining tickets: %v", remainingTickets)

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
