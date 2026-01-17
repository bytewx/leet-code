func countElements(arr []int) int {
    seen := make(map[int]struct{}, len(arr))

    for _, value := range arr {
        seen[value] = struct{}{} 
    } 

    count := 0

    for _, value := range arr {
        if _, ok := seen[value+1]; ok {
            count++
        }
    } 

    return count
}
