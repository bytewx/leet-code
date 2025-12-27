func canMakeArithmeticProgression(arr []int) bool {
    if len(arr) == 0 || len(arr) == 1 {
        return false
    }

    if len(arr) == 2 {
        return true
    }

    sort.Ints(arr)

    commonDifference := arr[1] - arr[0]

    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i - 1] != commonDifference {
            return false
        }
    }
    
    return true
}
