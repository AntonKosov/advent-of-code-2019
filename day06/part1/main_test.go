package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "273985", "input.txt")
}

func TestExample(t *testing.T) {
	test.AssertStringInput(t, run, "42", `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`)
}
