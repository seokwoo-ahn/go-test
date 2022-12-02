package channel

import "fmt"

type Channels struct {
	input chan int
	term  chan int
}

func (c *Channels) Select1() {
	select {
	case c.input <- 1:
		fmt.Println("ch1 inserted")
	case <-c.term:
		fmt.Println("dead")
	}
}

func (c *Channels) Select2() {
	select {
	case <-c.term:
		fmt.Println("dead")
	case c.input <- 1:
		fmt.Println("ch1 inserted")
	}
}

func (c *Channels) Terminate() {
	close(c.term)
}
