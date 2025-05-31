package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "90744714", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, run, "24176176", "80871224585914546619083218645595")
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, run, "73745418", "19617804207202209144916044189917")
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, run, "52432133", "69317163492948606335995924319873")
}
