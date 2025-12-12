package part02

import (
	"fmt"
	"maps"
	"slices"

	"github.com/trekawek/advent-of-code/aoc-2025/day09/part01"
)

type field struct {
	points []part01.Pair
	field  [][]rune

	xs []int
	ys []int
	sx map[int]int
	sy map[int]int
}

func newField(points []part01.Pair) *field {
	f := field{
		points: points,
	}
	f.fillXs()
	f.fillSx()
	f.initField()
	f.drawLines()
	return &f
}

func (f *field) fillXs() {
	xsm := map[int]bool{}
	ysm := map[int]bool{}
	for _, p := range f.points {
		xsm[p.X] = true

		ysm[p.Y] = true
	}
	f.xs = slices.Collect(maps.Keys(xsm))
	f.ys = slices.Collect(maps.Keys(ysm))
	slices.Sort(f.xs)
	slices.Sort(f.ys)
}

func (f *field) fillSx() {
	f.sx = map[int]int{}
	f.sy = map[int]int{}
	for i, x := range f.xs {
		f.sx[x] = i
	}
	for j, y := range f.ys {
		f.sy[y] = j
	}
}

func (f *field) initField() {
	for range f.ys {
		var line []rune
		for range f.xs {
			line = append(line, '.')
		}
		f.field = append(f.field, line)
	}
}

func (f *field) drawLines() {
	for i, p := range f.points {
		q := f.points[(i+1)%len(f.points)]
		px := f.sx[p.X]
		py := f.sy[p.Y]
		qx := f.sx[q.X]
		qy := f.sy[q.Y]

		if px == qx {
			for y := min(py, qy) + 1; y < max(py, qy); y++ {
				f.field[y][px] = 'X'
			}
		} else {
			for x := min(px, qx) + 1; x < max(px, qx); x++ {
				f.field[py][x] = 'X'
			}
		}
		f.field[py][px] = '#'
	}
}

func (f *field) findPointInside() (int, int) {
	for j := range f.ys {
		for i := range f.xs {
			if f.field[j][i] != '.' {
				continue
			}
			reachedBorder := f.fillPolygon(i, j, '.', '?')
			f.clearField()
			if !reachedBorder {
				return i, j
			}
		}
	}
	return -1, -1
}

var directions [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func (f *field) fillPolygon(i, j int, empty, fill rune) bool {
	if f.field[j][i] != empty {
		return false
	}

	f.field[j][i] = fill
	reachedBorder := false
	for _, d := range directions {
		k := i + d[0]
		l := j + d[1]
		if k < 0 || l < 0 || l >= len(f.field) || k >= len(f.field[l]) {
			reachedBorder = true
		} else if f.fillPolygon(k, l, empty, fill) {
			reachedBorder = true
		}
	}
	return reachedBorder
}

func (f *field) clearField() {
	for j := range f.field {
		for i, c := range f.field[j] {
			if c == '?' {
				f.field[j][i] = '.'
			}
		}
	}
}

func (f *field) isRectInside(p1, p2 part01.Pair) bool {
	xmin := min(f.sx[p1.X], f.sx[p2.X])
	xmax := max(f.sx[p1.X], f.sx[p2.X])
	ymin := min(f.sy[p1.Y], f.sy[p2.Y])
	ymax := max(f.sy[p1.Y], f.sy[p2.Y])

	if p1.X == 4 && p2.X == 11 && p1.Y == 10 && p2.Y == 6 {
		fmt.Printf("so?!\n")
	}

	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if f.field[y][x] == '.' {
				return false
			}
		}
	}
	return true
}

func (f *field) printWithRect(p1, p2 part01.Pair) {
	xmin := min(f.sx[p1.X], f.sx[p2.X])
	xmax := max(f.sx[p1.X], f.sx[p2.X])
	ymin := min(f.sy[p1.Y], f.sy[p2.Y])
	ymax := max(f.sy[p1.Y], f.sy[p2.Y])

	var field2 [][]rune
	for j, line := range f.field {
		var line2 []rune
		for i, c := range line {
			if i >= xmin && i <= xmax && j >= ymin && j <= ymax {
				line2 = append(line2, 'O')
			} else {
				line2 = append(line2, c)
			}
		}
		field2 = append(field2, line2)
	}
	field2[f.sy[p1.Y]][f.sx[p1.X]] = 'o'
	field2[f.sy[p2.Y]][f.sx[p2.X]] = 'o'
	print(field2)
}

func print(field [][]rune) {
	for _, line := range field {
		fmt.Printf("%s\n", string(line))
	}
	fmt.Println()
}

func Solve(lines []string) int {
	f := newField(part01.Parse(lines))

	i, j := f.findPointInside()
	if i != -1 {
		f.fillPolygon(i, j, '.', 'X')
	}

	var max int
	for _, p1 := range f.points {
		for _, p2 := range f.points {
			a := part01.Area(p1, p2)
			if f.isRectInside(p1, p2) && a > max {
				max = a
			}
		}
	}
	return max
}
