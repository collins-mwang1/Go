package main

import "fmt"

func main() {
	var a [6]int
	a[5] = 50
	fmt.Println("set:", a)
	fmt.Println("get:", a[5])

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two d:", twoD)

}
