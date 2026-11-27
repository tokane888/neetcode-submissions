func twoSum(numbers []int, target int) []int {
	for i, iNum := range numbers {
		j := i + 1
		k := len(numbers)-1
		for j <= k {
			mid := j + (k - j) / 2
			switch {
				case iNum + numbers[mid] == target:
					return []int{i+1, mid+1}
				case iNum + numbers[mid] < target:
					j = mid + 1
				default:
					k = mid - 1
			}
		}
	}

	return nil
}
