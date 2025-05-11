package main

import (
	"fmt"
	"io"
	stdmath "math"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	Run(os.Stdin, os.Stdout)
	fmt.Println()
}

func Run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) [2][]section {
	lines := input.Lines(reader)
	makeSections := func(line string) []section {
		parts := strings.Split(line, ",")
		return slice.Map(parts, func(s string) section {
			return section{dir: rune(s[0]), dst: transform.StrToInt(s[1:])}
		})
	}

	return [2][]section{makeSections(lines[0]), makeSections(lines[1])}
}

func process(wires [2][]section) int {
	wire0 := map[math.Vector2[int]]bool{}
	wire0Pos := math.NewVector2(0, 0)
	for _, section := range wires[0] {
		dir := dirs[section.dir]
		for range section.dst {
			wire0Pos = wire0Pos.Add(dir)
			wire0[wire0Pos] = true
		}
	}

	minDst := stdmath.MaxInt
	wire1Pos := math.NewVector2(0, 0)
	for _, section := range wires[1] {
		dir := dirs[section.dir]
		for range section.dst {
			wire1Pos = wire1Pos.Add(dir)
			if !wire0[wire1Pos] {
				continue
			}
			if dst := wire1Pos.ManhattanDst(math.NewVector2(0, 0)); dst > 0 && dst < minDst {
				minDst = dst
			}
		}
	}

	return minDst
}

type section struct {
	dir rune
	dst int
}

var dirs map[rune]math.Vector2[int]

func init() {
	dirs = map[rune]math.Vector2[int]{
		'U': math.NewVector2(0, -1),
		'R': math.NewVector2(1, 0),
		'D': math.NewVector2(0, 1),
		'L': math.NewVector2(-1, 0),
	}
}
