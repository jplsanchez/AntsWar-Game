package display

import (
	"antswar/game"
	"fmt"

	"github.com/fatih/color"
)

type ConsoleDisplay struct {
	player                 game.Team
	highlightX, highlightY int
	textColor              color.Color
	board                  game.GameBoard
	messageLog             *MessageLog
}

func (cd *ConsoleDisplay) UpdateBoard(gb game.GameBoard) {
	cd.board = gb
	cd.UpdateDisplay()
}

func (cd *ConsoleDisplay) DisplayMessage(message string) {
	switch cd.player {
	case game.TeamRed:
		cd.textColor = *color.New(color.FgRed)
	case game.TeamBlack:
		cd.textColor = *color.New(color.FgBlue)
	default:
		cd.textColor = *color.New(color.FgWhite)
	}
	cd.messageLog = cd.messageLog.Append(message, cd.textColor)
	cd.UpdateDisplay()
}

func (cd *ConsoleDisplay) AskForString(messages ...string) (string, error) {
	for _, message := range messages {
		cd.DisplayMessage(message)
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err == nil {
		cd.messageLog = cd.messageLog.Append(input, *color.New(color.FgWhite))
	}
	return input, err
}

func (cd *ConsoleDisplay) SetPlayer(player game.Team) {
	cd.player = player
}

func (cd *ConsoleDisplay) SetHighlight(x, y int) {
	cd.highlightX = x
	cd.highlightY = y
}

func (cd *ConsoleDisplay) UpdateDisplay() {
	if cd.messageLog == nil {
		cd.messageLog = &MessageLog{}
	}
	CallClear()
	cd.PrintBoard(cd.board)
	cd.messageLog.Print()
}

func (cd ConsoleDisplay) PrintBoard(gb game.GameBoard) {
	piecePrint := func(gb game.GameBoard, i int, j int) {
		val := gb[j][i].Value
		var primary, secondary *color.Color

		secondary = FgColorByTeam(gb[j][i].Team)

		if cd.highlightX == j && cd.highlightY == i {
			primary = FgColorByTeam(gb[j][i].Team).Add(color.BgGreen)
		} else {
			primary = secondary
		}

		switch {
		case val == 0:
			primary.Printf("Q")
		case val >= 1 && val <= 9:
			primary.Printf("%d", gb[j][i].Value)
		default:
			primary.Printf(" ")
		}
		secondary.Printf(" ")
	}

	middleFieldPrint := func(i int, gb game.GameBoard) {
		if i == gb.Height()/2 {
			fmt.Println("- - - - -")
		}
	}

	for i := 0; i < gb.Height(); i++ {
		middleFieldPrint(i, gb)
		for j := 0; j < gb.Width(); j++ {
			piecePrint(gb, i, j)
		}
		fmt.Println()
	}
}

func FgColorByTeam(team game.Team) *color.Color {
	switch team {
	case game.TeamRed:
		return color.New(color.FgRed)
	case game.TeamBlack:
		return color.New(color.FgBlue)
	default:
		return color.New(color.FgWhite)
	}
}
