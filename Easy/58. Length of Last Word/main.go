func lengthOfLastWord(s string) int {
    words := strings.Fields(strings.TrimSpace(s))
    lastWord := words[len(words) - 1]

    return len(lastWord)
}
