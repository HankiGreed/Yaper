package main

import "fmt"

func main() {
	prod := 1
	gcd  :=20
	for i := 19; i > 0; i-- {
		prod *= i
		gcd = 
	}
	fmt.Println(prod)
	fmt.Println(findGCD(20, 15))
}

func findGCD(x, y int) int {
	for x%y != 0 {
		a := y
		y = x % y
		x = a
	}
	return y
}
