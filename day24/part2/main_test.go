package main

import (
	"io"
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func runTest(mins int) func(reader io.Reader, writer io.Writer) {
	return func(reader io.Reader, writer io.Writer) {
		run(reader, writer, mins)
	}
}

func TestInput(t *testing.T) {
	test.AssertFileInput(t, runTest(minutes), "1902", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, runTest(10), "99", `....#
#..#.
#..##
..#..
#....
`)
}
