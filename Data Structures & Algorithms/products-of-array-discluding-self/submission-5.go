func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))
	product := 1
	for _, num := range nums {
		result = append(result, product)
		product *= num
	}

	product = 1
	for i := len(nums)-1; i >= 0; i-- {
		result[i] = result[i] * product
		product *= nums[i]
	}

	return result
}
