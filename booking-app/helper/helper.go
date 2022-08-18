package helper

import "strings"

// includes common, helper and shared functions to the main package/app

func ValidateUserInput(paramFirstName string, paramLastName string, paramEmail string, paramUserTickets uint, paramRemainingTickets uint) (bool, bool, bool) {
	// firstName and LastName validation: needs to enter at least 2 characters
	// var isValidName bool =  len(firstName) >= 2 && len(lastName) >= 2
	isValidName := len(paramFirstName) >= 2 && len(paramLastName) >= 2
	// needs to enter a valid email format containg "@" sign:
	isValidEmail := strings.Contains(paramEmail, "@")
	// needs to enter correct number of ticket (positive + greater than 0) and user buy is lest or equal to remaing tickets
	isValidTicketNumber := paramUserTickets > 0 && paramUserTickets <= paramRemainingTickets

	return isValidName, isValidEmail, isValidTicketNumber // the output of this function are 3 boolean, that's why we have (bool, bool, bool) at the beginning
} // end validateUserInpuut function
