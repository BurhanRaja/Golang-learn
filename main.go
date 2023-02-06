package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 100
	var remainingTickets uint = 100

	// Initialize Array
	// var bookings [100]string
	// var bookings = [100]string{}
	// var bookings = [100]string{"Burhan", "Bijal", "Samrath"}
	
	// Initialize Slice
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	fmt.Printf("Hello World to Our %v booking Application.\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket from here to attend.")

	var firstName string
	var lastName string
	var userEmail string
	var userTicket uint

	fmt.Print("Enter your first name: \n")
	fmt.Scan(&firstName) // passing a pointer instead of a userName variable
	fmt.Print("Enter your last name: \n")
	fmt.Scan(&lastName) // passing a pointer instead of a userName variable
	fmt.Print("Enter your email name: \n")
	fmt.Scan(&userEmail) // passing a pointer instead of a userName variable
	fmt.Print("Enter the no of tickets required: \n")
	fmt.Scan(&userTicket) // passing a pointer instead of a userName variable

	remainingTickets -= userTicket
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slic: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("The Slice Type: %T \n", bookings)
	fmt.Printf("The Slice Size: %v \n", len(bookings))
	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("The Array Type: %T \n", bookings)
	// fmt.Printf("The Array Size: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v  \n", firstName, lastName, userTicket, userEmail)
	fmt.Printf("%v tickets are remaining for %v", remainingTickets, conferenceName)

	fmt.Printf("These are all the bookings %v", bookings);
}
