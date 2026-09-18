func manacher (s string) []int {
	t := "#"
	for _, char := range s { t += string(char) + "#" }
	n := len(t)
	p := make([]int, n)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		if i < r {
			p[i] = min(r-i, p[l+r-i])
		}
		for i + p[i] + 1 < n && i - p[i] - 1 > -1 && t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}
		if i + p[i] > r {
			l, r = i - p[i], i + p[i]
		}
	}
	return p
}

func countSubstrings(s string) int {
    p := manacher(s)
	res := 0
	for _, val := range p {
		res += (val + 1) / 2
	}
	return res
}
