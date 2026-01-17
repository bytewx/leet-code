func repeatedCharacter(s string) byte {
    var seen [26]bool

    for i := 0; i < len(s); i++ {
        idx := s[i] - 'a'
        if seen[idx] {
            return s[i]
        } 
        seen[idx] = true
    }

    return 0
}
