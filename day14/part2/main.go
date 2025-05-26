package main

import (
	"fmt"
	"io"
	"maps"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const (
	cargo = 1_000_000_000_000
	ore   = "ORE"
	fuel  = "FUEL"
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
	quantity uint64
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
				quantity: uint64(transform.StrToInt(parts[0])),
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

func process(reactions []Reaction) uint64 {
	graph := buildGraph(reactions)
	quantity := map[ChemicalName]uint64{ore: cargo}
	fuelQuantity := uint64(0)

	for amount := uint64(cargo); amount > 0; {
		q := maps.Clone(quantity)
		if !produceChemical(fuel, amount, graph, q) {
			amount /= 2
			continue
		}
		fuelQuantity += amount
		quantity = q
	}

	return fuelQuantity
}

func produceChemical(
	chemicalName ChemicalName, requiredQuantity uint64, graph map[ChemicalName]Reaction,
	quantity map[ChemicalName]uint64,
) bool {
	if chemicalName == ore {
		if quantity[ore] < requiredQuantity {
			return false
		}

		quantity[ore] -= requiredQuantity

		return true
	}

	existingChemicals := quantity[chemicalName]
	if existingChemicals >= requiredQuantity {
		quantity[chemicalName] -= requiredQuantity
		return true
	}

	reaction := graph[chemicalName]
	outputQuantity := reaction.outputChemical.quantity
	requiredNewChemicals := requiredQuantity - existingChemicals
	batches := requiredNewChemicals / outputQuantity
	if requiredNewChemicals%outputQuantity != 0 {
		batches++
	}

	for _, c := range reaction.inputChemicals {
		if !produceChemical(c.name, batches*c.quantity, graph, quantity) {
			return false
		}
	}

	quantity[chemicalName] = existingChemicals + batches*outputQuantity - requiredQuantity

	return true
}

func buildGraph(reactions []Reaction) map[ChemicalName]Reaction {
	res := make(map[ChemicalName]Reaction, len(reactions))
	for _, reaction := range reactions {
		res[reaction.outputChemical.name] = reaction
	}

	return res
}
