func twoSum(numbers []int, target int) []int {
	// mid := len(numbers) / 2
	// start, end := 0, len(numbers) - 1
	// for {
	// 	if numbers[mid] < target {
	// 		end = mid
	// 	} else if numbers[mid] > target {
	// 		start = mids
	// 	} else {
	// 		break
	// 	}
	// 	mid = (start + end) / 2
	// }

	// if numbers[0]

	start, end := 0, len(numbers) - 1
	for {
		if numbers[start] + numbers[end] > target {
			end--
		}
		if numbers[start] + numbers[end] < target {
			start++
		}
		if numbers[start] + numbers[end] == target {
			break
		}
	}
	return []int{start + 1, end + 1}
}
