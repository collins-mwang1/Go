package main

import "fmt"

func main() {

	i := -7
	for i <= 5 {
		fmt.Println(i)
		i = i + 1

	}
	for j := 0; j < 6; j++ {
		fmt.Println(j)

	}
	for i = range 3 {
		fmt.Println("range", i)

	}
}
