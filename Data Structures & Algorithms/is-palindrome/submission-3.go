func isAlphaNumeric(b byte) bool {
	return ('a' <= b && b <= 'z') ||
		('A' <= b && b <= 'Z') ||
		('0' <= b && b <= '9')
}

func isPalindrome(s string) bool {
	strLen := len(s)
	leftIdx := 0
	rightIdx := strLen-1
	sRune := []rune(s)
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
		left := unicode.ToLower(sRune[leftIdx])
		right := unicode.ToLower(sRune[rightIdx])
		if left != right {
			return false
		}
		leftIdx++
		rightIdx--
	}

	return true
}
