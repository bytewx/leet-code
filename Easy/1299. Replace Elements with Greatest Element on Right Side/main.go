func replaceElements(arr []int) []int {
    max := -1

    for i := len(arr)-1; i >= 0; i-- {
        newMax := max
        if arr[i] > max {
            max = arr[i]
        }
        arr[i] = newMax
    }

    return arr
}
