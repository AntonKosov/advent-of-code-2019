package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

type Tile rune
type Move int
type Status int

const (
	moveNorth Move = 1
	moveSouth Move = 2
	moveWest  Move = 3
	moveEast  Move = 4

	statusWall  Status = 0
	statusMoved Status = 1
	statusFound Status = 2

	tileEmpty  Tile = '.'
	tileWall   Tile = '#'
	tileTarget Tile = 'X'
)

var dirs []math.Vector2[int]
var dir2move map[math.Vector2[int]]Move
var status2tile map[Status]Tile

func init() {
	north := math.NewVector2(0, -1)
	east := math.NewVector2(1, 0)
	south := math.NewVector2(0, 1)
	west := math.NewVector2(-1, 0)
	dirs = []math.Vector2[int]{north, east, south, west}
	dir2move = map[math.Vector2[int]]Move{
		north: moveNorth,
		east:  moveEast,
		south: moveSouth,
		west:  moveWest,
	}
	status2tile = map[Status]Tile{
		statusWall:  tileWall,
		statusMoved: tileEmpty,
		statusFound: tileTarget,
	}
}

type Plan map[math.Vector2[int]]Tile

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := program.Parse(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func process(code []int64) int {
	input := make(chan int64)
	defer close(input)

	output := make(chan int64)
	defer close(output)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		program.Run(ctx, code, input, output)
	}()

	startPos := math.NewVector2(0, 0)
	plan := Plan{startPos: tileEmpty}
	move := func(dir math.Vector2[int]) Status {
		input <- int64(dir2move[dir])
		return Status(<-output)
	}
	drawPlan(plan, startPos, move)

	cancel()
	wg.Wait()

	return fillWithOxygen(plan, findOxygenPosition(plan))
}

func drawPlan(plan Plan, pos math.Vector2[int], move func(dir math.Vector2[int]) Status) {
	for _, dir := range dirs {
		newPos := pos.Add(dir)
		if _, ok := plan[newPos]; ok {
			continue
		}

		status := move(dir)
		plan[newPos] = status2tile[status]
		if status == statusWall {
			continue
		}
		drawPlan(plan, newPos, move)

		move(dir.Mul(-1))
	}
}

func findOxygenPosition(plan Plan) math.Vector2[int] {
	for pos, tile := range plan {
		if tile == tileTarget {
			return pos
		}
	}

	panic("target location not found")
}

func fillWithOxygen(plan Plan, startPos math.Vector2[int]) int {
	minutes := 0
	positions := []math.Vector2[int]{startPos}
	for len(positions) > 0 {
		nextPositions := make([]math.Vector2[int], 0, len(positions)*len(dirs))
		for _, pos := range positions {
			for _, dir := range dirs {
				newPos := pos.Add(dir)
				if tile, ok := plan[newPos]; ok && tile == tileEmpty {
					plan[newPos] = tileWall
					nextPositions = append(nextPositions, newPos)
				}
			}
		}
		if len(nextPositions) > 0 {
			minutes++
		}
		positions = nextPositions
	}

	return minutes
}
