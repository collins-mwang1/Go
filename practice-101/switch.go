package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Print("Write ", i, "as ")
	switch i {
	case 1:
		println(" one")
	case 2:
		println(" two")
	case 3:
		println(" four")
		fallthrough
	case 4:

		fmt.Println(" You won bronze ")

	}
	switch time.Now().Weekday() {
	case time.Friday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 10:
		fmt.Println("It's Morning")
	case t.Hour() < 13:
		fmt.Println("it's Noon")
	case t.Hour() < 16:
		fmt.Println("Its afternoon")
	case t.Hour() < 19:
		fmt.Println("It's Evening")
	default:
		fmt.Println("It's Night")
	}

}
