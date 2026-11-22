import "slices"

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	maxSequence := 0
	sequenceCount := 0
	previousNum := 0
	for idx, num := range nums {
		if idx != 0 {
			if previousNum == num {
				continue
			}
			if previousNum + 1 != num {
				if sequenceCount > maxSequence {
					maxSequence = sequenceCount
				}
				sequenceCount = 0
			}
		}
		previousNum = num
		sequenceCount++
	}
	if sequenceCount > maxSequence {
		maxSequence = sequenceCount
	}

	return maxSequence
}
