package main

import (
	"fmt"
	"io"
	"os"

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

type Layer [10]int // color -> count

func process(encodedImage []rune, width, height int) int {
	var bestLayer Layer
	bestLayer[0] = len(encodedImage)
	decodeLayers(encodedImage, width, height, func(layer Layer) {
		if layer[0] < bestLayer[0] {
			bestLayer = layer
		}
	})

	return bestLayer[1] * bestLayer[2]
}

func decodeLayers(encodedImage []rune, width, height int, output func(Layer)) {
	layerSize := width * height
	layersCount := len(encodedImage) / layerSize
	for i := range layersCount {
		var layer Layer
		for j := 0; j < layerSize; j++ {
			idx := i*layerSize + j
			digit := int(encodedImage[idx] - '0')
			layer[digit]++
		}
		output(layer)
	}
}
