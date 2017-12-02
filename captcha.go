package day1

import (
	"strconv"
)

func Sum(seq string) int {
	return 0
}

func SumChars(prev, cur rune) int {
	if prev == cur {
		dig, err := strconv.Atoi(strconv.QuoteRune(prev))
		if err != nil {
			return 0
		}
		return dig
	}
	return 0
}
