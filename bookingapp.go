package main

import (
	"Booking-App/helper"
	"fmt"
	"time"
)

const totalTickets = 50

var remainingTickets uint = 50
var conferenceName string = "Go Conference"

// var bookings []string
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {

	greetUser()

	// for remainingTickets < 0 {
	// For loop with condition
	//}

	// for loop with infinite
	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.IsValidInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookConference(firstName, lastName, userTicket, email)
			sendTicket(firstName, lastName, email, userTicket)

			firstNames := getFirstName()
			fmt.Printf("The first name of our bookings are %v \n", firstNames)

			noTicketsAvailable := remainingTickets == 0

			if noTicketsAvailable {
				fmt.Println("Our conference are booked out. come for next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name which you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid. Email doesn't contain @ symbol")
			}
			if !isValidTicket {
				fmt.Println("Number of ticket entered is invalid")
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have totol of %v ticket and %v tickets are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	// To find the data types
	// fmt.Printf("Type of totalTickets is %T, remainingTickets is %T and conferenceName is %T \n", totTicket, remTicket, confName)
}

func getFirstName() []string {
	// for index, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }
	// In above syntax, index is not used. so it will error.
	// whenever some variable you are using and if it is unused, in that place we can use _
	firstNames := []string{}

	for _, booking := range bookings {
		// var names = strings.Fields(booking) // For array of string
		// firstNames = append(firstNames, booking["firstName"]) // For Map
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookConference(firstName string, lastName string, userTicket uint, email string) {
	remainingTickets = totalTickets - userTicket

	// Create a Map
	// userData := make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, email string, noOfTicket uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", noOfTicket, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: \n %v to email \n %v \n", ticket, email)
	fmt.Println("##################")
}
