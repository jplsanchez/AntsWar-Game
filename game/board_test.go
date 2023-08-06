package game

import "testing"

func TestGameBoard_Height(t *testing.T) {
	var board GameBoard
	result := board.Height()
	expected := 8

	if result != expected {
		t.Errorf("Result was %d, but expected value was %d ", result, expected)
	}
}

func TestGameBoard_Width(t *testing.T) {
	var board GameBoard
	result := board.Width()
	expected := 5

	if result != expected {
		t.Errorf("Result was %d, but expected value was %d ", result, expected)
	}
}
