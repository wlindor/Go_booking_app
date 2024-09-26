package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and only %v tickets left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}
