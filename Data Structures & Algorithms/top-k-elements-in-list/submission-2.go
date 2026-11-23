import "slices"

func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)	// num => count
	for _, num := range nums {
		numCount[num]++
	}

	result := make([]int, 0, len(numCount))
	for num := range numCount {
		result = append(result, num)
	}
	slices.SortFunc(result, func(a, b int) int {
		return numCount[b] - numCount[a]
	})

	return result[:k]
}
