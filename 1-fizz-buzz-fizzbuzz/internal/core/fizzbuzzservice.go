package core

type FizzBuzzService interface {
	Execute()
}

type fizzBuzzService struct {
	p FizzBuzzPrinter
}

var _ FizzBuzzService = &fizzBuzzService{}

func NewFizzBuzzService(printer FizzBuzzPrinter) FizzBuzzService {
	return &fizzBuzzService{
		p: printer,
	}
}

func (fs *fizzBuzzService) Execute() {
	min := 1
	max := 100

	for i := min; i <= max; i++ {
		fs.p.Print(fizzBuzz(i))
	}
}
