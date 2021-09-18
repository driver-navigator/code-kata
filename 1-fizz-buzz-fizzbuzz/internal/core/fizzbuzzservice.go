package core

type FizzBuzzService interface {
	Execute()
}

type fizzBuzzService struct {
	start   int32
	end     int32
	printer FizzBuzzPrinter
}

var _ FizzBuzzService = &fizzBuzzService{}

func NewFizzBuzzService(printer FizzBuzzPrinter, start int32, end int32) FizzBuzzService {
	return &fizzBuzzService{
		printer: printer,
		start:   start,
		end:     end,
	}
}

func (fs *fizzBuzzService) Execute() {
	for i := fs.start; i <= fs.end; i++ {
		fs.printer.Print(fizzBuzz(i))
	}
}
