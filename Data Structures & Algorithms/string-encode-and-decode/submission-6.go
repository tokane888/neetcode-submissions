type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		// [num]#[str]
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	result := make([]string, 0)
	for i < len(encoded) {
		lenStartIdx := i
		var strLen int
		for {
			i++
			if encoded[i] == '#' {
				strLen, _ = strconv.Atoi(string(encoded[lenStartIdx:i]))
				break
			}
		}
		i++
		str := encoded[i:i+strLen]
		result = append(result, str)
		i = i + strLen
	}

	return result
}
