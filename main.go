package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 100
	remainingTickets := 100

	fmt.Printf("Hello World to Our %v booking Application.\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket from here to attend.")
	fmt.Print(quote.Go())

	var username string
	var userTicket int

	username = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v ticket", username, userTicket)
}
