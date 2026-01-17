func missingNumber(nums []int) int {
    result := len(nums)

    for index, value := range nums {
        result ^= index ^ value
    }

    return result
}

// nums = [3,0,1]
// n = 3
// res = 3
// i = 0: res = 3 ^ 0 ^ 3 = 0
// i = 1: res = 0 ^ 1 ^ 0 = 1
// i = 2: res = 1 ^ 2 ^ 1 = 2
