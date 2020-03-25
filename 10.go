package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumOfPrimes())
}

func SumOfPrimes() int {
	x := 2000000
	var pr [2000000]bool
	sum := 0
	i := 2
	for ; i*i < x; i++ {
		if pr[i] == false {
			sum += i
			for j := 2 * i; j < x; j += i {
				pr[j] = true
			}
		}
	}

	for ; i < x; i++ {
		if pr[i] == false {
			sum += i
		}
	}
	return sum
}
