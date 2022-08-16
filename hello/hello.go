package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

// Write code to call the Hello function we created in greetings module (folder), then print the function's return value.

// In this code, you:
// Declare a main package. In Go, code executed as an application must be in a main package.
// Import two packages: example.com/greetings and the fmt package. This gives your code access to functions in those packages. Importing example.com/greetings (the package contained in the module you created earlier) gives you access to the Hello function. You also import fmt, with functions for handling input and output text (such as printing text to the console).
// Get a greeting by calling the greetings package’s Hello function.
/* Configure the log package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
Assign both of the Hello return values, including the error, to variables.
Change the Hello argument from Gladys’s name to an empty string, so you can try out your error-handling code.
Look for a non-nil error value. There's no sense continuing in this case.
Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program. */

/* Create a names variable as a slice type holding three names.
Pass the names variable as the argument to the Hellos function. */
