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

type fizzBuzzMock struct {
}

var processMock func(int32) string

func (fz *fizzBuzzMock) Process(n int32) string {
	return processMock(n)
}

func TestFizzBuzzService_CallsPrintExpectedNumberOfTimes(t *testing.T) {
	var got int32
	var want int32 = 50

	var fizzBuzzMock FizzBuzzer = &fizzBuzzMock{}
	var fizzBuzzPrinterMock FizzBuzzPrinter = &fizzBuzzPrinterMock{}
	var sut FizzBuzzService = &fizzBuzzService{
		start:    1,
		end:      50,
		printer:  fizzBuzzPrinterMock,
		fizzBuzz: fizzBuzzMock,
	}

	printMock = func(s string) {
		got++
	}

	processMock = func(n int32) string {
		return ""
	}

	sut.Execute()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFizzBuzzService_CallsFizzBuzzProcessExpectedNumberOfTimes(t *testing.T) {
	var got int32
	var want int32 = 50

	var fizzBuzzMock FizzBuzzer = &fizzBuzzMock{}
	var fizzBuzzPrinterMock FizzBuzzPrinter = &fizzBuzzPrinterMock{}
	var sut FizzBuzzService = &fizzBuzzService{
		start:    1,
		end:      50,
		printer:  fizzBuzzPrinterMock,
		fizzBuzz: fizzBuzzMock,
	}

	printMock = func(s string) {}

	processMock = func(n int32) string {
		got++
		return ""
	}

	sut.Execute()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
