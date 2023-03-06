package algorithm

func FillKmpTable(input string) []int {
	res := make([]int, len(input))
	res[0] = 0

	index := 0

	for i := 1; i < len(input); i++ {
		for index > 0 && input[index] != input[i] {
			index = res[index-1]
		}

		if input[index] == input[i] {
			index++
			res[i] = index
		}
	}
	return res
}

func KMP(input, pattern string) int {
	res := 0
	index := 0

	table := FillKmpTable(pattern)

	for i := 0; i < len(input); i++ {
		for index > 0 && input[i] != pattern[index] {
			index = table[index-1]
		}
		if input[i] == pattern[index] {
			if index == len(pattern)-1 {
				return i - len(pattern) + 1
			} else {
				index++
			}
		}
	}
	return res
}
