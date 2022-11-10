package channel

import "fmt"

func Select() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int)

	select {
	case ch1 <- 1:
		fmt.Println("ch1 inserted")
	case <-ch2:
		fmt.Println("dead")
	}

	fmt.Println(<-ch1)
}
