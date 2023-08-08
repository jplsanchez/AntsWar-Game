package display

import (
	"antswar/game"
	"fmt"

	"github.com/fatih/color"
)

type Displayer interface {
	SetPlayer(player game.Team)
	SetHighlight(x, y int)
	DisplayBoard(gb game.GameBoard)
	DisplayMessage(message string)
	AskForCoordinates(messages ...string) (x, y int, err error)
	AskForString(messages ...string) (string, error)
}

type ConsoleDisplay struct {
	player                 game.Team
	highlightX, highlightY int
	textColor              color.Color
}

func (cd *ConsoleDisplay) SetPlayer(player game.Team) {
	cd.player = player
}

func (cd *ConsoleDisplay) SetHighlight(x, y int) {
	cd.highlightX = x
	cd.highlightY = y
}

func (cd ConsoleDisplay) DisplayBoard(gb game.GameBoard) {
	CallClear()
	for i := 0; i < gb.Height(); i++ {
		cd.DisplayMiddle(i, gb)
		for j := 0; j < gb.Width(); j++ {
			cd.HandleAndDisplay(gb, i, j)
		}
		fmt.Println()
	}
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
	cd.textColor.Println(message)
}

func (cd ConsoleDisplay) AskForCoordinates(messages ...string) (x, y int, err error) {
	for _, message := range messages {
		cd.DisplayMessage(message)
	}
	cd.DisplayMessage("x,y: ")
	_, err = fmt.Scanf("%d,%d", &x, &y)
	return x, y, err
}

func (cd ConsoleDisplay) AskForString(messages ...string) (string, error) {
	for _, message := range messages {
		cd.DisplayMessage(message)
	}
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

func (ConsoleDisplay) DisplayMiddle(i int, gb game.GameBoard) {
	if i == gb.Height()/2 {
		fmt.Println("- - - - -")
	}
}

func (cd ConsoleDisplay) HandleAndDisplay(gb game.GameBoard, i int, j int) {
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
