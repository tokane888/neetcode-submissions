func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)	// value => idx
	for idx, num := range nums {
		if prevIdx, ok := seen[target - num]; ok {
			return []int{prevIdx, idx}
		}
		seen[num] = idx
	}

	// Something wrong
	return nil
}
