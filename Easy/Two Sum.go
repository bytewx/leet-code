// 1-st solution

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return []int{}
}

// 2-nd solution

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numsMap[complement]; found {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return nil
}
