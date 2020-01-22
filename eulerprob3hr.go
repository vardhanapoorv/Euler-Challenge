package main

import (
	"bufio"
	"fmt"
	"math"
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
	for T > 0 {
		var N int
		scanf("%d\n", &N)
		var highest int
		for N%2 == 0 {
			highest = 2
			N = N / 2
		}

		for i := 3; i <= int(math.Sqrt(float64(N))); i = i + 2 {
			for N%i == 0 {
				highest = i
				N = N / i
			}
		}

		if N > 2 {
			highest = N
		}
		fmt.Println(highest)
		T--
	}
}
