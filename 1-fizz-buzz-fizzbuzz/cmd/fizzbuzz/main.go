package main

import (
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/core"
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/printer"
)

func main() {
	print := func(s string) {
		println(s)
	}

	fizzBuzz := core.NewFizzBuzz()
	fizzBuzzPrinter := printer.NewFizzBuzzPrinter(print)
	fizzBuzzService := core.NewFizzBuzzService(1, 100, fizzBuzz, fizzBuzzPrinter)
	fizzBuzzService.Execute()
}
