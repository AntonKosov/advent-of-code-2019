package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	lines := input.Lines()
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
