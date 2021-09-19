package main

import (
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/core"
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/printer"
)

const (
	start = 1
	end   = 100
)

func main() {
	print := func(s string) {
		println(s)
	}

	fizzBuzzer := core.NewFizzBuzzer()
	fizzBuzzPrinter := printer.NewFizzBuzzPrinter(print)
	fizzBuzzService := core.NewFizzBuzzService(start, end, fizzBuzzer, fizzBuzzPrinter)
	fizzBuzzService.Execute()
}
