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
	var T int
	scanf("%d\n", &T)
	maxVal := 105000
	var primes []int
	isNotPrime := make([]bool, maxVal)
	for i := 2; i*i < maxVal; i++ {
		if isNotPrime[i] == false {
			for j := i * i; j < maxVal; j += i {
				isNotPrime[j] = true
			}
		}
	}
	for i := 2; i < maxVal; i++ {
		if isNotPrime[i] == false {
			primes = append(primes, i)
		}
	}

	for T > 0 {
		var N int
		scanf("%d\n", &N)
		fmt.Println(primes[N-1])
		T--
	}
}
