package game

import "fmt"

const Middle_Col_Index = 2

type GameBoard [5][8]Card

func (b GameBoard) Width() int {
	test := len(b)
	return test
}

func (b GameBoard) Height() int {
	test := len(b[0])
	return test
}

func (b *GameBoard) LoopThroughBoard(action func(i, j int)) {
	for i := 0; i < b.Width(); i++ {
		for j := 0; j < b.Height(); j++ {
			action(i, j)
		}
	}
}

func (b *GameBoard) LoadEmptyBoard() {

	b.LoopThroughBoard(func(i, j int) { b[i][j].Value = 0 })
}

func (b *GameBoard) LoadInitialBoard() {
	redDeck := NewDeck(TeamRed).Shuffle()
	blackDeck := NewDeck(TeamBlack).Shuffle(42)

	b.LoopThroughBoard(func(i, j int) {
		if b.TrySetEmptySpace(i, j) {
			return
		}

		if j >= b.Height()/2 {
			b[i][j] = redDeck.DrawCard()
			return
		}
		if j < b.Height()/2 {
			b[i][j] = blackDeck.DrawCard()
			return
		}
		b[i][j].Value = -1
	})
}

func (b *GameBoard) TrySetEmptySpace(i int, j int) bool {
	if j == 0 && i == Middle_Col_Index {
		b[i][j].Value = -1
		return true
	}
	if j == b.Height()-1 && i == Middle_Col_Index {
		b[i][j].Value = -1
		return true
	}
	return false
}

func (b GameBoard) CountQueens() []Team {
	queensAlive := make([]Team, 0)
	b.LoopThroughBoard(func(i, j int) {
		if b[i][j].Value == 0 {
			queensAlive = append(queensAlive, b[i][j].Team)
		}
	})
	return queensAlive
}

func (b GameBoard) GetCard(pos Position) Card {
	return b[pos.x][pos.y]
}

func (b *GameBoard) SetCard(pos Position, card Card) {
	b[pos.x][pos.y] = card
}

func (b GameBoard) Equals(comparison GameBoard) error {
	for i := range b {
		for j := range b[i] {
			if b[i][j] != comparison[i][j] {
				return fmt.Errorf("boards are not equal at i=%v, j=%v", i, j)
			}
		}
	}
	return nil
}
