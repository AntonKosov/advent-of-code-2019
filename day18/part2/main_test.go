package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2019/aoc/test"
)

func TestInput(t *testing.T) {
	test.AssertFileInput(t, run, "2066", "input.txt")
}

func TestExample1(t *testing.T) {
	test.AssertStringInput(t, run, "8", `#######
#a.#Cd#
##...##
##.@.##
##...##
#cB#Ab#
#######
`)
}

func TestExample2(t *testing.T) {
	test.AssertStringInput(t, run, "24", `###############
#d.ABC.#.....a#
######...######
######.@.######
######...######
#b.....#.....c#
###############
`)
}

func TestExample3(t *testing.T) {
	test.AssertStringInput(t, run, "32", `#############
#DcBa.#.GhKl#
#.###...#I###
#e#d#.@.#j#k#
###C#...###J#
#fEbA.#.FgHi#
#############
`)
}

func TestExample4(t *testing.T) {
	test.AssertStringInput(t, run, "72", `#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba...BcIJ#
#####.@.#####
#nK.L...G...#
#M###N#H###.#
#o#m..#i#jk.#
#############
`)
}
