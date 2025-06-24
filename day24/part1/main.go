package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
)

const side = 5

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func posToIdx(x, y int) int {
	return y*side + x
}

type Area uint32

func (a Area) Infest(x, y int) Area {
	return a | (1 << posToIdx(x, y))
}

func (a Area) Infested(x, y int) bool {
	if x < 0 || y < 0 || x >= side || y >= side {
		return false
	}

	return a&(1<<posToIdx(x, y)) != 0
}

func (a Area) Rating() int {
	rating := 0
	mul := 1
	for y := range side {
		for x := range side {
			if a.Infested(x, y) {
				rating += mul
			}
			mul = mul << 1
		}
	}

	return rating
}

func (a Area) InfestedAround(x, y int) int {
	check := func(x, y int) int {
		if a.Infested(x, y) {
			return 1
		}
		return 0
	}

	return check(x-1, y) + check(x+1, y) + check(x, y-1) + check(x, y+1)
}

func read(reader io.Reader) Area {
	lines := input.Lines(reader)
	lines = lines[:len(lines)-1]

	var area Area
	for y, line := range lines {
		for x, v := range line {
			if v == '#' {
				area = area.Infest(x, y)
			}
		}
	}

	return area
}

func process(area Area) int {
	areas := map[Area]bool{}
	for {
		var nextArea Area
		for y := range side {
			for x := range side {
				if ia := area.InfestedAround(x, y); area.Infested(x, y) {
					if ia != 1 {
						continue
					}
				} else if ia == 0 || ia > 2 {
					continue
				}
				nextArea = nextArea.Infest(x, y)
			}
		}

		area = nextArea
		if areas[area] {
			return area.Rating()
		}
		areas[area] = true
	}
}
