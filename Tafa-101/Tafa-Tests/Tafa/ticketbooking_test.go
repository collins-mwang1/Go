package main

import (
	"testing"
)

func bookingTikos(totalTikos *int, userTikos *int, remainingTikos *int) {
	*remainingTikos = *totalTikos - *userTikos

}
func TestBookingProcess(t *testing.T) {
	var totalTikos int = 300
	var remainingTikos int
	userTiko := int(27)
	bookingTikos(&totalTikos, &userTiko, &remainingTikos)
	if userTiko > totalTikos {
		t.Errorf("we have %v tickets kindly try a lower value", totalTikos)
		return
	}
	if remainingTikos != 30 {
		t.Errorf("Expected Tickets to be 30, but Got %d", remainingTikos)
	}
}
