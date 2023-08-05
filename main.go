package main

type Team uint8

func main() {
	display := ConsoleDisplay{}
	run(display)
}

func run(d Displayer) {
	var board GameBoard
	board.LoadInitialBoard()

	d.Display(board)
}
