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
	for T > 0 {
		var N int
		count := 1
		scanf("%d\n", &N)
		for i := 1; ; i++ {
			for j := 2; j <= N; j++ {
				if i%j != 0 {
					break
				} else {
					count++
				}
			}
			if count == N {
				fmt.Println(i)
				break
			}
			count = 1
		}
		T--
	}
}
