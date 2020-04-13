package main

import "fmt"

func main() {
	prod := 1
	for i := 2; i <= 20; i++ {
		prod = (i * prod) / (findGCD(i, prod))
	}
	fmt.Println(prod)
}

func findGCD(x, y int) int {
	if x == 0 {
		return y
	} else {
		return findGCD(y%x, x)
	}
}
