package algorithm

import (
	"testing"
)

func TestBirthdayProblem(t *testing.T) {
	for i := 1; i < 100; i++ {
		if BirthdayProblem(i) > 0.5 {
			t.Logf("n = %d, p = %f\n", i, BirthdayProblem(i))
			break
		}
	}
}
