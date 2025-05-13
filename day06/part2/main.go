package main

import (
	"fmt"
	"io"
	"iter"
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

func read(reader io.Reader) map[string]string {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	orbits := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, ")")
		orbits[parts[1]] = parts[0]
	}

	return orbits
}

func process(orbits map[string]string) int {
	myIndex := map[string]int{}
	for dst, obj := range buildPath(orbits, "YOU") {
		myIndex[obj] = dst
	}

	for dst, obj := range buildPath(orbits, "SAN") {
		if myDst, ok := myIndex[obj]; ok {
			return dst + myDst
		}
	}

	panic("path not found")
}

func buildPath(orbits map[string]string, start string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		current := start
		for dst := 0; current != "COM"; dst++ {
			current = orbits[current]
			if !yield(dst, current) {
				return
			}
		}
	}
}
