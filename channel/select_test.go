package channel

import (
	"fmt"
	"testing"
)

func TestSelect1(t *testing.T) {
	channels := &Channels{
		input: make(chan int, 2),
		term:  make(chan int),
	}
	channels.Terminate()
	channels.Select1()
}

func TestSelect2(t *testing.T) {
	channels := &Channels{
		input: make(chan int, 2),
		term:  make(chan int),
	}
	channels.Terminate()
	channels.Select2()
}

// select 문의 case는 완전 랜덤하게 실행됨
func TestSelectLoop(t *testing.T) {
	fmt.Println(">>>>>Test Select 1 Start")
	for i := 0; i < 20; i++ {
		TestSelect1(t)
	}
	fmt.Println(">>>>>Test Select 2 Start")
	for i := 0; i < 20; i++ {
		TestSelect2(t)
	}
}
