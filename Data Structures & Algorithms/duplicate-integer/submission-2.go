func hasDuplicate(nums []int) bool {
    appeared := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := appeared[num]; exists {
			return true
		}
		appeared[num] = struct{}{}
	}
	return false
}
