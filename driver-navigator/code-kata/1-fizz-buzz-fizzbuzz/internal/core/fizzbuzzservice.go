package core

type fizzBuzzService struct {
	start    int32
	end      int32
	fizzBuzz FizzBuzzer
	printer  FizzBuzzPrinter
}

var _ FizzBuzzService = &fizzBuzzService{}

func NewFizzBuzzService(start, end int32, fizzBuzz FizzBuzzer, printer FizzBuzzPrinter) FizzBuzzService {
	return &fizzBuzzService{
		start:    start,
		end:      end,
		fizzBuzz: fizzBuzz,
		printer:  printer,
	}
}

func (fs *fizzBuzzService) Execute() {
	for i := fs.start; i <= fs.end; i++ {
		fs.printer.Print(fs.fizzBuzz.Process(i))
	}
}
