package matrix

import (
	"testing"
)

func TestMatrix_ClonesCorrectLength(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
		{true, true},
		{false, false},
	}

	matrixB := Clone(matrixA)

	want := []int{4, 2}
	got := []int{len(matrixB), len(matrixB[0])}

	if want[0] != got[0] {
		t.Errorf("len of matrix got %v, wanted %v", got, want)
	}

	if want[1] != got[1] {
		t.Errorf("len of matrix[0] got %v, wanted %v", got, want)
	}
}

func TestMatrix_ClonesCorrectValues(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	got := Clone(matrixA)

	if !got[0][0] {
		t.Errorf("got %v at got[0][0], wanted true", got[0][0])
	}

	if got[0][1] {
		t.Errorf("got %v at got[0][1], wanted false", got[0][1])
	}

	if got[1][0] {
		t.Errorf("got %v at got[1][0], wanted false", got[1][0])
	}

	if !got[1][1] {
		t.Errorf("got %v at got[1][1], wanted true", got[1][1])
	}
}

func TestMatrix_ClonedByValue(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	got := Clone(matrixA)

	matrixA[0][0] = false
	matrixA[0][1] = true

	if !got[0][0] {
		t.Errorf("got %v at got[0][0], wanted true", got[0][0])
	}

	if got[0][1] {
		t.Errorf("got %v at got[0][1], wanted false", got[0][1])
	}

	if got[1][0] {
		t.Errorf("got %v at got[1][0], wanted false", got[1][0])
	}

	if !got[1][1] {
		t.Errorf("got %v at got[1][1], wanted true", got[1][1])
	}
}

func TestMatrix_AreEqual(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	matrixB := [][]bool{
		{true, false},
		{false, true},
	}

	got := AreEqual(matrixA, matrixB)

	if !got {
		t.Errorf("expected matrix %v equal to %v", matrixA, matrixB)
	}
}

func TestMatrix_AreNotEqual(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	matrixB := [][]bool{
		{false, false},
		{false, true},
	}

	got := AreEqual(matrixA, matrixB)

	if got {
		t.Errorf("expected matrix %v not equal to %v", matrixA, matrixB)
	}
}

func TestMatrix_OfDifferentNumberOfRows_AreNotEqual(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	matrixB := [][]bool{
		{true, false},
		{false, true},
		{false, true},
	}

	got := AreEqual(matrixA, matrixB)

	if got {
		t.Errorf("expected matrix %v not equal to %v", matrixA, matrixB)
	}
}

func TestMatrix_OfDifferentNumberOfCols_AreNotEqual(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	matrixB := [][]bool{
		{true, false},
		{false, true, false},
	}

	got := AreEqual(matrixA, matrixB)

	if got {
		t.Errorf("expected matrix %v not equal to %v", matrixA, matrixB)
	}
}

func TestMatrix_WhenExists_CanFind(t *testing.T) {
	matrixA := [][]bool{
		{true, false},
		{false, true},
	}

	expectedEle := true

	got := Exists(matrixA, expectedEle)

	if !got {
		t.Errorf("expected to find %v in matrix %v but not found", matrixA, expectedEle)
	}
}

func TestMatrix_WhenNotExists_CanNotFind(t *testing.T) {
	matrixA := [][]bool{
		{true, true},
		{true, true},
	}

	expectedEle := false

	got := Exists(matrixA, expectedEle)

	if got {
		t.Errorf("expected to not find %v in matrix %v but found", matrixA, expectedEle)
	}
}
