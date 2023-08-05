package main

import (
	"antswar/display"
	"antswar/game"
)

func main() {
	display := display.ConsoleDisplay{}
	run(display)
}

func run(d display.Displayer) {
	var board game.GameBoard
	board.LoadInitialBoard()

	d.Display(board)
}
