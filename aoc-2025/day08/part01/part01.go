package part01

import (
	"cmp"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Box [3]int

type Pair struct {
	B1 int
	B2 int
	d  float64
}

func (b *Box) dist(c *Box) float64 {
	d := 0
	for i := range 3 {
		d += (b[i] - c[i]) * (b[i] - c[i])
	}
	return math.Sqrt(float64(d))
}

func Parse(lines []string) []Box {
	var result []Box
	for _, l := range lines {
		var p Box
		for i, e := range strings.Split(l, ",") {
			n, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			p[i] = n
		}
		result = append(result, p)
	}
	return result
}

func GetSortedPairs(boxes []Box) []Pair {
	var pairs []Pair
	for i := range len(boxes) {
		b1 := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			b2 := boxes[j]
			pairs = append(pairs, Pair{
				i, j, b1.dist(&b2),
			})
		}
	}
	slices.SortFunc(pairs, func(p1, p2 Pair) int {
		return cmp.Compare(p1.d, p2.d)
	})
	return pairs
}

func InitializeSets(length int) []map[int]bool {
	var connections []map[int]bool
	for i := range length {
		m := map[int]bool{i: true}
		connections = append(connections, m)
	}
	return connections
}

func getUniqueSets(sets []map[int]bool) []map[int]bool {
	var uniqueSets []map[int]bool
	for i, c := range sets {
		if i == min(c) {
			uniqueSets = append(uniqueSets, c)
		}
	}
	return uniqueSets
}

func Merge(connections []map[int]bool, i, j int) bool {
	if connections[i][j] {
		return false
	}
	n := map[int]bool{}
	for e := range connections[i] {
		n[e] = true
	}
	for e := range connections[j] {
		n[e] = true
	}
	for e := range n {
		connections[e] = n
	}
	return true
}

func min(m map[int]bool) int {
	return slices.Min(slices.Collect(maps.Keys(m)))
}

func Solve(lines []string, pairsCount int) int {
	boxes := Parse(lines)
	sets := InitializeSets(len(boxes))
	pairs := GetSortedPairs(boxes)

	for i := range pairsCount {
		p := pairs[i]
		Merge(sets, p.B1, p.B2)
	}

	var setSizes []int
	for _, c := range getUniqueSets(sets) {
		setSizes = append(setSizes, len(c))
	}
	slices.Sort(setSizes)
	slices.Reverse(setSizes)

	s := 1
	for i := range 3 {
		s *= setSizes[i]
	}
	return s
}
