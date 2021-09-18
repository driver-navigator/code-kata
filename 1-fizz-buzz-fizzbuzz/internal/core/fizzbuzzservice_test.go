package core

import (
	"testing"
)

type fizzBuzzPrinterMock struct {
}

var printMock func(s string)

func (fs *fizzBuzzPrinterMock) Print(s string) {
	printMock(s)
}

func TestFizzBuzzService_PrintsExpectedNumberOfTimes(t *testing.T) {
	var got int32
	var want int32 = 50

	var fizzBuzzPrinterMock FizzBuzzPrinter = &fizzBuzzPrinterMock{}
	var sut fizzBuzzService = fizzBuzzService{
		printer: fizzBuzzPrinterMock,
		start:   1,
		end:     50,
	}

	printMock = func(s string) {
		got++
	}

	sut.Execute()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
