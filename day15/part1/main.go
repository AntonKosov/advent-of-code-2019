package main

import (
	"context"
	"fmt"
	"io"
	"iter"
	"os"
	"sync"

	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/path"
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

	return distance(plan, startPos, findOxygenPosition(plan))
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

func distance(plan Plan, start, target math.Vector2[int]) int {
	pathToTarget := path.AStar(
		start,
		target,
		func(v1, v2 math.Vector2[int]) bool {
			return v1.ManhattanDst(target) < v2.ManhattanDst(target)
		},
		func(v math.Vector2[int]) iter.Seq[math.Vector2[int]] {
			return func(yield func(math.Vector2[int]) bool) {
				for _, dir := range dirs {
					nextPos := v.Add(dir)
					if tile, ok := plan[nextPos]; ok && tile != tileWall {
						if !yield(nextPos) {
							return
						}
					}
				}
			}
		},
	)

	return len(pathToTarget) - 1
}
