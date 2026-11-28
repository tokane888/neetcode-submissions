import "slices"

func threeSum(nums []int) [][]int {
	// nums=[-1,0,1,2,-1,-4,-2,-3,3,0,4]
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	count := make(map[int]int)	// num => count
	for _, num := range sorted {
		count[num]++
	}

	i := 0
	result := [][]int{}
	// sorted=[-4,-3,-2,-1,-1,0,0,1,2,3,4]
	// len(sorted) = 11
	for i < len(sorted) {
		// count
		// -4 => 1
		// -3 => 1
		// -2 => 1
		// -1 => 2
		// 0  => 2
		// 1  => 1
		// 2  => 1
		// 3  => 1
		// 4  => 1
		count[sorted[i]]--
		if i > 0 && sorted[i] == sorted[i-1] {
			i++
			continue
		}
		// loop2: i = 1, j = 2
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
