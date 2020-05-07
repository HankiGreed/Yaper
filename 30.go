package main

import (
	"fmt"
	"math"
)

func main() {
	for i := int32(2); i < math.MaxInt32; i++ {
		if checkFifthPowerSum(i) == true {
			fmt.Println(i)
		}
	}
}

func checkFifthPowerSum(n int32) bool {
	copyOfN := n

	for ; copyOfN > 0 && n >= 0; copyOfN /= 10 {
		n -= int32(math.Pow(float64(copyOfN%10), float64(5)))
	}
	if copyOfN == 0 && n == 0 {
		return true
	}

	return false
}
