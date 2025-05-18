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
	test.AssertFileInput(t, runWithSize(inputWidth, inputHeight), "828", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, runWithSize(3, 2), "1", "123456789012")
}
