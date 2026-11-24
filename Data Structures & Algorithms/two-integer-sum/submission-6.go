func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)	// num => index
	for i, num := range nums {
		if _, ok := seen[target-num]; ok {
			return []int{seen[target-num], i}
		}
		seen[num] = i
	}

	return nil
}
