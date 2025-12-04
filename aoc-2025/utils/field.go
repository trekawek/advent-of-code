package utils

import "strings"

type Field [][]rune

func (f *Field) Get(x, y int) rune {
	return (*f)[y][x]
}

func (f *Field) Set(x, y int, r rune) {
	(*f)[y][x] = r
}

func (f *Field) Width() int {
	return len((*f)[0])
}

func (f *Field) Height() int {
	return len(*f)
}

func (f *Field) Iterate(yield func(int, int) bool) {
	for x := range f.Width() {
		for y := range f.Height() {
			if !yield(x, y) {
				return
			}
		}
	}
}

func (f *Field) IterateAdjacent(x, y int) func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for _, a := range [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}} {
			i := x + a[0]
			j := y + a[1]
			if i < 0 || j < 0 || i >= f.Width() || j >= f.Height() {
				continue
			}
			if !yield(i, j) {
				break
			}
		}
	}
}

func (f *Field) String() string {
	b := strings.Builder{}
	for _, l := range *f {
		b.WriteString(string(l))
		b.WriteRune('\n')
	}
	return b.String()
}

func ParseField(lines []string) *Field {
	var f [][]rune
	for _, l := range lines {
		f = append(f, []rune(l))
	}
	g := Field(f)
	return &g
}
