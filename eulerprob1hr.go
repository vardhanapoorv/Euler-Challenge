package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N int
		fmt.Scan(&N)
		var n1, n2, n3, s1, s2, s3 int
		if N%3 == 0 {
			n1 = ((N / 3) - 1)
		} else {
			n1 = N / 3
		}
		if N%5 == 0 {
			n2 = ((N / 5) - 1)
		} else {
			n2 = N / 5
		}
		if N%15 == 0 {
			n3 = ((N / 15) - 1)
		} else {
			n3 = N / 15
		}
		s1 = (3 * n1 * (n1 + 1)) / 2
		s2 = (5 * n2 * (n2 + 1)) / 2
		s3 = (15 * n3 * (n3 + 1)) / 2

		s := s1 + s2 - s3
		fmt.Println(s)
		T--
	}
}
