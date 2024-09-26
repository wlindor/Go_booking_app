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
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter how many tickets you would like to purchase: ")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("We have a total of %v tickets and only %v tickets left\n", conferenceTickets, remainingTickets)
}
