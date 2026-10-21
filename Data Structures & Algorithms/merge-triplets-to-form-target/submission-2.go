func mergeTriplets(triplets [][]int, target []int) bool {
	check := []int{0, 0, 0}
	for _, t := range triplets {
		if t[0] > target[0] ||
		t[1] > target[1] ||
		t[2] > target[2] { continue }
		check[0] = max(t[0], check[0])
		check[1] = max(t[1], check[1])
		check[2] = max(t[2], check[2])
	}
	
	for i := range target {
		if check[i] != target[i] {return false}
	}
	return true
}
