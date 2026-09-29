package main

import "fmt"

func main() {
	if 22/7 == 0 {
		fmt.Println("Pi is even")
	} else {
		fmt.Println("Pi is complicated")
	}

	if n := 5; n < 0 {
		fmt.Println("Negative number", n)
	} else if n < 10 {
		fmt.Println("This is a 1 digit number", n)

	}

}
