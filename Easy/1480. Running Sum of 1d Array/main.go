func runningSum(nums []int) []int {
    prefix := []int{nums[0]}

    for i := 1; i < len(nums); i++ {
        prefix = append(prefix, prefix[len(prefix) - 1] + nums[i]) 
    }

    return prefix
}
