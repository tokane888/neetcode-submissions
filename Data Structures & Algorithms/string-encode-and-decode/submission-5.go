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
	i := 0
	lenStartIdx := 0
	result := make([]string, 0)
	for i < len(encoded) {
		for encoded[i] != '#' {
			i++
		}
		strLen, _ := strconv.Atoi(encoded[lenStartIdx:i])
		i++
		result = append(result, encoded[i:i+strLen])
		i += strLen
		lenStartIdx = i
	}

	return result
}
