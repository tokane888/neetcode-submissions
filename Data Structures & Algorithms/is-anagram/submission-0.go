import "maps"

func isAnagram(s string, t string) bool {
	charS := make(map[rune]int)
	for _, c := range s {
		charS[c]++
	}
	charT := make(map[rune]int)
	for _, c := range t {
		charT[c]++
	}

	return maps.Equal(charS, charT)
}
