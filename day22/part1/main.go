package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const CardsCount = 10_007

type Action uint8

const (
	NewStackAction Action = iota
	CutAction
	IncrementAction
)

type Deal struct {
	action Action
	value  int
}

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []Deal {
	lines := input.Lines(reader)

	return slice.Map(lines[:len(lines)-1], func(line string) Deal {
		if strings.HasPrefix(line, "deal into") {
			return Deal{action: NewStackAction}
		}
		value := transform.StrToInts(line)[0]
		switch {
		case strings.HasPrefix(line, "cut"):
			return Deal{action: CutAction, value: value}
		case strings.HasPrefix(line, "deal with"):
			return Deal{action: IncrementAction, value: value}
		}

		panic("unrecognized action")
	})
}

func process(deals []Deal) int {
	deck := make([]int, CardsCount)
	newDeck := make([]int, CardsCount)
	for i := range CardsCount {
		deck[i] = i
	}

	for _, deal := range deals {
		value := deal.value
		switch deal.action {
		case NewStackAction:
			for i := range CardsCount / 2 {
				j := CardsCount - i - 1
				deck[i], deck[j] = deck[j], deck[i]
			}
		case CutAction:
			if value > 0 {
				copy(newDeck[CardsCount-value:], deck[:value])
				copy(newDeck, deck[value:])
			} else {
				copy(newDeck, deck[CardsCount+value:])
				copy(newDeck[-value:], deck[:CardsCount+value])
			}
			deck, newDeck = newDeck, deck
		case IncrementAction:
			for i, v := range deck {
				newDeck[(i*value)%CardsCount] = v
			}
			deck, newDeck = newDeck, deck
		default:
			panic("unknown action")
		}
	}

	for i, v := range deck {
		if v == 2019 {
			return i
		}
	}

	panic("card not found")
}
