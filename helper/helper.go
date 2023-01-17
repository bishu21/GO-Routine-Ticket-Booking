package helper

import "strings"

func ValidUserInputs(remainingTicket uint, firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 &&
		userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
