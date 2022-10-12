package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("Bookings: %v \n", bookings)
	fmt.Printf("Bookings type %T \n", bookings)
	fmt.Println("Bookings length ", len(bookings))

	remainingTickets -= userTickets

	fmt.Printf("%v remining tickets for %v\n", remainingTickets, conferenceName)
}
