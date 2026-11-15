func hasDuplicate(nums []int) bool {
	appeared := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := appeared[num]; ok {
			return true
		}
		appeared[num] = struct{}{}
	}

	return false
}
