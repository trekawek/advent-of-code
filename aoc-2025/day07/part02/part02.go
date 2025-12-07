package part02

func Solve(lines []string) int {
	beams := map[int]int{}
	for _, line := range lines {
		for i, c := range line {
			switch c {
			case '.':
				continue
			case 'S':
				beams[i] = 1
			case '^':
				if beams[i] > 0 {
					beams[i-1] += beams[i]
					beams[i+1] += beams[i]
					beams[i] = 0
				}
			}
		}
	}

	count := 0
	for _, v := range beams {
		count += v
	}
	return count
}
