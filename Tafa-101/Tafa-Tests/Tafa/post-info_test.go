package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	request := httptest.NewRequest(http.MethodPut, "/bookings", nil)
	rr := httptest.NewRecorder()

	handleBookings(rr, request)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("espected 405 error but Got %d", rr.Code)

	}

}
