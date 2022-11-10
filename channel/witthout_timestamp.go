package channel

import (
	"fmt"
	"time"
)

func WithoutTimeStamp() {
	ch := make(chan int)

	go func() {
		for {
			for i := 0; ; i++ {
				ch <- i
				fmt.Println("channel input:", ch, time.Now())
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("tick:", <-ch)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
}
