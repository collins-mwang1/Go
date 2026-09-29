package main

import (
	"testing"
)

func TestPointersinTafa(t *testing.T) {
	user := createUserData("Roy", "012457485", 21)
	if user == nil {
		t.Fatal("Expected pointer to userData, but got nil")

	}

}
