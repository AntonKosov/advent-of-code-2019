package main

import (
	"fmt"
	"io"
	"iter"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
)

const phaseCount = 100

var pattern = []int{0, 1, 0, -1}

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []int {
	return slice.Map([]byte(input.Lines(reader)[0]), func(b byte) int {
		return int(b - '0')
	})
}

func process(signal []int) string {
	for range phaseCount {
		signal = genNewSignal(signal)
	}

	return string(slice.Map(signal[:8], func(i int) rune {
		return rune(i + '0')
	}))
}

func genNewSignal(signal []int) []int {
	newSignal := make([]int, len(signal))
	for i := range signal {
		sum := 0
		for j, v := range patternValues(i+1, len(signal)) {
			sum += v * signal[j]
		}

		newSignal[i] = math.Abs(sum) % 10
	}

	return newSignal
}

func patternValues(duplicates, length int) iter.Seq2[int, int] {
	return func(yield func(idx int, value int) bool) {
		first := true
		idx := 0
		for {
			for _, v := range pattern {
				for range duplicates {
					if first {
						first = false
						continue
					}
					if !yield(idx, v) {
						return
					}
					idx++
					if idx == length {
						return
					}
				}
			}
		}
	}
}
