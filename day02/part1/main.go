package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2019/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	lines := input.Lines()
	nums := strings.Split(lines[0], ",")

	return slice.Map(nums, func(num string) int { return transform.StrToInt(num) })
}

func process(nums []int) int {
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
