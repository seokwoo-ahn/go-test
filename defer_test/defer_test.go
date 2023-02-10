package defer_test

import "testing"

func TestDefer(t *testing.T) {
	num := 20
	ch := make(chan int, num)
	InsertValueWithDefer(num, &ch)

	PopChannel(&ch)
}
