func twoSum(numbers []int, target int) []int {
	l := 0
	lenNum := len(numbers)
	r := lenNum-1
	for l < lenNum {
		for r > l {
			if numbers[l] + numbers[r] == target {
				return []int{l+1,r+1}
			}
			if numbers[l] + numbers[r] < target {
				l++
			} else {
				r--
			}
		}
	}

	return nil
}
