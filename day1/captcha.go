package day1

import (
	"strconv"
)

func SumSequence(seq string) int {
	var summable []int
	for v := 1; v < len(seq); v++ {
		second := rune(seq[v])
		first := rune(seq[v-1])
		if IsSame(first, second) {
			summable = append(summable, Convert(second))
		}
	}

	// check last digit against first (circular)
	first := rune(seq[0])
	last := rune(seq[len(seq)-1])
	if IsSame(first, last) {
		summable = append(summable, Convert(last))
	}

	return SumRange(summable)

}

func SumRange(summable []int) int {
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
