func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)	// [val, idx]
	for idx, num := range nums {
		if _, ok := hash[target-num]; ok {
			return []int{hash[target - num],  idx}
		}
		hash[num] = idx
	}

	// something wrong
	return nil
}
