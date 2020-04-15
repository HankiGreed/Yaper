package main

import "fmt"

func main() {
	numToString := map[int]int{
		1: 3, 2: 3, 3: 5, 4: 4, 5: 4, 6: 3, 7: 5, 8: 5, 9: 4, 10: 3,
		11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8,
		20: 6, 30: 6, 40: 6, 50: 5, 60: 5, 70: 7, 80: 6, 90: 6, 100: 7, 1000: 11}

	sumOfChar := 0
	for i := 1; i <= 1000; i++ {
		if numToString[i] == 0 {
			j := i
			if j/100 >= 1 {
				sumOfChar = sumOfChar + numToString[j/100] + numToString[100] + 3
				fmt.Println(j/100, j%100)
				j %= 100
			}
			sumOfChar += numToString[(j/10)*10]
			sumOfChar += numToString[j%10]
		} else {
			sumOfChar += numToString[i]
		}

	}

	fmt.Println(sumOfChar)

}
