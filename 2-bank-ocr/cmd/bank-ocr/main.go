package main

import (
	"fmt"

	"github.com/driver-navigator/code-kata/2-bank-ocr/internal/core"
	filereader "github.com/driver-navigator/code-kata/2-bank-ocr/internal/file-reader"
)

func main() {
	fileData := filereader.ReadFile("input.txt")
	arrayOfDots := core.FileParser(fileData)

	var output string
	for _, v := range arrayOfDots {
		fmt.Println("################")
		fmt.Println(v)
		arr := unpack(v...)
		fmt.Println(arr)
		digit := core.ParseDigit(arr)
		output += fmt.Sprint(digit)
	}
	println(output)
}

func unpack(values ...string) [9]string {
	fmt.Println(values)
	var arr [9]string
	for i, v := range values {
		println(i)
		arr[i] = v
	}
	return arr
}
