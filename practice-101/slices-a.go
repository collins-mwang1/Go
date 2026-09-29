package main

import (
	"fmt"
)

func main() {
	var participantNames []string
	var ages []int
	numberOfParticipants := 3
	validAge := 18

	fmt.Println("Welcome To GT Hackathon")
	fmt.Println("Kindly Register with your name.")
	fmt.Println("All Participants must be 18 and above.")

	for {
		var name string
		var age int

		fmt.Print("\nEnter your name: ")
		fmt.Scan(&name)

		fmt.Print("Enter your age: ")
		fmt.Scan(&age)

		if age < validAge {
			fmt.Println("You are below the age limit!")
			continue
		}

		participantNames = append(participantNames, name)
		ages = append(ages, age)
		fmt.Println("Thank you for registering,", name)

		remainingSlots := numberOfParticipants - len(participantNames)
		fmt.Printf("Slots remaining: %d\n", remainingSlots)

		if remainingSlots == 0 {
			fmt.Println("\nHackathon is full!")
			break
		}
	}

	fmt.Println("\nRegistered Participants:", participantNames)
}
