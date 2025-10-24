func removeDuplicates(nums []int) int {
	var (
		fast int
		k    int
	)
	l := len(nums)
	temp := make([]int, len(nums))

	for i := 0; i < l; i++ {
		if i < l-1 {
			fast = nums[i+1]
		}

		if nums[i] == nums[fast] {
			k++
			temp[k] = nums[i]
		} else {
			temp[k] = nums[i]
		}

		i++
	}

	return k
}
