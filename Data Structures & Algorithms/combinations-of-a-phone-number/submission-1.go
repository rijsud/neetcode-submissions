func letterCombinations(digits string) []string {
	res := []string{}
	subset := []string{}
	n := len(digits)
	if digits == "" {return res}

	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, strings.Join(subset, ""))
			return
		}
		
		digitSet := digitsMap(digits[i])
		for j := 0; j < len(digitSet); j++ {
			subset = append(subset, digitSet[j])
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0)
	return res
}

func digitsMap(i byte) []string {
	switch i {
		case '2':
		return []string{"a", "b", "c"}
		case '3':
		return []string{"d", "e", "f"}
		case '4':
		return []string{"g", "h", "i"}
		case '5':
		return []string{"j", "k", "l"}
		case '6':
		return []string{"m", "n", "o"}
		case '7':
		return []string{"p", "q", "r", "s"}
		case '8':
		return []string{"t", "u", "v"}
		case '9':
		return []string{"w", "x", "y", "z"}
		default:
		return []string{}
	}
}