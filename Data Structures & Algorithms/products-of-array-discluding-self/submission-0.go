func productExceptSelf(nums []int) []int {
    prefixProd := make([]int, len(nums))
    suffixProd := make([]int, len(nums))
    output := make([]int, len(nums))

    prefixProd[0] = 1
    suffixProd[len(nums)-1] = 1

    product := 1
    index := 0
    for i := 1; i < len(nums); i++ {
        prefixProd[i] = product * nums[index]
        product = prefixProd[i]
        index++
    }

    product = 1
    index = len(nums) - 1
    for i := len(nums) - 2; i >= 0; i-- {
        suffixProd[i] = product * nums[index]
        product = suffixProd[i]
        index--
    }

    for i := 0; i < len(nums); i++ {
        output[i] = prefixProd[i] * suffixProd[i]
    }

    return output
}
