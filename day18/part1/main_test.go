package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "4676", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, run, "8", `#########
#b.A.@.a#
#########
`)
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, run, "86", `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`)
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, run, "132", `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`)
}

func TestExample4(t *testing.T) {
	test.AssertStringInput(t, run, "136", `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`)
}

func TestExample5(t *testing.T) {
	test.AssertStringInput(t, run, "81", `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`)
}
