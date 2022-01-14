package helper

import "strings"

func DataValidation(firstName, lastName, email string, userTickets, remianingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNoofTickets := userTickets > 0 && userTickets <= remianingTickets
	return isValidName, isValidEmail, isValidNoofTickets
}