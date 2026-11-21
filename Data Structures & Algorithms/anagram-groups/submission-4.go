import "slices"

func groupAnagrams(strs []string) [][]string {
	patternLists := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		patternLists[string(runes)] = append(patternLists[string(runes)], str)
	}
	result := make([][]string, 0, len(patternLists))
	for _, stringSlice := range patternLists {
		result = append(result, stringSlice)
	}

	return result
}
