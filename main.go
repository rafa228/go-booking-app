package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	//"strings"
)

const conferenceTickets int = 100

var conferenceName = "GO Conference"
var remainingTickets uint = 100
var bookings = make([]map[string]string, 0)

func main() {
	//fmt.Printf("Welcome to %v Booking App\n", conferenceName)
	//greetUsers func
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInputs()

		//validate users input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//bookTicket function
			//bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			remainingTickets = remainingTickets - uint(userTickets)

			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = email
			userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Println()
			fmt.Printf("%v tickets remaining for this %v\n", remainingTickets, conferenceName)

			//printFirstName func
			firstNames := getFirstNames()

			fmt.Printf("These are all our bookings: %v\n", firstNames)
			fmt.Println()

			if remainingTickets == 0 {
				fmt.Println("Our conference ticket is booked out, come back next year")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short, please input a valid name")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is invalid, please input a valid format")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you want to book is 0 or greater than remaining tickets, our remaining ticket is %v", remainingTickets)
				fmt.Println()
			}
			//fmt.Printf("Sorry you can only book %v tickets, you book %v tickets\n", remainingTickets, userTickets)
			//fmt.Println("Your input data is invalid, please check it and try again")
		}

	}

	//userTickets = 2
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInputs() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Hello, please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book...")
	fmt.Scan(&userTickets)
	fmt.Println()

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets int, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//create a map for users
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Println()
	fmt.Printf("%v tickets remaining for this %v\n", remainingTickets, conferenceName)
}
