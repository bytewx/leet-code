func sortedSquares(nums []int) []int {
	var result []int

	for num := range nums {
		result = append(result, nums[num]*nums[num])
	}

	sort.Ints(result)

	return result
}
