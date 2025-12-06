package part02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trekawek/advent-of-code/aoc-2025/day06/part01"
)

func reduce(s []int, o part01.Operator) int {
	r := o.Id
	for _, n := range s {
		r = o.Func(r, n)
	}
	return r
}

func iterateOverChars(lines []string) func(yield func(rune) bool) {
	return func(yield func(rune) bool) {
		width := len(lines[0])
		for i := width - 1; i >= 0; i-- {
			if !yield(' ') {
				return
			}
			for _, line := range lines {
				runes := []rune(line)
				if i >= len(runes) {
					continue
				}
				r := runes[i]
				if r == ' ' {
					continue
				}
				// put a space before operator
				isDigit := r >= '0' && r <= '9'
				if isDigit {
					if !yield(r) {
						return
					}
				} else {
					if !(yield(' ') && yield(r) && yield(' ') && yield('X')) {
						return
					}
				}
			}
		}
	}
}

func eval(str string) int {
	s := strings.Split(str, " ")
	var numbers []int
	for i, e := range s {
		if i == len(s)-1 {
			break
		}
		n, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	o := s[len(s)-1]
	if len(o) != 1 {
		panic(fmt.Errorf("invalid operator: %s", o))
	}

	op := part01.NewOperator([]rune(o)[0])
	return reduce(numbers, op)
}

func Solve(lines []string) int {
	b := strings.Builder{}
	for c := range iterateOverChars(lines) {
		b.WriteRune(c)
	}

	s := 0
	for _, e := range strings.Split(b.String(), "X") {
		t := strings.TrimSpace(e)
		if t == "" {
			continue
		}
		s += eval(t)
	}
	return s
}
