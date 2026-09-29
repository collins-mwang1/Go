package main

import (
	"fmt"
)

type userData struct {
	userName    string
	phoneNumber string
	userTickets uint
}
type volunteerData struct {
	volunteerName  string
	volunteerEmail string
}

func createUserData(name string, phone string, tickets uint) *userData {
	return &userData{
		userName:    name,
		phoneNumber: phone,
		userTickets: tickets,
	}
}
func bookingProcess(remainingTickets *uint, userTickets uint) {
	*remainingTickets = *remainingTickets - userTickets

}
func createvolunteerData(fullName string, email string) *volunteerData {
	return &volunteerData{
		volunteerName:  fullName,
		volunteerEmail: email,
	}

}
func volunteerOnboarding(volunteersNeeded uint, remainingVolunteers *uint, registeredVolunteers uint) {
	*remainingVolunteers = volunteersNeeded - registeredVolunteers
}
func main() {
	var conferenceName = " Tafa Business Conference"
	const conferenceTickets = 30
	var remainingTickets uint = 30
	var bookings []*userData
	var response uint

	//volunteers
	var volunteersNeeded uint = 7
	var remainingVolunteers uint = 7

	var volunteerName string
	var volunteerEmail string
	var registeredVolunteers uint
	var onboarding []*volunteerData

	fmt.Println("Jambo, Welcome to", conferenceName, "application")
	fmt.Printf("The %v , where we Build and scale\n", conferenceName)

	fmt.Println("would you like to Purchase a ticket or Volunteer at the event")
	fmt.Println("Reply with - 1 - To Purchase tickets or - 2 - to Volunteer")
	fmt.Scan(&response)

	if response == 1 {
		fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingTickets, "remaining tickets")
		fmt.Println(" Get your Tickets here ")
	} else if response == 2 {
		fmt.Println("Thank you for showing Interest we have ", remainingVolunteers, "spots left")
		fmt.Println("Please Enter your  Full name")

		fmt.Scan(&volunteerName)

		fmt.Println("Enter your email")
		fmt.Scan(&volunteerEmail)

		volunteerOnboarding(remainingVolunteers, &registeredVolunteers, volunteersNeeded)
		var volunteerPointer *volunteerData = createvolunteerData(volunteerName, volunteerEmail)

		onboarding = append(onboarding, volunteerPointer)

		fmt.Println("Thank you for Registering we look forward to working with you")
		return

	}

	for {
		var userName string
		var userTickets uint
		var phoneNumber string

		// deatails entry

		fmt.Println("Enter your Name")
		fmt.Scan(&userName)
		fmt.Println("Enter your phone number(e.g. 0712345678):)")
		fmt.Scan(&phoneNumber)

		if len(phoneNumber) < 10 {
			fmt.Println("Invalid phone number, Numbers must be atleast 10 long")
			continue

		}

		fmt.Println("Enter Number of tickets to book")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry but, we only have %v Tickets left\n", remainingTickets)
			continue

		}

		bookingProcess(&remainingTickets, userTickets)

		var userPointer *userData = createUserData(userName, phoneNumber, userTickets)

		bookings = append(bookings, userPointer)

		//pointer info check

		fmt.Println("\n--- Booking Summary ---")
		fmt.Printf("Thank you %v for purchasing %v tickets.\n", userPointer.userName, userPointer.userTickets)

		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		fmt.Println("\n--- All Bookings ---")
		for i, booking := range bookings {

			fmt.Printf("[%d] %v (%v tickets) - Phone: %v\n", i+1, booking.userName, booking.userTickets, booking.phoneNumber)
		}
		fmt.Println("------------------------\n")

		if remainingTickets == 0 {
			fmt.Println("We are fully booked!")
			break
		}
	}
}
