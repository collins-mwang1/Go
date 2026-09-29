package main

import (
	"fmt"
	"testing"
)

func TestIntergers(t *testing.T) {
	sum := add(2, 2)
	expected := 4

	if sum == expected {
		fmt.Print("Test passed can add")
	} else if sum != expected {
		t.Errorf("expected '%v' but Got %v", expected, sum)

	}

}
