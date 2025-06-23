package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

const (
	computersCount    = 50
	natAddress        = 255
	idleComputerPause = time.Millisecond
	kickTimeout       = 10 * time.Millisecond
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := program.Parse(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

type CompIO struct {
	input        chan int64
	outputBuffer []int64
}

func NewCompIO() *CompIO {
	return &CompIO{input: make(chan int64, 100)}
}

func (c *CompIO) ReadInput() int64 {
	select {
	case v := <-c.input:
		return v
	default:
		time.Sleep(idleComputerPause) // give other goroutines opportunity to run
		return -1
	}
}

func (c *CompIO) AddInput(values ...int64) {
	for _, v := range values {
		c.input <- v
	}
}

func (c *CompIO) AddOutput(value int64) []int64 {
	c.outputBuffer = append(c.outputBuffer, value)
	if len(c.outputBuffer) < 3 {
		return nil
	}

	buffer := c.outputBuffer
	c.outputBuffer = nil

	return buffer
}

type Output struct {
	source int
	value  int64
}

func outputHandler(ctx context.Context, compsIO [computersCount]*CompIO, output <-chan Output, ans chan<- int64) {
	nat := math.Vector2[int64]{}
	var prevKickY *int64
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-output:
			compIO := compsIO[v.source]
			buffer := compIO.AddOutput(v.value)
			if buffer == nil {
				continue
			}
			address, x, y := buffer[0], buffer[1], buffer[2]
			if address == natAddress {
				nat = math.NewVector2(x, y)
				continue
			}
			compsIO[int(address)].AddInput(x, y)
		case <-time.After(kickTimeout):
			y := nat.Y
			if prevKickY != nil && *prevKickY == y {
				ans <- y
				return
			}
			prevKickY = &y
			compsIO[0].AddInput(nat.X, y)
		}
	}
}

func startPrograms(ctx context.Context, code []int64, output chan<- Output) (compsIO [computersCount]*CompIO) {
	for address := range computersCount {
		compIO := NewCompIO()
		compIO.AddInput(int64(address))
		compsIO[address] = compIO

		codeCopy := make([]int64, len(code))
		copy(codeCopy, code)

		go program.Run(
			ctx,
			codeCopy,
			compIO.ReadInput,
			func(v int64) { output <- Output{source: address, value: v} },
		)
	}

	return compsIO
}

func process(code []int64) int64 {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	output := make(chan Output, 100)
	compsIO := startPrograms(ctx, code, output)

	ans := make(chan int64, 1)
	go outputHandler(ctx, compsIO, output, ans)

	return <-ans
}
