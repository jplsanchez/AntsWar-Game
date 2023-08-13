package display

import (
	"antswar/game"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
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
	ui.Disp.Init()
	ui.Disp.SetHighlight(-1, -1)
	ui.Disp.UpdateBoard(*ui.Game.Board)
	ui.Disp.SetPlayer(ui.Game.States.Team)

	for !ui.Game.States.GameFinished {
		switch ui.Game.States.Stage {
		case game.StartTurnStage:
			ui.startPlayersTurnRoutine()
			ui.Disp.DisplayMessage("It's your turn! Team:" + ui.Game.States.Team.String())
			ui.Game.NextGameStage()

		case game.ChooseCardStage:
			if err := ui.chooseCardRoutine(); err != nil {
				ui.handleRoutineError(err)
			} else {
				ui.Game.NextGameStage()
			}

		case game.ChooseActionStage:
			if err := ui.chooseActionRoutine(); err != nil {
				ui.handleRoutineError(err)
			} else {
				ui.Game.NextGameStage()
			}

		case game.ResolveActionStage:
			if err := ui.resolveActionRoutine(); err != nil {
				ui.handleRoutineError(err)
			} else {
				ui.Game.NextGameStage()
			}

		case game.EndTurnStage:
			ui.endTurnRoutine()
			ui.Game.NextGameStage()
		}
	}
}

func (ui *UIManager) handleRoutineError(err error) {
	if err.Error() == "back" {
		ui.Game.ReturnGameStage()
		return
	}
	ui.Disp.DisplayMessage(err.Error())
}

func (ui *UIManager) startPlayersTurnRoutine() {
	ui.Game.StartTurn()
	ui.Disp.SetHighlight(-1, -1)
	ui.Disp.SetPlayer(ui.Game.States.Team)
	ui.Disp.UpdateBoard(*ui.Game.Board)
}

func (ui *UIManager) chooseCardRoutine() error {
	x, y, err := ui.getCoordinatesFromUser("Choose a card to play")
	if err != nil {
		log.Println(err)
		return err
	}

	err = ui.Game.ChooseCard(x, y)
	if err != nil {
		log.Println(err)
		return err
	}

	ui.Disp.SetHighlight(x, y)
	ui.Disp.UpdateBoard(*ui.Game.Board)
	return nil
}

func (ui *UIManager) chooseActionRoutine() error {
	action, err := ui.Disp.AskForString("Choose an action to play (move, swap, march, attack)")
	if err != nil {
		log.Println(err)
		return err
	}
	if strings.Contains(action, "back") {
		return errors.New("back")
	}

	switch {
	case strings.Contains(action, "move"):
		ui.Game.States.ActionSelected = &game.Move{}
	case strings.Contains(action, "swap"):
		ui.Game.States.ActionSelected = &game.Swap{}
	case strings.Contains(action, "march"):
		ui.Game.States.ActionSelected = &game.March{}
	case strings.Contains(action, "attack"):
		ui.Game.States.ActionSelected = &game.Attack{}
	default:
		return errors.New("invalid action")
	}

	err = ui.Game.ChooseAction(ui.Game.States.ActionSelected, ui.Game.States.SelectedPosition)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ui *UIManager) resolveActionRoutine() error {
	x, y, err := ui.getCoordinatesFromUser("Choose a target")
	if err != nil {
		log.Println(err)
		return err
	}

	err = ui.Game.ResolveAction(x, y)
	if err != nil {
		log.Println(err)
		return err
	}

	ui.Disp.UpdateBoard(*ui.Game.Board)
	return nil
}

func (ui *UIManager) endTurnRoutine() {
	message := ui.Game.CheckIfGameFinished()
	if message != "" {
		ui.Disp.DisplayMessage(message)
	}
}

func (ui UIManager) getCoordinatesFromUser(messages ...string) (int, int, error) {
	str, err := ui.Disp.AskForString(messages...)
	if err != nil {
		return -1, -1, err
	}
	if str == "back" {
		return -1, -1, errors.New("back")
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
