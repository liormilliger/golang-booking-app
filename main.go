package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask users for their name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)
}
