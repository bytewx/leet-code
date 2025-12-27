func removeElement(nums []int, val int) int {
	k := 0
	n := len(nums)
	temp := make([]int, len(nums))

	for i := 0; i < n; i++ {
        if nums[i] != val {
            temp[k] = nums[i]
            k++
        }
	}

	copy(nums, temp)

	return k
}
