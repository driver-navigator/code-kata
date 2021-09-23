package core

import (
	"testing"
)

func TestDigitParser_CanParseZero(t *testing.T) {
	var input = [9]string{"", "_", "", "|", "", "|", "|", "_", "|"}

	got := ParseDigit(input)
	want := int8(0)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseOne(t *testing.T) {
	var input = [9]string{"", "", "", "", "", "|", "", "", "|"}

	got := ParseDigit(input)
	want := int8(1)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseTwo(t *testing.T) {
	var input = [9]string{"", "_", "", "", "_", "|", "|", "_", ""}

	got := ParseDigit(input)
	want := int8(2)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseThree(t *testing.T) {
	var input = [9]string{"", "_", "", "", "_", "|", "", "_", "|"}

	got := ParseDigit(input)
	want := int8(3)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseFour(t *testing.T) {
	var input = [9]string{"", "", "", "|", "_", "|", "", "", "|"}

	got := ParseDigit(input)
	want := int8(4)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseFive(t *testing.T) {
	var input = [9]string{"", "_", "", "|", "_", "", "", "_", "|"}

	got := ParseDigit(input)
	want := int8(5)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseSix(t *testing.T) {
	var input = [9]string{"", "_", "", "|", "_", "", "|", "_", "|"}

	got := ParseDigit(input)
	want := int8(6)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseSeven(t *testing.T) {
	var input = [9]string{"", "_", "", "", "|", "", "", "", "|"}

	got := ParseDigit(input)
	want := int8(7)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseEight(t *testing.T) {
	var input = [9]string{"", "_", "", "|", "_", "|", "|", "_", "|"}

	got := ParseDigit(input)
	want := int8(8)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDigitParser_CanParseNine(t *testing.T) {
	var input = [9]string{"_", "", "", "|", "_", "|", "", "_", "|"}

	got := ParseDigit(input)
	want := int8(9)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
