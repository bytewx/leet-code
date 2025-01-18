// 1-st solution
// Time complexity - O(log(n))
// Space complexity - O(log(n))

func isPalindrome(x int) bool {
    numberString := strconv.Itoa(x)
    stringRunes := []rune(numberString)

    for i, j := 0, len(stringRunes) - 1; i < len(stringRunes) / 2; i, j = i + 1, j - 1 {
        stringRunes[i], stringRunes[j] = stringRunes[j], stringRunes[i]
    }

    reversedNumber := string(stringRunes)

    return reversedNumber == numberString
}   

// 2-nd solution
// Time complexity - O(log(n))
// Space complexity - O(1)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    if x != 0 && x % 10 == 0 {
        return false
    }

    reversed := 0
    original := x

    for x > 0 {
        reversed = reversed * 10 + x % 10
        x = x / 10
    }

    return original == reversed
}
