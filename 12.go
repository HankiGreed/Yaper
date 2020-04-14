package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1
	//for j := 0; j < 10; j++ {
	//fmt.Println(nthTriangleNo(j))
	//fmt.Println(noOfFactors(j))
	//}
	for ; ; i++ {
		if n := noOfFactors(nthTriangleNo(i)); n > 500 {
			fmt.Println(nthTriangleNo(i), n)
			return
		}
	}
}

func nthTriangleNo(n int) int {
	return (n * (n + 1)) / 2
}

func noOfFactors(n int) int {
	no := 2
	if n == 1 {
		return 1
	} else {
		for i := 2; i < int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				no += 2
			}
		}
		return no
	}
}
