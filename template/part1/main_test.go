package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "", "example")
}
