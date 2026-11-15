import "slices"

func groupAnagrams(strs []string) [][]string {
	anagramLists := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		anagram := string(runes)
		if anagramLists[anagram] == nil {
			anagramLists[anagram] = make([]string, 0)
		}
		anagramLists[anagram] = append(anagramLists[anagram], str)
	}

	result := make([][]string, 0, len(anagramLists))
	for _, anagramList := range anagramLists {
		result = append(result, anagramList)
	}

	return result
}
