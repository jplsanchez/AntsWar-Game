package game

import (
	"testing"
)

var NoCard Card = Card{-1, -1}

func TestAttack_CanDo(t *testing.T) {
	type args struct {
		pos Position
		gb  GameBoard
	}
	tests := []struct {
		name    string
		a       Attack
		args    args
		wantErr bool
	}{
		{
			name: "Attack with no adjacent cards",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
		{
			name: "Attack with enemy adjacent cards with lower value",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: false,
		},
		{
			name: "Attack with enemy adjacent cards with higher value",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{8, 0}, {9, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
		{
			name: "Attack with only allies adjacent cards",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{8, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.CanDo(tt.args.pos, tt.args.gb); (err != nil) != tt.wantErr {
				t.Errorf("Attack.CanDo() error = %v,  want %v", err, tt.wantErr)
			}
		})
	}
}

func TestMove_CanDo(t *testing.T) {
	type args struct {
		pos Position
		gb  GameBoard
	}
	tests := []struct {
		name    string
		m       Move
		args    args
		wantErr bool
	}{
		{
			name: "Move with adjacent empty space",
			m:    Move{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{0, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{1, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: false,
		},
		{
			name: "Move with no adjacent empty space",
			m:    Move{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{0, 0}, {1, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{1, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.CanDo(tt.args.pos, tt.args.gb); (err != nil) != tt.wantErr {
				t.Errorf("Move.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSwap_CanDo(t *testing.T) {
	type args struct {
		pos Position
		gb  GameBoard
	}
	tests := []struct {
		name    string
		s       Swap
		args    args
		wantErr bool
	}{
		{
			name: "Swap with no adjacent cards",
			s:    Swap{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
		{
			name: "Swap with only enemy adjacent cards",
			s:    Swap{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{8, 0}, {9, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
		{
			name: "Swap with allies adjacent cards",
			s:    Swap{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{8, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CanDo(tt.args.pos, tt.args.gb); (err != nil) != tt.wantErr {
				t.Errorf("Swap.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarch_CanDo(t *testing.T) {
	type args struct {
		pos Position
		gb  GameBoard
	}
	tests := []struct {
		name    string
		m       March
		args    args
		wantErr bool
	}{
		{
			name: "March with no adjacent cards",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: true,
		},
		{
			name: "March with adjacent cards vertically",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, NoCard, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "March with adjacent cards vertically and with enemy team value on the empty space",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, NoCard, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "March with adjacent cards horizontally",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{NoCard, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "March with adjacent empty space after enemy card",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, NoCard, NoCard},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.CanDo(tt.args.pos, tt.args.gb); (err != nil) != tt.wantErr {
				t.Errorf("March.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarch_DoAction_Errors(t *testing.T) {
	tests := []struct {
		name    string
		m       March
		wantErr bool
	}{
		{
			name: "March vertically",
			m: March{
				pos:       Position{0, 0},
				targetPos: Position{0, 3},
				gb: &GameBoard{
					{{0, 0}, {2, 0}, {2, 0}, NoCard, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "March horizontally",
			m: March{
				pos:       Position{0, 0},
				targetPos: Position{4, 0},
				gb: &GameBoard{
					{{0, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{NoCard, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: false,
		},
		{
			name: "March diagonally must generate error",
			m: March{
				pos:       Position{0, 0},
				targetPos: Position{4, 1},
				gb: &GameBoard{
					{{0, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, NoCard, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.DoAction(); (err != nil) != tt.wantErr {
				t.Errorf("March.DoAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestMarch_DoAction_BoardMove(t *testing.T) {
	tests := []struct {
		name      string
		m         March
		wantBoard GameBoard
	}{
		{
			name: "March vertically",
			m: March{
				pos:       Position{0, 0},
				targetPos: Position{0, 3},
				gb: &GameBoard{
					{{0, 0}, {1, 0}, {2, 0}, NoCard, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantBoard: GameBoard{
				{NoCard, {0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
			},
		},
		{
			name: "March horizontally",
			m: March{
				pos:       Position{0, 0},
				targetPos: Position{4, 0},
				gb: &GameBoard{
					{{0, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{1, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{3, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{NoCard, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
			},
			wantBoard: GameBoard{
				{NoCard, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{0, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{1, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				{{3, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.DoAction(); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
			if err := tt.m.gb.Equals(tt.wantBoard); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
		})
	}
}

func TestSwap_DoAction(t *testing.T) {
	tests := []struct {
		name      string
		s         Swap
		wantBoard GameBoard
	}{
		{
			name: "swap vertically",
			s: Swap{
				pos:       Position{0, 0},
				targetPos: Position{0, 1},
				gb: &GameBoard{
					{{1, 0}, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{3, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{{2, 0}, {1, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{{3, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
		{
			name: "swap horizontally",
			s: Swap{
				pos:       Position{0, 0},
				targetPos: Position{1, 0},
				gb: &GameBoard{
					{{1, 0}, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{3, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{{3, 0}, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{{1, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DoAction(); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
			if err := tt.s.gb.Equals(tt.wantBoard); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
		})
	}
}

func TestMove_DoAction(t *testing.T) {
	tests := []struct {
		name      string
		m         Move
		wantBoard GameBoard
	}{
		{
			name: "move vertically",
			m: Move{
				pos:       Position{0, 0},
				targetPos: Position{0, 1},
				gb: &GameBoard{
					{{0, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{NoCard, {0, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
		{
			name: "move horizontally",
			m: Move{
				pos:       Position{0, 0},
				targetPos: Position{1, 0},
				gb: &GameBoard{
					{{0, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{{0, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.DoAction(); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
			if err := tt.m.gb.Equals(tt.wantBoard); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
		})
	}
}

func TestAttack_DoAction_BoardMove(t *testing.T) {
	tests := []struct {
		name      string
		a         Attack
		wantBoard GameBoard
	}{
		{
			name: "attack vertically",
			a: Attack{
				pos:       Position{0, 0},
				targetPos: Position{0, 1},
				gb: &GameBoard{
					{{2, 0}, {1, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{NoCard, {2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
		{
			name: "attack horizontally",
			a: Attack{
				pos:       Position{0, 0},
				targetPos: Position{1, 0},
				gb: &GameBoard{
					{{2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{{1, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{{2, 0}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
		{
			name: "attack equal power",
			a: Attack{
				pos:       Position{0, 0},
				targetPos: Position{0, 1},
				gb: &GameBoard{
					{{1, 0}, {1, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantBoard: GameBoard{
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.DoAction(); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
			if err := tt.a.gb.Equals(tt.wantBoard); err != nil {
				t.Errorf("March.DoAction() error = %v", err)
			}
		})
	}
}

func TestAttack_DoAction_Errors(t *testing.T) {
	tests := []struct {
		name    string
		a       Attack
		wantErr bool
	}{
		{
			name: "attack higher power",
			a: Attack{
				pos:       Position{0, 0},
				targetPos: Position{0, 1},
				gb: &GameBoard{
					{{1, 0}, {9, 1}, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
					{NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard, NoCard},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.DoAction(); (err != nil) != tt.wantErr {
				t.Errorf("March.DoAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
