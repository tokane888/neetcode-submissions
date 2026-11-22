func productExceptSelf(nums []int) []int {
	// TODO: Think about 0 case

	productsFromLeft := make(map[int]int)	// idx => product from left to idx-1
	product := 1
	for idx, num := range nums {
		productsFromLeft[idx] = product
		product *= num
	}

	productsFromRight := make(map[int]int)	// idx => product from right to idx+1
	product = 1
	for idx := len(nums)-1; idx >= 0; idx-- {
		productsFromRight[idx] = product
		product *= nums[idx]
	}

	result := make([]int, 0, len(nums))
	for idx, _ := range nums {
		result = append(result, productsFromLeft[idx]*productsFromRight[idx])
	}

	return result
}
