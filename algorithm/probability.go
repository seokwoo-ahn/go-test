package algorithm

func BirthdayProblem(n int) float64 {
	var res float64
	var p float64 = 1
	for i := 1; i < n; i++ {
		p *= float64(365-i) / 365
	}
	res = 1 - p
	return res
}
