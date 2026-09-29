package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostBooking(t *testing.T) {
	bookingData := BookingRequest{
		UserName:    "Magnetto           steward",
		PhoneNumber: "0786510813",
		UserTickets: 130,
	}
	jsonBytes, err := json.Marshal(bookingData)
	if err != nil {
		t.Errorf("failed to encode : %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(jsonBytes))

	rr := httptest.NewRecorder()

	handleBookings(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected 201 created but got %d and a response %s", rr.Code, rr.Body.String())

	}

}
