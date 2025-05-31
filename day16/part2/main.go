package main

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const phaseCount = 100

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
	return slices.Repeat(slice.Map([]byte(input.Lines(reader)[0]), func(b byte) int {
		return int(b - '0')
	}), 10000)
}

func process(signal []int) int {
	offset := parseNum(signal, 7)
	if offset < len(signal)/2 {
		panic("offset is too small")
	}

	signal = signal[offset:]
	for range phaseCount {
		signal = genNewSignal(signal)
	}

	return parseNum(signal, 8)
}

func parseNum(signal []int, count int) int {
	return transform.StrToInt(string(slice.Map(signal[:count], func(i int) rune {
		return rune(i + '0')
	})))
}

func genNewSignal(signal []int) []int {
	sum := 0
	newSignal := make([]int, len(signal))
	for i := len(signal) - 1; i >= 0; i-- {
		sum = (sum + signal[i]) % 10
		newSignal[i] = sum
	}

	return newSignal
}
