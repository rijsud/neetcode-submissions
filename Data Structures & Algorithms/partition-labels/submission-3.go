func partitionLabels(s string) []int {
    found := map[byte]int{}
	for i := range s {
		found[s[i]] = i
	}

	res := []int{}
	check, size := 0, 0
	for i := range s {
		size++
		if found[s[i]] > check {
			check = found[s[i]]
		}
		if i == check {
			res = append(res, size)
			size = 0
		}
	}

	return res
}
