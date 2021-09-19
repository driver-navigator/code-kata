package printer

import (
	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/core"
)

type fizzBuzzPrinter struct {
	print func(s string)
}

var _ core.FizzBuzzPrinter = &fizzBuzzPrinter{}

func NewFizzBuzzPrinter(print func(s string)) core.FizzBuzzPrinter {
	return &fizzBuzzPrinter{
		print: print,
	}
}

func (p *fizzBuzzPrinter) Print(s string) {
	p.print(s)
}
