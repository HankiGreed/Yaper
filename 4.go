package main

import "fmt"

func main() {
	v := 999 * 999

	for v >= 0 {
		if checkPalindrome(v) {
			if checkFactors(v) {
				fmt.Println(v)

				break
			} else {
				v--
			}
		} else {
			v--
		}
	}

}

func checkPalindrome(num int) bool {
	var sl []int
	for num > 0 {

		a := num % 10
		sl = append(sl, a)
		num = num / 10
	}

	for i := 0; i < len(sl)/2; i++ {
		if sl[i] == sl[len(sl)-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkFactors(num int) bool {
	for i := 999; i > 99; i-- {
		if (num/i) > 999 && i*i < num {
			break
		}
		if num%i == 0 {
			return true
		}
	}
	return false
}
