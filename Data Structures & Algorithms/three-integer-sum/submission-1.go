import "slices"

func threeSum(nums []int) [][]int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	results := make([][]int, 0)
	i := 0
	// data1: [0,0,0,0]
	// loop1: i,j,k,sum = 0,1,3,0
	// loop2: i,j,k,sum = 1,2,2,
	//   i = 4
	for i < len(sorted)-2 {
		if sorted[i] > 0 {
			break
		}
		for i > 0 && i < len(sorted)-2 && sorted[i] == sorted[i-1] {
			i++
		}
		j := i + 1
		k := len(sorted) - 1
		for j < k {
			sum := sorted[i] + sorted[j] + sorted[k]
			switch {
			case sum == 0:
				results = append(results, []int{sorted[i], sorted[j], sorted[k]})
				j++
				k--
				for j < k && sorted[j] == sorted[j-1] {
					j++
				}
			case sum < 0:
				j++
			default:
				k--
			}
		}
		i++
	}

	return results
}
