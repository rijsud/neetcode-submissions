func productExceptSelf(nums []int) []int {
    output := make([]int, len(nums))

    product := 1

    for index, number := range (nums) {
        output[index] = product
        product *= number
    }

    product = 1

    for i := len(nums) - 1; i > -1; i-- {
        output[i] *= product
        product *= nums[i]
    }

    return output
}
