package main

import (
	"io"
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func runWithInputInstruction(inputInstruction int) func(io.Reader, io.Writer) {
	return func(r io.Reader, w io.Writer) {
		run(r, w, inputInstruction)
	}
}

func TestInput(t *testing.T) {
	test.AssertFileInput(t, runWithInputInstruction(5), "12077198", "input.txt")
}

const example = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

func TestExample999(t *testing.T) {
	test.AssertStringInput(t, runWithInputInstruction(7), "999", example)
}

func TestExample1000(t *testing.T) {
	test.AssertStringInput(t, runWithInputInstruction(8), "1000", example)
}

func TestExample1001(t *testing.T) {
	test.AssertStringInput(t, runWithInputInstruction(9), "1001", example)
}
