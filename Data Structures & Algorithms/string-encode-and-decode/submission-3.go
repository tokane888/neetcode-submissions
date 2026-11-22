type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	numStartIdx := 0
	results := make([]string, 0)
	idx := 0
	for idx < len(encoded) {
		if encoded[idx] == '#' {
			num := encoded[numStartIdx:idx]
			strLen, _ := strconv.Atoi(num)
			results = append(results, encoded[idx+1:idx+1+strLen])

			idx += strLen
			numStartIdx = idx + 1
		}
		idx++
	}

	return results
}
