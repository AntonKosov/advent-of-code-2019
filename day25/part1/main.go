package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/AntonKosov/advent-of-code-2019/day15/part1/program"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inputCh := make(chan rune)
	outputCh := make(chan rune)
	go program.Run(
		context.Background(),
		program.Parse(file),
		func() int64 { return int64(<-inputCh) },
		func(v int64) { outputCh <- rune(v) },
	)

	manualMode(inputCh, outputCh)
}

func manualMode(inputCh chan<- rune, outputCh <-chan rune) {
	go func() {
		for {
			fmt.Printf("%v", string(<-outputCh))
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		for _, v := range command {
			inputCh <- rune(v)
		}
	}
}
