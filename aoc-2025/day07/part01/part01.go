package part01

func Solve(lines []string) int {
	beams := map[int]bool{}
	count := 0
	for _, line := range lines {
		for i, c := range line {
			switch c {
			case '.':
				continue
			case 'S':
				beams[i] = true
			case '^':
				if beams[i] {
					beams[i] = false
					beams[i-1] = true
					beams[i+1] = true
					count++
				}
			}
		}
	}
	return count
}
