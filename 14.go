package main

import "fmt"

func main() {
	fmt.Println(maxCollatzSequenceLengthUpto(1000000))
}

func maxCollatzSequenceLengthUpto(n int) (int, int) {
	collatzTable := make([]int, n+1)
	num := 0
	i := 2
	maxLength := 0

	for ; i < n+1; i++ {
		if collatzTable[i] == 0 {
			//fmt.Println(i)
			j := i
			le := 1
			for j != 1 {
				if j%2 == 0 {
					j /= 2
					le += 1
					if j <= n {
						if collatzTable[j] != 0 {
							collatzTable[i] = collatzTable[j] + le - 1
							break
						}
					}
				} else {
					j = 3*j + 1
					le += 1
					if j <= n {
						if collatzTable[j] != 0 {
							collatzTable[i] = collatzTable[j] + le - 1
							break
						}
					}
				}
			}
			if collatzTable[i] == 0 {
				collatzTable[i] = le
			}
			if collatzTable[i] > maxLength {
				num = i
				maxLength = collatzTable[i]
			}
		}
	}

	//for w, i := range collatzTable {
	//fmt.Println(w, i)
	//}
	return num, maxLength
}
