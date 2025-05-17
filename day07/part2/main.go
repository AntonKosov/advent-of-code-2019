package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer) {
	computer := read(reader)
	answer := process(computer)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []int {
	lines := input.Lines(reader)

	program := slice.Map(
		strings.Split(lines[0], ","),
		func(num string) int { return transform.StrToInt(num) },
	)

	return program
}

func process(program []int) int {
	phases := []int{5, 6, 7, 8, 9}
	maxOutput := 0
	slice.Permute(phases, func() bool {
		output := calcOutput(program, phases)
		maxOutput = max(maxOutput, output)
		return true
	})

	return maxOutput
}

func calcOutput(program, phases []int) int {
	lastOutput := 0
	amplifiers := slice.Map(phases, func(phase int) *Amplifier {
		return NewAmplifier(program, phase)
	})

	for input, i := 0, 0; ; i++ {
		ampIdx := i % len(amplifiers)
		amplifier := amplifiers[ampIdx]
		output, halted := amplifier.Run(input)
		if halted {
			return lastOutput
		}
		if ampIdx == len(amplifiers)-1 {
			lastOutput = output
		}
		input = output
	}
}

type Amplifier struct {
	initialPhase *int
	input        int
	lastOutput   int
	halted       bool
	program      []int
	position     int
	commands     map[int]func() (offset int)
}

const outputOpcode = 4

func NewAmplifier(program []int, initialPhase int) *Amplifier {
	pc := make([]int, len(program))
	copy(pc, program)

	amplifier := Amplifier{
		program:      pc,
		initialPhase: &initialPhase,
	}

	amplifier.commands = map[int]func() (offset int){
		1:            amplifier.sumCommand,
		2:            amplifier.mulCommand,
		3:            amplifier.inputCommand,
		outputOpcode: amplifier.outputCommand,
		5:            amplifier.jumpIfTrueCommand,
		6:            amplifier.jumpIfFalseCommand,
		7:            amplifier.lessThanCommand,
		8:            amplifier.equalsCommand,
		99:           amplifier.haltCommand,
	}

	return &amplifier
}

func (a *Amplifier) Run(input int) (output int, halted bool) {
	if a.halted {
		panic("amplifier is halted")
	}

	a.input = input

	for {
		opcode := a.program[a.position] % 100
		offset := a.commands[opcode]()
		a.position += offset
		if a.halted {
			return 0, true
		}
		if opcode == outputOpcode {
			return a.lastOutput, false
		}
	}
}

func (a *Amplifier) sumCommand() (offset int) {
	a.set(3, a.get(1)+a.get(2))
	return 4
}

func (a *Amplifier) mulCommand() (offset int) {
	a.set(3, a.get(1)*a.get(2))
	return 4
}

func (a *Amplifier) inputCommand() (offset int) {
	v := a.input
	if pv := a.initialPhase; pv != nil {
		v = *pv
		a.initialPhase = nil
	}
	a.set(1, v)

	return 2
}

func (a *Amplifier) outputCommand() (offset int) {
	a.lastOutput = a.get(1)

	return 2
}

func (a *Amplifier) haltCommand() (offset int) {
	a.halted = true
	return 1
}

func (a *Amplifier) jumpIfTrueCommand() (offset int) {
	if a.get(1) != 0 {
		return a.get(2) - a.position
	}

	return 3
}

func (a *Amplifier) jumpIfFalseCommand() (offset int) {
	if a.get(1) == 0 {
		return a.get(2) - a.position
	}

	return 3
}

func (a *Amplifier) lessThanCommand() (offset int) {
	value := 0
	if a.get(1) < a.get(2) {
		value = 1
	}

	a.set(3, value)

	return 4
}

func (a *Amplifier) equalsCommand() (offset int) {
	value := 0
	if a.get(1) == a.get(2) {
		value = 1
	}

	a.set(3, value)

	return 4
}

func (a *Amplifier) get(paramNum int) int {
	idx := a.position + paramNum
	if a.paramValue(paramNum) == 1 {
		return a.program[idx]
	}

	return a.program[a.program[idx]]
}

func (a *Amplifier) set(paramNum, value int) {
	idx := a.position + paramNum
	if a.paramValue(paramNum) == 0 {
		idx = a.program[idx]
	}

	a.program[idx] = value
}

func (a *Amplifier) paramValue(paramNum int) int {
	return (a.program[a.position] / math.Pow(10, uint(paramNum)+1)) % 10
}
