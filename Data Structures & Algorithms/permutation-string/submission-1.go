func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	// s2Map := make(map[rune]int)

	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
	}
	// for _, letter := range s2 {
	// 	s2Map[letter]++
	// }
	// for key, val := range s1Map {
	// 	_, found := s2Map[key]
	// 	if !found
	// }

	need := len(s1Map)
	for i := 0; i < len(s2); i++ {
		s2Map := make(map[byte]int)
		cur := 0
		for j := i; j < len(s2); j++ {
			s2Map[s2[j]]++
			if s1Map[s2[j]] < s2Map[s2[j]] { break }
			if s1Map[s2[j]] == s2Map[s2[j]] { cur++ }
			if cur == need { return true }
		}
	}

	return false
}
