package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	Run(os.Stdin, os.Stdout)
	fmt.Println()
}

func Run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []int {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) int { return transform.StrToInt(line) })
}

func process(masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += mass/3 - 2
	}

	return sum
}
