package utils

import (
	"os"
	"strconv"
)

func GetPartNumber() int {
	if len(os.Args) != 2 {
		panic("Please pass part number")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	return i
}
