package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findLargestPrimeFactors(600851475143))
}

func findLargestPrimeFactors(x int) int {

	for x%2 == 0 {
		x /= 2
	}
	maxprime := -1

	for i := 3; i < int(math.Sqrt(float64(x))); i += 2 {
		for x%i == 0 {
			maxprime = i
			x = x / i
		}
	}

	if x > 2 {
		maxprime = x
	}
	return maxprime
}
