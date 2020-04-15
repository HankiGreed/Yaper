package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(calcDigitSum(calcFactorial(100)))
}

func calcFactorial(n int64) string {
	fact := big.NewInt(1)
	nn := big.NewInt(n)
	for i := 1; i != 0; i = nn.Cmp(big.NewInt(1)) {
		fact.Mul(fact, nn)
		nn.Sub(nn, big.NewInt(1))
	}
	return fact.String()
}

func calcDigitSum(s string) int {
	sum := 0

	for _, i := range s {
		sum += (int(i) - 48)
	}
	return sum
}
