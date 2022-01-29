package helper

import "strings"

// In order to export the functionality, we need to give first letter capital for variable or methods.
func IsValidInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
