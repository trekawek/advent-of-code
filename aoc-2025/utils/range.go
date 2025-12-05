package utils

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
