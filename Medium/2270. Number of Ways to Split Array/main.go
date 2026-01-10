func waysToSplitArray(nums []int) int {
    prefix := []int{nums[0]}
    ans := 0
    
    for i := 1; i < len(nums); i++ {
        prefix = append(prefix, prefix[len(prefix) - 1] + nums[i])
    }

    for i := 0; i < len(nums) - 1; i++ {
        leftSection := prefix[i]
        rightSection := prefix[len(prefix) - 1] - prefix[i]

        if leftSection >= rightSection {
            ans++
        }
    }

    return ans
}
