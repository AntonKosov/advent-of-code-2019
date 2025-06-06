package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

const tileEmpty = '.'

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	code := program.Parse(reader)
	answer := process(code)
	fmt.Fprint(writer, answer)
}

func process(code []int64) int {
	scaffold := readScaffold(code)

	return alignmentParameters(scaffold)
}

func alignmentParameters(scaffold [][]rune) int {
	var sum int
	for y := 1; y < len(scaffold)-1; y++ {
		for x := 1; x < len(scaffold[y])-1; x++ {
			if scaffold[y][x] == tileEmpty || len(scaffold[y+1]) == 0 {
				continue
			}
			if scaffold[y-1][x] == tileEmpty ||
				scaffold[y+1][x] == tileEmpty ||
				scaffold[y][x-1] == tileEmpty ||
				scaffold[y][x+1] == tileEmpty {
				continue
			}
			sum += x * y
		}
	}

	return sum
}

func readScaffold(code []int64) [][]rune {
	var sb strings.Builder

	output := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for out := range output {
			sb.WriteByte(byte(out))
		}
	}()

	program.Run(context.Background(), code, nil, output)

	close(output)
	wg.Wait()

	return slice.Map(strings.Split(sb.String(), "\n"), func(line string) []rune {
		return []rune(line)
	})
}
