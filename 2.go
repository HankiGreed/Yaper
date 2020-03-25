package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 0

	for b < 4000000 {
		d := a
		a = b
		b = d + b
		fmt.Printf("%v %v \n", a, b)
		if a%2 == 0 {
			c += a
			fmt.Println(a)
		}
	}
	if b%2 == 0 {
		c += b
		fmt.Println(b)
	}
	fmt.Println(c)
}
