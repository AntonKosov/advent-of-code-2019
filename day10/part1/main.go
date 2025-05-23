package main

import (
	"fmt"
	"io"
	"iter"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
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

func read(reader io.Reader) [][]rune {
	lines := input.Lines(reader)

	return slice.Map(lines[:len(lines)-1], func(line string) []rune { return []rune(line) })
}

func process(asteroids [][]rune) int {
	minInvisibleAsteroids := len(asteroids) * len(asteroids[0])
	totalAsteroids := 0
	for y, row := range asteroids {
		for x, v := range row {
			if v == empty {
				continue
			}
			totalAsteroids++
			stationPos := math.NewVector2(x, y)
			invisibleAsteroids := countInisibleAsteroids(asteroids, stationPos)
			minInvisibleAsteroids = min(minInvisibleAsteroids, invisibleAsteroids)
		}
	}

	return totalAsteroids - minInvisibleAsteroids - 1
}

const (
	empty    = '.'
	asteroid = '#'
)

func countInisibleAsteroids(asteroids [][]rune, stationPos math.Vector2[int]) int {
	invisible := map[math.Vector2[int]]bool{}
	validPos := func(pos math.Vector2[int]) bool {
		return pos.X >= 0 && pos.Y >= 0 && pos.X < len(asteroids[0]) && pos.Y < len(asteroids)
	}
	for repeat, r := true, 1; repeat; r++ {
		repeat = false
		for pos := range cellsAround(stationPos, r) {
			if !validPos(pos) {
				continue
			}
			repeat = true
			if invisible[pos] {
				continue
			}
			firstVisible := asteroids[pos.Y][pos.X] == asteroid
			offset := pos.Sub(stationPos)
			for np := pos.Add(offset); validPos(np); np = np.Add(offset) {
				if asteroids[np.Y][np.X] == empty {
					continue
				}
				if !firstVisible {
					firstVisible = true
					continue
				}
				invisible[np] = true
			}
		}
	}

	return len(invisible)
}

func cellsAround(center math.Vector2[int], radius int) iter.Seq[math.Vector2[int]] {
	dirs := []math.Vector2[int]{
		math.NewVector2(1, 0),
		math.NewVector2(0, 1),
		math.NewVector2(-1, 0),
		math.NewVector2(0, -1),
	}

	return func(yield func(math.Vector2[int]) bool) {
		pos := center.Sub(math.NewVector2(radius, radius))
		for i := range radius * 2 * 4 {
			if !yield(pos) {
				return
			}
			dirIdx := i / (radius * 2)
			pos = pos.Add(dirs[dirIdx])
		}
	}
}
