func reverseWords(s string) string {
    bytes := []byte(s)

    start := 0

    for i := 0; i <= len(bytes); i++ {
        if i == len(bytes) || bytes[i] == ' ' {
            left, right := start, i - 1

            for left < right {
                bytes[left], bytes[right] = bytes[right], bytes[left]

                left++
                right--
            }

            start = i + 1
        }
    }

    return string(bytes)
}
