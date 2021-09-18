package core

import "fmt"

func fizzBuzz(n int32) string {
	if n < 1 || n > 100 {
		return ""
	}

	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprint(n)
}

type FizzBuzzPrinter interface {
	Print(string)
}
