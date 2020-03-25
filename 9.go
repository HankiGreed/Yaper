package main

import "fmt"

func main() {
	m, n := 2, 1

	a, b, c := 0, 0, 0
	for m < 300 {
		n = 0
		for n < m {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if a+b+c == 1000 {
				fmt.Println(a * b * c)
				return
			}
			n++
		}
		m++
	}
}
