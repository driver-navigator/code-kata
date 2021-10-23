package player

const (
	LOVE    = 0
	FIFTEEN = 1
	THIRTY  = 2
	FOURTY  = 3
)

type Player interface {
	Scored()
	LostAdvantage()
	HasAdvantage() bool
	Name() string
	Score() int
}

type player struct {
	name         string
	score        int
	hasAdvantage bool
}

func New(name string) Player {
	return &player{
		name: name,
	}
}

func (p *player) Scored() {
	switch p.score {
	case LOVE:
		p.score = FIFTEEN
	case FIFTEEN:
		p.score = THIRTY
	case THIRTY:
		p.score = FOURTY
	case FOURTY:
		p.hasAdvantage = true
	}
}

func (p *player) LostAdvantage() {
	p.hasAdvantage = false
}

func (p *player) Name() string {
	return p.name
}

func (p *player) Score() int {
	return p.score
}

func (p *player) HasAdvantage() bool {
	return p.hasAdvantage
}
