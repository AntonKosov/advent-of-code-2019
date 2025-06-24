package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
)

const (
	side    = 5
	minutes = 200
)

func main() {
	run(os.Stdin, os.Stdout, minutes)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer, mins int) {
	inputData := read(reader)
	answer := process(inputData, mins)
	fmt.Fprint(writer, answer)
}

func posToIdx(x, y int) int {
	return y*side + x
}

type Area uint32

func (a Area) Infest(x, y int) Area {
	if x < 0 || y < 0 || x >= side || y >= side {
		panic("incorrect position")
	}

	return a | (1 << posToIdx(x, y))
}

func (a Area) Infested(x, y int) bool {
	return a&(1<<posToIdx(x, y)) != 0
}

func (a Area) infestedTileCount(x, y int) int {
	if a.Infested(x, y) {
		return 1
	}

	return 0
}

func (a Area) infestedRowCount(y int) int {
	count := 0
	for x := range side {
		count += a.infestedTileCount(x, y)
	}

	return count
}

func (a Area) infestedColumnCount(x int) int {
	count := 0
	for y := range side {
		count += a.infestedTileCount(x, y)
	}

	return count
}

func (a Area) InfestedAround(x, y int, lowerArea, higherArea Area) int {
	count := 0

	// top
	switch {
	case y == 0:
		count += higherArea.infestedTileCount(2, 1)
	case y == 3 && x == 2:
		count += lowerArea.infestedRowCount(4)
	default:
		count += a.infestedTileCount(x, y-1)
	}

	// bottom
	switch {
	case y == 4:
		count += higherArea.infestedTileCount(2, 3)
	case y == 1 && x == 2:
		count += lowerArea.infestedRowCount(0)
	default:
		count += a.infestedTileCount(x, y+1)
	}

	// left
	switch {
	case x == 0:
		count += higherArea.infestedTileCount(1, 2)
	case x == 3 && y == 2:
		count += lowerArea.infestedColumnCount(4)
	default:
		count += a.infestedTileCount(x-1, y)
	}

	// right
	switch {
	case x == 4:
		count += higherArea.infestedTileCount(3, 2)
	case x == 1 && y == 2:
		count += lowerArea.infestedColumnCount(0)
	default:
		count += a.infestedTileCount(x+1, y)
	}

	return count
}

func (a Area) CountInfestedTiles() int {
	return math.BitsCount(uint32(a))
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

func updatedArea(currentArea, lowerArea, higherArea Area) (area Area) {
	for y := range side {
		for x := range side {
			if x == 2 && y == 2 {
				continue
			}
			if ia := currentArea.InfestedAround(x, y, lowerArea, higherArea); currentArea.Infested(x, y) {
				if ia != 1 {
					continue
				}
			} else if ia == 0 || ia > 2 {
				continue
			}
			area = area.Infest(x, y)
		}
	}

	return area
}

func process(area Area, mins int) int {
	areas := map[int]Area{0: area}
	minLevel, maxLevel := 0, 0
	for range mins {
		nextAreas := make(map[int]Area, len(areas)+2)
		nextMinLevel, nextMaxLevel := minLevel, maxLevel
		for lvl := minLevel - 1; lvl <= maxLevel+1; lvl++ {
			area := updatedArea(areas[lvl], areas[lvl-1], areas[lvl+1])
			if area.CountInfestedTiles() == 0 {
				continue
			}
			nextAreas[lvl] = area
			nextMinLevel, nextMaxLevel = min(nextMinLevel, lvl), max(nextMaxLevel, lvl)
		}

		areas = nextAreas
		minLevel, maxLevel = nextMinLevel, nextMaxLevel
	}

	infestedTiles := 0
	for _, area := range areas {
		infestedTiles += area.CountInfestedTiles()
	}

	return infestedTiles
}
