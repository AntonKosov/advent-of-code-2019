package main

import (
	"io"
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func runWithSteps(steps int) func(io.Reader, io.Writer) {
	return func(reader io.Reader, writer io.Writer) {
		run(reader, writer, steps)
	}
}

func TestInput(t *testing.T) {
	test.AssertFileInput(t, runWithSteps(inputSteps), "7138", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, runWithSteps(10), "179", `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`)
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, runWithSteps(100), "1940", `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`)
}
