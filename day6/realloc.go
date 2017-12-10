package day6

type memory []int
type memMap map[string]bool

// NumReallocsCycle will calculate part 1 of the chore, where we determine the
// number of memory reallocations that are needed before we see our first
// repetition.
func NumReallocsCycle(mem memory) int {
	return -1
}

func reallocateSegments(mem memory) memory {
	return memory{}
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
