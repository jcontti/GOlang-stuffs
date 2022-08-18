package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// creating an empty list of maps types
var bookings_maps = make([]map[string]string, 0) // we need to define the initial size of the slice, it will increase because is a slice so 0 it's ok.

// creating a struct, where we can mix diferents types likee string and unit among others.
type userDataStruct struct {
	// define the list of keys
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
} // struct can be similar to "class" concept in POO, but without inheritance. So it simpler and cleaner than using maps.

// when we are using multiple-threads we need to synchronizing them, because maybe main thread finish and do not wait for ther other thread to do theeir jobs...
var wg = sync.WaitGroup{
	// Waits for the launched gorooutine to finish
	// Package "sync" provides basic synchronization functionality
	// WaitGroup has 3 functions that we can call:
	// - Add(): sets the number of gorouutines to wait for
	// - Wait(): blocks until the Waitgroup counter is 0. It needs to be place at the end of the "main thread"
	// - Done(): Decrements the WaitGroup counter by 1. So this is called by the gorouutine to indicate that it's finished
}

func main() {
	var conferenceName = "GOlang Conference"
	// another form of var declaration, but only for variables, and we cannot define there the type.
	conferenceName2 := "GOlang Conference number 2"
	fmt.Println(conferenceName2)

	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint type DO NOT accept negatives values.

	// calling new functions and passing an input parameters
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		// call getUserInput() function
		firstName, lastName, email, userTickets := getUserInput()

		// Validate User Input function and returning 3 booleans values.
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// all conditions must be true to continue booking tickets...
		if isValidName && isValidEmail && isValidTicketNumber {

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

			// create an empty map for a user
			var userData = make(map[string]string) // bear in mind that the type key and value must be the same. string-string, int-int, etc...
			// adding key-pair values to a map
			userData["firstNameKeyName"] = firstName // ["firstNameKeyNAme"] is KeyName and the value is the variable firstName
			userData["lastNameKeyName"] = lastName
			userData["emailKeyName"] = email
			userData["numberOfTicketsKeyName"] = strconv.FormatUint(uint64(userTickets), 10) // takes our uinit value and format to a string as decimal number (10 represent decimal number, 16 represent hexadecimalnumbers)
			bookings_maps = append(bookings_maps, userData)                                  // adding a map to our list

			bookings_slice = append(bookings_slice, firstName+" "+lastName)
			fmt.Printf("The first value of the slice: %v\n", bookings_slice[0])
			fmt.Printf("The Type of the slice: %T\n", bookings_slice)
			fmt.Printf("The lenght of the slice: %v\n", len(bookings_slice))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// using "send email" function in another thread, so now our application is concurrent
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) // "go" starts a new goroutine, which is a lightweight thread managed by the Go runtime
			// so now, ouur sendTicket function runs in "background"

			// call function print first names and passing the booking slice as input parameter
			firstNames := getFirstNames(bookings_slice) //it would return a slice with the information needed
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program using break statement
				fmt.Println("Our conference is booked out. Come back next year.")
				break // terminate the for loop
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			} // if the name is not valid

			if !isValidEmail {
				fmt.Println("your email is not valid, must contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you enter is not valid")
			}
			// break // if we use break, the program ends, but we need to give the opportunity to re-enter the num of tickets
			continue // causee loop to skip the rest of its body and immediately re-testing its condition going to the next iteration.
		}
	}
	wg.Wait()
	// How to use Switch instead of multiple if-else
	/* 	city := "London"

	   	switch city {
	   	case "New York":
	   		// execute code for new york value
	   	case "Singapore", "Hong Kong":
	   		// execute some code here for Singapore and Hong Kong
	   	case "London", "Berlin":
	   		// execute some code here for London and Berlin
	   	case "Mexico":
	   		// same for mexico..
	   	default:
	   		fmt.Println("No valid city selected")
	   	} */

} // end main function

func greetUsers(paramConfName string, paramConfTickets int, paramRemainingTickets uint) {
	fmt.Printf("Welcome to our %v - func greetUsers() \n", paramConfName)

	// how to know the type of variable? we can use %T
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", paramConfTickets, paramRemainingTickets, paramConfName)
	fmt.Println("Welcome to", paramConfName, "booking application")
	// same but using Printf, %v is default format of display (we have for numberes, and others specific types), and after the comma we pass what variable.
	fmt.Printf("Welcome to %v booking app\n", paramConfName) // el \n es para una nueva linea
	fmt.Println("We have total of", paramConfTickets, "tickets and", paramRemainingTickets, "are still available")
	// same but using Printf
	fmt.Printf("We have total of %v tickets and %v are STILL available.\n", paramConfTickets, paramRemainingTickets)
	fmt.Println("Get your tickets here to attend")
} // end greetUsers function

// the first part "getFirstNames(paramBookingsSlice []string)" is creating a function that receiv a slice as parameter
// the second part "[]string" is saying that will RETURN an output type slice to where the function was called
func getFirstNames(paramBookingsSlice []string) []string {
	firstNames := []string{} // create a slice
	// go and scan bookings_slice, save the value in "booking"
	//for index, booking := range bookings_slice {
	// the above code give us problem because we do not use "index", so that's why we'll use the Blank Identifier "_" to ignore a variable we don't want to use
	for _, booking := range paramBookingsSlice {
		var fullName = strings.Fields(booking) // strings.Fields will split string separated with space, and will return a SLICE with the split elements

		firstNames = append(firstNames, fullName[0]) // adding only the first name to our slice firstNames
	}

	return firstNames
} // end getFirstNames function

// this function do not need input parameters, but would need output (return) values.
func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

// function to send tickets by email
func sendTicket(paramUserTickets uint, paramFirstName string, paramLastName string, paramEmail string) {
	time.Sleep(10 * time.Second) // simulating 10 seconds delay

	// if we can save the output string of a Print function, we shouuld use Sprintf
	var ticket = fmt.Sprintf("%v tickets for %v %v", paramUserTickets, paramFirstName, paramLastName)
	// so when I use thee variable ticket, it wouuld return the output with al the %v replaced.
	fmt.Println("##############################")
	fmt.Printf("Sending ticket:\n  %v \nto email address %v\n", ticket, paramEmail)
	fmt.Println("##############################")
	wg.Done()
}

/* NOTES:
- Go programs are organized into packages, a package is a collection of Go files..
- We have only one kind of LOOP in Go which is "for"
- in Go we NEED to use Blank Identifier "_" underscore to explicit variables that we won't use.
- A Go function can return multiple values
- Variables and Functions defined outside any function, can be accessed in all other files within the same package.
- Whenever we need to use a function of another package, we must explicit import it. (applied to our homemade package as well)
- When we want that one function from one package can be used in another package, we need to EXPORT that function. For that we just need to CAPITALIZE the first letter of the function
- We can also EXPORT: variables, functions, constants, types to be used in another packages.
- Concurrency: we use "go" before calling a function to make that function starts in another thread.
*/
