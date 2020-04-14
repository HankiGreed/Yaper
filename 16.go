package main

import (
	"fmt"
)

func main() {
	fmt.Println(digitSumTwoToThePowerOf(1000))
}

func digitSumTwoToThePowerOf(n int) int {
	digits := make([]int, 500)

	digits[0] = 1
	//m := 1
	carry := 0
	for j := 0; j < n; j++ {
		for i := 0; i < 500; i++ {
			n := digits[i]*2 + carry
			digits[i] = n % 10
			carry = n / 10
		}
	}
	sum := 0
	for _, i := range digits {
		sum += i
	}
	return sum
}
