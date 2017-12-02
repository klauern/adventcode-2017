package day1

import (
	"strconv"
)

func Sum(seq string) int {
	var summable []int
	for v, _ := range seq[1:] {
		if IsSame(rune(seq[v]), rune(seq[v-1])) {
			summable = append(summable, Convert(rune(seq[v])))
		}
	}

	// check last digit against first (circular)
	if IsSame(rune(seq[len(seq)-1]), rune(seq[0])) {
		summable = append(summable, Convert(rune(seq[0])))
	}

	sum := 0
	for _, v := range summable {
		sum += v
	}
	return sum
}

func IsSame(prev, cur rune) bool {
	if prev == cur {
		return true
	}
	return false
}

func Convert(digit rune) int {
	d, err := strconv.Atoi(string(digit))
	if err != nil {
		return 0
	}
	return d
}

func SumChars(first, second rune) int {
	f, err := strconv.Atoi(string(first))
	if err != nil {
		return 0
	}
	s, err := strconv.Atoi(string(second))
	if err != nil {
		return 0
	}
	return f + s
}
