func isAlphaNumeric(b byte) bool {
	return ('a' <= b && b <= 'z') ||
		('A' <= b && b <= 'Z') ||
		('0' <= b && b <= '9')
}

func isPalindrome(s string) bool {
	strLen := len(s)
	leftIdx := 0
	rightIdx := strLen-1
	for {
		for leftIdx < strLen && !isAlphaNumeric(s[leftIdx]) {
			leftIdx++
		}
		for rightIdx >= 0 && !isAlphaNumeric(s[rightIdx]) {
			rightIdx--
		}
		if leftIdx >= rightIdx  {
			break
		}
		left := unicode.ToLower(rune(s[leftIdx]))
		right := unicode.ToLower(rune(s[rightIdx]))
		if left != right {
			return false
		}
		leftIdx++
		rightIdx--
	}

	return true
}
