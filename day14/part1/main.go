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

const (
	ore  = "ORE"
	fuel = "FUEL"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

type ChemicalName string

type Chemical struct {
	name     ChemicalName
	quantity int
}

type Reaction struct {
	inputChemicals []Chemical
	outputChemical Chemical
}

func read(reader io.Reader) []Reaction {
	lines := input.Lines(reader)
	parseChemicals := func(chemicals string) []Chemical {
		parts := strings.Split(chemicals, ", ")
		return slice.Map(parts, func(chemical string) Chemical {
			parts := strings.Split(chemical, " ")
			return Chemical{
				name:     ChemicalName(parts[1]),
				quantity: transform.StrToInt(parts[0]),
			}
		})
	}

	return slice.Map(lines[:len(lines)-1], func(line string) Reaction {
		parts := strings.Split(line, " => ")
		return Reaction{
			inputChemicals: parseChemicals(parts[0]),
			outputChemical: parseChemicals(parts[1])[0],
		}
	})
}

func process(reactions []Reaction) int {
	graph := buildGraph(reactions)
	quantity := map[ChemicalName]int{}

	return countOre(fuel, 1, graph, quantity)
}

func countOre(
	chemicalName ChemicalName, requiredQuantity int, graph map[ChemicalName]Reaction,
	quantity map[ChemicalName]int,
) int {
	if chemicalName == ore {
		return requiredQuantity
	}

	existingChemicals := quantity[chemicalName]
	if existingChemicals >= requiredQuantity {
		quantity[chemicalName] -= requiredQuantity
		return 0
	}

	ore := 0
	reaction := graph[chemicalName]
	outputQuantity := reaction.outputChemical.quantity
	requiredNewChemicals := requiredQuantity - existingChemicals
	batches := requiredNewChemicals / outputQuantity
	if requiredNewChemicals%outputQuantity != 0 {
		batches++
	}

	for _, c := range reaction.inputChemicals {
		ore += countOre(c.name, batches*c.quantity, graph, quantity)
	}

	quantity[chemicalName] = existingChemicals + batches*outputQuantity - requiredQuantity

	return ore
}

func buildGraph(reactions []Reaction) map[ChemicalName]Reaction {
	res := make(map[ChemicalName]Reaction, len(reactions))
	for _, reaction := range reactions {
		res[reaction.outputChemical.name] = reaction
	}

	return res
}
