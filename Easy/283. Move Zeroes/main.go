func moveZeroes(nums []int) {    
    if len(nums) == 0 || len(nums) == 1 {
        return
    }

    lastNonZero := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZero] = nums[i]

            lastNonZero++
        }
    }

    for i := lastNonZero; i < len(nums); i++ {
        nums[i] = 0
    }
}
