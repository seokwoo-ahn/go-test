package algorithm

func FillKmpTable(input string) *[]int {
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
	return &res
}
