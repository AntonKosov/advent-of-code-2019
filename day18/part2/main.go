package main

import (
	"fmt"
	"io"
	stdmath "math"
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

type Tile rune

func (t Tile) Key() bool {
	return t >= 'a' && t <= 'z'
}

const (
	entryTile   Tile = '@'
	wallTile    Tile = '#'
	passageTile Tile = '.'
)

type Maze [][]Tile

func read(reader io.Reader) Maze {
	lines := input.Lines(reader)
	return slice.Map(lines[:len(lines)-1], func(line string) []Tile { return []Tile(line) })
}

var dirs = []math.Vector2[int]{
	math.NewVector2(0, -1),
	math.NewVector2(1, 0),
	math.NewVector2(0, 1),
	math.NewVector2(-1, 0),
}

type Keys uint32

func (k Keys) Add(i Tile) (Keys, bool) {
	bit := Keys(1 << (i - 'a'))
	if k&bit != 0 {
		return 0, false
	}

	return k | bit, true
}

func (k Keys) Open(i Tile) bool {
	return k&(1<<(i-'A')) != 0
}

func initialState(maze Maze) (positions [4]math.Vector2[int], allKeys Keys) {
	for y, row := range maze {
		for x, tile := range row {
			if tile == entryTile {
				maze[y-1][x] = wallTile
				maze[y+1][x] = wallTile
				maze[y][x-1] = wallTile
				maze[y][x+1] = wallTile
				positions[0] = math.NewVector2(x-1, y-1)
				positions[1] = math.NewVector2(x+1, y-1)
				positions[2] = math.NewVector2(x+1, y+1)
				positions[3] = math.NewVector2(x-1, y+1)
				continue
			}
			if tile.Key() {
				allKeys, _ = allKeys.Add(tile)
			}
		}
	}

	return positions, allKeys
}

type ProcessedKey struct {
	positions     [4]math.Vector2[int]
	collectedKeys Keys
}

func findMinSteps(
	maze Maze, minSteps *int, steps int, startPositions [4]math.Vector2[int], keys, allKeys Keys,
	processed map[ProcessedKey]int,
) {
	if steps >= *minSteps {
		return
	}

	if keys == allKeys {
		*minSteps = steps
		return
	}

	pk := ProcessedKey{positions: startPositions, collectedKeys: keys}
	if s, ok := processed[pk]; ok && s <= steps {
		return
	}
	processed[pk] = steps

	for posIdx, position := range startPositions {
		steps := steps
		visited := map[math.Vector2[int]]bool{position: true}
		positions := []math.Vector2[int]{position}
		for len(positions) > 0 {
			steps++
			nextPositions := make([]math.Vector2[int], 0, len(positions)*2)
			for _, pos := range positions {
				for _, dir := range dirs {
					pos := pos.Add(dir)
					if visited[pos] {
						continue
					}
					visited[pos] = true

					tile := maze[pos.Y][pos.X]
					if tile == wallTile {
						continue
					}

					if tile != passageTile {
						if tile.Key() {
							if nk, added := keys.Add(tile); added {
								positionsCopy := startPositions
								positionsCopy[posIdx] = pos
								findMinSteps(maze, minSteps, steps, positionsCopy, nk, allKeys, processed)
								continue
							}
						} else if !keys.Open(tile) { // door
							continue
						}
					}

					nextPositions = append(nextPositions, pos)
				}

				positions = nextPositions
			}
		}
	}
}

func process(maze Maze) int {
	positions, allKeys := initialState(maze)
	minSteps := stdmath.MaxInt
	findMinSteps(maze, &minSteps, 0, positions, Keys(0), allKeys, map[ProcessedKey]int{})

	return minSteps
}
