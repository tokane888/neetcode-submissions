func isAnagram(s string, t string) bool {
	count := [26]int{}	// count each alphabet appeared in string([0] => 'a')
	for i := 0; i < len(s); i++ {
		count[int(s[i] - 'a')]++
	}
	for i := 0; i < len(t); i++ {
		count[int(t[i] - 'a')]--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}

	return true
}
