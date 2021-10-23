package main

import (
	"github.com/driver-navigator/code-kata/2-bank-ocr/internal/core"
	filereader "github.com/driver-navigator/code-kata/2-bank-ocr/internal/file-reader"
)

func main() {
	ocrLine := filereader.ReadFile("input2.txt")
	ocrChars := core.ParseOcrLine(ocrLine)

	var output string
	for _, ocrChar := range ocrChars {
		ch := core.ParseOcrChar(ocrChar)
		output += ch
	}
	println(output)
}
