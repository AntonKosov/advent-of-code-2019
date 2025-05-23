package main

import (
	"fmt"
	"io"
	"iter"
	builtinmath "math"
	"os"
	"sort"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
)

const targetAsteroidNumber = 200

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
	stationPos := findBestStationLocation(asteroids)
	var targetAsteroids []math.Vector2[int]
	for i := 1; ; i++ {
		if len(targetAsteroids) == 0 {
			targetAsteroids = visibleAsteroids(asteroids, stationPos)
			sortTargets(targetAsteroids, stationPos)
			if len(targetAsteroids) == 0 {
				panic("asteroids not found")
			}
		}

		currentAsteroid := targetAsteroids[0]
		targetAsteroids = targetAsteroids[1:]

		if i == targetAsteroidNumber {
			return currentAsteroid.X*100 + currentAsteroid.Y
		}

		asteroids[currentAsteroid.Y][currentAsteroid.X] = empty
	}
}

func sortTargets(targets []math.Vector2[int], stationPos math.Vector2[int]) {
	up := math.NewVector2(0, -1)
	angle := func(target math.Vector2[int]) float64 {
		direction := target.Sub(stationPos)
		rad := up.AngleRad(direction)
		if direction.X < 0 {
			rad = 2*builtinmath.Pi - rad
		}

		return rad
	}

	sort.Slice(targets, func(i, j int) bool {
		return angle(targets[i]) < angle(targets[j])
	})
}

func findBestStationLocation(asteroids [][]rune) math.Vector2[int] {
	maxVisibleAsteroids := 0
	var bestStationPos math.Vector2[int]
	for y, row := range asteroids {
		for x, v := range row {
			if v == empty {
				continue
			}
			stationPos := math.NewVector2(x, y)
			visibleAsteroids := visibleAsteroids(asteroids, stationPos)
			if count := len(visibleAsteroids); count > maxVisibleAsteroids {
				maxVisibleAsteroids = count
				bestStationPos = stationPos
			}
		}
	}

	return bestStationPos
}

const (
	empty    = '.'
	asteroid = '#'
)

func visibleAsteroids(asteroids [][]rune, stationPos math.Vector2[int]) []math.Vector2[int] {
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

	var visible []math.Vector2[int]
	for y, row := range asteroids {
		for x, v := range row {
			pos := math.NewVector2(x, y)
			if v == empty || invisible[pos] {
				continue
			}
			visible = append(visible, pos)
		}
	}

	return visible
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
