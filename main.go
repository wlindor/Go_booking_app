package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and only %v tickets left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var emailAddress string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter how many tickets you would like to purchase: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v for booking your %v tickets. You will receive a confirmation email at %v \n", userName, userTickets, emailAddress)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("We have a total of %v tickets and only %v tickets left\n", conferenceTickets, remainingTickets)
}
