package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const whiteColor = 1

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	program := read(reader)
	fmt.Fprint(writer, process(program))
}

func read(reader io.Reader) []int64 {
	return slice.Map(
		strings.Split(input.Lines(reader)[0], ","),
		func(num string) int64 { return transform.StrToInt64(num) },
	)
}

func process(program []int64) string {
	pos := math.NewVector2(0, 0)
	dir := math.NewVector2(0, -1)
	paintedPanels := map[math.Vector2[int]]int64{pos: whiteColor}

	input := func() int64 {
		if color, ok := paintedPanels[pos]; ok {
			return color
		}

		return 0
	}

	output := func() func(int64) {
		paint := true
		return func(value int64) {
			if paint {
				paintedPanels[pos] = value
			} else {
				if value == 0 {
					dir = dir.RotateLeft()
				} else {
					dir = dir.RotateRight()
				}
				pos = pos.Add(dir)
			}
			paint = !paint
		}
	}

	computer := NewComputer(program, input, output())
	computer.Run()

	return panelsToImage(paintedPanels)
}

func panelsToImage(panels map[math.Vector2[int]]int64) string {
	topLeft, bottomRight := panelBounds(panels)
	var sb strings.Builder
	for r := topLeft.Y; r <= bottomRight.Y; r++ {
		if sb.Len() > 0 {
			sb.WriteRune('\n')
		}
		for c := topLeft.X; c <= bottomRight.X; c++ {
			v := ' '
			if panels[math.NewVector2(c, r)] == whiteColor {
				v = '#'
			}
			sb.WriteRune(v)
		}
	}

	return sb.String()
}

func panelBounds(panels map[math.Vector2[int]]int64) (topLeft, bottomRight math.Vector2[int]) {
	for pos := range panels {
		topLeft.X = min(topLeft.X, pos.X)
		topLeft.Y = min(topLeft.Y, pos.Y)
		bottomRight.X = max(bottomRight.X, pos.X)
		bottomRight.Y = max(bottomRight.Y, pos.Y)
	}

	return topLeft, bottomRight
}

type Computer struct {
	program      []int64
	position     int
	input        func() int64
	output       func(int64)
	relativeBase int
	commands     map[int]func() (offset int, halt bool)
}

func NewComputer(program []int64, input func() int64, output func(int64)) *Computer {
	c := Computer{program: program, input: input, output: output}
	c.commands = map[int]func() (offset int, halt bool){
		1:  c.sumCommand,
		2:  c.mulCommand,
		3:  c.inputCommand,
		4:  c.outputCommand,
		5:  c.jumpIfTrueCommand,
		6:  c.jumpIfFalseCommand,
		7:  c.lessThanCommand,
		8:  c.equalsCommand,
		9:  c.adjustRelativeBaseCommand,
		99: c.haltCommand,
	}

	return &c
}

func (c *Computer) Run() {
	for {
		opcode := int(c.program[c.position] % 100)
		offset, halt := c.commands[opcode]()
		c.position += offset
		if halt {
			return
		}
	}
}

func (c *Computer) sumCommand() (offset int, halt bool) {
	c.set(3, c.get(1)+c.get(2))
	return 4, false
}

func (c *Computer) mulCommand() (offset int, halt bool) {
	c.set(3, c.get(1)*c.get(2))
	return 4, false
}

func (c *Computer) inputCommand() (offset int, halt bool) {
	c.set(1, c.input())
	return 2, false
}

func (c *Computer) outputCommand() (offset int, halt bool) {
	c.output(c.get(1))
	return 2, false
}

func (c *Computer) haltCommand() (offset int, halt bool) {
	return 1, true
}

func (c *Computer) jumpIfTrueCommand() (offset int, halt bool) {
	if c.get(1) != 0 {
		return int(c.get(2)) - c.position, false
	}

	return 3, false
}

func (c *Computer) jumpIfFalseCommand() (offset int, halt bool) {
	if c.get(1) == 0 {
		return int(c.get(2)) - c.position, false
	}

	return 3, false
}

func (c *Computer) lessThanCommand() (offset int, halt bool) {
	var value int64
	if c.get(1) < c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4, false
}

func (c *Computer) equalsCommand() (offset int, halt bool) {
	var value int64
	if c.get(1) == c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4, false
}

func (c *Computer) adjustRelativeBaseCommand() (offset int, halt bool) {
	c.relativeBase += int(c.get(1))

	return 2, false
}

func (c *Computer) get(paramNum int) int64 {
	paramPos := c.position + paramNum
	switch pv := c.paramValue(paramNum); pv {
	case 0:
		paramPos = int(c.program[paramPos])
	case 1:
		// do nothing, the value is correct
	case 2:
		paramPos = c.relativeBase + int(c.program[paramPos])
	default:
		panic(fmt.Sprintf("unknown parameter type: %v", pv))
	}

	if paramPos+1 > len(c.program) {
		return 0
	}

	return c.program[paramPos]
}

func (c *Computer) set(paramNum int, value int64) {
	paramPos := c.position + paramNum
	switch pv := c.paramValue(paramNum); pv {
	case 0:
		paramPos = int(c.program[paramPos])
	case 1:
		// do nothing, the value is correct
	case 2:
		paramPos = c.relativeBase + int(c.program[paramPos])
	default:
		panic(fmt.Sprintf("unknown parameter type: %v", pv))
	}

	if expand := paramPos - len(c.program) + 1; expand > 0 {
		c.program = append(c.program, make([]int64, expand)...)
	}

	c.program[paramPos] = value
}

func (c *Computer) paramValue(paramNum int) int {
	return (int(c.program[c.position]) / math.Pow(10, uint(paramNum)+1)) % 10
}
