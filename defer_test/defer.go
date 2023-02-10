package defer_test

import "fmt"

func InsertValue(val int, ch *chan int) {
	*ch <- val
}

func PopChannel(ch *chan int) {
	for len(*ch) != 0 {
		fmt.Println(<-*ch)
	}
}

func InsertValueWithDefer(num int, ch *chan int) {
	for i := 0; i < num; i++ {
		defer InsertValue(i, ch)
	}
}
