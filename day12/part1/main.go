package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

const inputSteps = 1000

func main() {
	run(os.Stdin, os.Stdout, inputSteps)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer, steps int) {
	inputData := read(reader)
	answer := process(inputData, steps)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []Planet {
	lines := input.Lines(reader)

	return slice.Map(lines[:len(lines)-1], func(line string) Planet {
		parts := transform.StrToInts(line)
		return Planet{position: math.NewVector3(parts[0], parts[1], parts[2])}
	})
}

func process(planets []Planet, steps int) int {
	for range steps {
		movePlanets(planets)
	}

	return totalEnergy(planets)
}

type Planet struct {
	position, velocity math.Vector3[int]
}

func (p Planet) potentialEnergy() int {
	return math.Abs(p.position.X) + math.Abs(p.position.Y) + math.Abs(p.position.Z)
}

func (p Planet) kineticEnergy() int {
	return math.Abs(p.velocity.X) + math.Abs(p.velocity.Y) + math.Abs(p.velocity.Z)
}

func (p Planet) totalEnergy() int {
	return p.potentialEnergy() * p.kineticEnergy()
}

func movePlanets(planets []Planet) {
	for i := range planets {
		cp := &planets[i]
		for _, ap := range planets {
			cp.velocity.X += math.Sign(ap.position.X - cp.position.X)
			cp.velocity.Y += math.Sign(ap.position.Y - cp.position.Y)
			cp.velocity.Z += math.Sign(ap.position.Z - cp.position.Z)
		}
	}

	for i := range planets {
		planet := &planets[i]
		planet.position = planet.position.Add(planet.velocity)
	}
}

func totalEnergy(planets []Planet) int {
	sum := 0
	for _, p := range planets {
		sum += p.totalEnergy()
	}

	return sum
}
