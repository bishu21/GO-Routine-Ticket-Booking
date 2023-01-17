package main

import (
	"fmt"
	"ticket-booking/helper"
	"time"
)

const ticket int = 50

var conferanceName = "Go Conferance"
var remainingTicket uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

func main() {

	greetUsers()

	for remainingTicket > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInputs(remainingTicket, firstName, lastName, email, userTicket)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)
			go sendTicket(userTicket, firstName, lastName, email)
			// call function print firstNames
			firstNames := printFirstNames()
			fmt.Printf("These are firstnames all bookings :%v\n", firstNames)

			if remainingTicket == 0 {
				// end the program
				fmt.Printf(" Our connferance is booked out. Come next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not conatin @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid.")
			}
		}

	}
}

func printFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v meeting.\n", conferanceName)
	fmt.Printf("Total available tickets are %v\n", ticket)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email :")
	fmt.Scan(&email)
	fmt.Println("Enter No. of tickets")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string,
	lastName string, email string) {
	remainingTicket = remainingTicket - userTicket

	// create a map for a user

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	// create a struct for a User
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation email on %v. \n",
		firstName, lastName,
		userTicket, email)
	fmt.Printf("%v tickets remaining for conferance.\n", remainingTicket)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTicket, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket %v\n to email address %v\n", ticket, email)
	fmt.Println("#################")
}
