package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets uint8 = 50

var conferenceName string = "Go Conference"
var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

// this is an struct in golang
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

func main() {
	greetUsers()

	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// ask user for their info
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	isValidName, isValidEmail, isValidTicket := validateUserInputs(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicket {

		bookTicket(firstName, lastName, email, userTickets)
		firstNames := getFirstNames()
		fmt.Printf("Firsts names of bookings are: %v\n", firstNames)

		wg.Add(1)
		go sendTickets(firstName, lastName, email, userTickets)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out, Come Back next year!")
			// break
		}
	} else {
		errorInputMessages(isValidName, isValidEmail, isValidTicket)
	}
	wg.Wait()
}

// Firsts messages for users
func greetUsers() {
	// print welcome messages
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// function for validate user input data
func validateUserInputs(firstName string, lastName string, email string, userTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= 50

	return isValidName, isValidEmail, isValidTicket
}

// logic for each user book
func bookTicket(firstName string, lastName string, email string, userTickets uint8) {
	remainingTickets -= userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remining tickets for %v\n", remainingTickets, conferenceName)
}

// split names of bookings
func getFirstNames() []string {
	firstNames := []string{}
	// for each in golang
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// messages of error when any user input is fail
func errorInputMessages(isValidName bool, isValidEmail bool, isValidTicket bool) {
	if !isValidName {
		fmt.Println("First name or last name you entered is too short")
	}
	if !isValidEmail {
		fmt.Println("Email address you entered doesn't contain @ sign")
	}
	if !isValidTicket {
		fmt.Println("Number of tickets you entered is invalid")
	}
}

func sendTickets(firstName string, lastName string, email string, userTickets uint8) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
