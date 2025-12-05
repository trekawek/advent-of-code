package utils

import "fmt"

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

func (r *Range) Contains(i uint64) bool {
	return i >= r.From && i <= r.To
}

func ParseRange(s string) *Range {
	r := Range{}
	n, err := fmt.Sscanf(s, "%d-%d", &r.From, &r.To)
	if err != nil {
		panic(err)
	}
	if n == 0 {
		return nil
	}
	return &r
}
