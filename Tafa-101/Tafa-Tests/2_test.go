package main

import (
	"fmt"
	"testing"
)

func TestCreateUserData2(t *testing.T) {
	result := createUserData("paul", "07656500000655", 333)
	if result.userName != "aul" {
		t.Errorf("expected name 'aul' but got %s\n", result.userName)
	}
	if result.phoneNumber == "0765650000655" {
		fmt.Printf("Correct number '%s' registered\n", result.phoneNumber)

	} else {
		fmt.Println("Eiii the program needs work")
	}
	if result.userTickets == 333 {
		fmt.Println("We are cruising Nicely")

	} else {
		fmt.Printf("Eii there's a problem \n We expected 333 we got	' %v '	", result.userTickets)
	}

}
