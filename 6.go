package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(int(differenceOfSum(100)))
}

func differenceOfSum(x float64) float64 {
	return math.Pow((x*(x+1))/2, 2) - ((x * (x + 1) * (2*x + 1)) / 6)
}
