package main // // Declare a main package (a package is a way to group functions)

import (
	"fmt" // Import the popular fmt package, which contains functions for formatting text, including printing to the console.

	"rsc.io/quote" // When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code. Visiting pkg.go.dev and searching there.
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
} // Implement a main function to print a message to the console. A main function executes by default when you run the main package.
