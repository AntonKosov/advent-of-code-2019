package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

const squareSize = 100

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
	for y, minX, maxX := 5, 0, 0; ; y++ {
		for ; !affectedPoint(code, minX, y); minX++ {
		}

		maxX = max(maxX, minX)
		for x := maxX + 1; affectedPoint(code, x, y); x++ {
			maxX = x
		}

		x := maxX - squareSize + 1
		bottomY := y + squareSize - 1
		if x < 0 || !affectedPoint(code, x, bottomY) {
			continue
		}

		for i := x - 1; i >= 0 && affectedPoint(code, i, bottomY); i-- {
			x = i
		}

		return x*10_000 + y
	}
}

func affectedPoint(code []int64, x, y int) bool {
	input := make(chan int64, 2)
	defer close(input)
	input <- int64(x)
	input <- int64(y)

	output := make(chan int64)
	defer close(output)

	go program.Run(
		context.Background(),
		code,
		func() int64 { return <-input },
		func(v int64) { output <- v },
	)

	return <-output == 1
}
