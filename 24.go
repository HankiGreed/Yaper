package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(nthDigitPermutation([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1000000-1))
}

func nthDigitPermutation(array []int, n int) []int {
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	//fmt.Println(array, n)

	if n == 0 {
		return array
	}
	noOfDigits := len(array)
	noOfPerms := returnFactorial(noOfDigits - 1)
	firstdig := []int{array[n/noOfPerms]}

	array[n/noOfPerms] = array[len(array)-1]
	array = array[:len(array)-1]

	return append(firstdig, nthDigitPermutation(array, n%noOfPerms)...)
}

func returnFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * returnFactorial(n-1)
}
func arrayToString(array []int) string {
	s := ""
	for _, i := range array {
		s += strconv.Itoa(i)
	}
	return s
}
