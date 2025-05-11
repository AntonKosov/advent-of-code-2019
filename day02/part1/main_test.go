package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, Run, "5482655", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, Run, "3500", "1,9,10,3,2,3,11,0,99,30,40,50\n")
}
