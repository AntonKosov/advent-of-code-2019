package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const (
	CardsCount     = int64(119_315_717_514_047)
	ShuffleCount   = int64(101_741_582_076_661)
	TargetPosition = int64(2020)
)

type Action string

const (
	NewStackAction  Action = "rev"
	CutAction       Action = "cut"
	IncrementAction Action = "inc"
)

type Deal struct {
	action Action
	value  int64
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
			return Deal{action: CutAction, value: int64(value)}
		case strings.HasPrefix(line, "deal with"):
			return Deal{action: IncrementAction, value: int64(value)}
		}

		panic("unrecognized action")
	})
}

func mul(v1, v2, mod int64) int64 {
	return new(big.Int).Mod(
		new(big.Int).Mul(big.NewInt(v1), big.NewInt(v2)),
		big.NewInt(mod),
	).Int64()
}

func add(v1, v2, mod int64) int64 {
	return new(big.Int).Mod(
		new(big.Int).Add(big.NewInt(v1), big.NewInt(v2)),
		big.NewInt(mod),
	).Int64()
}

func ruleKey(a1, a2 Action) string {
	return string(a1 + a2)
}

var simplificationRules = map[string]func(*[]Deal, Deal, Deal){
	ruleKey(NewStackAction, NewStackAction): func(*[]Deal, Deal, Deal) {
		// nothing to do
	},
	ruleKey(IncrementAction, IncrementAction): func(deals *[]Deal, d1, d2 Deal) {
		*deals = append(*deals, Deal{action: IncrementAction, value: mul(d1.value, d2.value, CardsCount)})
	},
	ruleKey(CutAction, IncrementAction): func(deals *[]Deal, d1, d2 Deal) {
		*deals = append(
			*deals,
			Deal{action: IncrementAction, value: d2.value},
			Deal{action: CutAction, value: mul(d1.value, d2.value, CardsCount)},
		)
	},
	ruleKey(CutAction, CutAction): func(deals *[]Deal, d1, d2 Deal) {
		*deals = append(*deals, Deal{action: CutAction, value: add(d1.value, d2.value, CardsCount)})
	},
	ruleKey(NewStackAction, IncrementAction): func(deals *[]Deal, d1, d2 Deal) {
		*deals = append(
			*deals,
			Deal{action: IncrementAction, value: CardsCount - d2.value},
			Deal{action: CutAction, value: d2.value},
		)
	},
	ruleKey(NewStackAction, CutAction): func(deals *[]Deal, d1, d2 Deal) {
		*deals = append(
			*deals,
			Deal{action: CutAction, value: -d2.value},
			Deal{action: NewStackAction},
		)
	},
}

func simplify(deals []Deal) []Deal {
	for {
		simplified := false
		result := make([]Deal, 0, len(deals))
		for inc, i := 0, 0; i < len(deals); i += inc {
			inc = 1
			d1 := deals[i]
			if i == len(deals)-1 {
				result = append(result, d1)
				break
			}
			d2 := deals[i+1]

			if rule, ok := simplificationRules[ruleKey(d1.action, d2.action)]; ok {
				simplified = true
				inc = 2
				rule(&result, d1, d2)
				continue
			}

			result = append(result, d1)
		}

		if !simplified {
			return result
		}

		deals = result
	}
}

func process(deals []Deal) int64 {
	cache := [][]Deal{simplify(deals)}
	for 1<<len(cache) <= ShuffleCount/2 {
		lastDeals := cache[len(cache)-1]
		cp := make([]Deal, len(lastDeals)*2)
		copy(cp, lastDeals)
		copy(cp[len(lastDeals):], lastDeals)
		cache = append(cache, simplify(cp))
	}

	result := cache[len(cache)-1]
	for rest := ShuffleCount - (1 << (len(cache) - 1)); rest > 0; {
		count := int64(1 << (len(cache) - 1))
		if count > rest {
			cache = cache[:len(cache)-1]
			continue
		}
		result = append(result, cache[len(cache)-1]...)
		rest -= count
		result = simplify(result)
	}

	return sourceIndex(result, TargetPosition)
}

func sourceIndex(deals []Deal, targetIndex int64) int64 {
	idx := targetIndex
	for i := len(deals) - 1; i >= 0; i-- {
		deal := deals[i]
		value := deal.value
		switch action := deal.action; action {
		case NewStackAction:
			idx = CardsCount - idx - 1
		case CutAction:
			idx = getSourceCutIndex(value, idx)
		case IncrementAction:
			idx = mul(idx, modularInverse(value, CardsCount), CardsCount)
		default:
			panic(fmt.Sprintf("unknown action: %v", action))
		}
	}

	return idx
}

func getSourceCutIndex(value, targetIndex int64) int64 {
	if value < 0 {
		value += CardsCount
	}

	if targetIndex >= CardsCount-value {
		return targetIndex - (CardsCount - value)
	}

	return value + targetIndex
}

func modularInverse(a, n int64) int64 {
	t, newT := int64(0), int64(1)
	r, newR := n, a
	for newR != 0 {
		q := r / newR
		t, newT = newT, t-q*newT
		r, newR = newR, r-q*newR
	}

	if r > 1 {
		panic("no solution")
	}

	return (t + n) % n
}
