package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, Run, "6484", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, Run, "30", "R8,U5,L5,D3\nU7,R6,D4,L4\n")
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, Run, "610", "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83\n")
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, Run, "410", `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7
`)
}
