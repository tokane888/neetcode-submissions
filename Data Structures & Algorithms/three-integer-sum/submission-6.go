import "slices"

func threeSum(nums []int) [][]int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	count := make(map[int]int)	// num => count
	for _, num := range sorted {
		count[num]++
	}

	i := 0
	result := [][]int{}
	for i < len(sorted) {
		count[sorted[i]]--
		if i > 0 && sorted[i] == sorted[i-1] {
			i++
			continue
		}
		for j := i + 1; j < len(sorted) - 1; j++ {
			count[sorted[j]]--
			if j > i + 1 && sorted[j] == sorted[j-1] {
				continue
			}
			target := -(sorted[i] + sorted[j])
			if count[target] != 0 {
				result = append(result, []int{sorted[i], sorted[j], target})
			}
		}

		for j := i + 1; j < len(sorted) - 1; j++ {
			count[sorted[j]]++
		}
		i++
	}

	return result
}
