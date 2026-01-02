func repeatedNTimes(nums []int) int {
    m := make(map[int]int)

    for _, num := range nums {
        m[num]++

        if m[num] > 1 {
            return num
        }
    }

    return -1
}
