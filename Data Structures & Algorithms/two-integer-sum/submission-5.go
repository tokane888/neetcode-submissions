func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for idx, num := range nums {
		if _, ok := seen[target-num]; ok {
			return []int{seen[target-num], idx}
		}
		seen[num] = idx
	}

	return nil
}
