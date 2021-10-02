package core

import (
	"testing"
)

func TestGameOfLife_WhenAliveCellHasFewerThanTwoAliveNeighbors_ItDies(t *testing.T) {
	sut := NewGameOfLife(2, 2)
	sut.SetCurrentGeneration([][]bool{
		{ALIVE, DEAD},
		{DEAD, ALIVE},
	})
	sut.AdvanceGeneration()

	got := sut.CurrentGeneration()
	want := [][]bool{
		{DEAD, DEAD},
		{DEAD, DEAD},
	}

	if !areEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGameOfLife_WhenAliveCellHasMoreThanThreeAliveNeighbors_ItDies(t *testing.T) {
	sut := NewGameOfLife(3, 2)
	sut.SetCurrentGeneration([][]bool{
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
		{ALIVE, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.CurrentGeneration()
	want := [][]bool{
		{ALIVE, ALIVE},
		{DEAD, DEAD},
		{ALIVE, ALIVE},
	}

	if !areEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGameOfLife_WhenAliveCellHasTwoAliveNeighbors_ItStaysAlive(t *testing.T) {
	sut := NewGameOfLife(2, 3)
	sut.SetCurrentGeneration([][]bool{
		{ALIVE, DEAD, ALIVE},
		{DEAD, ALIVE, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.CurrentGeneration()
	want := [][]bool{
		{DEAD, ALIVE, DEAD},
		{DEAD, ALIVE, DEAD},
	}

	if !areEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGameOfLife_WhenAliveCellHasThreeAliveNeighbors_ItStaysAlive(t *testing.T) {
	sut := NewGameOfLife(3, 2)
	sut.SetCurrentGeneration([][]bool{
		{DEAD, ALIVE},
		{ALIVE, ALIVE},
		{DEAD, ALIVE},
	})
	sut.AdvanceGeneration()

	got := sut.CurrentGeneration()
	want := [][]bool{
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
	}

	if !areEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGameOfLife_WhenDeadCellHasExactleyThreeAliveNeighbors_ItBecomesAlive(t *testing.T) {
	sut := NewGameOfLife(3, 3)
	sut.SetCurrentGeneration([][]bool{
		{DEAD, DEAD, ALIVE},
		{DEAD, ALIVE, ALIVE},
		{DEAD, DEAD, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.CurrentGeneration()
	want := [][]bool{
		{DEAD, ALIVE, ALIVE},
		{DEAD, ALIVE, ALIVE},
		{DEAD, DEAD, DEAD},
	}

	if !areEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func areEqual(arr1, arr2 [][]bool) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}
