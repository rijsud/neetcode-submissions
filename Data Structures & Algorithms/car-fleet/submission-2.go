func carFleet(target int, position []int, speed []int) int {
	stack := []float64{}
	cars := [][]int{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, []int{position[i], speed[i]})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	for i := 0; i < len(position); i++ {
		time := float64(target - cars[i][0]) / float64(cars[i][1])
		stack = append(stack, time)
		if len(stack) > 1 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}
