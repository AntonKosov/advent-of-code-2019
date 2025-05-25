package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const planetsCount = 4

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) [planetsCount]math.Vector3[int] {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]
	if len(lines) != planetsCount {
		panic("incorrect input")
	}

	var planets [planetsCount]math.Vector3[int]
	for i := range planetsCount {
		parts := transform.StrToInts(lines[i])
		planets[i] = math.NewVector3(parts[0], parts[1], parts[2])
	}

	return planets
}

func process(planets [planetsCount]math.Vector3[int]) uint64 {
	return math.LCM(
		uint64(cycleSize([planetsCount]int{planets[0].X, planets[1].X, planets[2].X, planets[3].X})),
		uint64(cycleSize([planetsCount]int{planets[0].Y, planets[1].Y, planets[2].Y, planets[3].Y})),
		uint64(cycleSize([planetsCount]int{planets[0].Z, planets[1].Z, planets[2].Z, planets[3].Z})),
	)
}

func cycleSize(positions [planetsCount]int) int {
	currentPositions := positions
	var velocities [planetsCount]int
	for i := 1; ; i++ {
		for i, cp := range currentPositions {
			for _, ap := range currentPositions {
				velocities[i] += math.Sign(ap - cp)
			}
		}

		for i := range currentPositions {
			currentPositions[i] += velocities[i]
		}

		if velocities == [planetsCount]int{} && positions == currentPositions {
			return i
		}
	}
}
