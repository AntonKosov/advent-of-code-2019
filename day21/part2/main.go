package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
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

// @ABCDEFGHI
// ^...^...##
// ^...#^...#
// ^#..^##.^.
// Don't jump:
// ^###^..###.#..##
// ^#.#^.##..
//
// !A || ((!B || !C) && D && H)
const springscript = `NOT B T
NOT C J
OR T J
AND D J
AND H J
NOT A T
OR T J
RUN
`

func process(code []int64) int64 {
	input := make(chan int64, len(springscript))
	defer close(input)
	for _, v := range springscript {
		input <- int64(v)
	}

	output := make(chan int64)
	defer close(output)

	var damage int64

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range output {
			if v > 0xff {
				damage = v
				return
			}
			fmt.Print(string(rune(v)))
		}
	}()

	program.Run(
		context.Background(),
		code,
		func() int64 { return <-input },
		func(v int64) { output <- v },
	)

	wg.Wait()

	return damage
}
