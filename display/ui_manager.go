package display

import (
	"antswar/game"
	"errors"
	"regexp"
	"strconv"
)

type UIManager struct {
	Game *game.GameManager
	Disp Displayer
}

func NewUI(d Displayer) *UIManager {
	var board game.GameBoard
	board.LoadInitialBoard()

	return &UIManager{
		Game: game.NewGameManager(game.TeamBlack, &board),
		Disp: d,
	}
}

func (ui *UIManager) Run() {
	ui.Disp.SetHighlight(-1, -1)
	ui.Disp.UpdateBoard(*ui.Game.Board)

	for !ui.Game.States.GameFinished {
		ui.Disp.DisplayMessage(" ")
		ui.Disp.SetPlayer(ui.Game.States.Team)

		switch ui.Game.States.Stage {
		case game.StartTurnStage:
			ui.StartPlayersTurnRoutine()
		case game.ChooseCardStage:
			ui.ChooseCardRoutine()
		case game.ChooseActionStage:
			ui.ChooseActionRoutine()
		case game.ResolveActionStage:
			ui.ResolveActionRoutine()
		case game.EndTurnStage:
			ui.EndTurnRoutine()
		}

		ui.Game.NextGameStage()
	}
}

func (ui *UIManager) StartPlayersTurnRoutine() {
	ui.Game.StartTurn()
	ui.Disp.SetHighlight(-1, -1)
	ui.Disp.UpdateBoard(*ui.Game.Board)
	ui.Disp.DisplayMessage("It's your turn! Team:" + ui.Game.States.Team.String())
}

func (ui *UIManager) ChooseCardRoutine() {
	x, y, err := ui.GetCoordinatesFromUser("Choose a card to play")
	if err != nil {
		ui.Disp.DisplayMessage("Error reading coordinates")
		ui.ChooseCardRoutine()
		return
	}
	err = ui.Game.ChooseCard(x, y)
	if err != nil {
		ui.Disp.DisplayMessage(err.Error())
		ui.ChooseCardRoutine()
		return
	}
	ui.Disp.SetHighlight(x, y)
	ui.Disp.UpdateBoard(*ui.Game.Board)
}

func (ui *UIManager) ChooseActionRoutine() {
	action, err := ui.Disp.AskForString("Choose an action to play (move, swap, march, attack)")
	if err != nil {
		ui.Disp.DisplayMessage("Error reading action")
		ui.ChooseActionRoutine()
		return
	}

	switch action {
	case "move":
		ui.Game.States.ActionSelected = &game.Move{}
	case "swap":
		ui.Game.States.ActionSelected = &game.Swap{}
	case "march":
		ui.Game.States.ActionSelected = &game.March{}
	case "attack":
		ui.Game.States.ActionSelected = &game.Attack{}
	default:
		ui.Disp.DisplayMessage("Invalid action")
		ui.ChooseActionRoutine()
		return
	}

	err = ui.Game.ChooseAction(ui.Game.States.ActionSelected, ui.Game.States.SelectedPosition)
	if err != nil {
		ui.Disp.DisplayMessage(err.Error())
		ui.ChooseActionRoutine()
		return
	}
}

func (ui *UIManager) ResolveActionRoutine() {
	x, y, err := ui.GetCoordinatesFromUser("Choose a target")
	if err != nil {
		ui.Disp.DisplayMessage("Error reading coordinates")
		ui.ResolveActionRoutine()
		return
	}
	err = ui.Game.ResolveAction(x, y)
	if err != nil {
		ui.Disp.DisplayMessage(err.Error())
		ui.ResolveActionRoutine()
		return
	}
	ui.Disp.UpdateBoard(*ui.Game.Board)
}

func (ui *UIManager) EndTurnRoutine() {
	message := ui.Game.CheckIfGameFinished()

	if message != "" {
		ui.Disp.DisplayMessage("Turn ended")
		return
	}

	ui.Disp.DisplayMessage(message)
}

func (ui UIManager) GetCoordinatesFromUser(messages ...string) (int, int, error) {
	str, err := ui.Disp.AskForString(messages...)
	if err != nil {
		return -1, -1, err
	}
	re := regexp.MustCompile("[0-9],[0-9]")
	if re.MatchString(str) {
		x, err := strconv.Atoi(str[0:1])
		if err != nil {
			return -1, -1, err
		}
		y, err := strconv.Atoi(str[2:3])
		if err != nil {
			return -1, -1, err
		}
		return x, y, nil
	}
	return -1, -1, errors.New("invalid coordinates, should be in format \"x,y\"")
}
