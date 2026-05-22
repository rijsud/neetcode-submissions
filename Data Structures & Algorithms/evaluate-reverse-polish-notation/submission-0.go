func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}
		if token == "+" {
			stack[len(stack)-2] += stack[len(stack)-1]
		}
		if token == "-" {
			stack[len(stack)-2] -= stack[len(stack)-1]
		}
		if token == "*" {
			stack[len(stack)-2] *= stack[len(stack)-1]
		}
		if token == "/" {
			if stack[len(stack)-1] == 0 {
				stack[len(stack)-2] = 0
			} else {
				stack[len(stack)-2] /= stack[len(stack)-1]
			}
		}
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}
