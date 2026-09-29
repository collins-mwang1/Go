package main

import (
	"fmt"
	"math"
)

const s = "Wonder"

func main() {
	fmt.Println(s)
	const n = 22 / 7
	const b = 50000
	const d = 3e20 / b
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
