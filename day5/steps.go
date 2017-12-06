package day5

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func loadInput(steps string) []int {
	sary := strings.Split(steps, "\n")
	stepAry := make([]int, len(sary))
	for i, v := range sary {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		stepAry[i] = val
	}
	return stepAry
}

func howManySteps(steps []int) int {
	i := 0
	stepCount := 0

	for {
		stepCount++
		jump := steps[i]
		steps[i] = steps[i] + 1
		i += jump
		if i < 0 || i >= len(steps) {
			break
		}
		// get the value to jump ahead/behind for
		// ++ the current value in steps[i]
		// change i to be i + current
	}
	return stepCount
}

func testStepsInput() int {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return -1
	}
	steps := loadInput(string(f))
	return howManySteps(steps)
}

func howManyStepsWithCaveats(steps []int) int {
	i := 0
	stepCount := 0

	for {
		stepCount++
		jump := steps[i]
		if jump >= 3 || jump <= -3 {
			steps[i] = steps[i] - 1
		} else {
			steps[i] = steps[i] + 1
		}
		i += jump
		if i < 0 || i >= len(steps) {
			break
		}
		// get the value to jump ahead/behind for
		// ++ the current value in steps[i]
		// change i to be i + current
	}
	return stepCount
}

func testStepsInputPart2() int {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return -1
	}
	steps := loadInput(string(f))
	return howManyStepsWithCaveats(steps)
}
