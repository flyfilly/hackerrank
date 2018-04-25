package main

import (
	"strconv"
	"testing"
)

func TestDiagonalDiff(t *testing.T) {
	slice := [][]int{
		[]int{11, 2, 4},
		[]int{4, 5, 6},
		[]int{10, 8, -12},
	}

	expected := 15
	actual := diagonalDifference(slice)

	if expected != actual {
		t.Errorf("Tests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}

	slice = [][]int{
		[]int{11, 2, 4, 66},
		[]int{4, 5, 6, 44},
		[]int{10, 8, -12, 17},
		[]int{101, 18, -122, -17},
	}

	expected = 194
	actual = diagonalDifference(slice)

	if expected != actual {
		t.Errorf("Tests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}

}
