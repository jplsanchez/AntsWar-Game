package display

import (
	"antswar/game"
)

type Displayer interface {
	DisplayBoard(gb game.GameBoard)
	DisplayMessage(message string)
	SetPlayer(player game.Team)
	SetHighlight(x, y int)
	AskForString(messages ...string) (string, error)
	AskForCoordinates(messages ...string) (int, int, error)
}
