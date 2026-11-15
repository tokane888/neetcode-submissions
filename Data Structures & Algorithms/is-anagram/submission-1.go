func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCounter := make(map[rune]int)
	for _, c := range s {
		charCounter[c]++
	}
	for _, c := range t {
		charCounter[c]--
	}
	
	for _, count := range charCounter {
		if count != 0 {
			return false
		}
	}
	return true
}
