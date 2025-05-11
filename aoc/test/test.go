package test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func AssertFileInput(t *testing.T, run func(io.Reader, io.Writer), expectedResult string, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("canot open file: %v", err)
		return
	}
	defer file.Close()

	AssertReaderInput(t, run, expectedResult, file)
}

func AssertStringInput(t *testing.T, run func(io.Reader, io.Writer), expectedResult, input string) {
	AssertReaderInput(t, run, expectedResult, strings.NewReader(input))
}

func AssertReaderInput(t *testing.T, run func(io.Reader, io.Writer), expectedResult string, reader io.Reader) {
	var buf bytes.Buffer
	run(reader, &buf)
	actualResult := buf.String()
	if actualResult != expectedResult {
		t.Errorf("expected result: %v, actual result: %v", expectedResult, actualResult)
	}
}
