package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handleInformation(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("expected ststus 200, got %d", rr.Code)

	}
	var response InfoResponse
	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Errorf("failed to get response: %v", err)

	}

	if response.ConferenceName != "Tafa" {
		t.Errorf("expected name to be Tafa but got %s", response.ConferenceName)

	}

}
