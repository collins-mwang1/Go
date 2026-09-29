package main

import "fmt"

var x []int

func main() {
	x = []int{1, 2, 3}
	x = append(x, 10)
	y := make([]int, 5)
	y = append(y, 5, 11, 23, 5, 6)
	fmt.Println(x)
	println(y)
	fmt.Printf("%v\n", y)

}
