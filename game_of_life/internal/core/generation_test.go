package core

import (
	"testing"

	"github.com/driver-navigator/code-kata/game-of-life/pkg/matrix"
)

func TestGeneration_SetsGenerationCorrectly(t *testing.T) {
	want := [][]bool{
		{ALIVE, DEAD},
		{DEAD, ALIVE},
	}

	sut, err := NewGeneration(2, 2, want)

	if err != nil {
		t.Errorf("expected to set the current generation correctly, but got %v", err.Error())
	}

	got := sut.Population()

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGeneration_WhenIncorrectRows_ReturnsError(t *testing.T) {
	_, err := NewGeneration(3, 2, [][]bool{
		{ALIVE, DEAD},
		{DEAD, ALIVE},
	})

	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestGeneration_WhenIncorrectCols_ReturnsError(t *testing.T) {
	_, err := NewGeneration(2, 1, [][]bool{
		{ALIVE, DEAD},
		{DEAD, ALIVE},
	})

	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestGeneration_WhenAliveCellHasFewerThanTwoAliveNeighbors_ItDies(t *testing.T) {
	sut, _ := NewGeneration(2, 2, [][]bool{
		{ALIVE, DEAD},
		{DEAD, ALIVE},
	})
	sut.AdvanceGeneration()

	got := sut.Population()
	want := [][]bool{
		{DEAD, DEAD},
		{DEAD, DEAD},
	}

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGeneration_WhenAliveCellHasMoreThanThreeAliveNeighbors_ItDies(t *testing.T) {
	sut, _ := NewGeneration(3, 2, [][]bool{
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
		{ALIVE, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.Population()
	want := [][]bool{
		{ALIVE, ALIVE},
		{DEAD, DEAD},
		{ALIVE, ALIVE},
	}

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGeneration_WhenAliveCellHasTwoAliveNeighbors_ItStaysAlive(t *testing.T) {
	sut, _ := NewGeneration(2, 3, [][]bool{
		{ALIVE, DEAD, ALIVE},
		{DEAD, ALIVE, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.Population()
	want := [][]bool{
		{DEAD, ALIVE, DEAD},
		{DEAD, ALIVE, DEAD},
	}

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGeneration_WhenAliveCellHasThreeAliveNeighbors_ItStaysAlive(t *testing.T) {
	sut, _ := NewGeneration(3, 2, [][]bool{
		{DEAD, ALIVE},
		{ALIVE, ALIVE},
		{DEAD, ALIVE},
	})
	sut.AdvanceGeneration()

	got := sut.Population()
	want := [][]bool{
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
		{ALIVE, ALIVE},
	}

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGeneration_WhenDeadCellHasExactleyThreeAliveNeighbors_ItBecomesAlive(t *testing.T) {
	sut, _ := NewGeneration(3, 3, [][]bool{
		{DEAD, DEAD, ALIVE},
		{DEAD, ALIVE, ALIVE},
		{DEAD, DEAD, DEAD},
	})
	sut.AdvanceGeneration()

	got := sut.Population()
	want := [][]bool{
		{DEAD, ALIVE, ALIVE},
		{DEAD, ALIVE, ALIVE},
		{DEAD, DEAD, DEAD},
	}

	if !matrix.AreEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
