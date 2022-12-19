package main

import "fmt"

var s string

func main() {
	sum := 0
	fmt.Scanln(&s)
	for i := 0; i < len(s); i++ {
		sum = sum + int(s[i])
	}
	if sum > 650 {
		fmt.Println("It is Tasty!")
	} else {
		fmt.Println("Ohh.. No :(")
	}
}
