package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

const (
	tileEmpty = '.'
	maxLength = 20
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	code := program.Parse(reader)
	answer := process(code)
	fmt.Fprint(writer, answer)
}

func process(code []int64) int64 {
	readCode := make([]int64, len(code))
	copy(readCode, code)
	scaffold := readScaffold(readCode)
	actions := buildActions(scaffold)

	code[0] = 2 // Set the program to use the robot

	return notify(code, splitActions(actions))
}

func splitActions(actions []string) []string {
	var programA Program
	for aIdx, aAction := range actions {
		if !programA.Add(aAction) {
			break
		}
		bStartIdx := aIdx + 1
		for {
			l, ok := programA.Matches(actions[bStartIdx:])
			if !ok {
				break
			}
			bStartIdx += l
		}
		var programB Program
		for bIdx := bStartIdx; bIdx < len(actions); bIdx++ {
			if !programB.Add(actions[bIdx]) {
				break
			}
			cStartIdx := bIdx + 1
			for {
				l, ok := programA.Matches(actions[cStartIdx:])
				if !ok {
					l, ok = programB.Matches(actions[cStartIdx:])
					if !ok {
						break
					}
				}
				cStartIdx += l
			}
			var programC Program
			for cIdx := cStartIdx; cIdx < len(actions); cIdx++ {
				if !programC.Add(actions[cIdx]) {
					break
				}
				if s, ok := getSequence(&programA, &programB, &programC, actions); ok {
					return []string{
						s.String(), programA.String(), programB.String(), programC.String(), "n",
					}
				}
			}
		}
	}

	panic("cannot split actions into A, B, and C programs")
}

func getSequence(a, b, c *Program, actions []string) (*Program, bool) {
	var sequence Program
	applyProgram := func(p *Program, name string) bool {
		if length, ok := p.Matches(actions); ok {
			sequence.Add(name)
			actions = actions[length:]
			return true
		}

		return false
	}
	for len(actions) > 0 {
		if !applyProgram(a, "A") && !applyProgram(b, "B") && !applyProgram(c, "C") {
			return nil, false
		}
	}

	return &sequence, true
}

type Program struct {
	length  int
	actions []string
}

func (p *Program) Add(action string) bool {
	newLength := p.length + len(action)
	if len(p.actions) > 0 {
		newLength++
	}

	if newLength > maxLength {
		return false
	}

	p.length = newLength
	p.actions = append(p.actions, action)

	return true
}

func (p *Program) Matches(actions []string) (int, bool) {
	if len(actions) < len(p.actions) {
		return 0, false
	}

	for i, a := range p.actions {
		if actions[i] != a {
			return 0, false
		}
	}

	return len(p.actions), true
}

func (p *Program) String() string {
	return strings.Join(p.actions, ",")
}

func buildActions(scaffold [][]rune) []string {
	var actions []string
	pos, dir := findStart(scaffold)
	canMove := func(dir math.Vector2[int]) bool {
		nextPos := pos.Add(dir)
		return nextPos.X >= 0 && nextPos.Y >= 0 && nextPos.Y < len(scaffold) && nextPos.X < len(scaffold[nextPos.Y]) &&
			scaffold[nextPos.Y][nextPos.X] != tileEmpty
	}
	addAction := func(action string) {
		actions = append(actions, action)
	}
	rotate := func() bool {
		if l := dir.RotateLeft(); canMove(l) {
			dir = l
			addAction("L")
			return true
		}

		if r := dir.RotateRight(); canMove(r) {
			dir = r
			addAction("R")
			return true
		}

		return false
	}

	if !canMove(dir) {
		if !rotate() {
			panic("cannot move in any direction")
		}
	}

	for {
		stepCount := 0
		for canMove(dir) {
			pos = pos.Add(dir)
			stepCount++
		}

		addAction(fmt.Sprintf("%v", stepCount))

		if !rotate() {
			break
		}
	}

	return actions
}

func findStart(scaffold [][]rune) (pos, dir math.Vector2[int]) {
	robotTile := map[rune]bool{'^': true, '>': true, 'v': true, '<': true}
	var tile2dir = map[rune]math.Vector2[int]{
		'^': {X: 0, Y: -1},
		'>': {X: 1, Y: 0},
		'v': {X: 0, Y: 1},
		'<': {X: -1, Y: 0},
	}

	for y, row := range scaffold {
		for x, tile := range row {
			if robotTile[tile] {
				return math.NewVector2(x, y), tile2dir[tile]
			}
		}
	}

	panic("robot not found")
}

func notify(code []int64, robotInput []string) int64 {
	input := make(chan int64)
	defer close(input)
	go func() {
		for _, v := range strings.Join(robotInput, "\n") {
			input <- int64(v)
		}
		input <- '\n'
	}()

	output := make(chan int64)
	defer close(output)

	ans := make(chan int64, 1)

	program.Run(
		context.Background(),
		code,
		func() int64 { return <-input },
		func(v int64) {
			if v > 255 {
				ans <- v
			}
		},
	)

	return <-ans
}

func readScaffold(code []int64) [][]rune {
	var sb strings.Builder
	program.Run(context.Background(), code, nil, func(v int64) { sb.WriteByte(byte(v)) })

	scaffold := slice.Map(strings.Split(sb.String(), "\n"), func(line string) []rune {
		return []rune(line)
	})

	// remove the last empty lines
	return slice.Filter(scaffold, func(line []rune) bool { return len(line) > 0 })
}
