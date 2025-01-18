// 1-st solution
// Time complexity - O(n^2)
// Space complexity - O(n)

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
// Time complexity - O(n)
// Space complexity - O(n)

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
