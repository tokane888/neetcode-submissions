func longestConsecutive(nums []int) int {
	numsSet := make(map[int]struct{})
	for _, num := range nums {
		numsSet[num] = struct{}{}
	}

	if len(nums) == 0 {
		return 0
	}
	maxSequenceLen := 1
	for num := range numsSet {
		if _, ok := numsSet[num-1]; ok {
			continue
		}
		count := 1
		for {
			if _, ok := numsSet[num+count]; ok {
				count++
				if count > maxSequenceLen {
					maxSequenceLen = count
				}
			} else {
				break
			}
		}
	}

	return maxSequenceLen
}
