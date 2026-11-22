func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0
	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}

		count := 1
		for {
			if _, ok := set[num+count]; ok {
				count++
			} else {
				longest = max(longest, count)
				break
			}
		}
	}

	return longest
}
