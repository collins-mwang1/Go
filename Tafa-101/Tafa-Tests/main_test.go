package main

import (
	"testing"
)

func TestCreateUserData(t *testing.T) {
	result := createUserData("collins", "0768510810", 2)
	if result.userName != "collins" {
		t.Errorf("Expected name 'collins', but got '%s'", result.userName)

	}
}
