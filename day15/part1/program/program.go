package program

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func Parse(reader io.Reader) []int64 {
	return slice.Map(
		strings.Split(input.Lines(reader)[0], ","),
		func(num string) int64 { return transform.StrToInt64(num) },
	)
}

type computer struct {
	code         []int64
	position     int
	input        func() int64
	output       func(int64)
	relativeBase int
	commands     map[int]func(context.Context) (offset int)
}

func Run(ctx context.Context, code []int64, input func() int64, output func(int64)) {
	c := computer{code: code, input: input, output: output}
	c.commands = map[int]func(context.Context) (offset int){
		1: c.sumCommand,
		2: c.mulCommand,
		3: c.inputCommand,
		4: c.outputCommand,
		5: c.jumpIfTrueCommand,
		6: c.jumpIfFalseCommand,
		7: c.lessThanCommand,
		8: c.equalsCommand,
		9: c.adjustRelativeBaseCommand,
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			opcode := int(c.code[c.position] % 100)
			if opcode == 99 {
				return
			}
			offset := c.commands[opcode](ctx)
			c.position += offset
		}
	}
}

func (c *computer) sumCommand(context.Context) (offset int) {
	c.set(3, c.get(1)+c.get(2))
	return 4
}

func (c *computer) mulCommand(context.Context) (offset int) {
	c.set(3, c.get(1)*c.get(2))
	return 4
}

func (c *computer) inputCommand(ctx context.Context) (offset int) {
	ch := make(chan struct{})
	go func() {
		c.set(1, c.input())
		close(ch)
	}()
	select {
	case <-ctx.Done():
		return 0
	case <-ch:
		return 2
	}
}

func (c *computer) outputCommand(ctx context.Context) (offset int) {
	ch := make(chan struct{})
	go func() {
		c.output(c.get(1))
		close(ch)
	}()
	select {
	case <-ctx.Done():
		return 0
	case <-ch:
		return 2
	}
}

func (c *computer) jumpIfTrueCommand(context.Context) (offset int) {
	if c.get(1) != 0 {
		return int(c.get(2)) - c.position
	}

	return 3
}

func (c *computer) jumpIfFalseCommand(context.Context) (offset int) {
	if c.get(1) == 0 {
		return int(c.get(2)) - c.position
	}

	return 3
}

func (c *computer) lessThanCommand(context.Context) (offset int) {
	var value int64
	if c.get(1) < c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4
}

func (c *computer) equalsCommand(context.Context) (offset int) {
	var value int64
	if c.get(1) == c.get(2) {
		value = 1
	}

	c.set(3, value)

	return 4
}

func (c *computer) adjustRelativeBaseCommand(context.Context) (offset int) {
	c.relativeBase += int(c.get(1))

	return 2
}

func (c *computer) get(paramNum int) int64 {
	paramPos := c.position + paramNum
	switch pv := c.paramValue(paramNum); pv {
	case 0:
		paramPos = int(c.code[paramPos])
	case 1:
		// do nothing, the value is correct
	case 2:
		paramPos = c.relativeBase + int(c.code[paramPos])
	default:
		panic(fmt.Sprintf("unknown parameter type: %v", pv))
	}

	if paramPos+1 > len(c.code) {
		return 0
	}

	return c.code[paramPos]
}

func (c *computer) set(paramNum int, value int64) {
	paramPos := c.position + paramNum
	switch pv := c.paramValue(paramNum); pv {
	case 0:
		paramPos = int(c.code[paramPos])
	case 1:
		// do nothing, the value is correct
	case 2:
		paramPos = c.relativeBase + int(c.code[paramPos])
	default:
		panic(fmt.Sprintf("unknown parameter type: %v", pv))
	}

	if expand := paramPos - len(c.code) + 1; expand > 0 {
		c.code = append(c.code, make([]int64, expand)...)
	}

	c.code[paramPos] = value
}

func (c *computer) paramValue(paramNum int) int {
	return (int(c.code[c.position]) / math.Pow(10, uint(paramNum)+1)) % 10
}
