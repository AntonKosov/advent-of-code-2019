package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
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

func read(reader io.Reader) []string {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	return lines
}

func process(_ []string) int {
	return -1
}
