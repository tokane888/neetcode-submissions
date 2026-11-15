func twoSum(nums []int, target int) []int {
	// [value, idx] of appeared num
    appearedNum := make(map[int]int)
	// [[0,4], [1,5], [2,6]]
	for idx, num := range nums {
		if prevIdx, ok := appearedNum[target - num]; ok {
			return []int{prevIdx, idx}
		}
		// appearedNum = [[4,0], [5,1], [6,2]]
		appearedNum[num] = idx
	}

	fmt.Println("Something wrong")
	return []int{0,0}
}
