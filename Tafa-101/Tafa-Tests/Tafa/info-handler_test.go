package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInformationHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handleInformation(rr, request)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("GOT ststus %v, and i wanted %v", status, http.StatusOK)

	}
	var response InfoResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to decode : %v", err)
	}
	if response.RemainingTickets != 56 {
		t.Errorf("Got remaining tickets %d , But wanted 56", response.RemainingTickets)

	}

}
