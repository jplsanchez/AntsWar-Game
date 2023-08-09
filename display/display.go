package display

import (
	"antswar/game"
)

type Displayer interface {
	Init()
	UpdateBoard(gb game.GameBoard)
	DisplayMessage(message string)
	AskForString(messages ...string) (string, error)
	SetPlayer(player game.Team)
	SetHighlight(x, y int)
}
