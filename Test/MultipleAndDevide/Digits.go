package main

import (
	"fmt"
	"math"
)

var f int

func main() {
	fmt.Scanln(&f)
	fmt.Println(100 / f)
	fmt.Println(math.Sqrt(float64(f)), ": It is square root of included integer")
	fmt.Println(f*f, ": It is Square of included integer")
}
