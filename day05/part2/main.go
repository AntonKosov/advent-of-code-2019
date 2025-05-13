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
	run(os.Stdin, os.Stdout, 5)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer, inputInstruction int) {
	computer := read(reader, inputInstruction)
	fmt.Fprint(writer, computer.Run())
}

func read(reader io.Reader, inputInstruction int) *Computer {
	lines := input.Lines(reader)

	program := slice.Map(
		strings.Split(lines[0], ","),
		func(num string) int { return transform.StrToInt(num) },
	)

	return &Computer{
		inputInstruction: inputInstruction,
		program:          program,
		position:         0,
	}
}

type Computer struct {
	inputInstruction int
	program          []int
	position         int
	lastOutput       int
}

func (c *Computer) Run() int {
	type commandFunc func() (offset int, halt bool)
	commands := map[int]commandFunc{
		1:  c.sumCommand,
		2:  c.mulCommand,
		3:  c.inputCommand,
		4:  c.outputCommand,
		5:  c.jumpIfTrueCommand,
		6:  c.jumpIfFalseCommand,
		7:  c.lessThanCommand,
		8:  c.equalsCommand,
		99: c.haltCommand,
	}

	for {
		opcode := c.program[c.position] % 100
		offset, halt := commands[opcode]()
		c.position += offset
		if halt {
			return c.lastOutput
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
	c.set(1, c.inputInstruction)
	return 2, false
}

func (c *Computer) outputCommand() (offset int, halt bool) {
	c.lastOutput = c.get(1)
	return 2, false
}

func (c *Computer) haltCommand() (offset int, halt bool) {
	return 1, true
}

func (c *Computer) jumpIfTrueCommand() (offset int, halt bool) {
	if c.get(1) != 0 {
		return c.get(2) - c.position, false
	}

	return 3, false
}

func (c *Computer) jumpIfFalseCommand() (offset int, halt bool) {
	if c.get(1) == 0 {
		return c.get(2) - c.position, false
	}

	return 3, false
}

func (c *Computer) lessThanCommand() (offset int, halt bool) {
	value := 0
	if c.get(1) < c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4, false
}

func (c *Computer) equalsCommand() (offset int, halt bool) {
	value := 0
	if c.get(1) == c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4, false
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
