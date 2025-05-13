package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) map[string][]string {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	orbits := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ")")
		c := parts[0]
		orbits[c] = append(orbits[c], parts[1])
	}

	return orbits
}

func process(orbits map[string][]string) int {
	count := 0
	distance := 0
	current := []string{"COM"}
	for len(current) > 0 {
		count += distance * len(current)
		next := make([]string, 0, len(current))
		for _, center := range current {
			next = append(next, orbits[center]...)
		}

		distance++
		current = next
	}

	return count
}
