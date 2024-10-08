package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and only %v tickets left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var emailAddress string
	var userTickets uint
	var lastName string

	for {
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailAddress)

		fmt.Println("Enter how many tickets you would like to purchase: ")
		fmt.Scan(&userTickets)
		fmt.Printf("Thank you %v %v for booking your %v tickets. You will receive a confirmation email at %v \n", userName, lastName, userTickets, emailAddress)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName+" "+lastName)
		fmt.Printf("We have a total of %v tickets left\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("Here are all of the bookings: %v\n", firstNames)
	}
}
