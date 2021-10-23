package tennisgame_test

import (
	"testing"

	"github.com/driver-navigator/code-kata/4-tennis/internal/core/player"
	"github.com/driver-navigator/code-kata/4-tennis/internal/core/tennisgame"
)

const (
	PLY_ONE = "player 1"
	PLY_TWO = "player 2"
)

func newTestGame(playerOneScore, playerTwoScore int) tennisgame.TennisGame {
	g := tennisgame.New(PLY_ONE, PLY_TWO)

	g = setPlayerScore(g, PLY_ONE, playerOneScore)
	g = setPlayerScore(g, PLY_TWO, playerTwoScore)

	return g
}

func setPlayerScore(g tennisgame.TennisGame, name string, score int) tennisgame.TennisGame {
	for {
		currentScore := g.Player(name).Score()
		if currentScore == score {
			break
		}
		g.Player(name).Scored()
	}
	return g
}

func TestTennisGame_GivenBothPlayersHaveScoreFourty_WhenPlayerOneScores_ThenPlayerOneGainsAdvantage(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)

	//act
	game.Score(PLY_ONE)

	//assert
	want := true
	got := game.Player(PLY_ONE).HasAdvantage()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenBothPlayersHaveScoreFourty_WhenPlayerTwoScores_ThenPlayerTwoGainsAdvantage(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)

	//act
	game.Score(PLY_TWO)

	//assert
	want := true
	got := game.Player(PLY_TWO).HasAdvantage()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerOneHasAdvantage_WhenPlayerOneScores_ThenPlayerOneWins(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)
	game.Player(PLY_ONE).Scored()

	//act
	game.Score(PLY_ONE)

	//assert
	want := PLY_ONE
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerTwoHasAdvantage_WhenPlayerTwoScores_ThenPlayerTwoWins(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)
	game.Player(PLY_TWO).Scored()

	//act
	game.Score(PLY_TWO)

	//assert
	want := PLY_TWO
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerOneHasAdvantage_WhenPlayerTwoScores_ThenGameIsDeduced(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)
	game.Player(PLY_ONE).Scored()

	//act
	game.Score(PLY_TWO)

	//assert
	want := true
	got := game.IsDeduced()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerTwoHasAdvantage_WhenPlayerOneScores_ThenGameIsDeduced(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.FOURTY)
	game.Player(PLY_TWO).Scored()

	//act
	game.Score(PLY_ONE)

	//assert
	want := true
	got := game.IsDeduced()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerOneHasScoreFourtyAndPlayerTwoHasScoreThirty_WhenPlayerTwoScores_ThenGameIsDeduced(t *testing.T) {
	//arrange
	game := newTestGame(player.FOURTY, player.THIRTY)

	//act
	game.Score(PLY_TWO)

	//assert
	want := true
	got := game.IsDeduced()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerTwoHasScoreFourtyAndPlayerOneHasScoreThirty_WhenPlayerOneScores_ThenGameIsDeduced(t *testing.T) {
	//arrange
	game := newTestGame(player.THIRTY, player.FOURTY)

	//act
	game.Score(PLY_ONE)

	//assert
	want := true
	got := game.IsDeduced()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerOneHasScoreThirtyAndPlayerTwoHasScoreFifteen_WhenPlayerOneScores_ThenPlayerOneWins(t *testing.T) {
	//arrange
	game := newTestGame(player.THIRTY, player.FIFTEEN)

	//act
	game.Score(PLY_ONE)

	//assert
	want := PLY_ONE
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerTwoHasScoreThirtyAndPlayerOneHasScoreFifteen_WhenPlayerTwoScores_ThenPlayerTwoWins(t *testing.T) {
	//arrange
	game := newTestGame(player.FIFTEEN, player.THIRTY)

	//act
	game.Score(PLY_TWO)

	//assert
	want := PLY_TWO
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerOneHasScoreThirtyAndPlayerTwoHasScoreLove_WhenPlayerOneScores_ThenPlayerOneWins(t *testing.T) {
	//arrange
	game := newTestGame(player.THIRTY, player.LOVE)

	//act
	game.Score(PLY_ONE)

	//assert
	want := PLY_ONE
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenPlayerTwoHasScoreThirtyAndPlayerOneHasScoreLove_WhenPlayerTwoScores_ThenPlayerTwoWins(t *testing.T) {
	//arrange
	game := newTestGame(player.LOVE, player.THIRTY)

	//act
	game.Score(PLY_TWO)

	//assert
	want := PLY_TWO
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenBothPlayerHaveScoreThirty_WhenPlayerOneScores_NoWinnerYet(t *testing.T) {
	//arrange
	game := newTestGame(player.THIRTY, player.THIRTY)

	//act
	game.Score(PLY_ONE)

	//assert
	want := ""
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTennisGame_GivenBothPlayerHaveScoreThirty_WhenPlayerTwoScores_NoWinnerYet(t *testing.T) {
	//arrange
	game := newTestGame(player.THIRTY, player.THIRTY)

	//act
	game.Score(PLY_TWO)

	//assert
	want := ""
	got := game.Winner()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
