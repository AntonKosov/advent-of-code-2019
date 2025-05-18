package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2019/aoc/input"
)

const (
	inputWidth  = 25
	inputHeight = 6
)

func main() {
	run(os.Stdin, os.Stdout, inputWidth, inputHeight)
	fmt.Println()
}

func run(reader io.Reader, writer io.Writer, width, height int) {
	inputData := read(reader)
	answer := process(inputData, width, height)
	fmt.Fprint(writer, answer)
}

func read(reader io.Reader) []rune {
	return []rune(input.Lines(reader)[0])
}

func process(encodedImage []rune, width, height int) Image {
	image := NewImage(width, height)

	for i, pxl := range encodedImage {
		if pxl == transparentColor {
			continue
		}
		x := i % width
		y := (i / width) % height
		if image[y][x] != transparentColor {
			continue
		}
		image[y][x] = pxl
	}

	return image
}

const transparentColor = '2'

type Image [][]rune

func NewImage(width, height int) Image {
	image := make(Image, height)
	for i := range image {
		row := make([]rune, width)
		for j := range row {
			row[j] = transparentColor
		}
		image[i] = row
	}

	return image
}

func (i Image) String() string {
	var sb strings.Builder
	for _, row := range i {
		if sb.Len() > 0 {
			sb.WriteRune('\n')
		}

		sb.WriteString(string(row))
	}

	return sb.String()
}
