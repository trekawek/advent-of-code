package part01

import (
	"fmt"
	"strconv"
	"strings"
)

type Operator struct {
	r    string
	Func func(int, int) int
	Id   int
}

func NewOperator(r rune) Operator {
	switch r {
	case '+':
		return Operator{string(r), func(a, b int) int { return a + b }, 0}
	case '*':
		return Operator{string(r), func(a, b int) int { return a * b }, 1}
	default:
		panic(fmt.Errorf("Invalid operator %c", r))
	}
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func removeEmpty(s []string) []string {
	var t []string
	for _, e := range s {
		if e != "" {
			t = append(t, e)
		}
	}
	return t
}

func Parse(lines []string) ([][]int, []Operator) {
	var numbers [][]int
	var operators []Operator

	for _, l := range lines {
		k := strings.Split(l, " ")
		k = removeEmpty(k)
		if isNumber(k[0]) {
			var m []int
			for _, n := range k {
				if n == "" {
					continue
				}
				n, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				m = append(m, n)
			}
			numbers = append(numbers, m)
		} else {
			for _, o := range k {
				if o == "" {
					continue
				}
				if len(o) > 1 {
					panic(fmt.Errorf("invalid operator: %s", o))
				}
				operators = append(operators, NewOperator([]rune(o)[0]))
			}
		}
	}
	return numbers, operators
}

func Solve(lines []string) int {
	numbers, operators := Parse(lines)

	var results []int
	for i, o := range operators {
		r := o.Id
		for _, n := range numbers {
			r = o.Func(r, n[i])
		}
		results = append(results, r)
	}

	var result int
	for _, r := range results {
		result += r
	}

	return result
}
