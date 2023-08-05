package main

import "testing"

func TestBoardHeight(t *testing.T) {
	var board GameBoard
	result := board.Height()
	expected := 8

	if result != expected {
		t.Errorf("Result was %d, but expected value was %d ", result, expected)
	}
}

func TestBoardWidth(t *testing.T) {
	var board GameBoard
	result := board.Width()
	expected := 5

	if result != expected {
		t.Errorf("Result was %d, but expected value was %d ", result, expected)
	}
}
