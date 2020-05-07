package main

import "fmt"

func main() {

}

func isDigitCancellingFraction(numerator, denominator int) bool {

}

func findGCD(x, y int) int {
	if x == 0 {
		return y
	} else {
		return findGCD(y%x, x)
	}
}
