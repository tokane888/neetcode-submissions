func productExceptSelf(nums []int) []int {
	product := 1
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		result = append(result, product)
		product *= num
	}

	product = 1
	for idx := len(nums)-1; idx >= 0; idx-- {
		result[idx] *= product
		product *= nums[idx]
	}

	return result
}
