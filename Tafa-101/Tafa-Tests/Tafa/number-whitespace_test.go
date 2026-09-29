package main

import (
	"fmt"
	"strings"
)

func ValidatePhone(phone string) error {
	cleanPhone := strings.TrimSpace(phone)
	if len(cleanPhone) < 10 {
		return fmt.Errorf("phone number must be at least 10 digits long")
	}
	return nil
}
