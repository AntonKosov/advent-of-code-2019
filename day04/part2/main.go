package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	start, end := read(reader)
	answer := process(start, end)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) (start, end int) {
	line := input.Lines(reader)[0]

	return transform.StrToInt(line[:6]), transform.StrToInt(line[7:])
}

func process(start, end int) int {
	count := 0
nextNum:
	for i := start; i <= end; i++ {
		doubleDigits := false
		countEqualDigits := 0
		for num, rigthDigit := i/10, i%10; num != 0; num /= 10 {
			leftDigit := num % 10
			if leftDigit > rigthDigit {
				continue nextNum
			}
			if leftDigit == rigthDigit {
				if countEqualDigits == 0 && num <= 9 {
					doubleDigits = true
					break
				}
				countEqualDigits = max(2, countEqualDigits+1)
				continue
			}
			rigthDigit = leftDigit
			doubleDigits = doubleDigits || countEqualDigits == 2
			countEqualDigits = 0
		}
		if doubleDigits {
			count++
		}
	}

	return count
}
