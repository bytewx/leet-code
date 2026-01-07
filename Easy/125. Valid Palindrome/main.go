func isPalindrome(s string) bool {
    i, j := 0, len(s)-1

    for i < j {
        for i < j && !isAlnum(s[i]) {
            i++
        }

        for i < j && !isAlnum(s[j]) {
            j--
        }

        if i < j {
            if toLowerASCII(s[i]) != toLowerASCII(s[j]) {
                return false
            }
            i++
            j--
        }
    }
    return true
}

func isAlnum(b byte) bool {
    return (b >= 'a' && b <= 'z') ||
        (b >= 'A' && b <= 'Z') ||
        (b >= '0' && b <= '9')
}

func toLowerASCII(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        return b + ('a' - 'A')
    }
    return b
}
