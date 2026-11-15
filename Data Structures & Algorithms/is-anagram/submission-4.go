import "maps"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	setS := make(map[rune]int)
	setT := make(map[rune]int)
	for _, charS := range s {
		setS[charS]++
	}
	for _, charT := range t {
		setT[charT]++
	}

	return maps.Equal(setS, setT)	
}
