func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	output := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			output[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return output
}
