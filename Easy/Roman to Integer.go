func romanToInt(s string) int {
	var number int

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	byteSlice := []byte(s)

	for i := 0; i < len(byteSlice); i++ {
		val := romanMap[string(byteSlice[i])]

		if i < len(byteSlice)-1 {
			nextVal := romanMap[string(byteSlice[i+1])]

			if val < nextVal {
				number += nextVal - val

				i++

				continue
			}
		}

		number += val
	}

	return number
}
