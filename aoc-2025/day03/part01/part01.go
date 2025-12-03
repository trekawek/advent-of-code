package part01

import (
	"fmt"
	"slices"
	"strconv"
)

func getMaxDigit(s string) (int, rune) {
	r := slices.Max([]rune(s))
	i := slices.Index([]rune(s), r)
	return i, r
}

func getMaxJolt(s string) int {
	l := len(s)
	i1, d1 := getMaxDigit(s[:l-1])
	_, d2 := getMaxDigit(s[i1+1:])
	r := fmt.Sprintf("%c%c", d1, d2)
	n, err := strconv.Atoi(r)
	if err != nil {
		panic(err)
	}
	return n
}

func Solve(lines []string) int {
	s := 0
	for _, l := range lines {
		s += getMaxJolt(l)
	}
	return s
}
