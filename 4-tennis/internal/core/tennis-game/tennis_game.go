package tennisgame

import (
	"github.com/driver-navigator/code-kata/4-tennis/internal/core/player"
	"github.com/driver-navigator/code-kata/4-tennis/pkg/arr"
)

const (
	WIN_AT_SCORE    = player.FOURTY
	DEDUCE_AT_SCORE = player.FOURTY
)

type TennisGame interface {
	Deuce()
	Score(string)
	IsDeuce() bool
	Winner() string
	Player(string) player.Player
}

type tennisGame struct {
	playerOne player.Player
	playerTwo player.Player
	isDeuce   bool
	winner    string
}

func New(playerOne, playerTwo string) TennisGame {
	return &tennisGame{
		playerOne: player.New(playerOne),
		playerTwo: player.New(playerTwo),
	}
}

func (t *tennisGame) Score(name string) {
	if name == t.playerOne.Name() && t.playerOne.HasAdvantage() {
		t.winner = t.playerOne.Name()
	}

	if name == t.playerTwo.Name() && t.playerTwo.HasAdvantage() {
		t.winner = t.playerTwo.Name()
	}

	if t.playerOne.Name() == name {
		if t.playerTwo.HasAdvantage() {
			t.playerTwo.LostAdvantage()
		} else {
			t.playerOne.Scored()
		}
	} else {
		if t.playerOne.HasAdvantage() {
			t.playerOne.LostAdvantage()
		} else {
			t.playerTwo.Scored()
		}
	}

	if name == t.playerOne.Name() && t.playerOne.Score() == WIN_AT_SCORE && t.playerTwo.Score() <= WIN_AT_SCORE-2 {
		t.winner = t.playerOne.Name()
	}

	if name == t.playerTwo.Name() && t.playerTwo.Score() == WIN_AT_SCORE && t.playerOne.Score() <= WIN_AT_SCORE-2 {
		t.winner = t.playerTwo.Name()
	}

	if arr.AllAreEqual(DEDUCE_AT_SCORE, t.playerOne.Score(), t.playerTwo.Score()) {
		t.Deuce()
	}
}

func (t *tennisGame) IsDeuce() bool {
	return t.isDeuce
}

func (t *tennisGame) Deuce() {
	t.isDeuce = true
}

func (t *tennisGame) Player(name string) player.Player {
	if name == t.playerOne.Name() {
		return t.playerOne
	}
	return t.playerTwo
}

func (t *tennisGame) Winner() string {
	return t.winner
}
