package core

type FizzBuzzer interface {
	Process(int32) string
}

type FizzBuzzService interface {
	Execute()
}

type FizzBuzzPrinter interface {
	Print(string)
}
