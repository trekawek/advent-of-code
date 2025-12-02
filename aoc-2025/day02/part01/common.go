package part01

import (
	"fmt"
	"strings"
)

type Range struct {
	From uint64
	To   uint64
}

func (r *Range) Iterate(yield func(uint64) bool) {
	i := r.From
	for i <= r.To {
		if !yield(i) {
			break
		}
		i++
	}
}

func ParseLines(lines []string) []Range {
	var ranges []Range
	for _, l := range lines {
		for _, v := range strings.Split(l, ",") {
			w := strings.TrimSpace(v)
			if w == "" {
				continue
			}
			r := Range{}
			fmt.Sscanf(w, "%d-%d", &r.From, &r.To)
			ranges = append(ranges, r)
		}
	}
	return ranges
}
