package main

import (
	"fmt"
)

func main() {
	a := []int{1}
	b := []int{1}

	for i := 2; ; {
		temp := b
		b = addUsingArrays(a, b)
		i++
		a = temp
		if len(b) >= 1000 {
			fmt.Println(i)
			break
		}
	}
}

func addUsingArrays(a, b []int) []int {
	if len(a) < len(b) {
		a, b = b, a
	}

	result := make([]int, len(a)+1)

	carry := 0
	i := 1
	for ; i <= len(a); i++ {
		temp := 0
		if i <= len(b) {
			temp = a[len(a)-i] + b[len(b)-i] + carry
		} else {
			temp = a[len(a)-i] + carry
		}
		if temp >= 10 {
			carry = temp / 10
		} else {
			carry = 0
		}
		result[len(result)-i] = temp % 10
	}
	if carry != 0 {
		result[len(result)-i] += carry
	} else {
		result = result[1:]
	}
	return result
}
