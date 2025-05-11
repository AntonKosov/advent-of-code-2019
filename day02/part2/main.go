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

func main() {
	Run(os.Stdin, os.Stdout)
	fmt.Println()
}

func Run(reader io.Reader, writer io.Writer) {
	inputData := read(reader)
	answer := process(inputData)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []int {
	lines := input.Lines(reader)
	nums := strings.Split(lines[0], ",")

	return slice.Map(nums, func(num string) int { return transform.StrToInt(num) })
}

func process(nums []int) int {
	attempt := make([]int, len(nums))
	for noun := 0; ; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(attempt, nums)
			attempt[1], attempt[2] = noun, verb
			if run(attempt) == 19690720 {
				return noun*100 + verb
			}
		}
	}
}

func run(nums []int) int {
	for i := 0; i < len(nums); i += 4 {
		switch opcode := nums[i]; opcode {
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		case 99:
			return nums[0]
		default:
			panic(fmt.Sprintf("unknown opcode: %v", opcode))
		}
	}

	panic("something went wrong")
}
