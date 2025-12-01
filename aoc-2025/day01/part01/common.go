package part01

import (
	"fmt"
)

const START = 50
const MAX = 100

type Move struct {
	d rune
	v int
}

func (m *Move) String() string {
	return fmt.Sprintf("%c%d", m.d, m.v)
}

func ParseMoves(lines []string) []Move {
	var moves []Move
	for _, l := range lines {
		m := Move{}
		fmt.Sscanf(l, "%c%d", &m.d, &m.v)
		moves = append(moves, m)
	}
	return moves
}
