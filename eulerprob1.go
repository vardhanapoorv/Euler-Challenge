package main

import (
	"fmt"
)

func multiplesof3and5(val int) int {
	sum := 0
	for i := 0; i < val; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(multiplesof3and5(49))
}
