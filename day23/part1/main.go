package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

const (
	computersCount = 50
	targetAddress  = 255
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

func read(reader io.Reader) []int64 {
	return program.Parse(reader)
}

type CompIO struct {
	inputMutex   sync.Mutex
	outputMutex  sync.Mutex
	input        []int64
	outputBuffer []int64
}

func (c *CompIO) ReadInput() int64 {
	c.inputMutex.Lock()
	defer c.inputMutex.Unlock()

	if len(c.input) == 0 {
		return -1
	}

	v := c.input[0]
	c.input = c.input[1:]

	return v
}

func (c *CompIO) AddInput(values ...int64) {
	c.inputMutex.Lock()
	defer c.inputMutex.Unlock()

	c.input = append(c.input, values...)
}

func (c *CompIO) AddOutput(value int64) []int64 {
	c.outputMutex.Lock()
	defer c.outputMutex.Unlock()

	c.outputBuffer = append(c.outputBuffer, value)
	if len(c.outputBuffer) < 3 {
		return nil
	}

	buffer := c.outputBuffer
	c.outputBuffer = c.outputBuffer[:0]

	return buffer
}

func process(code []int64) int64 {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	compsIO := make(map[int]*CompIO, computersCount)
	for i := range computersCount {
		compsIO[i] = &CompIO{}
	}

	ans := make(chan int64)
	allRunning := make(chan struct{})
	for address := range computersCount {
		codeCopy := make([]int64, len(code))
		copy(codeCopy, code)

		compIO := compsIO[address]
		compIO.AddInput(int64(address))

		go program.Run(
			ctx,
			codeCopy,
			compIO.ReadInput,
			func(v int64) {
				<-allRunning
				buffer := compIO.AddOutput(v)
				if buffer == nil {
					return
				}
				address, x, y := buffer[0], buffer[1], buffer[2]
				if address == targetAddress {
					ans <- y
					return
				}
				targetCompIO := compsIO[int(address)]
				targetCompIO.AddInput(x, y)
			},
		)
	}

	close(allRunning)

	return <-ans
}
