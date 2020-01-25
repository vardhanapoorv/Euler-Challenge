package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func checkPalindrome(str string) bool {
	length := len(str)
	counter := length - 1
	for i := 0; i < length/2; i++ {
		if str[counter] != str[i] {
			return false
		}
		counter--
	}
	return true
}

func checkFactors3DigitOrNot(num int) bool {
	var threeDigitFactors []int
	for i := 100; i < 1000; i++ {
		if num%i == 0 {
			threeDigitFactors = append(threeDigitFactors, i)
		}
	}

	for i := range threeDigitFactors {
		if int(math.Floor(float64(num/threeDigitFactors[i]))) >= 100 && int(math.Floor(float64(num/threeDigitFactors[i]))) < 1000 {
			return true
		}
	}

	return false
}

func main() {
	defer writer.Flush()
	var T int
	scanf("%d\n", &T)
	for T > 0 {
		var N int
		scanf("%d\n", &N)
		for {
			N--
			str := strconv.Itoa(N)
			if checkPalindrome(str) && checkFactors3DigitOrNot(N) {
				break
			}
		}
		fmt.Println(N)
		T--
	}
}
