package part01

import (
	"fmt"
	"strings"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func ParseLines(lines []string) []utils.Range {
	var ranges []utils.Range
	for _, l := range lines {
		for _, v := range strings.Split(l, ",") {
			w := strings.TrimSpace(v)
			if w == "" {
				continue
			}
			r := utils.Range{}
			fmt.Sscanf(w, "%d-%d", &r.From, &r.To)
			ranges = append(ranges, r)
		}
	}
	return ranges
}
