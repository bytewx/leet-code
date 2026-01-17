func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }

    var seen [26]bool
    count := 0

    for i := 0; i < len(sentence); i++ {
        idx := sentence[i] - 'a'
        if !seen[idx] {
            seen[idx] = true
            count++
            
            if count == 26 {
                return true
            }
        }
    }

    return false
}
