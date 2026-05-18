func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	countT := make(map[rune]int)
	for _, c := range t {
		countT[c]++
	}

	res := []int{-1, -1}
	resLen := int(^uint(0) >> 1)
	for i := 0; i < len(s); i++ {
		countS := make(map[rune]int)
		for j := i; j < len(s); j++ {
			countS[rune(s[j])]++

			flag := true
			for c, cnt := range countT {
				if cnt > countS[c] {
					flag = false
					break
				}
			}

			if flag && (j-i+1) < resLen {
				resLen = j - i + 1
				res = []int{i, j}
			}
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]
}
