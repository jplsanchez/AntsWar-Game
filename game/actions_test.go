package game

import (
	"testing"
)

func TestAttack_CanDo(t *testing.T) {
	type args struct {
		pos  Position
		gb   GameBoard
		team Team
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
					{{9, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
		{
			name: "Attack with enemy adjacent cards with lower value",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 1}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: false,
		},
		{
			name: "Attack with enemy adjacent cards with higher value",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{8, 0}, {9, 1}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
		{
			name: "Attack with only allies adjacent cards",
			a:    Attack{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{8, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.CanDo(tt.args.pos, tt.args.gb, tt.args.team); (err != nil) != tt.wantErr {
				t.Errorf("Attack.CanDo() error = %v,  want %v", err, tt.wantErr)
			}
		})
	}
}

func TestMove_CanDo(t *testing.T) {
	type args struct {
		pos  Position
		gb   GameBoard
		team Team
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
					{{0, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: false,
		},
		{
			name: "Move with no adjacent empty space",
			m:    Move{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{0, 0}, {1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.CanDo(tt.args.pos, tt.args.gb, tt.args.team); (err != nil) != tt.wantErr {
				t.Errorf("Move.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSwap_CanDo(t *testing.T) {
	type args struct {
		pos  Position
		gb   GameBoard
		team Team
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
					{{9, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
		{
			name: "Swap with only enemy adjacent cards",
			s:    Swap{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{8, 0}, {9, 1}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: true,
		},
		{
			name: "Swap with allies adjacent cards",
			s:    Swap{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{9, 0}, {2, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{8, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
					{{-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}, {-1, 0}},
				},
				team: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CanDo(tt.args.pos, tt.args.gb, tt.args.team); (err != nil) != tt.wantErr {
				t.Errorf("Swap.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarch_CanDo(t *testing.T) {
	type args struct {
		pos  Position
		gb   GameBoard
		team Team
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
				team: 0,
			},
			wantErr: true,
		},
		{
			name: "March with adjacent cards vertically",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {-1, 0}, {2, 0}, {-1, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
				team: 0,
			},
			wantErr: false,
		},
		{
			name: "March with adjacent cards vertically and with enemy team value on the empty space",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, {-1, 1}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
				team: 0,
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
					{{-1, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
				team: 0,
			},
			wantErr: false,
		},
		{
			name: "March with adjacent empty space after enemy card",
			m:    March{},
			args: args{
				pos: Position{0, 0},
				gb: GameBoard{
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {-1, 1}, {-1, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
					{{2, 0}, {2, 0}, {2, 0}, {2, 0}, {2, 1}, {2, 1}, {2, 1}, {2, 1}},
				},
				team: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.CanDo(tt.args.pos, tt.args.gb, tt.args.team); (err != nil) != tt.wantErr {
				t.Errorf("March.CanDo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
