package part02

import (
	"slices"
	"strconv"
	"strings"
)

const LENGTH = 12

func getMaxDigit(s string, fromIndex int) (int, rune) {
	t := []rune(s[fromIndex:])
	r := slices.Max(t)
	i := slices.Index(t, r) + fromIndex
	return i, r
}

func getMaxJolt(s string) uint64 {
	l := len(s)

	b := strings.Builder{}
	i := 0
	maxIndex := 0

	for i < 12 {
		j, d := getMaxDigit(s[:l-(LENGTH-i-1)], maxIndex)
		b.WriteRune(d)
		maxIndex = j + 1
		i++
	}

	n, err := strconv.ParseUint(b.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func Solve(lines []string) uint64 {
	var s uint64
	for _, l := range lines {
		s += getMaxJolt(l)
	}
	return s
}
