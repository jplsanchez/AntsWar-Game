package game

type Position struct {
	x int
	y int
}

type GameManager struct {
	turns        int
	teamState    Team
	cardSelected Position
}

func NewGameManager(startingTeam Team) *GameManager {
	gm := GameManager{turns: 0, teamState: startingTeam}
	return &gm
}

func (gm *GameManager) ChooseCard(x, y int) {
	gm.cardSelected = Position{x: x, y: y}
}
