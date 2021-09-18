package main

import (
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/core"
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/printer"
)

func main() {
	print := func(s string) {
		println(s)
	}

	fizzBuzzPrinter := printer.NewFizzBuzzPrinter(print)
	fizzBuzzService := core.NewFizzBuzzService(fizzBuzzPrinter, 1, 100)
	fizzBuzzService.Execute()
}
