package display

import "antswar/game"

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
	for !ui.Game.States.GameFinished {
		ui.Game.NextGameStage()

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
	}
}

func (ui *UIManager) StartPlayersTurnRoutine() {}
func (ui *UIManager) ChooseCardRoutine()       {}
func (ui *UIManager) ChooseActionRoutine()     {}
func (ui *UIManager) ResolveActionRoutine()    {}
func (ui *UIManager) EndTurnRoutine()          {}
