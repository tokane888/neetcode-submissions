import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRune := []rune(s)
	slices.Sort(sRune)
	tRune := []rune(t)
	slices.Sort(tRune)

	return slices.Equal(sRune, tRune)
}
