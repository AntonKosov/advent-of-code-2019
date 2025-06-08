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
	width  = 50
	height = 50
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

func process(code []int64) int {
	count := 0
	for y := range height {
		for x := range width {
			if affectedPoint(code, x, y) {
				count++
			}
		}
	}

	return count
}

func affectedPoint(code []int64, x, y int) bool {
	input := make(chan int64, 2)
	defer close(input)
	input <- int64(x)
	input <- int64(y)

	output := make(chan int64)
	defer close(output)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		program.Run(context.Background(), code, input, output)
	}()

	affected := <-output == 1
	wg.Wait()

	return affected
}
