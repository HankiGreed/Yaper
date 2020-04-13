package main

import (
	"fmt"
)

func main() {
	N := 10001
	fmt.Println(returnNthPrime(N))
}

func returnNthPrime(n int) int {
	i := 3
	for ; n >= 2; i += 2 {
		if checkIfPrime(i) == true {
			n--
		}
	}
	return i - 2
}

func checkIfPrime(x int) bool {
	if x > 1 {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
