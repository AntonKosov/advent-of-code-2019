package main

import (
	"context"
	"fmt"
	"io"
	"os"

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

	output := make(chan int64, 1)
	defer close(output)

	go program.Run(
		context.Background(),
		code,
		func() int64 { return <-input },
		func(v int64) { output <- v },
	)

	return <-output == 1
}
