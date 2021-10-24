package main

import (
	"fmt"
	"os"

	tennisgame "github.com/driver-navigator/code-kata/4-tennis/internal/core/tennis-game"
	"github.com/driver-navigator/code-kata/4-tennis/pkg/clearscreen"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	playerOne = "SANDEEP"
	playerTwo = "AMIT"
)

func main() {
	g := tennisgame.New(playerOne, playerTwo)

	updateScoreboard(g)
	for {
		whoScored := getWhoScored()
		g.Score(whoScored)

		logs = append(logs, fmt.Sprintf("'%s' scored.", whoScored))
		updateScoreboard(g)

		winner := g.Winner()
		if winner != "" {
			break
		}
	}
}

func getWhoScored() string {
	for {
		var who string
		fmt.Scanln(&who)

		switch who {
		case "1":
			return playerOne
		case "2":
			return playerTwo
		}
		println("Invalid input. Please try again.")
	}
}

func updateScoreboard(gm tennisgame.TennisGame) {
	clearscreen.Clear()
	plyOne := gm.Player(playerOne)
	plyTwo := gm.Player(playerTwo)

	winner := gm.Winner()

	gameStatus := "In progress"

	if winner != "" {
		gameStatus = fmt.Sprintf("'%s' has won", winner)
	} else if gm.IsDeuce() {
		gameStatus += " (Deuce)"
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Tennis Scoreboard",
		"Tennis Scoreboard",
	}, table.RowConfig{AutoMerge: true})
	t.AppendRow(table.Row{
		plyOne.Name(),
		plyTwo.Name(),
	})
	t.AppendSeparator()
	t.AppendRows([]table.Row{
		{
			scoreMap[plyOne.Score()],
			scoreMap[plyTwo.Score()],
		},
		{
			advMap[plyOne.HasAdvantage()],
			advMap[plyTwo.HasAdvantage()],
		},
	})
	t.AppendSeparator()
	for _, v := range logs {
		t.AppendRow(table.Row{
			v,
			v,
		}, table.RowConfig{AutoMerge: true})
	}
	t.AppendFooter(table.Row{
		gameStatus,
		gameStatus,
	}, table.RowConfig{AutoMerge: true})
	t.SetStyle(table.StyleRounded)
	t.Render()

	ti := table.NewWriter()
	ti.SetOutputMirror(os.Stdout)
	ti.AppendHeader(table.Row{
		"Instructions",
	})
	ti.AppendRows([]table.Row{
		{
			fmt.Sprintf("Enter %s when %s has scored", "1", playerOne),
		},
		{
			fmt.Sprintf("Enter %s when %s has scored", "2", playerTwo),
		},
	})
	ti.AppendFooter(table.Row{
		"Please enter below: ",
	})
	ti.SetStyle(table.StyleColoredBlackOnCyanWhite)
	ti.Render()
}

var advMap = map[bool]string{
	true:  "*Has Advantage*",
	false: "",
}

var scoreMap = map[int]string{
	0: "LOVE",
	1: "FIFTEEN",
	2: "THIRTY",
	3: "FOURTY",
}

var logs = make([]string, 0)
