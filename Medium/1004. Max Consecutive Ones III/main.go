func longestOnes(nums []int, k int) int {
    left := 0
    ans := 0

    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            k--
        }

        for k < 0 {
            if nums[left] == 0 {
                k++
            }

            left++
        }

        if right - left + 1 > ans {
            ans = right - left + 1
        }
    }

    return ans
}
