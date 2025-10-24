func finalValueAfterOperations(operations []string) int {
	var x int

	for _, operation := range operations {
		if operation[0] == '-' || operation[2] == '-' {
			x--
		} else {
			x++
		}
	}

	return x
}
