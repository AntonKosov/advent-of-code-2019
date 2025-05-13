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

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	computer := read(reader)
	fmt.Fprint(writer, computer.Run())
}

func read(reader io.Reader) *Computer {
	lines := input.Lines(reader)

	program := slice.Map(
		strings.Split(lines[0], ","),
		func(num string) int { return transform.StrToInt(num) },
	)

	return &Computer{
		inputInstruction: 1,
		program:          program,
		position:         0,
	}
}

type Computer struct {
	inputInstruction int
	program          []int
	position         int
}

func (c *Computer) Run() int {
	var lastOutput *int
	type commandFunc func() (length int, output *int, halt bool)
	commands := map[int]commandFunc{
		1:  c.sumCommand,
		2:  c.mulCommand,
		3:  c.inputCommand,
		4:  c.outputCommand,
		99: c.haltCommand,
	}

	for {
		opcode := c.program[c.position] % 100
		length, output, halt := commands[opcode]()
		c.position += length
		if halt {
			return *lastOutput
		}
		if output != nil {
			lastOutput = output
		}
	}
}

func (c *Computer) sumCommand() (length int, output *int, halt bool) {
	c.set(3, c.get(1)+c.get(2))
	return 4, nil, false
}

func (c *Computer) mulCommand() (length int, output *int, halt bool) {
	c.set(3, c.get(1)*c.get(2))
	return 4, nil, false
}

func (c *Computer) inputCommand() (length int, output *int, halt bool) {
	c.set(1, c.inputInstruction)
	return 2, nil, false
}

func (c *Computer) outputCommand() (length int, output *int, halt bool) {
	value := c.get(1)
	return 2, &value, false
}

func (c *Computer) haltCommand() (length int, output *int, halt bool) {
	return 1, nil, true
}

func (c *Computer) get(paramNum int) int {
	idx := c.position + paramNum
	if c.paramValue(paramNum) == 1 {
		return c.program[idx]
	}

	return c.program[c.program[idx]]
}

func (c *Computer) set(paramNum, value int) {
	idx := c.position + paramNum
	if c.paramValue(paramNum) == 0 {
		idx = c.program[idx]
	}

	c.program[idx] = value
}

func (c *Computer) paramValue(paramNum int) int {
	return (c.program[c.position] / math.Pow(10, uint(paramNum)+1)) % 10
}
