package input

import (
	"io"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/must"
)

func Raw(reader io.Reader) []byte {
	return must.Return(io.ReadAll(reader))
}

func Lines(reader io.Reader) []string {
	return strings.Split(string(Raw(reader)), "\n")
}
