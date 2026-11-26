func twoSum(numbers []int, target int) []int {
	// binary search
	// [i, iNum, j, k, mid] = 0,2,1,2,1
	// [i, iNum, j, k, mid] = 0,2,2,2,1
	for i, iNum := range numbers {
		j := i + 1
		k := len(numbers) - 1
		for {
			mid := (j + k) / 2
			if iNum + numbers[mid] < target {
				if j == mid + 1 {
					break
				}
				j = mid + 1
			} else if iNum + numbers[mid] > target {
				if k == mid {
					break
				}
				k = mid
			} else if iNum + numbers[mid] == target {
				return []int{i+1, mid+1}
			}
		}
	}

	return nil
}
