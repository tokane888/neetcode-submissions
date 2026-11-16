import "slices"

func groupAnagrams(strs []string) [][]string {
	anagramList := make(map[string][]string)	// pattern => anagrams
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		anagramList[string(runes)] = append(anagramList[string(runes)], str)
	}

	result := make([][]string, 0)
	for _, anagrams := range anagramList {
		result = append(result, anagrams)
	}

	return result
}
