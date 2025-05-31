package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "82994322", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, run, "84462026", "03036732577212944063491565474664")
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, run, "78725270", "02935109699940807407585447034323")
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, run, "53553731", "03081770884921959731165446850517")
}
