package player_test

import (
	"testing"

	"github.com/driver-navigator/code-kata/4-tennis/internal/core/player"
)

var (
	testPlayerName = "player"
)

func newTestPlayer(score int) player.Player {
	testPlayer := player.New(testPlayerName)
	for {
		currentScore := testPlayer.Score()
		if currentScore == score {
			break
		}
		testPlayer.Scored()
	}
	return testPlayer
}

func TestPlayer_GivenPlayerScoreIsLove_WhenPlayerScores_ThenPlayerScoreBecomesFifteen(t *testing.T) {
	//arrange
	testPlayer := newTestPlayer(player.LOVE)

	//act
	testPlayer.Scored()

	//assert
	want := player.FIFTEEN
	got := testPlayer.Score()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPlayer_GivenPlayerScoreIsFifteen_WhenPlayerScores_ThenPlayerScoreBecomesThiry(t *testing.T) {
	//arrange
	testPlayer := newTestPlayer(player.FIFTEEN)

	//act
	testPlayer.Scored()

	//assert
	want := player.THIRTY
	got := testPlayer.Score()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPlayer_GivenPlayerScoreIsThirty_WhenPlayerScores_ThenPlayerScoreBecomesFourty(t *testing.T) {
	//arrange
	testPlayer := newTestPlayer(player.THIRTY)

	//act
	testPlayer.Scored()

	//assert
	want := player.FOURTY
	got := testPlayer.Score()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPlayer_GivenPlayerScoreIsFoutry_WhenPlayerScores_ThenPlayerGainsAdvantage(t *testing.T) {
	//arrange
	testPlayer := newTestPlayer(player.FOURTY)

	//act
	testPlayer.Scored()

	//assert
	want := true
	got := testPlayer.HasAdvantage()

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
