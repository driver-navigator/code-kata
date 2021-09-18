package printer

import (
	"testing"

	"github.com/driver-navigator/code-kata/1-fizz-buzz-fizzbuzz/internal/core"
)

func TestFizzBuzzPrinter_PrintsAsExpected(t *testing.T) {
	got := ""
	want := "test str"

	var sut core.FizzBuzzPrinter = &fizzBuzzPrinter{
		print: func(s string) {
			got = s
		},
	}

	sut.Print(want)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
