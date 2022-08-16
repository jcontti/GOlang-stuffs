package main

import "fmt"

func main() {
	var conferenceName = "GOlang Conference"
	// another form of var declaration, but only for variables, and we cannot define there the type.
	conferenceName2 := "GOlang Conference number 2"
	fmt.Println(conferenceName2)

	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint type DO NOT accept negatives values.

	// how to know the type of variable? we can use %T
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking application")
	// same but using Printf, %v is default format of display (we have for numberes, and others specific types), and after the comma we pass what variable.
	fmt.Printf("Welcome to %v booking app\n", conferenceName) // el \n es para una nueva linea

	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// same but using Printf
	fmt.Printf("We have total of %v tickets and %v are STILL available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name and saving it into firstName variable previously created
	fmt.Println("Enter your first name: ")
	fmt.Scan(firstName)  // this would be empty, because we need to use the "pointer" known also as "special variable"
	fmt.Scan(&firstName) // using the pointer

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
