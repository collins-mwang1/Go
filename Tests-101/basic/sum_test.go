package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	problem := Sum(145, 23)
	answer := 168

	if problem != answer {
		fmt.Println("Incorrect answer. Sum function faulty")
	} else {
		fmt.Println("Correct answer")
	}
}
