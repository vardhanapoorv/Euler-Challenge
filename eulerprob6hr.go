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
		absVal := int(math.Pow(float64((N*(N+1))/2), float64(2))) - (N*(N+1)*(2*N+1))/6
		fmt.Println(absVal)
		T--
	}
}
