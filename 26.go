package main

import (
	"fmt"
)

func main() {
	max := -1
	maxIndex := 0
	for i := 3; i <= 1000; i += 1 {
		if b, length := returnRecurringSequenceLength(i); b == true {
			if length > max {
				max = length
				maxIndex = i
			}
		}
	}

	fmt.Println(max, maxIndex)
}

func returnRecurringSequenceLength(n int) (bool, int) {

	dividend := 1
	seen := []int{dividend}
	seen = append(seen, dividend)

	div := dividend

	for {
		if div == 0 {
			return false, 0
		}
		div = (div % n) * 10
		//fmt.Println(div)
		if b, index := checkIfExists(div, seen); b == true {
			if len(seen)-index == 0 {

				return false, len(seen) - index
			}
			return true, len(seen) - index
		}
		seen = append(seen, div)
	}

}

func checkIfExists(n int, array []int) (bool, int) {
	for i, item := range array {
		if item == n {
			return true, i
		}
	}
	return false, 0
}
