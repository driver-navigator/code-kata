package core

import (
	"fmt"
	"testing"
)

func TestFizzBuzz_WhenMultipleOfThree_ReturnsFizz(t *testing.T) {
	sut := NewFizzBuzzer()
	testData := [...]int32{3, 6, 9, 12}
	for _, n := range testData {
		got := sut.Process(n)
		want := "Fizz"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestFizzBuzz_WhenMultipleOfFive_ReturnsBuzz(t *testing.T) {
	sut := NewFizzBuzzer()
	testData := [...]int32{5, 10, 20, 100}
	for _, n := range testData {
		got := sut.Process(n)
		want := "Buzz"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestFizzBuzz_WhenMultipleOfThreeAndFive_ReturnsFizzBuzz(t *testing.T) {
	sut := NewFizzBuzzer()
	testData := [...]int32{15, 30, 45}
	for _, n := range testData {
		got := sut.Process(n)
		want := "FizzBuzz"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestFizzBuzz_WhenNotMultipleOfThreeOrAndFive_ReturnsInput(t *testing.T) {
	sut := NewFizzBuzzer()
	testData := [...]int32{1, 2, 4, 8, 11, 19, 43}
	for _, n := range testData {
		got := sut.Process(n)
		want := fmt.Sprint(n)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestFizzBuzz_WhenLessThanOneOrGreaterThanHundred_ReturnsEmptyString(t *testing.T) {
	sut := NewFizzBuzzer()
	testData := [...]int32{-2, 0, 101}
	for _, n := range testData {
		got := sut.Process(n)
		want := ""

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
