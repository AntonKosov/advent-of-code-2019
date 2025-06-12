package main

import (
	"fmt"
	"io"
	"os"
	"sort"

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

var dirs = []math.Vector2[int]{
	math.NewVector2(0, -1),
	math.NewVector2(0, 1),
	math.NewVector2(1, 0),
	math.NewVector2(-1, 0),
}

type Tile rune

const (
	wallTile    Tile = '#'
	passageTile Tile = '.'
	spaceTile   Tile = ' '
)

type Maze [][]Tile

type Portal struct {
	label string
	entry math.Vector2[int]
	exit  math.Vector2[int]
}

func read(reader io.Reader) Maze {
	lines := input.Lines(reader)

	return slice.Map(lines[:len(lines)-1], func(line string) []Tile { return []Tile(line) })
}

func holePosition(maze Maze) (leftTop, bottomRight math.Vector2[int]) {
	leftTop = math.NewVector2(len(maze[0])/2, len(maze)/2)
	bottomRight = leftTop
	move := func(pos *math.Vector2[int], dir math.Vector2[int]) {
		for {
			np := pos.Add(dir)
			if v := maze[np.Y][np.X]; v == wallTile || v == passageTile {
				return
			}
			*pos = np
		}
	}

	move(&leftTop, math.NewVector2(-1, 0))
	move(&leftTop, math.NewVector2(0, -1))
	move(&bottomRight, math.NewVector2(1, 0))
	move(&bottomRight, math.NewVector2(0, 1))

	return leftTop, bottomRight
}

func scanPortals(maze Maze, firstLetter, secondLetterOffset, dir, exitOffset, entryOffset math.Vector2[int], count int) []Portal {
	var portals []Portal
	for i := range count {
		pos := firstLetter.Add(dir.Mul(i))
		l1 := maze[pos.Y][pos.X]
		if l1 == spaceTile {
			continue
		}
		secondLetterPos := pos.Add(secondLetterOffset)
		l2 := maze[secondLetterPos.Y][secondLetterPos.X]
		if l2 == spaceTile {
			continue
		}
		entryPos := pos.Add(entryOffset)
		portals = append(portals, Portal{
			label: string(l1) + string(l2),
			entry: entryPos,
			exit:  pos.Add(exitOffset),
		})
	}

	return portals
}

func parsePortals(maze Maze) []Portal {
	holeLeftTop, holeBottomRight := holePosition(maze)
	holeWidth, holeHeight := holeBottomRight.Sub(holeLeftTop).X+1, holeBottomRight.Sub(holeLeftTop).Y+1
	zero := math.NewVector2(0, 0)
	up := math.NewVector2(0, -1)
	down := math.NewVector2(0, 1)
	left := math.NewVector2(-1, 0)
	right := math.NewVector2(1, 0)
	width, height := len(maze[0])-4, len(maze)-4
	// outer top edge
	portals := scanPortals(maze, math.NewVector2(2, 0), down, right, down.Mul(2), down, width)
	// outer left edge
	portals = append(portals, scanPortals(maze, math.NewVector2(0, 2), right, down, right.Mul(2), right, height)...)
	// outer bottom edge
	portals = append(portals, scanPortals(maze, math.NewVector2(2, len(maze)-2), down, right, up, zero, width)...)
	// outer right edge
	portals = append(portals, scanPortals(maze, math.NewVector2(len(maze[0])-2, 2), right, down, left, zero, height)...)
	// inner top edge
	portals = append(portals, scanPortals(maze, holeLeftTop, down, right, up, zero, holeWidth)...)
	// inner left edge
	portals = append(portals, scanPortals(maze, holeLeftTop, right, down, left, zero, holeHeight)...)
	// inner bottom edge
	portals = append(portals, scanPortals(maze, math.NewVector2(holeLeftTop.X, holeBottomRight.Y-1), down, right, down.Mul(2), down, holeWidth)...)
	// inner right edge
	portals = append(portals, scanPortals(maze, math.NewVector2(holeBottomRight.X-1, holeLeftTop.Y), right, down, right.Mul(2), right, holeHeight)...)

	return portals
}

func parseTeleports(portals []Portal) (idx map[math.Vector2[int]]math.Vector2[int], entry, exit Portal) {
	sort.Slice(portals, func(i, j int) bool { return portals[i].label < portals[j].label })
	idx = make(map[math.Vector2[int]]math.Vector2[int], len(portals)-2)
	for i := 1; i < len(portals)-1; i += 2 {
		p1, p2 := portals[i], portals[i+1]
		idx[p1.entry] = p2.exit
		idx[p2.entry] = p1.exit
	}

	entry = portals[0]
	exit = portals[len(portals)-1]

	return
}

func process(maze Maze) int {
	portals := parsePortals(maze)
	teleports, entryPortal, exitPortal := parseTeleports(portals)
	steps := 0
	maze[entryPortal.entry.Y][entryPortal.entry.X] = wallTile
	positions := []math.Vector2[int]{entryPortal.exit}
	visited := map[math.Vector2[int]]bool{}
	for len(positions) > 0 {
		steps++
		nextPositions := make([]math.Vector2[int], 0, len(positions)*2)
		for _, position := range positions {
			for _, dir := range dirs {
				position := position.Add(dir)
				if position == exitPortal.exit {
					return steps
				}

				tile := maze[position.Y][position.X]
				if tile == wallTile || visited[position] {
					continue
				}

				if tile != passageTile {
					position = teleports[position]
					if visited[position] {
						continue
					}
				}

				visited[position] = true
				nextPositions = append(nextPositions, position)
			}
		}

		positions = nextPositions
	}

	panic("path not found")
}
