import "slices"

func groupAnagrams(strs []string) [][]string {
	patterns := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		patterns[string(runes)] = append(patterns[string(runes)], str)
	}
	result := make([][]string, 0, len(patterns))
	for _, strList := range patterns {
		result = append(result, strList)
	}

	return result
}
