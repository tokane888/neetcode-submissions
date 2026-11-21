func productExceptSelf(nums []int) []int {
	var productOfAll int
	productOfAll = 1
	for _, num := range nums {
		productOfAll *= num 
	}
	result := make([]int, 0, len(nums))
	for idx, num := range nums {
		if num != 0 {
			result = append(result, productOfAll / num)
		} else {
			i := 0
			product := 1
			for i < len(nums) {
				if i != idx {
					product *= nums[i]
				}
				i++
			}
			result = append(result, product)
		}
	}

	return result
}
