func getAverages(nums []int, k int) []int {
    if k == 0 {
        return nums
    }

    div, sum := 2 * k + 1, 0
    prefix := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        if i < k || i >= len(nums) - k {
            prefix[i] -= 1
            
            continue
        }

        if sum == 0 {
            for j := i - k; j <= i + k; j++ {
                sum += nums[j]
            }
        } else {
            sum -= nums[i - k - 1]
            sum += nums[i + k]
        }

        prefix[i] = sum / div
    }

    return prefix
}
