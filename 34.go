package main

import (
	"fmt"
	"math"
)

/*
You can estabilish a upper bound for this program which is 9!*7,Cus 9!*8 also gives a
7 digit no.
*/
func main() {

	for i := 0; i <= math.MaxInt32; i++ {
		if checkDigitFactorialSum(i) == true {
			fmt.Println(i)
		}
	}
}

func checkDigitFactorialSum(n int) bool {
	factorials := [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	sum := 0
	for i := n; i > 0; i /= 10 {
		sum += factorials[(i % 10)]
	}
	return sum == n
}
