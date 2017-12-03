package day1

import (
	"strconv"
)

// SumSequence will sum a sequence of digits based on the rule that if the current digit
// is the same as the previous one, then it can be summed.  The collection of these
// is summed together.
func SumSequence(seq string) int {
	jump := 1
	return SumOverRange(seq, jump)
}

// SumRange takes an array of []int's and sums them.
func SumRange(summable []int) int {
	sum := 0
	for _, v := range summable {
		sum += v
	}
	return sum
}

// IsSame determines if the two runes are equal.
func IsSame(prev, cur rune) bool {
	return prev == cur
}

// Convert will convert a rune to an integer digit, returning 0 if it errors.
func Convert(digit rune) int {
	d, err := strconv.Atoi(string(digit))
	if err != nil {
		return 0
	}
	return d
}

// SumChars will take two runes, convert them to digits, and sum them together.
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

// SumOverRange implements part 2 of the day 1 stuff.
func SumOverRange(seq string, jump int) int {
	var summable []int
	l := len(seq)
	for i := 0; i < l; i++ {
		first := rune(seq[i])
		second := rune(seq[(i+jump)%l])
		if first == second {
			summable = append(summable, Convert(first))
		}
	}
	return SumRange(summable)
}

// SumSequenceMod2 is like SumSequence, but will sum and calculate based on the second part of the
// puzzle.
func SumSequenceMod2(seq string) int {
	jump := len(seq) / 2
	return SumOverRange(seq, jump)
}
