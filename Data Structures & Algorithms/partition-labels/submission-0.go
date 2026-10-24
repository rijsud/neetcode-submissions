func partitionLabels(s string) []int {
    found := map[byte]int{}
	for i := range s {
		found[s[i]] = i
	}

	res := []int{}
	size := 0
	check := 0
	for i := range s {
		if check == 0 {
			check = found[s[i]]
			size = 0
		}

		size++

		if i == check {
			check = 0
			res = append(res, size)
			continue
		}

		if found[s[i]] > check {
			check = found[s[i]]
		}
	}

	return res
}
