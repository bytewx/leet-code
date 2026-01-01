func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {
        return 0
    }

    left := 0
    right := 0
    curr := 1
    ans := 0

    for ; right < len(nums); right++ {
        curr *= nums[right]

        for curr >= k {
            curr /= nums[left]
            left++
        }

        ans += right - left + 1
    }

    return ans
}   
