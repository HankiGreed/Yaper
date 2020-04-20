package main

import "fmt"

func main() {
	fmt.Println(returnSpriralSumSquare(500))
}

func returnSpriralSumSquare(n int) int {
	sum := 1
	lastNum := 1

	for i := 1; i <= n; i++ {

		for j := 1; j <= 4; j++ {
			sum += lastNum + (i * 2)
			lastNum += (2 * i)
		}
	}
	return sum
}
