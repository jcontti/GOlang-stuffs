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
	var userTickets uint
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

	remainingTickets = remainingTickets - userTickets

	// using Arrays (FIXED SIZE > how many elements the array can hold)
	// var bookings = [50]string{"John", "Nicole", "Peter"}
	// var bookings = [50]string{}
	var bookings [50]string
	//bookings[0] = "John"
	//bookings[1] = "Nicole"
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value of the array: %v\n", bookings[0])
	fmt.Printf("The second value of the array: %v\n", bookings[1])
	fmt.Printf("The Type of the array: %T\n", bookings)
	fmt.Printf("The lenght of the array: %v\n", len(bookings))

	// using Slice (is like an Array, but IS NOT fixed, so we don't waste memory, is resized when needed)
	// also, we do not need to use index to ADD elements, we use append() instead
	var bookings_slice []string
	// bookings_slice := []string{}

	bookings_slice = append(bookings_slice, firstName+" "+lastName)
	fmt.Printf("The first value of the slice: %v\n", bookings_slice[0])
	fmt.Printf("The Type of the slice: %T\n", bookings_slice)
	fmt.Printf("The lenght of the slice: %v\n", len(bookings_slice))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
