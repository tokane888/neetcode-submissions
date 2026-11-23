func isAlphaNumeric(b byte) bool {
	return ('0' <= b && b <= '9') ||
		('a' <= b && b <= 'z') ||
		('A' <= b && b <= 'Z')
}

func isPalindrome(s string) bool {
	// check numbers from right and left
	// ignore non-alphanumeric char
	// if rune on left side and right side is not same, return false
	leftIdx := 0
	rightIdx := len(s) - 1
	for leftIdx < rightIdx {
		for leftIdx < rightIdx && !isAlphaNumeric(s[leftIdx]) {
			leftIdx++
		}
		for rightIdx > leftIdx && !isAlphaNumeric(s[rightIdx]) {
			rightIdx--
		}
		if unicode.ToLower(rune(s[leftIdx])) != unicode.ToLower(rune(s[rightIdx])) {
			return false
		}
		leftIdx++
		rightIdx--
	}

	return true
}
