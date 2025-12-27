func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		complement := target - nums[i]

		if index, found := numsMap[complement]; found {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return nil
}
