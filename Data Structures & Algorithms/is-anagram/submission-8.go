func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	seen := make(map[rune]int)
	for _, r := range s {
		seen[r]++
	}
	for _, r := range t {
		seen[r]--
	}

	for _, runeCount := range seen {
		if runeCount != 0 {
			return false
		}
	}
	return true
}
