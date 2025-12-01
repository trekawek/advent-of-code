package utils

func Abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func Mod(a, d int) int {
	var r = a % d
	if r < 0 {
		return r + d
	}
	return r
}
