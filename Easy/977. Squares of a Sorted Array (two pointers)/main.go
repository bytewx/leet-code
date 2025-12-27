func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))
    left := 0
    right := len(nums) - 1
    position := len(nums) - 1

    for left <= right {
        if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
            result[position] = nums[right] * nums[right]
            right--
        } else {
            result[position] = nums[left] * nums[left]
            left++
        }

        position--
    }
    
    return result
}
