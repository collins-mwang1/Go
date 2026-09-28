package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookings(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/bookings", nil)
	rr := httptest.NewRecorder()
	handleBookings(rr, request)
	if rr.Code != http.StatusOK {
		t.Errorf("expected '200' but got %d", rr.Code)

	}

}
