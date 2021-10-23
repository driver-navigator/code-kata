package arr_test

import (
	"testing"

	"github.com/driver-navigator/code-kata/4-tennis/pkg/arr"
)

func TestAllAreEqual_GivenEqualValues_ReturnsTrue(t *testing.T) {
	want := true
	got := arr.AllAreEqual(40, 40, 40, 40)

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAllAreEqual_GivenUnequalValues_ReturnsFalse(t *testing.T) {
	want := false
	got := arr.AllAreEqual(40, 40, 40, 10)

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
