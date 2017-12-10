package day6

type memory []int

// NumReallocsCycle will calculate part 1 of the chore, where we determine the
// number of memory reallocations that are needed before we see our first
// repetition.
func NumReallocsCycle(mem []int) int {
	return -1
}

func reallocateSegments(mem []int) []int {
	return []int{}
}

func (m memory) maxIdx() int {
	maxIdx := 0
	for i, v := range m {
		if v > m[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
