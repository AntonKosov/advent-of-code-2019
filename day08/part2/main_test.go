package main

import (
	"io"
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func runWithSize(width, height int) func(io.Reader, io.Writer) {
	return func(r io.Reader, w io.Writer) {
		run(r, w, width, height)
	}
}

func TestInput(t *testing.T) {
	test.AssertFileInput(t, runWithSize(inputWidth, inputHeight), `1111010000111000011011110
0001010000100100001010000
0010010000111000001011100
0100010000100100001010000
1000010000100101001010000
1111011110111000110010000`, "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, runWithSize(2, 2), "01\n10", "0222112222120000")
}
