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
		sum += calc(mass)
	}

	return sum
}

func calc(mass int) int {
	if mass <= 0 {
		return 0
	}

	fuel := max(0, mass/3-2)

	return fuel + calc(fuel)
}
