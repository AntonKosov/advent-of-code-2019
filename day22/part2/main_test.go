package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "41653717360577", "input.txt")
}
