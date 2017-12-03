package day2

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const spreadsheet = `5	1	9	5
7	5	3
2	4	6	8`

const partTwoSpreadsheet = `5	9	2	8
9	4	7	3
3	8	6	5`

// DifferenceLargestSmallest will find the largest and smallest numbers in a string,
// and apply a difference on them to find the largest difference value.
func DifferenceLargestSmallest(line string) int {
	ints := parseLine(line, "\t")
	sort.Ints(ints)
	return ints[len(ints)-1] - ints[0]
}

func parseLine(line, splitBy string) []int {
	var values []int

	s := strings.Split(line, splitBy)
	// TODO: write parser

	for _, v := range s {
		val, err := strconv.Atoi(v)
		if err == nil {
			values = append(values, val)
		}
	}
	return values
}

// ChecksumTestSpreadsheet performs the checksum for the example `spreadsheet` value above.
func ChecksumTestSpreadsheet() int {
	return checksumString(spreadsheet)
}

func checksumString(str string) int {
	lines := strings.Split(str, "\n")
	tot := 0
	for _, v := range lines {
		tot += DifferenceLargestSmallest(v)
	}

	return tot
}

// ChecksumInput will checksum the input.txt file.
func ChecksumInput() int {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return -1
	}
	str := string(file)
	cksum := checksumString(str)
	fmt.Println("Checksum of input.txt is ", cksum)
	return cksum
}

func findEvenlyDivisible(str string) []int {
	ints := parseLine(str, "\t")
	for _, v := range ints {
		for _, u := range ints[1:] {
			if v > u && v%u == 0 {
				return []int{v, u}
			}
			if u > v && u%v == 0 {
				return []int{v, u}
			}
		}
	}
	return []int{}
}

func getDivisibleChecksumForLine(str string) int {
	ints := findEvenlyDivisible(str)

	if ints[0] > ints[1] {
		return ints[0] / ints[1]
	}
	return ints[1] / ints[0]
}

func divisibleChecksumString(str string) int {
	lines := strings.Split(str, "\n")
	tot := 0
	for _, v := range lines {
		tot += getDivisibleChecksumForLine(v)
	}

	return tot
}

func getChecksumFromPartTwo() int {
	return divisibleChecksumString(partTwoSpreadsheet)
}

func getDivisibleChecksumFromInput() int {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return 0
	}
	str := string(file)
	return divisibleChecksumString(str)
}
