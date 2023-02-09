package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 100
	var remainingTickets uint = 100

	// ? Initialize Array
	// var bookings [100]string
	// var bookings = [100]string{}
	// var bookings = [100]string{"Burhan", "Bijal", "Samrath"}

	// ? Initialize Slice
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	// Greetings to User
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// Infinite For Loop
	for {
		firstName, lastName, userEmail, userTicket := inputUser()

		isValidName, isValidEmail, isValidTicket := validateUserInp(firstName, lastName, userEmail, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookingDone(remainingTickets, userTicket, bookings, firstName, lastName, userEmail, conferenceName)

			// Get First Names
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names o the bookings are: %v\n", firstNames)

			// ? If/Else
			if remainingTickets == 0 {
				fmt.Println("Our Conference is sold out. Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Printf("Please FirstName and LastName should contain at most 2 chars.")
			}
			if !isValidEmail {
				fmt.Printf("Please Enter valid email.")
			}
			if !isValidTicket {
				fmt.Printf("Please Enter correct ticket number.")
			}
		}
	}

	// ? Switch
	// city := "London"
	// switch city {
	// case "London":
	// Logic
	// case "New York":
	// Logic
	// case "Singapore":
	// Login
	// }
}

func greetUser(confName string, confTickets uint, remainTickets uint) {
	fmt.Printf("Hello World to Our %v booking Application.\n", confName)
	fmt.Printf("We have a total of %v Tickets and %v are still available.\n", confTickets, remainTickets)
	fmt.Println("Get your ticket from here to attend.")
}

func getFirstNames(bookings []string) []string {
	// ? For Loop
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInp(firstName string, lastName string, userEmail string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	// ? Validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@.")
	isValidTicket := userTicket <= remainingTickets && userTicket > 0

	return isValidName, isValidEmail, isValidTicket
}

func inputUser() (string, string, string, uint) {
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

	return firstName, lastName, userEmail, userTicket
}

func bookingDone(remainingTickets uint, userTicket uint, bookings []string, firstName string, lastName string, userEmail string, conferenceName string) {
	
	remainingTickets -= userTicket

	// ? Array Based
	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("The Array Type: %T \n", bookings)
	// fmt.Printf("The Array Size: %v \n", len(bookings))

	// ? Slice Based
	bookings = append(bookings, firstName+" "+lastName)
	// fmt.Printf("The whole slic: %v \n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("The Slice Type: %T \n", bookings)
	// fmt.Printf("The Slice Size: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v  \n", firstName, lastName, userTicket, userEmail)
	fmt.Printf("%v tickets are remaining for %v", remainingTickets, conferenceName)
}
