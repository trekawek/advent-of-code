package utils

import (
	"os"
	"strings"
)

func ReadAllLines(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(input)), "\n")
}
