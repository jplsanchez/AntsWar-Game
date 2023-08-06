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

	gm := game.NewGameManager(game.TeamBlack, &board)

	for gm.Turns = 0; gm.States.GameFinished; gm.Turns++ { // TODO

	}
}
