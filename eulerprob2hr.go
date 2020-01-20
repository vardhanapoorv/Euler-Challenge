package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()
	var T uint64
	scanf("%d\n", &T)
	for T > 0 {
		var N uint64
		var fibTerm, sumEven, fib1, fib2 uint64
		fib1 = 0
		sumEven = 2
		fib2 = 2
		scanf("%d\n", &N)
		// E(k) = 4E(k-1) + E(k-2)
		for fibTerm <= N {
			fibTerm = fib1 + 4*fib2
			if fibTerm <= N {
				sumEven += fibTerm
			}
			fib1 = fib2
			fib2 = fibTerm
		}
		fmt.Println(sumEven)
		T--
	}
}
