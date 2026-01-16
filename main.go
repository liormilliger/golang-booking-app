package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want? ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive a confirmation email at %v\n", email)

		fmt.Printf("%v remaining tickets for the %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings %v\n", bookings)
	}
}
