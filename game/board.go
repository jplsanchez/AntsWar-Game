package game

const Middle_Col_Index = 2

type GameBoard [5][8]Card

func (b GameBoard) Width() int {
	return len(b)
}

func (b GameBoard) Height() int {
	return len(b[0])
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
