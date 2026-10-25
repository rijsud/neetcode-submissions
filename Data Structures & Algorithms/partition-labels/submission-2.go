func partitionLabels(s string) []int {
    found := map[byte]int{}
	for i := range s {
		found[s[i]] = i
	}

	res := []int{}
	check, size := 0, 0
	for i := range s {
		if check < found[s[i]] {
			check = found[s[i]]
		}
		size++
		if check == i {
			res = append(res, size)
			size = 0
		}
	}

	return res
}
