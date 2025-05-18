package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "3454977209", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, run, "1091204-1100110011001008100161011006101099", "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, run, "1219070632396864", "1102,34915192,34915192,7,4,7,99,0")
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, run, "1125899906842624", "104,1125899906842624,99")
}
