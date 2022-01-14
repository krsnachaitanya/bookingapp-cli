package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var confName = "GO-CONF"
const confTickets uint = 50
var remianingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// Wait for the multi thread process to finish
var wg = sync.WaitGroup{}

func main () {
	greetUser()

	for remianingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidNoofTickets := helper.DataValidation(firstName, lastName, email, userTickets, remianingTickets)

		if isValidName && isValidEmail && isValidNoofTickets {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			
			if remianingTickets >= 1 {
				fmt.Printf("Only %v tickets are available for %v\n", remianingTickets, confName)
			}
			
			// Print people attending conference
			fmt.Printf("People attending this conference: %v\n", getFirstNames())

			// Stop taking bookings if no tickets are available
			if remianingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First name or last name is too short!\n")
			}
			if !isValidEmail {
				fmt.Printf("email address you entered is invalid!\n")
			}
			if !isValidNoofTickets {
				fmt.Printf("Number of tickets you entered is invalid!\n")
			}
		}
	}
	wg.Wait()
}

func greetUser () {
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remianingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames () []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput () (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket (userTickets uint, firstName, lastName, email string) (uint, []userData) {
	remianingTickets = remianingTickets - userTickets

	// Create a map for user
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
		
	// Add user details to array
	bookings = append(bookings, userData)
	
	// Print Booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	return remianingTickets, bookings
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("####################################")
	wg.Done()
}