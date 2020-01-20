package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N int
		var fibTerm int
		fib1 := 1
		sumEven := 2
		fib2 := 2
		fmt.Scan(&N)
		for fibTerm <= N {
			fibTerm = fib1 + fib2
			if fibTerm%2 == 0 && fibTerm <= N {
				sumEven += fibTerm
			}
			fib1 = fib2
			fib2 = fibTerm
		}
		fmt.Println(sumEven)
		T--
	}
}
