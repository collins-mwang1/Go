package main

import (
	"testing"
)

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
	}{
		{name: "Valid Phone", phone: "0768510821"},
		{name: "incomplete", phone: "076850"},
		{name: "Space", phone: "          "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err != nil {
				t.Errorf("Validation failed for '%s': %v", tt.phone, err)
			}
		})
	}
}
