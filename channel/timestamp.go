package channel

import (
	"fmt"
	"time"
)

func WithTimeStamp() {
	ch := make(chan int)

	go func() {
		for {
			for i := 0; ; i++ {
				ch <- i
				fmt.Println("channel input:", ch, time.Now())
			}
		}
	}()

	//ticker에서 t를 가져오는 시점으로 인해 print문의 출력이 어긋남

	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick:", <-ch, t)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
}

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
