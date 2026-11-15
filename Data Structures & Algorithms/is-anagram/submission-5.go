func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeCount := make(map[rune]int)
	for _, sRune := range s {
		runeCount[sRune]++
	}
	for _, tRune := range t {
		runeCount[tRune]--
	}
	
	for _, rc := range runeCount {
		if rc != 0 {
			return false
		}
	}

	return true
}
